package app

import (
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/gin-gonic/gin"
	ad "github.com/saaitt/snappfood_test/agents/domain"
	"github.com/saaitt/snappfood_test/db"
	od "github.com/saaitt/snappfood_test/orders/domain"
	"github.com/saaitt/snappfood_test/services/kafka_broker"
	td "github.com/saaitt/snappfood_test/trips/domain"
	"gorm.io/gorm"
)

type App struct {
	engine   *gin.Engine
	db       *gorm.DB
	consumer *kafka_broker.MessageConsumer
	producer *kafka.Producer
	config   ServiceConfig

	orderUc od.UseCase
	orderHr od.Handler

	tripUc td.UseCase
	tripHr td.Handler

	agentUc ad.UseCase
	agentHr ad.Handler
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
	a.InitAgents()
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
	a.db = db.New(a.config.DB.Host,
		a.config.DB.User,
		a.config.DB.Password,
		a.config.DB.DbName,
		a.config.DB.Port,
	)
}
