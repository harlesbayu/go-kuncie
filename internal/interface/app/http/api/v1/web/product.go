package web

import (
	"github.com/harlesbayu/kuncie/internal/interface/usecase/product"
	"github.com/harlesbayu/kuncie/internal/shared/constants"
	"github.com/harlesbayu/kuncie/internal/shared/utils"
	"github.com/labstack/echo/v4"
	"github.com/spf13/cast"
)

type ProductHandler struct {
	productService product.Service
}

func NewSubmissionHandler(pro product.Service) *ProductHandler {
	return &ProductHandler{
		productService: pro,
	}
}

func (h *ProductHandler) ListProductHandler(ctx echo.Context) (err error) {
	var sess = ctx.Get(constants.AppSessionRequest).(*utils.Session)
	var request = new(product.ListProductRequest)
	request.Page = cast.ToInt64(ctx.QueryParam("page"))
	request.PerPage = cast.ToInt64(ctx.QueryParam("perPage"))

	if err = ctx.Bind(request); err != nil {
		return sess.ResponseInvalidRequest(ctx, err.Error())
	}

	products, err := h.productService.ListProduct(sess, request)
	if err != nil {
		return sess.ResponseInternalError(ctx, err.Error())
	}

	response := utils.CreateHttpResponse(constants.StatusOK, "Success!", products)
	return sess.ResponseOK(ctx, response)
}

func (h *ProductHandler) ScanProductHandler(ctx echo.Context) (err error) {
	var sess = ctx.Get(constants.AppSessionRequest).(*utils.Session)
	var request = new(product.ScanProductRequest)
	if err = ctx.Bind(request); err != nil {
		return sess.ResponseInvalidRequest(ctx, err.Error())
	}

	products, err := h.productService.ScanProduct(sess, request)
	if err != nil {
		return sess.ResponseInternalError(ctx, err.Error())
	}

	response := utils.CreateHttpResponse(constants.StatusOK, "Success!", products)
	return sess.ResponseOK(ctx, response)
}
