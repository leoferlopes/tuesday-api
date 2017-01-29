package configuration

import (
	"io/ioutil"
	"fmt"
	"encoding/json"
	"github.com/leoferlopes/tuesday-api/api/types"
	"sync"
)

var config Configuration
var menu *types.Menu
var filePath string = "config.json"

var l sync.Mutex

func init() {
	loadConfig()
}

func readFile(object interface{}, filePath string) error {
	dat, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error when reading file ", filePath)
		fmt.Println(err.Error())
	}
	return json.Unmarshal(dat, object)
}

func loadConfig() {
	err := readFile(&config, filePath)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func GetAddress() string {
	return config.Address
}

func GetMenu() *types.Menu {
	l.Lock()
	if menu == nil{
		var m types.Menu
		readFile(&m, config.Menu)
		menu = &m
	}
	l.Unlock()
	return menu
}
