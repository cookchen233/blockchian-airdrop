package utility

type GenericError interface {

}
type ParamError struct {
	Err error
}
func (err ParamError) Error() string{
	return err.Err.Error()
}

type NotAllowedError struct {
	Err error
}
func (err NotAllowedError) Error() string{
	return err.Err.Error()
}

type FatalError struct {
	Err error
}
func (err FatalError) Error() string{
	return err.Err.Error()
}

type UnexpectedError struct {
	Err error
}
func (err UnexpectedError) Error() string{
	return err.Err.Error()
}

type NoLoginError struct {
	Err error
}
func (err NoLoginError) Error() string{
	return err.Err.Error()
}
