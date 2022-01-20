package models

type Book struct {
	Title   string `json:"title" faker:"word"`
	Summary string `json:"summary" faker:"sentence"`
	Pages   int64  `json:"pages" faker:"unix_time"`
}
