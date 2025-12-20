package services

import (
	"context"
	"hesdastore/api-ppob/common/helper"
	"hesdastore/api-ppob/constants"
	"hesdastore/api-ppob/domain/dto"
	"hesdastore/api-ppob/repositories"
)

type ProductServiceImpl struct {
	repository repositories.IRepoRegistry
}

func NewProductServiceImpl(repo repositories.IRepoRegistry) ProductService {
	return &ProductServiceImpl{
		repository: repo,
	}
}

func (s *ProductServiceImpl) FindAllPrabayar(ctx context.Context, role int) ([]*dto.ProductResponse, error) {
	products, err := s.repository.Product().FindAllPrabayar(ctx)
	if err != nil {
		return nil, err
	}

	data := make([]*dto.ProductResponse, 0, len(products))
	for _, p := range products {
		price := helper.GetPriceProductByRole(role, *p)
		data = append(data, &dto.ProductResponse{
			ProductCode: p.ProductCode,
			ProductName: p.ProductName,
			Category: &dto.CategoryResponse{
				Name: p.Category.Name,
			},
			Brand: &dto.BrandResponse{
				Name: p.Brand.Name,
				Logo: constants.UploadBrandUrl + "/" + p.Brand.Logo,
			},
			Type:        p.Type,
			Price:       price,
			Status:      helper.GetStatusProductByRole(role, *p),
			StartCutOff: *p.StartCutOff,
			EndCutOff:   *p.EndCutOff,
			Description: p.Description,
		})
	}
	return data, nil
}

func (s *ProductServiceImpl) FindAllPascabayar(ctx context.Context, role int) ([]*dto.ProductResponse, error) {
	products, err := s.repository.Product().FindAllPascabayar(ctx)
	if err != nil {
		return nil, err
	}

	data := make([]*dto.ProductResponse, 0, len(products))
	for _, p := range products {
		data = append(data, &dto.ProductResponse{
			ProductCode: p.ProductCode,
			ProductName: p.ProductName,
			Category: &dto.CategoryResponse{
				Name: p.Category.Name,
			},
			Brand: &dto.BrandResponse{
				Name: p.Brand.Name,
				Logo: constants.UploadBrandUrl + "/" + p.Brand.Logo,
			},
			Type:        p.Type,
			Admin:       *p.Admin,
			Commission:  *p.Commission,
			Status:      helper.GetStatusProductByRole(role, *p),
			Description: p.Description,
		})
	}
	return data, nil
}

func (s *ProductServiceImpl) FindByProductCode(ctx context.Context, productCode string, role int) (*dto.ProductResponse, error) {
	data, err := s.repository.Product().FindByProductCode(ctx, productCode)
	if err != nil {
		return nil, err
	}

	var product *dto.ProductResponse
	if data.Metode == "prepaid" {
		price := helper.GetPriceProductByRole(role, *data)
		product = &dto.ProductResponse{
			ProductCode: data.ProductCode,
			ProductName: data.ProductName,
			Category: &dto.CategoryResponse{
				Name: data.Category.Name,
			},
			Brand: &dto.BrandResponse{
				Name: data.Brand.Name,
				Logo: constants.UploadBrandUrl + "/" + data.Brand.Logo,
			},
			Type:        data.Type,
			Price:       price,
			Status:      helper.GetStatusProductByRole(role, *data),
			StartCutOff: *data.StartCutOff,
			EndCutOff:   *data.EndCutOff,
			Description: data.Description,
		}
	} else {
		product = &dto.ProductResponse{
			ProductCode: data.ProductCode,
			ProductName: data.ProductName,
			Category: &dto.CategoryResponse{
				Name: data.Category.Name,
			},
			Brand: &dto.BrandResponse{
				Name: data.Brand.Name,
				Logo: constants.UploadBrandUrl + "/" + data.Brand.Logo,
			},
			Type:        data.Type,
			Admin:       *data.Admin,
			Commission:  *data.Commission,
			Status:      helper.GetStatusProductByRole(role, *data),
			Description: data.Description,
		}
	}

	return product, nil
}
