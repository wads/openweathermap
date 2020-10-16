package openweathermap

import (
	"compress/gzip"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type City struct {
	ID      int
	Name    string
	Country string
	Coord   Coord
}

type OWMCities struct {
	Cities []City
	Len    int
}

var cityList OWMCities

func NewOWMCityList() (*OWMCities, error) {
	if len(cityList.Cities) == 0 {
		raw, err := loadCityListJSON()
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(raw, &cityList)
		if err != nil {
			return nil, err
		}
	}
	return &cityList, nil
}

func loadCityListJSON() ([]byte, error) {
	resp, err := http.Get(cityListURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	zr, err := gzip.NewReader(resp.Body)
	if err != nil {
		return nil, err
	}
	defer zr.Close()

	return ioutil.ReadAll(zr)
}

func (o *OWMCities) UnmarshalJSON(data []byte) error {
	type tempCity struct {
		ID      float32
		Name    string
		Country string
		Coord   Coord
	}
	var cities []tempCity

	if err := json.Unmarshal(data, &cities); err != nil {
		return err
	}

	for _, city := range cities {

		o.Cities = append(
			o.Cities,
			City{ID: int(city.ID), Name: city.Name, Country: city.Country, Coord: city.Coord},
		)
	}
	o.Len = len(cities)

	return nil
}
