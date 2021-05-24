package generator

import (
	"errors"
	"os"
	"xlsx_to_db_schema/constants"
)

type Generator struct {
	tableName string
	csvRows   *[][]string
	destPath  string
	Err       error
}

type columnDB struct {
	physicalName string
	cType        string
	notNull      string
	primaryKey   string
}

func NewGenerator(tableName string, csvRows *[][]string, destPath string) *Generator {
	return &Generator{tableName: tableName, csvRows: csvRows, destPath: destPath}
}

// Execute - Generate to export DB Schema
func (generator *Generator) Execute(){
	var createColumns []string
	for index, row := range *generator.csvRows {
		// Skip when row is header
		if index == 0 {
			continue
		}
		column := &columnDB{
			physicalName: row[0],
			cType: row[1],
			notNull: row[2],
			primaryKey: row[3],
		}
		createColumns = append(createColumns, createColumn(column.physicalName, column.cType,
			column.notNull, column.primaryKey))
	}

	// Write output to file
	output := createTable(generator.tableName, createColumns)
	var (
		err error
		file *os.File
	)
	file, err = os.Create(generator.destPath)
	if err != nil {
		generator.Err = errors.New(constants.CanNotCreateFile)
		return
	}

	_, err = file.WriteString(output)
	if err != nil {
		generator.Err = errors.New(constants.CanNotWriteFile)
		return
	}
}
