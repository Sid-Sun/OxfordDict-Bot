package api

// Response defines the reponse structure for Dictionaries API
type Response struct {
	Metadata struct {
		Provider string `json:"provider" bson:"provider"`
	} `json:"metadata" bson:"metadata"`
	Results []struct {
		LexicalEntries []struct {
			Language        string `json:"language" bson:"language"`
			LexicalCategory struct {
				Text string `json:"text" bson:"text"`
			} `json:"lexicalCategory" bson:"lexicalCategory"`
			Entries []struct {
				Senses []struct {
					Definitions []string `json:"definitions" bson:"definitions"`
					Examples    []struct {
						Text string `json:"text" bson:"text"`
					} `json:"examples" bson:"examples"`
				} `json:"senses" bson:"senses"`
			} `json:"entries" bson:"entries"`
		} `json:"lexicalEntries" bson:"lexicalEntries"`
	} `json:"results" bson:"results"`
}

// NumberOfDefinitions returns the number of definitions for query in response
func (r Response) NumberOfDefinitions() int {
	return len(r.Results[0].LexicalEntries[0].Entries[0].Senses)
}

// IsEmpty checks if Response is empty and if it is, returns true
func (r Response) IsEmpty() bool {
	return len(r.Results) == 0
}
