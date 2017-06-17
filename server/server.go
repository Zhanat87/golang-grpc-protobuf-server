package server

import (
	"github.com/Zhanat87/go/grpc"
	"github.com/Zhanat87/golang-grpc-protobuf-server/currency"
	"github.com/Zhanat87/golang-grpc-protobuf-server/weather"
	"github.com/Zhanat87/golang-grpc-protobuf-server/liveness"
)

// server is used to implement grpc.GrpcServiceServer.
type Server struct {}

func (s *Server) CheckGrpcServerLiveness(in *grpc.EmptyRequest, stream grpc.GrpcService_CheckGrpcServerLivenessServer) error {
	return liveness.CheckGrpcServerLiveness(in, stream)
}

func (s *Server) GetExchangeRates(in *grpc.EmptyRequest, stream grpc.GrpcService_GetExchangeRatesServer) error {
	return currency.GetExchangeRates(in, stream)
}

func (s *Server) GetWeatherInfo(in *grpc.WeatherRequest, stream grpc.GrpcService_GetWeatherInfoServer) error {
	return weather.GetWeatherInfo(in, stream)
}