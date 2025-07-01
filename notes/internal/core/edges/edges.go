package edges

type EdgeRepo interface {
}

type Module struct {
	repo EdgeRepo
}

func New(r EdgeRepo) *Module {
	return &Module{repo: r}
}

func (m *Module) GetGraph() {

}

func (m *Module) CreateEdge() {

}
