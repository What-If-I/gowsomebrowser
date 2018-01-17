package layout

import (
	"github.com/What-If-I/gowsomebrowser/proto/layout"
)

func GetTestLayout() layout.Grid {
	black := layout.Color{Value: "FFFFFF"}
	text := layout.Text{Content: "Test Content", Size: "10", Color: &black}

	px40 := layout.Units{Value: 40, Type: layout.Units_PIXEL}
	percent100 := layout.Units{Value: 100, Type:layout.Units_PERCENT}
	area := layout.Area{Width: &px40, Height: &px40}

	textArea := layout.TextBox{Size: &area, Text: &text, Color: &black}
	elements := []*layout.Element{{&layout.Element_Textbox{Textbox: &textArea}}}
	elementsLayout := []*layout.ElementLayout{
		{ElemPosition: 1, RowStart: 2, RowEnd: 2, ColStart: 2, ColEnd: 2},
	}
	grid := layout.Grid{
		Size: &layout.Area{Width:&percent100, Height:&percent100}, Color: &black,
		Rows: 3, Columns: 3, Elements: elements, Elementslayout: elementsLayout}

	return grid
}
