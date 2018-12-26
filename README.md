# utils
Utilities Library For the Habitat Project, it includes:

# Yaml Parser

A Yaml file Parser and Yaml File creator that creates an easy to use tree model out of .yaml file and allows you to create a tree model of yaml nodes that is persisted to .yaml file.

## Yaml Parser Example

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

SqlST - An slq.ST wrapper that makes it easy to create sql parepare statements without the need to manage the parameters numbers and types.



