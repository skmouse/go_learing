package go_learn

import (
	"fmt"
	"testing"
)

//func TestHelloWorld(t *testing.T) {
//	t.Log("helle world")
//}
//
//func TestVar(t *testing.T) {
//	//var a =1
//	//var b =2
//	//t.Log(&a)
//	//t.Log(&b)
//	//t.Log(Max(1, 2))
//	//t.Log(Swap("a", "b"))
//	//outPut("test11111")
//
//
//}

func Max(a int, b int) int {

	var result int
	if a > b {
		result = a
	} else {
		result = b
	}

	return result
}

func Swap(x, y string) (string, string) {

	return y, x
}

func outPut(name string) {
	fmt.Println(name)
}

//func TestArray(t *testing.T)  {
//	var arr  = [5] int{1,2,3,4}
//	for i:=0; i<len(arr); i++ {
//		t.Log(arr[i])
//	}
//
//	for k,v :=range arr {
//		t.Log(k, v, arr[k])
//	}
//
//}

//func TestSlice(t *testing.T)  {
//	slice := []int{1,2,3,4,5}
//	t.Log(slice, len(slice), cap(slice))
//	newSlice := append(slice,10)
//	t.Log(newSlice, len(newSlice), cap(newSlice))
//}
//
//func doSomeThing(s1 []int )  {
//	if len(s1 ) > 0 {
//		s1[0] = 1
//	}
//
//}
//
//func TestRange(t *testing.T)  {
//	//kvs := map[string]string {"a": "apple", "b": "banana"}
//	//
//	//for k,v := range kvs {
//	//
//	//	t.Log(k,v)
//	//}
//
//	m :=make(map[int] string)
//	m[0] = "a"
//	m[1] = "b"
//	m[2] = "c"
//	m[3] = "d"
//
//	t.Log(m)
//	for k,v := range m{
//		t.Log(k,v)
//	}
//
//	t.Log(len(m))
//
//}

//func TestMap(t *testing.T) {
//
//	countryCapitalMap:= map[string]string{}
//	countryCapitalMap [ "France" ] = "巴黎"
//	countryCapitalMap [ "Italy" ] = "罗马"
//	countryCapitalMap [ "Japan" ] = "东京"
//	countryCapitalMap [ "India " ] = "新德里"
//
//	//t.Log(countryCapitalMap,len(countryCapitalMap))
//	//
//	//for k,v := range countryCapitalMap{
//	//
//	//	t.Log(k,v,countryCapitalMap[k])
//	//
//	//}
//}
//
//type Phone interface {
//	call()
//}
//
//type NokiaPhone struct {
//}
//
//
//type IPhone struct {
//}
//
//
//func (nokiaPhone NokiaPhone) call() {
//	fmt.Println("I am Nokia, I can call you!")
//}
//
//func (iPhone  IPhone) call() {
//	fmt.Println("I am iPhone, I can call you!")
//}
//
//func TestInterface(t *testing.T) {
//
//	var phone Phone
//	phone = new(NokiaPhone)
//	phone.call()
//
//	phone = new(IPhone)
//
//	phone.call()
//}

func TestVar(t *testing.T) {
	////ages:=[]string{"10", "20", "30"}
	//
	//ages := map[string] int{"张三":15,"李四":20,"王武":36}
	//for k, v := range ages {
	//	t.Log(k, v, ages[k])
	//}
	//
	//u := User{"张三",20}
	//t :=reflect.TypeOf(u)
	//
	//t.Log(t)

}

type User struct {
	Name string
	Age  int
}

func TestFF(t *testing.T) {
	//u := User{"fengwenen",20}
	//(&u).modify()
	//fmt.Println(u.string())
	//
	//var b bytes.Buffer
	//fmt.Fprint(&b,"Hello World")
	//fmt.Println(b.String())

	//var l like
	//var a User
	//l =  a
	//l.basketball()
	//var b User1
	//l = b
	//l.basketball()

	//
	//
	//var c cat
	////值作为参数传递
	//invoke(&c)

	fmt.Println("可以直接调用,名字为：", ad.name)
	fmt.Println("也可以通过内部类型调用,名字为：", ad.user.name)
	fmt.Println("但是新增加的属性只能直接调用，级别为：", ad.level)

}

type user struct {
	name  string
	email string
}

type admin struct {
	user
	level string
}

type animal interface {
	printInfo()
}

func invoke(a animal) {
	a.printInfo()
}

func (c *cat) printInfo() {
	fmt.Println("a cat")
}

type cat int

type User1 User

type like interface {
	basketball()
	//volleyball()

}

func (a User) basketball() {
	println(a.Name + "like 篮球")
}

func (b User1) basketball() {
	println(b.Name + "not like 篮球")
}

func (u User) string() string {
	return "the person name is " + u.Name
}

func (u *User) modify() {
	u.Name = "李四"
}

func printf1(a ...interface{}) {

	for _, v := range a {
		fmt.Print(v)
	}

}
