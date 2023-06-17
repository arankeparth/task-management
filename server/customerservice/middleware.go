package customerApi

import (
	"context"
	"plantrip-backend/server/spec/customerspec"

	"github.com/go-kit/kit/endpoint"
)

func (ch *CustomerHandler) JWTMiddleware() endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			req := request.(customerspec.GetOffersRequest)
			err := ch.VerifyJwt(req.TokenString, req.PublicKey)
			if err != nil {
				return nil, err
			}
			return next(ctx, request)
		}
	}
}
