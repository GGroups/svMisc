package cosurl

import (
	"context"
	"errors"

	"github.com/go-kit/kit/endpoint"
)

type MallInfo struct {
	ChannelId string `json:"channelId"`
}

type ObjUrlRequest struct {
	MallInfo MallInfo `json:"mallInfo"`
	Type     string   `json:"type"`
	UrlObjs  []UrlObj `json:"urlObjs"`
}

type ObjUrlResponse struct {
	UrlObjs []UrlObj `json:"datas"`
	Msg     string   `json:"msg"`
	RetCode string   `json:"retode"`
}

func MakeObjUrlEndPoint(sv IObjUrl) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		r, ok := request.(ObjUrlRequest)
		if !ok {
			return ObjUrlResponse{}, nil
		}
		if r.Type != "wx" {
			return nil, errors.New(INPUTE_RROR + `not "wx"`)
		}
		return ObjUrlResponse{UrlObjs: sv.GetObjUrlItems(r), Msg: "ok", RetCode: "0"}, nil
	}
}
