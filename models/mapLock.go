package models

type MapArrayElemType string

const (
	ArrayElemTypeCommon MapArrayElemType = ""       //  默认，数组元素普通值类型，直接替换
	ArrayElemTypeInc    MapArrayElemType = "inc"    // 数组元素自增类型
	ArrayElemTypeDec    MapArrayElemType = "dec"    // 数组元素自减类型
	ArrayElemTypeAcc    MapArrayElemType = "acc"    // 数组元素累加类型
	ArrayElemTypeSubAcc MapArrayElemType = "subacc" // 数组元素累减类型
	ArrayElemTypeMax    MapArrayElemType = "max"    // 数组元素最大值类型
	ArrayElemTypeMin    MapArrayElemType = "min"    // 数组元素最小值类型
)
