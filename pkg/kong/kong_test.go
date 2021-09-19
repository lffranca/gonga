package kong

import (
	"github.com/lffranca/gonga/pkg/kong/service"
	"log"
	"testing"
)

func Test_kong_GetInformationRoutes(t *testing.T) {
	gateway, errGateway := NewKong("http://localhost:8001")
	if errGateway != nil {
		t.Error(errGateway)
		return
	}

	informationRoutes, errInformation := gateway.InformationRoutes()
	if errInformation != nil {
		t.Error(errInformation)
		return
	}

	log.Println("gateway: ", informationRoutes)
}

func Test_kong_GetHealthRoutes(t *testing.T) {
	gateway, errGateway := NewKong("http://localhost:8001")
	if errGateway != nil {
		t.Error(errGateway)
		return
	}

	item, errItem := gateway.HealthRoutes()
	if errItem != nil {
		t.Error(errItem)
		return
	}

	log.Println("item: ", item)
}

func Test_kong_List(t *testing.T) {
	gateway, errGateway := NewKong("http://localhost:8001")
	if errGateway != nil {
		t.Error(errGateway)
		return
	}

	tagName := "managed-by-ingress-controller"

	item, errItem := gateway.ListAllServices(nil, &tagName)
	if errItem != nil {
		t.Error(errItem)
		return
	}

	log.Println("item: ", item)
}

func Test_kong_Item(t *testing.T) {
	gateway, errGateway := NewKong("http://localhost:8001")
	if errGateway != nil {
		t.Error(errGateway)
		return
	}

	itemName := "20bb5f94-2f6d-48b6-ae30-96ca6182ea97"

	item, errItem := gateway.ServiceByNameOrID(&itemName)
	if errItem != nil {
		t.Error(errItem)
		return
	}

	log.Println("item: ", item)
}

func Test_kong_Create(t *testing.T) {
	gateway, errGateway := NewKong("http://localhost:8001")
	if errGateway != nil {
		t.Error(errGateway)
		return
	}

	itemToCreate  := &service.Form{
		Name: NewString("service-create-from-gonga"),
		Protocol: NewString("http"),
		Host: NewString("localhost"),
		Port: NewInt(80),
	}

	item, errItem := gateway.CreateService(itemToCreate)
	if errItem != nil {
		t.Error(errItem)
		return
	}

	log.Println("item: ", item)
}

func Test_kong_Update(t *testing.T) {
	gateway, errGateway := NewKong("http://localhost:8001")
	if errGateway != nil {
		t.Error(errGateway)
		return
	}

	itemToCreate  := &service.Form{
		Protocol: NewString("http"),
	}

	serviceName := "service-create-from-gonga"

	item, errItem := gateway.UpsertService(&serviceName, itemToCreate)
	if errItem != nil {
		t.Error(errItem)
		return
	}

	log.Println("item: ", item)
}

func Test_kong_Delete(t *testing.T) {
	gateway, errGateway := NewKong("http://localhost:8001")
	if errGateway != nil {
		t.Error(errGateway)
		return
	}

	serviceName := "service-create-from-gonga"

	if err := gateway.DeleteService(&serviceName); err != nil {
		t.Error(err)
		return
	}

	log.Println("deleted item: ", serviceName)
}


