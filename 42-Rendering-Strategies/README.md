# Architectural Manifesto: Rendering Strategies (SPA, SSR, SSG)

## 1. The Engineering Paradigm
Rendering strategy defines the trade-off between "Time to First Byte" (TTFB), "Time to Interactive" (TTI), and "Server Computational Overhead." 

## 2. Comparison Matrix (The Principal's Decision Matrix)
| Strategy | Best For | SEO | Runtime Cost | UX Strength |
| :--- | :--- | :--- | :--- | :--- |
| **SPA** | Dashboard, SaaS, Apps | Low (unless hashed) | Low (CDN only) | High (Transitions) |
| **SSR** | Dynamic, SEO-heavy apps | High | High (Server Load) | Medium (FCP) |
| **SSG** | Docs, Blogs, Marketing | Maximum | Zero (Static) | Maximum (CDN) |

## 3. The "Hydration" Problem
The modern architectural challenge is **Hydration**. When using SSR/SSG, the server delivers HTML, but the client must "re-hydrate" (attach event listeners). This creates a "dead zone" where the user sees the page but cannot interact. We minimize this by using "Islands Architecture" or "Partial Hydration."

## 4. Operational Strategy
- Default to SSG for marketing and content.
- Use SSR for personalized, authenticated dynamic data.
- Use SPA only for application-heavy interfaces where the initial load penalty is acceptable.
