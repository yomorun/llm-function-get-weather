package main

import (
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"

	"github.com/yomorun/yomo/serverless"
)

// Description defines the Function Calling, this will be combine to prompt automatically. As openweathermap.org does not support request api by `city_name`, we have ask LLM to translate it to Latitude and Longitude.
func Description() string {
	return `Get current weather for a giving city. If no city is provided, you should ask to clarify the city. 
	If the city name is given, you should convert city name to Latitude and Longitude geo coordinates, 
	keep Latitude and Longitude in decimal format.`
}

// Param defines the input parameters of the Function Calling, this will be combine to prompt automatically
type Param struct {
	City      string  `json:"city" jsonschema:"description=The city name to get the weather for"`
	Latitude  float64 `json:"latitude" jsonschema:"description=The latitude of the city, in decimal format, range should be in (-90, 90)"`
	Longitude float64 `json:"longitude" jsonschema:"description=The longitude of the city, in decimal format, range should be in (-180, 180)"`
}

// InputSchema decalares the type of Function Calling Parameter, required by YoMo
func InputSchema() any {
	return &Param{}
}

// DataTags declares the type of data this stateful serverless observe, required by YoMo
func DataTags() []uint32 {
	return []uint32{0x30}
}

// Handler is the main entry point of the Function Calling when LLM response with `tool_call`
func Handler(ctx serverless.Context) {
	var p Param
	// deserilize the parameters from llm tool_call response
	ctx.ReadLLMArguments(&p)

	// call the openweathermap api and write the result back to LLM
	result := requestOpenWeatherMap(p.Latitude, p.Longitude)
	ctx.WriteLLMResult(result)

	slog.Info("get-weather", "city", p.City, "rag", result)
}

func requestOpenWeatherMap(lat, lon float64) string {
	const apiURL = "https://api.openweathermap.org/data/2.5/weather?lat=%f&lon=%f&appid=%s&units=metric"
	apiKey := os.Getenv("OPENWEATHERMAP_API_KEY")
	url := fmt.Sprintf(apiURL, lat, lon, apiKey)
	// send api request to openweathermap.org
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return "can not get the weather information at the moment"
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return "can not get the weather information at the moment"
	}

	return string(body)
}
