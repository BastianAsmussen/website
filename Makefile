.PHONY: all build css dev clean

all: css build

css:
	npm install
	npx tailwindcss -i static/css/input.css -o static/css/style.css --minify

build:
	go build -o bin/website .

dev:
	npx tailwindcss -i static/css/input.css -o static/css/style.css --watch &
	go run .

clean:
	rm -rf bin/ node_modules/ static/css/style.css
