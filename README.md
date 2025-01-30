# Project Structure:
This project follows DDD structure and is heavily inspired from:
- https://github.com/sklinkert/go-ddd
- https://github.com/percybolmer/ddd-go.git
    - branch clean-architecture

# Dependencies:

### ORM
- [sqlc](https://github.com/sqlc-dev/sqlc)
---
### MIGRATION
- [goose](https://github.com/pressly/goose)
- [goose doc](https://pressly.github.io/goose/documentation/annotations/)
        
    - **Usage:** \
    goose {Flag} {Command} {File-name} {File-format} \
    goose -dir cmd/migrate create user-table sql
    
    - **Note:** \
    More complex statements (PL/pgSQL) that have `semicolons` within them must be annotated with `-- +goose StatementBegin` and `-- +goose StatementEnd` to be properly recognized.
---
### HOT_RELOAD
- [air](https://github.com/air-verse/air)
    - **Issue 1:**\
    add this line to `.bashrc` ===> **`alias air='$(go env GOPATH)/bin/air'`**