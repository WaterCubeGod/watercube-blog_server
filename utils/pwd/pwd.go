package pwd

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

// HashEncode 加密密码
/*
 * hash 加密密码
 * @param string pwd 待加密的明文密码
 */
func HashEncode(pwd string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

// ComparePasswords 验证密码
/*
 * 验证 hash 密码
 * @param string hashedPwd 已加密的hash密码
 * @param string sourcePwd 确认密码
 */
func ComparePasswords(hashedPwd string, sourcePwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(sourcePwd))
	if err != nil {
		return false
	}
	return true
}
