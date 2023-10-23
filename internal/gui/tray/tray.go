/*
 *  MIT License
 *
 *  Copyright (c) 2023 Konstantin Nezhbert
 *
 *  Permission is hereby granted, free of charge, to any person obtaining a copy
 *  of this software and associated documentation files (the "Software"), to deal
 *  in the Software without restriction, including without limitation the rights
 *  to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 *  copies of the Software, and to permit persons to whom the Software is
 *  furnished to do so, subject to the following conditions:
 *
 *  The above copyright notice and this permission notice shall be included in all
 *  copies or substantial portions of the Software.
 *
 *  THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 *  IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 *  FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 *  AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 *  LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 *  OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 *  SOFTWARE.
 */

package tray

import (
	"fmt"
	"github.com/getlantern/systray"
	"impomoro/internal/services/time_services"
	"time"
)

var tomatoTime = 1500
var quitChan = make(chan bool)

func StartTrayApp() {
	systray.Run(onReady, onExit)
}

func onReady() {
	systray.SetTitle(time_services.SecondsToMinutes(tomatoTime))
	startItem := systray.AddMenuItem("Start", "Start or continue a tomato")
	pauseItem := systray.AddMenuItem("Pause", "Suspend the tomato")
	pauseItem.Disable()
	stopItem := systray.AddMenuItem("Stop", "Stop the tomato")
	systray.AddSeparator()
	exitItem := systray.AddMenuItem("Exit", "Close the application")

	go func() {
		for {
			select {
			case <-exitItem.ClickedCh:
				systray.Quit()
			case <-startItem.ClickedCh:
				fmt.Printf("Tomato is launched at %s o'clock\n", time.Now().Format("15:04:05"))
				startItem.Disable()
				pauseItem.Enable()
				go func(curTime *int) {
					for start := *curTime; start != 0; start-- {
						select {
						case <-quitChan:
							return
						default:
							time.Sleep(1 * time.Second)
							*curTime--
							systray.SetTitle(time_services.SecondsToMinutes(*curTime))
						}
					}
				}(&tomatoTime)
			case <-pauseItem.ClickedCh:
				fmt.Printf("Tomato is paused at %s o'clock\n", time.Now().Format("15:04:05"))
				quitChan <- true
				pauseItem.Disable()
				startItem.Enable()
			case <-stopItem.ClickedCh:
				fmt.Printf("Tomato is stopped at %s o'clock\n", time.Now().Format("15:04:05"))
				quitChan <- true
				startItem.Enable()
				tomatoTime = 1500
				systray.SetTitle(time_services.SecondsToMinutes(tomatoTime))
			}
		}
	}()
}

func onExit() {
	fmt.Println("Application stopped on " + time.Now().Format("2 Jan 2006 at 15:04:05"))
	close(quitChan)
}
