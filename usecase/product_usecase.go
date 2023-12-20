package usecase

import (
	"context"
	"net/http"
	"strings"

	"github.com/fadilahonespot/library/errors"
	"github.com/fadilahonespot/simple-api/entity"
	"github.com/fadilahonespot/simple-api/repository"
	"github.com/fadilahonespot/simple-api/usecase/dto"
	"github.com/fadilahonespot/simple-api/utils/logger"
	"github.com/fadilahonespot/simple-api/utils/paginate"
)

type ProductUsecase interface {
	CreateProduct(ctx context.Context, req dto.ProductRequest) (err error)
	GetListProduct(ctx context.Context, param paginate.Pagination) (resp []dto.ProductListResponse, count int64, err error)
	GetDetailProduct(ctx context.Context, productId string) (resp dto.DetailProductResponse, err error)
	UpdateProduct(ctx context.Context, productId string, req dto.ProductRequest) (err error)
	DeleteProduct(ctx context.Context, productId string) (err error)
}

type defaultProductUsecase struct {
	productRepo repository.ProductRepository
}

func NewProductRepository(productRepo repository.ProductRepository) ProductUsecase {
	return &defaultProductUsecase{productRepo: productRepo}
}

func (s *defaultProductUsecase) CreateProduct(ctx context.Context, req dto.ProductRequest) (err error) {
	productData, _ := s.productRepo.GetProductByTitle(ctx, req.Title)
	if productData.Title != "" {
		logger.Error(ctx, "product is already exist")
		err = errors.SetError(http.StatusBadRequest, "Product is already exist")
		return
	}
	reqProduct := entity.Product{
		Title:       req.Title,
		Description: req.Description,
		Rating:      req.Rating,
		Image:       req.Image,
	}
	err = s.productRepo.CreateProduct(ctx, &reqProduct)
	if err != nil {
		logger.Error(ctx, "error creating product", err.Error())
		err = errors.SetError(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	return
}

func (s *defaultProductUsecase) GetListProduct(ctx context.Context, param paginate.Pagination) (resp []dto.ProductListResponse, count int64, err error) {
	data, count, err := s.productRepo.GetListProduct(ctx, param)
	if err != nil {
		logger.Error(ctx, "error getting product list", err.Error())
		err = errors.SetError(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	for i := 0; i < len(data); i++ {
		resp = append(resp, dto.ProductListResponse{
			ID:          data[i].ID,
			Title:       data[i].Title,
			Description: data[i].Description,
			Rating:      data[i].Rating,
			Image:       data[i].Image,
		})
	}

	return
}

func (s *defaultProductUsecase) GetDetailProduct(ctx context.Context, productId string) (resp dto.DetailProductResponse, err error) {
	data, err := s.productRepo.GetProductById(ctx, productId)
	if err != nil {
		logger.Error(ctx, "error getting product", err.Error())
		err = errors.SetError(http.StatusNotFound, http.StatusText(http.StatusNotFound))
		return
	}

	resp = dto.DetailProductResponse{
		ID:          data.ID,
		Title:       data.Title,
		Description: data.Description,
		Rating:      data.Rating,
		Image:       data.Image,
		CreatedAt:   data.CreatedAt,
		UpdatedAt:   data.UpdatedAt,
		DeletedAt:   data.DeletedAt,
	}

	return
}

func (s *defaultProductUsecase) UpdateProduct(ctx context.Context, productId string, req dto.ProductRequest) (err error) {
	productData, err := s.productRepo.GetProductById(ctx, productId)
	if err != nil {
		logger.Error(ctx, "failed to get product: ", err.Error())
		err = errors.SetError(http.StatusNotFound, http.StatusText(http.StatusNotFound))
		return
	}

	if !strings.EqualFold(productData.Title, req.Title) {
		productData, _ := s.productRepo.GetProductByTitle(ctx, req.Title)
		if productData.Title != "" {
			logger.Error(ctx, "product title is already exist")
			err = errors.SetError(http.StatusBadRequest, "product is already exist")
			return
		}
	}

	productData.Title = req.Title
	productData.Description = req.Description
	productData.Rating = req.Rating
	productData.Image = req.Image
	err = s.productRepo.UpdateProduct(ctx, productData)
	if err != nil {
		logger.Error(ctx, "failed to update product", err.Error())
		err = errors.SetError(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	return
}

func (s *defaultProductUsecase) DeleteProduct(ctx context.Context, productId string) (err error) {
	_, err = s.productRepo.GetProductById(ctx, productId)
	if err != nil {
		logger.Error(ctx, "failed to get product: ", err.Error())
		err = errors.SetError(http.StatusNotFound, http.StatusText(http.StatusNotFound))
		return
	}

	err = s.productRepo.DeleteProduct(ctx, productId)
	if err != nil {
		logger.Error(ctx, "failed to delete product: ", err.Error())
		err = errors.SetError(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	return
}
