package helper

import "github.com/putukrisna6/url-shortener/shortener/models"

type RedirectSerializer interface {
	Decode(input []byte) (*models.Redirect, error)
	Encode(input *models.Redirect) ([]byte, error)
}
