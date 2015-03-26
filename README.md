Golang Web Application with Mongo connection
# Build Image
docker build -t daocloud/go-mongo .


## Below Mongo Connection Env must be set
* MONGO_NAME  /* Mongo DB Name */
* MONGO_CONN  /* Mongo DB Connection */

# Run Container
docker run --link your_mongo:mongo -e MONGO_CONN=user:pass@localhost:2345 -e MONGO_NAME=daocloud -d -p 80:80 daocloud/go-mongo

# That's it
