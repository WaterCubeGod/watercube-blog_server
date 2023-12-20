package user_ser

import (
	"fmt"
	"gvb_server/global"
	"gvb_server/utils/jwts"
	"time"
)

type UserService struct {
}

func (*UserService) Logout(claims *jwts.CustomClaims, token string) error {
	exp := claims.ExpiresAt
	now := time.Now()

	diff := exp.Time.Sub(now)
	err := global.RDB.Set(fmt.Sprintf("logout_%s", token), "", diff).Err()
	return err
}
