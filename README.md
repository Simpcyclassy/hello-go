## The assigment

For this assignment you can use the language and environment of choice between python, go or nodejs.

Note: Only use the standard library for each of the above. Third party dependencies are allowed only for testing or for development tools.

## Requirements

We ask you to create a simple web server that should do the following:

* Runs locally on port 8000 and accepts `GET` requests at the index URL `/`
* It checks that the request has a query parameter called `favoriteTree` with a valid value
* For a successful request, returns a properly encoded HTML document with the following content:

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

## Improvements

- [x] If there is no `favouriteTree` URL param return 'Please tell me your favorite tree'
- [x] Write tests
- [x] To use HTML templates (GO has a templating)
- [x] Add other end points
- [x] Use Go modules
- [ ] Restructure the folder (for example a separate file for handlers)
- [ ] Add concurrent requests
- [ ] Add a config and pass constants as command line arguments

