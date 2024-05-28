# llm-function-get-weather

This example demostrates how to implement OpenAI Function Calling in Strong-typed language, easy to write and maintain. follow the instructions below, you will extend OpenAI gpt-4o with openweathermap.org API.

## Development on local

### Install YoMo

```sh
curl -fsSL "https://get.yomo.run" | sh
```

### Add your OpenAI SK to configuration

```sh
vim yomo.yml
```

> [!TIP]
> If you do not have an OpenAI API Key, you can try [vivgrid.com](https://vivgrid.com), it's free for developers with generous OpenAI Token every month.

### Start YoMo

```sh
yomo serve -c yomo.yml
```

### Get your api-key from OpenWeatherMap.org

grab your api-key from [openweathermap.org](https://openweathermap.org) free, and add it to your `.env` file:

```sh
YOMO_SFN_NAME=llm_fn_get_weather
YOMO_SFN_ZIPPER=localhost:9000
YOMO_SFN_CREDENTIAL=SECRET
OPENWEATHERMAP_API_KEY=<ADD_HERE>
```

### Run your function calling

```sh
yomo run app.go
```

### Try it!

```sh
curl -X POST -H "Content-Type: application/json" -d '{"prompt":"Hows the weather like in Paris today?"}' http://127.0.0.1:9000/invoke |jq    
{
  "Content": "The weather in Paris today is clear with a clear sky. Here are the detailed conditions:\n\n- **Temperature:** 9.89°C (feels like 9.02°C)\n- **Humidity:** 91%\n- **Pressure:** 1020 hPa\n- **Wind:** 2.06 m/s coming from the southwest (210°)\n- **Visibility:** 10 km\n- **Cloudiness:** 0% (clear sky)\n- **Sunrise:** 05:27 AM\n- **Sunset:** 08:41 PM\n\nOverall, it's a cool and clear day in Paris.",
  "ToolCalls": null,
  "ToolMessages": null,
  "FinishReason": "stop",
  "TokenUsage": {
    "prompt_tokens": 274,
    "completion_tokens": 125
  },
  "AssistantMessage": null
}
```

