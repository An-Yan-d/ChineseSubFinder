//go:build windows

package extend_log

import (
	"github.com/An-Yan-d/ChineseSubFinder/pkg/settings"
	"github.com/sirupsen/logrus"
)

type ExtendLog struct {
}

func (e *ExtendLog) AddHook(log *logrus.Logger, extendLog settings.ExtendLog) {

	if extendLog.SysLog.Enable == true {
		log.Infoln("Skip Add Syslog Hook, Windows Not Support it !")
	}
}
