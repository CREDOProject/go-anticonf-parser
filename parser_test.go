package goanticonfparser

import (
	"os"
	"testing"
)

func TestParser(t *testing.T) {
	fileContent, err := os.ReadFile("test/test1.configure")
	if err != nil {
		t.Error(err)
	}
	filestring := string(fileContent)
	kv := Parse(filestring)

	expects := map[string]string{
		"PKG_CONFIG_NAME": "odbc",
		"PKG_DEB_NAME":    "unixodbc-dev",
		"PKG_RPM_NAME":    "unixODBC-devel",
		"PKG_CSW_NAME":    "unixodbc_dev",
		"PKG_PACMAN_NAME": "unixodbc",
		"PKG_BREW_NAME":   "unixodbc",
		"PKG_TEST_HEADER": "<sql.h>",
		"PKG_LIBS":        "${PKG_LIBS:--lodbc}",
		"VARIABLE_1S":     "Hello",
	}

	if len(expects) != len(kv) {
		t.Fatalf("Expected length of %d, got %d", len(expects), len(kv))
	}

	keys := make([]string, 0, len(expects))
	for k2 := range expects {
		keys = append(keys, k2)
	}

	for _, k := range keys {
		if kv[k] != expects[k] {
			t.Errorf("Expected \"%s\" for \"%s\", got \"%s\".", expects[k], k, kv[k])
		}
	}

}
