package schedule

import (
	"github.com/robfig/cron/v3"
	"log"
)

type Schedule struct {
	c *cron.Cron
}

func (s *Schedule) Init() {
	s.c = cron.New()
	entryId, err := s.c.AddFunc("@every 5s", func() { log.Println("Heartbeat") })
	if err != nil {
		log.Panicf("Can't add job by %d entry id", entryId)
	}
}

func (s *Schedule) Start() {
	s.c.Start()
}
