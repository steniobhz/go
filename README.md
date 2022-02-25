# go
docker build -t go-app-img .
docker run -d -p 3333:3000 --name go-app-container go-app-img
