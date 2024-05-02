# `hryzun/status`

Simple and flexible error handling package.

## Install

```
go get -u github.com/osyah/hryzun
```

## Example

### Status Code

```go
const (
    CodeNotFound status.Code = iota
    CodeInternal
    // ...
)

func example() error {
    return CodeNotFound
}
```

### Status

```go
func example() error {
    return status.New(CodeNotFound, "example not found")
}
```

### Handler

```go
func handler(err error) {
    switch t := err.(type) {
    case status.Status:
        // ...
    case status.Code:
        // ..
    default:
        // ...
    }
}
```
