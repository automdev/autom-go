// Code generated from the Autom OpenAPI document. DO NOT EDIT.

package autom

import "context"

// BingSearch — Bing Search. Query Bing Search via API and retrieve structured organic results (title, link, snippet, position). Credit cost: 1 credit per call.
//
// GET / POST /v1/bing/search
func (c *Client) BingSearch(ctx context.Context, params *BingSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["bing_search"], params.params(), options)
}

// BraveImages — Brave Images. Search Brave Images via API and retrieve structured image results with pagination and SafeSearch. Credit cost: 1 credit per call.
//
// GET / POST /v1/brave/images
func (c *Client) BraveImages(ctx context.Context, params *BraveImagesParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["brave_images"], params.params(), options)
}

// BraveSearch — Brave Search. Search Brave Search via API and retrieve structured organic results with pagination. Credit cost: 1 credit per call.
//
// GET / POST /v1/brave/search
func (c *Client) BraveSearch(ctx context.Context, params *BraveSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["brave_search"], params.params(), options)
}

// BraveVideos — Brave Videos. Search Brave Videos via API and retrieve structured video results with pagination and SafeSearch. Credit cost: 1 credit per call.
//
// GET / POST /v1/brave/videos
func (c *Client) BraveVideos(ctx context.Context, params *BraveVideosParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["brave_videos"], params.params(), options)
}

// DuckduckgoSearch — DuckDuckGo Search. Query DuckDuckGo via API and retrieve structured organic results with region and language. Credit cost: 1 credit per call.
//
// GET / POST /v1/duckduckgo/search
func (c *Client) DuckduckgoSearch(ctx context.Context, params *DuckduckgoSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["duckduckgo_search"], params.params(), options)
}

// FinderGoogleCountries — Google Countries. Search Google countries and retrieve country codes (free Finder endpoint).
//
// GET /v1/finder/google-countries
func (c *Client) FinderGoogleCountries(ctx context.Context, params *FinderGoogleCountriesParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["finder_google_countries"], params.params(), options)
}

// FinderGoogleLanguages — Google Languages. Search Google languages and retrieve language codes (free Finder endpoint).
//
// GET /v1/finder/google-languages
func (c *Client) FinderGoogleLanguages(ctx context.Context, params *FinderGoogleLanguagesParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["finder_google_languages"], params.params(), options)
}

// FinderGoogleLocations — Google Locations. Find supported Google locations and retrieve location IDs and country codes (free endpoint).
//
// GET /v1/finder/google-locations
func (c *Client) FinderGoogleLocations(ctx context.Context, params *FinderGoogleLocationsParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["finder_google_locations"], params.params(), options)
}

// GoogleImages — Google Images. Search Google Images via API and retrieve image results (URLs, titles, sources, dimensions). Credit cost: 1 credit per call.
//
// GET / POST /v1/google/images
func (c *Client) GoogleImages(ctx context.Context, params *GoogleImagesParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["google_images"], params.params(), options)
}

// GoogleJobs — Google Jobs. Scrape Google Jobs results via API and retrieve structured job listings with title, company, location, and pagination. Credit cost: 1 credit per call.
//
// GET / POST /v1/google/jobs
func (c *Client) GoogleJobs(ctx context.Context, params *GoogleJobsParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["google_jobs"], params.params(), options)
}

// GoogleMaps — Google Maps. Search Google Maps via API and retrieve structured place results (address, website, phone). Credit cost: 1 credit per call.
//
// GET / POST /v1/google/maps
func (c *Client) GoogleMaps(ctx context.Context, params *GoogleMapsParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["google_maps"], params.params(), options)
}

// GoogleNews — Google News. Search Google News via API and retrieve structured news results with pagination. Credit cost: 1 credit per call.
//
// GET / POST /v1/google/news
func (c *Client) GoogleNews(ctx context.Context, params *GoogleNewsParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["google_news"], params.params(), options)
}

// GoogleSearch — Google Search. Scrape Google Search results via API and retrieve structured SERP data with pagination. Credit cost: 1 credit per call.
//
// GET / POST /v1/google/search
func (c *Client) GoogleSearch(ctx context.Context, params *GoogleSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["google_search"], params.params(), options)
}

// GoogleSearchAI — Google AI Mode. Query Google AI Mode via API. JSON with text_blocks, references, and shopping_results. Credit cost: 3 credits per call.
//
// GET / POST /v1/google/search/ai
func (c *Client) GoogleSearchAI(ctx context.Context, params *GoogleSearchAIParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["google_search_ai"], params.params(), options)
}

