package openweathermap

import (
	"net/url"
	"strconv"
)

const (
	oneCallURL = "https://api.openweathermap.org/data/2.5/onecall"
)

type Config struct {
	APIKey string
	Mode   string
	Units  string
	Lang   string
}

type degree float64

func (d degree) String() string {
	return strconv.FormatFloat(float64(d), 'f', -1, 64)
}

type Coordinates struct {
	Lat degree `json:"lat"`
	Lon degree `json:"lon"`
}

func (c *Coordinates) Valid() bool {
	if c.Lat < -90 && c.Lat > 90 {
		return false
	}
	return true
}

type OptionParameters interface {
	urlValues() url.Values
}

type Weather struct {
	ID          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type Main struct {
	Temp      float64 `json:"temp"`
	Pressure  int     `json:"pressure"`
	Humidity  int     `json:"humidity"`
	TempMin   float64 `json:"temp_min"`
	TempMax   float64 `json:"temp_max"`
	SeaLevel  float64 `json:"sea_level"`
	GrndLevel float64 `json:"grnd_level"`
}

type Wind struct {
	Speed float64 `json:"speed"`
	Deg   int     `json:"deg"`
}

type Clouds struct {
	All int `json:"all"`
}

type Rain struct {
	OneH   float64 `json:"1h"`
	ThreeH float64 `json:"3h"`
}

type Snow struct {
	OneH   float64 `json:"1h"`
	ThreeH float64 `json:"3h"`
}

type Sys struct {
	Type    int     `json:"type"`
	ID      int     `json:"id"`
	Message float64 `json:"message"`
	Country string  `json:"country"`
	Sunrise int     `json:"sunrise"`
	Sunset  int     `json:"sunset"`
}

func validMode(mode string) bool {
	return mode != ""
}

func validUnits(units string) bool {
	return units != ""
}

func validLang(lang string) bool {
	return lang != ""
}

func apiURL(config Config, url string, params OptionParameters) string {
	values := params.urlValues()
	values.Set("appid", config.APIKey)

	if validMode(config.Mode) {
		values.Set("mode", config.Mode)
	}

	if validUnits(config.Units) {
		values.Set("units", config.Units)
	}

	if validLang(config.Lang) {
		values.Set("lang", config.Lang)
	}

	return url + "?" + values.Encode()
}
