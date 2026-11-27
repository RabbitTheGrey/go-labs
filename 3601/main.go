package main

import (
	"fmt"
)

const PassStatus = "pass"
const FailStatus = "fail"

type HealthCheck struct {
	ServiceID int
	status    string
}

func GenerateCheck() []HealthCheck {
	var services []HealthCheck
	var status string

	for i := 0; i < 5; i++ {
		if i%2 == 0 {
			status = PassStatus
		} else {
			status = FailStatus
		}

		services = append(services, HealthCheck{
			ServiceID: i,
			status:    status,
		})
	}

	return services
}

func main() {
	fmt.Println("Тут будет выведен идентификатор")

	hc := GenerateCheck()

	for _, service := range hc {
		if service.status == PassStatus {
			fmt.Println(service.ServiceID)
		}
	}
}
