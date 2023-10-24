package main

import (
	"PatronDeDisenio/state"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()

	var widgetLabel *widget.Label
	var bombilla *state.BombillaSimple

	myWindow := myApp.NewWindow("Control de Bombilla")

	bombilla = state.NewBombillaSimple()

	widgetLabel = widget.NewLabel("Bombilla OFF")

	boton := widget.NewButton("Cambiar estado", func() {
		bombilla.PressButton(widgetLabel)
	})

	content := container.New(layout.NewVBoxLayout(), boton, widgetLabel)
	myWindow.SetContent(content)

	myWindow.ShowAndRun()
}
