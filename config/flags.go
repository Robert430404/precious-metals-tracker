package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
)

type Flags struct {
	configPath string

	AddHoldingRan bool   `json:"add_holding_ran"`
	GoldAPIKey    string `json:"gold_api_key"`
}

func (self *Flags) Hydrate(path string) {
	self.configPath = path

	flags, err := os.Open(path + "/flags.json")
	if err != nil {
		self.loadDefaults()
		return
	}

	defer flags.Close()

	flagBytes, err1 := io.ReadAll(flags)
	if err1 != nil {
		self.loadDefaults()
		return
	}

	var flagStruct Flags
	err2 := json.Unmarshal(flagBytes, &flagStruct)
	if err2 != nil {
		self.loadDefaults()
		return
	}

	self.AddHoldingRan = flagStruct.AddHoldingRan
	self.GoldAPIKey = flagStruct.GoldAPIKey
}

func (self *Flags) SetAddHoldingRan(flag bool) {
	self.AddHoldingRan = flag
	err := self.writeFile()
	if err != nil {
		fmt.Printf("there was a problem writing the flag json: %v", err)
		return
	}
}

func (self *Flags) SetGoldAPIKey(flag string) {
	self.GoldAPIKey = flag
	err := self.writeFile()
	if err != nil {
		fmt.Printf("there was a problem writing the flag json: %v", err)
		return
	}
}

func (self *Flags) writeFile() error {
	blob, err := json.Marshal(self)
	if err != nil {
		return errors.New(fmt.Sprintf("there was a problem writing the flag file: %v", err))
	}

	err = os.WriteFile(self.configPath+"/flags.json", blob, 0644)

	return err
}

func (self *Flags) loadDefaults() {
	self.AddHoldingRan = false
	self.GoldAPIKey = "no_key_provided"
}
