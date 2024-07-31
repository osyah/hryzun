# `hryzun/container`

Simple and flexible dependency container package.

## Install

```
go get -u github.com/osyah/hryzun
```

## Example

### Basic

```go
type Service interface {
    Send(string)
}

type UserService struct{}

func (UserService) Send(string) { /* ... */ }

func main() {
    base := container.New()

    base.RegisterHandler("user_service", func(*container.Base) any {
        return &UserService{}
    })

    client := container.Get[Service](base, "user_service")
}
```

### With closer

```go
type Database interface {
    Close() error
}

type Postgres struct{}

func (Postgres) Close() error { /* ... */ }

func main() {
    base := container.New()

    base.RegisterHandler("postgres", func(b *container.Base) any {
        var postgres Postgres

        b.RegisterCloser("postgres", postgres.Close)

        return &postgres
    })

    postgres := container.Get[Database](base, "postgres")
    // ...

    if err := base.Close(); err != nil {
        // ...
    }
}
```
