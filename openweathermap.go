package owm

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

var Mode = map[string]string{
	"xml":  "xml",
	"html": "html",
}

var Units = map[string]string{
	"imperial": "Temperature in Fahrenheit and wind speed in miles/hour",
	"metric":   "Temperature in Celsius and wind speed in meter/sec",
	"standard": "Temperature in Kelvin and wind speed in meter/sec",
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
	boxCityURL  = "https://api.openweathermap.org/data/2.5/box/city"
	cityListURL = "http://bulk.openweathermap.org/sample/city.list.json.gz"
	currentURL  = "https://api.openweathermap.org/data/2.5/weather"
	findURL     = "https://api.openweathermap.org/data/2.5/find"
	groupURL    = "https://api.openweathermap.org/data/2.5/group"
	oneCallURL  = "https://api.openweathermap.org/data/2.5/onecall"
)

type APICallError struct {
	COD     string `json:"cod"`
	Message string `json:"message"`
}

func (a APICallError) Error() string {
	return fmt.Sprintf("%s, (cod=%s)", a.Message, a.COD)
}

type Config struct {
	APIKey string
	Mode   string
	Units  string
	Lang   string
}

type ConfigOption func(*Config)

func ModeOption(mode string) ConfigOption {
	return func(c *Config) {
		c.Mode = mode
	}
}

func UnitsOption(units string) ConfigOption {
	return func(c *Config) {
		c.Units = units
	}
}

func LangOption(lang string) ConfigOption {
	return func(c *Config) {
		c.Lang = lang
	}
}

func NewConfig(apiKey string, opts ...ConfigOption) *Config {
	c := &Config{APIKey: apiKey}

	for _, opt := range opts {
		opt(c)
	}

	return c
}

type Params interface {
	urlValues() url.Values
}

type OwmAPI struct {
	Config   *Config
	Endpoint string
	Params   Params
}

func NewOwmAPI(config *Config, endpoint string) *OwmAPI {
	api := &OwmAPI{Config: config, Endpoint: endpoint}

	return api
}

func (a *OwmAPI) apiURL() string {
	var values url.Values
	if a.Params != nil {
		values = a.Params.urlValues()
	} else {
		values = url.Values{}
	}

	values.Set("appid", a.Config.APIKey)

	if ValidateMode(a.Config.Mode) {
		values.Set("mode", a.Config.Mode)
	}

	if ValidateUnits(a.Config.Units) {
		values.Set("units", a.Config.Units)
	}

	if ValidateLang(a.Config.Lang) {
		values.Set("lang", a.Config.Lang)
	}

	query := values.Encode()

	return fmt.Sprintf("%s?%s", a.Endpoint, query)
}

func (a *OwmAPI) get(dest interface{}) error {
	res, err := http.Get(a.apiURL())
	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return err
	}

	if res.StatusCode != 200 {
		return handleAPICallError(body)
	}

	return json.Unmarshal(body, &dest)
}

func handleAPICallError(respBody []byte) error {
	apiCallError := APICallError{}
	err := json.Unmarshal(respBody, &apiCallError)
	if err != nil {
		return err
	}
	return apiCallError
}

func ValidateConfig(c *Config) bool {
	return len(c.APIKey) > 0
}

func ValidateMode(mode string) bool {
	_, ok := Mode[mode]
	return ok
}

func ValidateUnits(units string) bool {
	_, ok := Units[units]
	return ok
}

func ValidateLang(lang string) bool {
	_, ok := Lang[lang]
	return ok
}
