package repository

// Repo repository
type Repo struct {
	conn string
}

// NewRepository constructor
func NewRepository(conn string) *Repo {
	return &Repo{
		conn: conn,
	}
}
