package openweathermap

import "strconv"

type degree float64

func (d degree) String() string {
	return strconv.FormatFloat(float64(d), 'f', -1, 64)
}

type Coord struct {
	Lat degree `json:"lat"`
	Lon degree `json:"lon"`
}

func (c *Coord) Validate() bool {
	if c.Lat < -90 && c.Lat > 90 {
		return false
	}
	return true
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
