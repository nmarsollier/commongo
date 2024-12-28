package mktools

import (
	"github.com/golang/mock/gomock"
	"github.com/nmarsollier/commongo/test/mockgen"
)

func DefaultMockRabbitChannel(ctrl *gomock.Controller, times int) *mockgen.MockRabbitChannel {
	channel := mockgen.NewMockRabbitChannel(ctrl)
	channel.EXPECT().ExchangeDeclare(gomock.Any(), gomock.Any()).Return(nil).Times(times)
	return channel
}
