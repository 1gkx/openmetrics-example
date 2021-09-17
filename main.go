package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"net/http"
)

type YamlData struct {
	Currencies []struct{
		Name string `yaml:"name" json:"name"`
		Value uint64 `yaml:"value" json:"value"`
	} `yaml:"currencies" json:"currencies"`
}

func main() {

	yamlData, err := ioutil.ReadFile("currencies.yaml")
	if err != nil {
		panic(err)
	}

	var data YamlData
	if err := yaml.Unmarshal(yamlData, &data); err != nil {
		panic(err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)

		// formatted to currencies_name_ usd
		// formatted to currencies_value_ 70

		unmarshalData := make(map[string]uint64)
		for _, c := range data.Currencies {
			key := fmt.Sprintf("currency{name=%s}", c.Name)
			unmarshalData[key] = c.Value
		}

		response := unmarshal(unmarshalData)

		w.Write(response)
	})

	http.Handle("/", r)
	if err := http.ListenAndServe(":8180", nil); err != nil {
		panic(err)
	}
}

func unmarshal(data map[string]uint64) []byte {
	var resultString string

	for k, v := range data {
		resultString += fmt.Sprintf("%s %d\n", k, v)
	}

	return []byte(resultString)
}
