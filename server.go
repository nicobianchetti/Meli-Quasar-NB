package main

import (
	"os"

	"github.com/nicobianchetti/Meli-Quasar-NB/src/router"
)

const (
	_portDefualt = "5000"
	_ipDefault   = ""
)

//Ayuda memoria docker local
// docker-compose --compatibility up --build
//docker run --name redis-test -p 6379:6379 -d redis
// docker exec -it redis-test sh
// redis-cli
// SELECt <db>

func main() {

	// ip, err := externalIP()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(ip)

	httpRouter := router.NewMuxRouter()
	router.SetupRoutesSatellite(httpRouter)

	serverPort := os.Getenv("PORT")

	if serverPort == "" {
		serverPort = _portDefualt
	}

	serverIP := os.Getenv("IP")

	if serverIP == "" {
		serverIP = _ipDefault
	}

	addr := serverIP + ":" + serverPort

	httpRouter.SERVE(addr)

}

// //Debug Ip dentro de container trabajando con docker-compose
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
