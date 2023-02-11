package save_articles_telegram

import (
	"articles-tbot/files_storage"
	"articles-tbot/lib/e"
	"articles-tbot/storage"
	"errors"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

const (
	RndCmd   = "/rnd"
	HelpCmd  = "/help"
	StartCmd = "/start"
)

func (p *Processor) doCmd(text string, chatID int, username string) error {
	text = strings.TrimSpace(text)

	log.Printf("got new command '%s' from '%s'", text, username)

	if isAddCmd(text) {
		return p.savePage(chatID, text, username)
	}

	switch text {
	case RndCmd:
		return p.sendRandom(chatID, username)
	case HelpCmd:
		return p.sendHelp(chatID)
	case StartCmd:
		return p.startConversation(chatID, username)
	default:
		return p.tg.SendMessage(chatID, msgUnknownCommand)
	}
}

func (p *Processor) savePage(chatID int, pageURL string, username string) (err error) {
	defer func() { err = e.WrapIfErr("can't execute command: save page", err) }()

	page := &storage.Page{
		Content:  pageURL,
		UserName: username,
	}

	isExist, err := p.storage.IsExist(page)
	if err != nil {
		return err
	}

	if isExist {
		return p.tg.SendMessage(chatID, msgAlreadyExists)
	}

	if err := p.storage.Save(page); err != nil {
		return err
	}

	if err := p.tg.SendMessage(chatID, msgSaved); err != nil {
		return err
	}

	return nil
}

func (p *Processor) sendRandom(chatID int, username string) (err error) {
	defer func() { err = e.WrapIfErr("can't do command: can't send random", err) }()

	page, err := p.storage.PickRandom(username)
	if err != nil && !errors.Is(err, storage.ErrNoSavePages) {
		return err
	}

	if errors.Is(err, storage.ErrNoSavePages) {
		return p.tg.SendMessage(chatID, msgNoSavedPages)
	}

	if err := p.tg.SendMessage(chatID, page.Content); err != nil {
		return err
	}

	return p.storage.Remove(page)
}

func (p *Processor) sendHelp(chatID int) error {
	return p.tg.SendMessage(chatID, msgHelp)
}

func (p *Processor) startConversation(chatID int, username string) error {
	if err := createUserDir(username); err != nil {
		return err
	}
	return p.tg.SendMessage(chatID, msgHello)
}

func createUserDir(username string) (err error) {
	defer func() { err = e.WrapIfErr("can't create user dir", err) }()

	fPath := filepath.Join(files_storage.StoragePath, username)

	if err := os.Mkdir(fPath, files_storage.DefaultPerm); err != nil {
		return err
	}

	return nil
}

func isAddCmd(text string) bool {
	return isURL(text)
}

func isURL(text string) bool {
	u, err := url.Parse(text)

	return err == nil && u.Host != ""
}
