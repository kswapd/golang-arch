package mytools

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
)

const (
	MessageTypeSuccess = "1000"
	MessageTypeERROR   = "9999"
	CurrentVersion     = "1.0.0"
)

type Data struct {
	MessageType string      `json:"messageType"`
	Version     string      `json:"version"`
	Time        string      `json:"time"`
	Body        interface{} `json:"body"`
	ErrInfo     string      `json:"errInfo"`
}

var (
	// Version is the version of the binary
	staticDir       string = "./mytools/static"
	portalIndexFile string = "index.html"
	ConstPublicErr         = Data{MessageType: MessageTypeERROR, Version: CurrentVersion, ErrInfo: "Unknown Error"}
)

type GoldenTD struct {
	Name      string  `json:"name"`
	TradeTime float64 `json:"tradeTime"`
	ReqTime   int64   `json:"reqTime"`
	Price     float64 `json:"price"`
	Unit      string  `json:"unit"`
}

func handler(w http.ResponseWriter, r *http.Request) {

	/*tmpl, err := template.ParseFiles(fmt.Sprintf("%s%s", staticPath, portalIndex))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}*/

	// Create a file server handler for the specified directory.
	//staticDir := fmt.Sprintf("%s%s", staticPath, portalIndex)
	fs := http.FileServer(http.Dir(staticDir))

	// Handle requests to the root path by serving files from the static directory.
	http.Handle("/", fs)
}

func currentMilliseconds() (int64, string) {
	cTime := time.Now()
	nano := cTime.UnixNano()
	t := time.Unix(0, nano)
	tStr := t.Format(time.DateTime) // "2006" is the reference year in Go's time formatting
	_, offsetSec := cTime.Zone()
	//log.Infof("Offset: %+03d:%02d\n", offsetHours, offsetMinutes)
	timeStr := fmt.Sprintf("%s %+03d:%02d\n", tStr, offsetSec/3600, (offsetSec%3600)/60)
	log.Infof("time:%d, %s", nano/1e6, timeStr)
	return nano / 1e6, timeStr
}

func getGoldTD() (*GoldenTD, error) {
	milSec, _ := currentMilliseconds()
	url := fmt.Sprintf("https://api.jijinhao.com/quoteCenter/realTime.htm?codes=JO_9753&_=%d", milSec)
	h := map[string]string{
		"accept":                   "*/*",
		"accept-language":          "en-US,en;q=0.9,zh-CN;q=0.8,zh;q=0.7,zh-TW;q=0.6",
		"referer":                  "https://quote.cngold.org/'",
		"sec-fetch-dest":           "script",
		"sec-fetch-mode":           "no-cors",
		"sec-fetch-site":           "cross-site",
		"sec-fetch-storage-access": "active",
		"user-agent":               "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/136.0.0.0 Safari/537.36",
	}
	if respData, err := CallAPI("GET", url, h, nil); err == nil {
		str := string(respData)
		log.Infof("Get response success:%s", str)
		targetStr := "var quote_json = "
		if strings.Contains(str, targetStr) {
			jsonMsg := str[strings.Index(str, targetStr)+len(targetStr):]
			var data map[string]interface{}
			err := json.Unmarshal([]byte(jsonMsg), &data)
			if err != nil {
				log.Error(err)
				return nil, err
			}

			JO_9753 := data["JO_9753"].(map[string]interface{})

			tdTime := JO_9753["time"].(float64)
			tdPrice := JO_9753["q63"].(float64)
			tdName := JO_9753["showName"].(string)
			unit := JO_9753["unit"].(string)
			return &GoldenTD{
				TradeTime: tdTime,
				Price:     tdPrice,
				Name:      tdName,
				ReqTime:   milSec,
				Unit:      unit,
			}, nil
		} else {
			errMsg := fmt.Sprintf("Can not find target string to convert to json, return data: %s", str)
			log.Error(errMsg)
			return nil, errors.New(errMsg)
		}
	} else {
		log.Error(err)
		return nil, err

	}
}

