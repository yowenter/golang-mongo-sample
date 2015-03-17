Golang Web Application with Mongo connection
# Build Image
docker build -t daocloud/go-mongo .

# Run Container
docker run -d -p 80:80 daocloud/go-mongo

# That's it
