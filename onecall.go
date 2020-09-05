package openweathermap

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

type CurrentAndForecastWeather struct {
	Lat            float64    `json:"lat"`
	Lon            float64    `json:"lon"`
	Timezone       string     `json:"timezone"`
	TimezoneOffset int64      `json:"timezone_offset"`
	Current        Current    `json:"current"`
	Minutely       []Minutely `json:"minutely"`
	Hourly         []Hourly   `json:"hourly"`
	Daily          []Daily    `json:"daily"`
}

type Current struct {
	Dt         int64     `json:"dt"`
	Sunrise    int64     `json:"sunrise"`
	Sunset     int64     `json:"sunset"`
	Temp       float64   `json:"temp"`
	FeelsLike  float64   `json:"feels_like"`
	Pressure   int64     `json:"pressure"`
	Humidity   int64     `json:"humidity"`
	DewPoint   float64   `json:"dew_point"`
	Clouds     int64     `json:"clouds"`
	Uvi        float64   `json:"uvi"`
	Visibility int64     `json:"visibility"`
	WindSpeed  float64   `json:"wind_speed"`
	WindGust   float64   `json:"wind_gust"`
	WindDeg    int64     `json:"wind_deg"`
	Rain       Rain      `json:"rain"`
	Snow       Snow      `json:"snow"`
	Weather    []Weather `json:"weather"`
}

type Minutely struct {
	Dt            int64   `json:"dt"`
	Precipitation float64 `json:"precipitation"`
}

type Hourly struct {
	Dt         int64     `json:"dt"`
	Temp       float64   `json:"temp"`
	FeelsLike  float64   `json:"feels_like"`
	Pressure   int64     `json:"pressure"`
	Humidity   int64     `json:"humidity"`
	DewPoint   float64   `json:"dew_point"`
	Clouds     int64     `json:"clouds"`
	Visibility int64     `json:"visibility"`
	WindSpeed  float64   `json:"wind_speed"`
	WindGust   float64   `json:"wind_gust"`
	WindDeg    int64     `json:"wind_deg"`
	Rain       Rain      `json:"rain"`
	Snow       Snow      `json:"snow"`
	Weather    []Weather `json:"weather"`
}

type Daily struct {
	Dt         int64     `json:"dt"`
	Temp       Temp      `json:"temp"`
	FeelsLike  FeelsLike `json:"feels_like"`
	Pressure   int64     `json:"pressure"`
	Humidity   int64     `json:"humidity"`
	DewPoint   float64   `json:"dew_point"`
	Clouds     int64     `json:"clouds"`
	Visibility int64     `json:"visibility"`
	WindSpeed  float64   `json:"wind_speed"`
	WindGust   float64   `json:"wind_gust"`
	WindDeg    int64     `json:"wind_deg"`
	Rain       float64   `json:"rain"`
	Snow       Snow      `json:"snow"`
	Weather    []Weather `json:"weather"`
}

type TimeOfDay struct {
	Morn  float64 `json:"morn"`
	Day   float64 `json:"day"`
	Eve   float64 `json:"eve"`
	Night float64 `json:"night"`
}

type Temp struct {
	TimeOfDay
	Min float64 `json:"min"`
	Max float64 `json:"max"`
}

type FeelsLike struct {
	TimeOfDay
}

type OneCallApi struct {
	Config Config
	URL    string
	Params OneCallApiParams
}

type OneCallApiParams struct {
	Coord   Coordinates
	Exclude string
}

func (p OneCallApiParams) urlValues() url.Values {
	values := url.Values{}
	values.Set("lat", strconv.FormatFloat(p.Coord.Lat, 'f', -1, 64))
	values.Set("lon", strconv.FormatFloat(p.Coord.Lon, 'f', -1, 64))
	values.Set("exclude", p.Exclude)
	return values
}

func NewOneCallApi(config Config) *OneCallApi {
	params := OneCallApiParams{}
	return &OneCallApi{config, oneCallURL, params}
}

func (a *OneCallApi) CurrentAndForecastByCoordinates(coord Coordinates) (*CurrentAndForecastWeather, error) {
	a.Params.Coord = coord

	response := CurrentAndForecastWeather{}

	url := qqq(a.Config, a.URL, a.Params)
	fmt.Println(url)

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	fmt.Println(res)
	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}