package guipigCore

type Runnable interface {
	Run([]Schema)
}

type Migrator struct {
	DbType  string
	Migrate func(schemas []Schema)
}

func (m Migrator) Run(schemas []Schema) {
	m.Migrate(schemas)
}
