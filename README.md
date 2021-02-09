# gorm-opentracing

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