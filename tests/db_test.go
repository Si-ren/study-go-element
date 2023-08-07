package main

import (
	"testing"

	gomock "go.uber.org/mock/gomock"
)

func TestGetFromDB(t *testing.T) {
	// mock的固定写法
	ctrl := gomock.NewController(t)
	defer ctrl.Finish() // 断言 DB.Get() 方法是否被调用

	// 打桩，为接口指定传参并且返回值
	// 还有很多参数选择
	// m.EXPECT().Get(gomock.Any()).Return(60, errors.New("test Any"))
	// m.EXPECT().Get(gomock.Not("Apple")).Return(60, errors.New("test Any"))
	// m.EXPECT().Get(gomock.Nil()).Return(60, errors.New("test Any"))
	// 返回值也有两种Return,DoAndReturn  DoAndReturn()
	// m.EXPECT().Get(gomock.Eq("Tom")).Return(100, errors.New("not exist"))
	// 还可以加入调用次数Times(3)
	m := NewMockDB(ctrl)
	m.EXPECT().Get(gomock.Eq("Tom")).DoAndReturn(func(key string) (int, error) {
		if key == "Tom" {
			return -1, nil
		}
		return 0, nil
	}).Times(3)
	// 这里是个测试用例，当Get()返回err时，v等于-1，如果不等于-1，那么报错
	if v := GetFromDB(m, "Tom"); v != -1 {
		t.Fatal("expected -1, but got", v)
	}
	// mock 还支持顺序
}
