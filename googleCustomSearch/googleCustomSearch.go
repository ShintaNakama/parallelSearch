package googleCustomSearch

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"parallelSearch/config"
	"sync"
)

var key string = config.Config.ApiKey
var cx string = config.Config.Cx

const search_base_uri = "https://www.googleapis.com/"

func GoogleSearch(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	base, _ := url.Parse(search_base_uri)
	reference, _ := url.Parse("customsearch/v1?key=" + key + "&cx=" + cx)
	endpoint := base.ResolveReference(reference).String()
	fmt.Println(endpoint)
	req, _ := http.NewRequest("GET", endpoint, nil)
	q := req.URL.Query()
	q.Add("q", s)
	q.Add("num", "3")
	fmt.Println(q.Encode())
	req.URL.RawQuery = q.Encode()
	// httpのクライアント生成
	var client *http.Client = &http.Client{}
	// 実行
	resp, _ := client.Do(req)
	// ioutilでresposeを読む
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
