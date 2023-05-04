package userinterface

import (
	"bufio"
	"fmt"
	"os"

	"TPs-DDS/TP1/development/controller"
)

func ShowMenu() {
	fmt.Println("\nIngrese Caso de Uso: ")
	fmt.Println("1) Logín de Guia.")
	fmt.Println("2) Registrar Guía.")
	fmt.Println("3) Registrar Reserva.")
	fmt.Println("4) Salir.")

	fmt.Print("\nOpcion Elegida: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	option := scanner.Text()

	switch option {
	case "1":
		UserLogin()
		ShowMenu()
	case "2":
		InProcess()
	case "3":
		InProcess()
	case "4":
		Exit()
	default:
		fmt.Println("\nEsa opcion es incorrecta, elija nuevamente.")
		ShowMenu()
	}
}

func InProcess() {
	fmt.Println("En proceso!")
}

func Exit() {
	fmt.Println("\nGracias por utilizar nuestros servicios!")
}

func UserLogin() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("\nIngrese Correo: ")
	scanner.Scan()
	email := scanner.Text()

	fmt.Print("Ingrese contraseña: ")
	scanner.Scan()
	password := scanner.Text()

	if controller.CheckEmailAndPassword(email, password) {
		fmt.Println("\nEl correo y la contraseña son correctas. Bienvenido!")
	} else {
		fmt.Println("\nEl correo o la contraseña ingresados no son correctos!")
	}
	fmt.Println("Fin del caso de uso!")
}
