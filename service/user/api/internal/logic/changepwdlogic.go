package logic

import (
	"book/common/errorx"
	"context"
	"github.com/tal-tech/go-zero/tools/goctl/model/sql/test/model"
	"strings"

	"book/service/user/api/internal/svc"
	"book/service/user/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type ChangepwdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChangepwdLogic(ctx context.Context, svcCtx *svc.ServiceContext) ChangepwdLogic {
	return ChangepwdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChangepwdLogic) Changepwd(req types.ChangePwdReq) error {
	// todo: add your logic here and delete this line
	if len(strings.TrimSpace(req.Password))==0{
		return errorx.NewDefaultError("参数错误")
	}
	//session信息读取用户ID
	//用户ID读取userInfo
	//修改用户密码，并保存到数据库
	//但是不知道怎么通过session获取用户ID艹
	userInfo,err :=l.svcCtx.UserModel.FindOneByName(req.Username)
	switch err {
	case nil:
	case model.ErrNotFound:
		return  errorx.NewDefaultError("用户名不存在")
	default:
		return err
	}
	userInfo.Password=req.Password
	err=l.svcCtx.UserModel.Update(*userInfo)
	if err != nil {
		return err
	}
	return nil
}
