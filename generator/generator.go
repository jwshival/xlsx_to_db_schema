package generator

import "os"

type Generator struct {
	file     *os.File
	destPath string
	Err      error
}

func NewGenerator(file *os.File, destPath string) *Generator {
	return &Generator{file: file, destPath: destPath}
}

// Execute - Generate to export DB Schema
func (generator Generator) Execute() {
	var err error
	generator.Err = err
}
