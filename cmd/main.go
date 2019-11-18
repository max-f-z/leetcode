package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println(isSubsequence("axc", "ahbgdc"))

	data := make(map[string][]string, 3)
	data["os_username"] = []string{"zhangfuqing"}
	data["os_password"] = []string{"Polaris123"}

	resp, err := http.PostForm("https://jira.devops.viewchain.net/login.jsp", data)

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(resp)

}
