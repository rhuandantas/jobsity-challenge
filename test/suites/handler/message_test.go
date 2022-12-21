package test_handler

import (
	"chat-jobsity/internal/handler"
	mock_service "chat-jobsity/test/mock/service"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Stooq Command", func() {
	var (
		ctrl       *gomock.Controller
		msgMgr     *mock_service.MockMessageManager
		msgHandler handler.MessageHandler
	)
	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		msgMgr = mock_service.NewMockMessageManager(ctrl)
		msgHandler = handler.NewMessageHandler(msgMgr)
	})

	_ = Describe("Call run", func() {
		It("should fail, when passing wrong command code", func() {
			msgMgr.EXPECT().ManageMessage(gomock.Any()).Return("mockMessage", nil)
			msg := "{\"text\":\"mockTest\"}"
			msg, err := msgHandler.HandleMessage([]byte(msg))
			Expect(err).Should(BeNil())
			Expect(msg).Should(Equal("mockMessage"))
		})
	})
})
