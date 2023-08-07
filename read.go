package Google_Trends_CLI

import (
	"io/ioutil"
	"os"
)

func readGoogleTrends() []byte {
	rep := getGoogleTrends()
	data, err := ioutil.ReadAll(rep.Body)
	if err != nil {
		println(err)
		os.Exit(1)
	}
	return data
}
