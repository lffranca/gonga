package kong

import (
	"encoding/json"
	"github.com/lffranca/gonga/pkg/kong/route"
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

	// tagName := "apimanagement-via"

	item, errItem := gateway.ListAllRoutes(nil, nil)
	if errItem != nil {
		t.Error(errItem)
		return
	}

	jsonResult, _ := json.Marshal(item)
	log.Println("item: ", string(jsonResult))
}

func Test_kong_Item(t *testing.T) {
	gateway, errGateway := NewKong("http://localhost:8001")
	if errGateway != nil {
		t.Error(errGateway)
		return
	}

	itemName := "apione-develop.manager.02"

	item, errItem := gateway.RouteByNameOrID(&itemName)
	if errItem != nil {
		t.Error(errItem)
		return
	}

	jsonResult, _ := json.Marshal(item)
	log.Println("item: ", string(jsonResult))
}

func Test_kong_Create(t *testing.T) {
	gateway, errGateway := NewKong("http://localhost:8001")
	if errGateway != nil {
		t.Error(errGateway)
		return
	}

	itemToCreate := &route.Form{
		Name: NewString("route-test"),
		Paths: []string{"/servicetest01", "/servicetest02"},
		Service: &route.ItemID{ID: NewString("427163d3-84f2-412f-bbb5-62285fd9646a")},
	}

	item, errItem := gateway.CreateRoute(itemToCreate)
	if errItem != nil {
		t.Error(errItem)
		return
	}

	jsonResult, _ := json.Marshal(item)
	log.Println("item: ", string(jsonResult))
}

func Test_kong_Update(t *testing.T) {
	gateway, errGateway := NewKong("http://localhost:8001")
	if errGateway != nil {
		t.Error(errGateway)
		return
	}

	itemToUpdate := &route.Form{
		Tags: []string{"testtest", "test1"},
	}

	username := "route-test"

	item, errItem := gateway.UpsertRoute(&username, itemToUpdate)
	if errItem != nil {
		t.Error(errItem)
		return
	}

	jsonResult, _ := json.Marshal(item)
	log.Println("item: ", string(jsonResult))
}

func Test_kong_Delete(t *testing.T) {
	gateway, errGateway := NewKong("http://localhost:8001")
	if errGateway != nil {
		t.Error(errGateway)
		return
	}

	serviceName := "route-test"

	if err := gateway.DeleteRoute(&serviceName); err != nil {
		t.Error(err)
		return
	}

	log.Println("deleted item: ", serviceName)
}
