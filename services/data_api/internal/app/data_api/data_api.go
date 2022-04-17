package dataapi

import (
	api "HackDavisFood/services/data_api/pb/food/data_api/v1alpha"

	"google.golang.org/grpc"
)

// Server data api grpc server
type Server struct {
	api.UnimplementedDataApiServiceServer
}

// StartDataAPIServer starts the data api grpc server
func StartDataAPIServer(s *grpc.Server) {
	parseConfig()
	setupCockroach()
	api.RegisterDataApiServiceServer(s, &Server{})
}
