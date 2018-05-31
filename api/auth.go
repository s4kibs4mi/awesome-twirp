package api

import (
	"github.com/twitchtv/twirp"
	"context"
	"fmt"
	"net/http"
)

/**
 * := Coded with love by Sakib Sami on 29/5/18.
 * := root@sakib.ninja
 * := www.sakib.ninja
 * := Coffee : Dream : Code
 */

func NewAuthHook() *twirp.ServerHooks {
	hook := &twirp.ServerHooks{}
	hook.RequestReceived = func(ctx context.Context) (context.Context, error) {
		fmt.Println("Auth Hook Request Received...")

		appKeyValues := ctx.Value("Appkey")
		if appKeyValues == nil {
			return ctx, twirp.NewError(twirp.Aborted, "Unauthorized request")
		}
		appKeys := appKeyValues.([]string)
		if len(appKeys) > 0 && appKeys[0] == "1" {
			return ctx, nil
		}
		return ctx, twirp.NewError(twirp.Aborted, "Unauthorized request")
	}
	return hook
}

// ForwardHttpHeaders forwards header from http request to twirp request
func ForwardHttpHeaders(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		for k, v := range r.Header {
			ctx = context.WithValue(ctx, k, v)
		}
		r = r.WithContext(ctx)
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
