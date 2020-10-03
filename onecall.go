package openweathermap

import (
	"errors"
	"net/url"
)

type OneCallAPI struct {
	*OwmAPI
}

func NewOneCallAPI(config *Config) (*OneCallAPI, error) {
	if !ValidateConfig(config) {
		return nil, errors.New("Invalid Config")
	}

	return &OneCallAPI{NewOwmAPI(config, oneCallURL)}, nil
}

type oneCallParams struct {
	coord   *Coord
	exclude string
}

func (o *oneCallParams) urlValues() url.Values {
	values := url.Values{}

	if o.coord != nil {
		values.Set("lat", o.coord.Lat.String())
		values.Set("lon", o.coord.Lon.String())
	}

	if o.exclude != "" {
		values.Set("exclude", o.exclude)
	}

	return values
}

type OneCallOption func(*oneCallParams)

func ExcludeOption(exclude string) OneCallOption {
	return func(o *oneCallParams) {
		o.exclude = "exclude"
	}
}

func (o *OneCallAPI) GetCurrentAndForecast(coord *Coord, opts ...OneCallOption) (*CurrentAndForecastWeather, error) {
	if !ValidateCoord(coord) {
		return nil, errors.New("Invalid Coord value")
	}

	params := &oneCallParams{coord: coord}

	for _, opt := range opts {
		opt(params)
	}

	o.Params = params

	weather := &CurrentAndForecastWeather{}
	err := o.get(weather)

	return weather, err
}
