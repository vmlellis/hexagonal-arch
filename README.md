# Hexagonal-arch

- Generating the mock:

```bash
mockgen -destination=application/mocks/application.go -source=application/contract/product.go application
```

- Run the tests:

```bash
go test -v ./...
```
