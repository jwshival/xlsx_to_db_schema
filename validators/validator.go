package validators

import "os"

type Validator struct {
	adapter    string
	sourcePath string
	destPath   string
	Err        error
}

func NewValidator(adapter string, sourcePath string, destPath string) *Validator {
	return &Validator{adapter: adapter, sourcePath: sourcePath, destPath: destPath}
}

func (validator Validator) Valid() *os.File {
	var err error
	err = checkFileExtension(validator.sourcePath)
	if err != nil {
		return nil
	}

	err = checkSupportedAdapter(validator.adapter)
	if err != nil {
		validator.Err = err
		return nil
	}

	err = checkFilePath(validator.sourcePath, validator.destPath)
	if err != nil {
		validator.Err = err
		return nil
	}

	err = checkFileHeader(validator.sourcePath)
	if err != nil {
		validator.Err = err
		return nil
	}

	// Open file
	return nil
}
