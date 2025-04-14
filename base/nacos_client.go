package base

import (
	"fmt"
	"time"

	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	_ "github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
)

func getService(client naming_client.INamingClient, param vo.GetServiceParam) {
	service, err := client.GetService(param)
	if err != nil {
		panic("GetService failed!" + err.Error())
	}
	fmt.Printf("GetService,param:%+v, result:%+v \n\n", param, service)
}

func selectAllInstances(client naming_client.INamingClient, param vo.SelectAllInstancesParam) {
	instances, err := client.SelectAllInstances(param)
	if err != nil {
		panic("SelectAllInstances failed!" + err.Error())
	}
	fmt.Printf("SelectAllInstance,param:%+v, result:%+v \n\n", param, instances)
}

func selectInstances(client naming_client.INamingClient, param vo.SelectInstancesParam) {
	instances, err := client.SelectInstances(param)
	if err != nil {
		panic("SelectInstances failed!" + err.Error())
	}
	fmt.Printf("SelectInstances,param:%+v, result:%+v \n\n", param, instances)
}

func NacosClientTest() {
	sc := []constant.ServerConfig{
		*constant.NewServerConfig("console.nacos.io", 8848, constant.WithContextPath("/nacos")),
	}

	//create ClientConfig
	cc := *constant.NewClientConfig(
		constant.WithNamespaceId(""),
		constant.WithTimeoutMs(5000),
		constant.WithNotLoadCacheAtStart(true),
		constant.WithLogDir("./log"),
		constant.WithCacheDir("./cache"),
		constant.WithLogLevel("debug"),
	)

	// create naming client
	client, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)

	time.Sleep(1 * time.Second)

	//Get service with serviceName, groupName , clusters
	getService(client, vo.GetServiceParam{
		ServiceName: "client1",
		GroupName:   "DEFAULT_GROUP",
		Clusters:    []string{"DEFAULT"},
	})

	selectInstances(client, vo.SelectInstancesParam{
		ServiceName: "client1",
		GroupName:   "DEFAULT_GROUP",
		Clusters:    []string{"DEFAULT"},
		HealthyOnly: true,
	})

	if err != nil {
		panic(err)
	}
}
