package mmhm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
)

const (
	MessageTypeSuccess    = "1000"
	MessageTypeERROR      = "9999"
	CurrentVersion        = "1.0.0"
	base                  = 10
	bitSize               = 64
	lshmDevUrlPrefix      = "http://47.97.217.191:8080/restcloud/user_center/apiV2/person2x"
	lshmPreUrlPrefix      = "https://ipaas-pre-gw.hnlshm.com/restcloud/user_center/apiV2/person2x"
	lshmPrdUrlPrefix      = "https://ipaas-gw.hnlshm.com/restcloud/user_center/apiV2/person2x"
	lshmAppHeaderKey      = "appKey"
	lshmDevAppHeaderValue = "68351dcc5d25bb1f1af25b23"
	lshmOrgAPIUri         = "/organization/queryOrg"
	lshmPersonAPIUri      = "/person/queryPerson"
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

type Page struct {
	PageNo      int64    `json:"pageNo"`
	PageSize    int64    `json:"pageSize"`
	OrderFields []string `json:"orderFields"`
	Order       string   `json:"order"`
	AutoCount   bool     `json:"autoCount"`
	TotalCount  int64    `json:"totalCount"`
	TotalPages  int64    `json:"totalPages"`
}
type QueryEntity struct {
	ReqSourceCode string   `json:"reqSourceCode"`
	BizDomain     string   `json:"bizDomain"`
	OrgIdList     []string `json:"orgIdList"`
}
type LshmQueryInfo struct {
	QueryEntity `json:"entity"`
	Page        `json:"page"`
}

type LshmOrg struct {
	ID                 string   `json:"id"`
	BizDomain          string   `json:"bizDomain"`
	BelongBrand        string   `json:"belongBrand"`
	Code               string   `json:"code"`
	OrgName            string   `json:"orgName"`
	IsEnable           int      `json:"isEnable"`
	OrgType            string   `json:"orgType"`
	ParentId           string   `json:"parentId"`
	DepartmentId       string   `json:"departmentId"`
	ParentDepartmentId int64    `json:"parentDepartmentId"`
	OrgLevel           string   `json:"orgLevel"`
	TreeCode           string   `json:"treeCode"`
	Children           []string `json:"children"`
	BeisonOrgLevel     string   `json:"beisonOrgLevel"`
}

type LshmPerson struct {
	UserId           string `json:"userId"`
	RealName         string `json:"realName"`
	Phone            string `json:"phone"`
	ThirdUid         string `json:"thirdUid"`
	Email            string `json:"email"`
	UserStatus       string `json:"userStatus"`
	Sex              string `json:"sex"`
	DirectLeaderId   int64  `json:"directLeaderId"`
	DirectLeaderName string `json:"directLeaderName"`
	Code             string `json:"code"`
	//JobInfoList      string `json:"jobInfoList"`
}

type LshmOrgRespData struct {
	Page
	Records []LshmOrg `json:"records"`
}
type LshmOrgRespInfo struct {
	RespCode        string `json:"respCode"`
	RespMsg         string `json:"respMsg"`
	LshmOrgRespData `json:"data"`
}

type LshmPersonRespData struct {
	Page
	Records []LshmPerson `json:"records"`
}
type LshmPersonRespInfo struct {
	RespCode           string `json:"respCode"`
	RespMsg            string `json:"respMsg"`
	LshmPersonRespData `json:"data"`
}

func getLshmOrgInfo() ([]LshmOrg, error) {

	var defaultPageSize int64 = 100
	reqSourceCode := "2011"
	bizDomain := "10"
	reqUrl := fmt.Sprintf("%s%s", lshmDevUrlPrefix, lshmOrgAPIUri)
	log.Infof("getLshmOrgInfo:%s", reqUrl)
	h := map[string]string{
		"accept":         "*/*",
		lshmAppHeaderKey: lshmDevAppHeaderValue,
	}
	q := LshmQueryInfo{
		QueryEntity: QueryEntity{
			ReqSourceCode: reqSourceCode,
			BizDomain:     bizDomain,
		},
		Page: Page{
			PageNo:   1,
			PageSize: defaultPageSize,
		},
	}
	if qm, err := json.Marshal(q); err != nil {
		log.Errorf("getLshmOrgInfo, error: %s", err)
		return nil, err
	} else {
		if respData, err := CallAPI("POST", reqUrl, h, qm); err != nil {
			log.Errorf("getLshmOrgInfo, error: %s", err)
			return nil, err
		} else {
			//str := string(respData)
			log.Infof("Get lshm org info success, data len: %d", len(respData))
			resp := LshmOrgRespInfo{}
			allOrg := make([]LshmOrg, 0)
			if err := json.Unmarshal(respData, &resp); err != nil {
				log.Errorf("getLshmOrgInfo, error: %s", err)
				return nil, err
			} else {
				allOrg = append(allOrg, resp.LshmOrgRespData.Records...)
				//log.Infof("Get org info:%+v", resp.LshmOrgRespData.Records)
				if resp.LshmOrgRespData.TotalPages > 1 {
					for i := int64(2); i <= resp.LshmOrgRespData.TotalPages; i++ {
						q.PageNo = i
						qm, _ = json.Marshal(q)
						if respData, err := CallAPI("POST", reqUrl, h, qm); err != nil {
							log.Error("getLshmOrgInfo, %d, error: %s", i, err)
							return nil, err
						} else {
							//str := string(respData)
							log.Infof("Get lshm org info success, data len: %d, %d", i, len(respData))
							if err := json.Unmarshal(respData, &resp); err != nil {
								log.Errorf("getLshmOrgInfo, error: %s", err)
								return nil, err
							} else {
								//log.Infof("Get org info success:\n %+v", resp)
								allOrg = append(allOrg, resp.LshmOrgRespData.Records...)
							}

						}
					}

				}
				return allOrg, nil
			}

		}
	}

}

func getLshmPersonInfo() ([]LshmPerson, error) {

	var defaultPageSize int64 = 100
	reqSourceCode := "2011"
	bizDomain := "10"
	reqUrl := fmt.Sprintf("%s%s", lshmDevUrlPrefix, lshmPersonAPIUri)
	log.Infof("getLshmPersonInfo:%s", reqUrl)
	h := map[string]string{
		"accept":         "*/*",
		lshmAppHeaderKey: lshmDevAppHeaderValue,
	}
	q := LshmQueryInfo{
		QueryEntity: QueryEntity{
			ReqSourceCode: reqSourceCode,
			BizDomain:     bizDomain,
		},
		Page: Page{
			PageNo:   1,
			PageSize: defaultPageSize,
		},
	}
	if qm, err := json.Marshal(q); err != nil {
		log.Errorf("getLshmPersonInfo, error: %s", err)
		return nil, err
	} else {
		if respData, err := CallAPI("POST", reqUrl, h, qm); err != nil {
			log.Errorf("getLshmPersonInfo, error: %s", err)
			return nil, err
		} else {
			//str := string(respData)
			log.Infof("Get lshm person info success, data len: %d", len(respData))
			resp := LshmPersonRespInfo{}
			allPerson := make([]LshmPerson, 0)
			if err := json.Unmarshal(respData, &resp); err != nil {
				log.Errorf("getLshmPersonInfo, error: %s", err)
				return nil, err
			} else {
				allPerson = append(allPerson, resp.LshmPersonRespData.Records...)
				if resp.LshmPersonRespData.TotalPages > 1 {
					for i := int64(2); i <= resp.LshmPersonRespData.TotalPages; i++ {
						q.PageNo = i
						qm, _ = json.Marshal(q)
						if respData, err := CallAPI("POST", reqUrl, h, qm); err != nil {
							log.Errorf("getLshmPersonInfo, %d, error: %s", i, err)
							return nil, err
						} else {
							//str := string(respData)
							log.Infof("Get lshm person info success, data len: %d,%d", i, len(respData))
							if err := json.Unmarshal(respData, &resp); err != nil {
								log.Errorf("getLshmPersonInfo, error: %s", err)
								return nil, err
							} else {
								allPerson = append(allPerson, resp.LshmPersonRespData.Records...)
							}

						}
					}

				}
				return allPerson, nil
			}

		}
	}

}

func getLshmOrgInfobk() {
	log.Info("getLshmOrgInfo start")
	url := "http://47.97.217.191:8080/restcloud/user_center/apiV2/person2x/organization/queryOrg"
	url2 := "https://ipaas-pre-gw.hnlshm.com/restcloud/user_center/apiV2/person2x/organization/queryOrg"
	url3 := "https://ipaas-pre-gw.hnlshm.com/restcloud/user_center/apiV2/person2x/organization/queryOrg"
	url4 := "http://47.97.217.191:8080/restcloud/user_center/apiV2/person2x/person/queryPerson"

	h := map[string]string{
		"accept":          "*/*",
		"accept-language": "en-US,en;q=0.9,zh-CN;q=0.8,zh;q=0.7,zh-TW;q=0.6",
		"appKey":          "68351dcc5d25bb1f1af25b23",
	}
	q := LshmQueryInfo{
		QueryEntity: QueryEntity{
			ReqSourceCode: "2011",
			BizDomain:     "10",
		},
		Page: Page{
			PageNo:   200,
			PageSize: 2,
		},
	}
	_ = url
	_ = url2
	_ = url3
	_ = url4
	//orgId":"900910566",
	if qser, err := json.Marshal(q); err == nil {
		log.Info("qser:", string(qser))
		if respData, err := CallAPI("POST", url, h, qser); err == nil {
			str := string(respData)
			log.Infof("Get org info success:\n %s", str)

			resp := LshmOrgRespInfo{}
			//allOrg := make([]LshmOrg, 0)
			if err := json.Unmarshal(respData, &resp); err != nil {
				log.Errorf("getLshmOrgInfo Unmarshal error: %s", err)
			} else {
				log.Infof("Get org info success:\n %+v", resp)
			}

		} else {
			log.Error(err)
			return
		}
		/*if respData, err := CallAPI("POST", url4, h, qser); err == nil {
			str := string(respData)
			log.Infof("Get person info success:\n %s", str)
		} else {
			log.Error(err)
			return
		}*/

	} else {
		log.Error(err)
	}

	log.Info("getLshmOrgInfo finish")

}

func If(condition bool, trueVal, falseVal interface{}) interface{} {
	if condition {
		return trueVal
	}
	return falseVal
}

// Define the tree structure with children field
type LshmOrgTreeNode struct {
	LshmOrg
	Level    int
	Children []*LshmOrgTreeNode
}

func ProcessOrgs(orgs []LshmOrg) ([]LshmOrg, error) {
	orgMap := make(map[string]*LshmOrgTreeNode)

	// Step 1: Initialize all nodes with Level = 0 (unassigned)
	for _, org := range orgs {
		orgMap[org.ID] = &LshmOrgTreeNode{
			LshmOrg:  org,
			Level:    0,
			Children: []*LshmOrgTreeNode{},
		}
	}

	// Step 2: Link children and build roots
	roots := []*LshmOrgTreeNode{}

	for id, node := range orgMap {
		if parentNodeId := node.ParentId; parentNodeId != "" {
			if parentNode, exists := orgMap[parentNodeId]; exists {
				parentNode.LshmOrg.Children = append(parentNode.LshmOrg.Children, node.ID)
				parentNode.Children = append(parentNode.Children, node)
			} else {
				//lshm use -2 as root id, we need to convert to -1 to be compatible with ml
				log.Warnf("Parent not found for org %s with ParentId %s", id, parentNodeId)
				node.ParentId = "-1"
				roots = append(roots, node)
			}
		} else {
			//Other node with empty parent id also treated as root id, todo.
			node.ParentId = "-5"
			roots = append(roots, node)
		}
	}

	// Step 3: Assign levels using BFS
	queue := make([]*LshmOrgTreeNode, 0)

	// Set root level to 1
	for _, root := range roots {
		root.Level = 1
		root.OrgLevel = strconv.Itoa(root.Level)
		root.TreeCode = "/"
		queue = append(queue, root)
	}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for _, child := range current.Children {
			child.Level = current.Level + 1
			child.OrgLevel = strconv.Itoa(child.Level)
			child.TreeCode = fmt.Sprintf("%s%s/", current.TreeCode, current.ID)
			queue = append(queue, child)
		}
	}

	//flatten tree by bfs
	var result []LshmOrg
	var queue2 []*LshmOrgTreeNode

	// Initialize queue with roots
	for _, root := range roots {
		queue2 = append(queue2, root)
	}

	// BFS traversal
	for len(queue2) > 0 {
		current := queue2[0]
		queue2 = queue2[1:]
		result = append(result, current.LshmOrg)
		queue2 = append(queue2, current.Children...)
	}
	return result, nil

}
func RunGetMsg() error {
	log.Infof("Starting org sync....")
	startTime := time.Now()
	orgsTmp, _ := getLshmOrgInfo()
	var orgs []LshmOrg
	var err error

	if orgs, err = ProcessOrgs(orgsTmp); err != nil {
		log.Errorf("ProcessOrgs error: %s", err)
	}

	var minParOrgId = int64(100)
	var maxParOrgId = int64(-100)

	var minOrgId = int64(10000000)
	var maxOrgId = int64(-100)

	var minOrgLevel = int64(100)
	var maxOrgLevel = int64(-100)
	for _, org := range orgs {
		numId, err := strconv.ParseInt(org.ID, 10, 64)
		if err != nil {
			//log.Errorf("Error converting org.ID: %s to int64, %s, %+v", org.ID, err, org)
			continue
			//return err
			//numId =
		} else {
			minOrgId = If(numId < minOrgId, numId, minOrgId).(int64)
			maxOrgId = If(numId > maxOrgId, numId, maxOrgId).(int64)
		}
		numParentId, err := strconv.ParseInt(org.ParentId, 10, 64)
		if err != nil {
			log.Errorf("Error converting org.ParentId: %s to int64, %s, %+v", org.ParentId, err, org)
			continue
			//return err
		} else {
			minParOrgId = If(numParentId < minParOrgId, numParentId, minParOrgId).(int64)
			maxParOrgId = If(numParentId > maxParOrgId, numParentId, maxParOrgId).(int64)
		}
		orgLevel, err := strconv.ParseInt(org.OrgLevel, 10, 64)
		if err != nil {
			log.Errorf("Error converting org.OrgLevel: %s to int64, %s, %+v", org.OrgLevel, err, org)
			continue
		} else {
			minOrgLevel = If(orgLevel < minOrgLevel, orgLevel, minOrgLevel).(int64)
			maxOrgLevel = If(orgLevel > maxOrgLevel, orgLevel, maxOrgLevel).(int64)
		}

		if strings.Contains(org.TreeCode, "null") {
			log.Infof("org.TreeCode contains null: %+v", org)
		}

		if orgLevel == 14 {
			log.Infof("level 14: %+v", org)
		}
		if numId == 539753 {
			log.Infof("org 539753: %+v", org)
		}

		if numParentId < 0 {
			log.Infof("numParentId < 0: %+v", org)
		}
	}
	elapsedTime := time.Since(startTime)
	log.Infof("Finish org sync len:%d, %v", len(orgs), elapsedTime)
	log.Infof("Finish org stat: parent org:(%d, %d), org:(%d, %d), level:(%d, %d)", minParOrgId, maxParOrgId, minOrgId, maxOrgId, minOrgLevel, maxOrgLevel)
	return nil

	/*log.Infof("Starting person sync....")
	startTime2 := time.Now()
	b, _ := getLshmPersonInfo()
	elapsedTime2 := time.Since(startTime2)
	log.Infof("Finish person sync len:%d, %v", len(b), elapsedTime2)*/

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
