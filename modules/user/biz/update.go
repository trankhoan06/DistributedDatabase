package biz

import (
	"context"
	"esim/common"
	"esim/modules/user/model"
)

func (biz *UserCommon) NewUpdateUser(ctx context.Context, data *model.UpdateUser) error {
	if err := biz.user.UpdateUser(ctx, data); err != nil {
		return common.ErrUpdate(err)
	}
	return nil
}
