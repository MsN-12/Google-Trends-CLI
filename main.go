package Google_Trends_CLI

import (
	"encoding/xml"
)

type RSS struct {
	XMLName xml.Name `xml:"xml_name"`
	Channel *Channel `xml:"channel"`
}
type Channel struct {
	Title    string  `xml:"title"`
	ItemList []Items `xml:"item_list"`
}
type Items struct {
	Title     string `xml:"title"`
	Link      string `xml:"link"`
	Traffic   string `xml:"traffic"`
	NewsItems []News `xml:"news_items"`
}
type News struct {
	HeadLine     string `xml:"head_line"`
	HeadLineLink string `xml:"head_line_link"`
}

func main() {
	var r RSS
	data := readGoogleTrends()
	err := xml.Unmarshal(data, &r)
	if err != nil {
		println(err)
	}
}
