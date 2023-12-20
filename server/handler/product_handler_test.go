package handler

import (
	"errors"
	"net/http"
	"testing"

	"github.com/fadilahonespot/simple-api/usecase/dto"
	"github.com/fadilahonespot/simple-api/usecase/mocks"
	"github.com/fadilahonespot/simple-api/utils/logger"
	mockUtils "github.com/fadilahonespot/simple-api/utils/mocks"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

func TestProductHandler_AddProduct(t *testing.T) {
	logger.NewLogger()

	tests := []struct {
		name             string
		createProductErr error
		bodyRequest      interface{}
		wantErr          bool
	}{
		{
			name:        "error binding data",
			bodyRequest: map[string]string{"rating": "1"},
			wantErr:     true,
		},
		{
			name: "error validate data: title is empty",
			bodyRequest: dto.ProductRequest{
				Title:       "",
				Description: "Taburan ayam gurih nikmat di setiap kemasan",
			},
			wantErr: true,
		},
		{
			name: "create product failed",
			bodyRequest: dto.ProductRequest{
				Title:       "Mie indomi Rasa ayam Bawang",
				Description: "Taburan ayam gurih nikmat di setiap kemasan",
			},
			createProductErr: errors.New("create product error"),
			wantErr:          true,
		},
		{
			name: "create product success",
			bodyRequest: dto.ProductRequest{
				Title:       "Mie indomi Rasa ayam Bawang",
				Description: "Taburan ayam gurih nikmat di setiap kemasan",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			productUsecase := new(mocks.ProductUsecase)
			productUsecase.On("CreateProduct", mock.Anything, mock.Anything).Return(tt.createProductErr).Once()

			ctx, _ := mockUtils.MockEcho(http.MethodPost, "/products", nil, tt.bodyRequest)
			svc := NewProductHandler(productUsecase)
			if err := svc.AddProduct(ctx); (err != nil) != tt.wantErr {
				t.Errorf("ProductHandler.AddProduct() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestProductHandler_GetListProduct(t *testing.T) {
	logger.NewLogger()
	uidStr := "a1b91cb9-c4a5-408f-ad28-5f32e197d954"
	uid, _ := uuid.Parse(uidStr)

	tests := []struct {
		name      string
		listResp  []dto.ProductListResponse
		listCount int64
		listErr   error
		wantErr   bool
	}{
		{
			name:    "error get list product",
			listErr: errors.New("error get list product"),
			wantErr: true,
		},
		{
			name: "succes get list product",
			listResp: []dto.ProductListResponse{
				{
					ID:          uid,
					Title:       "Mie indomi Rasa ayam Bawang",
					Description: "Taburan ayam gurih nikmat di setiap kemasan",
					Rating:      8.1,
					Image:       "http://google.com/image.jpg",
				},
			},
			listCount: 1,
			wantErr:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			productUsecase := new(mocks.ProductUsecase)
			productUsecase.On("GetListProduct", mock.Anything, mock.Anything).Return(tt.listResp, tt.listCount, tt.listErr).Once()

			ctx, _ := mockUtils.MockEcho(http.MethodGet, "/products", nil, nil)
			svc := NewProductHandler(productUsecase)

			if err := svc.GetListProduct(ctx); (err != nil) != tt.wantErr {
				t.Errorf("ProductHandler.GetListProduct() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestProductHandler_GetProductDetail(t *testing.T) {
	logger.NewLogger()
	uidStr := "a1b91cb9-c4a5-408f-ad28-5f32e197d954"
	uid, _ := uuid.Parse(uidStr)

	tests := []struct {
		name              string
		productDetailResp dto.DetailProductResponse
		productDetailErr  error
		wantErr           bool
	}{
		{
			name:             "error get product detail",
			productDetailErr: errors.New("error get product detail"),
			wantErr:          true,
		},
		{
			name: "success get product detail",
			productDetailResp: dto.DetailProductResponse{
				ID:          uid,
				Title:       "Mie indomi Rasa ayam Bawang",
				Description: "Taburan ayam gurih nikmat di setiap kemasan",
				Rating:      8.1,
				Image:       "http://google.com/image.jpg",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			productUsecase := new(mocks.ProductUsecase)
			productUsecase.On("GetDetailProduct", mock.Anything, mock.Anything).Return(tt.productDetailResp, tt.productDetailErr).Once()

			ctx, _ := mockUtils.MockEcho(http.MethodGet, "/products", nil, nil)
			svc := NewProductHandler(productUsecase)

			if err := svc.GetProductDetail(ctx); (err != nil) != tt.wantErr {
				t.Errorf("ProductHandler.GetProductDetail() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestProductHandler_UpdateProduct(t *testing.T) {
	logger.NewLogger()
	uidStr := "a1b91cb9-c4a5-408f-ad28-5f32e197d954"

	tests := []struct {
		name        string
		bodyRequest interface{}
		updateErr   error
		wantErr     bool
	}{
		{
			name:        "error binding data",
			bodyRequest: map[string]string{"rating": "1"},
			wantErr:     true,
		},
		{
			name: "error validate data: title is empty",
			bodyRequest: dto.ProductRequest{
				Title:       "",
				Description: "Taburan ayam gurih nikmat di setiap kemasan",
			},
			wantErr: true,
		},
		{
			name: "update product failed",
			bodyRequest: dto.ProductRequest{
				Title:       "Mie indomi Rasa ayam Bawang",
				Description: "Taburan ayam gurih nikmat di setiap kemasan",
			},
			updateErr: errors.New("update product failed"),
			wantErr:   true,
		},
		{
			name: "update product success",
			bodyRequest: dto.ProductRequest{
				Title:       "Mie indomi Rasa ayam Bawang",
				Description: "Taburan ayam gurih nikmat di setiap kemasan",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			productUsecase := new(mocks.ProductUsecase)
			productUsecase.On("UpdateProduct", mock.Anything, mock.Anything, mock.Anything).Return(tt.updateErr).Once()

			ctx, _ := mockUtils.MockEcho(http.MethodPut, "/products/"+uidStr, nil, tt.bodyRequest)
			svc := NewProductHandler(productUsecase)

			if err := svc.UpdateProduct(ctx); (err != nil) != tt.wantErr {
				t.Errorf("ProductHandler.UpdateProduct() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestProductHandler_DeleteProduct(t *testing.T) {
	logger.NewLogger()
	uidStr := "a1b91cb9-c4a5-408f-ad28-5f32e197d954"
	tests := []struct {
		name      string
		deleteErr error
		wantErr   bool
	}{
		{
			name:      "error deleting product",
			deleteErr: errors.New("error deleting product"),
			wantErr:   true,
		},
		{
			name:    "success deleting product",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			productUsecase := new(mocks.ProductUsecase)
			productUsecase.On("DeleteProduct", mock.Anything, mock.Anything).Return(tt.deleteErr).Once()

			ctx, _ := mockUtils.MockEcho(http.MethodDelete, "/products/"+uidStr, nil, nil)
			svc := NewProductHandler(productUsecase)

			if err := svc.DeleteProduct(ctx); (err != nil) != tt.wantErr {
				t.Errorf("ProductHandler.DeleteProduct() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
