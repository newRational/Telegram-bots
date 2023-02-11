package storage

import (
	"articles-tbot/lib/e"
	"crypto/sha1"
	"errors"
	"fmt"
	"io"
)

type ArticlesStorage interface {
	Save(p *Page) error
	PickRandom(userName string) (*Page, error)
	Remove(p *Page) error
	IsExist(p *Page) (bool, error)
}

type ScheduleStorage interface {
	PickSchedule(dayName, userName string) (*Page, error)
}

var ErrNoSavePages = errors.New("no saved pages")

type Page struct {
	Content  string
	UserName string
}

func (p Page) Hash() (string, error) {
	h := sha1.New()

	if _, err := io.WriteString(h, p.Content); err != nil {
		return "", e.Wrap("can't calculate hash", err)
	}

	if _, err := io.WriteString(h, p.UserName); err != nil {
		return "", e.Wrap("can't calculate hash", err)
	}

	return fmt.Sprintf("%x", h.Sum(nil)), nil
}
