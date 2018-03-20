# flakeseries-example

> An example of writing to a FlakeSeries table with [gocassa](https://github.com/monzo/gocassa)

## Installation

```
make install-dependencies
```

## Usage

Start Cassandra in a Docker container:

```
make start-containers
```

Create the keyspace and table:

```
make create-tables
```

Run the script:

```
go run main.go "Hello world"
```

Read the data:

```
make cqlsh
select * from posts.posts_flakeseries;
```
