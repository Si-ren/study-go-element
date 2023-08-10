package main

import (
	"fmt"
	"log"

	"github.com/hashicorp/consul/api"
)

func main() {
	config := api.DefaultConfig()
	config.Token = "d0391863-9328-6521-50ce-163348e728f8"
	// config.Token = "543f101c-3d99-3d4f-4e56-c7000017aabf"
	// 创建一个Consul客户端
	client, err := api.NewClient(config)

	if err != nil {
		log.Fatal(err)
	}

	// 创建查询选项
	queryOpts := &api.QueryOptions{}
	// 使用服务名从Consul中发现服务实例
	// services, _, err := client.Catalog().Service("web", "", queryOpts)
	// for _, service := range services {
	// 	fmt.Printf("Service ID: %s, Address: %s, Port: %d\n", service.ServiceID, service.Address, service.ServicePort)
	// }

	// 太离谱了。需要token带有web的read权限，仅write权限不行。不知道是不是官方文档写错了
	entries, _, err := client.Health().Service("web", "", false, queryOpts)
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println(entries)
	// 打印所有服务实例的地址和端口
	for _, entry := range entries {
		fmt.Printf("Service ID: %s, Address: %s, Port: %d\n", entry.Service.ID, entry.Service.Address, entry.Service.Port)
	}
}
