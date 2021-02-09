module github.com/yeqown/gorm-opentracing/tests

go 1.15

require (
	github.com/HdrHistogram/hdrhistogram-go v1.0.1 // indirect
	github.com/jinzhu/now v1.1.1
	github.com/opentracing/opentracing-go v1.2.0
	github.com/pkg/errors v0.9.1 // indirect
	github.com/uber/jaeger-client-go v2.25.0+incompatible
	github.com/uber/jaeger-lib v2.4.0+incompatible // indirect
	github.com/yeqown/gorm-opentracing v1.0.0
	go.uber.org/atomic v1.7.0 // indirect
	gorm.io/driver/mysql v1.0.4
	gorm.io/gorm v1.20.12
)

replace github.com/yeqown/gorm-opentracing => ../
