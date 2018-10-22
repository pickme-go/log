package log

import "gitlab.mytaxi.lk/pickme/go/log"

type OtherLogger struct {
	log log.PrefixedLogger
}

func NewOtherLogger() OtherLogger  {
	return OtherLogger{
		log: log.Constructor.PrefixedLog(
			log.FileDepth(3),
			log.Prefixed(`otherlog`),
			),
	}
}

func (l *OtherLogger) Info(v ...interface{})  {
	l.log.Info(`test_prefix`, v)
}

func (l *OtherLogger) Trace(v ...interface{})  {
	l.log.Trace(`test_prefix`, v)
}
