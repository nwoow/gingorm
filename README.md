# Golang MVC-CRUD

To run the Code clone the code:

And then add dependecies this code uses:

  	"github.com/gin-gonic/gin" 
  	"github.com/jinzhu/gorm"
	
You can install all packages before running the code:

After installing packages simply run this:

	go run main.go


# Run the code using Docker:

If you wish to tun this code simple run this command to build image:

	docker build -t server .


If you wish to run the code simple run:

	docker run -p 8090:8080 server


To build docker image:

	docker build -t angelbee .

To run docker image with network host and port:

	docker run  -d --net=host -p 8080:8080 DT


Docker Remove all stopped container:

	docker rm $(docker ps -a -q)