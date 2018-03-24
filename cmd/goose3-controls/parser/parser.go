package parser

import (
	"encoding/json"
	"fmt"
	. "goose3-controls/cmd/goose3-controls/helper"
	"io/ioutil"
	"os"
)

func ToJson(c interface{}) string {
	bytes, err := json.Marshal(c)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return string(bytes)
}

func ToString(c Command) string {
	return ToJson(c)
}

func GetCommands() []Command {
	raw, err := ioutil.ReadFile("./test.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var c []Command
	json.Unmarshal(raw, &c)
	return c
}
