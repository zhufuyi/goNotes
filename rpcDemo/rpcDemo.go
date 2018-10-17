package rpcDemo

import "errors"

const (
	Add = "DemoService.Add"
	Sub = "DemoService.Sub"
	Mul = "DemoService.Mul"
	Div = "DemoService.Div"
)

type DemoService struct{}

type Args struct {
	A, B int
}

func (DemoService) Add(args *Args, result *int) error {
	*result = args.A + args.B
	return nil
}

func (DemoService) Sub(args *Args, result *int) error {
	*result = args.A - args.B
	return nil
}

func (DemoService) Mul(args *Args, result *int) error {
	*result = args.A * args.B
	return nil
}

func (DemoService) Div(args *Args, result *float64) error {
	if args == nil {
		return errors.New("args is nil")
	}
	if args.B == 0 {
		return errors.New("division by 0")
	}

	*result = float64(args.A) / float64(args.B)
	return nil
}
