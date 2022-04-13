package services

import "github.com/putukrisna6/url-shortener/shortener/models"

type RedirectService interface {
	Find(code string) (*models.Redirect, error)
	Store(redirect *models.Redirect) error
}
