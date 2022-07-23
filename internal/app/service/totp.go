package service

import (
	"bytes"
	"encoding/base64"
	"image/png"

	"github.com/pinyi-lee/core.account.2fa.git/internal/pkg/config"
	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
)

func generateKey(accountId string, serviceName string) (*otp.Key, error) {
	key, err := totp.Generate(totp.GenerateOpts{
		AccountName: accountId,
		Issuer:      serviceName,
	})
	if err != nil {
		return nil, err
	}

	return key, nil

}

func getQrCode(key *otp.Key) (string, error) {
	var buf bytes.Buffer
	img, err := key.Image(config.Env.ImageWidth, config.Env.ImageHeight)
	if err != nil {
		return "", err
	}

	err = png.Encode(&buf, img)
	if err != nil {
		return "", err
	}

	qrCode := base64.StdEncoding.EncodeToString(buf.Bytes())
	return qrCode, err
}

func verify(passcode string, secret string) bool {
	return totp.Validate(passcode, secret)
}
