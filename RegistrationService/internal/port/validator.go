package port

type Validator interface {
	Struct(s any) error // any = interface{}
}