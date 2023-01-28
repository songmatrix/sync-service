package database

var _ driver = (*mockDriver)(nil)

type mockDriver struct{}

func (d *mockDriver) connect() error      { return nil }
func (d *mockDriver) assertSchema() error { return nil }
func (d *mockDriver) Close()              {}
