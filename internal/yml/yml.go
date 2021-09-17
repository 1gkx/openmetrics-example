package yml

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type YamlData struct {
	Currencies []struct{
		Name string `yaml:"name" json:"name"`
		Value uint64 `yaml:"value" json:"value"`
	} `yaml:"currencies" json:"currencies"`
}

func ParseFile(filepath string) (*YamlData, error) {

	yamlData, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	var data YamlData
	if err := yaml.Unmarshal(yamlData, &data); err != nil {
		return nil, err
	}

	for i := range data.Currencies {
		data.Currencies[i].Name = fmt.Sprintf("currency{name=%s}", data.Currencies[i].Name)
	}

	return &data, nil
}

func (yd *YamlData) Unmarshal() []byte {
	var resultString string

	for _, v := range yd.Currencies {
		resultString += fmt.Sprintf("%s %d\n", v.Name, v.Value)
	}

	return []byte(resultString)
}

