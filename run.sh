docker build -t bff-go .
docker run -d -p 8080:8080 bff-go