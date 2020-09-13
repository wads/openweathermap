package openweathermap

import (
	"errors"
	"net/url"
)

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

type OneCallApi struct {
	*OwmAPI
}

func NewOneCallApi(config *Config) (*OneCallApi, error) {
	if !ValidateConfig(config) {
		return nil, errors.New("Invalid Config value")
	}

	return &OneCallApi{
		&OwmAPI{
			Config:   config,
			Endpoint: oneCallURL,
			Params:   OneCallApiParams{},
		},
	}, nil
}

func (a *OneCallApi) CurrentAndForecastByCoordinates(coord Coordinates, exclude string) (*CurrentAndForecastWeather, error) {
	if !coord.Validate() {
		return nil, errors.New("Invalid Coordinates value")
	}

	a.Params = OneCallApiParams{Coord: &coord, Exclude: exclude}

	weather := &CurrentAndForecastWeather{}
	err := a.get(weather)
	return weather, err
}
