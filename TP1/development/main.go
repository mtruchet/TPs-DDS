package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	/*	reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter text: ")
		text, _ := reader.ReadString('\n')
		fmt.Println(text)
	*/

	fmt.Print("Ingrese Correo: ")
	text, _ := reader.ReadString('\n')
	fmt.Println(text)
}
