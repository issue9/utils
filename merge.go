// SPDX-License-Identifier: MIT

package utils

import "reflect"

// Merge 合并所有的元素
//
// 后一个元素的非零值将取代前一个元素中的值。
// deep 是否递归合并子元素。
// 合并的元素只支持结构体或是结构体指针。
func Merge(deep bool, elems ...interface{}) error {
	if len(elems) < 2 {
		panic("参数 elems 数量必须大于 1")
	}

	vals := make([]reflect.Value, len(elems), len(elems))
	var typ reflect.Type
	for index, elem := range elems {
		val := reflect.ValueOf(elem)
		for val.Kind() == reflect.Ptr {
			val = val.Elem()
		}

		// 类型检测
		if index == 0 {
			typ = val.Type()
		} else if typ != val.Type() {
			panic("参数拥有不同的类型")
		}

		vals[index] = val
	}

	for i := 1; i < len(vals); i++ {
		if err := merge(deep, vals[0], vals[i]); err != nil {
			return err
		}
	}
	return nil
}

// 将 v1, v2 合并，其中 v2 中的非零值将覆盖 v1 中对应的值。
// deep 是否将合并功能作用于成员变量。
// 确保 v1.Type()==v2.Type()
// 若是 map，则当作普通的成员变量，直接赋值。
func merge(deep bool, v1, v2 reflect.Value) error {
	if !v1.IsValid() || !v2.IsValid() {
		return nil
	}

	num := v2.NumField()
	for i := 0; i < num; i++ {
		var err error
		switch {
		case v1.Type().Field(i).Anonymous: // 匿名
			err = merge(deep, v1.Field(i), v2.Field(i))
		case v1.Field(i).Kind() == reflect.Ptr:
			if !deep {
				v1.Field(i).Elem().Set(v2.Field(i).Elem())
				continue
			}

			if v1.Field(i).IsNil() { // v1 为 nil 时，需要先初始化该结构体
				if v2.Field(i).IsNil() { // 双方均为nil时，无须任何操作
					continue
				}
				typ := v1.Field(i).Type().Elem()
				v1.Field(i).Set(reflect.New(typ))
			}
			err = merge(true, v1.Field(i).Elem(), v2.Field(i).Elem())
		case v1.Field(i).Kind() == reflect.Struct: // 嵌套
			if !deep {
				v1.Field(i).Set(v2.Field(i))
				continue
			}
			err = merge(true, v1.Field(i), v2.Field(i))
		default:
			if !v1.Field(i).CanSet() { // 过滤不可导出字段
				continue
			}

			k := v2.Field(i).Kind()
			if k == reflect.Slice || k == reflect.Map || k == reflect.Array {
				if v2.Field(i).Len() > 0 {
					v1.Field(i).Set(v2.Field(i))
				}
				continue

			}

			// v2 若是零值，则不合并
			if v2.Field(i).Interface() == reflect.Zero(v2.Field(i).Type()).Interface() {
				continue
			}
			v1.Field(i).Set(v2.Field(i))
		}

		if err != nil {
			return err
		}
	}

	return nil
}
