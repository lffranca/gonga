package mapper

import (
	"github.com/kong/go-kong/kong"
	"github.com/lffranca/gonga/domain"
	"time"
)

func CertificateEntityToModel(item *domain.Certificate) *kong.Certificate {
	if item == nil {
		return nil
	}

	var createdAt *int64
	if item.CreatedAt != nil {
		t := item.CreatedAt.Unix()
		createdAt = &t
	}

	return &kong.Certificate{
		ID:        item.ID,
		Cert:      item.Cert,
		Key:       item.Key,
		CreatedAt: createdAt,
		SNIs:      item.SNIs,
		Tags:      item.Tags,
	}
}

func CertificateModelToEntity(item *kong.Certificate) *domain.Certificate {
	if item == nil {
		return nil
	}

	var createdAt *time.Time
	if item.CreatedAt != nil {
		t := time.UnixMilli(*item.CreatedAt)
		createdAt = &t
	}

	return &domain.Certificate{
		ID:        item.ID,
		Cert:      item.Cert,
		Key:       item.Key,
		CreatedAt: createdAt,
		SNIs:      item.SNIs,
		Tags:      item.Tags,
	}
}
