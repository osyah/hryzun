# `hryzun/config`

Simple, flexible, and fast configuration engine.

## Install

```
go get -u github.com/osyah/hryzun/config
```

## Context

To reduce the number of application dependencies, the package contains adapters — each responsible for predefined data formats. Allowing developers to quickly and easily add support for necessary formats, avoiding adding all possible dependencies.

## Best practices

### Naming convention

Using the `snake_case` naming convention for all configuration keys is the best solution, given all possible data formats.

### Final format

It is recommended to use configuration tools that are able to convert content into more common data formats (for example, JSON), this will avoid specific and unpopular dependencies.

## Example

### Simple 

```json
// config.json
{
  "address": "localhost:8000",
  "body_limit": 1024
}
```

```go
type Config struct{
    Address   string `cfg:"address"`
    BodyLimit int    `cfg:"body_limit"`
}

func main() {
    cfg, err := config.New[Config]("config.json")
    if err != nil { /* ... */ }

    // ...
}
```

Please note that the use of `cfg` annotations is mandatory for each field of the structure to be decoded, otherwise it may lead to an unexpected result.

### With custom adapter

```go
type Adapter struct{}

func (Adapter)  Peek(s string) bool { /* ... */}

func (Adapter) Unmarshal(data []byte, v any) error { /* ... */}

func main() {
    cfg, err := config.New[Config]("config.json", &Adapter{})
    if err != nil { /* ... */ }
}
```

### Simple unmarshal

```go
result, err := config.Unmarshal(&UnmarshalInput{
    Type: ".json",
    Data: []byte(`{"hello": "world"}`)
    Adapters: []config.Adapter{&config.JSON{}},
})
if err != nil { /* ... */ }
```
