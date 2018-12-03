package configlib

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
)

func UnMarshalFile(filename string, result interface{}) (err error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}

	return UnMarshal(data, result)
}

func UnMarshal(data []byte, result interface{}) (err error) {
	t := reflect.TypeOf(result)
	v := reflect.ValueOf(result)

	if t.Kind() != reflect.Ptr {
		panic("please input a address!")
	}

	var sectionName string
	lines := strings.Split(string(data), "\n")
	lineNo := 0

	for _, line := range lines {
		lineNo++
		line = strings.Trim(line, "\t\r\n")
		if len(line) == 0 {
			continue
		}

		if line[0] == '#' || line[0] == ';' {
			continue
		}

		if line[0] == '[' {
			if len(line) <= 2 || line[len(line)-1] != ']' {
				tips := fmt.Sprintf("syntax error: invalid section: \"%s\" line: %d", line, lineNo)
				panic(tips)
			}

			sectionName = strings.TrimSpace(line[1 : len(line)-1])
			if len(sectionName) == 0 {
				tips := fmt.Sprintf("syntax error: invalid section: \"%s\" line: %d", line, lineNo)
				panic(tips)
			}

			fmt.Printf("section: %s\n", sectionName)
		} else {
			if len(sectionName) == 0 {
				tips := fmt.Sprintf("syntax error: invalid line: %s, lineNo: %d, not found section", line, lineNo)
				panic(tips)
			}

			equalIndex := strings.Index(line, "=")
			if equalIndex == -1 {
				tips := fmt.Sprintf("syntax error: invalid line: %s, lineNo: %d, not found \"=\"", line, lineNo)
				panic(tips)
			}

			// 取得ini 配置文件的key 和value 值，后续对传入的 Conf 结构体对象，
			// 进行reflect 取出信息后对照进行Conf 结构体的各个字段的赋值
			key := strings.TrimSpace(line[0:equalIndex])
			value := strings.TrimSpace(line[equalIndex+1:])

			if len(key) == 0 {
				tips := fmt.Sprintf("syntax error: invalid line: %s, lineNo: %d, key is empty", line, lineNo)
				panic(tips)
			}

			for i := 0; i < t.Elem().NumField(); i++ {
				// 获取嵌入的Conf 结构体的类型的信息
				tfs := t.Elem().Field(i)
				// 获取嵌入的Conf 结构体的值的信息
				vfs := v.Elem().Field(i)

				// 获取嵌入 Conf 结构体的类型的信息，取得他的tag 标记
				if tfs.Tag.Get("ini") != sectionName {
					continue
				}

				// 获取Conf 结构体的 嵌入结构体的类型
				tfsType := tfs.Type
				if tfsType.Kind() != reflect.Struct {
					tips := fmt.Sprintf("syntax error: filed %s is not struct", tfsType.Name())
					panic(tips)
				}

				// 进入各个嵌入的Conf 结构体的 内部的Filed 字段，进行各个字段的赋值
				for j := 0; j < tfsType.NumField(); j++ {
					// 获取各个字段的key的类型信息
					tfsKeyf := tfsType.Field(j)
					// 获取各个字段的value的值的信息
					vfsValuef := vfs.Field(j)

					// 获取各个字段的key 的 tag 的类型信息
					if tfsKeyf.Tag.Get("ini") != key {
						continue
					}

					switch tfsKeyf.Type.Kind() {
					case reflect.String:
						vfsValuef.SetString(value)
					case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
						fallthrough
					case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
						valueInt, err := strconv.ParseInt(value, 10, 64)
						if err != nil {
							tips := fmt.Sprintf("value: %s can not convert to int64, lineNo: %d\n", value, lineNo)
							panic(tips)
						}
						vfsValuef.SetInt(valueInt)
					case reflect.Float32, reflect.Float64:
						valueFloat, err := strconv.ParseFloat(value, 64)
						if err != nil {
							tips := fmt.Sprintf("value: %s can not convert to float64, lineNo: %d\n", value, lineNo)
							panic(tips)
						}
						vfsValuef.SetFloat(valueFloat)
					default:
						tips := fmt.Sprintf("key: \"%s\" can not convert to %v, lineNo: %d\n", key, tfsKeyf.Type.Kind(), lineNo)
						panic(tips)
					}
				}
				// 这里的break 的含义是，代码走到这里的话说明已经找到了相应的sectionName 以及该 section 下的key，此为唯一，不必再进行下面的循环了
				break
			}
		}
	}
	return
}

func MarshalFile(filename string, result interface{}) (err error) {
	data, err := Marshal(result)
	if err != nil {
		return
	}

	err = ioutil.WriteFile(filename, data, 0755)
	return
}

func Marshal(result interface{}) (data []byte, err error) {
	t := reflect.TypeOf(result)
	v := reflect.ValueOf(result)

	if t.Kind() != reflect.Struct {
		panic("please input struct type")
		return
	}

	var strSlice []string
	for i := 0; i < t.NumField(); i++ {
		// 取得嵌入Conf内部的结构体的类型的信息
		tf := t.Field(i)
		vf := v.Field(i)

		// 判断内嵌的类型是否为结构体
		if tf.Type.Kind() != reflect.Struct {
			continue
		}

		sectionName := tf.Name
		if len(tf.Tag.Get("ini")) > 0 {
			sectionName = tf.Tag.Get("ini")
		}

		sectionName = fmt.Sprintf("[%s]\n", sectionName)
		strSlice = append(strSlice, sectionName)

		for j := 0; j < tf.Type.NumField(); j++ {
			subTf := tf.Type.Field(j)
			if subTf.Type.Kind() == reflect.Struct || subTf.Type.Kind() == reflect.Ptr {
				// 跳过内嵌结构体包含结构体字段的情况
				continue
			}

			// 获取的是结构体的字段的名字
			subTfName := subTf.Name
			// 如果结构体存在Tag 的标记，则获取
			subTfTagName := subTf.Tag.Get("ini")

			if len(subTfTagName) > 0 {
				subTfName = subTfTagName
			}

			subVf := vf.Field(j)
			filedStr := fmt.Sprintf("%s=%v\n", subTfName, subVf.Interface())
			fmt.Printf("config: %s", filedStr)

			strSlice = append(strSlice, filedStr)
		}
	}

	for _, v := range strSlice {
		data = append(data, []byte(v)...)
	}

	return
}
