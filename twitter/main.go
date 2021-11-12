package twitter

import (
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetLevel(log.DebugLevel) // TODO: make this a flag
	log.Infof("log level: %v", log.GetLevel())
}
