package parsers

import (
	"articles-tbot/schedules_storage"
	"github.com/ConvertAPI/convertapi-go"
	"github.com/ConvertAPI/convertapi-go/config"
	"github.com/ConvertAPI/convertapi-go/param"
	"path/filepath"
)

func Convert(studyGroupName string) {
	config.Default.Secret = "FclSU1D7eRCoQWqN"

	from := filepath.Join(schedules_storage.StoragePath, studyGroupName, "schedule.pdf")
	to := filepath.Join(schedules_storage.StoragePath, studyGroupName, "schedule.txt")

	convertapi.ConvDef("pdf", "txt",
		param.NewPath("File", from, nil),
	).ToPath(to)
}

func clear(studyGroupName string) {

}
