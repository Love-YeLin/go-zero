package user

import (
	"context"
	"mail/internal/svc"
	"mail/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserLogic) GetUser(req *types.GetUserRequest) (resp *types.GetUserResponse, err error) {
	// todo: add your logic here and delete this line
	one, err := l.svcCtx.UC.UserModel.FindOne(l.ctx, 1)
	if err != nil {
		return nil, nil
	}
	return &types.GetUserResponse{Name: one.Name}, nil
}
