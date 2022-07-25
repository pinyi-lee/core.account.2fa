package test

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/pinyi-lee/core.account.2fa.git/internal/app/service"
	"github.com/pinyi-lee/core.account.2fa.git/internal/pkg/config"
	"github.com/pinyi-lee/core.account.2fa.git/internal/pkg/model"
	"github.com/stretchr/testify/assert"
)

func TestInitTotp(t *testing.T) {
	t.Run("[InitTotp] body error, should return 400", func(t *testing.T) {
		body := `{"accountId": "","serviceName": ""}`
		w, _ := HttpPost("/pviv/2fa/v1/totp/init", body, nil)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("[InitTotp] totp already created, should return 409", func(t *testing.T) {
		body := `{"accountId": "Init001","serviceName": "Init001"}`
		w, _ := HttpPost("/pviv/2fa/v1/totp/init", body, nil)
		assert.Equal(t, http.StatusOK, w.Code)
		initTotpRes := model.InitTotpRes{}
		json.Unmarshal([]byte(w.Body.String()), &initTotpRes)

		code, _ := service.GenerateCode(initTotpRes.Secret)

		body = "{\"accountId\": \"Init001\",\"serviceName\": \"Init001\",\"passcode\":\"" + code + "\"}"
		w, _ = HttpPost("/pviv/2fa/v1/totp/enable", body, nil)
		assert.Equal(t, http.StatusOK, w.Code)

		body = `{"accountId": "Init001","serviceName": "Init001"}`
		w, _ = HttpPost("/pviv/2fa/v1/totp/init", body, nil)
		assert.Equal(t, http.StatusConflict, w.Code)
	})

	t.Run("[InitTotp] success, should return 200", func(t *testing.T) {
		body := `{"accountId": "Init002","serviceName": "Init002"}`
		w, _ := HttpPost("/pviv/2fa/v1/totp/init", body, nil)
		assert.Equal(t, http.StatusOK, w.Code)
	})
}

func TestEnableTotp(t *testing.T) {
	t.Run("[EnableTotp] body error, should return 400", func(t *testing.T) {
		body := `{"accountId": "","serviceName": "","passcode": ""}`
		w, _ := HttpPost("/pviv/2fa/v1/totp/enable", body, nil)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("[EnableTotp] data not found, should return 400", func(t *testing.T) {
		body := `{"accountId": "Enable001","serviceName": "Enable001","passcode": "123456"}`
		w, _ := HttpPost("/pviv/2fa/v1/totp/enable", body, nil)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("[EnableTotp] totp already created, should return 409", func(t *testing.T) {
		body := `{"accountId": "Enable002","serviceName": "Enable002"}`
		w, _ := HttpPost("/pviv/2fa/v1/totp/init", body, nil)
		assert.Equal(t, http.StatusOK, w.Code)
		initTotpRes := model.InitTotpRes{}
		json.Unmarshal([]byte(w.Body.String()), &initTotpRes)

		code, _ := service.GenerateCode(initTotpRes.Secret)

		body = "{\"accountId\": \"Enable002\",\"serviceName\": \"Enable002\",\"passcode\":\"" + code + "\"}"
		w, _ = HttpPost("/pviv/2fa/v1/totp/enable", body, nil)
		assert.Equal(t, http.StatusOK, w.Code)

		w, _ = HttpPost("/pviv/2fa/v1/totp/enable", body, nil)
		assert.Equal(t, http.StatusConflict, w.Code)
	})

	t.Run("[EnableTotp] passcode verify fail, should return 400", func(t *testing.T) {
		body := `{"accountId": "Enable003","serviceName": "Enable003"}`
		w, _ := HttpPost("/pviv/2fa/v1/totp/init", body, nil)
		assert.Equal(t, http.StatusOK, w.Code)
		initTotpRes := model.InitTotpRes{}
		json.Unmarshal([]byte(w.Body.String()), &initTotpRes)

		body = `{"accountId": "Enable003","serviceName": "Enable003","passcode": "123456"}`
		w, _ = HttpPost("/pviv/2fa/v1/totp/enable", body, nil)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("[EnableTotp] success, should return 200", func(t *testing.T) {
		body := `{"accountId": "Enable004","serviceName": "Enable004"}`
		w, _ := HttpPost("/pviv/2fa/v1/totp/init", body, nil)
		assert.Equal(t, http.StatusOK, w.Code)
		initTotpRes := model.InitTotpRes{}
		json.Unmarshal([]byte(w.Body.String()), &initTotpRes)

		code, _ := service.GenerateCode(initTotpRes.Secret)

		body = "{\"accountId\": \"Enable004\",\"serviceName\": \"Enable004\",\"passcode\":\"" + code + "\"}"
		w, _ = HttpPost("/pviv/2fa/v1/totp/enable", body, nil)
		assert.Equal(t, http.StatusOK, w.Code)
	})
}

func TestDisableTotp(t *testing.T) {
	t.Run("[DisableTotp] body error, should return 400", func(t *testing.T) {
		body := `{"accountId": "","serviceName": "","passcode": ""}`
		w, _ := HttpPost("/pviv/2fa/v1/totp/disable", body, nil)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("[DisableTotp] data not found, should return 400", func(t *testing.T) {
		body := `{"accountId": "Disable001","serviceName": "Disable001","passcode": "123456"}`
		w, _ := HttpPost("/pviv/2fa/v1/totp/disable", body, nil)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("[DisableTotp] totp not already created, should return 409", func(t *testing.T) {
		body := `{"accountId": "Disable002","serviceName": "Disable002"}`
		w, _ := HttpPost("/pviv/2fa/v1/totp/init", body, nil)
		assert.Equal(t, http.StatusOK, w.Code)

		body = `{"accountId": "Disable002","serviceName": "Disable002","passcode": "123456"}`
		w, _ = HttpPost("/pviv/2fa/v1/totp/disable", body, nil)
		assert.Equal(t, http.StatusConflict, w.Code)
	})

	t.Run("[DisableTotp] passcode verify fail, should return 400", func(t *testing.T) {
		body := `{"accountId": "Disable003","serviceName": "Disable003"}`
		w, _ := HttpPost("/pviv/2fa/v1/totp/init", body, nil)
		assert.Equal(t, http.StatusOK, w.Code)
		initTotpRes := model.InitTotpRes{}
		json.Unmarshal([]byte(w.Body.String()), &initTotpRes)

		code, _ := service.GenerateCode(initTotpRes.Secret)

		body = "{\"accountId\": \"Disable003\",\"serviceName\": \"Disable003\",\"passcode\":\"" + code + "\"}"
		w, _ = HttpPost("/pviv/2fa/v1/totp/enable", body, nil)
		assert.Equal(t, http.StatusOK, w.Code)

		body = `{"accountId": "Disable003","serviceName": "Disable003","passcode": "123456"}`
		w, _ = HttpPost("/pviv/2fa/v1/totp/disable", body, nil)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("[DisableTotp] success, should return 200", func(t *testing.T) {
		body := `{"accountId": "Disable004","serviceName": "Disable004"}`
		w, _ := HttpPost("/pviv/2fa/v1/totp/init", body, nil)
		assert.Equal(t, http.StatusOK, w.Code)
		initTotpRes := model.InitTotpRes{}
		json.Unmarshal([]byte(w.Body.String()), &initTotpRes)

		code, _ := service.GenerateCode(initTotpRes.Secret)

		body = "{\"accountId\": \"Disable004\",\"serviceName\": \"Disable004\",\"passcode\":\"" + code + "\"}"
		w, _ = HttpPost("/pviv/2fa/v1/totp/enable", body, nil)
		assert.Equal(t, http.StatusOK, w.Code)

		body = "{\"accountId\": \"Disable004\",\"serviceName\": \"Disable004\",\"passcode\":\"" + code + "\"}"
		w, _ = HttpPost("/pviv/2fa/v1/totp/disable", body, nil)
		assert.Equal(t, http.StatusOK, w.Code)
	})

}

func TestVerifyTotp(t *testing.T) {
	t.Run("[VerifyTotp] body error, should return 400", func(t *testing.T) {
		body := `{"accountId": "","serviceName": "","passcode": ""}`
		w, _ := HttpPost("/pviv/2fa/v1/totp/verify", body, nil)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("[VerifyTotp] data not found, should return 400", func(t *testing.T) {
		body := `{"accountId": "Verify001","serviceName": "Verify001","passcode": "123456"}`
		w, _ := HttpPost("/pviv/2fa/v1/totp/verify", body, nil)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("[VerifyTotp] totp not already created, should return 409", func(t *testing.T) {
		body := `{"accountId": "Verify002","serviceName": "Verify002"}`
		w, _ := HttpPost("/pviv/2fa/v1/totp/init", body, nil)
		assert.Equal(t, http.StatusOK, w.Code)

		body = `{"accountId": "Verify002","serviceName": "Verify002","passcode": "123456"}`
		w, _ = HttpPost("/pviv/2fa/v1/totp/verify", body, nil)
		assert.Equal(t, http.StatusConflict, w.Code)
	})

	t.Run("[VerifyTotp] passcode verify fail, should return 400", func(t *testing.T) {
		body := `{"accountId": "Verify003","serviceName": "Verify003"}`
		w, _ := HttpPost("/pviv/2fa/v1/totp/init", body, nil)
		assert.Equal(t, http.StatusOK, w.Code)
		initTotpRes := model.InitTotpRes{}
		json.Unmarshal([]byte(w.Body.String()), &initTotpRes)

		code, _ := service.GenerateCode(initTotpRes.Secret)

		body = "{\"accountId\": \"Verify003\",\"serviceName\": \"Verify003\",\"passcode\":\"" + code + "\"}"
		w, _ = HttpPost("/pviv/2fa/v1/totp/enable", body, nil)
		assert.Equal(t, http.StatusOK, w.Code)

		body = `{"accountId": "Verify003","serviceName": "Verify003","passcode": "123456"}`
		w, _ = HttpPost("/pviv/2fa/v1/totp/verify", body, nil)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("[VerifyTotp] success, should return 200", func(t *testing.T) {
		body := `{"accountId": "Verify004","serviceName": "Verify004"}`
		w, _ := HttpPost("/pviv/2fa/v1/totp/init", body, nil)
		assert.Equal(t, http.StatusOK, w.Code)
		initTotpRes := model.InitTotpRes{}
		json.Unmarshal([]byte(w.Body.String()), &initTotpRes)

		code, _ := service.GenerateCode(initTotpRes.Secret)

		body = "{\"accountId\": \"Verify004\",\"serviceName\": \"Verify004\",\"passcode\":\"" + code + "\"}"
		w, _ = HttpPost("/pviv/2fa/v1/totp/enable", body, nil)
		assert.Equal(t, http.StatusOK, w.Code)

		body = "{\"accountId\": \"Verify004\",\"serviceName\": \"Verify004\",\"passcode\":\"" + code + "\"}"
		w, _ = HttpPost("/pviv/2fa/v1/totp/verify", body, nil)
		assert.Equal(t, http.StatusOK, w.Code)
	})
}

func TesGetTotpStatus(t *testing.T) {
	t.Run("[GetTotpStatus] body error, should return 400", func(t *testing.T) {
		w, _ := HttpGet("/pviv/2fa/v1/totp/status", nil)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("[GetTotpStatus] success (none), should return 200", func(t *testing.T) {
		w, _ := HttpGet("/pviv/2fa/v1/totp/status?accountId=Get001&serviceName=Get001", nil)
		assert.Equal(t, http.StatusOK, w.Code)
		getTotpStatusRes := model.GetTotpStatusRes{}
		json.Unmarshal([]byte(w.Body.String()), &getTotpStatusRes)
		assert.Equal(t, getTotpStatusRes.Status, string(config.None))
	})

	t.Run("[GetTotpStatus] success (init), should return 200", func(t *testing.T) {
		body := `{"accountId": "Get002","serviceName": "Get002"}`
		w, _ := HttpPost("/pviv/2fa/v1/totp/init", body, nil)
		assert.Equal(t, http.StatusOK, w.Code)

		w, _ = HttpGet("/pviv/2fa/v1/totp/status?accountId=Get002&serviceName=Get002", nil)
		assert.Equal(t, http.StatusOK, w.Code)
		getTotpStatusRes := model.GetTotpStatusRes{}
		json.Unmarshal([]byte(w.Body.String()), &getTotpStatusRes)
		assert.Equal(t, getTotpStatusRes.Status, string(config.Init))
	})

	t.Run("[GetTotpStatus] success (created), should return 200", func(t *testing.T) {
		body := `{"accountId": "Get003","serviceName": "Get003"}`
		w, _ := HttpPost("/pviv/2fa/v1/totp/init", body, nil)
		assert.Equal(t, http.StatusOK, w.Code)
		initTotpRes := model.InitTotpRes{}
		json.Unmarshal([]byte(w.Body.String()), &initTotpRes)

		code, _ := service.GenerateCode(initTotpRes.Secret)

		body = "{\"accountId\": \"Get003\",\"serviceName\": \"Get003\",\"passcode\":\"" + code + "\"}"
		w, _ = HttpPost("/pviv/2fa/v1/totp/enable", body, nil)
		assert.Equal(t, http.StatusOK, w.Code)

		w, _ = HttpGet("/pviv/2fa/v1/totp/status?accountId=Get003&serviceName=Get003", nil)
		assert.Equal(t, http.StatusOK, w.Code)
		getTotpStatusRes := model.GetTotpStatusRes{}
		json.Unmarshal([]byte(w.Body.String()), &getTotpStatusRes)
		assert.Equal(t, getTotpStatusRes.Status, string(config.Created))
	})

	t.Run("[GetTotpStatus] after delete, success (none), should return 200", func(t *testing.T) {
		body := `{"accountId": "Get004","serviceName": "Get004"}`
		w, _ := HttpPost("/pviv/2fa/v1/totp/init", body, nil)
		assert.Equal(t, http.StatusOK, w.Code)
		initTotpRes := model.InitTotpRes{}
		json.Unmarshal([]byte(w.Body.String()), &initTotpRes)

		code, _ := service.GenerateCode(initTotpRes.Secret)

		body = "{\"accountId\": \"Get004\",\"serviceName\": \"Get004\",\"passcode\":\"" + code + "\"}"
		w, _ = HttpPost("/pviv/2fa/v1/totp/enable", body, nil)
		assert.Equal(t, http.StatusOK, w.Code)

		body = "{\"accountId\": \"Get004\",\"serviceName\": \"Get004\",\"passcode\":\"" + code + "\"}"
		w, _ = HttpPost("/pviv/2fa/v1/totp/disable", body, nil)
		assert.Equal(t, http.StatusOK, w.Code)

		w, _ = HttpGet("/pviv/2fa/v1/totp/status?accountId=Get004&serviceName=Get004", nil)
		assert.Equal(t, http.StatusOK, w.Code)
		getTotpStatusRes := model.GetTotpStatusRes{}
		json.Unmarshal([]byte(w.Body.String()), &getTotpStatusRes)
		assert.Equal(t, getTotpStatusRes.Status, string(config.None))
	})
}
