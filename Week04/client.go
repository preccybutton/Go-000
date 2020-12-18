package main

import (
	"Go-000/Week04/api"
	"context"
	"fmt"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

func main(){
	conn, err := grpc.Dial(":8000", grpc.WithInsecure())
	if err != nil{
		errors.Wrapf(err, "连接server失败 地址: %v", ":8000")
		fmt.Println(err)
		return
	}
	defer conn.Close()

	client := api.NewGetTestDataServiceClient(conn)
	res, _ := client.GetData(context.Background(),&api.GetDataReq{Data:"aaa"})
	fmt.Println(res)

}
