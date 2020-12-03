package gowheel

import jsoniter "github.com/json-iterator/go"



var j jsoniter.API

func Jsoniter() jsoniter.API {
	return j
}

func init() {
	j = jsoniter.ConfigCompatibleWithStandardLibrary
}
