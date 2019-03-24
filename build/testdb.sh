docker run -d --name mongodb \
  -e MONGODB_USERNAME=reader -e MONGODB_PASSWORD=password \
  -e MONGODB_ROOT_PASSWORD=password \
  -e MONGODB_DATABASE=jedict \
  -v "$(pwd)"/build/data:/docker-entrypoint-initdb.d \
  -p 27017:27017 \
  bitnami/mongodb:latest