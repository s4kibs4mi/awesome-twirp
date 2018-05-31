package api

import (
	"context"
	"github.com/s4kibs4mi/awesome-twirp/protobuf/pdefs"
	"github.com/twitchtv/twirp"
	"net/http"
	"os"
	"syscall"
	"time"
	"os/signal"
	"fmt"
	"github.com/go-chi/chi"
)

/**
 * := Coded with love by Sakib Sami on 28/5/18.
 * := root@sakib.ninja
 * := www.sakib.ninja
 * := Coffee : Dream : Code
 */

type Server struct {
}

func (s *Server) SayHello(ctx context.Context, req *pdefs.ReqHello) (*pdefs.ResHello, error) {
	return &pdefs.ResHello{
		GoodBye: &pdefs.GoodBye{
			UserID:  req.GetHello().GetUserID(),
			Message: req.GetHello().GetMessage(),
		},
	}, nil
}

func RunServer() {
	addr := ":7007"
	twirpServer := &Server{}
	handler := pdefs.NewAwesomeTwirpServiceServer(twirpServer, twirp.ChainHooks(NewAuthHook()))

	routes := chi.NewRouter()
	routes.Use(ForwardHttpHeaders)
	routes.Mount(pdefs.AwesomeTwirpServicePathPrefix, handler)

	server := http.Server{
		Addr:    addr,
		Handler: routes,
	}

	gracefulShutdown := make(chan os.Signal)
	signal.Notify(gracefulShutdown, syscall.SIGINT, syscall.SIGKILL)

	go func() {
		fmt.Println(fmt.Sprintf("server has been started on : %s", addr))
		if err := server.ListenAndServe(); err != nil {
			fmt.Errorf("erver failed to start : %v", err)
		}
	}()

	<-gracefulShutdown

	ctx, _ := context.WithTimeout(context.Background(), time.Second*5)
	server.Shutdown(ctx)
	fmt.Println("Server has been stopped gracefully")
}
