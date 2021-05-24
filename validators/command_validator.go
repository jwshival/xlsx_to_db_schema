package validators

import (
	"errors"
	"xlsx_to_db_schema/constants"
)

/*
Check console from user.

1. Format console (Total arguments)
2. Supported adapters (mysql - postgresql)
*/

func checkSupportedAdapter(adapter string) error {
	if adapter != constants.MySQLAdapter && adapter != constants.PSQLAdapter {
		return errors.New(constants.NotSupportAdapter)
	}
	return nil
}
