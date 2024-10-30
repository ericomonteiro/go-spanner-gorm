# Demo application for:
- GoLang
- GCP Spanner
- Gorm

## References:
- [GCP Spanner](https://cloud.google.com/spanner)
- [Gorm](https://gorm.io/)
- [Gorm Spanner](https://cloud.google.com/spanner/docs/use-gorm?hl=pt-br)

### To execute the application:
- Clone the repository
- Run the following command:

The docker-compose file will start a local instance of the Spanner emulator and execute migrations.
```bash
docker-compose up -d
```

```bash
go mod tidy
```

```bash
go run main.go
```