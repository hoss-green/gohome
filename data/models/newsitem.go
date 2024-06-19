package models

type NewsItem struct {
	Id      string `db:"id" json:"id"`
	Created string `db:"created" json:"created"`
	Updated string `db:"updated" json:"updated"`
	Title   string `db:"title" json:"title"`
	Content string `db:"content" json:"content"`
}
