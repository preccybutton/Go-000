package biz

//数据库model
type TestData struct {
	data string
}

func (td *TestData) SetData(data string) error{
	td.data = data
	return nil
}

func (td *TestData) Data() string{
	if td != nil {
		return "td.data"
	}
	return ""
}


//持久化接口
type TestDataRepo interface{
	GetTestData(*TestData)
}

func NewTestDataUsercase(repo TestDataRepo) *TestDataUsercase{
	return &TestDataUsercase{repo:repo}
}

//业务struct
type TestDataUsercase struct {
	repo TestDataRepo
}

//业务的方法
func (td *TestDataUsercase) Get(t *TestData) {
	td.repo.GetTestData(t)
}
