package cmd

import (
	"context"
	"ggclass_log_service/src/config"
	"ggclass_log_service/src/routes"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"
)

func server() *cobra.Command {
	return &cobra.Command{
		Use: "server",
		Run: func(cmd *cobra.Command, args []string) {
			var wg sync.WaitGroup
			wg.Add(1)
			go runGrpc(context.Background(), &wg)
			//
			go runGateway(context.Background(), &wg)

			wg.Wait()
			log.Println("shutdown")

		},
	}
}

func runGrpc(ctx context.Context, wg *sync.WaitGroup) {
	cfg := config.GetConfig()
	lis, err := net.Listen("tcp", ":"+cfg.GrpcPort)
	if err != nil {
		log.Println("fail to lis", err)
		return
	}

	s := grpc.NewServer()
	reflection.Register(s)

	routes.RegisterGrpcServer(s)

	idleConnectionsClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint
		s.GracefulStop()
		close(idleConnectionsClosed)
	}()
	s.Serve(lis)
	<-idleConnectionsClosed
	wg.Done()
	log.Println("shutdown grpc server")
}

func runGateway(ctx context.Context, wg *sync.WaitGroup) {

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	conn, err := grpc.DialContext(context.Background(), "0.0.0.0:"+config.GetConfig().GrpcPort, grpc.WithBlock(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return
	}

	mux := http.NewServeMux()

	//gwmux := runtime.NewServeMux(runtime.WithErrorHandler(myhttp.ErrorHandler))
	gwmux := runtime.NewServeMux()

	mux.Handle("/", gwmux)

	//err = folderpb.RegisterFolderServiceHandler(ctx, gwmux, conn)

	err = routes.RegisterGrpcGateway(ctx, gwmux, conn)

	if err != nil {
		return
	}

	gwServer := &http.Server{
		Addr:    ":" + config.GetConfig().HttpPort,
		Handler: mux,
	}

	//idleConnectionsClosed := make(chan struct{})
	//go func() {
	//	sigint := make(chan os.Signal, 1)
	//	signal.Notify(sigint, os.Interrupt)
	//	<-sigint
	//	if err := gwServer.Shutdown(context.Background()); err != nil {
	//		log.Printf("HTTP Server Shutdown Error: %v", err)
	//	}
	//	close(idleConnectionsClosed)
	//}()
	if err := gwServer.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("HTTP server ListenAndServe Error: %v", err)
	}

	//<-idleConnectionsClosed
	//wg.Done()

	log.Printf("shutdown gateway server")

	return
}
