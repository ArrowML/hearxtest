package model

type DadJoke struct {
	Joke      string `json:"joke"`
	Punchline string `json:"punchline"`
	Rating    int    `json:"rating"`
}

type PaginatedDadJokes struct {
	Count int       `json:"total_jokes"`
	Page  int       `json:"page"`
	Jokes []DadJoke `json:"jokes"`
}
