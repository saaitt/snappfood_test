package app

import (
	"github.com/gin-gonic/gin"
	"github.com/saaitt/snappfood_test/db"
	od "github.com/saaitt/snappfood_test/orders/domain"
	"gorm.io/gorm"
)

type App struct {
	engine *gin.Engine
	db     *gorm.DB

	orderUc od.UseCase
	orderHr od.Handler
}

func New() *App {
	return &App{}
}

func (a *App) Initialize() {
	a.InitCore()
	a.InitModules()
	a.RegisterRoutes()
}

func (a *App) InitEngine() {
	a.engine = gin.Default()
}

func (a *App) InitDB() {
	a.db = db.New()
}

func (a *App) InitModules() {
	a.InitOrders()
}

func (a *App) InitCore() {
	a.InitEngine()
	a.InitDB()
}
