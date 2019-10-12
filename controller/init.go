package controller

import (
	"reflect"
)

func Manual(){
	RegisterRouter(reflect.ValueOf(&studentController{}))
	RegisterRouter(reflect.ValueOf(&honorController{}))
	RegisterRouter(reflect.ValueOf(&pictureController{}))
	RegisterRouter(reflect.ValueOf(&teacherController{}))
}