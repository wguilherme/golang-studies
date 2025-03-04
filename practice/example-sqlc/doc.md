go install github.com/jackc/tern/v2@latest
go env GOPATH
echo 'export PATH="$PATH:$(go env GOPATH)/bin"' >> ~/.zshrc

tern init
apagar migration de ex. gerada


após configurar arquivo tern.conf gerado pelo comando anterior, rodar a migration:


tern new create_table_authors para gerar arquivo de migration.


para aexecutar, podemos rodar tern migrate
caso os arquivos nao estejam na raiz, rodar tern migrate --migrations <path(arquivos sql)>  --config <path(arquivo tern.conf)>

no arquivo sqlc.yaml poderia informar o diretório de migraitons no lugar do arquivo único 001_create_table_authors.sql


## Executando migration

```bash
tern migrate --migrations ./sql/migrations --config ./tern.conf
```

output:

```bash
(⎈|N/A:default)➜ example-sqlc (main) ✗ tern migrate
2025-03-04 17:11:57 executing 001_create_table_authors.sql up
-- Write your migrate up statements here
CREATE TABLE authors (
  id   BIGSERIAL PRIMARY KEY,
  name text      NOT NULL,
  bio  text
);
``` 



após criar o arquivo query.sql com as queries, gerar código GO com: 
ps: garantir que o sqlc esteja instalado e requer go v1.21+

```bash
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
```

```bash
sqlc generate
```


## Comando que atualiza todas as dependências do projeto

```bash
go get -u ./...
```