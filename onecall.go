package openweathermap

import (
	"errors"
	"net/url"
)

type OneCallParams struct {
	coord   *Coord
	exclude string
}

func (o OneCallParams) urlValues() url.Values {
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

type OneCallOption func(*OneCallParams)

func ExcludeOption(exclude string) OneCallOption {
	return func(o *OneCallParams) {
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
			Params:   OneCallParams{},
		},
	}, nil
}

func (o *OneCall) CurrentAndForecastByCoord(coord Coord, opts ...OneCallOption) (*CurrentAndForecastWeather, error) {
	if !coord.Validate() {
		return nil, errors.New("Invalid Coord value")
	}
	params := &OneCallParams{coord: &coord}

	for _, opt := range opts {
		opt(params)
	}

	o.Params = params

	weather := &CurrentAndForecastWeather{}
	err := o.get(weather)

	return weather, err
}
