.PHONY: build test lint clean generate

GO := go

build:
	$(GO) build ./...

test:
	$(GO) test -v ./...

lint:
	$(GO) vet ./...

clean:
	rm -rf dist/

generate:
	# Run from the ecosystems root scripts/generate_sdks.sh instead
	@echo "Generating SDK from OpenAPI..."
	cd .. && ./scripts/generate_sdks.sh
