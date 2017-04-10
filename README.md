Orchid-micro
=================
Golang boilerplate using gin-gonic framework and gorm for microservice

### Installation
* Go to your `$GOPATH/src` and clone the directory using `git clone https://github.com/thedevsaddam/orchid.git`
or [download the zip file](https://github.com/thedevsaddam/orchid-micro/archive/master.zip)

* Install dependency manager `govendor` using the command below
    ```
    go get -u github.com/kardianos/govendor
    ```

* Go to the `$GOPATH/src/orchid/vendor` directory and install dependencies using `govendor sync` command

* Copy `.env.example` to `.env` and set your configurations.

* Run `go build` to build binary file and to start the application use `./orchid`

### Todo
- [ ] Job Queue
- [ ] Caching
- [ ] Localization
- [ ] Helpers
- [ ] OAuth2 server or JWT
- [ ] Fixing inconsistent codes
- [ ] Request validation
- [ ] Security
- [ ] Find out performance issues
- [ ] Benchmarking

### Credits
* Routing, middleware, route-group [gin-gonic](https://gin-gonic.github.io/gin)
* Object-relational mapping [gorm](https://github.com/jinzhu/gorm)
* Dependency management package [govendor](https://github.com/kardianos/govendor)
* Environment management  package [godotenv](https://github.com/joho/godotenv)

### License
The **Orchid-micro** is a open-source software licensed under the [MIT License](LICENSE.md).
