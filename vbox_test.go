package main

import (
	"fmt"
	"fyne.io/fyne/v2/widget"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("first")
	m.Run()
}

func makeUI2() (*widget.Label, *widget.Entry) {
	out := widget.NewLabel("Hello world!")
	in := widget.NewEntry()

	in.OnChanged = func(content string) {
		out.SetText("Hello " + content + "!")
	}
	return out, in
}

func TestGreeting(t *testing.T) {
	out, _ := makeUI2()

	if out.Text != "Hello world!" {
		t.Error("Incorrect initial greeting")
	}

	//fmt.Printf("%#v", in)
	fmt.Println("greeting")
}
