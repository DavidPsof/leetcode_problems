package main

import (
	"flag"
	"fmt"
	"github.com/DavidPsof/leetcode_problems/backend/config"
	applog "github.com/DavidPsof/leetcode_problems/backend/log"
	"github.com/DavidPsof/leetcode_problems/backend/server"
	log "github.com/sirupsen/logrus"
	"runtime"
	"strings"
)

func main() {
	defer recovery()

	configPathParam := flag.String("config", "conf.env", "path to config file")
	flag.Parse()

	configPath := strings.TrimSpace(*configPathParam)

	conf := config.InitConfig(configPath)

	applog.InitLog(conf.LogSettings)

	s, err := server.NewServer(*conf)
	if err != nil {
		log.Error(err)
	}

	s.Run()
}

func recovery() {
	r := recover()
	if r == nil {
		return
	}

	log.Fatal(fmt.Errorf("PANIC:'%v'\nRecovered in: %s", r, IdentifyPanic()))
}

func IdentifyPanic() string {
	var name, file string
	var line int
	var pc [16]uintptr

	n := runtime.Callers(3, pc[:])
	for _, pc := range pc[:n] {
		fn := runtime.FuncForPC(pc)
		if fn == nil {
			continue
		}
		file, line = fn.FileLine(pc)
		name = fn.Name()
		if !strings.HasPrefix(name, "runtime.") {
			break
		}
	}

	switch {
	case name != "":
		return fmt.Sprintf("%v:%v", name, line)
	case file != "":
		return fmt.Sprintf("%v:%v", file, line)
	}

	return fmt.Sprintf("pc:%x", pc)
}
