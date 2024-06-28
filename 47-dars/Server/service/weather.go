package service

import (
	"context"
	"log"
	pb "mymode/protos/weather/protos"
	"mymode/storage/postgres"
	"time"
)

type WeatherServer struct {
	pb.UnimplementedWeatherServer
	WDb postgres.NewWeather
}

func NewWeatherRepo(w postgres.NewWeather) *WeatherServer {
	return &WeatherServer{WDb: w}
}

func (W *WeatherServer) GetCurrentWeather(ctx context.Context, t *pb.Time) (*pb.Weather, error) {
	weather, err := W.WDb.GetCurrentWeather(t.Time)
	if err != nil {
		log.Println(err)
	}
	time.Sleep(2 * time.Second)
	return &weather, err
}

func (W *WeatherServer) GetWeatherForecast(ctx context.Context, day *pb.Day) (*pb.Weather, error) {
	weather, err := W.WDb.GetWeatherForecast(day.Day)
	if err != nil {
		log.Println(err)
	}
	time.Sleep(2 * time.Second)
	return &weather, err
}

func (W *WeatherServer) ReportWeatherCondition(ctx context.Context, weather *pb.Weather) (*pb.Status, error) {
	status, err := W.WDb.ReportWeatherCondition(*weather)
	if err != nil {
		log.Println(err)
	}
	time.Sleep(2 * time.Second)
	return &status, err
}
