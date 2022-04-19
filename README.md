# GO Simple clean architecture

Opensource Package:
- GORM
- VIPER
- ECHO
- MYSQL

## Instalation Step
```
#clone this repo
git@github.com:WailanTirajoh/go-simple-clean-architecture.git

#create .env from .example
cp .env.example .env

#create database
mysql -u root -p

create database 2022_godb;
exit;

go get .
go run .
```

## Code Pattern
This code style follow repository pattern, bootstrap with dependency injection from config, to repository, to service, to controller that will be consume by routes.


## Whats includes?
- Router (Echo)
- Middleware (Echo)
- Controller
- Model
- Request Handler
- JWT Authentication
- 

## What will be include ?
- Gates Permission / Authorization
- Event & listener
- Job & queue
- Tracer