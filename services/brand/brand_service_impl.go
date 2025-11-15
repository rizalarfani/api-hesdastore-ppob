package services

import (
	"context"
	"hesdastore/api-ppob/constants"
	"hesdastore/api-ppob/domain/dto"
	"hesdastore/api-ppob/repositories"
)

type BrandServiceImpl struct {
	repository repositories.IRepoRegistry
}

func NewBrandServiceImpl(repo repositories.IRepoRegistry) BrandService {
	return &BrandServiceImpl{
		repository: repo,
	}
}

func (s *BrandServiceImpl) FindAll(ctx context.Context) ([]*dto.BrandResponse, error) {
	brand, err := s.repository.Brand().FindAll(ctx)
	if err != nil {
		return nil, err
	}

	data := make([]*dto.BrandResponse, 0, len(brand))
	for _, b := range brand {
		data = append(data, &dto.BrandResponse{
			ID:   int(b.ID),
			Name: b.Name,
			Logo: constants.UploadBrandUrl + "/" + b.Logo,
		})
	}

	return data, nil
}
