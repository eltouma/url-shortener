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
By default, the server runs on: `http://localhost:8081`<br>
You can configure the base URL with the `SITEURL` environment variable
```bash
export SITEURL=http://your_own_site
```


clone the repository, then start the server:
```bash 
go run main.go
```

In another terminal window, create a short URL from a long one
```bash
curl -X  POST -d "url=http://example.com/some_very_long_url" http://localhost:8081
```
Response
```bash
http://localhost/8081/short_url
```

## Limitation ##
Since it's in-memory storage only, data is lost when the server stops.
