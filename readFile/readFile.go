package readFile

import (
	"fmt"
	"log"
	"os"
)

func ReadTextFile(fileName string) *os.File {
	var file, err = os.Open(fileName)
	fmt.Println(file)
	if err != nil {
		log.Fatalf("failed to open")

	}
	return file
}
