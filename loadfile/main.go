package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// 统计文本行数

func main()  {
	file,err := os.Open("myfile.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	count := 0
	for scanner.Scan() {
		log.Println(scanner.Text())
		count ++
	}
	fmt.Println("一共有",count,"行")
}