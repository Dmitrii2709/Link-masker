package maskservice

import (
	"io/ioutil"
	"log"
)

type FilePresenter struct {
	outputFile string
}

func NewFilePresenter(outputFile string) *FilePresenter {
	return &FilePresenter{outputFile: outputFile}
}

func (f *FilePresenter) present(mes []string) error {

	data := []byte(mes[0])

	err := ioutil.WriteFile(f.outputFile, data, 0600)

	if err != nil {
		//fmt.Println("Не могу записать файл\n")
		log.Fatal(err)
	}
	return nil
}
