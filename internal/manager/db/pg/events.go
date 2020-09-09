package pg

func (d *Driver) CreateEvent() error {
	_, err := d.db.GetDatabase()
	return err
}
