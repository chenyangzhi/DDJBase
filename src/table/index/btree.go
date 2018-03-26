package index

import (
	"common"
	"container/list"
	"fmt"
	"github.com/edsrzf/mmap-go"
	"io"
	"os"
	"sort"
	"strings"
	"sync/atomic"
	"time"
	"until"
	logger "until/xlog4go"
)

type Item interface {
	Less(than Item) bool
}

type items []Item

const (
	DefaultFreeListSize = 32
)

var (
	nilItems    = make(items, 16)
	nilChildren = make(children, 16)
)

type FreeList struct {
	freelist []*node
}

func NewFreeList(size int) *FreeList {
	return &FreeList{freelist: make([]*node, 0, size)}
}

func (f *FreeList) newNode() (n *node) {
	index := len(f.freelist) - 1
	if index < 0 {
		return new(node)
	}
	n = f.freelist[index]
	f.freelist[index] = nil
	f.freelist = f.freelist[:index]
	return
}

func (f *FreeList) freeNode(n *node) {
	if len(f.freelist) < cap(f.freelist) {
		f.freelist = append(f.freelist, n)
	}
}

type ItemIterator func(i Item) bool

func New(degree int, path string) *BTree {
	return NewWithFreeList(degree, NewFreeList(DefaultFreeListSize), path)
}

func RemoveFromLruCache(k uint32, cow *copyOnWriteContext) {
	node := cow.nodeIdMap[k]
	if len(node.children) != 0 {
		return
	}
	page := node.NodeToPage()
	page.WriteToMMapRegion(cow)
	d := cow.dirtyPage
	d.Delete(k)
	cow.freeNode(node)
}

func AddToLruCache(k uint32, cow *copyOnWriteContext) {
	if true {
		return
	}
	lru := cow.cache
	node := cow.nodeIdMap[k]
	if len(node.children) != 0 {
		return
	}
	if lru.Len() == 50 {
		for i := 0; i < 25; i++ {
			k, _, _ := lru.RemoveOldest()
			RemoveFromLruCache(k.(uint32), cow)
		}
	} else {
		lru.Add(k, nil)
	}
}

func NewWithFreeList(degree int, f *FreeList, path string) *BTree {
	if degree <= 1 {
		panic("bad degree")
	}
	cow := copyOnWriteContext{freelist: f}
	fs, err := os.OpenFile(path, os.O_RDWR, 0666)
	cow.f = fs
	common.Check(err)
	cow.mmapmap = make(map[uint32]mmap.MMap)
	cow.nodeIdMap = make(map[uint32]*node)
	cow.dirtyPage = make(common.Set)
	cow.mtPage, cow.metaMmap = GetMetaPage(fs)
	cow.emptyList = cow.mtPage.GetEmptyList()
	cow.NewPage = make(common.Set)
	cow.freelist = f
	cow.cache, _ = until.NewLRU(50, nil)
	return &BTree{
		degree: degree,
		cow:    &cow,
	}
}

func (s *items) insertAt(index int, item Item) {
	*s = append(*s, nil)
	if index < len(*s) {
		copy((*s)[index+1:], (*s)[index:])
	}
	(*s)[index] = item
}

func (s *items) removeAt(index int) Item {
	item := (*s)[index]
	copy((*s)[index:], (*s)[index+1:])
	(*s)[len(*s)-1] = nil
	*s = (*s)[:len(*s)-1]
	return item
}

func (s *items) pop() (out Item) {
	index := len(*s) - 1
	out = (*s)[index]
	(*s)[index] = nil
	*s = (*s)[:index]
	return
}

func (s *items) truncate(index int) {
	var toClear items
	*s, toClear = (*s)[:index], (*s)[index:]
	for len(toClear) > 0 {
		toClear = toClear[copy(toClear, nilItems):]
	}
}

func (s items) find(item Item) (index int, found bool) {
	i := sort.Search(len(s), func(i int) bool {
		return item.Less(s[i])
	})
	if i > 0 && !s[i-1].Less(item) {
		return i - 1, true
	}
	return i, false
}

type childrenNode struct {
	childNode   *node
	childNodeId uint32
}
type children []*childrenNode

func (s *children) insertAt(index int, n *node) {
	cn := childrenNode{
		childNode:   n,
		childNodeId: n.nodeId,
	}
	*s = append(*s, nil)
	if index < len(*s) {
		copy((*s)[index+1:], (*s)[index:])
	}
	(*s)[index] = &cn
}

