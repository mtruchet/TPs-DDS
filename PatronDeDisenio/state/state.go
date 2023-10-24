package state

import "fyne.io/fyne/v2/widget"

// State es una interfaz para los diferentes estados de la bombilla
type State interface {
	PressButton(bombilla *Bombilla, widgetLabel *widget.Label)
}

// Bombilla mantiene el estado actual y proporciona una forma de cambiarlo.
type Bombilla struct {
	state State
}

// NewBombilla inicializa una nueva Bombilla con un estado inicial
func NewBombilla() *Bombilla {
	return &Bombilla{state: &OffState{}}
}

// PressButton permite que el estado actual maneje el evento del botón
func (b *Bombilla) PressButton(widgetLabel *widget.Label) {
	b.state.PressButton(b, widgetLabel)
}

// OffState representa el estado de la bombilla apagada
type OffState struct{}

func (s *OffState) PressButton(bombilla *Bombilla, widgetLabel *widget.Label) {
	widgetLabel.SetText("Bombilla ON -> Modo 1")
	bombilla.state = &OnStateMode1{}
}

// OnStateMode1 representa el estado de la bombilla encendida en Modo 1
type OnStateMode1 struct{}

func (s *OnStateMode1) PressButton(bombilla *Bombilla, widgetLabel *widget.Label) {
	widgetLabel.SetText("Bombilla ON -> Modo 2")
	bombilla.state = &OnStateMode2{}
}

// OnStateMode2 representa el estado de la bombilla encendida en Modo 2
type OnStateMode2 struct{}

func (s *OnStateMode2) PressButton(bombilla *Bombilla, widgetLabel *widget.Label) {
	widgetLabel.SetText("Bombilla OFF")
	bombilla.state = &OffState{}
}
