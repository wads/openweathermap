package openweathermap

import (
	"errors"
	"net/url"
)

type OneCallOption func(url.Values)

func ExcludeOption(exclude string) OneCallOption {
	return func(v url.Values) {
		v.Set("exclude", exclude)
	}
}

type OneCallAPI struct {
	*OwmAPI
}

func NewOneCallAPI(config *Config) (*OneCallAPI, error) {
	if !ValidateConfig(config) {
		return nil, errors.New("Invalid Config value")
	}

	return &OneCallAPI{NewOwmAPI(config, oneCallURL)}, nil
}

func (o *OneCallAPI) CurrentAndForecast(coord *Coord, opts ...OneCallOption) (*CurrentAndForecastWeather, error) {
	if !ValidateCoord(coord) {
		return nil, errors.New("Invalid Coord value")
	}

	o.Params.Set("lat", coord.Lat.String())
	o.Params.Set("lon", coord.Lon.String())

	for _, opt := range opts {
		opt(o.Params)
	}

	weather := &CurrentAndForecastWeather{}
	err := o.get(weather)

	return weather, err
}
