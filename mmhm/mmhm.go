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
type JobInfo struct {
	Id              string `json:"id"`
	UserId          string `json:"userId"`
	BizDomain       string `json:"bizDomain"`
	BelongBrand     string `json:"belongBrand"`
	JobPostCode     string `json:"jobPostCode"`
	JobPostName     string `json:"jobPostName"`
	JobStatus       string `json:"jobStatus"`
	MainWork        string `json:"mainWork"`
	IsAgentMgm      string `json:"isAgentMgm"`
	JobPostId       string `json:"jobPostId"`
	StartWorkDay    string `json:"startWorkDay"`
	EndWorkDay      string `json:"endWorkDay"`
	IsActuralMgm    string `json:"isActuralMgm"`
	OrgName         string `json:"orgName"`
	OrgId           string `json:"orgId"`
	Code            string `json:"code"`
	AddrProvince    string `json:"addrProvince"`
	AddrCity        string `json:"addrCity"`
	AddrArea        string `json:"addrArea"`
	BelongBrandName string `json:"belongBrandName"`
}
type LshmPerson struct {
	UserId           string    `json:"userId"`
	RealName         string    `json:"realName"`
	Phone            string    `json:"phone"`
	ThirdUid         string    `json:"thirdUid"`
	Email            string    `json:"email"`
	UserStatus       string    `json:"userStatus"`
	Sex              string    `json:"sex"`
	DirectLeaderId   int64     `json:"directLeaderId"`
	DirectLeaderName string    `json:"directLeaderName"`
	Code             string    `json:"code"`
	OrgId            string    `json:"orgId"`
	JobInfoList      []JobInfo `json:"jobInfoList"`
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

	var defaultPageSize int64 = 400
	reqSourceCode := "2011"
	bizDomains := []string{"10", "20"}
	reqUrl := fmt.Sprintf("%s%s", lshmDevUrlPrefix, lshmOrgAPIUri)
	log.Infof("getLshmOrgInfo:%s", reqUrl)
	h := map[string]string{
		"accept":         "*/*",
		lshmAppHeaderKey: lshmDevAppHeaderValue,
	}
	allOrg := make([]LshmOrg, 0)

	for _, bizDomain := range bizDomains {
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

				if err := json.Unmarshal(respData, &resp); err != nil {
					log.Errorf("getLshmOrgInfo, error: %s", err)
					return nil, err
				} else {
					allOrg = append(allOrg, resp.LshmOrgRespData.Records...)
					if resp.LshmOrgRespData.TotalPages > 1 {
						for i := int64(2); i <= resp.LshmOrgRespData.TotalPages; i++ {
							q.PageNo = i
							qm, _ = json.Marshal(q)
							if respData, err := CallAPI("POST", reqUrl, h, qm); err != nil {
								log.Error("getLshmOrgInfo, %d, error: %s", i, err)
								return nil, err
							} else {
								//str := string(respData)
								log.Infof("Get lshm org info success, data len: %d,%d", i, len(respData))
								if err := json.Unmarshal(respData, &resp); err != nil {
									log.Errorf("getLshmOrgInfo, error: %s", err)
									return nil, err
								} else {
									allOrg = append(allOrg, resp.LshmOrgRespData.Records...)
								}

							}
						}

					}

					log.Infof("Get lshm org info from domain %s success, total: %d.", bizDomain, len(allOrg))
				}

			}
		}
	}
	return allOrg, nil

}

func getLshmPersonInfo() ([]LshmPerson, error) {

	var defaultPageSize int64 = 400
	reqSourceCode := "2011"
	bizDomains := []string{"10", "20"}
	reqUrl := fmt.Sprintf("%s%s", lshmDevUrlPrefix, lshmPersonAPIUri)
	log.Infof("getLshmPersonInfo:%s", reqUrl)
	h := map[string]string{
		"accept":         "*/*",
		lshmAppHeaderKey: lshmDevAppHeaderValue,
	}
	allPerson := make([]LshmPerson, 0)

	for _, bizDomain := range bizDomains {
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
					log.Infof("Get lshm person info from domain %s success, total: %d.", bizDomain, len(allPerson))

				}

			}
		}
	}
	//Retrive person organization info from job info list, we use the org from the job with newest work day as the person's organization.
	for m, person := range allPerson {
		newestJobIndex := 0
		newestWorkTime, _ := time.Parse(time.DateTime, "1900-01-01 00:00:00")
		for i, job := range person.JobInfoList {
			//log.Infof("Get person job info date:%s---%s", person.JobInfoList[0].StartWorkDay, person.JobInfoList[1].StartWorkDay)
			//log.Infof("get date:%s-%s, main work:%s, org:%s,%s", job.StartWorkDay, job.EndWorkDay, job.MainWork, job.OrgId, job.OrgName)
			curWorkTime, _ := time.Parse(time.DateTime, job.StartWorkDay)
			if curWorkTime.After(newestWorkTime) {
				newestWorkTime = curWorkTime
				newestJobIndex = i
			}
		}
		//log.Infof("get newest job:%s-%s, main work:%s, org:%s,%s", person.JobInfoList[newestJobIndex].StartWorkDay, person.JobInfoList[newestJobIndex].EndWorkDay, person.JobInfoList[newestJobIndex].MainWork, person.JobInfoList[newestJobIndex].OrgId, person.JobInfoList[newestJobIndex].OrgName)
		allPerson[m].OrgId = person.JobInfoList[newestJobIndex].OrgId
		//log.Infof("Get person info:%+v", person)
	}
	return allPerson, nil
}

