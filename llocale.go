package lloc

import (
	"encoding/json"
	"os"
	"sync"

	"github.com/LGYtech/lgo"
)

var (
	localizations        map[string]map[string]string
	fallbackLanguageCode string
	mutex                sync.RWMutex
)

// Initialize
func init() {
	localizations = make(map[string]map[string]string)
}

// SetFallbackLanguageCode Sets default language code
func SetFallbackLanguageCode(languageCode string) {
	fallbackLanguageCode = languageCode
}

// LoadFromFile Loads file content to related language code
func LoadFromFile(languageCode string, filePath string) *lgo.OperationResult {
	mutex.Lock()
	// Read file
	data, err := os.ReadFile(filePath)
	if err != nil {
		mutex.Unlock()
		return lgo.NewFailureWithReturnObject(err)
	}
	// Read file content
	var keyValues map[string]string
	err = json.Unmarshal(data, &keyValues)
	if err != nil {
		mutex.Unlock()
		return lgo.NewFailureWithReturnObject(err)
	}
	// Add file content
	if _, exists := localizations[languageCode]; !exists {
		localizations[languageCode] = make(map[string]string)
	}
	for key, value := range keyValues {
		localizations[languageCode][key] = value
	}
	mutex.Unlock()
	return lgo.NewSuccess(nil)
}

// Get Returns requested localization value
func Get(key string, languageCode string) string {
	mutex.RLock()
	var value string
	if _, exists := localizations[languageCode][key]; !exists {
		value = localizations[fallbackLanguageCode][key]
	} else {
		value = localizations[languageCode][key]
	}
	mutex.RUnlock()
	return value
}

// Reset localizations
func Reset() {
	mutex.Lock()
	localizations = make(map[string]map[string]string)
	mutex.Unlock()
}
