package configuration

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

/**
Define a simple database configuration
*/
type Configuration struct {
	//Load in the Params from json
	Params map[string]interface{}
}

//Provide a function to create a new one
func NewConfiguration(configFiles ...string) (*Configuration, error) {
	//Define a Configuration
	config := Configuration{}

	//Now march over each file
	for _, configFile := range configFiles {
		//Load in the file
		configFileStream, err := os.Open(configFile)

		if err != nil {
			return nil, err
		}
		//Get the json and add to the Params
		jsonParser := json.NewDecoder(configFileStream)
		jsonParser.Decode(&config.Params)
		configFileStream.Close()
	}

	//Return it
	return &config, nil
}

/**
 * Add function to get item
 */
func (config *Configuration) Get(key string) interface{} {
	//Get the key from the file
	param := config.Params[key]

	//Now see if it is specified in the env
	systemEnvVar := os.Getenv(key)

	//If it is not empty set it
	if len(systemEnvVar) != 0 {
		param = systemEnvVar
	}

	return param

}

/**
 * Add function to get item
 */
func (config *Configuration) GetString(key string) string {
	//Get the key from the
	return fmt.Sprint(config.Get(key))

}

/**
 * Add function to get item
 */
func (config *Configuration) GetInt(key string) (int, error) {
	//Get the key from the
	res, err := strconv.Atoi(config.GetString(key))

	return res, err

}

/**
 * Add function to get item
 */
func (config *Configuration) GetFloat(key string) (float64, error) {
	//Get the key from the
	res, err := strconv.ParseFloat(config.GetString(key), 64)

	return res, err

}

/**
 * Add function to get item
 */
func (config *Configuration) GetConfig(key string) *Configuration {
	//Get the child interface
	childConfigInterface := config.Get(key)

	//If childConfigInterface, return nil
	if childConfigInterface == nil {
		return nil
	}

	//Now cast it
	childConfig := childConfigInterface.(map[string]interface{})

	return &Configuration{childConfig}

}

/**
 * Add function to get item
 */
func (config *Configuration) GetStruct(key string, object interface{}) error {
	//Get the child interface
	childConfigInterface := config.Get(key)

	//If childConfigInterface, return nil
	if childConfigInterface == nil {
		return nil
	}

	//Now unmarshal
	jsonByte, err := json.Marshal(childConfigInterface)

	//If there is no error
	if err != nil {
		return err
	}

	//Now put the json back into the object
	err = json.Unmarshal(jsonByte, object)

	return err

}

/**
 * Add function to get item
 */
func (config *Configuration) GetStringArray(key string) []string {
	//Get the child interface
	childConfigInterface := config.Get(key)

	//Get as an array
	childArray := childConfigInterface.([]interface{})

	//Now build a new slice
	childStringArray := make([]string, 0)

	//Now march over each child array
	for _, child := range childArray {
		childStringArray = append(childStringArray, child.(string))

	}

	return childStringArray

}
