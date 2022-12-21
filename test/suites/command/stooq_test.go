package test_command

import (
	"chat-jobsity/internal/command"
	mock_client "chat-jobsity/test/mock/client"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Stooq Command", func() {
	var (
		ctrl     *gomock.Controller
		stooqCli *mock_client.MockStooqClient
		cmd      command.Command
	)
	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		stooqCli = mock_client.NewMockStooqClient(ctrl)
		cmd = command.NewStooqCommand(stooqCli)
	})

	_ = Describe("Call run", func() {
		It("should fail, when passing wrong command code", func() {
			stooqCli.EXPECT().GetStockDetails(gomock.Any()).Return("mockMessage", nil)
			msg, err := cmd.Run("mockParam")
			Expect(err).Should(BeNil())
			Expect(msg).Should(Equal("mockMessage"))
		})
	})
})
