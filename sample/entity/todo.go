package entity

type Todo struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Done    bool   `json:"done"`
}
