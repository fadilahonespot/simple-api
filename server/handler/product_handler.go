package handler

import (
	"net/http"

	"github.com/fadilahonespot/library/errors"
	"github.com/fadilahonespot/library/response"
	"github.com/fadilahonespot/simple-api/usecase"
	"github.com/fadilahonespot/simple-api/usecase/dto"
	"github.com/fadilahonespot/simple-api/utils/logger"
	"github.com/fadilahonespot/simple-api/utils/paginate"
	"github.com/labstack/echo/v4"
)

type ProductHandler struct {
	productUsecase usecase.ProductUsecase
}

func NewProductHandler(productUsecase usecase.ProductUsecase) ProductHandler {
	return ProductHandler{productUsecase: productUsecase}
}

func (h *ProductHandler) AddProduct(c echo.Context) (err error) {
	ctx := c.Request().Context()

	var req dto.ProductRequest
	err = c.Bind(&req)
	if err != nil {
		logger.Error(ctx, "error binding", err.Error())
		err = errors.SetError(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}

	err = c.Validate(req)
	if err != nil {
		logger.Error(ctx, "error validating", err.Error())
		err = errors.SetError(http.StatusBadRequest, err.Error())
		return
	}

	logger.Info(ctx, "[Request]", req)

	err = h.productUsecase.CreateProduct(ctx, req)
	if err != nil {
		return err
	}

	resp := response.ResponseSuccess(nil)
	return c.JSON(http.StatusOK, resp)
}

func (h *ProductHandler) GetListProduct(c echo.Context) (err error) {
	ctx := c.Request().Context()
	params := paginate.GetParams(c)
	
	data, count, err := h.productUsecase.GetListProduct(ctx, params)
	if err != nil {
		return err
	}

	resp := response.HandleSuccessWithPagination(float64(count), params.Limit, params.Page, data)
	return c.JSON(http.StatusOK, resp)
}

func (h *ProductHandler) GetProductDetail(c echo.Context) (err error) {
	ctx := c.Request().Context()
	productId := c.Param("productId")
	data, err := h.productUsecase.GetDetailProduct(ctx, productId)
	if err != nil {
		return
	}

	resp := response.ResponseSuccess(data)
	return c.JSON(http.StatusOK, resp)
}

func (h *ProductHandler) UpdateProduct(c echo.Context) (err error) {
	ctx := c.Request().Context()
	productId := c.Param("productId")

	var req dto.ProductRequest
	err = c.Bind(&req)
	if err != nil {
		logger.Error(ctx, "error binding", err.Error())
		err = errors.SetError(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}

	err = c.Validate(req)
	if err != nil {
		logger.Error(ctx, "error validating", err.Error())
		err = errors.SetError(http.StatusBadRequest, err.Error())
		return
	}

	logger.Info(ctx, "[Request]", req)

	err = h.productUsecase.UpdateProduct(ctx, productId, req)
	if err != nil {
		return err
	}

	resp := response.ResponseSuccess(nil)
	return c.JSON(http.StatusOK, resp)
}

func (h *ProductHandler) DeleteProduct(c echo.Context) (err error) {
	ctx := c.Request().Context()
	productId := c.Param("productId")
	err = h.productUsecase.DeleteProduct(ctx, productId)
	if err != nil {
		return
	}

	resp := response.ResponseSuccess(nil)
	return c.JSON(http.StatusOK, resp)
}
