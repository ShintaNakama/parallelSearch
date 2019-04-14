package googleMapPlace

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"parallelSearch/config"
)

type Results struct {
	Name string `json:"name"`
}
type RecivedResults struct {
	Results []Results `json:"results"`
}

var key string = config.Config.ApiKey

const place_base_uri = "https://maps.googleapis.com/"

func GooglePlaces(category string, place string, c chan []byte) {
	base, _ := url.Parse(place_base_uri)
	reference, _ := url.Parse("maps/api/place/textsearch/json?")
	endpoint := base.ResolveReference(reference).String()
	fmt.Println(endpoint)
	req, _ := http.NewRequest("GET", endpoint, nil)
	q := req.URL.Query()
	q.Add("query", category+"in"+place)
	q.Add("key", key)
	req.URL.RawQuery = q.Encode()
	// httpのクライアント生成
	var client *http.Client = &http.Client{}
	// 実行
	resp, _ := client.Do(req)
	// ioutilでresposeを読む
	body, _ := ioutil.ReadAll(resp.Body)
	// fmt.Printf("%T", body)
	c <- getDetail(body)
}

func getDetail(body []uint8) (v []byte) {
	var data RecivedResults
	if err := json.Unmarshal(body, &data); err != nil {
		fmt.Println(err)
	}
	// fmt.Println(data)
	v, _ = json.Marshal(data)
	fmt.Println(string(v))
	// for _, r := range data.Results {
	// 	fmt.Println(r)
	// }
	return v
}
