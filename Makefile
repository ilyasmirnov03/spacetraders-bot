build:
	rm -rf dist
	mkdir dist && cp .env dist/
	go build -o dist/spacetraders-bot