package mktools

import (
	"github.com/golang/mock/gomock"
	"github.com/nmarsollier/commongo/log"
	"github.com/nmarsollier/commongo/test/mockgen"
	"github.com/sirupsen/logrus"
)

func NewTestLogger(ctrl *gomock.Controller, withFieldCount int, errorCount int, infoCount int, dataCount int, warnCount int, fatalCount int) log.LogRusEntry {
	logger := mockgen.NewMockLogRusEntry(ctrl)
	logger.EXPECT().WithField(gomock.Any(), gomock.Any()).Return(logger).Times(withFieldCount)
	logger.EXPECT().Error(gomock.Any()).Return().Times(errorCount)
	logger.EXPECT().Info(gomock.Any()).Return().Times(infoCount)
	logger.EXPECT().Warn(gomock.Any()).Return().Times(warnCount)
	logger.EXPECT().Fatal(gomock.Any()).Return().Times(fatalCount)

	logger.EXPECT().Data().Return(logrus.Fields{
		log.LOG_FIELD_CORRELATION_ID: "correlationId",
	}).Times(dataCount)

	logger.EXPECT().CorrelationId().Return("correlationId").AnyTimes()

	return logger
}
