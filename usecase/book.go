package usecase

import (
	"context"
	"send2kobo/domain"
	"time"
)

type bookUsecase struct {
	bookRepository domain.BookRepository
	contextTimeout time.Duration
}

func NewBookUsecase(bookRepository domain.BookRepository, timeout time.Duration) domain.BookUsecase {
	return &bookUsecase{
		bookRepository: bookRepository,
		contextTimeout: timeout,
	}
}

func (bu *bookUsecase) Create(c context.Context, book *domain.Book) error {
	ctx, cancel := context.WithTimeout(c, bu.contextTimeout)
	defer cancel()
	return bu.bookRepository.Create(ctx, book)
}

func (bu *bookUsecase) Fetch(c context.Context) ([]domain.Book, error) {
	ctx, cancel := context.WithTimeout(c, bu.contextTimeout)
	defer cancel()
	return bu.bookRepository.Fetch(ctx)
}

func (bu *bookUsecase) GetByID(c context.Context, id string) (domain.Book, error) {
	ctx, cancel := context.WithTimeout(c, bu.contextTimeout)
	defer cancel()
	return bu.bookRepository.GetByID(ctx, id)
}

func (bu *bookUsecase) DeleteByID(c context.Context, id string) error {
	ctx, cancel := context.WithTimeout(c, bu.contextTimeout)
	defer cancel()
	return bu.bookRepository.DeleteByID(ctx, id)
}
