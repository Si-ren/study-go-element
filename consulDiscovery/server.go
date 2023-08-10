package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/hashicorp/consul/api"
)

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	// 在此处编写健康检查逻辑
	// 检查数据库连接、依赖服务的可用性等

	// 假设我们检查到服务正常运行
	// 返回 HTTP 200 OK 状态码
	w.WriteHeader(http.StatusOK)
}

func main() {
	config := api.DefaultConfig()
	config.Token = "543f101c-3d99-3d4f-4e56-c7000017aabf"
	// 创建一个Consul客户端
	client, err := api.NewClient(config)
	if err != nil {
		log.Fatal(err)
	}

	// 创建一个新的服务实例
	registration := new(api.AgentServiceRegistration)
	registration.ID = "web01"
	registration.Name = "web"
	registration.Address = "172.28.208.1"
	registration.Port = 8080
	registration.Check = &api.AgentServiceCheck{
		HTTP:     "http://172.28.208.1:8080/health", // 健康检查的URL
		Interval: "10s",                             // 检查的频率
		Timeout:  "3s",                              // 检查的超时时间
	}

	// 注册服务到Consul
	err = client.Agent().ServiceRegister(registration)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Service registered successfully with Consul")
	// 注册健康检查的处理函数
	http.HandleFunc("/health", healthCheckHandler)

	// 启动服务
	fmt.Println("Starting server on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
