# gorm-opentracing

opentracing support for gorm2

### Features

[ ] Record `SQL` in `span` logs.

[ ] Record `Table` in `span` tags.

[ ] Record `Error` in `span` tags and logs.

[ ] Register `Create` `Query` `Delete` `Update` `Row` `Raw` tracing callbacks. 

### Get Started

```go
func main() {
	var db *gorm.DB
	
	db.Use(gormopentracing.New())
}
```