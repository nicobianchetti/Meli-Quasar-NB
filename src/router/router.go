package router

import (
	"net/http"
	"os"

	"github.com/nicobianchetti/Meli-Quasar-NB/src/cache"
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

	// portRedis := "6379"
	// portRedis := "11105"

	// Credenciales prueba contra deploy de redis en Heroku (Redis Cloud)
	// pass := "nztCPqkAMdHEhwDJFgR2h8Ggdy2sWjuO"

	// addr := "redis-13325.c89.us-east-1-3.ec2.cloud.redislabs.com:13325"

	addr := os.Getenv("REDISHOST")

	pass := os.Getenv("REDISPASS")

	if addr == "" {
		addr = "localhost:6379"
	}

	if pass == "" {
		pass = ""
	}

	cacheSatellite := cache.NewRedisCache(addr, 0, 500, pass)
	// cacheSatellite := cache.NewRedisCache(ip+":"+portRedis, 1, 500, pass)
	// cacheSatellite := cache.NewRedisCache("localhost:6379", 1, 500, "")
	sSatellite := service.NewSatelliteService(cacheSatellite)
	cSatellite := controller.NewSatelliteController(sSatellite)

	NewSatelliteRouter(PATHPREFIX, cSatellite, httpRouter)

}

// //Debug IP dentro de docker
// func externalIP() (string, error) {
// 	ifaces, err := net.Interfaces()
// 	if err != nil {
// 		return "", err
// 	}
// 	for _, iface := range ifaces {
// 		if iface.Flags&net.FlagUp == 0 {
// 			continue // interface down
// 		}
// 		if iface.Flags&net.FlagLoopback != 0 {
// 			continue // loopback interface
// 		}
// 		addrs, err := iface.Addrs()
// 		if err != nil {
// 			return "", err
// 		}
// 		for _, addr := range addrs {
// 			var ip net.IP
// 			switch v := addr.(type) {
// 			case *net.IPNet:
// 				ip = v.IP
// 			case *net.IPAddr:
// 				ip = v.IP
// 			}
// 			if ip == nil || ip.IsLoopback() {
// 				continue
// 			}
// 			ip = ip.To4()
// 			if ip == nil {
// 				continue // not an ipv4 address
// 			}
// 			return ip.String(), nil
// 		}
// 	}
// 	return "", errors.New("are you connected to the network?")
// }