func (s *children) removeAt(index int) *childrenNode {
	n := (*s)[index]
	copy((*s)[index:], (*s)[index+1:])
	(*s)[len(*s)-1] = nil
	*s = (*s)[:len(*s)-1]
	return n
}

func (s *children) pop() (out *node) {
	index := len(*s) - 1
	out = (*s)[index].childNode
	(*s)[index] = nil
	*s = (*s)[:index]
	return
}

func (s *children) truncate(index int) {
	var toClear children
	*s, toClear = (*s)[:index], (*s)[index:]
	for len(toClear) > 0 {
		toClear = toClear[copy(toClear, nilChildren):]
	}
}

type node struct {
	nodeId   uint32
	items    items
	children children
	cow      *copyOnWriteContext
}

func (n *node) getReadableChild(i int) *node {
	if n.children[i].childNode == nil || len(n.children[i].childNode.items) == 0 {
		n.children[i].childNode = GetBTreeNodeById(n.children[i].childNodeId, n.cow)
		return n.children[i].childNode
	}
	return n.children[i].childNode
}

func (n *node) split(i int) (Item, *node) {
	AddToLruCache(n.nodeId, n.cow)
	item := n.items[i]
	next := n.cow.newNode()
	next.items = append(next.items, n.items[i+1:]...)
	n.items.truncate(i)
	if len(n.children) > 0 {
		next.children = append(next.children, n.children[i+1:]...)
		n.children.truncate(i + 1)
	}
	return item, next
}

func (n *node) maybeSplitChild(i, maxItems int) bool {
	AddToLruCache(n.nodeId, n.cow)
	if len(n.getReadableChild(i).items) < maxItems {
		return false
	}
	first := n.getReadableChild(i)
	item, second := first.split(maxItems / 2)
	n.items.insertAt(i, item)
	n.children.insertAt(i+1, second)
	n.cow.dirtyPage.Insert(n.nodeId)
	n.cow.dirtyPage.Insert(first.nodeId)
	n.cow.dirtyPage.Insert(second.nodeId)
	return true
}

func (n *node) insert(item Item, maxItems int) Item {
	i, found := n.items.find(item)
	AddToLruCache(n.nodeId, n.cow)
	if found {
		out := n.items[i]
		n.items[i] = item
		n.cow.dirtyPage.Insert(n.nodeId)
		return out
	}
	if len(n.children) == 0 {
		n.items.insertAt(i, item)
		n.cow.dirtyPage.Insert(n.nodeId)
		return nil
	}
	if n.maybeSplitChild(i, maxItems) {
		inTree := n.items[i]
		switch {
		case item.Less(inTree):
		case inTree.Less(item):
			i++
		default:
			out := n.items[i]
			n.items[i] = item
			return out
		}
	}
	return n.getReadableChild(i).insert(item, maxItems)
}

func (n *node) get(key Item) Item {
	AddToLruCache(n.nodeId, n.cow)
	i, found := n.items.find(key)
	if found {
		return n.items[i]
	} else if len(n.children) > 0 {
		return n.getReadableChild(i).get(key)
	}
	return nil
}

func min(n *node) Item {
	if n == nil {
		return nil
	}
	for len(n.children) > 0 {
		n = n.children[0].childNode
	}
	if len(n.items) == 0 {
		return nil
	}
	return n.items[0]
}

func max(n *node) Item {
	if n == nil {
		return nil
	}
	for len(n.children) > 0 {
		n = n.children[len(n.children)-1].childNode
	}
	if len(n.items) == 0 {
		return nil
	}
	return n.items[len(n.items)-1]
}

type toRemove int

const (
	removeItem toRemove = iota
	removeMin
	removeMax
)

