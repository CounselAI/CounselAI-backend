package counselai

type QueryReq struct {
	Query string `json:"query"`
}

type QueryRes struct {
	Data []struct {
		Unnamed0         int         `json:"Unnamed: 0"`
		ID               int         `json:"Id"`
		Name             string      `json:"Name"`
		Href             string      `json:"Href"`
		Docket           string      `json:"Docket"`
		Term             string      `json:"Term"`
		FirstParty       string      `json:"First_party"`
		SecondParty      string      `json:"Second_party"`
		Facts            string      `json:"Facts"`
		FactsLen         int         `json:"Facts_len"`
		MajorityVote     int         `json:"Majority_vote"`
		MinorityVote     int         `json:"Minority_vote"`
		FirstPartyWinner bool        `json:"First_party_winner"`
		DecisionType     string      `json:"Decision_type"`
		Disposition      interface{} `json:"Disposition"`
		IssueArea        string      `json:"Issue_area"`
	} `json:"data"`
}

type CompileReq struct {
	Ids []int `json:"ids"`
}
type CompileRes struct {
	Data string `json:"data"`
}

type GetCasesRes struct {
	Data []struct {
		Unnamed0         int         `json:"Unnamed: 0"`
		ID               int         `json:"Id"`
		Name             string      `json:"Name"`
		Href             string      `json:"Href"`
		Docket           string      `json:"Docket"`
		Term             string      `json:"Term"`
		FirstParty       string      `json:"First_party"`
		SecondParty      string      `json:"Second_party"`
		Facts            string      `json:"Facts"`
		FactsLen         int         `json:"Facts_len"`
		MajorityVote     int         `json:"Majority_vote"`
		MinorityVote     int         `json:"Minority_vote"`
		FirstPartyWinner bool        `json:"First_party_winner"`
		DecisionType     string      `json:"Decision_type"`
		Disposition      interface{} `json:"Disposition"`
		IssueArea        string      `json:"Issue_area"`
	} `json:"data"`
}
