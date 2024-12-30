package mktools

import (
	gomock "github.com/golang/mock/gomock"
	"github.com/nmarsollier/commongo/test/mockgen"
)

func ExpectFindOneError(coll *mockgen.MockCollection, err error, times int) {
	coll.EXPECT().FindOne(gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(
		func(arg1 interface{}, params interface{}, update interface{}) error {
			return err
		},
	).Times(times)
}

func ExpectUpdateOneError(coll *mockgen.MockCollection, err error, times int) {
	coll.EXPECT().UpdateOne(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(int64(0), err).Times(times)
}

func ExpectInsertOneError(coll *mockgen.MockCollection, err error, times int) {
	coll.EXPECT().InsertOne(gomock.Any(), gomock.Any()).Return("", err).Times(times)
}

func ExpectInsertOne(coll *mockgen.MockCollection, times int) {
	coll.EXPECT().InsertOne(gomock.Any(), gomock.Any()).Return("", nil).Times(times)
}

func ExpectFindOne[T any](coll *mockgen.MockCollection, userData *T, times int) {
	coll.EXPECT().FindOne(gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(
		func(arg1 interface{}, params interface{}, update *T) error {
			*update = *userData
			return nil
		},
	).Times(times)
}
