package service

type Service struct {
	Name, Type, Url, Method string
	Expect int
}

type Checker func(Service) error

var checkers = make(map[string]Checker)

func (s Service) Check() error {
	if err := checkers[s.Type](s); err != nil {
		return err
	}
	return nil
}

func init() {
	checkers["http"] = httpCall
	checkers["mongodb"] = mongoCall
}
