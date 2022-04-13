package shortener

import (
	"errors"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/putukrisna6/url-shortener/shortener/models"
	"github.com/putukrisna6/url-shortener/shortener/repositories"
	"github.com/putukrisna6/url-shortener/shortener/services"
	"github.com/teris-io/shortid"
)

var (
	ErrRedirectNotFound = "redirect not found"
	ErrRedirectInvalid  = "redirect invalid"

	validate = validator.New()
)

type redirectService struct {
	redirectRepo repositories.RedirectRepository
}

func NewRedirectService(redirectRepo repositories.RedirectRepository) services.RedirectService {
	return &redirectService{
		redirectRepo: redirectRepo,
	}
}

func (r *redirectService) Find(code string) (*models.Redirect, error) {
	return r.redirectRepo.Find(code)
}

func (r *redirectService) Store(redirect *models.Redirect) error {
	if err := validate.Struct(redirect); err != nil {
		return errors.New(ErrRedirectInvalid + "service.Redirect.Store")
	}
	redirect.Code = shortid.MustGenerate()
	redirect.CreatedAt = time.Now().UTC().Unix()
	return r.redirectRepo.Store(redirect)
}
