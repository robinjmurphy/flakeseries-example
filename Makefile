.PHONY: install-dependencies
install-dependencies:
	go get

.PHONY: start-containers
start-containers:
	docker-compose up -d

.PHONY: create-tables
create-tables:
	docker-compose exec cassandra cqlsh -f /cql/schema.cql

.PHONY: truncate-tables
truncate-tables:
	docker-compose exec cassandra cqlsh -e 'truncate posts.posts_flakeseries;'

.PHONY: drop-tables
drop-tables:
	docker-compose exec cassandra cqlsh -e 'drop keyspace posts;'

.PHONY: cqlsh
cqlsh:
	docker-compose exec cassandra cqlsh
