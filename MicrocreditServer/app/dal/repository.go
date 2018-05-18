package dal


type Repository interface {
	List(dest interface{}, query string) error
	Get(ID int, dest interface{}, query string) error
	Create(data interface{}, query string) error
	Update(ID int, data interface{}, query string) error
	Delete(ID int, query string) error
}
