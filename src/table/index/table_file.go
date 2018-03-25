package index

import (
	"bufio"
	"common"
	"fmt"
	"github.com/edsrzf/mmap-go"
	"io"
	"iowrapper"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

type Table struct {
	basePath  string
	table     string
	dataBase  string
	indexPath string
	dataPath  string
}

func NewTable(basePath, table, dataBase string) *Table {
	return &Table{
		basePath:  basePath,
		table:     table,
		dataBase:  dataBase,
		indexPath: "",
		dataPath:  "",
	}
}

func (table *Table) CreateIndexPath() string {
	ts := time.Now().Unix()
	table.indexPath = fmt.Sprintf("index_%v", ts)
	return fmt.Sprintf("%v/%v", table.GetTablePath(), table.indexPath)
}

func (table *Table) CreateTableDataPath() string {
	ts := time.Now().Unix()
	table.dataPath = fmt.Sprintf("data_%v", ts)
	return fmt.Sprintf("%v/%v", table.GetTablePath(), table.dataPath)
}
func (table Table) GetTablePath() string {
	return fmt.Sprintf("%v/%v/%v", table.basePath, table.dataBase, table.table)
}

func (table Table) GetIndexPath() string {
	return fmt.Sprintf("%v/%v/%v/%v", table.basePath, table.dataBase, table.table, table.indexPath)
}

func (table Table) GetTableDataPath() string {
	return fmt.Sprintf("%v/%v/%v/%v", table.basePath, table.dataBase, table.table, table.dataPath)
}

func CreateManifest(path string) *os.File {
	manifestPath := fmt.Sprintf("%v/%v", path, "manifest")
	return iowrapper.CreateDataFile(manifestPath)
}

func ManifestWrite(indexFileName, dataFileName string, totalRemoved uint64) {
	line := fmt.Sprintf("%v\t%v\t%v\n", indexFileName, dataFileName, totalRemoved)
	ManifeshHandle.WriteString(line)
}

func ManifestReadFileName() (string, string) {
	rd := bufio.NewReader(ManifeshHandle)
	var dataFileName, indexFileName string
	for {
		line, err := rd.ReadString('\n')
		if err != nil || io.EOF == err {
			break
		}
		arr := strings.Split(line, "\t")
		indexFileName = arr[0]
		dataFileName = arr[1]
	}
	return indexFileName, dataFileName
}

func (table *Table) CreateNewTable() *os.File {
	indexPath, dataPath := "", ""
	path := table.GetTablePath()
	os.MkdirAll(path, os.ModePerm)
	if iowrapper.PathExist(path + "/" + "manifest") {
		ManifeshHandle = CreateManifest(path)
		table.indexPath, table.dataPath = ManifestReadFileName()
		indexPath = table.GetIndexPath()
		dataPath = table.GetTableDataPath()
	} else {
		ManifeshHandle = CreateManifest(path)
	}
	f := table.CreateTable(indexPath, dataPath)
	dataPath = filepath.Base(table.GetTableDataPath())
	indexPath = filepath.Base(table.GetIndexPath())
	VaccumHolder = &VaccumData{}
	VaccumHolder.VaccumFlag = false
	VaccumHolder.VMutex = &sync.Mutex{}
	ManifestWrite(indexPath, dataPath, 0)
	return f
}

func (table *Table) CreateTable(indexPath, dataPath string) *os.File {
	var dataFile *os.File
	if !iowrapper.PathExist(dataPath) {
		dataPath = table.CreateTableDataPath()
		dataFile = iowrapper.CreateDataFile(dataPath)
	} else {
		dataFile = iowrapper.CreateDataFile(dataPath)
	}
	if iowrapper.PathExist(indexPath) {
		return dataFile
	}
	indexPath = table.CreateIndexPath()
	common.Check(iowrapper.CreateSparseFile(indexPath, 4096*1000000))
	f, err := os.OpenFile(indexPath, os.O_RDWR, 0666)
	common.Check(err)
	metaPage := NewMetaPage(INITROOTNULL, MAXPAGENUMBER/8, 0)
	bs := metaPage.ToBytes()
	mapregion, err := mmap.MapRegion(f, METAPAGEMAXLENGTH, mmap.RDWR, 0, 0)
	copy(mapregion, bs)
	mapregion.Flush()
	mapregion.Unmap()
	f.Close()
	return dataFile
}
