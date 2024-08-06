package main

//
//import (
//	db "TechStore/db/sqlc"
//	data "TechStore/internal/dto/data/paginated_result"
//	"TechStore/internal/dto/payload"
//	"math"
//	"net/http"
//	"strconv"
//)
//
//func (app *application) getProductsPaginatedHandler(w http.ResponseWriter, r *http.Request) {
//	pageNumber, _ := strconv.Atoi(app.readStrings(r.URL.Query(), "pageNumber", "1"))
//	pageSize, _ := strconv.Atoi(app.readStrings(r.URL.Query(), "pageSize", "10"))
//	minPrice, _ := strconv.Atoi(app.readStrings(r.URL.Query(), "minPrice", "0"))
//	maxPrice, _ := strconv.Atoi(app.readStrings(r.URL.Query(), "maxPrice", "0"))
//	searchTerm := app.readStrings(r.URL.Query(), "searchTerm", "")
//	label := app.readStrings(r.URL.Query(), "label", "")
//	sortBy := app.readStrings(r.URL.Query(), "sortBy", "price")
//	sortOrder := app.readStrings(r.URL.Query(), "sortOrder", "DESC")
//
//	params := db.GetProductsPaginatedParams{
//		SearchTerm: searchTerm,
//		MinPrice:   float64(minPrice),
//		MaxPrice:   float64(maxPrice),
//		Label:      label,
//		SortBy:     sortBy,
//		SortOrder:  sortOrder,
//		Limit:      int32(pageSize),
//		Offset:     int32(pageSize * (pageNumber - 1)),
//	}
//
//	products, err := app.queries.GetProductsPaginated(r.Context(), params)
//	if err != nil {
//		products = []db.GetProductsPaginatedRow{}
//	}
//
//	totalCount, err := app.queries.GetProductTotalCount(r.Context())
//	if err != nil {
//		totalCount = 0
//	}
//
//	totalPages := math.Ceil(float64(totalCount) / float64(pageSize))
//
//	paginatedResult := data.PaginatedResult{
//		Items:       products,
//		TotalCount:  int(totalCount),
//		HasNextPage: pageNumber < int(totalPages),
//		HasPrevPage: pageNumber > 1,
//		PageNumber:  pageNumber,
//		PageSize:    pageSize,
//	}
//
//	err = app.writeJson(w, http.StatusOK, payload.BaseResponse{
//		ResultCode:    SuccessCode,
//		ResultMessage: SuccessMessage,
//		Data:          paginatedResult,
//	}, nil)
//
//	if err != nil {
//		app.serverErrorResponse(w, err)
//	}
//}
//
//func (app *application) getProductDetailHandler(w http.ResponseWriter, r *http.Request) {
//	productId, err := app.readIDParam(r, "productId")
//	if err != nil {
//		app.badRequestErrorResponse(w, "Can not get product Id from params.")
//		return
//	}
//
//	product, err := app.queries.GetProductById(r.Context(), productId.String())
//	if err != nil {
//		app.notFoundErrorResponse(w, "Product does not exist!")
//		return
//	}
//
//	err = app.writeJson(w, http.StatusOK, payload.BaseResponse{
//		ResultCode:    SuccessCode,
//		ResultMessage: SuccessMessage,
//		Data:          product,
//	}, nil)
//
//	if err != nil {
//		app.serverErrorResponse(w, err)
//	}
//}
