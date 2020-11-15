package owm

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

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
	boxCityURL     = "api.openweathermap.org/data/2.5/box/city"
	cityListURL    = "bulk.openweathermap.org/sample/city.list.json.gz"
	currentURL     = "api.openweathermap.org/data/2.5/weather"
	findURL        = "api.openweathermap.org/data/2.5/find"
	groupURL       = "api.openweathermap.org/data/2.5/group"
	oneCallURL     = "api.openweathermap.org/data/2.5/onecall"
	oneCallPrevURL = "api.openweathermap.org/data/2.5/onecall/timemachine"
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
	Units  string
	Lang   string
}

type ConfigOption func(*Config)

func WithUnitsOption(units string) ConfigOption {
	return func(c *Config) {
		c.Units = units
	}
}

func WithLangOption(lang string) ConfigOption {
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
	setQuery(values url.Values)
}

type OwmAPI struct {
	Config *Config
	URL    string
	Params Params
}

func NewOwmAPI(config *Config, url string) *OwmAPI {
	return &OwmAPI{Config: config, URL: url}
}

func (a *OwmAPI) endpoint() string {
	if a.URL == "" {
		return ""
	}

	query := a.queryParams()
	if query == nil {
		return ""
	}

	return fmt.Sprintf("https://%s?%s", a.URL, query.Encode())
}

func (a *OwmAPI) queryParams() url.Values {
	if !ValidateAPIKey(a.Config.APIKey) {
		return nil
	}

	values := url.Values{}
	values.Set("appid", a.Config.APIKey)

	if ValidateUnits(a.Config.Units) {
		values.Set("units", a.Config.Units)
	}

	if ValidateLang(a.Config.Lang) {
		values.Set("lang", a.Config.Lang)
	}

	if a.Params != nil {
		a.Params.setQuery(values)
	}

	return values
}

func (a *OwmAPI) getAndLoad(dest interface{}) error {
	resp, err := a.get()
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return handleAPICallError(body)
	}

	return json.Unmarshal(body, &dest)
}

func (a *OwmAPI) get() (*http.Response, error) {
	endpoint := a.endpoint()
	if endpoint == "" {
		return nil, fmt.Errorf("URL is not set")
	}

	return http.Get(endpoint)
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
	return ValidateAPIKey(c.APIKey)
}

func ValidateAPIKey(apiKey string) bool {
	return len(apiKey) > 0
}

func ValidateUnits(units string) bool {
	_, ok := Units[units]
	return ok
}

func ValidateLang(lang string) bool {
	_, ok := Lang[lang]
	return ok
}
