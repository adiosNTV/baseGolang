package payload

//
//import (
//	db "TechStore/db/sqlc"
//	"TechStore/internal/pkg/validator"
//	"fmt"
//	"time"
//)
//
//type CreateOrderRequest struct {
//	Data struct {
//		Products []struct {
//			Id       string `json:"id"`
//			Quantity int    `json:"quantity"`
//		} `json:"products"`
//		Amount float64 `json:"amount"`
//		UserId string  `json:"user_id"`
//		Status int     `json:"status"`
//	} `json:"data"`
//}
//
//type GetOrderByIdResponse struct {
//	Id           string           `json:"id"`
//	Amount       int              `json:"amount"`
//	Status       int              `json:"status"`
//	CreatedAt    time.Time        `json:"createdAt"`
//	OrderDetails []db.OrderDetail `json:"orderDetails"`
//}
//
//func ValidateCreateOrderRequest(v *validator.Validator, request CreateOrderRequest) {
//	v.Check(request.Data.Amount != 0, "amount", "can not be zero !")
//	v.Check(request.Data.UserId != "", "userId", "must be provided")
//	v.Check(validator.Matches(request.Data.UserId, validator.UuidRx), "userId", "must be uuid!")
//	v.Check(request.Data.Status >= 0 && request.Data.Status <= 2, "status", "unknown")
//	v.Check(len(request.Data.Products) > 0, "products", "must be provided")
//
//	for _, product := range request.Data.Products {
//		v.Check(validator.Matches(product.Id, validator.UuidRx), fmt.Sprintf("productId:%s", product.Id), "must be uuid!")
//		v.Check(product.Quantity > 0, "product quantity", "must larger than zero!")
//	}
//}
