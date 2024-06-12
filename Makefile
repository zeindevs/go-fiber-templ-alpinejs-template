run: build
	./bin/main.exe

build: winres-make
	go build -ldflags="-w -s" -o bin/main.exe .

air:
	air -c .air.toml

templ:
	templ generate --watch

css:
	tailwindcss -i public/style.scss -o static/style.css -c public/tailwind.config.js --minify --watch