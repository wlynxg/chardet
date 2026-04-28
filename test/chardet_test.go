package chrdet_test

import (
	"encoding/json"
	"github.com/wlynxg/chardet"
	"os"
	"strings"
	"testing"
)

type PythonResult struct {
	Encoding   string  `json:"encoding"`
	Language   string  `json:"language"`
	Confidence float64 `json:"confidence"`
}

type ChardetResults struct {
	Metadata map[string]string       `json:"metadata"`
	Results  map[string]PythonResult `json:"results"`
}

func TestCompareWithPythonChardet(t *testing.T) {
	jsonFile, err := os.Open("encoding_results.json")
	if err != nil {
		t.Fatalf("Failed to open JSON file: %v", err)
	}
	defer jsonFile.Close()

	var pythonResults ChardetResults
	decoder := json.NewDecoder(jsonFile)
	if err := decoder.Decode(&pythonResults); err != nil {
		t.Fatalf("Failed to parse JSON: %v", err)
	}

	for filePath, pyResult := range pythonResults.Results {
		content, err := os.ReadFile(filePath)
		if err != nil {
			t.Errorf("Failed to read file %s: %v", filePath, err)
			continue
		}

		goResult := chardet.Detect(content)

		if !compareResults(pyResult, goResult) {
			t.Errorf("Encoding detection mismatch for %s:\nPython: %s\nGo: %s", filePath, pyResult.Encoding, goResult.Encoding)
		}
	}
}

func compareResults(pythonResult PythonResult, goResult chardet.Result) bool {
	wanted := strings.ToLower(pythonResult.Encoding)
	if strings.ToLower(goResult.Encoding) == wanted {
		return true
	}
	charset := strings.ToLower(goResult.Charset)
	if charset == wanted {
		return true
	}
	if wanted == "utf-16" && (charset == "utf-16le" || charset == "utf-16be") {
		return true
	}
	if wanted == "utf-32" && (charset == "utf-32le" || charset == "utf-32be") {
		return true
	}
	return false
}
