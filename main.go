package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

func main() {
	// app new 一个 app 实例
	myApp := app.New()
	// app 实例 new 一个 window 实例
	myWindow := myApp.NewWindow("List Data")

	// binding new 一个 binding BindStringList实例
	data := binding.BindStringList(
		&[]string{"Item 1", "Item 2", "Item 3"},
	)

	// widget new 一个 widget list 实例
	list := widget.NewListWithData(data,
		// 使用 部件 widget label 实例
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		// update item
		func(i binding.DataItem, o fyne.CanvasObject) {
			// 实例 o 断言 *widget.Label  值绑定  数据 i 断言 binding.String
			o.(*widget.Label).Bind(i.(binding.String))
		})

	// widget new 一个 button 实例
	add := widget.NewButton("Append", func() {
		// 点击 button 触发  data 内部 slice 添加 list
		val := fmt.Sprintf("Item %d", data.Length()+1)
		data.Append(val)
	})

	// fyne new 一个 size 实例
	// window 实例 resize 适应 size 实例
	myWindow.Resize(fyne.NewSize(300, 300))

	// Border 布局  fyne.CanvasObject
	myWindow.SetContent(container.NewBorder(nil, add, nil, nil, list))
	myWindow.ShowAndRun()
}