func (n *node) remove(item Item, minItems int, typ toRemove) Item {
	var i int
	var found bool
	AddToLruCache(n.nodeId, n.cow)
	switch typ {
	case removeMax:
		if len(n.children) == 0 {
			n.cow.dirtyPage.Insert(n.nodeId)
			return n.items.pop()
		}
		i = len(n.items)
	case removeMin:
		if len(n.children) == 0 {
			n.cow.dirtyPage.Insert(n.nodeId)
			return n.items.removeAt(0)
		}
		i = 0
	case removeItem:
		i, found = n.items.find(item)
		if len(n.children) == 0 {
			if found {
				n.cow.dirtyPage.Insert(n.nodeId)
				return n.items.removeAt(i)
			}
			return nil
		}
	default:
		panic("invalid type")
	}
	if len(n.getReadableChild(i).items) <= minItems {
		return n.growChildAndRemove(i, item, minItems, typ)
	}
	child := n.getReadableChild(i)
	if found {
		out := n.items[i]
		n.cow.dirtyPage.Insert(n.nodeId)
		n.items[i] = child.remove(nil, minItems, removeMax)
		return out
	}
	return child.remove(item, minItems, typ)
}
func (n *node) growChildAndRemove(i int, item Item, minItems int, typ toRemove) Item {
	AddToLruCache(n.nodeId, n.cow)
	if i > 0 && len(n.getReadableChild(i).items) > minItems {
		child := n.getReadableChild(i)
		stealFrom := n.getReadableChild(i - 1)
		stolenItem := stealFrom.items.pop()
		child.items.insertAt(0, n.items[i-1])
		n.items[i-1] = stolenItem
		if len(stealFrom.children) > 0 {
			child.children.insertAt(0, stealFrom.children.pop())
		}
		n.cow.dirtyPage.Insert(n.nodeId)
		n.cow.dirtyPage.Insert(child.nodeId)
		n.cow.dirtyPage.Insert(stealFrom.nodeId)
	} else if i < len(n.items) && len(n.getReadableChild(i+1).items) > minItems {
		child := n.getReadableChild(i)
		stealFrom := n.getReadableChild(i + 1)
		stolenItem := stealFrom.items.removeAt(0)
		child.items = append(child.items, n.items[i])
		n.items[i] = stolenItem
		if len(stealFrom.children) > 0 {
			child.children = append(child.children, stealFrom.children.removeAt(0))
		}
		n.cow.dirtyPage.Insert(n.nodeId)
		n.cow.dirtyPage.Insert(child.nodeId)
		n.cow.dirtyPage.Insert(stealFrom.nodeId)
	} else {
		if i >= len(n.items) {
			i--
		}
		child := n.getReadableChild(i)
		mergeItem := n.items.removeAt(i)
		mergeChild := n.children.removeAt(i + 1)
		child.items = append(child.items, mergeItem)
		child.items = append(child.items, mergeChild.childNode.items...)
		child.children = append(child.children, mergeChild.childNode.children...)
		n.cow.dirtyPage.Insert(n.nodeId)
		n.cow.dirtyPage.Insert(child.nodeId)
		n.cow.freeNode(mergeChild.childNode)
	}
	return n.remove(item, minItems, typ)
}

type direction int

const (
	descend = direction(-1)
	ascend  = direction(+1)
)

func (n *node) iterate(dir direction, start, stop Item, includeStart bool, hit bool, iter ItemIterator) (bool, bool) {
	var ok bool
	AddToLruCache(n.nodeId, n.cow)
	switch dir {
	case ascend:
		for i := 0; i < len(n.items); i++ {
			if start != nil && n.items[i].Less(start) {
				continue
			}
			if len(n.children) > 0 {
				if hit, ok = n.getReadableChild(i).iterate(dir, start, stop, includeStart, hit, iter); !ok {
					return hit, false
				}
			}
			if !includeStart && !hit && start != nil && !start.Less(n.items[i]) {
				hit = true
				continue
			}
			hit = true
			if stop != nil && !n.items[i].Less(stop) {
				return hit, false
			}
			if !iter(n.items[i]) {
				return hit, false
			}
		}
		if len(n.children) > 0 {
			if hit, ok = n.getReadableChild(len(n.children)-1).iterate(dir, start, stop, includeStart, hit, iter); !ok {
				return hit, false
			}
		}
	case descend:
		for i := len(n.items) - 1; i >= 0; i-- {
			if start != nil && !n.items[i].Less(start) {
				if !includeStart || hit || start.Less(n.items[i]) {
					continue
				}
			}
			if len(n.children) > 0 {
				if hit, ok = n.getReadableChild(i+1).iterate(dir, start, stop, includeStart, hit, iter); !ok {
					return hit, false
				}
			}
			if stop != nil && !stop.Less(n.items[i]) {
				return hit, false
			}
			hit = true
			if !iter(n.items[i]) {
				return hit, false
			}
		}
		if len(n.children) > 0 {
			if hit, ok = n.getReadableChild(0).iterate(dir, start, stop, includeStart, hit, iter); !ok {
				return hit, false
			}
		}
	}
	return hit, true
}

