package owm

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type degree float64

func (d degree) String() string {
	return strconv.FormatFloat(float64(d), 'f', -1, 64)
}

type Coord struct {
	Lat degree
	Lon degree
}

func ValidateCoord(coord *Coord) bool {
	if coord.Lat < -90 && coord.Lat > 90 {
		return false
	}

	if coord.Lon < -180 && coord.Lon > 180 {
		return false
	}

	return true
}

type BoundingBox struct {
	LatTop    degree
	LatBottom degree
	LonLeft   degree
	LonRight  degree
	Zoom      int
}

func (b *BoundingBox) String() string {
	return fmt.Sprintf("%f,%f,%f,%f,%d", b.LonLeft, b.LatBottom, b.LonRight, b.LatTop, b.Zoom)
}

func ValidateBoundingBox(bbox *BoundingBox) bool {
	if bbox.LatTop < -90 && bbox.LatTop > 90 {
		return false
	}

	if bbox.LatBottom < -90 && bbox.LatBottom > 90 {
		return false
	}

	if bbox.LonLeft < -180 && bbox.LonLeft > 180 {
		return false
	}

	if bbox.LonRight < -180 && bbox.LonRight > 180 {
		return false
	}

	return true
}

type Weather struct {
	ID          int
	Main        string
	Description string
	Icon        string
}

type Main struct {
	Temp      float32
	Pressure  int
	Humidity  int
	TempMin   float32 `json:"temp_min"`
	TempMax   float32 `json:"temp_max"`
	SeaLevel  float32 `json:"sea_level"`
	GrndLevel float32 `json:"grnd_level"`
}

type Wind struct {
	Speed float32
	Deg   int
}

type Clouds struct {
	All   int
	Today int
}

type Rain struct {
	OneH   float32 `json:"1h"`
	ThreeH float32 `json:"3h"`
}

type Snow struct {
	OneH   float32 `json:"1h"`
	ThreeH float32 `json:"3h"`
}

type Sys struct {
	Type    int
	ID      int
	Message float32
	Country string
	Sunrise int
	Sunset  int
}

type CurrentAndForecastWeather struct {
	Lat            float32
	Lon            float32
	Timezone       string
	TimezoneOffset int `json:"timezone_offset"`
	Current        Current
	Minutely       []Minutely
	Hourly         []Hourly
	Daily          []Daily
}

type PreviousWeather struct {
	Lat            float32
	Lon            float32
	Timezone       string
	TimezoneOffset int `json:"timezone_offset"`
	Current        Current
	Hourly         []Hourly
}

type Current struct {
	Dt         int
	Sunrise    int
	Sunset     int
	Temp       float32
	FeelsLike  float32 `json:"feels_like"`
	Pressure   int
	Humidity   int
	DewPoint   float32 `json:"dew_point"`
	Clouds     int
	Uvi        float32
	Visibility int
	WindSpeed  float32 `json:"wind_speed"`
	WindGust   float32 `json:"wind_gust"`
	WindDeg    int     `json:"wind_deg"`
	Rain       Rain
	Snow       Snow
	Weather    []Weather
}

type Minutely struct {
	Dt            int
	Precipitation float32
}

type Hourly struct {
	Dt         int
	Temp       float32
	FeelsLike  float32 `json:"feels_like"`
	Pressure   int
	Humidity   int
	DewPoint   float32 `json:"dew_point"`
	Clouds     int
	Visibility int
	WindSpeed  float32 `json:"wind_speed"`
	WindGust   float32 `json:"wind_gust"`
	WindDeg    int     `json:"wind_deg"`
	Rain       Rain
	Snow       Snow
	Weather    []Weather
}

type Daily struct {
	Dt         int
	Temp       Temp
	FeelsLike  FeelsLike `json:"feels_like"`
	Pressure   int
	Humidity   int
	DewPoint   float32 `json:"dew_point"`
	Clouds     int
	Visibility int
	WindSpeed  float32 `json:"wind_speed"`
	WindGust   float32 `json:"wind_gust"`
	WindDeg    int     `json:"wind_deg"`
	Rain       float32
	Snow       Snow
	Weather    []Weather
}

type TimeOfDay struct {
	Morn  float32
	Day   float32
	Eve   float32
	Night float32
}

type Temp struct {
	TimeOfDay
	Min float32
	Max float32
}

type FeelsLike struct {
	TimeOfDay
}

type CurrentWeather struct {
	Coord      Coord
	Weather    []Weather
	Base       string
	Main       Main
	Visibility int
	Wind       Wind
	Clouds     Clouds
	Rain       Rain
	Snow       Snow
	Dt         int
	Sys        Sys
	Timezone   int
	ID         int
	Name       string
	Cod        int
}

type CurrentCitiesWeather struct {
	Cod      int
	Calctime float32
	Cnt      int
	List     []CurrentWeather
}

func (c *CurrentCitiesWeather) UnmarshalJSON(data []byte) error {
	type Alias CurrentCitiesWeather
	a := &struct {
		Cod json.RawMessage
		*Alias
	}{
		Alias: (*Alias)(c),
	}
	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	var cod string
	if err := json.Unmarshal(a.Cod, &cod); err == nil {
		i, err := strconv.Atoi(cod)
		if err != nil {
			return err
		}
		c.Cod = i
	}

	return nil
}
