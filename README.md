# googleFileProxy
## Running with Docker
1. Clone the Repo
2. cd in the project folder
3. run `docker build -t googleproxy ./` and wait for image to build
4. run `docker run -d -i -t --name googleproxy -p 8080:8080 googleproxy` to start the container in daemon mode
5. navigate to http://localhost:8080/<endpoint of choice>

## Available endpoints
1. `/download` - endpoint for downloading a file
2. `/info` - endpoint for getting the information (metadata) of a file
3. `/list` - endpoint for listing the information for all files existing in the specified bucket
4. `/timesAccessed` - endpoint for listing all available filenames in the bucket and the number of times they were downloaded
5. `/accessorInfo` - endpoint for listing all available filenames in the bucket and the ips and agents of the ones requesting the file download

### /download
#### QueryParams:
1. `bucket` - default value defined in .env file
2. `file` - filename to be downloaded (with extension e.g. atlantis.avi)

### /info
#### QueryParams:
1. `bucket` - default value defined in .env file
2. `file` - filename to be queried (with extension e.g. atlantis.avi)

### /list
#### QueryParams:
1. `bucket` - default value defined in .env file
2. `perPage` - number of items per page (default 50)
3. `page` - page number (default 1)

### /timesAccessed
#### QueryParams:
1. `bucket` - default value defined in .env file

### /accessorInfo
#### QueryParams:
1. `bucket` - default value defined in .env file

## CLI command

### How to run the command:
`docker exec googleproxy ./command`

### Available Flags for the command:
1. `perPage` - number of items per page (default 50)
2. `page` - page number (default 1)
3. `all` - if present pagination is ignored and all items are getting returned
4. `namesOnly` - if present only the filenames are getting returned instead of the complete set of metadata
