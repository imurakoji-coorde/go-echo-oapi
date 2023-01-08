# oapi-gen
go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v1.11.0

# go install一覧
ls -l $(go env GOPATH)/bin

# oapi-codegen config

oapi-codegen -config api/models.cfg.yml oas.yml
oapi-codegen -config api/server.cfg.yml oas.yml

# docler-compose down with network
docker-compose down --remove-orphans