func handleData(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		if respData, err := getGoldTD(); err == nil {
			_, timeStr := currentMilliseconds()
			log.Infof("Get getGoldTD success:%v", respData)
			retData := Data{
				MessageType: MessageTypeSuccess,
				Version:     CurrentVersion,
				Body:        respData,
				Time:        timeStr,
			}
			srcMsg, err := json.Marshal(retData)
			if err != nil {
				log.Error(err)
				curErr := ConstPublicErr
				curErr.Time = timeStr
				curErr.ErrInfo = err.Error()
				curErrMsgByte, _ := json.Marshal(curErr)
				w.Write(curErrMsgByte)
			} else {
				w.Write(srcMsg)
			}
		} else {
			_, timeStr := currentMilliseconds()
			log.Error(err)
			curErr := ConstPublicErr
			curErr.Time = timeStr
			curErr.ErrInfo = err.Error()
			curErrMsgByte, _ := json.Marshal(curErr)
			w.Write(curErrMsgByte)
		}
	}
	return
}

type Page struct {
	PageNo   int64 `json:"pageNo"`
	PageSize int64 `json:"pageSize"`
}
type Entity struct {
	ReqSourceCode string   `json:"reqSourceCode"`
	BizDomain     string   `json:"bizDomain"`
	OrgIdList     []string `json:"orgIdList"`
}
type LshmQueryInfo struct {
	EntityQ Entity `json:"entity"`
	PageQ   Page   `json:"page"`
}

func getLshmOrgInfo() {
	log.Info("getLshmOrgInfo start")
	url := "http://47.97.217.191:8080/restcloud/user_center/apiV2/person2x/organization/queryOrg"
	url2 := "https://ipaas-pre-gw.hnlshm.com/restcloud/user_center/apiV2/person2x/organization/queryOrg"
	url3 := "https://ipaas-pre-gw.hnlshm.com/restcloud/user_center/apiV2/person2x/organization/queryOrg"

	url4 := "http://47.97.217.191:8080/restcloud/user_center/apiV2/person2x/person/queryPerson"
	_ = url
	_ = url2
	_ = url3
	_ = url4
	h := map[string]string{
		"accept":          "*/*",
		"accept-language": "en-US,en;q=0.9,zh-CN;q=0.8,zh;q=0.7,zh-TW;q=0.6",
		"appKey":          "68351dcc5d25bb1f1af25b23",
	}
	q := LshmQueryInfo{
		EntityQ: Entity{
			ReqSourceCode: "2011",
			BizDomain:     "10",
			OrgIdList: []string{
				"900910566",
				"900701733",
			},
		},
		PageQ: Page{
			PageNo:   1,
			PageSize: 2,
		},
	}
	//orgId":"900910566",
	if qser, err := json.Marshal(q); err == nil {
		/*if respData, err := CallAPI("POST", url, h, qser); err == nil {
			str := string(respData)
			log.Infof("Get org info success:\n%s", str)
		} else {
			log.Error(err)
			return
		}*/

		if respData, err := CallAPI("POST", url4, h, qser); err == nil {
			str := string(respData)
			log.Infof("Get person info success:\n%s", str)
		} else {
			log.Error(err)
			return
		}

	} else {
		log.Error(err)
	}

	log.Info("getLshmOrgInfo finish")

}
func RunHtmlView() {
	var port int = 8888
	log.Infof("Starting server on :%d", port)
	getLshmOrgInfo()
	fs := http.FileServer(http.Dir(staticDir))
	// Handle requests to the root path by serving files from the static directory.
	http.Handle("/", fs)
	http.HandleFunc("/api/data", handleData)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		log.Fatal(err)
	}
}

func CallAPI(method string, url string, headers map[string]string, body []byte) ([]byte, error) {

	request, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	//request.Header.Set("Content-Type", "application/json; charset=UTF-8")
	for key, value := range headers {
		request.Header.Set(key, value)
	}
	//client := &http.Client{}
	client := &http.Client{
		Transport: &http.Transport{
			Proxy: nil,
		},
	}
	response, err := client.Do(request)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	defer response.Body.Close()
	respData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return respData, nil
}
