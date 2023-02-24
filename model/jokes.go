package model

type DadJoke struct {
	Joke      string `json:"joke"`
	Punchline string `json:"punchline"`
	Rating    int    `json:"rating"`
}

type PaginatedDadJokes struct {
	Count int
	Page  int
	Items []DadJoke
}
