package railroad

type Result struct {
	State   interface{}
	Success bool
	Failure bool
	Err     error
}

type Applier func(Result) Result

func (r Result) Apply(app Applier) Result {
	if r.Failure {
		return r
	}

	return app(r)
}
