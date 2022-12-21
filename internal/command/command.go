package command

//go:generate mockgen -source=$GOFILE -package=mock_command -destination=../../test/mock/command/$GOFILE

type Command interface {
	Run(param string) (string, error)
}
