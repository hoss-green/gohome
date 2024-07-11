package models

type ProjectItem struct {
	Id        string `db:"id" json:"id"`
	Created   string `db:"created" json:"created"`
	Updated   string `db:"updated" json:"updated"`
	Title     string `db:"title" json:"title"`
	Summary   string `db:"summary" json:"summary"`
	Link      string `db:"link" json:"link"`
	Published string `db:"published" json:"published"`
}
