package coupons

import (
	"context"
	"errors"

	"github.com/go-kit/kit/endpoint"
)

type AuthTXRequest struct {
	Type string `json:"type"`
}

type AuthTXResponse struct {
	AuthTX  []AuthTX `json:"datas"`
	Msg     string   `json:"msg"`
	RetCode string   `json:"retode"`
}

func MakeAuthTXEndPoint(sv IAuthTX) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		r, ok := request.(AuthTXRequest)
		if !ok {
			return AuthTXResponse{}, nil
		}
		if r.Type != "wx" {
			return nil, errors.New(INPUTE_RROR + `not "wx"`)
		}
		return AuthTXResponse{AuthTX: sv.GetAuthTXItems(r.Type), Msg: "ok", RetCode: "0"}, nil
	}
}
