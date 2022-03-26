package utils_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"sort"
	"strings"
	"testing"
	"time"

	. "github.com/franela/goblin"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/wanchai23chai/go-shared-lib/errmsg"
	"github.com/wanchai23chai/go-shared-lib/logger"
	"github.com/wanchai23chai/go-shared-lib/response"
	"github.com/wanchai23chai/go-shared-lib/utils"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitServerTest() *fiber.App {
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
		ReadTimeout:           5 * time.Second,
		ErrorHandler:          response.DefaultErrorResponse,
	})
	if err := logger.New(); err != nil {
		panic(err)
	}
	return app
}
func InitDB() (*gorm.DB, sqlmock.Sqlmock) {
	sqlDB, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      sqlDB,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db, sqlMock
}

type TestCase struct {
	name     string
	input    interface{}
	expected interface{}
}

func TestCheckPagination(t *testing.T) {
	g := Goblin(t)
	assert := assert.New(t)
	g.Describe("Success Case", func() {
		var app *fiber.App
		var db *gorm.DB
		var SQLMock sqlmock.Sqlmock
		g.BeforeEach(func() {
			app = InitServerTest()
			db, SQLMock = InitDB()

		})
		g.AfterEach(func() {
			app.Shutdown()
		})
		testCaseHttp := []TestCase{
			{
				name: "limit=3 page=1",

				input: map[string]interface{}{
					"http_path": fmt.Sprintf("/api?limit=%v&page=%v", 3, 1),
				},
				expected: map[string]interface{}{
					"pagination": utils.Pagination{
						Limit: 3,
						Page:  1,
					},
					"meta": map[string]interface{}{
						"NextPage": "",
						"PrevPage": "",
					},
				},
			},
			{
				name: "limit=3 page=1 sort=id-asc",

				input: map[string]interface{}{
					"http_path": fmt.Sprintf("/api?limit=%v&page=%v&sort=id-asc", 3, 1),
				},
				expected: map[string]interface{}{
					"pagination": utils.Pagination{
						Limit:      3,
						Page:       1,
						Sort:       "id-asc",
						TotalRows:  10,
						TotalPages: 4,
					},
					"meta": map[string]interface{}{
						"NextPage": "/api?limit=3&page=2&sort=id-asc",
						"PrevPage": "",
					},
				},
			},
			{
				name: "limit=3 page=2 sort=id-asc",

				input: map[string]interface{}{
					"http_path": fmt.Sprintf("/api?limit=%v&page=%v&sort=id-asc", 3, 2),
				},
				expected: map[string]interface{}{
					"pagination": utils.Pagination{
						Limit:      3,
						Page:       2,
						Sort:       "id-asc",
						TotalRows:  10,
						TotalPages: 4,
					},
					"meta": map[string]interface{}{
						"NextPage": "/api?limit=3&page=3&sort=id-asc",
						"PrevPage": "/api?limit=3&page=1&sort=id-asc",
					},
				},
			},
			{
				name: "not provide query",

				input: map[string]interface{}{
					"http_path": "/api",
				},
				expected: map[string]interface{}{
					"pagination": utils.Pagination{
						Limit:      -1,
						Page:       1,
						TotalRows:  10,
						TotalPages: 1,
					},
					"meta": map[string]interface{}{
						"NextPage": "",
						"PrevPage": "",
					},
				},
			},
			{
				name: "limit=0 page=0",

				input: map[string]interface{}{
					"http_path": "/api?limit=0&page=0",
				},
				expected: map[string]interface{}{
					"pagination": utils.Pagination{
						Limit:      10,
						Page:       1,
						TotalRows:  10,
						TotalPages: 1,
					},
					"meta": map[string]interface{}{
						"NextPage": "",
						"PrevPage": "",
					},
				},
			},
			{
				name: "limit=10 page=2 total_page=1 total_raw=10",

				input: map[string]interface{}{
					"http_path": fmt.Sprintf("/api?limit=%v&page=%v&sort=id-asc", 10, 2),
				},
				expected: map[string]interface{}{
					"pagination": utils.Pagination{
						Limit:      10,
						Page:       2,
						Sort:       "id-asc",
						TotalRows:  10,
						TotalPages: 1,
					},
					"meta": map[string]interface{}{
						"NextPage": "",
						"PrevPage": "/api?limit=10&page=1&sort=id-asc",
					},
				},
			},
			{
				name: "limit=-1 total_raw=0",

				input: map[string]interface{}{
					"http_path": fmt.Sprintf("/api?limit=%v&sort=id-asc", -1),
				},
				expected: map[string]interface{}{
					"pagination": utils.Pagination{
						Limit:      -1,
						Page:       1,
						Sort:       "id-asc",
						TotalRows:  0,
						TotalPages: 1,
					},
					"meta": map[string]interface{}{
						"NextPage": "",
						"PrevPage": "",
					},
				},
			},
			{
				name: "[Provide other query includes=test] limit=10 page=2 total_page=1 total_raw=10",

				input: map[string]interface{}{
					"http_path": fmt.Sprintf("/api?limit=%v&page=%v&sort=id-asc&includes=test", 10, 2),
				},
				expected: map[string]interface{}{
					"pagination": utils.Pagination{
						Limit:      10,
						Page:       2,
						Sort:       "id-asc",
						TotalRows:  10,
						TotalPages: 1,
					},
					"meta": map[string]interface{}{
						"NextPage": "",
						"PrevPage": "/api?limit=10&page=1&sort=id-asc&includes=test",
					},
				},
			},
		}
		for caseNo, v := range testCaseHttp {
			item := v
			g.It(fmt.Sprintf("Case : %v (%v)", caseNo+1, item.name), func() {
				item_input := item.input.(map[string]interface{})
				item_expect := item.expected.(map[string]interface{})
				expectMeta := item_expect["meta"].(map[string]interface{})
				expectPagination := item_expect["pagination"].(utils.Pagination)
				app.Get("/api", func(ctx *fiber.Ctx) error {
					pg, err := utils.CheckPagination(ctx)
					if err != nil {
						return err
					}
					assert.Nil(err)
					pg.TotalPages = expectPagination.TotalPages
					pg.TotalRows = expectPagination.TotalRows
					metaData := utils.GenerateMetaData(&pg)
					expectNextPage := strings.Split(expectMeta["NextPage"].(string), "")
					gotNextPage := strings.Split(metaData.NextPage, "")
					sort.Strings(expectNextPage)
					sort.Strings(gotNextPage)
					assert.ElementsMatch(expectNextPage, gotNextPage, fmt.Sprintf("Invalid NextPage : case (%v)", item.name))
					expectPrevPage := strings.Split(expectMeta["PrevPage"].(string), "")
					gotPrevPage := strings.Split(metaData.PrevPage, "")
					sort.Strings(expectPrevPage)
					sort.Strings(gotPrevPage)
					assert.Equal(expectPrevPage, gotPrevPage, fmt.Sprintf("Invalid PrevPage : case (%v)", item.name))
					return ctx.Status(http.StatusOK).JSON(pg)
				})

				req := httptest.NewRequest("GET", item_input["http_path"].(string), nil)
				res, _ := app.Test(req)
				defer res.Body.Close()
				body, err_res := ioutil.ReadAll(res.Body)
				resJson, _ := json.Marshal(item_expect["pagination"])
				assert.Equal(fiber.StatusOK, res.StatusCode)
				assert.Nil(err_res)
				assert.JSONEq(string(resJson), string(body), fmt.Sprintf("Invalid Json Response : case (%v)", item.name))
			})
		}
		g.It("Test Gorm Paginate", func() {

			SQLMock.ExpectQuery("SELECT count(*) FROM `models`").WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(10))
			SQLMock.ExpectQuery("SELECT * FROM `models` ORDER BY id-asc OFFSET 1").WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(10))
			offset := 1
			sortList := []string{
				"id-asc",
			}
			pg := utils.Pagination{
				Limit:     -1,
				Page:      1,
				TotalRows: 10,
				Offset:    &offset,
				SortList:  &sortList,
			}
			type Model struct {
				ID   uint   `gorm:"column:id;type:int(10) unsigned;primaryKey" json:"id,omitempty"`
				Name string `gorm:"column:name_th;type:text" json:"name,omitempty"`
			}
			var out []Model
			db.Model(&Model{}).Scopes(utils.Paginate(&pg)).Find(&out)
			acct, _ := json.Marshal(out)
			assert.Equal("null", string(acct))
		})
	})

	g.Describe("Fail Case", func() {
		var app *fiber.App
		g.BeforeEach(func() {
			app = InitServerTest()
		})
		g.AfterEach(func() {
			app.Shutdown()
		})
		testCaseHttp := []TestCase{
			{
				name: "page=invalid",

				input: map[string]interface{}{
					"http_path": fmt.Sprintf("/?page=%v", "invalid"),
				},
				expected: map[string]interface{}{
					"error":     errmsg.ErrInvalid_Page_Format,
					"http_code": 400,
				},
			},
			{
				name: "limit=invalid",

				input: map[string]interface{}{
					"http_path": fmt.Sprintf("/?limit=%v", "invalid"),
				},
				expected: map[string]interface{}{
					"error":     errmsg.ErrInvalid_Limit_Format,
					"http_code": 400,
				},
			},
			{
				name: "sort=id-invalid",

				input: map[string]interface{}{
					"http_path": "/?sort=id-invalid",
				},
				expected: map[string]interface{}{
					"error":     errmsg.ErrInvalid_Sort_Format,
					"http_code": 400,
				},
			},
		}
		for _, v := range testCaseHttp {
			item := v
			g.It(item.name, func() {
				app.Get("/", func(ctx *fiber.Ctx) error {
					pg, err := utils.CheckPagination(ctx)
					if err != nil {
						return err
					}
					assert.Nil(err)
					return ctx.Status(http.StatusOK).JSON(pg)
				})
				item_input := item.input.(map[string]interface{})
				item_expect := item.expected.(map[string]interface{})
				req := httptest.NewRequest("GET", item_input["http_path"].(string), nil)
				res, _ := app.Test(req)
				defer res.Body.Close()
				resJson, err_res := ioutil.ReadAll(res.Body)
				expectJson, _ := json.Marshal(item_expect["error"])
				assert.Equal(item_expect["http_code"], res.StatusCode)
				assert.Nil(err_res)
				assert.JSONEq(string(expectJson), string(resJson))
			})
		}
	})

}
