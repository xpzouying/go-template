.PHONY: b
b:
	cd web && npm install && npm run build && cd - && \
	GIN_MODE=release /usr/local/go/bin/go build -ldflags "-s -w" -o ./build/go-template .
