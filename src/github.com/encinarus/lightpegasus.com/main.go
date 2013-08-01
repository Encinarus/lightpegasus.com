package main

import (
	"fmt"
	"github.com/encinarus/lightpegasus.com/feeds"
	"github.com/knieriem/markdown"
	//"os"
	//"bufio"
	"bytes"
	"strings"
	"time"
)

func main() {
	contents := []feeds.Entry{
		feeds.Entry{"my title", "http://blog.lightpegasus.com/1", "1",
			time.Now(), time.Now(),
			"Test Post",
			markdownFormat("No *really* this is a test post")},
		feeds.Entry{"my title", "http://blog.lightpegasus.com/2", "2",
			time.Now(), time.Now(),
			"Test Post",
			markdownFormat("No *really* this is a test post")},
		feeds.Entry{"my title", "http://blog.lightpegasus.com/3", "3",
			time.Now(), time.Now(),
			"Test Post",
			markdownFormat("No *really* this is a test post")},
	}
	feed := feeds.Feed{"FeedTitle", time.Now(), time.Now(), feeds.Author{"alek"},
		"feedId1", "http://lightpegasus.com", contents}

	feedString := feed.AsAtom()
	fmt.Println(feedString)
}

func markdownFormat(text string) string {
	p := markdown.NewParser(&markdown.Extensions{Smart: true})
	// w := bufio.NewWriter(os.Stdout)
	buffer := bytes.NewBufferString("")
	p.Markdown(strings.NewReader(text), markdown.ToHTML(buffer))
	return buffer.String()
}
