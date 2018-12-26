package examples

import (
	. "github.com/saichler/utils/golang"
	"io/ioutil"
	"os"
)

func Parse_yaml_file(yamlfile string) (*YamlNode, error){
	_,e := os.Stat(yamlfile)
	if e!=nil {
		return nil,e
	}

	buff,e := ioutil.ReadFile(yamlfile)
	if e!=nil {
		return nil,e
	}

	data := string(buff)

	modelRoot:=NewYamlRoot()

	Parse(data,yamlfile,modelRoot)

	return modelRoot,nil
}

func Create_yaml_file(yamlfile string) {
	modelRoot:=NewYamlRoot()

	child:=&YamlNode{}
	child.Init("tag","value",modelRoot.GetLvl()+1)
	modelRoot.AddChild(child)
	ioutil.WriteFile(yamlfile,[]byte(modelRoot.String()),0777)
}
