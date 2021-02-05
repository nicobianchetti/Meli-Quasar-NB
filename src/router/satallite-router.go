package router

import (
	"github.com/nicobianchetti/Meli-Quasar-NB/src/interfaces"
	"github.com/nicobianchetti/Meli-Quasar-NB/src/middleware"
)

type satelliteRouter struct {
	routerSatellite interfaces.ISatelliteController
}

//NewSatelliteRouter instanced routes for mutants
func NewSatelliteRouter(pathPrefix string, cSatellite interfaces.ISatelliteController, httpRouter IRouter) {
	rSatellite := satelliteRouter{cSatellite}
	rSatellite.RoutesSatellite(pathPrefix, httpRouter)
}

func (r *satelliteRouter) RoutesSatellite(pathPrefix string, httpRouter IRouter) {
	httpRouter.POST(pathPrefix+"/topsecret/", middleware.LogAndAuthentication(r.routerSatellite.TopSecret))
}
