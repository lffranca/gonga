package kong

import (
	"encoding/json"
	"github.com/lffranca/gonga/pkg/kong/consumer"
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

	tagName := "apimanagement-via"

	item, errItem := gateway.ListAllConsumers(nil, &tagName)
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

	itemName := "my-username"

	item, errItem := gateway.ConsumerByNameOrID(&itemName)
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

	itemToCreate := &consumer.Form{
		Username: NewString("consumer23"),
		CustomID: NewString("consumer23"),
		Tags:     []string{"test1", "test2"},
	}

	item, errItem := gateway.CreateConsumer(itemToCreate)
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

	itemToUpdate := &consumer.Form{
		Tags: []string{"testtest", "test1"},
	}

	username := "consumer23"

	item, errItem := gateway.UpsertConsumer(&username, itemToUpdate)
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

	serviceName := "consumer23"

	if err := gateway.DeleteConsumer(&serviceName); err != nil {
		t.Error(err)
		return
	}

	log.Println("deleted item: ", serviceName)
}
