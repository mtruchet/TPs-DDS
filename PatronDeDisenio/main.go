package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()

	myWindow := myApp.NewWindow("Control de Bombilla")

	bombillaText := widget.NewLabel("Bombilla OFF")

	boton := widget.NewButton("Encender/Apagar", func() {
		if bombillaText.Text == "Bombilla OFF" {
			bombillaText.SetText("Bombilla ON")
		} else {
			bombillaText.SetText("Bombilla OFF")
		}
	})

	content := container.New(layout.NewVBoxLayout(), boton, bombillaText)
	myWindow.SetContent(content)

	myWindow.ShowAndRun()
}
