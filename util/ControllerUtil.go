package util

import (
	"fmt"
	"net/http"
	"reflect"
)

func FormToStruct(r *http.Request, rv reflect.Value) {

	v := reflect.Indirect(rv)
	vt := v.Type()
	for i := 0; i < vt.NumField(); i++ {
		field := vt.Field(i)
		if field.Type.Kind()==reflect.Struct{
			panic("cannot convert struct now")
		}else{
			formName := field.Tag.Get("json") // form name
			formValue := r.FormValue(formName)
			fieldVal:=v.Field(i)

			if fieldVal.CanSet(){
				fmt.Println(vt.Field(i).Name)
				fmt.Println("can set")
				fieldVal.SetString(formValue)
				fmt.Println("value",fieldVal.Interface())
			}
		}


	}

}
