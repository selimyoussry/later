package later

// auth authorizes or not the client with a simple token
// IMPLEMENT ME
func (server *Server) auth(token string) error {

	if server.SecretKey != token {
		return Err_WrongToken(token)
	}

	return nil

}
