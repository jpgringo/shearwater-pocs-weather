package weather

import (
	weatherTypes "github.com/jpgringo/shearwater-pocs-weather/type_definitions"
	"sync"
	"weather/internal/open_weather"
)

func GetWeather(config weatherTypes.Config, respCh chan weatherTypes.CurrentWeather, wg *sync.WaitGroup) {
	var weatherFunc weatherTypes.GetWeatherFunc
	switch config.Service {
	case "OpenWeather":
		weatherFunc = open_weather.GetOpenWeather
	}
	if weatherFunc == nil {
		close(respCh)
		wg.Done()
	} else {
		weatherFunc(config, respCh, wg)
	}

}
