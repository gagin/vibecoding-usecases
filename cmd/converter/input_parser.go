package main
import (
    "encoding/json"
    "io/ioutil"
    "regexp"
)
// UseCase represents a single use case in USECASES.jsonc
type UseCase struct {
    ID                        string            json:"id"
    Actors                    []string          json:"actors"
    Description               string            json:"description"
    Preconditions             []string          json:"preconditions"
    Postconditions            []string          json:"postconditions"
    BasicFlow                 []string          json:"basic_flow"
    CoreScenario              string            json:"core_scenario"
    FeatureSpecificScenario   string            json:"feature_specific_scenario"
    Needs                     []string          json:"needs"
    ChoicesCompromises        map[string]string json:"choices_compromises_assumptions"
    Solution                  map[string]string json:"solution"
    MermaidFlow               string            json:"mermaid_flow"
}
// ParseJSONC reads and parses a JSONC file, stripping comments
func ParseJSONC(filePath string) ([]UseCase, error) {
    data, err := ioutil.ReadFile(filePath)
    if err != nil {
        return nil, err
    }

// Strip // comments
re := regexp.MustCompile(`(?m)^\s*//.*?\n`)
cleanData := re.ReplaceAll(data, []byte{})

var useCases []UseCase
if err := json.Unmarshal(cleanData, &useCases); err != nil {
    return nil, err
}
return useCases, nil

}
