package httpclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"golang-arch/utils"
)

const (
	templateVersion        string = "1.0"
	envESBasicAuthUsername string = "GO_ELASTICSEARCH_ALERTS_ES_USERNAME"
	envESBasicAuthPassword string = "GO_ELASTICSEARCH_ALERTS_ES_PASSWORD"
	defaultStateIndexAlias string = "log-alert-manager"
	defaultTimestampFormat string = time.RFC3339
	defaultBodyField       string = "hits.hits._source"
)

func GetNextQuery() { // nolint: funlen

	var client *http.Client
	client = &http.Client{
		Timeout: time.Second * 10,
	}

	payload := fmt.Sprintf(`{
  "query": {
    "bool": {
      "must": [
        {
          "term": {
            "rule_name": {
	      "value": %q
	    }
	  }
	}
      ]
    }
  },
  "sort": [
    {
      "next_query": {
	"order": "desc"
      }
    }
  ],
  "size": 1
}`, "galaxy-alarm")

	u, err := url.Parse("http://10.7.20.157:9200/log-alert-manager/_search")
	//u, err := url.Parse("https://www.baidu.com")
	if err != nil {
		fmt.Printf("error parsing URL: %v", err)
		return
	}
	query := u.Query()
	query.Add("filter_path", "hits.hits._source.next_query")
	u.RawQuery = query.Encode()

	req, err := http.NewRequest("GET", u.String(), bytes.NewBufferString(payload))
	//req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		fmt.Printf("error creating new HTTP request instance: %v", err)
		return
	}

	if bytes.NewBufferString(payload) != nil {
		req.Header.Add("Content-Type", "application/json")
	}
	resp, err := client.Do(req)

	if err != nil {
		fmt.Printf("error making HTTP do: %v", err)
		return
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	// b, err := ioutil.ReadAll(resp.Body)  Go.1.15 and earlier
	if err != nil {
		fmt.Printf("io error:%v", err)
		return
	}
	fmt.Println(string(b))
	//fmt.Printf("Get response:.\n", string(b))

	if resp.StatusCode != 200 {
		fmt.Printf("received non-200 response status (status: %q)", resp.Status)
		return
	}

	data := make(map[string]interface{})
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil { // nolint: govet
		fmt.Printf("error JSON-decoding HTTP response: %v", err)
		return
	}

	if len(data) < 1 {
		fmt.Printf("no records found for this rule")
		return
	}

	nextRaw := utils.Get(data, "hits.hits[0]._source.next_query")
	if nextRaw == nil {
		fmt.Printf("field 'next_query' not found")
		return
	}

	nextString, ok := nextRaw.(string)
	if !ok {
		fmt.Printf("'next_query' value could not be cast to string")
		return
	}

	t, err := time.Parse(defaultTimestampFormat, nextString)
	if err != nil {
		fmt.Printf("error parsing time: %v", err)
		return
	}
	fmt.Printf("Get time: %v", t)
}
