package controller

import (
	"fmt"
	"go.uber.org/zap"
	"net/http"
	"reflect"
	"sync"
)

type webRouter struct {
	serveMux *http.ServeMux
	midWare  []MidWare
}

//中间件
type MidWare func(handle http.Handler) http.Handler

func (r *webRouter) Use(mid MidWare) {
	r.midWare = append(r.midWare, mid)
}
func (r *webRouter) AddRouter(url string, handle http.Handler) {
	var finalHandle = handle
	for i := len(r.midWare) - 1; i >= 0; i-- {
		finalHandle = r.midWare[i](finalHandle)
	}
	r.serveMux.Handle(url, finalHandle)
}

//global webrouter
var DefaultWebRouter *webRouter = NewRouter(nil)

func NewRouter(serveMux *http.ServeMux) *webRouter {
	var webR *webRouter
	one := &sync.Once{}
	one.Do(func() {
		if serveMux == nil {
			serveMux = http.DefaultServeMux
		}
		webR = &webRouter{
			serveMux: serveMux,
		}
	})
	return webR
}
func (wr *webRouter) Run(addr string) {

	http.ListenAndServe(addr, wr.serveMux)
}

func RegisterRouter(value reflect.Value) {
	//methods:=value.NumMethod()
	struPath := (value.Elem().Type().PkgPath())
	//fmt.Println(struPath)
	for i := 0; i < value.NumMethod(); i++ {
		//fmt.Println(i)

		method := value.Method(i)
		methodName := value.Type().Method(i).Name //important
		if method.Type().NumIn() == 2 {
			resp := reflect.TypeOf((*http.ResponseWriter)(nil)).Elem()
			req := reflect.TypeOf((*http.Request)(nil)) //.Elem() 就会判断不相等 @see 1
			argA := method.Type().In(0)
			argB := method.Type().In(1)
			//	fmt.Println(argA == resp)
			//	fmt.Println(argB == req) //@1
			if argA == resp && argB == req {
				fmt.Println(struPath + "/" + methodName)
				DefaultWebRouter.AddRouter("/"+struPath+"/"+methodName, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					args := []reflect.Value{reflect.ValueOf(w), reflect.ValueOf(r)}
					method.Call(args)
					//fmt.Println("addHandle")
				}))
			}
		}
	}
	/*if method.Type().In(0).Name()=="ResponseWriter"{
		fmt.Println(method.Type().In(0).Name())
	}*/

}
//deprected  原本是 将 form 参数转换为指定结构体 暂时废弃
func structParse(r *http.Request, v *interface{}) {

	tv := reflect.Indirect(reflect.ValueOf(v))
	tp := tv.Type()
	fmt.Println(tp.NumField())
	for i := 0; i < tp.NumField(); i++ {
		field := tp.Field(i)
		fmt.Println(field.Name, field.Type.Kind(), field.Type.Name())
	/*	fieldFormName := field.Tag.Get("json")
		vfield := tv.Field(i) //值
		if field.Type.Kind() != reflect.Struct{//非结构体
			if field.Type.Name()=="int"{

			}else if field.Type.Name()=="Time"{}
		}*/
		//fmt.Println(field.Type.Implements(t2))
	}
}

func  controllerCheckError(err error)  {
	if err!=nil{
		logger,_:=zap.NewProduction()
		defer logger.Sync()
		logger.Error(err.Error())
	}

}