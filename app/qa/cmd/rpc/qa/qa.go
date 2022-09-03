// Code generated by goctl. DO NOT EDIT!
// Source: qa.proto

package qa

import (
	"context"

	"quora/app/qa/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AddAnswerReq           = pb.AddAnswerReq
	AddAnswerResp          = pb.AddAnswerResp
	AddQuestionReq         = pb.AddQuestionReq
	AddQuestionResp        = pb.AddQuestionResp
	Answer                 = pb.Answer
	DelAnswerReq           = pb.DelAnswerReq
	DelAnswerResp          = pb.DelAnswerResp
	DelQuestionReq         = pb.DelQuestionReq
	DelQuestionResp        = pb.DelQuestionResp
	GetAnswerByIdReq       = pb.GetAnswerByIdReq
	GetAnswerByIdResp      = pb.GetAnswerByIdResp
	GetAnswerListByPidReq  = pb.GetAnswerListByPidReq
	GetAnswerListByPidResp = pb.GetAnswerListByPidResp
	GetAnswerListByUidReq  = pb.GetAnswerListByUidReq
	GetAnswerListByUidResp = pb.GetAnswerListByUidResp
	GetQuestionByIdReq     = pb.GetQuestionByIdReq
	GetQuestionByIdResp    = pb.GetQuestionByIdResp
	Question               = pb.Question
	SearchAnswerReq        = pb.SearchAnswerReq
	SearchAnswerResp       = pb.SearchAnswerResp
	SearchQuestionReq      = pb.SearchQuestionReq
	SearchQuestionResp     = pb.SearchQuestionResp
	UpdateAnswerReq        = pb.UpdateAnswerReq
	UpdateAnswerResp       = pb.UpdateAnswerResp
	UpdateQuestionReq      = pb.UpdateQuestionReq
	UpdateQuestionResp     = pb.UpdateQuestionResp

	Qa interface {
		// -----------------------answer-----------------------
		AddAnswer(ctx context.Context, in *AddAnswerReq, opts ...grpc.CallOption) (*AddAnswerResp, error)
		UpdateAnswer(ctx context.Context, in *UpdateAnswerReq, opts ...grpc.CallOption) (*UpdateAnswerResp, error)
		DelAnswer(ctx context.Context, in *DelAnswerReq, opts ...grpc.CallOption) (*DelAnswerResp, error)
		GetAnswerById(ctx context.Context, in *GetAnswerByIdReq, opts ...grpc.CallOption) (*GetAnswerByIdResp, error)
		SearchAnswer(ctx context.Context, in *SearchAnswerReq, opts ...grpc.CallOption) (*SearchAnswerResp, error)
		GetAnswerListByQid(ctx context.Context, in *GetAnswerListByPidReq, opts ...grpc.CallOption) (*GetAnswerListByPidResp, error)
		GetAnswerListByUid(ctx context.Context, in *GetAnswerListByUidReq, opts ...grpc.CallOption) (*GetAnswerListByUidResp, error)
		// -----------------------question-----------------------
		AddQuestion(ctx context.Context, in *AddQuestionReq, opts ...grpc.CallOption) (*AddQuestionResp, error)
		UpdateQuestion(ctx context.Context, in *UpdateQuestionReq, opts ...grpc.CallOption) (*UpdateQuestionResp, error)
		DelQuestion(ctx context.Context, in *DelQuestionReq, opts ...grpc.CallOption) (*DelQuestionResp, error)
		GetQuestionById(ctx context.Context, in *GetQuestionByIdReq, opts ...grpc.CallOption) (*GetQuestionByIdResp, error)
		SearchQuestion(ctx context.Context, in *SearchQuestionReq, opts ...grpc.CallOption) (*SearchQuestionResp, error)
	}

	defaultQa struct {
		cli zrpc.Client
	}
)

func NewQa(cli zrpc.Client) Qa {
	return &defaultQa{
		cli: cli,
	}
}

// -----------------------answer-----------------------
func (m *defaultQa) AddAnswer(ctx context.Context, in *AddAnswerReq, opts ...grpc.CallOption) (*AddAnswerResp, error) {
	client := pb.NewQaClient(m.cli.Conn())
	return client.AddAnswer(ctx, in, opts...)
}

func (m *defaultQa) UpdateAnswer(ctx context.Context, in *UpdateAnswerReq, opts ...grpc.CallOption) (*UpdateAnswerResp, error) {
	client := pb.NewQaClient(m.cli.Conn())
	return client.UpdateAnswer(ctx, in, opts...)
}

func (m *defaultQa) DelAnswer(ctx context.Context, in *DelAnswerReq, opts ...grpc.CallOption) (*DelAnswerResp, error) {
	client := pb.NewQaClient(m.cli.Conn())
	return client.DelAnswer(ctx, in, opts...)
}

func (m *defaultQa) GetAnswerById(ctx context.Context, in *GetAnswerByIdReq, opts ...grpc.CallOption) (*GetAnswerByIdResp, error) {
	client := pb.NewQaClient(m.cli.Conn())
	return client.GetAnswerById(ctx, in, opts...)
}

func (m *defaultQa) SearchAnswer(ctx context.Context, in *SearchAnswerReq, opts ...grpc.CallOption) (*SearchAnswerResp, error) {
	client := pb.NewQaClient(m.cli.Conn())
	return client.SearchAnswer(ctx, in, opts...)
}

func (m *defaultQa) GetAnswerListByQid(ctx context.Context, in *GetAnswerListByPidReq, opts ...grpc.CallOption) (*GetAnswerListByPidResp, error) {
	client := pb.NewQaClient(m.cli.Conn())
	return client.GetAnswerListByQid(ctx, in, opts...)
}

func (m *defaultQa) GetAnswerListByUid(ctx context.Context, in *GetAnswerListByUidReq, opts ...grpc.CallOption) (*GetAnswerListByUidResp, error) {
	client := pb.NewQaClient(m.cli.Conn())
	return client.GetAnswerListByUid(ctx, in, opts...)
}

// -----------------------question-----------------------
func (m *defaultQa) AddQuestion(ctx context.Context, in *AddQuestionReq, opts ...grpc.CallOption) (*AddQuestionResp, error) {
	client := pb.NewQaClient(m.cli.Conn())
	return client.AddQuestion(ctx, in, opts...)
}

func (m *defaultQa) UpdateQuestion(ctx context.Context, in *UpdateQuestionReq, opts ...grpc.CallOption) (*UpdateQuestionResp, error) {
	client := pb.NewQaClient(m.cli.Conn())
	return client.UpdateQuestion(ctx, in, opts...)
}

func (m *defaultQa) DelQuestion(ctx context.Context, in *DelQuestionReq, opts ...grpc.CallOption) (*DelQuestionResp, error) {
	client := pb.NewQaClient(m.cli.Conn())
	return client.DelQuestion(ctx, in, opts...)
}

func (m *defaultQa) GetQuestionById(ctx context.Context, in *GetQuestionByIdReq, opts ...grpc.CallOption) (*GetQuestionByIdResp, error) {
	client := pb.NewQaClient(m.cli.Conn())
	return client.GetQuestionById(ctx, in, opts...)
}

func (m *defaultQa) SearchQuestion(ctx context.Context, in *SearchQuestionReq, opts ...grpc.CallOption) (*SearchQuestionResp, error) {
	client := pb.NewQaClient(m.cli.Conn())
	return client.SearchQuestion(ctx, in, opts...)
}
