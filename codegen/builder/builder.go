package builder

type Builder interface {
	Build() error
}

func Run(builders ...Builder) error {
	for _, b := range builders {
		if err := b.Build(); err != nil {
			return err
		}
	}
	return nil
}
