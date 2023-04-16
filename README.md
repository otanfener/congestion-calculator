# Congestion Tax Calculator

In this project our aim is to write a program that calculates congestion tax based on the given [requirements](docs/ASSIGMENT.md)
## Design
### Database Design

MongoDB is selected as source of datastore for the simplicity and ease of use. Since MongoDB is a NoSQL database,
database design in this application also centered around establishing NOSQL design paradigms. As seen from the image below, we
modeled cities as one collection, and used city name as our primary access point through application.
We flattened our data as much as possible by placing different properties such as tariffs, rules of a city together.

![design.png](docs%2Fdesign.png)

In addition to above diagram, example data under `data` folder can also give idea about the structure.

### Application Design

Application modeled as standard Golang microservice program using repository pattern. With this pattern, we can easily use
any other database implementation in the future if we want to.
## Prerequisites 
- `make` 
  - Mac OS: `brew install make`
  - Linux(Ubuntu): `sudo apt install make`
- `docker` 
  - Mac OS: `brew install --cask docker`
  - Linux(Ubuntu): `sudo apt install docker-ce docker-ce-cli containerd.io docker-compose-plugin`

##  Usage
The provided `Makefile` will handle all the heavy lifting, and provide automated way to run the application. 
It has the following commands:
- `make build`: will create docker image for congestion calculator.
- `make up`: will run all services defined in `docker-compose.yml`. It will also import example dataset under `data` directory to `MongoDB`
- `make tools`: will download tooling binary used in application, and place them under `bin` folder.
- `make lint`: will lint application using `golangci-lint` tool.
### Postman Collection
Example postman collection can be found under `docs` folder. It can be directly imported to Postman.


