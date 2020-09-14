package orm

type ORM interface {
	GetModel() Model
	Create()
}
