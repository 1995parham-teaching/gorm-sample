default:
    @just --list

# build gorm-sample binary
build:
    go build -o gorm-sample ./cmd/cli

# update go packages
update:
    @cd ./cmd/cli && go get -u

# set up the dev environment with docker-compose
dev cmd *flags:
    #!/usr/bin/env bash
    set -euxo pipefail
    if [ {{ cmd }} = 'down' ]; then
      docker compose -f ./docker-compose.yml down --remove-orphans
      docker compose -f ./docker-compose.yml rm
    elif [ {{ cmd }} = 'up' ]; then
      docker compose -f ./docker-compose.yml up --wait -d {{ flags }}
    else
      docker compose -f ./docker-compose.yml {{ cmd }} {{ flags }}
    fi

# run tests in the dev environment
test: (dev "up")
    just seed
    go run ./cmd/cli/main.go

seed: (dev "up")
    atlas migrate apply --env local

# connect into the dev environment database
database: (dev "up") (dev "exec" "database psql postgres://postgres:postgres@127.0.0.1:5432/pgsql?search_path=public&sslmode=disable")

# run golangci-lint
lint *flags:
    golangci-lint run -c .golangci.yml {{ flags }}
