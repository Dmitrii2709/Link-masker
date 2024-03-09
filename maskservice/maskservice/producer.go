package maskservice

import (
	"fmt"
	"io/ioutil"
)

type FileProducer struct {
	inputFile string
}

func NewFileProducer(inputFile string) *FileProducer {
	return &FileProducer{inputFile: inputFile}
}

func (f *FileProducer) produce() ([]string, error) {

	fileData, err := ioutil.ReadFile(f.inputFile)

	if err != nil {
		fmt.Println("Не могу прочитать файл\n", err)
	}

	message := []string{string(fileData)}

	return message, err
}
