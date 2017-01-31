package palikeypad

import "github.com/gopherjs/gopherjs/js"

const css = `
.keypad {
  position: absolute;
  border: 1px solid #ccc;
  padding: 10px;
  cursor: move;
  z-index: 21;
  text-align: center;
  background-color: #F0F8FF;
}

.keypad > div {
  display: block;
}`

var letters = [][]string{
	[]string{"q", "w", "e", "r", "t", "y", "u", "i", "o", "p"},
	[]string{"a", "s", "d", "f", "g", "h", "j", "k", "l"},
	[]string{"z", "x", "c", "v", "b", "n", "m"},
	[]string{"ā", "ḍ", "ī", "ḷ", "ṁ", "ṃ", "ñ", "ṇ", "ṭ", "ū", "ŋ", "ṅ"},
}

func insertKeypadStyle() {
	s := js.Global.Get("document").Call("createElement", "style")
	s.Set("innerHTML", css)
	head := js.Global.Get("document").Call("querySelector", "head")
	head.Call("appendChild", s)
}

// initialization function
func BindKeypad(inputId, keypadId string) {
	insertKeypadStyle()
	d := js.Global.Get("document")

	input := d.Call("querySelector", "#"+inputId)
	keypad := d.Call("querySelector", "#"+keypadId)
	keypad.Get("classList").Call("add", "keypad")

	for row := 0; row < len(letters); row++ {
		de := d.Call("createElement", "div")
		for i := 0; i < len(letters[row]); i++ {
			ie := d.Call("createElement", "input")
			ie.Set("type", "button")
			ie.Set("value", letters[row][i])
			ie.Call("addEventListener", "click", func(event *js.Object) {
				ori := input.Get("value").String()
				input.Set("value", ori+ie.Get("value").String())
			}, false)
			de.Call("appendChild", ie)
		}
		keypad.Call("appendChild", de)
	}
}
