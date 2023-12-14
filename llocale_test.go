package lloc

import (
	"testing"

	"github.com/LGYtech/lgo"
	"github.com/stretchr/testify/assert"
)

// TestSetFallbackLanguageCode tests the setting of a fallback language code
func TestSetFallbackLanguageCode(t *testing.T) {
	SetFallbackLanguageCode("en")
	assert.Equal(t, "en", fallbackLanguageCode, "Fallback language code should be 'en'")
}

// TestLoadFromFile tests loading localization data from a file
func TestLoadFromFile(t *testing.T) {
	// Test loading data
	result := LoadFromFile("en", "resources/en.json")
	assert.IsType(t, &lgo.OperationResult{}, result, "Result should be of type *lgo.OperationResult")
	assert.True(t, result.IsSuccess(), "Loading from file should succeed")

	result = LoadFromFile("tr", "resources/tr.json")
	assert.IsType(t, &lgo.OperationResult{}, result, "Result should be of type *lgo.OperationResult")
	assert.True(t, result.IsSuccess(), "Loading from file should succeed")

	// Verify data is loaded
	assert.Equal(t, "Hello", localizations["en"]["hello"], "Localization for 'hello' should be 'Hello'")
	assert.Equal(t, "Merhaba", localizations["tr"]["hello"], "Localization for 'hello' should be 'Merhaba'")
	assert.Equal(t, "Goodbye", localizations["en"]["goodbye"], "Localization for 'goodbye' should be 'Goodbye'")
	assert.Equal(t, "Güle güle", localizations["tr"]["goodbye"], "Localization for 'goodbye' should be 'Güle güle'")
}

// TestGet tests retrieving localized strings
func TestGet(t *testing.T) {
	SetFallbackLanguageCode("tr")
	LoadFromFile("en", "resources/en.json")
	LoadFromFile("tr", "resources/tr.json")

	// Test getting an existing key
	assert.Equal(t, "Hello", Get("hello", "en"), "Should return 'Hello' for 'hello' key in English")
	assert.NotEqual(t, "Hello", Get("hello", "es"), "Should not return 'Hello' for 'hello' key in Spanish if it's different")

	// Test fallback
	assert.Equal(t, "The resource", Get("onlyintr", "en"), "Should return fallback language value for non-existent key")
}

// TestReset tests resetting
func TestReset(t *testing.T) {
	SetFallbackLanguageCode("tr")
	LoadFromFile("en", "resources/en.json")

	// Test getting an existing key
	assert.Equal(t, "Hello", Get("hello", "en"), "Should return 'Hello' for 'hello' key in English")

	// Test Reset
	Reset()

	// Test getting an existing key
	assert.Equal(t, "", Get("hello", "en"), "Should return '' for 'hello' key since it has been reset")

	// Adding files again
	LoadFromFile("en", "resources/en.json")

	// Test getting an existing key
	assert.Equal(t, "Hello", Get("hello", "en"), "Should return 'Hello' for 'hello' key in English")
}
