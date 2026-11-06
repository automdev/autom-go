// Code generated from the Autom OpenAPI document. DO NOT EDIT.

package autom

// BingSearchParams holds the parameters of BingSearch (GET / POST /v1/bing/search).
// Required fields are plain values; optional fields are pointers (use autom.String, autom.Float, autom.Bool) or nil slices and maps.
type BingSearchParams struct {
	// The search query you want to run on Bing.
	Query string `json:"query" api:"query,required"`
	// Page number for pagination (1 = first page). Default: 1.
	Page *float64 `json:"page" api:"page"`
	// Number of results per page (e.g. 10, 20, 50). This value is a suggestion. Default: 10.
	Num *float64 `json:"num" api:"num"`
	// Search origin location (e.g. city).
	Location *string `json:"location" api:"location"`
	// GPS latitude for the search origin.
	Lat *float64 `json:"lat" api:"lat"`
	// GPS longitude for the search origin.
	Lon *float64 `json:"lon" api:"lon"`
	// Market code where results come from (e.g. en-US). Cannot be used together with cc.
	Mkt *string `json:"mkt" api:"mkt"`
	// Country code to search from (2-letter ISO 3166-1). Cannot be used together with mkt.
	Cc *string `json:"cc" api:"cc"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *BingSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 11)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Page != nil {
		values = append(values, param{"page", *p.Page})
	}
	if p.Num != nil {
		values = append(values, param{"num", *p.Num})
	}
	if p.Location != nil {
		values = append(values, param{"location", *p.Location})
	}
	if p.Lat != nil {
		values = append(values, param{"lat", *p.Lat})
	}
	if p.Lon != nil {
		values = append(values, param{"lon", *p.Lon})
	}
	if p.Mkt != nil {
		values = append(values, param{"mkt", *p.Mkt})
	}
	if p.Cc != nil {
		values = append(values, param{"cc", *p.Cc})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// BraveImagesParams holds the parameters of BraveImages (GET / POST /v1/brave/images).
// Required fields are plain values; optional fields are pointers (use autom.String, autom.Float, autom.Bool) or nil slices and maps.
type BraveImagesParams struct {
	// The search query you want to run on Brave Images.
	Query string `json:"query" api:"query,required"`
	// Page number for pagination (50 results per page). Default: 1.
	Page *float64 `json:"page" api:"page"`
	// Country code (e.g. fr, us). Mapped to Brave country.
	Gl *string `json:"gl" api:"gl"`
	// Language code (e.g. fr, en). Mapped to Brave lang.
	Hl *string `json:"hl" api:"hl"`
	// SafeSearch: active | off | moderate | strict. Google-style active maps to Brave strict.
	Safe *string `json:"safe" api:"safe"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *BraveImagesParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 8)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Page != nil {
		values = append(values, param{"page", *p.Page})
	}
	if p.Gl != nil {
		values = append(values, param{"gl", *p.Gl})
	}
	if p.Hl != nil {
		values = append(values, param{"hl", *p.Hl})
	}
	if p.Safe != nil {
		values = append(values, param{"safe", *p.Safe})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// BraveSearchParams holds the parameters of BraveSearch (GET / POST /v1/brave/search).
// Required fields are plain values; optional fields are pointers (use autom.String, autom.Float, autom.Bool) or nil slices and maps.
type BraveSearchParams struct {
	// The search query you want to run on Brave Search.
	Query string `json:"query" api:"query,required"`
	// Page number for pagination. Default: 1.
	Page *float64 `json:"page" api:"page"`
	// Country code, for example us, fr, gb, de. Default: "us".
	Gl *string `json:"gl" api:"gl"`
	// Language code, for example en, fr, de. Default: "en".
	Hl *string `json:"hl" api:"hl"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *BraveSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 7)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Page != nil {
		values = append(values, param{"page", *p.Page})
	}
	if p.Gl != nil {
		values = append(values, param{"gl", *p.Gl})
	}
	if p.Hl != nil {
		values = append(values, param{"hl", *p.Hl})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// BraveVideosParams holds the parameters of BraveVideos (GET / POST /v1/brave/videos).
// Required fields are plain values; optional fields are pointers (use autom.String, autom.Float, autom.Bool) or nil slices and maps.
type BraveVideosParams struct {
	// The search query you want to run on Brave Videos.
	Query string `json:"query" api:"query,required"`
	// Page number for pagination. Default: 1.
	Page *float64 `json:"page" api:"page"`
	// Country code (e.g. fr, us). Mapped to Brave country.
	Gl *string `json:"gl" api:"gl"`
	// Language code (e.g. fr, en). Mapped to Brave lang.
	Hl *string `json:"hl" api:"hl"`
	// SafeSearch: active | off | moderate | strict. Google-style active maps to Brave strict.
	Safe *string `json:"safe" api:"safe"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *BraveVideosParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 8)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Page != nil {
		values = append(values, param{"page", *p.Page})
	}
	if p.Gl != nil {
		values = append(values, param{"gl", *p.Gl})
	}
	if p.Hl != nil {
		values = append(values, param{"hl", *p.Hl})
	}
	if p.Safe != nil {
		values = append(values, param{"safe", *p.Safe})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// DuckduckgoSearchParams holds the parameters of DuckduckgoSearch (GET / POST /v1/duckduckgo/search).
// Required fields are plain values; optional fields are pointers (use autom.String, autom.Float, autom.Bool) or nil slices and maps.
type DuckduckgoSearchParams struct {
	// Search query sent to DuckDuckGo Search.
	Query string `json:"query" api:"query,required"`
	// Page number. Page 1 is the first page of results. Default: 1.
	Page *float64 `json:"page" api:"page"`
	// Country code, for example us, fr, gb, de, jp.
	Gl *string `json:"gl" api:"gl"`
	// Language code, for example en, fr, de, ja, zh.
	Hl *string `json:"hl" api:"hl"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *DuckduckgoSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 7)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Page != nil {
		values = append(values, param{"page", *p.Page})
	}
	if p.Gl != nil {
		values = append(values, param{"gl", *p.Gl})
	}
	if p.Hl != nil {
		values = append(values, param{"hl", *p.Hl})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// FinderGoogleCountriesParams holds the parameters of FinderGoogleCountries (GET /v1/finder/google-countries).
// Required fields are plain values; optional fields are pointers (use autom.String, autom.Float, autom.Bool) or nil slices and maps.
type FinderGoogleCountriesParams struct {
	// Filter countries by a search string (case-insensitive). Returns all countries whose code or name matches the specified query.
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *FinderGoogleCountriesParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// FinderGoogleLanguagesParams holds the parameters of FinderGoogleLanguages (GET /v1/finder/google-languages).
// Required fields are plain values; optional fields are pointers (use autom.String, autom.Float, autom.Bool) or nil slices and maps.
type FinderGoogleLanguagesParams struct {
	// Filter languages by a search string (case-insensitive). Returns all languages whose names contain the specified string.
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *FinderGoogleLanguagesParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// FinderGoogleLocationsParams holds the parameters of FinderGoogleLocations (GET /v1/finder/google-locations).
// Required fields are plain values; optional fields are pointers (use autom.String, autom.Float, autom.Bool) or nil slices and maps.
type FinderGoogleLocationsParams struct {
	// Search locations supported by Piloterr (cities, countries, etc.) by providing a query string. Returns matching locations ordered by reach.
	Query string `json:"query" api:"query,required"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *FinderGoogleLocationsParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 4)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// GoogleImagesParams holds the parameters of GoogleImages (GET / POST /v1/google/images).
// Required fields are plain values; optional fields are pointers (use autom.String, autom.Float, autom.Bool) or nil slices and maps.
type GoogleImagesParams struct {
	// Query you want to search on Google Images (e.g. "lion", "nike logo", "modern house exterior").
	Query string `json:"query" api:"query,required"`
	// Page number for pagination (1 = first page). Default: 1.
	Page *float64 `json:"page" api:"page"`
	// Google domain to use for the search. Defaults to google.com. Default: "google.com".
	GoogleDomain *string `json:"google_domain" api:"google_domain"`
	// Country to use for the Google search. Two-letter country code (e.g. us for the United States, uk for United Kingdom, or fr for France).
	Gl *string `json:"gl" api:"gl"`
	// Language to use for the Google search. Two or three-letter language code (e.g. en, es, fr), optionally followed by a region code (e.g. en-gb, es-419).
	Hl *string `json:"hl" api:"hl"`
	// Limit results to one or more countries. Use country{XX} with | as delimiter (e.g. countryFR|countryDE).
	Cr *string `json:"cr" api:"cr"`
	// Limit results to one or more languages. Use lang_{xx} with | as delimiter (e.g. lang_fr|lang_de).
	Lr *string `json:"lr" api:"lr"`
	// If true, returns a job and lets you fetch results asynchronously. Default: false.
	Async *bool `json:"async" api:"async"`
	// Adult content filter. active enables Google SafeSearch; off disables it. Omit to use Google default.
	Safe *string `json:"safe" api:"safe"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *GoogleImagesParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 12)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Page != nil {
		values = append(values, param{"page", *p.Page})
	}
	if p.GoogleDomain != nil {
		values = append(values, param{"google_domain", *p.GoogleDomain})
	}
	if p.Gl != nil {
		values = append(values, param{"gl", *p.Gl})
	}
	if p.Hl != nil {
		values = append(values, param{"hl", *p.Hl})
	}
	if p.Cr != nil {
		values = append(values, param{"cr", *p.Cr})
	}
	if p.Lr != nil {
		values = append(values, param{"lr", *p.Lr})
	}
	if p.Async != nil {
		values = append(values, param{"async", *p.Async})
	}
	if p.Safe != nil {
		values = append(values, param{"safe", *p.Safe})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// GoogleJobsParams holds the parameters of GoogleJobs (GET / POST /v1/google/jobs).
// Required fields are plain values; optional fields are pointers (use autom.String, autom.Float, autom.Bool) or nil slices and maps.
type GoogleJobsParams struct {
	// The job search query to run on Google Jobs (e.g. developer, barista new york).
	Query string `json:"query" api:"query,required"`
	// Advanced search filters string for Google Jobs (encoded filter chips).
	Uds *string `json:"uds" api:"uds"`
	// Search radius in kilometers. Does not strictly limit the radius, but influences result proximity.
	Lrad *float64 `json:"lrad" api:"lrad"`
	// Filter results by work type. Set to 1 to filter for work from home / remote jobs.
	Ltype *float64 `json:"ltype" api:"ltype"`
	// Human-readable location to originate the search from (e.g. New York, NY, Paris, France). If several locations match, the most popular one is used. Note: location and uule cannot be used together.
	Location *string `json:"location" api:"location"`
	// Google encoded location parameter. Use for precise geo targeting. Note: uule and location cannot be used together.
	Uule *string `json:"uule" api:"uule"`
	// Google domain to use for the search. Defaults to google.com. Default: "google.com".
	GoogleDomain *string `json:"google_domain" api:"google_domain"`
	// Country to use for the Google search. Two-letter country code (e.g. us for the United States, uk for United Kingdom, or fr for France). Default: "us".
	Gl *string `json:"gl" api:"gl"`
	// Language to use for the Google search. Two or three-letter language code (e.g. en, es, fr), optionally followed by a region code (e.g. en-gb, es-419). Default: "en".
	Hl *string `json:"hl" api:"hl"`
	// Limit results to one or more countries. Use country{XX} with | as delimiter (e.g. countryFR|countryDE).
	Cr *string `json:"cr" api:"cr"`
	// Limit results to one or more languages. Use lang_{xx} with | as delimiter (e.g. lang_fr|lang_de).
	Lr *string `json:"lr" api:"lr"`
	// Token to fetch the next page of results. Returned in the serpapi_pagination.next_page_token field of the previous response.
	NextPageToken *string `json:"next_page_token" api:"next_page_token"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *GoogleJobsParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 15)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Uds != nil {
		values = append(values, param{"uds", *p.Uds})
	}
	if p.Lrad != nil {
		values = append(values, param{"lrad", *p.Lrad})
	}
	if p.Ltype != nil {
		values = append(values, param{"ltype", *p.Ltype})
	}
	if p.Location != nil {
		values = append(values, param{"location", *p.Location})
	}
	if p.Uule != nil {
		values = append(values, param{"uule", *p.Uule})
	}
	if p.GoogleDomain != nil {
		values = append(values, param{"google_domain", *p.GoogleDomain})
	}
	if p.Gl != nil {
		values = append(values, param{"gl", *p.Gl})
	}
	if p.Hl != nil {
		values = append(values, param{"hl", *p.Hl})
	}
	if p.Cr != nil {
		values = append(values, param{"cr", *p.Cr})
	}
	if p.Lr != nil {
		values = append(values, param{"lr", *p.Lr})
	}
	if p.NextPageToken != nil {
		values = append(values, param{"next_page_token", *p.NextPageToken})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// GoogleMapsParams holds the parameters of GoogleMaps (GET / POST /v1/google/maps).
// Required fields are plain values; optional fields are pointers (use autom.String, autom.Float, autom.Bool) or nil slices and maps.
type GoogleMapsParams struct {
	// Search query for Google Maps (e.g. business type, brand, or place name).
	Query string `json:"query" api:"query,required"`
	// Page number for pagination (1 = first page). Default: 1.
	Page *float64 `json:"page" api:"page"`
	// Google domain to use for the search. Defaults to google.com. Default: "google.com".
	GoogleDomain *string `json:"google_domain" api:"google_domain"`
	// Country to use for the Google search. Two-letter country code (e.g. us for the United States, uk for United Kingdom, or fr for France).
	Gl *string `json:"gl" api:"gl"`
	// Language to use for the Google search. Two or three-letter language code (e.g. en, es, fr), optionally followed by a region code (e.g. en-gb, es-419).
	Hl *string `json:"hl" api:"hl"`
	// Limit results to one or more countries. Use country{XX} with | as delimiter (e.g. countryFR|countryDE).
	Cr *string `json:"cr" api:"cr"`
	// Limit results to one or more languages. Use lang_{xx} with | as delimiter (e.g. lang_fr|lang_de).
	Lr *string `json:"lr" api:"lr"`
	// Map center / geo hint in Google format (e.g. @43.5131985,1.4600501,60m). Helps focus results around a coordinate and zoom radius.
	Ll *string `json:"ll" api:"ll"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *GoogleMapsParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 11)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Page != nil {
		values = append(values, param{"page", *p.Page})
	}
	if p.GoogleDomain != nil {
		values = append(values, param{"google_domain", *p.GoogleDomain})
	}
	if p.Gl != nil {
		values = append(values, param{"gl", *p.Gl})
	}
	if p.Hl != nil {
		values = append(values, param{"hl", *p.Hl})
	}
	if p.Cr != nil {
		values = append(values, param{"cr", *p.Cr})
	}
	if p.Lr != nil {
		values = append(values, param{"lr", *p.Lr})
	}
	if p.Ll != nil {
		values = append(values, param{"ll", *p.Ll})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// GoogleNewsParams holds the parameters of GoogleNews (GET / POST /v1/google/news).
// Required fields are plain values; optional fields are pointers (use autom.String, autom.Float, autom.Bool) or nil slices and maps.
type GoogleNewsParams struct {
	// Query to search on Google News (keywords, company, topic).
	Query string `json:"query" api:"query,required"`
	// Page number for pagination (1 = first page). Default: 1.
	Page *float64 `json:"page" api:"page"`
	// Maximum number of results per page (e.g. 10). Default: 10.
	Num *float64 `json:"num" api:"num"`
	// Google domain to use for the search. Defaults to google.com. Default: "google.com".
	GoogleDomain *string `json:"google_domain" api:"google_domain"`
	// Country to use for the Google search. Two-letter country code (e.g. us for the United States, uk for United Kingdom, or fr for France).
	Gl *string `json:"gl" api:"gl"`
	// Language to use for the Google search. Two or three-letter language code (e.g. en, es, fr), optionally followed by a region code (e.g. en-gb, es-419).
	Hl *string `json:"hl" api:"hl"`
	// Limit results to one or more countries. Use country{XX} with | as delimiter (e.g. countryFR|countryDE).
	Cr *string `json:"cr" api:"cr"`
	// Limit results to one or more languages. Use lang_{xx} with | as delimiter (e.g. lang_fr|lang_de).
	Lr *string `json:"lr" api:"lr"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *GoogleNewsParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 11)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Page != nil {
		values = append(values, param{"page", *p.Page})
	}
	if p.Num != nil {
		values = append(values, param{"num", *p.Num})
	}
	if p.GoogleDomain != nil {
		values = append(values, param{"google_domain", *p.GoogleDomain})
	}
	if p.Gl != nil {
		values = append(values, param{"gl", *p.Gl})
	}
	if p.Hl != nil {
		values = append(values, param{"hl", *p.Hl})
	}
	if p.Cr != nil {
		values = append(values, param{"cr", *p.Cr})
	}
	if p.Lr != nil {
		values = append(values, param{"lr", *p.Lr})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// GoogleSearchParams holds the parameters of GoogleSearch (GET / POST /v1/google/search).
// Required fields are plain values; optional fields are pointers (use autom.String, autom.Float, autom.Bool) or nil slices and maps.
type GoogleSearchParams struct {
	// The search query you want to run on Google (supports advanced operators like site:, inurl:, intitle:).
	Query string `json:"query" api:"query,required"`
	// Advanced search filters string (e.g. time ranges: qdr:h last hour, qdr:d last day, qdr:w last week, qdr:m last month, qdr:y last year).
	Tbs *string `json:"tbs" api:"tbs"`
	// Adult content filter. active enables Google SafeSearch; off disables it. Omit to use Google default.
	Safe *string `json:"safe" api:"safe"`
	// Human-readable location to originate the search from (e.g. New York, NY, Paris, France). If several locations match, the most popular one is used. Note: location and uule cannot be used together.
	Location *string `json:"location" api:"location"`
	// Google encoded location parameter. Use for precise geo targeting. Note: uule and location cannot be used together.
	Uule *string `json:"uule" api:"uule"`
	// Google domain to use for the search. Defaults to google.com. Default: "google.com".
	GoogleDomain *string `json:"google_domain" api:"google_domain"`
	// Country to use for the Google search. Two-letter country code (e.g. us for the United States, uk for United Kingdom, or fr for France).
	Gl *string `json:"gl" api:"gl"`
	// Language to use for the Google search. Two or three-letter language code (e.g. en, es, fr), optionally followed by a region code (e.g. en-gb, es-419).
	Hl *string `json:"hl" api:"hl"`
	// Limit results to one or more countries. Use country{XX} with | as delimiter (e.g. countryFR|countryDE).
	Cr *string `json:"cr" api:"cr"`
	// Limit results to one or more languages. Use lang_{xx} with | as delimiter (e.g. lang_fr|lang_de).
	Lr *string `json:"lr" api:"lr"`
	// Page number / offset used for pagination. Default: 1.
	Page *float64 `json:"page" api:"page"`
	// Maximum number of results to return per page (deprecated). Default: 10.
	Num *float64 `json:"num" api:"num"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *GoogleSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 15)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Tbs != nil {
		values = append(values, param{"tbs", *p.Tbs})
	}
	if p.Safe != nil {
		values = append(values, param{"safe", *p.Safe})
	}
	if p.Location != nil {
		values = append(values, param{"location", *p.Location})
	}
	if p.Uule != nil {
		values = append(values, param{"uule", *p.Uule})
	}
	if p.GoogleDomain != nil {
		values = append(values, param{"google_domain", *p.GoogleDomain})
	}
	if p.Gl != nil {
		values = append(values, param{"gl", *p.Gl})
	}
	if p.Hl != nil {
		values = append(values, param{"hl", *p.Hl})
	}
	if p.Cr != nil {
		values = append(values, param{"cr", *p.Cr})
	}
	if p.Lr != nil {
		values = append(values, param{"lr", *p.Lr})
	}
	if p.Page != nil {
		values = append(values, param{"page", *p.Page})
	}
	if p.Num != nil {
		values = append(values, param{"num", *p.Num})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// GoogleSearchAIParams holds the parameters of GoogleSearchAI (GET / POST /v1/google/search/ai).
// Required fields are plain values; optional fields are pointers (use autom.String, autom.Float, autom.Bool) or nil slices and maps.
type GoogleSearchAIParams struct {
	// The question or search query you want Google AI Mode to answer (supports advanced operators like site:, inurl:, intitle:).
	Query string `json:"query" api:"query,required"`
	// Human-readable location to originate the search from (e.g. New York, NY, Paris, France). If several locations match, the most popular one is used. Note: location and uule cannot be used together.
	Location *string `json:"location" api:"location"`
	// Google encoded location parameter. Use for precise geo targeting. Note: uule and location cannot be used together.
	Uule *string `json:"uule" api:"uule"`
	// Google domain to use for the search. Defaults to google.com. Default: "google.com".
	GoogleDomain *string `json:"google_domain" api:"google_domain"`
	// Country to use for the Google search. Two-letter country code (e.g. us for the United States, uk for United Kingdom, or fr for France).
	Gl *string `json:"gl" api:"gl"`
	// Language to use for the Google search. Two or three-letter language code (e.g. en, es, fr), optionally followed by a region code (e.g. en-gb, es-419).
	Hl *string `json:"hl" api:"hl"`
	// Limit results to one or more countries. Use country{XX} with | as delimiter (e.g. countryFR|countryDE).
	Cr *string `json:"cr" api:"cr"`
	// Limit results to one or more languages. Use lang_{xx} with | as delimiter (e.g. lang_fr|lang_de).
	Lr *string `json:"lr" api:"lr"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *GoogleSearchAIParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 11)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Location != nil {
		values = append(values, param{"location", *p.Location})
	}
	if p.Uule != nil {
		values = append(values, param{"uule", *p.Uule})
	}
	if p.GoogleDomain != nil {
		values = append(values, param{"google_domain", *p.GoogleDomain})
	}
	if p.Gl != nil {
		values = append(values, param{"gl", *p.Gl})
	}
	if p.Hl != nil {
		values = append(values, param{"hl", *p.Hl})
	}
	if p.Cr != nil {
		values = append(values, param{"cr", *p.Cr})
	}
	if p.Lr != nil {
		values = append(values, param{"lr", *p.Lr})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// GoogleSearchAutocompleteParams holds the parameters of GoogleSearchAutocomplete (GET / POST /v1/google/search/autocomplete).
// Required fields are plain values; optional fields are pointers (use autom.String, autom.Float, autom.Bool) or nil slices and maps.
type GoogleSearchAutocompleteParams struct {
	// The search query for which you want to get autocomplete suggestions.
	Query string `json:"query" api:"query,required"`
	// Cursor position in the search query where you want to get autocomplete suggestions. Default: 0.
	Cp *float64 `json:"cp" api:"cp"`
	// Google domain to use for the search. Defaults to google.com. Default: "google.com".
	GoogleDomain *string `json:"google_domain" api:"google_domain"`
	// Country to use for the Google search. Two-letter country code (e.g. us for the United States, uk for United Kingdom, or fr for France).
	Gl *string `json:"gl" api:"gl"`
	// Language to use for the Google search. Two or three-letter language code (e.g. en, es, fr), optionally followed by a region code (e.g. en-gb, es-419).
	Hl *string `json:"hl" api:"hl"`
	// Limit results to one or more countries. Use country{XX} with | as delimiter (e.g. countryFR|countryDE).
	Cr *string `json:"cr" api:"cr"`
	// Limit results to one or more languages. Use lang_{xx} with | as delimiter (e.g. lang_fr|lang_de).
	Lr *string `json:"lr" api:"lr"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *GoogleSearchAutocompleteParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 10)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Cp != nil {
		values = append(values, param{"cp", *p.Cp})
	}
	if p.GoogleDomain != nil {
		values = append(values, param{"google_domain", *p.GoogleDomain})
	}
	if p.Gl != nil {
		values = append(values, param{"gl", *p.Gl})
	}
	if p.Hl != nil {
		values = append(values, param{"hl", *p.Hl})
	}
	if p.Cr != nil {
		values = append(values, param{"cr", *p.Cr})
	}
	if p.Lr != nil {
		values = append(values, param{"lr", *p.Lr})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// GoogleSearchLightParams holds the parameters of GoogleSearchLight (GET / POST /v1/google/search/light).
// Required fields are plain values; optional fields are pointers (use autom.String, autom.Float, autom.Bool) or nil slices and maps.
type GoogleSearchLightParams struct {
	// The search query you want to run on Google (supports advanced operators like site:, inurl:, intitle:).
	Query string `json:"query" api:"query,required"`
	// Google domain to use for the search. Defaults to google.com. Default: "google.com".
	GoogleDomain *string `json:"google_domain" api:"google_domain"`
	// Country to use for the Google search. Two-letter country code (e.g. us for the United States, uk for United Kingdom, or fr for France).
	Gl *string `json:"gl" api:"gl"`
	// Language to use for the Google search. Two or three-letter language code (e.g. en, es, fr), optionally followed by a region code (e.g. en-gb, es-419).
	Hl *string `json:"hl" api:"hl"`
	// Limit results to one or more countries. Use country{XX} with | as delimiter (e.g. countryFR|countryDE).
	Cr *string `json:"cr" api:"cr"`
	// Limit results to one or more languages. Use lang_{xx} with | as delimiter (e.g. lang_fr|lang_de).
	Lr *string `json:"lr" api:"lr"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *GoogleSearchLightParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 9)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.GoogleDomain != nil {
		values = append(values, param{"google_domain", *p.GoogleDomain})
	}
	if p.Gl != nil {
		values = append(values, param{"gl", *p.Gl})
	}
	if p.Hl != nil {
		values = append(values, param{"hl", *p.Hl})
	}
	if p.Cr != nil {
		values = append(values, param{"cr", *p.Cr})
	}
	if p.Lr != nil {
		values = append(values, param{"lr", *p.Lr})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// GoogleShoppingParams holds the parameters of GoogleShopping (GET / POST /v1/google/shopping).
// Required fields are plain values; optional fields are pointers (use autom.String, autom.Float, autom.Bool) or nil slices and maps.
type GoogleShoppingParams struct {
	// Query to search on Google Shopping (product name, brand, category).
	Query string `json:"query" api:"query,required"`
	// Page number for pagination (1 = first page). Default: 1.
	Page *float64 `json:"page" api:"page"`
	// Location to localize results (e.g. New York, United States). Generates a UULE parameter automatically.
	Location *string `json:"location" api:"location"`
	// Encoded location parameter (UULE). Use location instead if you don't have a pre-encoded value.
	Uule *string `json:"uule" api:"uule"`
	// Google domain to use for the search. Defaults to google.com. Default: "google.com".
	GoogleDomain *string `json:"google_domain" api:"google_domain"`
	// Country to use for the Google search. Two-letter country code (e.g. us for the United States, uk for United Kingdom, or fr for France). Default: "us".
	Gl *string `json:"gl" api:"gl"`
	// Language to use for the Google search. Two or three-letter language code (e.g. en, es, fr), optionally followed by a region code (e.g. en-gb, es-419). Default: "en".
	Hl *string `json:"hl" api:"hl"`
	// Limit results to one or more countries. Use country{XX} with | as delimiter (e.g. countryFR|countryDE).
	Cr *string `json:"cr" api:"cr"`
	// Limit results to one or more languages. Use lang_{xx} with | as delimiter (e.g. lang_fr|lang_de).
	Lr *string `json:"lr" api:"lr"`
	// Adult content filter. active enables Google SafeSearch; off disables it. Omit to use Google default.
	Safe *string `json:"safe" api:"safe"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *GoogleShoppingParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 13)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Page != nil {
		values = append(values, param{"page", *p.Page})
	}
	if p.Location != nil {
		values = append(values, param{"location", *p.Location})
	}
	if p.Uule != nil {
		values = append(values, param{"uule", *p.Uule})
	}
	if p.GoogleDomain != nil {
		values = append(values, param{"google_domain", *p.GoogleDomain})
	}
	if p.Gl != nil {
		values = append(values, param{"gl", *p.Gl})
	}
	if p.Hl != nil {
		values = append(values, param{"hl", *p.Hl})
	}
	if p.Cr != nil {
		values = append(values, param{"cr", *p.Cr})
	}
	if p.Lr != nil {
		values = append(values, param{"lr", *p.Lr})
	}
	if p.Safe != nil {
		values = append(values, param{"safe", *p.Safe})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// GoogleVideosParams holds the parameters of GoogleVideos (GET / POST /v1/google/videos).
// Required fields are plain values; optional fields are pointers (use autom.String, autom.Float, autom.Bool) or nil slices and maps.
type GoogleVideosParams struct {
	// The search query you want to run on Google Videos.
	Query string `json:"query" api:"query,required"`
	// Human-readable location to originate the search from. Note: location and uule cannot be used together.
	Location *string `json:"location" api:"location"`
	// Google encoded location parameter. Note: uule and location cannot be used together.
	Uule *string `json:"uule" api:"uule"`
	// Google domain to use for the search. Defaults to google.com. Default: "google.com".
	GoogleDomain *string `json:"google_domain" api:"google_domain"`
	// Country to use for the Google search. Two-letter country code (e.g. us for the United States, uk for United Kingdom, or fr for France).
	Gl *string `json:"gl" api:"gl"`
	// Language to use for the Google search. Two or three-letter language code (e.g. en, es, fr), optionally followed by a region code (e.g. en-gb, es-419).
	Hl *string `json:"hl" api:"hl"`
	// Limit results to one or more countries. Use country{XX} with | as delimiter (e.g. countryFR|countryDE).
	Cr *string `json:"cr" api:"cr"`
	// Limit results to one or more languages. Use lang_{xx} with | as delimiter (e.g. lang_fr|lang_de).
	Lr *string `json:"lr" api:"lr"`
	// Page number for pagination (1 = first page). Default: 1.
	Page float64 `json:"page" api:"page,required"`
	// Maximum number of results per page (e.g. 10). Default: 10.
	Num *float64 `json:"num" api:"num"`
	// Adult content filter. active enables Google SafeSearch; off disables it. Omit to use Google default.
	Safe *string `json:"safe" api:"safe"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *GoogleVideosParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 14)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Location != nil {
		values = append(values, param{"location", *p.Location})
	}
	if p.Uule != nil {
		values = append(values, param{"uule", *p.Uule})
	}
	if p.GoogleDomain != nil {
		values = append(values, param{"google_domain", *p.GoogleDomain})
	}
	if p.Gl != nil {
		values = append(values, param{"gl", *p.Gl})
	}
	if p.Hl != nil {
		values = append(values, param{"hl", *p.Hl})
	}
	if p.Cr != nil {
		values = append(values, param{"cr", *p.Cr})
	}
	if p.Lr != nil {
		values = append(values, param{"lr", *p.Lr})
	}
	values = append(values, param{"page", p.Page})
	if p.Num != nil {
		values = append(values, param{"num", *p.Num})
	}
	if p.Safe != nil {
		values = append(values, param{"safe", *p.Safe})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// JobsParams holds the parameters of Jobs (GET /v1/jobs/{id}).
// Required fields are plain values; optional fields are pointers (use autom.String, autom.Float, autom.Bool) or nil slices and maps.
type JobsParams struct {
	// The job_id returned by an endpoint called with is_async=true.
	ID string `json:"id" api:"id,required"`
}

func (p *JobsParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 1)
	if p.ID != "" {
		values = append(values, param{"id", p.ID})
	}
	return values
}

// SearchParams holds the parameters of Search (GET / POST /v1/search).
// Required fields are plain values; optional fields are pointers (use autom.String, autom.Float, autom.Bool) or nil slices and maps.
type SearchParams struct {
	// Search query.
	Query string `json:"query" api:"query,required"`
	// Engine order. Default auto tries google, brave, bing, duckduckgo, yahoo, then baidu. A comma-separated list is followed as given. ddg and duck mean duckduckgo. Unknown names are ignored. Default: "auto".
	Engines *string `json:"engines" api:"engines"`
	// Country code. Defaults to us when omitted. For example us, fr, gb, de.
	Gl *string `json:"gl" api:"gl"`
	// Language code. Defaults to en when omitted. For example en, fr, de.
	Hl *string `json:"hl" api:"hl"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *SearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 7)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Engines != nil {
		values = append(values, param{"engines", *p.Engines})
	}
	if p.Gl != nil {
		values = append(values, param{"gl", *p.Gl})
	}
	if p.Hl != nil {
		values = append(values, param{"hl", *p.Hl})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// StatsParams holds the parameters of Stats (GET /v1/stats).
// Required fields are plain values; optional fields are pointers (use autom.String, autom.Float, autom.Bool) or nil slices and maps.
type StatsParams struct {
	// API family to check, for example google. Provide at least engine or path.
	Engine *string `json:"engine" api:"engine"`
	// Specific endpoint to check, for example /v1/google/search. If you also pass engine, use an endpoint from that API.
	Path *string `json:"path" api:"path"`
}

func (p *StatsParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 2)
	if p.Engine != nil {
		values = append(values, param{"engine", *p.Engine})
	}
	if p.Path != nil {
		values = append(values, param{"path", *p.Path})
	}
	return values
}

// YahooSearchParams holds the parameters of YahooSearch (GET / POST /v1/yahoo/search).
// Required fields are plain values; optional fields are pointers (use autom.String, autom.Float, autom.Bool) or nil slices and maps.
type YahooSearchParams struct {
	// Search query sent to Yahoo Search.
	Query string `json:"query" api:"query,required"`
	// Page number. Page 1 is the first page of results. Default: 1.
	Page *float64 `json:"page" api:"page"`
	// Country code, for example us, fr, gb, de, jp.
	Gl *string `json:"gl" api:"gl"`
	// Language code, for example en, fr, de, ja, zh.
	Hl *string `json:"hl" api:"hl"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *YahooSearchParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 7)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Page != nil {
		values = append(values, param{"page", *p.Page})
	}
	if p.Gl != nil {
		values = append(values, param{"gl", *p.Gl})
	}
	if p.Hl != nil {
		values = append(values, param{"hl", *p.Hl})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}
