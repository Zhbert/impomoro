/*******************************************************************************
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
 ******************************************************************************/

package main_window

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"impomoro/internal/gui/resources"
	"impomoro/internal/gui/tray"
	"impomoro/internal/services/config"
	"impomoro/internal/services/notifications"
	"impomoro/internal/services/time_services"
	"log"
	"time"
)

var tomatoTime = getTomatoTime()
var quitChan = make(chan bool)
var confOpts = config.GetConfigOptions()
var shortPeriod = false

func StartMainWindow() {
	application := app.New()
	application.SetIcon(resources.TomatoIcon)

	window := application.NewWindow("impomoro")

	tray.MakeTray(application, window)

	content := container.NewPadded()
	verticalBoxLayout := container.NewVBox()
	buttonsLineLayout := container.NewHBox()

	timeLabel := widget.NewLabel(time_services.SecondsToMinutes(getTomatoTime()))
	timeLabel.Refresh()
	timeLabel.TextStyle.Bold = true
	timeLabel.Alignment = fyne.TextAlign(2)

	stopButton := widget.NewButton("STOP", nil)
	stopButton.SetIcon(resources.StopIcon)
	stopButton.Disable()

	pauseButton := widget.NewButton("PAUSE", nil)
	pauseButton.SetIcon(resources.PauseIcon)
	pauseButton.Disable()

	startButton := widget.NewButton("START", nil)
	startButton.SetIcon(resources.PlayIcon)

	progress := widget.NewProgressBar()
	progress.SetValue(0)
	progress.Max = float64(tomatoTime)

	stopButton.OnTapped = func() {
		log.Println("STOP button pressed")
		if !pauseButton.Disabled() {
			quitChan <- true
		}
		stopButton.Disable()
		pauseButton.Disable()
		startButton.Enable()

		tomatoTime = getTomatoTime()
		timeLabel.Text = time_services.SecondsToMinutes(tomatoTime)
		timeLabel.Refresh()
		progress.SetValue(0)
	}

	pauseButton.OnTapped = func() {
		log.Println("PAUSE button pressed")
		pauseButton.Disable()
		startButton.Enable()
		quitChan <- true
	}

	startButton.OnTapped = func() {
		log.Println("Timer started")

		stopButton.Enable()
		pauseButton.Enable()
		startButton.Disable()

		progress.SetValue(float64(tomatoTime))

		go func(curTime *int, label *widget.Label, progress *widget.ProgressBar) {
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
						progress.SetValue(float64(*curTime))
					}
				} else {
					notifications.ShowNotification(getNotificationMessage())

					startButton.Enable()
					stopButton.Disable()
					pauseButton.Disable()

					if shortPeriod {
						shortPeriod = false
					} else if !shortPeriod {
						shortPeriod = true
					}

					tomatoTime = getTomatoTime()
					label.Text = time_services.SecondsToMinutes(tomatoTime)
					label.Refresh()
					progress.Max = float64(tomatoTime)

					return
				}
			}
		}(&tomatoTime, timeLabel, progress)
	}

	buttonsLineLayout.Add(startButton)
	buttonsLineLayout.Add(pauseButton)
	buttonsLineLayout.Add(stopButton)

	buttonsLineLayout.Add(layout.NewSpacer())
	buttonsLineLayout.Add(timeLabel)

	buttonsLinePadded := container.NewPadded()
	buttonsLinePadded.Add(buttonsLineLayout)
	verticalBoxLayout.Add(buttonsLinePadded)

	progressLinePadded := container.NewPadded()
	pBarMaxLayout := container.NewStack()
	pBarMaxLayout.Add(progress)
	progressLinePadded.Add(pBarMaxLayout)
	verticalBoxLayout.Add(progressLinePadded)

	helpLine := container.NewHBox()
	helpLine.Add(layout.NewSpacer())
	helpButton := widget.NewButton("", func() {
		ShowHelpWindow(&window)
	})
	helpButton.Icon = theme.QuestionIcon()
	helpLine.Add(helpButton)
	verticalBoxLayout.Add(helpLine)

	content.Add(verticalBoxLayout)

	window.SetContent(content)
	window.Resize(fyne.NewSize(float32(confOpts.Display.Width), float32(confOpts.Display.Height)))

	window.CenterOnScreen()

	window.SetCloseIntercept(func() {
		window.Hide()
	})

	window.SetOnClosed(window.Hide)
	window.ShowAndRun()
}

func getTomatoTime() int {
	if shortPeriod {
		return confOpts.Time.ShortTime * 60
	}
	return confOpts.Time.LongTime * 60
}

func getNotificationMessage() (string, string) {
	if shortPeriod {
		return "The break is complete!", "Go to work!"
	} else if !shortPeriod {
		return "The tomato is complete!", "Take a break."
	}
	return "", ""
}
