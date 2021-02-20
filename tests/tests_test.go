package tests_test

import (
	"errors"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	jaegerlog "github.com/uber/jaeger-client-go/log"
	gormopentracing "gorm.io/plugin/opentracing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	. "gorm.io/gorm/utils/tests"
)

var DB *gorm.DB

func init() {
	var err error
	if DB, err = openTestConnection(); err != nil {
		log.Printf("failed to connect database, got error %v", err)
		os.Exit(1)
	}

	// prepare Database
	sqlDB, err := DB.DB()
	if err == nil {
		err = sqlDB.Ping()
	}
	if err != nil {
		log.Printf("failed to connect database, got error %v", err)
		os.Exit(1)
	}

	runMigrations()
	bootTestTracer()
	usePlugin()
}

func bootTestTracer() {
	// DONE(@yeqown) use jaeger tracer
	cfg := &config.Configuration{
		Sampler: &config.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		ServiceName: "gormopentracing",
		Reporter: &config.ReporterConfig{
			LogSpans: true,
			//LocalAgentHostPort:  "127.0.0.1:6381",
			BufferFlushInterval: 100 * time.Millisecond,
			CollectorEndpoint:   "http://127.0.0.1:14268/api/traces",
		},
	}

	tracer, _, err := cfg.NewTracer(
		config.Logger(jaegerlog.StdLogger),
		config.ZipkinSharedRPCSpan(true),
	)
	if err != nil {
		log.Printf("failed to use jaeger tracer plugin, got error %v", err)
		os.Exit(1)
	}

	opentracing.SetGlobalTracer(tracer)

	cfg2 := &config.Configuration{
		Sampler: &config.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		ServiceName: "mysql",
		Reporter: &config.ReporterConfig{
			LogSpans: true,
			//LocalAgentHostPort:  "127.0.0.1:6381",
			BufferFlushInterval: 100 * time.Millisecond,
			CollectorEndpoint:   "http://127.0.0.1:14268/api/traces",
		},
	}
	tracer2, _, err = cfg2.NewTracer(
		config.Logger(jaegerlog.StdLogger),
		config.ZipkinSharedRPCSpan(true),
	)
	if err != nil {
		log.Printf("failed to use jaeger tracer plugin, got error %v", err)
		os.Exit(1)
	}
}

var (
	tracer2 opentracing.Tracer
)

func usePlugin() {
	var p gorm.Plugin

	withTracer := os.Getenv("WITH_TRACER2")
	if withTracer == "" {
		p = gormopentracing.New(
			gormopentracing.WithLogResult(true),
		)
	} else {
		p = gormopentracing.New(
			gormopentracing.WithLogResult(true),
			gormopentracing.WithTracer(tracer2),
		)
	}

	if err := DB.Use(p); err != nil {
		log.Printf("failed to use gormopentracing plugin, got error %v", err)
		os.Exit(1)
	}
}

func openTestConnection() (db *gorm.DB, err error) {
	dbDSN := os.Getenv("GORM_DSN")
	dialect := os.Getenv("GORM_DIALECT")
	if dialect == "" {
		dialect = "mysql"
	}
	switch dialect {
	case "mysql":
		log.Println("testing mysql...")
		if dbDSN == "" {
			dbDSN = "gorm:gorm@tcp(localhost:9910)/gorm?charset=utf8&parseTime=True&loc=Local"
		}
		db, err = gorm.Open(mysql.Open(dbDSN), &gorm.Config{})
	default:
		return nil, errors.New("invalaid GORM_DIALECT")
	}

	if debug := os.Getenv("DEBUG"); debug == "true" {
		db.Logger = db.Logger.LogMode(logger.Info)
	} else if debug == "false" {
		db.Logger = db.Logger.LogMode(logger.Silent)
	}

	return
}

func runMigrations() {
	var err error
	allModels := []interface{}{&User{}, &Account{}, &Pet{}, &Company{}, &Toy{}, &Language{}}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(allModels), func(i, j int) { allModels[i], allModels[j] = allModels[j], allModels[i] })

	_ = DB.Migrator().DropTable("user_friends", "user_speaks")

	if err = DB.Migrator().DropTable(allModels...); err != nil {
		log.Printf("Failed to drop table, got error %v\n", err)
		os.Exit(1)
	}

	if err = DB.AutoMigrate(allModels...); err != nil {
		log.Printf("Failed to auto migrate, but got error %v\n", err)
		os.Exit(1)
	}

	for _, m := range allModels {
		if !DB.Migrator().HasTable(m) {
			log.Printf("Failed to create table for %#v\n", m)
			os.Exit(1)
		}
	}
}
