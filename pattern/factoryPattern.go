package pattern

type factory struct {
	Name string
	Age  int
	use  bool
}

//因为factory是小写首字母,所以只能在pattern包内使用
//可以使用工厂模式来解决

func NewFactory(str string, age int, use bool) *factory {
	return &factory{
		Name: str,
		Age:  age,
		use:  use,
	}
}

// GetUse
// 如果只需要获取factory中的use变量,那么需要使用一个方法返回才能外部获取
func (f *factory) GetUse() bool {
	return f.use
}
