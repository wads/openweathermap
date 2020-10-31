package owm

import (
	"errors"
	"net/url"
	"strconv"
	"strings"
)

type CurrentCitiesWeatherAPI struct {
	*OwmAPI
}

func NewCurrentCitiesWeatherAPI(config *Config) (*CurrentCitiesWeatherAPI, error) {
	if !ValidateConfig(config) {
		return nil, errors.New("Invalid Config")
	}

	return &CurrentCitiesWeatherAPI{NewOwmAPI(config, "")}, nil
}

type rectZoneParams struct {
	bbox *BoundingBox
}

func (r *rectZoneParams) urlValues() url.Values {
	values := url.Values{}

	if r.bbox != nil {
		values.Set("bbox", r.bbox.String())
	}

	return values
}

func (s *CurrentCitiesWeatherAPI) GetWithinRectZone(bbox *BoundingBox) (*CurrentCitiesWeather, error) {
	if !ValidateBoundingBox(bbox) {
		return nil, errors.New("Invalid BoundingBox")
	}

	s.Params = &rectZoneParams{bbox: bbox}
	s.URL = boxCityURL

	weather := &CurrentCitiesWeather{}
	err := s.get(weather)

	return weather, err
}

func (s *CurrentCitiesWeatherAPI) GetInCircle(coord *Coord) (*CurrentCitiesWeather, error) {
	if !ValidateCoord(coord) {
		return nil, errors.New("Invalid Coord")
	}

	s.Params = &coordParams{coord: coord}
	s.URL = findURL

	weather := &CurrentCitiesWeather{}
	err := s.get(weather)

	return weather, err
}

type cityIDsParams struct {
	ids []int
}

func (c *cityIDsParams) urlValues() url.Values {
	values := url.Values{}

	if len(c.ids) > 0 {
		ids := make([]string, len(c.ids))
		for i := range c.ids {
			ids[i] = strconv.Itoa(c.ids[i])
		}
		values.Set("id", strings.Join(ids, ","))
	}

	return values
}

func (s *CurrentCitiesWeatherAPI) GetByCityIDs(ids []int) (*CurrentCitiesWeather, error) {
	s.Params = &cityIDsParams{ids: ids}
	s.URL = groupURL

	weather := &CurrentCitiesWeather{}
	err := s.get(weather)

	return weather, err
}
