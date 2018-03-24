package index

import (
	"common"
	"fmt"
	"github.com/edsrzf/mmap-go"
	"iowrapper"
	"os"
	"path/filepath"
)

type Table struct {
	basePath   string
	table      string
	dataBase   string
	columnName string
	fileName   string
}

func NewTable(basePath, table, dataBase, columnName, fileName string) *Table {
	return &Table{
		basePath:   basePath,
		table:      table,
		dataBase:   dataBase,
		columnName: columnName,
		fileName:   fileName,
	}
}
func (table Table) GetTablePath() string {
	return fmt.Sprintf("%v/%v/%v/%v", table.basePath, table.dataBase, table.table, table.columnName)
}

func (table Table) GetTableDataPath() string {
	return fmt.Sprintf("%v/%v/%v/%v", table.basePath, table.dataBase, table.table, table.fileName)
}

func (table Table) CreateTable() *os.File {
	path := table.GetTablePath()
	dataPath := table.GetTableDataPath()
	dataFile := iowrapper.CreateDataFile(dataPath)
	os.MkdirAll(filepath.Base(path), os.ModePerm)
	if iowrapper.PathExist(path) {
		return dataFile
	}
	common.Check(iowrapper.CreateSparseFile(path, 4096*1000000))
	f, err := os.OpenFile(path, os.O_RDWR, 0666)
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
