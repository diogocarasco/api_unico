# api_unico
<img align="right" width="159px" src="https://raw.githubusercontent.com/gin-gonic/logo/master/color.png">
A GIN-GONIC REST API made for the Unico technical testing.


The requirement to use this repo is `docker` and `docker-compose >=v1.27`, if you need
to install this plugin, follow these steps:

- install [Docker](https://docs.docker.com/engine/install/ubuntu/)
- install [Docker compose](https://docs.docker.com/compose/install/)

What you need to run is...

Start the database container
```sh
docker-compose up unico-mysql
```

Start the application
```sh
./run.sh app
```

For testing, run:
```sh
./run.sh test
```


