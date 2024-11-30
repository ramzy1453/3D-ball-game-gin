package responses

type PlayerResponse struct {
	Status  int                    `json:"status"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

type ScoreBody struct {
	Score float32 `json:"score"`
}

type Leaderboard struct {
	Name      string  `json:"name"`
	BestScore float32 `json:"bestScore"`
}
