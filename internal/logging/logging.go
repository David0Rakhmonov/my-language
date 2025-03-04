package logging

import (
	"log"
	"os"
)

var logger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)

// Log выводит сообщение в лог
func Log(message string) {
	logger.Println(message)
}
