package openweathermap

import (
	"net/url"
	"strconv"
)

var ModeList = map[string]string{
	"xml":  "xml",
	"html": "html",
}

var UnitsList = map[string]string{
	"imperial": "Fahrenheit",
	"metric":   "Celsius",
}

var LanguageCodeList = map[string]string{
	"af":    "Afrikaans",
	"al":    "Albanian",
	"ar":    "Arabic",
	"az":    "Azerbaijani",
	"bg":    "Bulgarian",
	"ca":    "Catalan",
	"cz":    "Czech",
	"da":    "Danish",
	"de":    "German",
	"el":    "Greek",
	"en":    "English",
	"eu":    "Basque",
	"fa":    "Persian (Farsi)",
	"fi":    "Finnish",
	"fr":    "French",
	"gl":    "Galician",
	"he":    "Hebrew",
	"hi":    "Hindi",
	"hr":    "Croatian",
	"hu":    "Hungarian",
	"id":    "Indonesian",
	"it":    "Italian",
	"ja":    "Japanese",
	"kr":    "Korean",
	"la":    "Latvian",
	"lt":    "Lithuanian",
	"mk":    "Macedonian",
	"no":    "Norwegian",
	"nl":    "Dutch",
	"pl":    "Polish",
	"pt":    "Portuguese",
	"pt_br": "Português Brasil",
	"ro":    "Romanian",
	"ru":    "Russian",
	"sv":    "Swedish",
	"se":    "Swedish",
	"sk":    "Slovak",
	"sl":    "Slovenian",
	"sp":    "Spanish",
	"es":    "Spanish",
	"sr":    "Serbian",
	"th":    "Thai",
	"tr":    "Turkish",
	"ua":    "Ukrainian",
	"uk":    "Ukrainian",
	"vi":    "Vietnamese",
	"zh_cn": "Chinese Simplified",
	"zh_tw": "Chinese Traditional",
	"zu":    "Zulu",
}

const (
	oneCallURL = "https://api.openweathermap.org/data/2.5/onecall"
)

type Config struct {
	APIKey string
	Mode   string
	Units  string
	Lang   string
}

type Option func(*Config)

func Mode(mode string) Option {
	return func(c *Config) {
		c.Mode = mode
	}
}

func Units(units string) Option {
	return func(c *Config) {
		c.Units = units
	}
}

func Lang(lang string) Option {
	return func(c *Config) {
		c.Lang = lang
	}
}

func NewConfig(apiKey string, opts ...Option) *Config {
	c := &Config{APIKey: apiKey}

	for _, opt := range opts {
		opt(c)
	}

	return c
}

func (c *Config) Valid() bool {
	return len(c.APIKey) > 0
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
	_, ok := ModeList[mode]
	return ok
}

func validUnits(units string) bool {
	_, ok := UnitsList[units]
	return ok
}

func validLang(lang string) bool {
	_, ok := LanguageCodeList[lang]
	return ok
}

func apiURL(config *Config, url string, params OptionParameters) string {
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
