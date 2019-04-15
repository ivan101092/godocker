docker build . -t my-golang-app-image

docker run -p 3030:3001 -it --rm --name my-golang-app-run my-golang-app-image

docker-compose up