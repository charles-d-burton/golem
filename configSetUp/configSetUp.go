package configSetUp

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

const filePath string = "/home/travisws/text.txt"

type Options struct {
	Mode     string
	CommPort int
	DataPort int
	BindIP   string
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

//MakeYamlFile : chreates the YAML configuration file
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

//OpenYaml : opens the configuration file and bind to the struct.
func OpenYaml() {
	//for testing, change to your user name
	filename, err1 := filepath.Abs("/home/travisws/text.txt")
	check(err1)

	var o *Options

	yamlFile, err2 := ioutil.ReadFile(filename)
	check(err2)

	err3 := yaml.Unmarshal(yamlFile, &o)
	check(err3)

	fmt.Printf("IPV4: %#v\n", o.Mode)
	fmt.Printf("IPV4: %#v\n", o.CommPort)
	fmt.Printf("IPV4: %#v\n", o.DataPort)
	fmt.Printf("IPV4: %#v\n", o.BindIP)
}
