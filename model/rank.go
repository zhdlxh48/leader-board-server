package model

type Rank struct {
	UserName  string `json:"user_name"`
	UserScore int64  `json:"user_score"`
	Ranking   int64  `json:"rank"`
}
