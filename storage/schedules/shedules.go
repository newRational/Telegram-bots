package schedules

import (
	"articles-tbot/lib/e"
	"articles-tbot/schedules_storage"
	"articles-tbot/storage"
	"fmt"
	"os"
	"path/filepath"
)

type Storage struct {
	basePath string
}

func New(basePath string) Storage {
	return Storage{basePath: basePath}
}

func (s Storage) PickSchedule(userName, dayName string) (page *storage.Page, err error) {
	page = &storage.Page{
		Content:  "",
		UserName: userName,
	}

	userStudyGroupPath := filepath.Join(schedules_storage.StoragePath, userName+".txt")

	studyGroup, err := os.ReadFile(userStudyGroupPath)
	if err != nil {
		return page, e.Wrap(fmt.Sprintf("can't read file '%s'", userStudyGroupPath), err)
	}

	studyGroupSchedulePath := filepath.Join(
		schedules_storage.StoragePath,
		string(studyGroup),
		schedules_storage.ByDaysDir,
		dayName,
	)

	daySchedule, err := os.ReadFile(studyGroupSchedulePath)
	if err != nil {
		return page, e.Wrap(fmt.Sprintf("can't read file '%s'", studyGroupSchedulePath), err)
	}

	page.Content = string(daySchedule)

	return page, nil
}
