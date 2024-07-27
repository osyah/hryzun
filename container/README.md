# `hryzun/container`

Simple and flexible dependency container package.

## Install

```
go get -u github.com/osyah/hryzun
```

## Example

```go
type Service interface {
    Send(string)
}

type UserService struct{}

func (UserService) Send(string) { /* ... */ }

func main() {
    base := container.New()

    base.Register("user_service", func(*container.Base) any {
        return &UserService{}
    })

    client := container.Get[Service](base, "user_service")
}
```
