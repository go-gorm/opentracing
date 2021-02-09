module github.com/yeqown/gorm-opentracing/tests

go 1.15

require (
	github.com/jinzhu/now v1.1.1
	github.com/opentracing/opentracing-go v1.2.0
	github.com/uber/jaeger-client-go v2.25.0+incompatible
	github.com/yeqown/gorm-opentracing v1.0.0
	gorm.io/driver/mysql v1.0.4
	gorm.io/gorm v1.20.12
)

replace github.com/yeqown/gorm-opentracing => ../
