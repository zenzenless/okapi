package types

import (
	"fmt"
	"reflect"
	"strings"
)

func genTypeComment(T interface{}, comment string) {
	rt := reflect.TypeOf(T)
	numField := rt.NumField()
	comments := strings.Split(comment, "\n")
	fmt.Println("comments len:", len(comments))
	find := func(s string) string {
		for _, v := range comments {
			if strings.Contains(v, s) {
				return v
			}
		}
		return ""
	}
	str := fmt.Sprintf("type %s struct{\n", rt.Name())
	for i := 0; i < numField; i++ {
		f := rt.Field(i)
		name := f.Name
		tag := f.Tag.Get("json")
		c := find(tag)
		field := fmt.Sprintf("%s\n%s\t%s\t `json:\"%s\"`", c, name, f.Type, tag)
		//fmt.Println(field)
		str = str + "\n" + field
	}
	fmt.Printf("%s\n}\n", str)
}
