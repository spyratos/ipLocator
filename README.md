# ipLocator
## Running with Docker
1. Clone the Repo
2. cd in the project folder
3. run `docker build -t iplocator ./` and wait for image to build
4. run `docker run -d -i -t --name iplocator -p 8080:8080 iplocator` to start the container in daemon mode
5. navigate to http://localhost:8080/<endpoint of choice>
6. You can ssh in the container to check the log files with `docker exec -it iplocator /bin/ash`

## Available endpoints
1. `/locate` - POST endpoint for locating an IP

### /locate
#### BodyParams:
1. `ip` - ip you wish to locate


Content-Type header of the incoming request should be "application/json". It is assumed that this API won't accept "application/x-www-form-urlencoded" or "multipart/form-data" content-types.
