package app

import (
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/gin-gonic/gin"
	"github.com/saaitt/snappfood_test/db"
	od "github.com/saaitt/snappfood_test/orders/domain"
	"github.com/saaitt/snappfood_test/services/kafka_broker"
	td "github.com/saaitt/snappfood_test/trips/domain"
	"gorm.io/gorm"
)

type App struct {
	engine   *gin.Engine
	db       *gorm.DB
	consumer *kafka.Consumer
	producer *kafka.Producer
	config   ServiceConfig

	orderUc od.UseCase
	orderHr od.Handler

	tripUc td.UseCase
	tripHr td.Handler
}

func New(cfg ServiceConfig) *App {
	return &App{
		config: cfg,
	}
}

func (a *App) Initialize() {
	a.InitCore()
	a.InitModules()
	a.RegisterRoutes()
}

func (a *App) InitModules() {
	a.InitTrips()
	a.InitOrders()
}

func (a *App) InitCore() {
	a.InitEngine()
	a.InitDB()
	a.InitBroker()
}

func (a *App) InitEngine() {
	a.engine = gin.Default()
}

func (a *App) InitDB() {
	a.db = db.New()
}

func (a *App) InitBroker() {
	a.consumer = kafka_broker.Consumer(
		a.config.Broker.Server,
		a.config.Broker.GroupId,
		a.config.Broker.OffsetReset,
		a.config.Broker.Topics)
	a.producer = kafka_broker.Producer(a.config.Broker.Server)
}
