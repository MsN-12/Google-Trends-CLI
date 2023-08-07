package Google_Trends_CLI

import (
	"fmt"
	"net/http"
	"os"
)

func getGoogleTrends() *http.Response {
	res, err := http.Get("https://trends.google.com/trends/trendingsearches/daily/rss?geo=US&hl=en-US")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return res
}
