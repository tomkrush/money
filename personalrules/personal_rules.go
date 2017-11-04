package personalrules

import (
	"encoding/json"
	"io/ioutil"
)

type PersonalRules struct {
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

func New(path string) PersonalRules {
	b, err := ioutil.ReadFile(path)

	if err != nil {
		panic(err)
	}

	personalRules := PersonalRules{}

	err = json.Unmarshal(b, &personalRules)

	if err != nil {
		panic(err)
	}

	return personalRules
}
