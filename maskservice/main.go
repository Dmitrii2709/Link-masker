package main

import (
	"maskservice/maskservice"
	"os"
)

func main() {

	var inputFile, outputFile string

	inputFile = os.Args[1]
	outputFile = os.Args[2]

	/* аргумент [1]: файл "Text3001.txt" с сообщением,
	которое нужно прочитать	(имя файла может быть любым) */

	/* аргумент [2]: файл "Text3001NEW.txt" в который записываем
	замаскированное сообщение, либо вводим любое название для
	создания нового файла (имя файла может быть любым) */

	fileProd := maskservice.NewFileProducer(inputFile)
	filePres := maskservice.NewFilePresenter(outputFile)
	service := maskservice.NewServiceName(fileProd, filePres)
	err := service.Run()
	if err != nil {
		return
	}
}
