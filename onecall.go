package openweathermap

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

type OneCallApi struct {
	Config *Config
	URL    string
	Params OneCallApiParams
}

type OneCallApiParams struct {
	Coord   *Coordinates
	Exclude string
}

func (p OneCallApiParams) urlValues() url.Values {
	values := url.Values{}

	if p.Coord != nil {
		values.Set("lat", p.Coord.Lat.String())
		values.Set("lon", p.Coord.Lon.String())
	}

	if len(p.Exclude) > 0 {
		values.Set("exclude", p.Exclude)
	}

	return values
}

func NewOneCallApi(config *Config) (*OneCallApi, error) {
	if !validateConfig(config) {
		return nil, errors.New("Invalid Config value")
	}

	return &OneCallApi{
		Config: config,
		URL:    oneCallURL,
		Params: OneCallApiParams{},
	}, nil
}

func (a *OneCallApi) CurrentAndForecastByCoordinates(coord Coordinates) (*CurrentAndForecastWeather, error) {
	if !coord.Validate() {
		return nil, errors.New("Invalid Coordinates value")
	}
	a.Params.Coord = &coord

	url := apiURL(a.Config, a.URL, a.Params)
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return nil, err
	}

	response := CurrentAndForecastWeather{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
