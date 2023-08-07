package main

// go install go.uber.org/mock/mockgen@latest
// mockgen -source=db.go -destination=db_mock.go -package=main
// 会生成一个db_mock.go文件

// 为什么需要mock，因为避免在单元测试中，会把一些持久化的数据修改。
// 例如 单元测试要对数据库增删改查，但这个是不太好的。

// db.go
// 使用注释方式为单个接口生成mock
//go:generate mockgen -source=db.go -destination=./mock_db_interface.go -package=main DB
type DB interface {
	Get(key string) (int, error)
}

func GetFromDB(db DB, key string) int {
	if value, err := db.Get(key); err == nil {
		return value
	}

	return -1
}
