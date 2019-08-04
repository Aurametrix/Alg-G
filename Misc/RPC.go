// RPCServer ...
type RPCServer struct {
	addr string
	funcs map[string] reflect.Value
}

// Register the name of the function and its entries
func (s *RPCServer) Register(fnName string, fFunc interface{}) {
	if _,ok := s.funcs[fnName]; ok {
		return
	}

	s.funcs[fnName] = reflect.ValueOf(fFunc)
}
