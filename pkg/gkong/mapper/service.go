package mapper

import (
	"github.com/kong/go-kong/kong"
	"github.com/lffranca/gonga/domain"
	"time"
)

func ServiceEntitiesToModels(items []*domain.Service) (values []*kong.Service) {
	for _, item := range items {
		values = append(values, ServiceEntityToModel(item))
	}

	return
}

func ServiceEntityToModel(item *domain.Service) *kong.Service {
	if item == nil {
		return nil
	}

	var createdAt *int
	if item.CreatedAt != nil {
		t := int(item.CreatedAt.Unix())
		createdAt = &t
	}

	var updatedAt *int
	if item.UpdatedAt != nil {
		t := int(item.UpdatedAt.Unix())
		updatedAt = &t
	}

	return &kong.Service{
		ClientCertificate: CertificateEntityToModel(item.ClientCertificate),
		ConnectTimeout:    item.ConnectTimeout,
		CreatedAt:         createdAt,
		Host:              item.Host,
		ID:                item.ID,
		Name:              item.Name,
		Path:              item.Path,
		Port:              item.Port,
		Protocol:          item.Protocol,
		ReadTimeout:       item.ReadTimeout,
		Retries:           item.Retries,
		UpdatedAt:         updatedAt,
		URL:               item.URL,
		WriteTimeout:      item.WriteTimeout,
		Tags:              item.Tags,
		TLSVerify:         item.TLSVerify,
		TLSVerifyDepth:    item.TLSVerifyDepth,
		CACertificates:    item.CACertificates,
	}
}

func ServiceModelsToEntities(items []*kong.Service) (values []*domain.Service) {
	for _, item := range items {
		values = append(values, ServiceModelToEntity(item))
	}

	return
}

func ServiceModelToEntity(item *kong.Service) *domain.Service {
	if item == nil {
		return nil
	}

	var createdAt *time.Time
	if item.CreatedAt != nil {
		t := time.UnixMilli(int64(*item.CreatedAt))
		createdAt = &t
	}

	var updatedAt *time.Time
	if item.UpdatedAt != nil {
		t := time.UnixMilli(int64(*item.UpdatedAt))
		updatedAt = &t
	}

	return &domain.Service{
		ClientCertificate: CertificateModelToEntity(item.ClientCertificate),
		ConnectTimeout:    item.ConnectTimeout,
		CreatedAt:         createdAt,
		Host:              item.Host,
		ID:                item.ID,
		Name:              item.Name,
		Path:              item.Path,
		Port:              item.Port,
		Protocol:          item.Protocol,
		ReadTimeout:       item.ReadTimeout,
		Retries:           item.Retries,
		UpdatedAt:         updatedAt,
		URL:               item.URL,
		WriteTimeout:      item.WriteTimeout,
		Tags:              item.Tags,
		TLSVerify:         item.TLSVerify,
		TLSVerifyDepth:    item.TLSVerifyDepth,
		CACertificates:    item.CACertificates,
	}
}
