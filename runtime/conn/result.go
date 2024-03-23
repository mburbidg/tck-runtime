package conn

type Result interface {
	result()
}

type ValueResult struct {
	Value any
}

func (result *ValueResult) result() {}

type BindingTableResult struct {
}

func (result *BindingTableResult) result() {}

type OmittedResult struct {
}

func (result *OmittedResult) result() {}
