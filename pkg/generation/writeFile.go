package generation

import (
	"fmt"
	"os"
	"strings"

	"github.com/pkg/errors"
)

// Writer is the interface that handles io.
type Writer interface {
	WriteFile(typeName, content string) error
}

type fileWriter struct {
	basePath string
}

// NewFileWriter returns a new instance of a FileWriter.
func NewFileWriter(basePath string) Writer {
	return &fileWriter{
		basePath: basePath,
	}
}

func generateFilename(typeName string) string {
	return fmt.Sprintf("gen_%v_partialstructupdater.go", typeName)
}

func (fw *fileWriter) CreateOutputFile(typeName string) (outFile *os.File, err error) {
	fileName := strings.Join([]string{fw.basePath, generateFilename(typeName)}, "/")
	outFile, err = os.Create(fileName)

	if err != nil {
		err = errors.Wrapf(err, "error while creating file: %v", fileName)
	}

	return
}

func (fw *fileWriter) WriteFile(typeName, content string) error {
	file, err := fw.CreateOutputFile(typeName)
	if err != nil {
		return errors.Wrap(err, "error while creating the file")
	}

	fileContent := []byte(content)
	i, err := file.Write(fileContent)

	if err != nil {
		return errors.Wrap(err, "error while writing to file")
	}

	if len(fileContent) != i {
		return errors.Errorf("len(fileContent) != bytesWritten (%v/%v)", len(fileContent), i)
	}

	return nil
}
