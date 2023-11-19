package job

import (
	"github.com/engboustani/go-quartz/pkg/command"
	"log"
	"sync"
)

type Job struct {
	Entrypoint string
	User       string
	Commands   []command.Command
}

func (j *Job) ExecuteCommands() {
	if len(j.Commands) == 0 {
		log.Fatal("No command defined for job")
	}
	var waitGroup sync.WaitGroup
	waitGroup.Add(len(j.Commands))

	defer waitGroup.Wait()

	for _, com := range j.Commands {
		go func(c *command.Command) {
			defer waitGroup.Done()
			c.Execute()
		}(&com)
	}
}
