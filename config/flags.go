package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Flags struct {
	AddHoldingRan bool `json:"add_hodling_ran"`
}

func (self *Flags) Hydrate() {
	path := GetConfig().ConfigPath

	flags, err := os.Open(fmt.Sprintf("%q/flags.json", path))
	if err != nil {
		self.loadDefaults()
		return
	}

	defer flags.Close()

	flagBytes, err1 := ioutil.ReadAll(flags)
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
}

func (self *Flags) loadDefaults() {
	self.AddHoldingRan = false
}
