package responses

type PlayerResponse struct {
	Status  int                    `json:"status"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

type ScoreBody struct {
	Score int `json:"score"`
}

type Leaderboard struct {
	Name      string `json:"name"`
	BestScore int    `json:"bestScore"`
}
