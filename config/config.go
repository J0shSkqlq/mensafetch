package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Configuration struct {
	FavCanteen       int               `yaml:"favCanteen"`
	CustomShorthands map[string]string `yaml:"customShorthands"`
	DataToPrint      struct {
		Metadata string   `yaml:"metadata"`
		Meals    []string `yaml:"meals"`
	} `yaml:"dataToPrint"`
}

type FlagSet struct {
	Mensaname      string
	MensaId        int
	DayOffSet      int
	ConfigFileName string
}

func NewFlagSet(name string, id int, day int, cName string) *FlagSet {
	return &FlagSet{
		Mensaname:      name,
		MensaId:        id,
		DayOffSet:      day,
		ConfigFileName: cName,
	}
}

func ReadConfig(filename string) (*Configuration, error) {
	// Read the file
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var config Configuration
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
