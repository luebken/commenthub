package db

type Comment struct {
	Owner   string
	Repo    string
	IssueNr string
	Text    string
}
