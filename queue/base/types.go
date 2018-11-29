package base

type Object interface {}

type MatchFun func(data1, data2 Object) int

func WhatType(value interface{}) string  {
    switch value.(type) {
    case int:
        return "int"
    case string:
        return "字符串"
    default:
        return "未知"
    }
}