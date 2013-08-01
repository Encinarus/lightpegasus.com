package feeds

import (
	"bytes"
	"text/template"
  "time"
)

/*
   <?xml version="1.0" encoding="utf-8"?>
   <feed xmlns="http://www.w3.org/2005/Atom">

     <title>Example Feed</title>
     <link href="http://example.org/"/>
     <updated>2003-12-13T18:30:02Z</updated>
     <author>
       <name>John Doe</name>
     </author>
     <id>urn:uuid:60a76c80-d399-11d9-b93C-0003939e0af6</id>

     <entry>
       <title>Atom-Powered Robots Run Amok</title>
       <link href="http://example.org/2003/12/13/atom03"/>
       <id>urn:uuid:1225c695-cfb8-4ebb-aaaa-80da344efa6a</id>
       <updated>2003-12-13T18:30:02Z</updated>
       <summary>Some text.</summary>
     </entry>

   </feed>
*/

var atomTemplate = template.Must(template.New("AtomEntry").Parse(`<?xml version="1.0" encoding="utf-8"?>
  <feed xmlns="http://www.w3.org/2005/Atom">
    <title type="text">{{.Title}}</title>
    <subtitle type="html">mySubtitle</subtitle>
    <updated>{{.Updated.Format "` + time.RFC3339 + `" }}</updated>
    <id>{{.Id}}</id>
    <link href="{{.Link}}"/>

    {{range .Entries}}<entry>
      <title>{{.Title}}</title>
      <link href="{{.Link}}"/>
      <id>{{.Id}}</id>
      <updated>{{.Updated.Format "` + time.RFC3339 + `"}}</updated>
      <published>{{.Created.Format "` + time.RFC3339 + `"}}</published>
      <summary>{{.Summary}}</summary>
      <content type="xhtml" xml:lang="en">
        <div xmlns="http://www.w3.org/1999/xhtml">
          {{.ContentHtml}}
        </div>
      </content>
    </entry>
    {{end}}
  </feed>`))

func (feed *Feed) AsAtom() string {
	// TODO: Use a specific format for time fields (Created & Updated)
	
	buffer := bytes.NewBufferString("")
	err := atomTemplate.Execute(buffer, feed)
	if err != nil {
		panic(err)
	}

	return buffer.String()
}
