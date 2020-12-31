package lloc

import (
	"encoding/json"
	"io/ioutil"

	"github.com/LGYtech/lgo"
)

var localizations map[string]map[string]string
var defaultLanguageCode string

// Initialize Initializes size for map
func Initialize(size int) {
	localizations = make(map[string]map[string]string, size)
}

// SetDefaultLanguageCode Sets default language code
func SetDefaultLanguageCode(languageCode string) {
	defaultLanguageCode = languageCode
}

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

// Getd Returns requested localization value with default language code
func Getd(key string) string {
	return localizations[defaultLanguageCode][key]
}
