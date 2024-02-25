package main

import "fmt"

type (
	IButton interface {
		Press() string
	}

	ITextBox interface {
		ShowText() string
	}

	MacButton struct {
	}

	MacTextBox struct {
	}

	WinButton struct {
	}

	WinTextBox struct {
	}

	IFactory interface {
		CreateButton() IButton
		CreateTextBox() ITextBox
	}

	MacFactory struct {
	}

	WinFactory struct {
	}

	GUIAbstractFactory struct {
	}
)

func (m MacButton) Press() string {
	return "Mac button pressed"
}

func (m MacTextBox) ShowText() string {
	return "Text from Mac"
}

func (w WinButton) Press() string {
	return "Window button pressed"
}

func (w WinTextBox) ShowText() string {
	return "Text from window"
}

func (m MacFactory) CreateButton() IButton {
	return MacButton{}
}

func (m MacFactory) CreateTextBox() ITextBox {
	return MacTextBox{}
}

func (w WinFactory) CreateButton() IButton {
	return WinButton{}
}

func (w WinFactory) CreateTextBox() ITextBox {
	return WinTextBox{}
}

func (g GUIAbstractFactory) CreateFactory(osType string) IFactory {
	switch osType {
	case "win":
		return WinFactory{}
	default:
		return MacFactory{}
	}
}

func main() {
	factory := &GUIAbstractFactory{}
	button := factory.CreateFactory("win")
	fmt.Println(button.CreateButton().Press())
	fmt.Println(button.CreateTextBox().ShowText())

	fmt.Println()

	macbutton := factory.CreateFactory("macos")
	fmt.Println(macbutton.CreateButton().Press())
	fmt.Println(macbutton.CreateTextBox().ShowText())
}
