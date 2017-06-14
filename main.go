package main

import (
	"log"
	"net"
	"encoding/xml"

	"google.golang.org/grpc"

	pb "github.com/Zhanat87/go/grpc/currency"

	"github.com/Zhanat87/go/util"
)

const (
	port = ":50051"
)

// server is used to implement currency.CurrencyServer.
type server struct {}

type RssFeed struct {
	XMLName xml.Name  `xml:"rss"`
	Items   []RssItem `xml:"channel>item"`
}

type RssItem struct {
	XMLName     xml.Name `xml:"item"`
	pb.Currency
}

func (s *server) GetExchangeRates(in *pb.EmptyRequest, stream pb.Currency_GetExchangeRatesServer) error {
	var data []*pb.ExchangeRateResponse
	currencies := []string{"GBP", "USD", "EUR", "RUB"}

	var rssFeed = &RssFeed{}
	xmlDoc := util.FetchURL("http://www.nationalbank.kz/rss/rates_all.xml")
	util.ParseXML(xmlDoc, &rssFeed)
	for _, item := range rssFeed.Items {
		if util.InArray(item.Title, currencies) {
			data = append(data, &pb.ExchangeRateResponse{
				Title: item.Title,
				PubDate: item.PubDate,
				Description: item.Description,
				Quant: item.Quant,
				Index: item.Index,
				Change: item.Change,
			})
		}
	}
	if err := stream.Send(&pb.ExchangeRatesResponse{Data: data}); err != nil {
		return err
	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// Creates a new gRPC server
	s := grpc.NewServer()
	pb.RegisterCurrencyServer(s, &server{})
	s.Serve(lis)
}