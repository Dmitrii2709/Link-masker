package main

import (
	"maskservice"
	"os"
)

func main() {

	var inputText, outputText string

	inputText = os.Args[1]
	outputText = os.Args[2]

	/* аргумент [1]: файл "Text3001.txt" с сообщением,
	которое нужно прочитать	(имя файла может быть любым) */

	/* аргумент [2]: файл "Text3001NEW.txt" в который записываем
	замаскированное сообщение, либо вводим любое название для
	создания нового файла (имя файла может быть любым) */

	fileProd := maskservice.NewFileProducer(inputText)
	filePres := maskservice.NewFilePresenter(outputText)
	service := maskservice.NewServiceName(fileProd, filePres)
	service.Run()
}
