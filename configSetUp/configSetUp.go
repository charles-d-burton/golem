package configSetUp

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
)
//for testing, change to your user name
const filePath string = "/home/travisws/text.txt"

type Options struct {
	Mode     string
	CommPort int
	DataPort int
	BindIP   string
}
//Used to pass the values to the main file.
var Config Options

func check(e error) {
	if e != nil {
		panic(e)
	}
}

//chreates the startup YAML configuration file, if it dose not exist.
func MakeYamlFile() {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		//The default config settings
		p := Options{
			Mode:     "master",
			CommPort: 10001,
			DataPort: 10002,
			BindIP:   "0.0.0.0"}

		// Convert the Person struct to YAML.

		y, jerr := yaml.Marshal(p)
		check(jerr)

		fileOut := ioutil.WriteFile(filePath, y, 0644)
		check(fileOut)

	}
}

//opens the configuration file and bind to the struct and the variable Config.
func OpenYaml() {
	//for testing, change to your user name
	filename, err1 := filepath.Abs("/home/travisws/text.txt")
	check(err1)

	yamlFile, err2 := ioutil.ReadFile(filename)
	check(err2)

	err3 := yaml.Unmarshal(yamlFile, &Config)
	check(err3)

	//fmt.Printf("IPV4: %#v\n", o.Mode)
}