package index

import (
	"common"
	"encoding/binary"
)

type BtreeNodeItem struct {
	Key     uint64
	IdxId   uint64
	KeyType byte
}

func(item BtreeNodeItem) GetBtreeNodeItemID() uint64{
	return item.IdxId
}
func(item BtreeNodeItem) GetBtreeNodeItemKey() uint64{
	return item.Key
}
func NewBtreeNodeItem(key,idxId uint64, keyType byte) *BtreeNodeItem {
	return &BtreeNodeItem{
		Key:     key,
		IdxId:   idxId,
		KeyType: keyType,
	}
}

func (item BtreeNodeItem) Size() uint16 {
	return common.INT64_LEN + common.INT64_LEN + common.INT8_LEN
}

func (item BtreeNodeItem) ToBytes(bytes []byte) int32 {
	iStart, iEnd := 0, 0
	iEnd = iStart + common.INT64_LEN
	binary.LittleEndian.PutUint64(bytes[iStart:iEnd], item.Key)
	iStart = iEnd
	iEnd = iStart + common.INT64_LEN
	binary.LittleEndian.PutUint64(bytes[iStart:iEnd], item.IdxId)
	iStart = iEnd
	iEnd = iStart + common.BYTE_LEN
	bytes[iStart] = item.KeyType
	crc := common.Crc16(bytes[0:iEnd])
	iStart = iEnd
	iEnd = iStart + common.INT16_LEN
	binary.LittleEndian.PutUint16(bytes[iStart:iEnd], crc)
	return int32(iEnd)
}

func BytesToBtreeNodeItems(barr []byte, count uint16) []*BtreeNodeItem {
	items := make([]*BtreeNodeItem, count, count)
	iStart, iEnd := uint32(0), uint32(0)
	sentiel := uint32(0)
	for i := uint16(0); i < count; i++ {
		b := new(BtreeNodeItem)
		iStart = iEnd
		iEnd = iStart + common.INT64_LEN
		b.Key = binary.LittleEndian.Uint64(barr[iStart:iEnd])
		iStart = iEnd
		iEnd = iStart + common.INT64_LEN
		b.IdxId = binary.LittleEndian.Uint64(barr[iStart:iEnd])
		iStart = iEnd
		iEnd = iStart + common.BYTE_LEN
		b.KeyType = barr[iStart]
		crc_0 := common.Crc16(barr[sentiel:iEnd])
		iStart = iEnd
		iEnd = iStart + common.INT16_LEN
		crc_1 := binary.LittleEndian.Uint16(barr[iStart:iEnd])
		_assert(crc_0 == crc_1, "the BtreeNodeItems crc is failed")
		items[i] = b
		sentiel = iEnd
	}
	return items
}

func BatchBtreeNodeItemToBytes(items []*BtreeNodeItem) []byte {
	bytes := make([]byte, BLOCKSIZE, BLOCKSIZE)
	iStart, length := int32(0), int32(0)
	for _, item := range items {
		iStart = iStart + length
		length = item.ToBytes(bytes[iStart:])
	}
	return bytes[0 : iStart+length]
}
