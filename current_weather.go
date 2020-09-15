package openweathermap

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

var endpoint = "https://api.openweathermap.org/data/2.5/weather"
var unit = "metric"

type CurrentWeather struct {
	Coord    Coord     `json:"coord"`
	Weather  []Weather `json:"weather"`
	Base     string    `json:"base"`
	Main     Main      `json:"main"`
	Wind     Wind      `json:"wind"`
	Clouds   Clouds    `json:"clouds"`
	Rain     Rain      `json:"rain"`
	Snow     Snow      `json:"snow"`
	Dt       int       `json:"dt"`
	Sys      Sys       `json:"sys"`
	Timezone int       `json:"timezone"`
	ID       int       `json:"id"`
	Name     string    `json:"name"`
	Cod      int       `json:"cod"`
	Unit     string
	Apikey   string
}

func NewCurrentWeather(apikey, unit string) *CurrentWeather {
	return &CurrentWeather{Unit: unit, Apikey: apikey}
}

func (w *CurrentWeather) CurrentByCityID(cityID string) error {
	params := []string{
		"appid=" + w.Apikey,
		"id=" + cityID,
		"units=" + w.Unit,
	}

	url := fmt.Sprintf("%s?%s", endpoint, strings.Join(params, "&"))
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, &w)
	if err != nil {
		return err
	}
	return nil
}