func (n *node) Print(w io.Writer, level int) {
	fmt.Fprintf(w, "%sNODE:%s %d %s %d ", strings.Repeat("  ", level), "nodeid:", n.nodeId, "the item length : ", len(n.items))
	for _, o := range n.items {
		fmt.Fprintf(w, "{%d}", o.(*BtreeNodeItem).IdxId)
	}
	fmt.Fprintf(w, "\n")
	for i := range n.children {
		node := n.getReadableChild(i)
		node.Print(w, level+1)
	}
}

type BTree struct {
	degree int
	length int
	root   *node
	cow    *copyOnWriteContext
}

type copyOnWriteContext struct {
	freelist      *FreeList
	f             *os.File
	indexFilePath string
	dataFile      *os.File
	dataFilePath  string
	emptyList     []uint32
	NewPage       common.Set
	mmapmap       map[uint32]mmap.MMap
	nodeIdMap     map[uint32]*node
	mtPage        *MetaPage
	metaMmap      mmap.MMap
	dirtyPage     common.Set
	cache         *until.LRU
	curVPool      *ValuePool
	curFileOffset uint64
	fullQueue     *list.List
	flushQueue    *list.List
	flushFinish   bool
	emptyQueue    chan *ValuePool
	readRef       int64
}

func (t *BTree) Clone() (t2 *BTree) {
	cow1, cow2 := *t.cow, *t.cow
	out := *t
	t.cow = &cow1
	out.cow = &cow2
	return &out
}

func (t *BTree) maxItems() int {
	return t.degree*2 - 1
}

func (t *BTree) minItems() int {
	return t.degree - 1
}

func (c *copyOnWriteContext) newNode() (n *node) {
	n = c.freelist.newNode()
	n.cow = c
	n.nodeId = c.emptyList[len(c.emptyList)-1]
	c.NewPage.Insert(n.nodeId)
	_assert(len(c.emptyList) != 0, "the emptylist length is zero")
	c.emptyList = c.emptyList[0 : len(c.emptyList)-1]
	c.nodeIdMap[n.nodeId] = n
	return
}

func (c *copyOnWriteContext) freeNode(n *node) {
	if n.cow == c {
		c.emptyList = append(c.emptyList, n.nodeId)
		delete(c.NewPage, n.nodeId)
		n.items.truncate(0)
		n.children.truncate(0)
		n.cow = nil
		c.freelist.freeNode(n)
		delete(c.nodeIdMap, n.nodeId)
	}
}

func freeBtree(btr *BTree) {
	readRef := atomic.LoadInt64(&btr.cow.readRef)
	for {
		if readRef == 0 {
			logger.Info("Vaccum is finished and the index file %v and the data file %v is removed",
				btr.cow.indexFilePath, btr.cow.dataFilePath)
			btr.cow.dataFile.Close()
			os.Remove(btr.cow.dataFilePath)
			btr.cow.f.Close()
			os.Remove(btr.cow.indexFilePath)
			break
		}
		time.Sleep(100 * time.Millisecond)
	}
}

func Insert(key uint64, value []byte) *BtreeNodeItem {
	var bItem *BtreeNodeItem
	if VaccumHolder.VaccumFlag == true {
		VaccumHolder.VMutex.Lock()
		bItem = Curtr.InternalInsert(key, value)
		VaccumHolder.VMutex.Unlock()
	} else {
		bItem = Curtr.InternalInsert(key, value)
	}
	return bItem
}

func (tr *BTree) InternalInsert(key uint64, value []byte) *BtreeNodeItem {
	var out *BtreeNodeItem = nil
	cow := tr.cow
	offset := cow.curVPool.CurrentOffest
	fileOffset := cow.curVPool.EofOffset
	bItem := NewBtreeNodeItem(fileOffset+offset, key, 1, uint32(len(value)))
	result := tr.ReplaceOrInsert(bItem)
	iv := NewInternalValue(false, uint32(len(value)), key, value)
	bv := iv.ToBytes()
	lv := len(bv)
	currentOffset := offset + uint64(lv)
	if result != nil {
		cow.mtPage.TotalRemoved = cow.mtPage.TotalRemoved + uint64(result.(*BtreeNodeItem).KeyLength) + 15
		cow.curVPool.RemoveValues.PushBack(result)
		out = result.(*BtreeNodeItem)
	}
	if currentOffset < MAXPOOLSIZE {
		copy(cow.curVPool.MemeryTail[offset:currentOffset], bv)
		cow.curVPool.CurrentOffest = currentOffset
	} else {
		//the memery is full, so need to flush
		cow.fullQueue.PushBack(cow.curVPool)
		if cow.flushFinish == true && cow.fullQueue.Len() > MEMQUEUESIZE/2 {
			cow.flushQueue = cow.fullQueue
			cow.fullQueue = list.New()
			cow.flushFinish = false
			go Flush(cow)
		}
		cow.curVPool = <-cow.emptyQueue
		cow.curVPool.EofOffset = offset + fileOffset
		currentOffset = uint64(lv)
		copy(cow.curVPool.MemeryTail[0:currentOffset], bv)
		cow.curVPool.CurrentOffest = currentOffset
		// need vaccum
		if tr.cow.mtPage.TotalRemoved*2 > tr.cow.curFileOffset && tr.cow.curFileOffset > GCMINFILESIZE {
			if VaccumHolder.VaccumFlag == false {
				VaccumHolder.VaccumFlag = true
				go Vaccum(cow)
			}
		}
	}
	return out
}

