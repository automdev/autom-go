// Package autom is the official Go SDK for the Autom API (https://www.autom.dev).
//
// Google, Bing and Brave search results as structured JSON — web, maps, news, shopping, images, jobs and AI Mode.
//
// Create one client and reuse it:
//
//	client, err := autom.NewClient("", nil) // reads AUTOM_API_KEY
//	response, err := client.GoogleSearch(ctx, ...)
//
// Every documented endpoint has a typed method; Client.Request reaches any other path.
package autom
