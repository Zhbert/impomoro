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

package config

import (
	"log"
	"os"
	"os/user"
	"path/filepath"
)

func DetectConfigFile() bool {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	configPath := filepath.Join(usr.HomeDir, ".impomoro")
	_, err = os.Stat(configPath)
	if err != nil {
		if os.IsNotExist(err) {
			log.Println("Config file does not exist. Creating...")
			err = os.MkdirAll(configPath, 0777)
			if err != nil {
				log.Printf("Can't create config dir: %s\n", configPath)
				return false
			}
		}
	} else {
		confFilePath := filepath.Join(configPath, "/config")
		_, err = os.Stat(confFilePath)
		if err != nil {
			f, fileErr := os.Create(confFilePath)
			if fileErr != nil {
				log.Printf("Can't create config file: %s\n", confFilePath)
				return false
			}
			defer f.Close()
		}
	}
	return true
}