// GoogleSearchAutocomplete — Google Search Autocomplete. Get Google Search autocomplete suggestions via API for a given query. Credit cost: 1 credit per call.
//
// GET / POST /v1/google/search/autocomplete
func (c *Client) GoogleSearchAutocomplete(ctx context.Context, params *GoogleSearchAutocompleteParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["google_search_autocomplete"], params.params(), options)
}

// GoogleSearchLight — Google Search Light. Scrape lightweight Google Search results via API: fast, cost-efficient, organic results only. Credit cost: 1 credit per call.
//
// GET / POST /v1/google/search/light
func (c *Client) GoogleSearchLight(ctx context.Context, params *GoogleSearchLightParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["google_search_light"], params.params(), options)
}

// GoogleShopping — Google Shopping. Scrape Google Shopping results via API and retrieve structured product listings with prices, ratings, and merchant data. Credit cost: 1 credit per call.
//
// GET / POST /v1/google/shopping
func (c *Client) GoogleShopping(ctx context.Context, params *GoogleShoppingParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["google_shopping"], params.params(), options)
}

// GoogleVideos — Google Videos. Search Google Videos via API and retrieve structured video results with metadata. Credit cost: 1 credit per call.
//
// GET / POST /v1/google/videos
func (c *Client) GoogleVideos(ctx context.Context, params *GoogleVideosParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["google_videos"], params.params(), options)
}

// Jobs — Jobs. Poll the status, progress (0–100) and result of a job started with is_async=true. Polling is free. Poll until status is terminal: CompletedJob, Error, NotFound or BadRequest.
//
// GET /v1/jobs/{id}
func (c *Client) Jobs(ctx context.Context, params *JobsParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["jobs"], params.params(), options)
}

// Search — Search. One call walks engines in order and returns the first organic results. engines=auto tries Google, then Brave, Bing, DuckDuckGo, Yahoo, and Baidu. Omitted gl is us and omitted hl is en. There is no page parameter: the response is always pag… Credit cost: 1 credit per call.
//
// GET / POST /v1/search
func (c *Client) Search(ctx context.Context, params *SearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["search"], params.params(), options)
}

// Stats — Stats. Infrastructure health for an engine or endpoint over the last hour. Not scoped to your account or tenant. Enabled per account: HTTP 403 means contact support to activate it.
//
// GET /v1/stats
func (c *Client) Stats(ctx context.Context, params *StatsParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["stats"], params.params(), options)
}

// Usage — Usage. Retrieve real-time usage statistics and credit balance for your Autom account.
//
// GET /v1/usage
func (c *Client) Usage(ctx context.Context, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["usage"], nil, options)
}

// YahooSearch — Yahoo Search. Query Yahoo via API and retrieve structured organic results with country and language. Credit cost: 1 credit per call.
//
// GET / POST /v1/yahoo/search
func (c *Client) YahooSearch(ctx context.Context, params *YahooSearchParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["yahoo_search"], params.params(), options)
}

// endpoints lists every documented operation by its snake_case name.
var endpoints = map[string]endpoint{
	"bing_search":                {path: "/v1/bing/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"brave_images":               {path: "/v1/brave/images", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"brave_search":               {path: "/v1/brave/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"brave_videos":               {path: "/v1/brave/videos", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"duckduckgo_search":          {path: "/v1/duckduckgo/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"finder_google_countries":    {path: "/v1/finder/google-countries", methods: []string{"GET"}, json: true, required: []string{"query"}},
	"finder_google_languages":    {path: "/v1/finder/google-languages", methods: []string{"GET"}, json: true, required: []string{"query"}},
	"finder_google_locations":    {path: "/v1/finder/google-locations", methods: []string{"GET"}, json: true, required: []string{"query"}},
	"google_images":              {path: "/v1/google/images", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"google_jobs":                {path: "/v1/google/jobs", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"google_maps":                {path: "/v1/google/maps", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"google_news":                {path: "/v1/google/news", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"google_search":              {path: "/v1/google/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"google_search_ai":           {path: "/v1/google/search/ai", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"google_search_autocomplete": {path: "/v1/google/search/autocomplete", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"google_search_light":        {path: "/v1/google/search/light", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"google_shopping":            {path: "/v1/google/shopping", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"google_videos":              {path: "/v1/google/videos", methods: []string{"GET", "POST"}, json: true, required: []string{"query", "page"}},
	"jobs":                       {path: "/v1/jobs/{id}", methods: []string{"GET"}, json: true, required: []string{"id"}},
	"search":                     {path: "/v1/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"stats":                      {path: "/v1/stats", methods: []string{"GET"}, json: true, required: []string{}},
	"usage":                      {path: "/v1/usage", methods: []string{"GET"}, json: true, required: []string{}},
	"yahoo_search":               {path: "/v1/yahoo/search", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
}
