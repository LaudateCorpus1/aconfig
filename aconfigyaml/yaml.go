package aconfigyaml

import (
	"github.com/avast/aconfig"
	"os"

	"gopkg.in/yaml.v2"
)

// Decoder of YAML files for aconfig.
type Decoder struct{}

// New YAML decoder for aconfig.
func New() *Decoder { return &Decoder{} }

// DecodeFile implements aconfig.FileDecoder.
func (d *Decoder) DecodeFile(filename string) (map[string]interface{}, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var raw map[string]interface{}
	if err := yaml.NewDecoder(f).Decode(&raw); err != nil {
		return nil, err
	}
	return raw, nil
}

func (d *Decoder) FileFormat() aconfig.ConfigFormat {
	return aconfig.YamlFormat
}