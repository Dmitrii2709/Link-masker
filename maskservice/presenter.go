package maskservice

import (
	"fmt"
	"io/ioutil"
)

type FilePresenter struct {
	outputText string
}

func NewFilePresenter(outputText string) *FilePresenter {
	return &FilePresenter{outputText: outputText}
}

func (f *FilePresenter) present(mes []string) {

	data := []byte(mes[0])

	newFileData := ioutil.WriteFile(f.outputText, data, 0600)

	if newFileData != nil {
		fmt.Println("Не могу записать файл\n")
	}
}
