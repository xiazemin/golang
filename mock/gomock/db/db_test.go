package db

import (
	"github.com/golang/mock/gomock"
	"github.com/xiazemin/mock/gomock/db/source"
	"testing"
)

/*func TestObjDemo(t *testing.T) {
	Convey("test obj demo", t, func() {
		Convey("create obj", func() {
			ctrl := NewController(t)
			defer ctrl.Finish()
			mockRepo := mock_db.NewMockRepository(ctrl)
			mockRepo.EXPECT().Retrieve(Any()).Return(nil, ErrAny)
			mockRepo.EXPECT().Create(Any(), Any()).Return(nil)
			mockRepo.EXPECT().Retrieve(Any()).Return(objBytes, nil)
			stubs := StubFunc(&redisrepo.GetInstance, mockRepo)
			defer stubs.Reset()
			...
		})

		Convey("bulk create objs", func() {
			ctrl := NewController(t)
			defer ctrl.Finish()
			mockRepo := mock_db.NewMockRepository(ctrl)
			mockRepo.EXPECT().Create(Any(), Any()).Return(nil).Times(5)
			stubs := StubFunc(&redisrepo.GetInstance, mockRepo)
			defer stubs.Reset()
			...
		})

		Convey("bulk retrieve objs", func() {
			ctrl := NewController(t)
			defer ctrl.Finish()
			mockRepo := mock_db.NewMockRepository(ctrl)
			objBytes1 := ...
			objBytes2 := ...
			objBytes3 := ...
			objBytes4 := ...
			objBytes5 := ...
			mockRepo.EXPECT().Retrieve(Any()).Return(objBytes1, nil)
			mockRepo.EXPECT().Retrieve(Any()).Return(objBytes2, nil)
			mockRepo.EXPECT().Retrieve(Any()).Return(objBytes3, nil)
			mockRepo.EXPECT().Retrieve(Any()).Return(objBytes4, nil)
			mockRepo.EXPECT().Retrieve(Any()).Return(objBytes5, nil)
			stubs := StubFunc(&redisrepo.GetInstance, mockRepo)
			defer stubs.Reset()
		})
	})
}
*/
func TestCreate(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	//go build github.com/xiazemin/mock/gomock/db/source: no non-test Go files in /Users/didi/goLang/src/github.com/xiazemin/mock/gomock/db/source

	mockCreate := source.NewMockRepository(ctl)
	mockCreate.EXPECT().Retrieve("key1").Return(nil, nil)
	mockCreate.EXPECT().Create("key1", "value1").Return(nil)
	mockCreate.EXPECT().Retrieve("key1").Return("value1", nil)
}
