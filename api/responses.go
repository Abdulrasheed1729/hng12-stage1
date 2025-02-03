package api

type NumbersAPIResponse struct {
	Text   string `json:"text"`
	Found  bool   `json:"found"`
	Number int    `json:"number"`
	Type   string `json:"type"`
	Date   string `json:"date"`
}

type ErrorResponse struct {
	Number string `json:"number"`
	Error  bool   `json:"error"`
}

type ClassifyNumberResponse struct {
	Number     int      `json:"number"`
	IsPrime    bool     `json:"is_prime"`
	IsPerfect  bool     `json:"is_perfect"`
	Properties []string `json:"properties"`
	DigitSum   int      `json:"digit_sum"`
	FunFact    string   `json:"fun_fact"`
}
