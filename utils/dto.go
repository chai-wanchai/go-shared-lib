package utils

import "github.com/gofiber/fiber/v2"

type Pagination struct {
	Limit      int        `json:"limit"`
	Page       int        `json:"page"`
	Sort       string     `json:"sort,omitempty"`
	TotalRows  int64      `json:"total_rows"`
	TotalPages int        `json:"total_pages"`
	SortList   *[]string  `json:"-"`
	Offset     *int       `json:"-"`
	CTX        *fiber.Ctx `json:"-"`
}
type MetaData struct {
	Pagination
	NextPage string `json:"next_page"`
	PrevPage string `json:"prev_page"`
}
type UserAdminAuth struct {
	ID *int
}
