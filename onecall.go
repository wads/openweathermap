package owm

import (
	"errors"
	"fmt"
	"net/url"
)

type OneCallAPI struct {
	*OwmAPI
}

func NewOneCallAPI(config *Config) (*OneCallAPI, error) {
	if !ValidateConfig(config) {
		return nil, errors.New("Invalid Config")
	}

	return &OneCallAPI{NewOwmAPI(config, "")}, nil
}

type oneCallParams struct {
	coord   *Coord
	exclude string
}

func (o *oneCallParams) urlValues() url.Values {
	values := url.Values{}

	values.Set("lat", o.coord.Lat.String())
	values.Set("lon", o.coord.Lon.String())

	if o.exclude != "" {
		values.Set("exclude", o.exclude)
	}

	return values
}

type OneCallOption func(*oneCallParams)

func WithExcludeOption(exclude string) OneCallOption {
	return func(o *oneCallParams) {
		o.exclude = "exclude"
	}
}

func (o *OneCallAPI) GetWeather(coord *Coord, opts ...OneCallOption) (*CurrentAndForecastWeather, error) {
	if !ValidateCoord(coord) {
		return nil, errors.New("Invalid Coord value")
	}

	params := &oneCallParams{coord: coord}
	for _, opt := range opts {
		opt(params)
	}

	o.Params = params
	o.URL = oneCallURL

	weather := &CurrentAndForecastWeather{}
	err := o.get(weather)

	return weather, err
}

type oneCallPrevParams struct {
	coord *Coord
	dt    int64
}

func (o *oneCallPrevParams) urlValues() url.Values {
	values := url.Values{}

	values.Set("lat", o.coord.Lat.String())
	values.Set("lon", o.coord.Lon.String())
	values.Set("dt", fmt.Sprintf("%d", o.dt))

	return values
}

func (o *OneCallAPI) GetPrevWeather(coord *Coord, dt int64) (*PreviousWeather, error) {
	if !ValidateCoord(coord) {
		return nil, errors.New("Invalid Coord value")
	}

	params := &oneCallPrevParams{coord: coord, dt: dt}

	o.Params = params
	o.URL = oneCallPrevURL

	weather := &PreviousWeather{}
	err := o.get(weather)

	return weather, err
}
