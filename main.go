package main

import (
	"flag"
	"fmt"
	"weather/geo"
	"weather/weather"
)

func main() {
	city := flag.String("city", "Saint-Petersburg", "Город пользователя")
	format := flag.Int("format", 1, "Формат вывода (1, 2, 3)")
	flag.Parse()

	geoData, err := geo.GetMyLocation(*city)
	if err != nil {
		fmt.Println(err.Error())
	}
	weatherData := weather.GetWeather(*geoData, *format)
	fmt.Println(weatherData)
}
