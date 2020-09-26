package openweathermap

import (
	"errors"
	"net/url"
)

type currentParams struct {
	cityID   string
	cityName string
	state    string
	country  string
	coord    *Coord
	zipCode  string
}

func (c currentParams) urlValues() url.Values {
	values := url.Values{}

	if len(c.cityName) > 0 {
		name := c.cityName

		if len(c.state) > 0 {
			name += "," + c.state
		}

		if len(c.country) > 0 {
			name += "," + c.country
		}

		values.Set("q", name)
	}

	if len(c.cityID) > 0 {
		values.Set("id", c.cityID)
	}

	if c.coord != nil {
		values.Set("lat", c.coord.Lat.String())
		values.Set("lon", c.coord.Lon.String())
	}

	if len(c.zipCode) > 0 {
		values.Set("zip", c.zipCode)
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
	params := &currentParams{cityName: name}

	for _, opt := range opts {
		opt(params)
	}

	c.Params = params

	weather := &CurrentWeather{}
	err := c.get(weather)

	return weather, err
}

func (c *CurrentAPI) CurrentByCityID(id string) (*CurrentWeather, error) {
	params := &currentParams{cityID: id}

	c.Params = params

	weather := &CurrentWeather{}
	err := c.get(weather)

	return weather, err
}

func (c *CurrentAPI) CurrentByCoord(coord *Coord) (*CurrentWeather, error) {
	params := &currentParams{coord: coord}

	c.Params = params

	weather := &CurrentWeather{}
	err := c.get(weather)

	return weather, err
}

func (c *CurrentAPI) CurrentByZIPCode(zipCode string) (*CurrentWeather, error) {
	params := &currentParams{zipCode: zipCode}

	c.Params = params

	weather := &CurrentWeather{}
	err := c.get(weather)

	return weather, err
}
