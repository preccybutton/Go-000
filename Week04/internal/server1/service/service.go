package service

import (
	"Go-000/Week04/api"
	"Go-000/Week04/internal/server1/biz"
	"context"
)

type GetTestDataService struct {
	api.UnimplementedGetTestDataServiceServer
	td *biz.TestDataUsercase
}

func NewGetTestDataService(data *biz.TestDataUsercase) api.GetTestDataServiceServer{
	return &GetTestDataService{td:data}
}



func (gtdsr *GetTestDataService) GetData(ctx context.Context, req *api.GetDataReq) (rsp *api.GetDataRsp, err error){
	td := new(biz.TestData)
	rsp = new(api.GetDataRsp)
	rsp.Data = td.Data()
	return
}