package runtime

import (
	"os"
	"testing"
)

func TestNewForNonExistingFile(t *testing.T) {
	runtime, err := New("./tmp-config.yml")

	if err != nil {
		t.Errorf("Expected to create a new runtime with an empty configuration.")
	}

	if runtime == nil {
		t.Errorf("Expected to create a new runtime with an empty configuration.")
	}
}

func TestNewForExistingInputFile(t *testing.T) {
	// Create a test configuration file
	runtime, _ := New("./tmp-config.yml")
	runtime.SaveConfiguration()

	// Load the configuration
	runtime, err := New("./tmp-config.yml")

	if err != nil {
		t.Errorf("Expected to load the configuration.")
	}

	if runtime == nil {
		t.Errorf("Expected to load the configuration.")
	}

	t.Cleanup(func() {
		os.Remove("./tmp-config.yml")
	})
}
