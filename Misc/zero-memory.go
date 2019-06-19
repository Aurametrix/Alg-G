// reset resets all fields of the AuthenticationResponse before pooling it.
func (a* AuthenticationResponse) reset() {
    a.Token = ""
    a.UserID = ""
}

rsp := authPool.Get().(*AuthenticationResponse)
defer func() {
    rsp.reset()
    authPool.Put(rsp)
}()
