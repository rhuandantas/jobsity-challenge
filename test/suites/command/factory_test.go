package test_command

import (
	"chat-jobsity/internal/command"
	mock_command "chat-jobsity/test/mock/command"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Factory", func() {
	var (
		ctrl         *gomock.Controller
		stockCommand *mock_command.MockCommand
		cmdManager   command.Manager
	)
	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		stockCommand = mock_command.NewMockCommand(ctrl)
		cmdManager = command.NewCommandManager(stockCommand)
	})

	_ = Describe("Get Command", func() {
		It("should fail, when passing wrong command code", func() {
			cmd, err := cmdManager.GetCommand("mockCommand")
			Expect(err).ShouldNot(BeNil())
			Expect(cmd).Should(BeNil())
		})
		It("should be success, when passing valid command code", func() {
			cmd, err := cmdManager.GetCommand(command.Stock)
			Expect(err).Should(BeNil())
			Expect(cmd).ShouldNot(BeNil())
		})
	})
})
