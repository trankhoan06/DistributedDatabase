package middleware

import (
	"esim/component/tokenprovider"
	"esim/modules/user/storage"
)

type ModelMiddleware struct {
	authen *storage.SqlModel
	token  tokenprovider.TokenProvider
}

func NewModelMiddleware(au *storage.SqlModel, token tokenprovider.TokenProvider) *ModelMiddleware {
	return &ModelMiddleware{authen: au, token: token}
}
