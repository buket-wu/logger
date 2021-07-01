package logger

import "github.com/sirupsen/logrus"

func Trace(msg ...interface{}) {
	if disableGoId {
		logrus.Trace(msg...)
		return
	}

	logrus.WithField("ctxId", getCtxId()).Trace(msg...)
}

func Debug(msg ...interface{}) {
	if disableGoId {
		logrus.Debug(msg...)
		return
	}

	logrus.WithField("ctxId", getCtxId()).Debug(msg...)
}

func Info(msg ...interface{}) {
	if disableGoId {
		logrus.Info(msg...)
		return
	}

	logrus.WithField("ctxId", getCtxId()).Info(msg...)
}

func Warn(msg ...interface{}) {
	if disableGoId {
		logrus.Trace(msg...)
		return
	}

	logrus.WithField("ctxId", getCtxId()).Trace(msg...)
}

func Error(msg ...interface{}) {
	if disableGoId {
		logrus.Error(msg...)
		return
	}

	logrus.WithField("ctxId", getCtxId()).Error(msg...)
}

func Fatal(msg ...interface{}) {
	if disableGoId {
		logrus.Fatal(msg...)
		return
	}

	logrus.WithField("ctxId", getCtxId()).Fatal(msg...)
}

func Panic(msg ...interface{}) {
	if disableGoId {
		logrus.Panic(msg...)
		return
	}

	logrus.WithField("ctxId", getCtxId()).Panic(msg...)
}

func Tracef(format string, args ...interface{}) {
	if disableGoId {
		logrus.Tracef(format, args...)
		return
	}

	logrus.WithField("ctxId", getCtxId()).Tracef(format, args...)
}

func Debugf(format string, args ...interface{}) {
	if disableGoId {
		logrus.Debugf(format, args...)
		return
	}

	logrus.WithField("ctxId", getCtxId()).Debugf(format, args...)
}

func Infof(format string, args ...interface{}) {
	if disableGoId {
		logrus.Infof(format, args...)
		return
	}

	logrus.WithField("ctxId", getCtxId()).Infof(format, args...)
}

func Warnf(format string, args ...interface{}) {
	if disableGoId {
		logrus.Warnf(format, args...)
		return
	}

	logrus.WithField("ctxId", getCtxId()).Warnf(format, args...)
}

func Errorf(format string, args ...interface{}) {
	if disableGoId {
		logrus.Errorf(format, args...)
		return
	}

	logrus.WithField("ctxId", getCtxId()).Errorf(format, args...)
}

func Fatalf(format string, args ...interface{}) {
	if disableGoId {
		logrus.Fatalf(format, args...)
		return
	}

	logrus.WithField("ctxId", getCtxId()).Fatalf(format, args...)
}

func Panicf(format string, args ...interface{}) {
	if disableGoId {
		logrus.Panicf(format, args...)
		return
	}

	logrus.WithField("ctxId", getCtxId()).Panicf(format, args...)
}
