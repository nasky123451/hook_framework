package utils

import (
	"fmt"
)

type Printer struct{}

func NewPrinter() *Printer {
	return &Printer{}
}

func (p *Printer) PrintSection(title string) {
	fmt.Printf("\n========== %s ==========\n", title)
	fmt.Println("----------------------------------------")
}

func (p *Printer) PrintMessage(message string) {
	fmt.Println(message)
}

func (p *Printer) PrintError(err error) {
	fmt.Printf("[Error] %v\n", err)
}

func (p *Printer) PrintKeyValue(key string, value interface{}) {
	fmt.Printf("%s: %v\n", key, value)
}

func (p *Printer) PrintDivider() {
	fmt.Println("----------------------------------------")
}
