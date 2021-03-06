package commands

import (
	"cf/api"
	"cf/requirements"
	term "cf/terminal"
	"errors"
	"github.com/codegangsta/cli"
)

type RecentLogs struct {
	ui       term.UI
	appRepo  api.ApplicationRepository
	appReq   requirements.ApplicationRequirement
	logsRepo api.LogsRepository
}

func NewRecentLogs(ui term.UI, aR api.ApplicationRepository, lR api.LogsRepository) (l *RecentLogs) {
	l = new(RecentLogs)
	l.ui = ui
	l.appRepo = aR
	l.logsRepo = lR
	return
}

func (l *RecentLogs) GetRequirements(reqFactory requirements.Factory, c *cli.Context) (reqs []requirements.Requirement, err error) {
	if len(c.Args()) == 0 {
		err = errors.New("Incorrect Usage")
		l.ui.FailWithUsage(c, "logs")
		return
	}

	l.appReq = reqFactory.NewApplicationRequirement(c.Args()[0])
	reqs = []requirements.Requirement{l.appReq}
	return
}

func (l *RecentLogs) Run(c *cli.Context) {
	app := l.appReq.GetApplication()
	logs, err := l.logsRepo.RecentLogsFor(app)

	if err != nil {
		l.ui.Failed(err.Error())
		return
	}

	for _, log := range logs {
		l.ui.Say(string(log.GetMessage()))
	}
}
