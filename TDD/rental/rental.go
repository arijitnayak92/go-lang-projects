package rental

type Rental struct {
	users map[string]struct{}
}

func NewRental() *Rental {
	return &Rental{
		users: make(map[string]struct{}),
	}
}

func (r *Rental) CheckUser(username string) bool {
	_, ok := r.users[username]
	return ok

}

func (r *Rental) AddUser(username string) {
	_, ok := r.users[username]
	if !ok {
		r.users[username] = struct{}{}
	}
}
