package utils_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	. "github.com/franela/goblin"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/wanchai23chai/go-shared-lib/src/errmsg"
	"github.com/wanchai23chai/go-shared-lib/src/meta"
	"github.com/wanchai23chai/go-shared-lib/src/response"
	"github.com/wanchai23chai/go-shared-lib/src/utils"
)

func TestValidateBodyParser(t *testing.T) {
	g := Goblin(t)
	assert := assert.New(t)
	type BodyReq struct {
		Name string `json:"name" validate:"required"`
		Age  int    `json:"age" validate:"gte=0"`
	}

	g.Describe("Success Case", func() {
		var app *fiber.App
		g.BeforeEach(func() {
			app = InitServerTest()
		})
		g.AfterEach(func() {
			app.Shutdown()
		})
		g.It("Correct Body", func() {

			expect := BodyReq{
				Name: "unit-test",
				Age:  20,
			}
			app.Post("/api", func(ctx *fiber.Ctx) error {
				var inputBody = new(BodyReq)
				err := utils.ValidateBodyParser(ctx, inputBody)
				assert.Nil(err)
				return ctx.Status(http.StatusOK).JSON(inputBody)
			})
			payload, _ := json.Marshal(expect)
			req := httptest.NewRequest("POST", "/api", bytes.NewBuffer(payload))
			req.Header.Add("Content-Type", "application/json")
			res, _ := app.Test(req)
			defer res.Body.Close()
			resJson, err_res := ioutil.ReadAll(res.Body)
			expectJson, _ := json.Marshal(expect)
			assert.Equal(fiber.StatusOK, res.StatusCode)
			assert.Nil(err_res)
			assert.JSONEq(string(expectJson), string(resJson))
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
		g.It("Content type unsupport", func() {
			app.Post("/api", func(ctx *fiber.Ctx) error {
				var inputBody = new(BodyReq)
				err := utils.ValidateBodyParser(ctx, inputBody)
				assert.Equal(errmsg.ErrInvalidContentType, err)
				return ctx.Status(http.StatusOK).JSON(err)
			})
			payload := `{}`
			req := httptest.NewRequest("POST", "/api", strings.NewReader(payload))
			res, _ := app.Test(req)
			defer res.Body.Close()
			resJson, err_res := ioutil.ReadAll(res.Body)
			expectJson, _ := json.Marshal(errmsg.ErrInvalidContentType)
			assert.Equal(fiber.StatusOK, res.StatusCode)
			assert.Nil(err_res)
			assert.JSONEq(string(expectJson), string(resJson))
		})
		g.It("Content empty", func() {
			app.Post("/api", func(ctx *fiber.Ctx) error {
				var inputBody = new(BodyReq)
				err := utils.ValidateBodyParser(ctx, inputBody)
				assert.Nil(err)
				return ctx.Status(http.StatusOK).JSON(errmsg.ErrInvalidContentType)
			})
			req := httptest.NewRequest("POST", "/api", nil)
			res, _ := app.Test(req)
			defer res.Body.Close()
			resJson, err_res := ioutil.ReadAll(res.Body)
			expectJson, _ := json.Marshal(errmsg.ErrInvalidContentType)
			assert.Equal(fiber.StatusOK, res.StatusCode)
			assert.Nil(err_res)
			assert.JSONEq(string(expectJson), string(resJson))
		})
		g.It("Payload Invalid Type (not use validator)", func() {

			app.Post("/api", func(ctx *fiber.Ctx) error {
				var inputBody = new(BodyReq)
				err := utils.ValidateBodyParser(ctx, inputBody)
				//assert.Equal(errmsg.ErrInvalidContentType, err)
				return response.ResponseBadRequest(ctx, err)
			})
			payload := `{
				"name":"ddd",
				"age":"34d"
			}`
			req := httptest.NewRequest("POST", "/api", strings.NewReader(payload))
			req.Header.Add("Content-Type", "application/json")
			res, _ := app.Test(req)
			defer res.Body.Close()
			resJson, err_res := ioutil.ReadAll(res.Body)
			expect := meta.ErrorBadRequest.AppendMessage("Invalid field (age), it should be int.")
			expectJson, _ := json.Marshal(expect)
			assert.Equal(fiber.StatusBadRequest, res.StatusCode)
			assert.Nil(err_res)
			assert.JSONEq(string(expectJson), string(resJson))
		})
		g.It("Payload Invalid Type (use validator)", func() {

			app.Post("/api", func(ctx *fiber.Ctx) error {
				var inputBody = new(BodyReq)
				err := utils.ValidateBodyParser(ctx, inputBody)
				return response.ResponseBadRequest(ctx, err)
			})
			payload := `{
				"name":34,
				"age":1
			}`
			req := httptest.NewRequest("POST", "/api", strings.NewReader(payload))
			req.Header.Add("Content-Type", "application/json")
			res, _ := app.Test(req)
			defer res.Body.Close()
			resJson, err_res := ioutil.ReadAll(res.Body)
			expect := meta.ErrorBadRequest.AppendMessage("Invalid field (name), it should be string.")
			expectJson, _ := json.Marshal(expect)
			assert.Equal(fiber.StatusBadRequest, res.StatusCode)
			assert.Nil(err_res)
			assert.JSONEq(string(expectJson), string(resJson))
		})
		g.It("Payload Validation (use validator)", func() {

			app.Post("/api", func(ctx *fiber.Ctx) error {
				var inputBody = new(BodyReq)
				err := utils.ValidateBodyParser(ctx, inputBody)
				return response.ResponseBadRequest(ctx, err)
			})
			payload := `{
				"name":"34",
				"age":-99
			}`
			req := httptest.NewRequest("POST", "/api", strings.NewReader(payload))
			req.Header.Add("Content-Type", "application/json")
			res, _ := app.Test(req)
			defer res.Body.Close()
			resJson, err_res := ioutil.ReadAll(res.Body)
			expect := meta.ErrorBadRequest.AppendMessage("Invalid field (age), it should be gte 0.")
			expectJson, _ := json.Marshal(expect)
			assert.Equal(fiber.StatusBadRequest, res.StatusCode)
			assert.Nil(err_res)
			assert.JSONEq(string(expectJson), string(resJson))
		})
	})

}
