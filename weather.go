package weather

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type CurrentCondition struct {
	Cloudcover       string
	Humidity         string
	Observation_Time string
	PrecipMM         string
	Pressure         string
	Temp_C           string
	Temp_F           string
	Visibility       string
	WeatherCode      string
	WeatherDesc      []map[string]string
	WeatherIconUrl   []map[string]string
	WindDir16Point   string
	WindDirDegree    string
	WindSpeedKmph    string
	WindSpeedMiles   string
}

type PastWeather struct {
	Date           string
	PrecipMM       string
	TempMaxC       string
	TempMaxF       string
	TempMinC       string
	TempMinF       string
	WeatherCode    string
	WeatherDesc    []map[string]string
	WeatherIconUrl []map[string]string
	WinDir16Point  string
	WindDirDegree  string
	WindDirection  string
	WindSpeedKmph  string
	WindSpeedMiles string
}

type WeatherRequest struct {
	Query string
	Type  string
}

type Weather struct {
	Current_Condition []CurrentCondition
	Request           []WeatherRequest
	Weather           []PastWeather
}

type weatherContainer struct {
	Data Weather
}

var WeatherTypes = map[int]string{
	395: "Moderate or heavy snow in area with thunder",
	392: "Patchy light snow in area with thunder",
	389: "Moderate or heavy rain in area with thunder",
	386: "Patchy light rain in area with thunder",
	377: "Moderate or heavy showers of ice pellets",
	374: "Light showers of ice pellets",
	371: "Moderate or heavy snow showers",
	368: "Light snow showers",
	365: "Moderate or heavy sleet showers",
	362: "Light sleet showers",
	359: "Torrential rain shower",
	356: "Moderate or heavy rain shower",
	353: "Light rain shower",
	350: "Ice pellets",
	338: "Heavy snow",
	335: "Patchy heavy snow",
	332: "Moderate snow",
	329: "Patchy moderate snow",
	326: "Light snow",
	323: "Patchy light snow",
	320: "Moderate or heavy sleet",
	317: "Light sleet",
	314: "Moderate or Heavy freezing rain",
	311: "Light freezing rain",
	308: "Heavy rain",
	305: "Heavy rain at times",
	302: "Moderate rain",
	299: "Moderate rain at times",
	296: "Light rain",
	293: "Patchy light rain",
	284: "Heavy freezing drizzle",
	281: "Freezing drizzle",
	266: "Light drizzle",
	263: "Patchy light drizzle",
	260: "Freezing fog",
	248: "Fog",
	230: "Blizzard",
	227: "Blowing snow",
	200: "Thundery outbreaks in nearby",
	185: "Patchy freezing drizzle nearby",
	182: "Patchy sleet nearby",
	179: "Patchy snow nearby",
	176: "Patchy rain nearby",
	143: "Mist",
	122: "Overcast",
	119: "Cloudy",
	116: "Partly Cloudy",
	113: "Clear/Sunny",
}

var WeatherCategories = map[int]string{
	395: "snow",
	392: "snow",
	389: "rain",
	386: "rain",
	377: "snow",
	374: "snow",
	371: "snow",
	368: "snow",
	365: "rain",
	362: "rain",
	359: "rain",
	356: "rain",
	353: "rain",
	350: "snow",
	338: "snow",
	335: "snow",
	332: "snow",
	329: "snow",
	326: "snow",
	323: "snow",
	320: "rain",
	317: "rain",
	314: "rain",
	311: "rain",
	308: "rain",
	305: "rain",
	302: "rain",
	299: "rain",
	296: "rain",
	293: "rain",
	284: "rain",
	281: "rain",
	266: "rain",
	263: "rain",
	260: "covered",
	248: "covered",
	230: "snow",
	227: "snow",
	200: "rain",
	185: "rain",
	182: "rain",
	179: "snow",
	176: "rain",
	143: "covered",
	122: "covered",
	119: "covered",
	116: "covered",
	113: "clear",
}

func WeatherType(code int) (string, string) {
	t, ok := WeatherTypes[code]
	if ok != true {
		t = "Unknown"
	}

	c, ok := WeatherCategories[code]
	if ok != true {
		c = "clear"
	}

	return t, c
}

func GetWeather(location string, apikey string) (*Weather, error) {
	url := buildWeatherUrl(location, apikey)

	res, err := http.Get(url)
	if err != nil {
		return nil, errors.New("Weather API did not respond: " + err.Error())
	}

	j, err := ioutil.ReadAll(res.Body)
	res.Body.Close()

	if err != nil {
		return nil, errors.New("Weather API response empty")
	}

	return parseWeatherResponse(j)
}

func parseWeatherResponse(res []byte) (*Weather, error) {
	w := new(weatherContainer)
	err := json.Unmarshal(res, &w)

	if err != nil {
		return nil, errors.New("Error parsing response: " + err.Error())
	}

	return &w.Data, nil
}

func buildWeatherUrl(location string, apikey string) string {
	url := "http://free.worldweatheronline.com/feed/weather.ashx?"
	url += "format=json&num_of_days=5&key=" + apikey + "&q=" + location

	return url
}
