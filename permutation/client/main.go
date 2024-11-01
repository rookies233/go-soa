package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"net/http"
)

type Request struct {
	XMLName xml.Name `xml:"Request"`
	Nums    []int    `xml:"Nums>Num"`
}

type Response struct {
	XMLName      xml.Name `xml:"Response"`
	Permutations [][]int  `xml:"Permutations>Permutation"`
}

func main() {
	// 创建请求数据
	nums := []int{1, 2, 3}
	req := Request{Nums: nums}

	// 编码请求为 XML
	var buf bytes.Buffer
	if err := xml.NewEncoder(&buf).Encode(req); err != nil {
		fmt.Println("Error encoding request:", err)
		return
	}

	// 发送 HTTP POST 请求
	resp, err := http.Post("http://localhost:8080/soap", "text/xml", &buf)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// 解析响应
	var response Response
	if err := xml.NewDecoder(resp.Body).Decode(&response); err != nil {
		fmt.Println("Error decoding response:", err)
		return
	}

	// 打印结果
	fmt.Println("Permutations:", response.Permutations)
}
