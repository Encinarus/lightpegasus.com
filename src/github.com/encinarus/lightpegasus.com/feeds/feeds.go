package feeds

import (
	"time"
)

type Entry struct {
	Title       string
	Link        string
	Id          string
	Created     time.Time
	Updated     time.Time
	Summary     string
	ContentHtml string
}

type Author struct {
	Name string
}

type Feed struct {
	Title   string
	Created time.Time
	Updated time.Time
	Author  Author
	Id      string
	Link    string
	Entries []Entry
}
