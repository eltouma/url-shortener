## URL SHORTENER ##
Easily convert a long URL into short URL.


## Features ##
- create a short URL from a long URL
- redirect from a short URL to the original URL
- in-memory storage
- basic collision handling
- concurrent access protection with a mutex
- minimal home page


## How it works ##
- the client sends a long URL to `POST /shorten`
- a unique short identifier is generated
- the mapping between the short URL and the long URL is stored in memory
- when `GET /<id>` is called, the client is redirected to the original URL
- `GET /` serves a minimal home page


## Run locally ##
clone the repository, then start the server:
```bash 
go run main.go
```

By default, the server runs on `http://localhost:8081`
You can configure the base URL with the `SITEURL` environment variable


```bash
curl -X  POST -d "url=http://example.com/your_very_long_url" http://localhost:8081
```

