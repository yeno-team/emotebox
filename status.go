package emotebox

// Status Object
// @Description Default Status Object for presentation
type Status struct {
	Message string `json:"message" minLength:"4" maxLength:"20" example:"OK"`
	Code    int    `json:"code" example:"200"`
}

type StatusService interface {
	Status() (*Status, error)
}
