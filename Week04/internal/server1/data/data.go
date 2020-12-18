package data

import "Go-000/Week04/internal/server1/biz"

func NewTestDataRepo() biz.TestDataRepo{
	return new(testDataRepo)
}

type testDataRepo struct {
}

func (tdr *testDataRepo) GetTestData(*biz.TestData){

}