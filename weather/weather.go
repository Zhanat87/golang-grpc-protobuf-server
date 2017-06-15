package weather

import (
	"fmt"

	"github.com/Zhanat87/go/grpc"
	"github.com/Zhanat87/go/util"
)

const OPEN_WEATHER_MAP_API_KEY = "8b745f60c241e69915fbe7a9b4ab96b9"

func GetWeatherInfo(in *grpc.WeatherRequest, stream grpc.GrpcService_GetWeatherInfoServer) error {
	weatherInfoJson := util.FetchURL(fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?lat=%f&lon=%f&units=metric&appid=%s",
		in.Latitude, in.Longitude, OPEN_WEATHER_MAP_API_KEY))
	// https://stackoverflow.com/questions/13593519/how-do-i-parse-an-inner-field-in-a-nested-json-object-in-golang
	decoded := util.ParseJsonData(weatherInfoJson)
	// pull out the parents object
	main := decoded["main"].(map[string]interface{})
	weatherInfo := &grpc.WeatherResponse{
		Temp:     main["temp"].(float64),
		Pressure: main["pressure"].(float64),
		Humidity: main["humidity"].(float64),
	}

	if err := stream.Send(weatherInfo); err != nil {
		return err
	}
	return nil
}
