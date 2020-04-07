package main

import (
	"strconv"

	"fyne.io/fyne"

	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

var myTexts = make(map[int]string)
var myTabs []*widget.TabItem
var loadedTab int

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Entry Widget")

	for i := 0; i < 13; i++ {

		entryBox := widget.NewMultiLineEntry()

		scroll := widget.NewScrollContainer(
			entryBox,
		)

		tabItem := widget.NewTabItem("Tab "+strconv.Itoa(i), scroll)
		myTabs = append(myTabs, tabItem)
	}

	tabs := widget.NewTabContainer(
		myTabs...,
	)

	//leftContent := widget.NewTabContainer(
	//myTabs...,
	//)

	middleContent := fyne.NewContainerWithLayout(layout.NewBorderLayout(nil, nil, nil, nil), tabs)
	//rightContent := widget.NewMultiLineEntry()

	totalContent := widget.NewHBox(
		//leftContent,
		middleContent,
		//rightContent,
	)

	myWindow.SetContent(totalContent)
	myWindow.Resize(fyne.Size{Width: 720, Height: 720})
	myWindow.ShowAndRun()
}
