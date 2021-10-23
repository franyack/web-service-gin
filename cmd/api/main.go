package main

import "example/web-service-gin/internal/infraestructure/delivery/webapi"

func main()  {
	deliveryStrategy := webapi.New()
	deliveryStrategy.Start()
}