func (t *BTree) ReplaceOrInsert(item Item) Item {
	if item == nil {
		panic("nil item add to BTree")
	}
	if t.root == nil {
		t.root = t.cow.newNode()
		t.root.items = append(t.root.items, item)
		t.length++
		return nil
	} else {
		if len(t.root.items) >= t.maxItems() {
			item2, second := t.root.split(t.maxItems() / 2)
			oldroot := t.root
			t.cow.dirtyPage.Insert(oldroot.nodeId)
			t.root = t.cow.newNode()
			t.cow.mtPage.RootId = t.root.nodeId
			t.root.items = append(t.root.items, item2)
			old := &childrenNode{childNode: oldroot, childNodeId: oldroot.nodeId}
			sec := &childrenNode{childNode: second, childNodeId: second.nodeId}
			t.root.children = append(t.root.children, old, sec)
		}
	}
	out := t.root.insert(item, t.maxItems())
	if out == nil {
		t.length++
	}
	return out
}

func Delete(key uint64) []byte {
	var bItem BtreeNodeItem
	var bs []byte
	bItem.IdxId = key
	if VaccumHolder.VaccumFlag == true {
		VaccumHolder.VMutex.Lock()
		bs = Curtr.InternalDelete(bItem)
		VaccumHolder.VMutex.Unlock()
		VaccumHolder.VaccumFlag = false
	} else {
		bs = Curtr.InternalDelete(bItem)
	}
	return bs
}

func (t *BTree) InternalDelete(item Item) []byte {
	r := t.deleteItem(item, removeItem)
	t.cow.curVPool.RemoveValues.PushBack(r)
	t.cow.mtPage.TotalRemoved = t.cow.mtPage.TotalRemoved + uint64(r.(*BtreeNodeItem).KeyLength) + 15
	if r == nil {
		return nil
	}
	off := r.(*BtreeNodeItem).Key
	iv := ReadAt(off, t.cow)
	b := iv.Value
	return b
}

func (t *BTree) DeleteMin() []byte {
	r := t.deleteItem(nil, removeMin)
	t.cow.curVPool.RemoveValues.PushBack(r)
	off := r.(*BtreeNodeItem).Key
	iv := ReadAt(off, t.cow)
	b := iv.Value
	return b
}

func (t *BTree) DeleteMax() []byte {
	r := t.deleteItem(nil, removeMax)
	t.cow.curVPool.RemoveValues.PushBack(r)
	off := r.(*BtreeNodeItem).Key
	iv := ReadAt(off, t.cow)
	b := iv.Value
	return b
}

func (t *BTree) deleteItem(item Item, typ toRemove) Item {
	if t.root == nil || len(t.root.items) == 0 {
		return nil
	}
	out := t.root.remove(item, t.minItems(), typ)
	if len(t.root.items) == 0 && len(t.root.children) > 0 {
		oldroot := t.root
		t.cow.dirtyPage.Insert(oldroot.nodeId)
		t.root = t.root.children[0].childNode
		t.cow.mtPage.RootId = t.root.nodeId
		t.cow.freeNode(oldroot)
	}
	if out != nil {
		t.length--
	}
	return out
}

func (t *BTree) AscendRange(greaterOrEqual, lessThan Item, iterator ItemIterator) {
	if t.root == nil {
		return
	}
	t.root.iterate(ascend, greaterOrEqual, lessThan, true, false, iterator)
}

func (t *BTree) AscendLessThan(pivot Item, iterator ItemIterator) {
	if t.root == nil {
		return
	}
	t.root.iterate(ascend, nil, pivot, false, false, iterator)
}

