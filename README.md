## I-Invest

### Backend

#### Dependencies
- go 1.22 (Use gvm to make it easir changing versions on the fly)
- air (development only)
- Gin framework
- Postgres 16
- Metabase (For reports)
- Migrate CLI Tool latest (github.com/golang-migrate/migrate/cmd/migrate)

#### Instalation

```bash
cd backend && go mod tidy
```

```bash
cd infrastructure && ./up-infra.sh
```

#### Migrations
Install the migrate tool manually, it's not setup in go.mod since we made the decision to separate de migration process from the binary itself e execute as migrations

```bash
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
cd backend && ./migrate_db.sh up
```