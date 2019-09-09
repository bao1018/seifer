default: migrate-up run

run:
	.bin/air

create-migrate:
	.bin/migrate create -ext sql -dir migration -seq $(seq)

migrate-up:
	.bin/migrate -database ${pg_str} -path migration up


preset: preset-air preset-migrate

preset-air:
	mkdir -p .bin
	curl -fLo .bin/air https://raw.githubusercontent.com/cosmtrek/air/master/bin/darwin/air 
	chmod +x .bin/air

preset-migrate:
	mkdir -p .bin
	curl -fLo .bin/migrate.tar.gz https://github.com/golang-migrate/migrate/releases/download/v4.6.1/migrate.darwin-amd64.tar.gz
	tar xvzf .bin/migrate.tar.gz -C .bin
	mv .bin/migrate.darwin-amd64 .bin/migrate
	chmod +x .bin/migrate
	rm .bin/migrate.tar.gz
