package factory

import (
	"reflect"
	"testing"
)

func TestFactory(t *testing.T) {
	test := []struct {
		nameTest  string
		validUser string
		valimpl   GetUserCodes
		want      GetUserCodes
	}{
		{nameTest: "TestFactoryAdmin", validUser: "adminSecurity1", valimpl: &AdminUno{}, want: &AdminUno{}},
		{nameTest: "TestFactoryUser", validUser: "userSecurity1", valimpl: &UserUno{}, want: &UserUno{}},
	}
	for _, tt := range test {
		t.Run(tt.nameTest, func(t *testing.T) {
			if got := Factorys(tt.validUser); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("cannotparse() = %v, want %v", got, tt.want)
			}
		})
	}
}
