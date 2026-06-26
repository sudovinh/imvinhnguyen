# Simple Makefile for a Go project

# Build the application
# Tooling (go, templ, tailwindcss, air) is provided by the flox environment.
# Run inside `flox activate` (or prefix commands with `flox activate -- ...`).
all: build test

build:
	@echo "Building..."
	@templ generate
	@tailwindcss -i cmd/web/styles/input.css -o cmd/web/assets/css/output.css
	@go build -o main cmd/api/main.go

# Run the application
run:
	@go run cmd/api/main.go

# Test the application
test:
	@echo "Testing..."
	@go test ./... -v

# Clean the binary
clean:
	@echo "Cleaning..."
	@rm -f main

# Live Reload (air is provided by the flox environment)
watch:
	@air

OP_TF = op run --env-file=.env.terraform --

terraform-init:
	cd terraform && terraform init

terraform-plan:
	cd terraform && $(OP_TF) terraform plan

terraform-apply:
	cd terraform && $(OP_TF) terraform apply

terraform-destroy:
	cd terraform && $(OP_TF) terraform destroy

docker-build:
	docker build -t imvinhnguyen:local .

docker-run:
	docker run --rm -p 8080:8080 imvinhnguyen:local

.PHONY: all build run test clean watch \
        terraform-init terraform-plan terraform-apply terraform-destroy \
        docker-build docker-run
