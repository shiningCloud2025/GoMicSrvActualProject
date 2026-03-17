package main

import (
	"crypto/md5"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"io"
	"strings"

	"github.com/anaskhan96/go-password-encoder"
)

func genMd5(code string) string {
	Md5 := md5.New()
	_, _ = io.WriteString(Md5, code)
	return hex.EncodeToString(Md5.Sum(nil))
}

func main() {
	//// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	//dsn := "root:mysql~!@#$%^&*()_+@tcp(117.50.184.138:37210)/zyh_dev_roundshop?charset=utf8mb4&parseTime=True&loc=Local"
	//
	//newLogger := logger.New(
	//	log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
	//	logger.Config{
	//		SlowThreshold:             time.Second, // Slow SQL threshold
	//		LogLevel:                  logger.Info, // Log level
	//		IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
	//		ParameterizedQueries:      true,        // Don't include params in the SQL log
	//		Colorful:                  true,        // Disable color
	//	},
	//)
	//
	//// Globally mode
	//db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
	//	NamingStrategy: schema.NamingStrategy{
	//		SingularTable: true,
	//	},
	//	Logger: newLogger,
	//})
	//if err != nil {
	//	panic(err)
	//}
	//
	//db.AutoMigrate(&model.User{})
	// Using the default options
	//salt, encodedPwd := password.Encode("generic password", nil)
	//fmt.Println("salt is", salt)
	//fmt.Println("encodedPwd", encodedPwd)
	//check := password.Verify("generic password", salt, encodedPwd, nil)
	//fmt.Println(check) // true

	//// Using custom options
	options := &password.Options{16, 100, 32, sha512.New}
	salt, encodedPwd := password.Encode("generic password", options)
	newPassword := fmt.Sprintf("$pbdkf2-sha512$%s$%s", salt, encodedPwd)
	fmt.Println(len(newPassword))
	fmt.Println(newPassword)

	passwordInfo := strings.Split(newPassword, "$")
	fmt.Println(passwordInfo)
	fmt.Println("Encoded password:", encodedPwd)
	fmt.Println("Salt:", salt)
	check := password.Verify("generic password", passwordInfo[2], passwordInfo[3], options)
	fmt.Println(check) // true
	fmt.Println(genMd5("123456"))

	// 将用户密码变一下 随机字符串+用户密码
	// https://toolshu.com/crackmd5":暴力破解，把一些常见值放彩虹表 会去核对  可以通过加盐值去解决这个问题
	// e10adc3949ba59abbe56e057f20f883e
	// e10adc3949ba59abbe56e057f20f883e

}
