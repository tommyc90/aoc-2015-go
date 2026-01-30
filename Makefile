.PHONY: run
run:
	@go run ./cmd/app $(d) $(p) $(f)

.PHONY: debug
debug:
	@dlv debug ./cmd/app \
		--headless \
		--listen=:2345 \
		--api-version=2 \
		--continue=false \
		-- $(d) $(p) $(f)
