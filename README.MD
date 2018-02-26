# TrueVal

This service provides a REST API to calculate the true to size value for a pair of shoes, getting you the correct size no matter what brand you are looking for.

# Dependencies and Environment

  - Written in [Go](https://golang.org/) (1.6 default repo version)
  - [Postgres](https://www.postgresql.org/)
  - [Docker](https://www.docker.com/)
  - [Gorilla Mux](https://github.com/gorilla/mux)
  - [Gorm](https://github.com/jinzhu/gorm)

### Instalation

The project comes with a configuration file (config.json) you will need to configure Postgres credentials, db name and on what port you like to run the HTTP server, in this file the settings for testing are supported as well. I recommend to keep a separate database for testing.

You can run the project localy or using docker-compose up.


### Known issues
* Gorm doesn't like mappings to json so that had to be done manually
* Sometimes Postgres stops allowing connections when running on Docker, this could be because I am using it on a VM but haven't figured out why