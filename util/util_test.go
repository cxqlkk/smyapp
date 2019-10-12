package util

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

func TestFormToStruct(t *testing.T) {
	var x float64 = 3.4
	p := reflect.ValueOf(&x) // 注意：得到X的地址
	fmt.Println("type of p:", p.Type())
	fmt.Println("settability of p:" , p.CanSet())
	v := p.Elem()
	fmt.Println("settability of v:" , v.CanSet())
	v.SetFloat(7.1)
	fmt.Println(v.Interface())
	fmt.Println(x)
}
type personArr []struct{
	Address string `json:"address"`
	UserName string `json:"username"`
}

func TestOnce(t *testing.T){
	str:="[{\"address\":\"北京\",\"username\":\"kongyixueyuan\"}]"
	var p personArr
	json.Unmarshal([]byte(str),&p)
	for _,m:=range p{
		fmt.Println(m.Address,m.UserName)
	}
}
