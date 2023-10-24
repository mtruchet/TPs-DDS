package state

import "fyne.io/fyne/v2/widget"

const (
	Off     int = 0
	OnMode1 int = 1
	OnMode2 int = 2
)

// BombillaSimple mantiene el estado actual y proporciona una forma de cambiarlo.
type BombillaSimple struct {
	state int
}

// NewBombillaSimple inicializa una nueva Bombilla con un estado inicial
func NewBombillaSimple() *BombillaSimple {
	return &BombillaSimple{state: Off}
}

func (b *BombillaSimple) PressButton(widgetLabel *widget.Label) {
	switch b.state {
	case Off:
		widgetLabel.SetText("Bombilla ON -> Modo 1")
		b.state = OnMode1
	case OnMode1:
		widgetLabel.SetText("Bombilla ON -> Modo 2")
		b.state = OnMode2
	case OnMode2:
		widgetLabel.SetText("Bombilla OFF")
		b.state = Off
	}
}
