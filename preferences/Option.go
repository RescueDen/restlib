package preferences

import (
	"encoding/json"
	"log"
	"os"
)

type OptionType string

const (
	Int    OptionType = "int"
	String OptionType = "string"
	Float  OptionType = "float"
	Bool   OptionType = "bool"
)

//Restore an optiosn group from a file
func LoadOptionsGroup(jsonFile string) *OptionGroup {

	optGroup := &OptionGroup{}

	//Load in the file
	configFileStream, err := os.Open(jsonFile)

	if err == nil {
		//Get the json and add to the Params
		jsonParser := json.NewDecoder(configFileStream)
		jsonParser.Decode(&optGroup)
		configFileStream.Close()
	} else {
		log.Fatal(err)
	}

	return optGroup
}

type Option struct {
	//Store the name of th option
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`

	//Store the default value
	DefaultValue string `json:"defaultValue"`

	//Store the type
	Type OptionType `json:"type"`

	//Store min max types if possible
	MaxValue float64 `json:"maxValue,omitempty"`
	MinValue float64 `json:"minValue,omitempty"`

	//Store a list of options
	Selection []string `json:"selection,omitempty"`

	//Set if it is a hidden setting
	Hidden bool `json:"hidden"`
}

/**
Simply stores a group of options for easy display
*/
type OptionGroup struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`

	//Store a list of options
	Options []Option `json:"options"`

	//We can also old other groups
	SubGroups []OptionGroup `json:"subgroups"`
}
