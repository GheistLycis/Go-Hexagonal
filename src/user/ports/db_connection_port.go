package user

type DBConnectionPort interface {
	// Get saves into destiny the first entry in DB that matches conditions
	Get(destiny any, conditions ...any) error

	// Insert adds the given model into its respective table
	Insert(value any) error

	// List saves into destiny all entries in DB that match conditions
	List(destiny any, conditions ...any) error

	// Raw executes raw SQL query with given queryArgs. If destiny is provided, it saves the result into it
	Raw(destiny any, query string, queryArgs ...any) error

	// Update updates the given model. If entry is not found in its respective table, it returns an error
	Update(value any) error

	// Upsert updates the given model. If entry is not found in its respective table, it is created instead
	Upsert(value any) error
}
