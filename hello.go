package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "net/url"
)

func main() {
    urlValues := url.Values{}
    urlValues.Add("name","zhaofan")
    urlValues.Add("age","223")
    resp, _ := http.PostForm("https://www.baidu.com",urlValues)
    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println(string(body))
}
