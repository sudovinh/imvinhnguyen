# imvinhnguyen

Source for [imvinhnguyen.com](https://imvinhnguyen.com) — Vinh Nguyen's personal
links / landing page. A single-page Go app (Echo + templ + Tailwind) with a
futuristic deep-navy theme, deployed to DigitalOcean App Platform.

## Status

[![CI](https://github.com/sudovinh/imvinhnguyen/actions/workflows/ci.yaml/badge.svg)](https://github.com/sudovinh/imvinhnguyen/actions/workflows/ci.yaml)
[![Deploy App](https://github.com/sudovinh/imvinhnguyen/actions/workflows/deploy-app.yaml/badge.svg)](https://github.com/sudovinh/imvinhnguyen/actions/workflows/deploy-app.yaml)

## Editing site content

All page content — bio, the animated identity line, social icons, the featured
video, and the link buttons — lives in **[`content/content.yaml`](content/content.yaml)**.
No Go code required.

To add a link button, add an entry under `quick_links`:

```yaml
quick_links:
  - label: My Ebay Store
    url: https://www.ebay.com/usr/vinhsellstuff
  - label: My YouTube          # <- new
    url: https://www.youtube.com/@imvinhnguyen
```

Commit and push to `main`; the site rebuilds and redeploys automatically
(~2 min). The file is validated in CI, so a malformed edit fails the build
instead of shipping a broken page.

## Development

This project uses [flox](https://flox.dev) to manage its dev environment
(Go, templ, tailwindcss, air).

```sh
flox activate            # enter the environment
PORT=8080 make watch     # live reload at http://localhost:8080
```

Or prefix individual commands: `flox activate -- make build`.

### Make targets

| Command | Description |
| --- | --- |
| `make build` | templ generate → Tailwind → `go build` |
| `PORT=8080 make run` | run without live reload |
| `PORT=8080 make watch` | run with live reload (air) |
| `make test` | run the test suite |
| `make clean` | remove the built binary |

## Deployment

Deployed to **DigitalOcean App Platform** as a Docker container, driven by
GitHub Actions:

- **`ci.yaml`** — vet, test, build, govulncheck, and a Docker boot-and-smoke test on every PR and push to `main`.
- **`deploy-app.yaml`** — deploys to production on push to `main`.
- **`deploy-preview.yaml`** / **`delete-preview.yaml`** — spin up and tear down an ephemeral preview app per PR.

The container image is built from the [`Dockerfile`](Dockerfile) (multi-stage,
CGO-free, non-root). Infrastructure is managed with Terraform in
[`terraform/`](terraform/); the app spec lives in [`.do/app.yaml`](.do/app.yaml).

```sh
make terraform-plan      # preview infra changes (1Password-backed DO token)
make terraform-apply
```

## Tech stack

- **[Go](https://go.dev)** + **[Echo](https://echo.labstack.com)** — HTTP server
- **[templ](https://templ.guide)** — type-safe HTML components
- **[Tailwind CSS](https://tailwindcss.com)** — styling
- **[modernc.org/sqlite](https://pkg.go.dev/modernc.org/sqlite)** — pure-Go SQLite (used only by `/health`; no CGO)
- **flox** · **Docker** · **Terraform** · **DigitalOcean App Platform**

