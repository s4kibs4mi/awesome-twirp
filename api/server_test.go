package api

import (
	"testing"
	"github.com/s4kibs4mi/awesome-twirp/protobuf/pdefs"
	"net/http"
	"context"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/metadata"
	"fmt"
	"github.com/twitchtv/twirp"
)

/**
 * := Coded with love by Sakib Sami on 28/5/18.
 * := root@sakib.ninja
 * := www.sakib.ninja
 * := Coffee : Dream : Code
 */

func TestProtobufClient(t *testing.T) {
	client := pdefs.NewAwesomeTwirpServiceProtobufClient("http://localhost:7007", &http.Client{})
	header := make(http.Header)
	header.Set("UID", "uDRlDxQYbFVXarBvmTncBoWKcZKqrZTY")
	header.Set("AppKey", "1")
	ctx := context.Background()
	ctx, err := twirp.WithHTTPRequestHeaders(ctx, header)

	fmt.Println(ctx)

	resp, err := client.SayHello(ctx, &pdefs.ReqHello{
		Hello: &pdefs.Hello{
			UserID:  "1",
			Message: "Hello World",
		},
	})
	assert.Nil(t, err)
	assert.NotNil(t, resp)
	t.Log(resp.GetGoodBye())
}

func TestJSONClient(t *testing.T) {
	client := pdefs.NewAwesomeTwirpServiceJSONClient("http://localhost:7007", &http.Client{})
	ctx := metadata.NewIncomingContext(context.TODO(), metadata.Pairs("user_id", "2"))
	resp, err := client.SayHello(ctx, &pdefs.ReqHello{
		Hello: &pdefs.Hello{
			UserID:  "1",
			Message: "Hello World",
		},
	})
	assert.Nil(t, err)
	assert.NotNil(t, resp)
	t.Log(resp.GetGoodBye())
}
