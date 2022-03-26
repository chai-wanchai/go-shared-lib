package utils

import (
	"fmt"
	"go-shared-lib/pkg/errmsg"
	"math"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CheckPagination(ctx *fiber.Ctx) (Pagination, error) {
	result := Pagination{}
	pages, err_p := strconv.Atoi(ctx.Query("page", "1"))
	if err_p != nil {
		return result, errmsg.ErrInvalid_Page_Format
	}
	limits, err_l := strconv.Atoi(ctx.Query("limit", "-1"))
	if err_l != nil {
		return result, errmsg.ErrInvalid_Limit_Format
	}
	sort := ctx.Query("sort", "")
	result.Limit = limits
	result.Page = pages
	result.Sort = sort
	result.CTX = ctx
	err := ValidatePagination(&result)
	if err != nil {
		return result, err
	}
	return result, nil
}
func (m *MetaData) AddQueryAuto(ctx *fiber.Ctx) *MetaData {
	if ctx.Context().QueryArgs().String() == "" {
		return m
	}
	rawQuery := QueryStringToMap(ctx.Context().QueryArgs().String())
	var conditionQuery map[string]string
	var listQuery []string
	if m.NextPage != "" {
		conditionQuery = QueryStringToMap(m.NextPage)
	}
	if m.PrevPage != "" {
		conditionQuery = QueryStringToMap(m.PrevPage)
	}
	for k, v := range rawQuery {
		if _, found := conditionQuery[k]; !found {
			listQuery = append(listQuery, fmt.Sprint(k, "=", v))
		}
	}
	m.AddQuery(listQuery...)
	return m
}

func (m *MetaData) EditUri(ctx *fiber.Ctx) *MetaData {
	host := string(ctx.Request().URI().Path())
	if m.NextPage != "" {
		m.NextPage = fmt.Sprint(host, m.NextPage)
	}
	if m.PrevPage != "" {
		m.PrevPage = fmt.Sprint(host, m.PrevPage)
	}
	m.AddQueryAuto(ctx)
	return m
}

func GenerateMetaData(pagination *Pagination) MetaData {
	prev := map[string]interface{}{
		"page": pagination.Page - 1,
	}
	next := map[string]interface{}{
		"page": pagination.Page + 1,
	}
	prevPage := fmt.Sprint("?limit=", pagination.Limit)
	nextPage := fmt.Sprint("?limit=", pagination.Limit)

	if pagination.Page-1 <= 0 || pagination.Page-1 > pagination.TotalPages {
		prevPage = ""
	} else if pagination.Page > pagination.TotalPages {
		prev["page"] = pagination.TotalPages
	}
	if pagination.Page > pagination.TotalPages || pagination.Page+1 > pagination.TotalPages {
		nextPage = ""
	}
	if pagination.Sort != "" {
		next["sort"] = pagination.Sort
		prev["sort"] = pagination.Sort
	}
	if prevPage != "" {
		for key, item := range prev {
			prevPage = fmt.Sprint(prevPage, "&", key, "=", item)
		}
	}
	if nextPage != "" {
		for key, item := range next {
			nextPage = fmt.Sprint(nextPage, "&", key, "=", item)
		}
	}

	result := MetaData{
		Pagination: *pagination,
		NextPage:   nextPage,
		PrevPage:   prevPage,
	}
	if pagination.CTX != nil {
		result.EditUri(pagination.CTX)
	}
	return result
}

func (m *MetaData) AddQuery(query ...string) *MetaData {
	var queryString string
	for _, v := range query {
		queryString = fmt.Sprint(queryString, "&", v)
	}
	if m.NextPage != "" {
		m.NextPage = fmt.Sprint(m.NextPage, queryString)
	}
	if m.PrevPage != "" {
		m.PrevPage = fmt.Sprint(m.PrevPage, queryString)
	}
	return m
}

func QueryStringToMap(query string) map[string]string {
	var mapQuery = make(map[string]string)
	if strings.Contains(query, "?") {
		query = strings.Split(query, "?")[1]
	}
	for _, v := range strings.Split(query, "&") {
		item := strings.Split(v, "=")
		mapQuery[item[0]] = item[1]
	}
	return mapQuery
}
func ValidatePagination(pagination *Pagination) error {
	var sortFeild []string
	if pagination.Page <= 0 {
		pagination.Page = 1
	}
	if pagination.Limit == 0 {
		pagination.Limit = 10
	}
	if pagination.Sort != "" {
		sortAllowKey := map[string]string{"asc": "asc", "desc": "desc"}
		listSort := strings.Split(pagination.Sort, ",")
		for _, v := range listSort {
			item := strings.Split(v, "-")
			if val, ok := sortAllowKey[item[1]]; ok {
				sortFeild = append(sortFeild, fmt.Sprint(item[0], " ", val))
			} else {
				return errmsg.ErrInvalid_Sort_Format
			}
		}
	}
	offset := (pagination.Page - 1) * pagination.Limit
	totalPages := int(math.Ceil(float64(pagination.TotalRows) / float64(pagination.Limit)))
	pagination.TotalPages = totalPages
	pagination.Offset = &offset
	pagination.SortList = &sortFeild
	return nil
}

func Paginate(pagination *Pagination) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		db.Count(&pagination.TotalRows)
		var totalPages int = 0
		if pagination.TotalRows != 0 {
			totalPages = int(math.Ceil(float64(pagination.TotalRows) / float64(pagination.Limit)))
			if pagination.Limit < 0 {
				totalPages = 1
			}
		}
		pagination.TotalPages = totalPages
		for _, v := range *pagination.SortList {
			db.Order(v)
		}
		return db.Offset(*pagination.Offset).Limit(pagination.Limit)
	}
}
