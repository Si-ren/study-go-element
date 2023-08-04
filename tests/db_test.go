package main

import (
	"errors"
	"testing"

	gomock "go.uber.org/mock/gomock"
)

func TestGetFromDB(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish() // 断言 DB.Get() 方法是否被调用

	m := NewMockDB(ctrl)
	m.EXPECT().Get(gomock.Eq("Tom")).Return(100, errors.New("not exist"))

	// 这里是个测试用例，当Get()返回err时，v等于-1，如果不等于-1，那么报错
	if v := GetFromDB(m, "Tom"); v != -1 {
		t.Fatal("expected -1, but got", v)
	}
	// m.EXPECT().Get(gomock.Eq("Test")).DoAndReturn(func(DB &MockDB{}, string) int {

	// })
}
