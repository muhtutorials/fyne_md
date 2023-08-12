package main

import (
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/test"
	"testing"
)

func Test_MakeUI(t *testing.T) {
	var testCfg config

	edit, preview := testCfg.MakeUI()

	test.Type(edit, "Hello")

	if preview.String() != "Hello" {
		t.Error("Failed -- did not find expected value in preview")
	}
}

func Test_RunApp(t *testing.T) {
	var testCfg config

	testApp := test.NewApp()
	testWin := testApp.NewWindow("Test MarkDown")

	edit, preview := testCfg.MakeUI()

	testCfg.createMenuItems(testWin)

	testWin.SetContent(container.NewHSplit(edit, preview))

	testApp.Run()

	test.Type(edit, "hey")

	if preview.String() != "hey" {
		t.Error("Failed -- did not find expected value in preview")
	}
}
