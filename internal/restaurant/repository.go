package restaurant

type Repository interface {
	GetRestaurants() ([]*Restaurant, error)
}