func getLshmPersonInfodddd() ([]LshmPerson, error) {

	var defaultPageSize int64 = 500
	reqSourceCode := "2011"
	bizDomains := []string{"10", "20"}
	//bizDomain := "20"
	reqUrl := fmt.Sprintf("%s%s", lshmDevUrlPrefix, lshmPersonAPIUri)
	log.Infof("getLshmPersonInfo:%s", reqUrl)
	h := map[string]string{
		"accept":         "*/*",
		lshmAppHeaderKey: lshmDevAppHeaderValue,
	}
	allPerson := make([]LshmPerson, 0)
	for _, bizDomain := range bizDomains {
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

				if err := json.Unmarshal(respData, &resp); err != nil {
					log.Errorf("getLshmPersonInfo, error: %s", err)
					return nil, err
				} else {
					allPerson = append(allPerson, resp.LshmPersonRespData.Records...)

					//log.Infof("Get person info:%+v", resp.LshmPersonRespData.Records)
					for m, person := range resp.LshmPersonRespData.Records {
						newestJobIndex := 0
						newestWorkTime, _ := time.Parse(time.DateTime, "2000-01-01 00:00:00")
						for i, job := range person.JobInfoList {
							//log.Infof("Get person job info date:%s---%s", person.JobInfoList[0].StartWorkDay, person.JobInfoList[1].StartWorkDay)
							//log.Infof("get date:%s-%s, main work:%s, org:%s,%s", job.StartWorkDay, job.EndWorkDay, job.MainWork, job.OrgId, job.OrgName)
							curWorkTime, _ := time.Parse(time.DateTime, job.StartWorkDay)
							if curWorkTime.After(newestWorkTime) {
								newestWorkTime = curWorkTime
								newestJobIndex = i
							}
						}
						//log.Infof("get newest job:%s-%s, main work:%s, org:%s,%s", person.JobInfoList[newestJobIndex].StartWorkDay, person.JobInfoList[newestJobIndex].EndWorkDay, person.JobInfoList[newestJobIndex].MainWork, person.JobInfoList[newestJobIndex].OrgId, person.JobInfoList[newestJobIndex].OrgName)
						resp.LshmPersonRespData.Records[m].OrgId = person.JobInfoList[newestJobIndex].OrgId
						//log.Infof("Get person info:%+v", person)
						//log.Infof("Get person in %d info:%+v", m, resp.LshmPersonRespData.Records[m])
					}
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
					log.Infof("Get person size %d in domain %s", len(allPerson), bizDomain)
				}

			}
		}
	}
	return allPerson, nil

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
				node.ParentId = "-2"
				roots = append(roots, node)
			}
		} else {
			//Other nodes with empty parent id also treated as root id, todo.
			node.ParentId = "-5"
			roots = append(roots, node)
		}
	}

	// Step 3: Assign levels using BFS
	queue := make([]*LshmOrgTreeNode, 0)

	// Set root level to 1
	maxChildrenNum := 0
	realRootIndex := 0
	for i, root := range roots {
		root.Level = 1
		root.OrgLevel = strconv.Itoa(root.Level)
		root.TreeCode = "/"
		if root.ParentId == "-2" && maxChildrenNum < len(root.Children) {
			realRootIndex = i
			maxChildrenNum = len(root.Children)
		}
		queue = append(queue, root)
	}

	log.Infof("Get real roots info: %d, %d, %s, %s", realRootIndex, maxChildrenNum, queue[realRootIndex].ID, queue[realRootIndex].OrgName)
	queue[realRootIndex].ParentId = "-1"

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

	log.Infof("orgs len--------:%d, %d", len(orgsTmp), len(orgs))
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
		if org.ID == "900910966" {
			log.Infof("org 900910966: %+v", org)
		}
	}
	elapsedTime := time.Since(startTime)
	log.Infof("Finish org sync len:%d, %v", len(orgs), elapsedTime)
	log.Infof("Finish org stat: parent org:(%d, %d), org:(%d, %d), level:(%d, %d), time:%v", minParOrgId, maxParOrgId, minOrgId, maxOrgId, minOrgLevel, maxOrgLevel, elapsedTime)
	return nil

	/*log.Infof("Starting person sync....")
	startTime2 := time.Now()
	persons, _ := getLshmPersonInfo()
	for i, p := range persons {
		if len(p.JobInfoList) == 0 {
			log.Infof("No info list %d:::%+v", i, p)
		}
	}
	elapsedTime2 := time.Since(startTime2)
	log.Infof("Finish person sync len:%d, %v", len(persons), elapsedTime2)
	return nil*/

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
