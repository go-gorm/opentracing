# gorm-opentracing

[![Go Report Card](https://goreportcard.com/badge/github.com/yeqown/gorm-opentracing)](https://goreportcard.com/report/github.com/yeqown/gorm-opentracing) [![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/yeqown/gorm-opentracing)

opentracing support for gorm2

### Features

- [x] Record `SQL` in `span` logs.

- [x] Record `Result` in `span` logs.
  
- [x] Record `Table` in `span` tags.

- [x] Record `Error` in `span` tags and logs.

- [x] Register `Create` `Query` `Delete` `Update` `Row` `Raw` tracing callbacks. 

### Get Started

```go
func main() {
	var db *gorm.DB
	
	db.Use(gormopentracing.New())
}
```

### Plugin options

```go
// WithLogResult log result into span log, default: disabled.
func WithLogResult(logResult bool)
```