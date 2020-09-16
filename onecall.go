package openweathermap

import (
	"errors"
	"net/url"
)

type oneCallParams struct {
	coord   *Coord
	exclude string
}

func (o oneCallParams) urlValues() url.Values {
	values := url.Values{}

	if o.coord != nil {
		values.Set("lat", o.coord.Lat.String())
		values.Set("lon", o.coord.Lon.String())
	}

	if len(o.exclude) > 0 {
		values.Set("exclude", o.exclude)
	}

	return values
}

type OneCallOption func(*oneCallParams)

func ExcludeOption(exclude string) OneCallOption {
	return func(o *oneCallParams) {
		o.exclude = exclude
	}
}

type OneCall struct {
	*OwmAPI
}

func NewOneCall(config *Config) (*OneCall, error) {
	if !ValidateConfig(config) {
		return nil, errors.New("Invalid Config value")
	}

	return &OneCall{
		&OwmAPI{
			Config:   config,
			Endpoint: oneCallURL,
		},
	}, nil
}

func (o *OneCall) CurrentAndForecastByCoord(coord Coord, opts ...OneCallOption) (*CurrentAndForecastWeather, error) {
	if !ValidateCoord(coord) {
		return nil, errors.New("Invalid Coord value")
	}
	params := &oneCallParams{coord: &coord}

	for _, opt := range opts {
		opt(params)
	}

	o.Params = params

	weather := &CurrentAndForecastWeather{}
	err := o.get(weather)

	return weather, err
}
