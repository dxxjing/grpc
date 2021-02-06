package service

import (
	"context"
	black "grpc-test/grpc-black/proto"
)

type BlackService struct {

}

func (b BlackService) List (ctx context.Context, req *black.ListReq) (*black.ListRsp, error) {
	_, _, keyword := req.Index, req.Count, req.Keyword
	if keyword != "" {
		//todo 搜索
	}
	//模拟从存储读出数据
	var infoList = []*black.Info{
		{
			Uid: 1,
			Name: "jdx",
		},
		{
			Uid: 2,
			Name: "tom",
		},
	}

	return &black.ListRsp{
		Info: infoList,
	},nil
}

func (b BlackService) Add (ctx context.Context, req *black.AddReq) (*black.AddRsp, error) {
	params := req.GetInfo()
	uid, name := params.GetUid(), params.GetName()
	//todo 写入存储
	_ = uid
	_ = name

	return &black.AddRsp{Id: 1111}, nil
}
