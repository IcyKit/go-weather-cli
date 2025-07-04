package main

import (
	"flag"
	"fmt"
	"io"
	"strings"
)

func main() {
	city := flag.String("city", "Saint-Petersburg", "Город пользователя")
	format := flag.String("format", "1", "Формат вывода (1, 2, 3)")
	flag.Parse()

	fmt.Println(*city, *format)

	r := strings.NewReader("Привет! Меня зовут Никита")
	block := make([]byte, 4)
	for {
		_, err := r.Read(block)
		fmt.Printf("%q\n", block)
		if err == io.EOF {
			break
		}
	}
}
