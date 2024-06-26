.DEFAULT_GOAL := build

brew:
	brew install golangci-lint
	brew install staticcheck
	brew install gofumpt
	brew install protobuf

install:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/google/addlicense@latest

clean:
	rm -rf dist/
	rm -rf tmp/
	rm -f coverage.out
	rm -f result.json

init-dependency:
	go get -u golang.org/x/crypto@v0.16.0
	go get -u github.com/go-playground/validator/v10@v10.15.1
	go get -u github.com/google/uuid@v1.3.1
	go get -u github.com/google/go-cmp/cmp@v0.6.0
	go get -u github.com/davecgh/go-spew/spew@v1.1.1
	go get -u github.com/xeipuuv/gojsonschema@v1.2.0
	go get -u github.com/jinzhu/copier@v0.4.0
	go get -u go.uber.org/zap@v1.26.0
	go get -u github.com/go-playground/validator/v10
	go get -u github.com/pressly/goose/v3
	go get -u github.com/lib/pq
	go get -u github.com/jackc/pgx
	go get -u gorm.io/driver/postgres
	go get -u moul.io/zapgorm2
	go get -u github.com/dgraph-io/badger/v4
	go get -u google.golang.org/grpc@v1.59.0
	go get -u github.com/spf13/cobra@v1.8.0
	go get -u github.com/spf13/viper
	go get -u github.com/spf13/viper v1.18.2
	go get -u github.com/stretchr/testify@v1.9.0
	go get -u github.com/fatih/color@v1.16.0
	go get -u get gopkg.in/yaml.v2@v2.4.0
	go get -u github.com/DATA-DOG/go-sqlmock@1.5.2

mod:
	go mod download
	go mod tidy

protoc:
	protoc internal/agents/services/aap/endpoints/api/v1/*.proto --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative --proto_path=.
	protoc internal/agents/services/pap/endpoints/api/v1/*.proto --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative --proto_path=.
	protoc internal/agents/services/pdp/endpoints/api/v1/*.proto --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative --proto_path=.

check:
	staticcheck  ./...

lint:
	go vet ./...
	gofmt -s -w **/**.go
	gofumpt -l -w .
	golangci-lint run --disable-all --enable staticcheck

lint-fix:
	gofmt -s -w **/**.go
	go vet ./...
	gofumpt -l -w .
	golangci-lint run ./... --fix

test:
	go test ./...

teste2e:
	export E2E="TRUE" && GOFLAGS="-count=1" go test ./e2e/...

coverage:
	go test -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out
	go tool cover -html=coverage.out
	rm coverage.out

coverage-plugin:
	go test -coverprofile=coverage.out ./plugin/...
	go tool cover -func=coverage.out
	go tool cover -html=coverage.out
	rm coverage.out

converage-%:
	go test -coverprofile=coverage.out ./...

converage-json:
	go test -json -coverprofile=coverage.out ./... > result.json

build-release:
	mkdir -p dist
	go build -o dist/server-all-in-one ./cmd/server-all-in-one
	go build -o dist/server-aap ./cmd/server-aap
	go build -o dist/server-pap ./cmd/server-pap
	go build -o dist/server-idp ./cmd/server-idp
	go build -o dist/server-pip ./cmd/server-pip
	go build -o dist/server-prp ./cmd/server-prp
	go build -o dist/server-pdp ./cmd/server-pdp
	go build -o dist/permguard ./cmd/cli

build-docker:
	docker stop permguard || true && docker rm permguard || true
	docker build -t permguard .

run-release:
	go run ./cmd/server-all-in-one

run-docker:
	docker run --name permguard permguard

build:  clean mod build-release

run:  clean mod lint-fix run-release

docker:  clean mod lint-fix run-docker

postgres-up:
	docker-compose -f ./docker-compose/docker-compose-postgres.yml -p permguard-postgres up -d
postgres-down:
	docker-compose -p permguard-postgres down

# disallow any parallelism (-j) for Make. This is necessary since some
# commands during the build process create temporary files that collide
# under parallel conditions.
.NOTPARALLEL:

.PHONY: clean mod lint lint-fix release alll
