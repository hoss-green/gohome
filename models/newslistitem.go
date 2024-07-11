package models

type NewsListItem struct {
	Id    string
	Created string `db:"created" json:"created"`
	Updated string `db:"updated" json:"updated"`
	Published string `db:"published" json:"published"`
	Title string
}
