package main

import (
	"fmt"
)

type RenderType string

const (
	Static RenderType = "SSG"
	Server RenderType = "SSR"
	Client RenderType = "SPA"
)

// RequestContext represents the incoming traffic metadata.
type RequestContext struct {
	Path        string
	IsBot       bool // SEO Crawler
	IsAuth      bool // Authenticated user
}

// StrategyEngine decides the rendering path at the Edge.
func StrategyEngine(ctx RequestContext) RenderType {
	// 1. If it's a bot, we MUST provide pre-rendered content (SEO).
	if ctx.IsBot {
		return Server
	}

	// 2. If it's a dynamic user dashboard, we need fresh data.
	if ctx.IsAuth {
		return Client
	}

	// 3. For public marketing pages, serve from Edge Cache (Max Performance).
	return Static
}

func main() {
	// Example scenarios
	seoRequest := RequestContext{Path: "/home", IsBot: true}
	userRequest := RequestContext{Path: "/dashboard", IsAuth: true}

	fmt.Printf("Path: %s -> Strategy: %s\n", seoRequest.Path, StrategyEngine(seoRequest))
	fmt.Printf("Path: %s -> Strategy: %s\n", userRequest.Path, StrategyEngine(userRequest))
}
