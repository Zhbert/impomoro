/*
 * MIT License
 *
 * Copyright (c) 2023 Konstantin Nezhbert
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

package main_window

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"log"
	"net/url"
)

func getContent() fyne.CanvasObject {
	verticalLayout := container.NewVBox()

	titleLine := container.NewHBox()
	titleLine.Add(layout.NewSpacer())
	titleLabel := widget.NewLabel("impomoro — simple pomodoro timer")
	titleLine.Add(titleLabel)
	titleLine.Add(layout.NewSpacer())
	verticalLayout.Add(titleLine)

	aboutLabel := widget.NewLabel("Written in Go by Konstantin Nezhbert")
	verticalLayout.Add(aboutLabel)

	licenseLabel := widget.NewLabel("Licensed by MIT license")
	verticalLayout.Add(licenseLabel)

	urlLink, err := url.Parse("https://github.com/Zhbert/impomoro")
	if err != nil {
		log.Println("Cannot pasre URL string")
	}

	contactsLabel := widget.NewHyperlink("GitHub", urlLink)
	verticalLayout.Add(contactsLabel)

	return verticalLayout
}

func ShowHelpWindow(win *fyne.Window) {
	dialog.ShowCustom("About", "OK", getContent(), *win)
	log.Println("Showed help window")
}
