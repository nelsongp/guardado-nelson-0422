package factory

//Factorys sera el factory que utilizaremos para identificar el tipo de usuario a utilizar
func Factorys(user string) GetUserCodes {
	switch user {
	case "adminSecurity1":
		return &AdminUno{}
	case "userSecurity1":
		return &UserUno{}
	default:
		return nil
	}
}
