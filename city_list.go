package openweathermap

import (
	"compress/gzip"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type OpenWeatherMapCity struct {
	City []City
}

type City struct {
	Id      int         `json:"id"`
	Name    string      `json:"name"`
	Country string      `json:"country"`
	Coord   Coordinates `json:"coord"`
}

const cityListURL = "http://bulk.openweathermap.org/sample/city.list.json.gz"

var openWeatherMapCity *OpenWeatherMapCity

func (c Coordinates) GetLatitude() float64 {
	return c.Lat
}

func (c Coordinates) GetLongitude() float64 {
	return c.Lon
}

func NewOpenWeatherMapCity() (*OpenWeatherMapCity, error) {
	if openWeatherMapCity == nil {
		raw, err := loadCityListJSON()
		if err != nil {
			return nil, err
		}

		var cities []City
		err = json.Unmarshal(raw, &cities)
		if err != nil {
			return nil, err
		}
		openWeatherMapCity = &OpenWeatherMapCity{cities}
	}
	return openWeatherMapCity, nil
}

func loadCityListJSON() ([]byte, error) {
	resp, err := http.Get(cityListURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	zr, err := gzip.NewReader(resp.Body)
	if err != nil {
		return nil, err
	}
	defer zr.Close()

	return ioutil.ReadAll(zr)
}
