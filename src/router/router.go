package router

import (
	"net/http"

	"github.com/nicobianchetti/Meli-Quasar-NB/src/controller"
	"github.com/nicobianchetti/Meli-Quasar-NB/src/service"
)

const (
	//PATHPREFIX .
	PATHPREFIX = "/quasar"
)

type funcHandler func(w http.ResponseWriter, r *http.Request)

//IRouter .
type IRouter interface {
	SERVE(port string)
	GET(uri string, f funcHandler)
	POST(uri string, f funcHandler)
	PUT(uri string, f funcHandler)
	PATCH(uri string, f funcHandler)
}

//SetupRoutesSatellite .
func SetupRoutesSatellite(httpRouter IRouter) {
	/*
		Defino driver de base de datos
	*/
	// driver := database.Postgres
	// database.New(driver)
	// db := database.DB()

	/*
		Conexión entidad Mutant
	*/
	// rMutant := repository.NewMutantRepository(db)
	sSatellite := service.NewSatelliteService()
	cSatellite := controller.NewSatelliteController(sSatellite)

	NewSatelliteRouter(PATHPREFIX, cSatellite, httpRouter)

}
