package lloc

import (
	"encoding/json"
	"io/ioutil"

	"github.com/LGYtech/lgo"
)

var localizations map[string]map[string]string

// LoadFromFile Loads file content to related language code
func LoadFromFile(languageCode string, filePath string) *lgo.OperationResult {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return lgo.NewFailureWithReturnObject(err)
	}
	var keyValues map[string]string
	err = json.Unmarshal(data, &keyValues)
	if err != nil {
		return lgo.NewFailureWithReturnObject(err)
	}
	localizations[languageCode] = keyValues
	return lgo.NewSuccess(nil)
}

// Get Returns requested localization value
func Get(key string, languageCode string) string {
	return localizations[languageCode][key]
}
