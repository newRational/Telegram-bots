package show_schedule_telegram

import (
	"articles-tbot/lib/e"
	schedules_storage "articles-tbot/schedules_storage"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

const (
	HelpCmd  = "/help"
	StartCmd = "/start"
)

var nameOfDays = []string{
	"Понедельник",
	"Вторник",
	"Среда",
	"Четверг",
	"Пятница",
	"Суббота",
	"Воскресенье",
}

func (p *ScheduleProcessor) doCmd(text string, chatID int, username string) error {
	text = strings.TrimSpace(text)

	log.Printf("got new command '%s' from '%s'", text, username)

	for _, nameOfDay := range nameOfDays {
		if matched, _ := regexp.MatchString(text, nameOfDay); matched {
			if err := p.showSchedule(nameOfDay, username, chatID); err != nil {
				return err
			}
		}
	}

	return nil
}

func (p *ScheduleProcessor) showSchedule(dayName string, username string, chatID int) error {
	daySchedule, err := p.storage.PickSchedule(dayName, username)
	if err != nil {
		return err
	}

	return p.tg.SendMessage(chatID, daySchedule.Content)
}

func (p *ArticlesProcessor) readSchedule(day string, username string) (string, error) {
	userStudyGroupPath := filepath.Join(schedules_storage.StoragePath, username+".txt")

	studyGroup, err := os.ReadFile(userStudyGroupPath)
	if err != nil {
		return "", e.Wrap(fmt.Sprintf("can't read file '%s'", userStudyGroupPath), err)
	}

	studyGroupSchedulePath := filepath.Join(
		schedules_storage.StoragePath,
		string(studyGroup),
		schedules_storage.ByDaysDir,
		day,
	)

	daySchedule, err := os.ReadFile(studyGroupSchedulePath)
	if err != nil {
		return "", e.Wrap(fmt.Sprintf("can't read file '%s'", studyGroupSchedulePath), err)
	}

	return string(daySchedule), nil
}
