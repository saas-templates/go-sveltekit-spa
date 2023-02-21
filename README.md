# Go + SvelteKit (Static)

## Goals

1. `/` should be pre-rendered since this acts as landing page with static content.
2. `/app/*` route hierarchy must be login-protected. Since this is dynamic, this whole part can be SPA.
3. `go build main.go` should produce a single binary with UI embedded.
4. Logged in user is identified by calling `/api/me` endpoint which uses cookies (not managed by svelte) to restore session and return user or 401 (logged out).

No. 1 & 2 are done by setting `prerender=true` in `/+layout.ts` and `fallback: "app/index.html"` in adaptor config. This produces:

```plaintext
build
├── _app
├── app
│   └── index.html       -- acts as the SPA entrypoint, all child routing is client-side
├── favicon.png
└── index.html           -- prerendered html file
```

## Issue

Even with the auth check guard, the `load()` function in `src/routes/app/+page.ts` gets invoked even when the user is not logged in.

This can be confirmed by running the go binary with `-nologin` flag.
