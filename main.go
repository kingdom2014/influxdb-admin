package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"influx-admin/admin"
	"io/ioutil"
)

func main() {
	config := admin.NewConfig()

	yamlFile, err := ioutil.ReadFile("conf/admin.yaml")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		err = yaml.Unmarshal(yamlFile, &config)
		if err != nil {
			fmt.Println(err.Error())
		}
	}

	server := admin.NewService(config)
	defer server.Close()

	err = server.Open()
	if err != nil {
		fmt.Println(err.Error())
	}
}
