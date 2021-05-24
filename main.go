package main

import (
	"fmt"
	"xlsx_to_db_schema/boot"
	"xlsx_to_db_schema/constants"
	gen "xlsx_to_db_schema/generator"
	"xlsx_to_db_schema/validators"
)

func main() {
	var adapter, tableName, sourcePath, destPath string
	boot.InitArguments(&adapter, &tableName, &sourcePath, &destPath)

	validator := validators.NewValidator(adapter, sourcePath, destPath)
	csvRows := validator.Valid()
	if validator.Err != nil {
		fmt.Println(validator.Err.Error())
		return
	}

	gen := gen.NewGenerator(tableName, csvRows, destPath)
	gen.Execute()
	if gen.Err != nil {
		fmt.Println(gen.Err.Error())
		return
	}
	fmt.Println(constants.ExecuteDone)
}
