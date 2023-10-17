VERSION := 0.0.1

.PHONY: dev
dev: dev-site dev-server
	@echo "Starting live reload setup for CatBook v$(VERSION)"

.PHONY: build
build: build-site build-server
	@echo "Building CatBook v$(VERSION)"

.PHONY: dev-site
dev-site:
	@cd server; go run .

.PHONY: dev-server
dev-server:
	@cd site; npm run dev

.PHONY: build-site
build-site:
	@cd site; npm run build

.PHONY: build-server
build-server:
	@cd server; go build