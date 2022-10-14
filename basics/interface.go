package basics

import "fmt"

type Writer interface {
	write()
}

type ConsoleWriter struct {
}
type FileWriter struct {
}

func Interface() {
	consoleWriter := ConsoleWriter{}
	write(consoleWriter)

	fileWriter := FileWriter{}
	write(fileWriter)
}

func write(writer Writer) {
	writer.write()
}

func (consoleWriter ConsoleWriter) write() {
	fmt.Println("Write content using consoleWriter")
}

func (fileWriter FileWriter) write() {
	fmt.Println("Write content using fileWriter")
}
