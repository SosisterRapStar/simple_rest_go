package tags

import "context"

type TagRepo interface {
	Save(ctx context.Context, tag string) (*Tag, error)
}

type Module struct {
	repo TagRepo
}

func New(r TagRepo) *Module {
	return &Module{repo: r}
}

func (m *Module) Create(ctx context.Context, name string) (*Tag, error) {
	t, err := m.repo.Save(ctx, name)
	// ошибки нужно будет оборачивать
	if err != nil {
		return nil, err
	}
	return t, nil
}
