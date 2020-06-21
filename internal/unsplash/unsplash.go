package unsplash

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var (
	url = "https://api.unsplash.com/photos/random?orientation=landscape"
)

func GetRandomImageUrl() string {
	return requestRandomImage().Urls.Raw
}

func requestRandomImage() Response {
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("accept-version", "v1")
	req.Header.Add("authorization", "Client-ID "+os.Getenv("UNSPLASH_API_SECRET"))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Fatal(err)
	}

	return response
}
