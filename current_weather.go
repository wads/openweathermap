package openweathermap

import (
	"errors"
	"net/url"
)

type CurrentOption func(*currentParams)

func StateOption(state string) CurrentOption {
	return func(c *currentParams) {
		c.state = state
	}
}

func CountryOption(country string) CurrentOption {
	return func(c *currentParams) {
		c.country = country
	}
}

type CurrentAPI struct {
	*OwmAPI
}

func NewCurrentAPI(config *Config) (*CurrentAPI, error) {
	if !ValidateConfig(config) {
		return nil, errors.New("Invalid Config value")
	}

	return &CurrentAPI{NewOwmAPI(config, currentURL)}, nil
}

type currentParams struct {
	country string
	name    string
	state   string
}

func (c currentParams) urlValues() url.Values {
	values := url.Values{}

	if len(c.name) > 0 {
		name := c.name

		if len(c.state) > 0 {
			name += "," + c.state
		}

		if len(c.country) > 0 {
			name += "," + c.country
		}

		values.Set("q", name)
	}

	return values
}

func (c *CurrentAPI) CurrentByCityName(name string, opts ...CurrentOption) (*CurrentWeather, error) {
	params := &currentParams{name: name}

	for _, opt := range opts {
		opt(params)
	}

	c.Params = params.urlValues()

	weather := &CurrentWeather{}
	err := c.get(weather)

	return weather, err
}

func (c *CurrentAPI) CurrentByCityID(id string) (*CurrentWeather, error) {
	c.Params.Set("id", id)

	weather := &CurrentWeather{}
	err := c.get(weather)

	return weather, err
}

func (c *CurrentAPI) CurrentByCoord(coord *Coord) (*CurrentWeather, error) {
	if !ValidateCoord(coord) {
		return nil, errors.New("Invalid Coord value")
	}

	c.Params.Set("lat", coord.Lat.String())
	c.Params.Set("lon", coord.Lon.String())

	weather := &CurrentWeather{}
	err := c.get(weather)

	return weather, err
}

func (c *CurrentAPI) CurrentByZIPCode(zipCode string) (*CurrentWeather, error) {
	c.Params.Set("zip", zipCode)

	weather := &CurrentWeather{}
	err := c.get(weather)

	return weather, err
}
