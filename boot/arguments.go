package boot

import (
	"flag"
	"xlsx_to_db_schema/constants"
)


func InitArguments(adapter *string, sourcePath *string, destPath *string) {
	flag.StringVar(*&adapter, constants.MySQLAdapter, constants.MySQLAdapter, constants.AdapterName)
	flag.StringVar(sourcePath, constants.SourcePathName, constants.SourcePathDefault, constants.SourcePathName)
	flag.StringVar(destPath, constants.DestPathName, constants.DestPathDefault, constants.DestPathName)
	flag.Parse()
}
