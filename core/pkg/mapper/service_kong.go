package mapper

import (
	"github.com/kong/go-kong/kong"
	"github.com/lffranca/gonga/core/entity"
)

func ServicesKongToEntity(items []*kong.Service) []entity.Service {
	var services []entity.Service
	for _, serv := range items {
		services = append(services, *ServiceKongToEntity(serv))
	}

	return services
}

func ServiceKongToEntity(item *kong.Service) *entity.Service {
	var clientCertificateID *string
	if item.ClientCertificate != nil {
		clientCertificateID = item.ClientCertificate.ID
	}

	return &entity.Service{
		ID:                  item.ID,
		Name:                item.Name,
		Protocol:            item.Protocol,
		Retries:             item.Retries,
		Port:                item.Port,
		Tags:                item.Tags,
		CACertificates:      item.CACertificates,
		Path:                item.Path,
		ConnectTimeout:      item.ConnectTimeout,
		WriteTimeout:        item.WriteTimeout,
		ReadTimeout:         item.ReadTimeout,
		TLSVerify:           item.TLSVerify,
		ClientCertificateID: clientCertificateID,
		Host:                item.Host,
		TLSVerifyDepth:      item.TLSVerifyDepth,
		CreatedAt:           item.CreatedAt,
		UpdatedAt:           item.UpdatedAt,
	}
}
