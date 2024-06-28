package postgres

import (
	"database/sql"
	"fmt"
	pb "mymode/protos/weather/protos"
)

type NewWeather struct {
	Db *sql.DB
}

func NewWeatherRepo(db *sql.DB) *NewWeather {
	return &NewWeather{Db: db}
}

func (W *NewWeather) GetCurrentWeather(time string) (pb.Weather, error) {
	weather := pb.Weather{}
	err := W.Db.QueryRow(`SELECT * FROM WeatherTime where time = $1`, time).Scan(
		&weather.T, &weather.Weath, &weather.Temperature, &weather.Wind, &weather.Damp)
	fmt.Println(weather)
	return weather, err
}

func (W *NewWeather) GetWeatherForecast(day string) (pb.Weather, error) {
	weather := pb.Weather{}
	err := W.Db.QueryRow(`SELECT * FROM WeatherDay where day = $1`, day).Scan(
		&weather.Day, &weather.Weath, &weather.Temperature, &weather.Wind, &weather.Damp)
	return weather, err
}

func (W *NewWeather) ReportWeatherCondition(weather pb.Weather) (pb.Status, error) {
	_, err := W.Db.Exec(`INSERT INTO WeatherTime(time, weather, temp, damp, wind) VALUES($1, $2, $3, $4, $5)`,
		weather.T, weather.Weath, weather.Temperature, weather.Damp, weather.Wind)
	if err != nil || weather.T == "" || weather.Weath == "" {
		return pb.Status{Status: false}, err
	}
	return pb.Status{Status: true}, nil
}
