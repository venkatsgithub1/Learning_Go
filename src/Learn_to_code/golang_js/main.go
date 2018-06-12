package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"strings"
)

func getGravatarHash(email string) string {
	h := md5.New()
	email = strings.ToLower(strings.TrimSpace(email))
	io.WriteString(h, email)
	return hex.EncodeToString(h.Sum(nil))
}

func main() {
	gravatarHash := getGravatarHash("johndoe@xyz.com")
	fmt.Fprintln(os.Stderr, "Enter your name:")
	var name string
	fmt.Scanln(&name)

	fmt.Println(`<!DOCTYPE html>
		<html>
			<head>
			</head>
			<body>
				<img src="http://www.gravatar.com/avatar/` + gravatarHash + `?d=identicon">
			</body>
		</html>
		`)
}
