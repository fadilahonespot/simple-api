package usecase

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/fadilahonespot/simple-api/entity"
	"github.com/fadilahonespot/simple-api/repository/mocks"
	"github.com/fadilahonespot/simple-api/usecase/dto"
	"github.com/fadilahonespot/simple-api/utils/logger"
	"github.com/fadilahonespot/simple-api/utils/paginate"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

func Test_defaultProductUsecase_CreateProduct(t *testing.T) {
	ctx := context.TODO()
	logger.NewLogger()

	type args struct {
		ctx context.Context
		req dto.ProductRequest
	}
	tests := []struct {
		name             string
		args             args
		getProductResp   *entity.Product
		getProductErr    error
		createProductErr error
		wantErr          bool
	}{
		{
			name: "produt already exists",
			args: args{
				ctx: ctx,
				req: dto.ProductRequest{
					Title:       "Mie indomi Rasa ayam Bawang",
					Description: "Taburan ayam gurih nikmat di setiap kemasan",
				},
			},
			getProductResp: &entity.Product{
				Title: "Mie indomi Rasa ayam Bawang",
			},
			wantErr: true,
		},
		{
			name: "create product error",
			args: args{
				ctx: ctx,
				req: dto.ProductRequest{
					Title:       "Mie indomi Rasa ayam Bawang",
					Description: "Taburan ayam gurih nikmat di setiap kemasan",
				},
			},
			getProductResp: &entity.Product{
				Title: "",
			},
			createProductErr: errors.New("create product error"),
			wantErr:          true,
		},
		{
			name: "create product success",
			args: args{
				ctx: ctx,
				req: dto.ProductRequest{
					Title:       "Mie indomi Rasa ayam Bawang",
					Description: "Taburan ayam gurih nikmat di setiap kemasan",
				},
			},
			getProductResp: &entity.Product{
				Title: "",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			productRepo := new(mocks.ProductRepository)
			productRepo.On("GetProductByTitle", mock.Anything, mock.Anything).Return(tt.getProductResp, tt.getProductErr).Once()
			productRepo.On("CreateProduct", mock.Anything, mock.Anything).Return(tt.createProductErr).Once()

			svc := NewProductRepository(productRepo)
			if err := svc.CreateProduct(tt.args.ctx, tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("defaultProductUsecase.CreateProduct() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_defaultProductUsecase_GetListProduct(t *testing.T) {
	ctx := context.TODO()
	logger.NewLogger()

	uid, _ := uuid.Parse("a1b91cb9-c4a5-408f-ad28-5f32e197d954")

	type args struct {
		ctx   context.Context
		param paginate.Pagination
	}
	tests := []struct {
		name        string
		args        args
		listProduct []entity.Product
		listCount   int64
		listErr     error
		wantResp    []dto.ProductListResponse
		wantCount   int64
		wantErr     bool
	}{
		{
			name: "failed to get list product",
			args: args{
				ctx: ctx,
				param: paginate.Pagination{
					Page:  1,
					Limit: 10,
				},
			},
			listErr: errors.New("failed to get list product"),
			wantErr: true,
		},
		{
			name: "success get list product",
			args: args{
				ctx: ctx,
				param: paginate.Pagination{
					Page:  1,
					Limit: 10,
				},
			},
			listProduct: []entity.Product{
				{
					ID:          uid,
					Title:       "Mie indomi Rasa ayam Bawang",
					Description: "Taburan ayam gurih nikmat di setiap kemasan",
					Rating:      8.1,
					Image:       "http://google.com/image.jpg",
				},
			},
			listCount: 1,
			wantCount: 1,
			wantResp: []dto.ProductListResponse{
				{
					ID:          uid,
					Title:       "Mie indomi Rasa ayam Bawang",
					Description: "Taburan ayam gurih nikmat di setiap kemasan",
					Rating:      8.1,
					Image:       "http://google.com/image.jpg",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			productRepo := new(mocks.ProductRepository)
			productRepo.On("GetListProduct", mock.Anything, mock.Anything).Return(tt.listProduct, tt.listCount, tt.listErr).Once()

			svc := NewProductRepository(productRepo)
			gotResp, gotCount, err := svc.GetListProduct(tt.args.ctx, tt.args.param)
			if (err != nil) != tt.wantErr {
				t.Errorf("defaultProductUsecase.GetListProduct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("defaultProductUsecase.GetListProduct() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
			if gotCount != tt.wantCount {
				t.Errorf("defaultProductUsecase.GetListProduct() gotCount = %v, want %v", gotCount, tt.wantCount)
			}
		})
	}
}

func Test_defaultProductUsecase_GetDetailProduct(t *testing.T) {
	ctx := context.TODO()
	logger.NewLogger()
	uidStr := "a1b91cb9-c4a5-408f-ad28-5f32e197d954"
	uid, _ := uuid.Parse(uidStr)

	type args struct {
		ctx       context.Context
		productId string
	}
	tests := []struct {
		name           string
		args           args
		getProductResp *entity.Product
		getProductErr  error
		wantResp       dto.DetailProductResponse
		wantErr        bool
	}{
		{
			name: "product not found",
			args: args{
				ctx:       ctx,
				productId: uidStr,
			},
			getProductErr: errors.New("product not found"),
			wantErr:       true,
		},
		{
			name: "get product success",
			args: args{
				ctx:       ctx,
				productId: uidStr,
			},
			getProductResp: &entity.Product{
				ID:          uid,
				Title:       "Mie indomi Rasa ayam Bawang",
				Description: "Taburan ayam gurih nikmat di setiap kemasan",
				Rating:      8.1,
				Image:       "http://google.com/image.jpg",
			},
			wantResp: dto.DetailProductResponse{
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
			productRepo := new(mocks.ProductRepository)
			productRepo.On("GetProductById", mock.Anything, mock.Anything).Return(tt.getProductResp, tt.getProductErr).Once()

			svc := NewProductRepository(productRepo)
			gotResp, err := svc.GetDetailProduct(tt.args.ctx, tt.args.productId)
			if (err != nil) != tt.wantErr {
				t.Errorf("defaultProductUsecase.GetDetailProduct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("defaultProductUsecase.GetDetailProduct() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}

func Test_defaultProductUsecase_UpdateProduct(t *testing.T) {
	ctx := context.TODO()
	logger.NewLogger()
	uidStr := "a1b91cb9-c4a5-408f-ad28-5f32e197d954"
	uid, _ := uuid.Parse(uidStr)

	type args struct {
		ctx       context.Context
		productId string
		req       dto.ProductRequest
	}
	tests := []struct {
		name                string
		args                args
		getProductResp      *entity.Product
		getProductErr       error
		getProductTitleResp *entity.Product
		getProductTitleErr  error
		updateProductErr    error
		wantErr             bool
	}{
		{
			name: "product not found",
			args: args{
				ctx:       ctx,
				productId: uidStr,
				req: dto.ProductRequest{
					Title:       "Mie indomi Rasa ayam Bawang",
					Description: "Taburan ayam gurih nikmat di setiap kemasan",
					Rating:      8.1,
					Image:       "http://google.com/image.jpg",
				},
			},
			getProductErr: errors.New("product not found"),
			wantErr:       true,
		},
		{
			name: "product title already exist",
			args: args{
				ctx:       ctx,
				productId: uidStr,
				req: dto.ProductRequest{
					Title: "Mie indomi Rasa Soto",
				},
			},
			getProductResp: &entity.Product{
				ID:    uid,
				Title: "Mie indomi Rasa ayam Bawang",
			},
			getProductTitleResp: &entity.Product{
				Title: "Mie indomi Rasa Soto",
			},
			wantErr: true,
		},
		{
			name: "error update product",
			args: args{
				ctx:       ctx,
				productId: uidStr,
				req: dto.ProductRequest{
					Title: "Mie indomi Rasa ayam Bawang",
				},
			},
			getProductResp: &entity.Product{
				ID:    uid,
				Title: "Mie indomi Rasa ayam Bawang",
			},
			updateProductErr: errors.New("error update product"),
			wantErr:          true,
		},
		{
			name: "success update product",
			args: args{
				ctx:       ctx,
				productId: uidStr,
				req: dto.ProductRequest{
					Title: "Mie indomi Rasa ayam Bawang",
				},
			},
			getProductResp: &entity.Product{
				ID:    uid,
				Title: "Mie indomi Rasa ayam Bawang",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			productRepo := new(mocks.ProductRepository)
			productRepo.On("GetProductById", mock.Anything, mock.Anything).Return(tt.getProductResp, tt.getProductErr).Once()
			productRepo.On("GetProductByTitle", mock.Anything, mock.Anything).Return(tt.getProductTitleResp, tt.getProductTitleErr).Once()
			productRepo.On("UpdateProduct", mock.Anything, mock.Anything).Return(tt.updateProductErr).Once()

			svc := NewProductRepository(productRepo)
			if err := svc.UpdateProduct(tt.args.ctx, tt.args.productId, tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("defaultProductUsecase.UpdateProduct() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_defaultProductUsecase_DeleteProduct(t *testing.T) {
	ctx := context.TODO()
	logger.NewLogger()
	uidStr := "a1b91cb9-c4a5-408f-ad28-5f32e197d954"
	uid, _ := uuid.Parse(uidStr)

	type args struct {
		ctx       context.Context
		productId string
	}
	tests := []struct {
		name             string
		args             args
		getProductResp   *entity.Product
		getProductErr    error
		deleteProductErr error
		wantErr          bool
	}{
		{
			name: "product not found",
			args: args{
				ctx:       ctx,
				productId: uidStr,
			},
			getProductErr: errors.New("product not found"),
			wantErr:       true,
		},
		{
			name: "error delete product",
			args: args{
				ctx:       ctx,
				productId: uidStr,
			},
			getProductResp: &entity.Product{
				ID:          uid,
				Title:       "Mie indomi Rasa ayam Bawang",
				Description: "Taburan ayam gurih nikmat di setiap kemasan",
				Rating:      8.1,
				Image:       "http://google.com/image.jpg",
			},
			deleteProductErr: errors.New("failed delete product"),
			wantErr:          true,
		},
		{
			name: "success delete product",
			args: args{
				ctx:       ctx,
				productId: uidStr,
			},
			getProductResp: &entity.Product{
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
			productRepo := new(mocks.ProductRepository)
			productRepo.On("GetProductById", mock.Anything, mock.Anything).Return(tt.getProductResp, tt.getProductErr).Once()
			productRepo.On("DeleteProduct", mock.Anything, mock.Anything).Return(tt.deleteProductErr).Once()

			svc := NewProductRepository(productRepo)
			if err := svc.DeleteProduct(tt.args.ctx, tt.args.productId); (err != nil) != tt.wantErr {
				t.Errorf("defaultProductUsecase.DeleteProduct() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
