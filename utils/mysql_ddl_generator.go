package utils

import (
	"fmt"
	"strings"
)

/*
Generate to create DLL String
 */

func CreateTable(tableName string, engine string, charset string, collate string, createColumns []string) string{
	tmpCreateTable := fmt.Sprintf("CREATE TABLE %s (\n", tableName)
	tmpCreateColumns :=  strings.Join(createColumns, ",\n")
	tmpDestCreateTable := fmt.Sprintf("\n)")
	tmpOptions := fmt.Sprintf("ENGINE=%s AUTO_INCREMENT=34119 DEFAULT" +
		"CHARSET=%s COLLATE=%s ROW_FORMAT=DYNAMIC;", engine, charset, collate)
	table := []string{tmpCreateTable, tmpCreateColumns, tmpDestCreateTable, tmpOptions}
	return strings.Join(table, "")
}

func CreateColumn(columnName string, columnType string, options map[string]string) string {
	tmpCreateColumn := fmt.Sprintf("%s %s ", columnName, columnType)
	for _, value := range options {
		tmpCreateColumn+=" "+value;
	}
	return tmpCreateColumn
}