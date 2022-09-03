package tool

import (
	"quora/app/user/cmd/rpc/pb"
)

func ParseToken(token string) {

}

func GenerateToken(in *pb.GenerateTokenReq, expire int64) (*pb.GenerateTokenResp, error) {
	return nil, nil
}

func getJwtToken(secretKey string, iat, seconds int64, in *pb.GenerateTokenReq) (string, error) {
	return "", nil

}
