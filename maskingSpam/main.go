package main

import (
	"fmt"
	"maskingSpam/maskingSpam"
	"os"
)

func main() {

	message := os.Args[1]
	/* аргумент [1]: в аргумент при запуске вводим строку,
	которую нужно будет замаскировать */

	fmt.Println(maskingSpam.MaskingSpam(message))
}
