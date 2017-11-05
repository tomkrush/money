package rules

import (
	"encoding/json"
	"io/ioutil"
)

type Rules struct {
	MonthlyIncome int `json:"monthlyIncome"`
	Categories    []struct {
		Find    string   `json:"find"`
		Replace string   `json:"replace"`
		Some    []string `json:"some"`
	} `json:"categories"`
	Transactions []struct {
		Contains    string   `json:"contains,omitempty"`
		Replace     string   `json:"replace,omitempty"`
		Category    string   `json:"category,omitempty"`
		Need        bool     `json:"need,omitempty"`
		FindReplace string   `json:"findReplace,omitempty"`
		Remove      string   `json:"remove,omitempty"`
		Some        []string `json:"some,omitempty"`
	} `json:"transactions"`
}

func New(path string) Rules {
	b, err := ioutil.ReadFile(path)

	if err != nil {
		panic(err)
	}

	rules := Rules{}

	err = json.Unmarshal(b, &rules)

	if err != nil {
		panic(err)
	}

	return rules
}