func AscendIterate(b BtreeNodeItem, tr *BTree, includeStart bool) []*InternalValue {
	offs := make([]uint64, 0, 100)
	count := 0
	var result []*InternalValue
	tr.AscendGreaterOrEqual(b, func(a Item) bool {
		offs = append(offs, a.(*BtreeNodeItem).Key)
		count++
		if count >= 100 {
			count = 0
			return false
		}
		return true
	}, includeStart)
	result = BatchReadAt(offs, tr.cow)
	return result
}

func (t *BTree) AscendGreaterOrEqual(pivot Item, iterator ItemIterator, includeStart bool) {
	if t.root == nil {
		return
	}
	t.root.iterate(ascend, pivot, nil, includeStart, false, iterator)
}

func (t *BTree) Ascend(iterator ItemIterator) {
	if t.root == nil {
		return
	}
	t.root.iterate(ascend, nil, nil, false, false, iterator)
}

func (t *BTree) DescendRange(lessOrEqual, greaterThan Item, iterator ItemIterator) {
	if t.root == nil {
		return
	}
	t.root.iterate(descend, lessOrEqual, greaterThan, true, false, iterator)
}

func (t *BTree) DescendLessOrEqual(pivot Item, iterator ItemIterator) {
	if t.root == nil {
		return
	}
	t.root.iterate(descend, pivot, nil, true, false, iterator)
}

func (t *BTree) DescendGreaterThan(pivot Item, iterator ItemIterator) {
	if t.root == nil {
		return
	}
	t.root.iterate(descend, nil, pivot, false, false, iterator)
}

func (t *BTree) Descend(iterator ItemIterator) {
	if t.root == nil {
		return
	}
	t.root.iterate(descend, nil, nil, false, false, iterator)
}

func (tr *BTree) GetByKey(key uint64) []byte {
	var b BtreeNodeItem
	var v []byte
	atomic.AddInt64(&(tr.cow.readRef), 1)
	b.IdxId = key
	item := tr.Get(&b)
	if item != nil {
		bItem := (tr.Get(&b)).(*BtreeNodeItem)
		offset := bItem.GetBtreeNodeItemKey()
		iv := ReadAt(offset, tr.cow)
		v = iv.Value
	} else {
		v = []byte{}
	}
	atomic.AddInt64(&(tr.cow.readRef), -1)
	return v
}

func (t *BTree) Get(key Item) Item {
	if t.root == nil {
		return nil
	}
	return t.root.get(key)
}

func (t *BTree) Min() Item {
	return min(t.root)
}

func (t *BTree) Max() Item {
	return max(t.root)
}

func (t *BTree) Has(key Item) bool {
	return t.Get(key) != nil
}

func (t *BTree) Len() int {
	return t.length
}

type Int int

func (a Int) Less(b Item) bool {
	return a < b.(Int)
}
func (a BtreeNodeItem) Less(b Item) bool {
	if t, ok := b.(BtreeNodeItem); ok {
		return a.IdxId < t.IdxId
	} else {
		return a.IdxId < b.(*BtreeNodeItem).IdxId
	}
}
func (tr BTree) GetDirtyPage() *common.Set {
	return &tr.cow.dirtyPage
}

func (tr BTree) GetRootNode() *node {
	return tr.root
}

func (tr BTree) GetNodeIds() []uint32 {
	arr := make([]uint32, 0, 32)
	for k := range tr.cow.nodeIdMap {
		arr = append(arr, k)
	}
	return arr
}

func (tr BTree) Commit() {
	d := tr.cow.dirtyPage
	fmt.Println("the dirty page length is", len(d))
	tr.cow.mtPage.RootId = tr.root.nodeId
	tr.cow.mtPage.EmptyPageCount = uint32(len(tr.cow.emptyList))
	tr.cow.mtPage.SetEmptyPage(tr.cow.NewPage)
	copy(tr.cow.metaMmap, tr.cow.mtPage.ToBytes())
	for k := range d {
		node := tr.cow.nodeIdMap[k]
		page := node.NodeToPage()
		page.WriteToMMapRegion(tr.cow)
	}
	tr.cow.dirtyPage = make(common.Set)
	fileOffset := tr.cow.curVPool.EofOffset + tr.cow.curVPool.CurrentOffest
	tr.cow.fullQueue.PushBack(tr.cow.curVPool)
	tr.cow.flushQueue = tr.cow.fullQueue
	tr.cow.fullQueue = list.New()
	Flush(tr.cow)
	tr.cow.curVPool = <-tr.cow.emptyQueue
	tr.cow.curVPool.EofOffset = fileOffset
	logger.Info("this time commit finished!")
}
