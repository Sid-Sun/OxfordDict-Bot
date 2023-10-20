package api

// Response defines the reponse structure for Dictionaries API
type Response struct {
	Word    string `json:"word" bson:"word"`
	Results []struct {
		Definition   string   `json:"definition" bson:"definition"`
		PartOfSpeech string   `json:"partOfSpeech" bson:"partOfSpeech"`
		Examples     []string `json:"examples" bson:"examples"`
		Derivation   []string `json:"derivation" bson:"derivation"`
	} `json:"results" bson:"results"`
	// Pronunciation struct {
	// 	All string `json:"all" bson:"all"`
	// }
}

// NumberOfDefinitions returns the number of definitions for query in response
func (r Response) NumberOfDefinitions() int {
	return len(r.Results)
}

// IsEmpty checks if Response is empty and if it is, returns true
func (r Response) IsEmpty() bool {
	return len(r.Results) == 0
}
