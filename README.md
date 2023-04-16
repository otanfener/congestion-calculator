# Congestion Tax Calculator

The aim of this project is to write a program that calculates congestion tax based on the given [requirements](docs/ASSIGMENT.md).

For assumptions and questions made during the scope of this project please see [questions](docs/QUESTIONS.md)
## Design
### Database Design

MongoDB was chosen as the datastore for its simplicity and ease of use. As MongoDB is a NoSQL database, the database design for this application was based on establishing NoSQL design paradigms. As shown in the image below, we modeled cities as one collection and used the city name as our primary access point through the application. We flattened our data as much as possible by grouping different properties, such as tariffs and rules of a city, together. With this design, we add any other city easily, and in addition we can also modify rules for a city as well.

![design.png](docs%2Fdesign.png)

In addition to the diagram above, the example dataset under the `data` folder can also provide an idea about the structure.
### Application Design
The application is modeled as a standard Golang microservice program using the repository pattern. With this pattern, we can easily use any other database implementation in the future if we want to.
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
- `make down`: will stop all services and cleanup. 
- `make tools`: will download tooling binary used in application, and place them under `bin` folder.
- `make lint`: will lint application using `golangci-lint` tool.

Please make sure that Docker daemon is up and running.
### Postman Collection
Example postman collection can be found under `docs` folder. It can be directly imported to Postman.

