# Autom Go SDK

The official Go SDK for the [Autom](https://www.autom.dev) SERP API.

Google, Bing and Brave search results as structured JSON — web, maps, news, shopping, images, jobs and AI Mode.

The SDK exposes **23 typed methods**, one per documented endpoint, plus a generic `Request()`. It has no dependencies outside the standard library.

## Requirements

- Go 1.22 or newer

## Installation

```sh
go get github.com/automdev/autom-go
```

## Get your API key

[Sign up for Autom](https://app.autom.dev/register) or sign in, then copy your API key from the [dashboard](https://app.autom.dev). `NewClient("", nil)` reads the `AUTOM_API_KEY` environment variable.

## Quick start

```go
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/automdev/autom-go"
)

func main() {
	client, err := autom.NewClient("", nil) // reads AUTOM_API_KEY
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	response, err := client.GoogleSearch(ctx, &autom.GoogleSearchParams{Query: "web scraping api"})
	if err != nil {
		log.Fatal(err)
	}
	data, err := response.Map()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(data)
}
```

Parameter structs mirror the [API reference](https://docs.autom.dev). Required fields are plain values; optional fields are pointers — use `autom.String`, `autom.Float`, `autom.Int` and `autom.Bool`.

Decode into your own types with `response.Decode(&value)`.

## Async jobs

Any call accepts `is_async: true` (useful past the 60-second server limit): the response is then a `job_id`. `WaitForJob` polls `GET /v1/jobs/{id}` until a terminal status (`CompletedJob`, `Error`, `NotFound`, `BadRequest`):

```go
submitted, err := client.GoogleSearch(ctx, &autom.GoogleSearchParams{Query: "web scraping api", IsAsync: autom.Bool(true)})
if err != nil {
	log.Fatal(err)
}
var accepted struct{ JobID string `json:"job_id"` }
if err := submitted.Decode(&accepted); err != nil {
	log.Fatal(err)
}
job, err := client.WaitForJob(ctx, accepted.JobID, &autom.WaitForJobOptions{
	OnProgress: func(job map[string]any) { fmt.Println(job["status"], job["progress"]) },
})
```

## GET or POST

Methods use GET when the endpoint supports it (parameters in the query string); `WithMethod("POST")` sends them as a JSON body.

```go
response, err := client.BingSearch(ctx, &autom.BingSearchParams{Query: "test"}, autom.WithMethod("POST"))
```

## Any endpoint

```go
response, err := client.Request(ctx, "GET", "/v1/usage", nil)
```

## Errors

```go
response, err := client.GoogleSearch(ctx, &autom.GoogleSearchParams{Query: "web scraping api"})
if apiError, ok := autom.IsAPIError(err); ok {
	log.Printf("HTTP %d: %s", apiError.StatusCode, apiError.Message)
}
```

| Error | When |
| --- | --- |
| `*ConfigurationError` | Missing API key, invalid options or unsupported method. |
| `*SerializationError` | A required parameter is missing or a value can't be sent. |
| `*TransportError` | Network failure, timeout or cancelled context (`errors.Is(err, context.DeadlineExceeded)`). |
| `*APIError` | Non-2xx status (`StatusCode`, `Message`, `Response`). |
| `*ResponseDecodingError` | A success response could not be decoded. |

## Retries

Requests are retried up to 3 times after a network error, a `408`, `500`, `502`, `503` or `504`, or a rate limit — which the API returns as `401` (or `429`) with a `Rate limit exceeded` message. Per-key quota errors (`…: quota daily`) are not retried, nor are `400`, `402`, `404` or client-side timeouts. The delay starts at 1 second and grows 1.5× per retry (capped at 30 seconds); a `Retry-After` header wins.

```go
client, err := autom.NewClient("", &autom.ClientOptions{
	MaxRetries: 5,                      // negative disables retries
	RetryDelay: 500 * time.Millisecond,
})

// or for one request
response, err := client.GoogleSearch(ctx, &autom.GoogleSearchParams{Query: "web scraping api"}, autom.WithMaxRetries(0))
```

## Batch requests

`autom.Batch` runs a function per item with at most 10 calls in flight and returns one `BatchResult` per item, in input order. A failing item never stops the others.

```go
inputs := []string{"web scraping api", "serp api", "rank tracker"}
results := autom.Batch(ctx, inputs, 5, func(ctx context.Context, value string) (*autom.Response, error) {
	return client.GoogleSearch(ctx, &autom.GoogleSearchParams{Query: value})
})
for _, result := range results {
	if result.Err != nil {
		log.Printf("%s: %v", inputs[result.Index], result.Err)
		continue
	}
	fmt.Println(string(result.Value.Body))
}
```

| Constant | Default | Description |
| --- | --- | --- |
| `DefaultConcurrency` | `10` | Tasks run at the same time by batch |
| `DefaultTimeout` | `120 * time.Second` | Request timeout |
| `MaxRetries` | `3` | Retry attempts after a network error, a 408/5xx or a rate limit |
| `RetryBackoffFactor` | `1.5` | Exponential backoff multiplier |
| `DefaultRetryDelay` | `time.Second` | Delay before the first retry |

## Timeouts and custom HTTP clients

```go
client, err := autom.NewClient("", &autom.ClientOptions{
	Timeout:    60 * time.Second, // default 120 s; negative disables
	HTTPClient: &http.Client{Transport: myTransport},
})
```

Every method also honours the deadline and cancellation of the `context.Context` you pass.

## Endpoints

23 methods, generated from the [OpenAPI document](https://app.autom.dev/openapi).

| Method | HTTP | Description |
| --- | --- | --- |
| `BingSearch` | GET / POST `/v1/bing/search` | Query Bing Search via API and retrieve structured organic results (title, link, snippet, position). |
| `BraveImages` | GET / POST `/v1/brave/images` | Search Brave Images via API and retrieve structured image results with pagination and SafeSearch. |
| `BraveSearch` | GET / POST `/v1/brave/search` | Search Brave Search via API and retrieve structured organic results with pagination. |
| `BraveVideos` | GET / POST `/v1/brave/videos` | Search Brave Videos via API and retrieve structured video results with pagination and SafeSearch. |
| `DuckduckgoSearch` | GET / POST `/v1/duckduckgo/search` | Query DuckDuckGo via API and retrieve structured organic results with region and language. |
| `FinderGoogleCountries` | GET `/v1/finder/google-countries` | Search Google countries and retrieve country codes (free Finder endpoint). |
| `FinderGoogleLanguages` | GET `/v1/finder/google-languages` | Search Google languages and retrieve language codes (free Finder endpoint). |
| `FinderGoogleLocations` | GET `/v1/finder/google-locations` | Find supported Google locations and retrieve location IDs and country codes (free endpoint). |
| `GoogleImages` | GET / POST `/v1/google/images` | Search Google Images via API and retrieve image results (URLs, titles, sources, dimensions). |
| `GoogleJobs` | GET / POST `/v1/google/jobs` | Scrape Google Jobs results via API and retrieve structured job listings with title, company, location, and pagination. |
| `GoogleMaps` | GET / POST `/v1/google/maps` | Search Google Maps via API and retrieve structured place results (address, website, phone). |
| `GoogleNews` | GET / POST `/v1/google/news` | Search Google News via API and retrieve structured news results with pagination. |
| `GoogleSearch` | GET / POST `/v1/google/search` | Scrape Google Search results via API and retrieve structured SERP data with pagination. |
| `GoogleSearchAI` | GET / POST `/v1/google/search/ai` | Query Google AI Mode via API. JSON with text_blocks, references, and shopping_results. |
| `GoogleSearchAutocomplete` | GET / POST `/v1/google/search/autocomplete` | Get Google Search autocomplete suggestions via API for a given query. |
| `GoogleSearchLight` | GET / POST `/v1/google/search/light` | Scrape lightweight Google Search results via API: fast, cost-efficient, organic results only. |
| `GoogleShopping` | GET / POST `/v1/google/shopping` | Scrape Google Shopping results via API and retrieve structured product listings with prices, ratings, and merchant data. |
| `GoogleVideos` | GET / POST `/v1/google/videos` | Search Google Videos via API and retrieve structured video results with metadata. |
| `Jobs` | GET `/v1/jobs/{id}` | Poll the status, progress (0–100) and result of a job started with is_async=true. Polling is free. Poll until status is terminal: Completed… |
| `Search` | GET / POST `/v1/search` | One call walks engines in order and returns the first organic results. engines=auto tries Google, then Brave, Bing, DuckDuckGo, Yahoo, and… |
| `Stats` | GET `/v1/stats` | Infrastructure health for an engine or endpoint over the last hour. Not scoped to your account or tenant. Enabled per account: HTTP 403 mea… |
| `Usage` | GET `/v1/usage` | Retrieve real-time usage statistics and credit balance for your Autom account. |
| `YahooSearch` | GET / POST `/v1/yahoo/search` | Query Yahoo via API and retrieve structured organic results with country and language. |

## License

MIT — see [LICENSE](LICENSE).
