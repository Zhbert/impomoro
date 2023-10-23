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
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/cmd/fyne_demo/data"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"impomoro/internal/gui/tray"
	"impomoro/internal/services/time_services"
	"log"
	"time"
)

var tomatoTime = 1500
var quitChan = make(chan bool)

func StartMainWindow() {
	app := app.New()
	app.SetIcon(data.FyneLogo)
	tray.MakeTray(app)
	window := app.NewWindow("impomoro")

	content := container.NewPadded()
	verticalBoxLayout := container.NewVBox()
	buttonsLineLayout := container.NewHBox()

	timeLabel := widget.NewLabel("25:00")
	timeLabel.TextStyle.Bold = true
	timeLabel.Alignment = fyne.TextAlign(2)

	startButton := widget.NewButton("START", func() {
		log.Println("Timer started")
		go func(curTime *int, label *widget.Label) {
			for range time.Tick(time.Second) {
				if *curTime > 0 {
					select {
					case <-quitChan:
						log.Println("Timer stopped")
						return
					default:
						*curTime--
						label.Text = time_services.SecondsToMinutes(tomatoTime)
						label.Refresh()
					}
				} else {
					return
				}
			}
		}(&tomatoTime, timeLabel)
	})
	stopButton := widget.NewButton("STOP", func() {
		log.Println("STOP button pressed")
		quitChan <- true
	})
	pauseButton := widget.NewButton("PAUSE", func() {
		log.Println("STOP button pressed")
	})

	buttonsLineLayout.Add(startButton)
	buttonsLineLayout.Add(pauseButton)
	buttonsLineLayout.Add(stopButton)

	buttonsLineLayout.Add(layout.NewSpacer())
	buttonsLineLayout.Add(timeLabel)

	verticalBoxLayout.Add(buttonsLineLayout)

	content.Add(verticalBoxLayout)
	window.SetContent(content)
	window.Resize(fyne.NewSize(400, 150))
	window.ShowAndRun()
}