## The assignment

## Requirements

We ask you to create a simple web server that should do the following:

- Runs locally on port 8000 and accepts `GET` requests at the index URL `/`
- It checks that the request has a query parameter called `favoriteTree` with a valid value
- For a successful request, returns a properly encoded HTML document with the following content:

If `favoriteTree` was specified (e.g. a call like `127.0.0.1:8000/?favoriteTree=baobab`):

```
It's nice to know that your favorite tree is a <value of "favoriteTree" from the url>
```

if not specified (e.g. a call like `127.0.0.1:8000/`):

```
Please tell me your favorite tree
```

## Run tests

```
go test ./...
```

## Run tests Verbose

```
go test ./... -v
```

## Run tests with coverage of all packages

```
go test -coverpkg=all ./...
```

## Call the service with CURL

```sh
curl localhost:8000/tree?favoriteTree=baobab
```

# Running with Docker

Build Docker container

```sh
docker build -t server-go .
```

Look for the docker images available

```sh
docker images
```

```sh
docker run -p 8000:8000 -it server-go
```

Look for the docker images available

```sh
docker images
```

```sh
docker run -p 8000:8000 -it server-go
```

## Improvements

- [x] If there is no `favouriteTree` URL param return 'Please tell me your favorite tree'
- [x] Write tests
- [x] To use HTML templates (GO has a templating)
- [x] Add other end points
- [x] Use Go modules
- [x] Restructure the folder (for example a separate file for handlers)
- [x] Add concurrent requests

## Next steps

- [ ] Add Database / Cache
- [ ] Get data from upstream APIs, save to DB and return to user
- [ ] Add post/update/delete endpoints
- [ ] Add authentication and user types (JWT)

- [ ] Add a config and pass constants as command line arguments
- [ ] Add metrics (Promethus & Grafana)
- [ ] Run in K8s

## Notes

* Covid data api expects country, doesn't work with city
* service not erroring if responses not found - will send empty body
* panics if we return empty byte array from makeCountryInfoRequests
    - how to recover panics and log
    - writing checks to avoid the panic
* showing coverage from a separate tests package

