package openweathermap

import (
	"net/url"
)

var Mode = map[string]string{
	"xml":  "xml",
	"html": "html",
}

var Units = map[string]string{
	"imperial": "Fahrenheit",
	"metric":   "Celsius",
}

var Lang = map[string]string{
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

func ModeOption(mode string) Option {
	return func(c *Config) {
		c.Mode = mode
	}
}

func UnitsOption(units string) Option {
	return func(c *Config) {
		c.Units = units
	}
}

func LangOption(lang string) Option {
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

type Parameters interface {
	urlValues() url.Values
}

func validateConfig(c *Config) bool {
	return len(c.APIKey) > 0
}

func validateMode(mode string) bool {
	_, ok := Mode[mode]
	return ok
}

func validateUnits(units string) bool {
	_, ok := Units[units]
	return ok
}

func validateLang(lang string) bool {
	_, ok := Lang[lang]
	return ok
}

func apiURL(config *Config, url string, params Parameters) string {
	values := params.urlValues()
	values.Set("appid", config.APIKey)

	if validateMode(config.Mode) {
		values.Set("mode", config.Mode)
	}

	if validateUnits(config.Units) {
		values.Set("units", config.Units)
	}

	if validateLang(config.Lang) {
		values.Set("lang", config.Lang)
	}

	return url + "?" + values.Encode()
}
