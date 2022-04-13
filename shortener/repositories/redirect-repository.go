package repositories

import "github.com/putukrisna6/url-shortener/shortener/models"

type RedirectRepository interface {
	Find(code string) (*models.Redirect, error)
	Store(redirect *models.Redirect) error
}
