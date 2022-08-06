package myhttp

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"log"
	"net/http"
)

func ErrorHandler(ctx context.Context, mux *runtime.ServeMux, marshaler runtime.Marshaler, writer http.ResponseWriter, request *http.Request, err error) {
	log.Println(err)
}
