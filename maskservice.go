package maskservice

import (
	"fmt"
	"io/ioutil"
	"os"
)

type producer interface {
	produce() ([]string, error)
}

type presenter interface {
	present(string)
}

type Service struct {
	prod    producer
	message string
	pres    presenter
}

type argument struct {
	arg1 string
	arg2 string
}

func (ar argument) produce() ([]string, error) {

	fileData, err := ioutil.ReadFile(ar.arg1)

	if err != nil {
		fmt.Println("Не могу прочитать файл\n", err)
	}

	message := []string{string(fileData)}

	return message, err
}

func (ar argument) present(mes []string) {

	data := []byte(mes[0])

	newFileData := ioutil.WriteFile(ar.arg2, data, 0600)

	if newFileData != nil {
		fmt.Println("Не могу записать файл\n")
	}
}

func Run() {
	var pr argument = argument{arg1: os.Args[1], arg2: os.Args[2]}
	/* аргумент [1]: файл "Text3001.txt" с сообщением,
	которое нужно прочитать	(имя файла может быть любым) */

	/* аргумент [2]: файл "Text3001NEW.txt" в который записываем
	замаскированное сообщение, либо вводим любое название для
	создания нового файла (имя файла может быть любым) */

	message, err := pr.produce()
	if err != nil {
		fmt.Println(err)
	}

	var mess Service = Service{message: message[0]}

	newMessage := []string{mess.maskingSpam(message[0])}
	pr.present(newMessage)
}

func (s Service) maskingSpam(a string) string {

	s1 := "http://"
	s2 := " "
	s3 := "*"
	var x []byte

	for i := 0; i < len(a); i++ {
		if len(a) > len(s1) && i < (len(a)-6) && a[i:i+len(s1)] == s1 {
			x = append(x, a[i:i+len(s1)]...)

			for j := i + len(s1); j < len(a); j++ {
				if a[j] != s2[0] {
					x = append(x, s3[0])
				} else {
					x = append(x, s2[0])
					i += len(x) - len(a[:i+1])
					break
				}
			}
			i += len(x) - len(a[:i+1])
		} else {
			x = append(x, a[i])
		}
	}
	return string(x)
}
