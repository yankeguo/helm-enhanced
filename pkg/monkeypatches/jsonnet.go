package monkeypatches

import (
	"encoding/json"
	"strings"

	"github.com/google/go-jsonnet"
	"gopkg.in/yaml.v3"
)

func isJsonnetFile(filePath string) bool {
	filePath = strings.ToLower(filePath)
	return strings.HasSuffix(filePath, ".jsonnet") ||
		strings.HasSuffix(filePath, ".libsonnet")
}

func TranslateJsonnet(bytes []byte, filePath string) ([]byte, error) {
	if !isJsonnetFile(filePath) {
		return bytes, nil
	}

	j, err := jsonnet.MakeVM().EvaluateFile(filePath)
	if err != nil {
		return nil, err
	}

	var m map[string]any
	if err = json.Unmarshal([]byte(j), &m); err != nil {
		return nil, err
	}

	return yaml.Marshal(m)
}
