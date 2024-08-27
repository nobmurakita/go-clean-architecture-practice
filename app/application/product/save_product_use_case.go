package product

import (
	"context"
	productDomain "go-clean-architecture-practice/domain/product"
)

type SaveProductUseCase struct {
	// TODO productRepo
}

type SaveProductUseCaseInputDto struct {
	OwnerID     string
	Name        string
	Description string
	Price       int64
	Stock       int
}

type SaveProductUseCaseOutputDto struct {
	ID          string
	OwnerID     string
	Name        string
	Description string
	Price       int64
	Stock       int
}

func (uc *SaveProductUseCase) Run(
	ctx context.Context,
	input SaveProductUseCaseInputDto,
) (*SaveProductUseCaseOutputDto, error) {
	p, err := productDomain.NewProduct(input.OwnerID, input.Name, input.Description, input.Price, input.Stock)
	if err != nil {
		return nil, err
	}

	// TODO productRepo.Save

	return &SaveProductUseCaseOutputDto{
		ID:          p.ID(),
		OwnerID:     p.OwnerID(),
		Name:        p.Name(),
		Description: p.Description(),
		Price:       p.Price(),
		Stock:       p.Stock(),
	}, nil
}
