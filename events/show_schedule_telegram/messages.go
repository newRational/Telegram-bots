package show_schedule_telegram

const msgHelp = `I'm a schedule-shower bot. Just send me a first letters
of name of the day (or the word "today") on which you want to get schedule`

const msgHello = "Hi there! \n\n" + msgHelp

const (
	msgUnknownCommand           = "Unknown command"
	msgNotSetStudyGroup         = "First of all send me the number of claimed study group"
	msgSaved                    = "Saved!"
	msgNotScheduleOfPickedGroup = "I'm sorry, but I don't have claimed study group schedule"
)
