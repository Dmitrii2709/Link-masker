package maskservice

import (
	"fmt"
	"io/ioutil"
)

type FileProducer struct {
	inputText string
}

func NewFileProducer(inputText string) *FileProducer {
	return &FileProducer{inputText: inputText}
}

func (f *FileProducer) produce() ([]string, error) {

	fileData, err := ioutil.ReadFile(f.inputText)

	if err != nil {
		fmt.Println("Не могу прочитать файл\n", err)
	}

	message := []string{string(fileData)}

	return message, err
}
