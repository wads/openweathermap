package openweathermap

import (
	"errors"
	"net/url"
)

type OneCallParams struct {
	Coord   *Coordinates
	Exclude string
}

func (p OneCallParams) urlValues() url.Values {
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

func (a *OneCall) CurrentAndForecastByCoordinates(coord Coordinates, exclude string) (*CurrentAndForecastWeather, error) {
	if !coord.Validate() {
		return nil, errors.New("Invalid Coordinates value")
	}

	a.Params = OneCallParams{Coord: &coord, Exclude: exclude}

	weather := &CurrentAndForecastWeather{}
	err := a.get(weather)
	return weather, err
}
