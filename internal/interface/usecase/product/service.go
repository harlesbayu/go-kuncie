package product

import "github.com/harlesbayu/kuncie/internal/shared/utils"

type Service interface {
	ListProduct(sess *utils.Session, request *ListProductRequest) (list ListProductResponse, err error)
	ScanProduct(sess *utils.Session, request *ScanProductRequest) (res *ScanProductResponse, err error)
}
