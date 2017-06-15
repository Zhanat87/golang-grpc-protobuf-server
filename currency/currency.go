package currency

import (
	"encoding/xml"

	"github.com/Zhanat87/go/grpc"

	"github.com/Zhanat87/go/util"
)

type RssFeed struct {
	XMLName xml.Name  `xml:"rss"`
	Items   []RssItem `xml:"channel>item"`
}

type RssItem struct {
	XMLName     xml.Name `xml:"item"`
	Currency
}

type Currency struct {
	Title       string  `json:"title" xml:"title"`
	PubDate     string  `json:"pubDate" xml:"pubDate"`
	Description float32 `json:"description" xml:"description"`
	Quant       int32   `json:"quant" xml:"quant"`
	Index       string  `json:"index" xml:"index"`
	Change      float32 `json:"change" xml:"change"`
}

func GetExchangeRates(in *grpc.EmptyRequest, stream grpc.GrpcService_GetExchangeRatesServer) error {
	var data []*grpc.ExchangeRateResponse
	currencies := []string{"GBP", "USD", "EUR", "RUB"}

	var rssFeed = &RssFeed{}
	xmlDoc := util.FetchURL("http://www.nationalbank.kz/rss/rates_all.xml")
	util.ParseXML(xmlDoc, &rssFeed)
	for _, item := range rssFeed.Items {
		if util.InArray(item.Title, currencies) {
			data = append(data, &grpc.ExchangeRateResponse{
				Title: item.Title,
				PubDate: item.PubDate,
				Description: item.Description,
				Quant: item.Quant,
				Index: item.Index,
				Change: item.Change,
			})
		}
	}
	if err := stream.Send(&grpc.ExchangeRatesResponse{Data: data}); err != nil {
		return err
	}
	return nil
}