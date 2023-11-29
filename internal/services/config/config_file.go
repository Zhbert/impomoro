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

package config

import (
	"gopkg.in/yaml.v3"
	"impomoro/internal/services/config/structs"
	"log"
	"os"
	"os/user"
	"path/filepath"
)

/******************************************************************************
* Contents of the configuration file
******************************************************************************/

const (
	folderName = ".impomoro"
	fileName   = "config.yml"
)

/******************************************************************************
* Creating a default configuration file
******************************************************************************/

func DetectConfigFile() {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	configPath := filepath.Join(usr.HomeDir, folderName)
	_, err = os.Stat(configPath)
	if err != nil {
		if os.IsNotExist(err) {
			log.Println("Config file does not exist. Creating...")
			err = os.MkdirAll(configPath, 0777)
			if err != nil {
				log.Printf("Can't create config dir: %s\n", configPath)
			}
		}
	} else {
		confFilePath := filepath.Join(configPath, fileName)
		_, err = os.Stat(confFilePath)
		if err != nil {
			createDefaultConfig(confFilePath)
		}
	}
}

func createDefaultConfig(confFilePath string) {
	file, err := os.OpenFile(confFilePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("error opening/creating file: %v", err)
	}
	defer file.Close()

	enc := yaml.NewEncoder(file)

	err = enc.Encode(getDefaultConfigStruct())
	if err != nil {
		log.Fatalf("error encoding: %v", err)
	}

}

func getDefaultConfigStruct() structs.ConfigOptions {
	return structs.ConfigOptions{
		Display: struct {
			Width  int `yaml:"width"`
			Height int `yaml:"height"`
		}(struct {
			Width  int
			Height int
		}{Width: 400, Height: 100}),
		Time: struct {
			LongTime  int `yaml:"longTime"`
			ShortTime int `yaml:"shortTime"`
		}(struct {
			LongTime  int
			ShortTime int
		}{LongTime: 25, ShortTime: 5}),
	}
}

/******************************************************************************
* Getting data from a configuration file
******************************************************************************/

func GetConfigOptions() structs.ConfigOptions {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	configFile := filepath.Join(usr.HomeDir, folderName, fileName)

	var configOpts structs.ConfigOptions

	yamlFile, err := os.ReadFile(configFile)
	if err != nil {
		log.Fatal(err)
		return getDefaultConfigStruct()
	}
	err = yaml.Unmarshal(yamlFile, &configOpts)
	if err != nil {
		log.Fatal(err)
		return getDefaultConfigStruct()
	}

	return configOpts
}
