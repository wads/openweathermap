package openweathermap

import (
	"errors"
	"net/url"
)

type currentParams struct {
	name    string
	state   string
	country string
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

	return &CurrentAPI{
		&OwmAPI{
			Config:   config,
			Endpoint: currentURL,
		},
	}, nil
}

func (c *CurrentAPI) CurrentByCityName(name string, opts ...CurrentOption) (*CurrentWeather, error) {
	params := &currentParams{name: name}

	for _, opt := range opts {
		opt(params)
	}

	c.Params = params

	weather := &CurrentWeather{}
	err := c.get(weather)

	return weather, err
}
