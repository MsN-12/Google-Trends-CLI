package Google_Trends_CLI

import (
	"encoding/xml"
	"fmt"
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
		fmt.Println(err)
	}
	fmt.Println("\n Bellow are all the google trends for today !")
	fmt.Println("----------------------------------------------")

	for i := range r.Channel.ItemList {
		rank := i + 1
		fmt.Println("[+]", rank)
		fmt.Println("Searching Term:", r.Channel.ItemList[i].Title)
		fmt.Println("Link of trend: ", r.Channel.ItemList[i].Link)
		fmt.Println("Headline: ", r.Channel.ItemList[i].NewsItems[0].HeadLine)
		fmt.Println("Link to article: ", r.Channel.ItemList[i].NewsItems[0].HeadLineLink)
		fmt.Println("----------------------------------------------")

	}
}
