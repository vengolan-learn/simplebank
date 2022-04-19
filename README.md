# simplebank

## 1 Databases

### 1.1 docker commands to setup postgres

```shell
$docker pull postgres:alpine
$docker run --name mypostgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:alpine
```

Login to the console of the container instance

```shell
$docker exec -it  mypostgres bash 

//postgress client in the container shell 
#psql -U root

//check the logs in the container 
$docker logs mypostgres
```

### 1.2 Install golang-migrate

```shell
$curl -L <https://github.com/golang-migrate/migrate/releases/download/v4.15.1/migrate.linux-amd64.tar.gz> | tar xvz
$v migrate ~/go/bin/migrate

```

#### Create initial migration files

```shell
//dir that stores all  db related stuff
mkddr -p db/migration


~/go/bin/migrate create -ext sql -dir db/migration -seq init_schema
```

### 1.3 ORMs overview

`db\sql` -

`gorm` -similar to hibernate. Need to define entities as structs, and takes care

`sqlx` -

`sqlc` -

### 1.4 sqlc

#### Installation

### 1.5 Testing

#### Test Main

1. Install postgres driver

    sqlx is a generic library and we need to install database specific driver to use it

    ```shell
    //Install lib/pq driver
    $go get github.com/lib/pq
    ```

2. Install `testify` package to compare test results

    ```shell
    $go get github.com/stretchr/testify
    ```
