package main

import (
	"net/http"
	"html/template"
	"strings"
	"io"
)

var(
	i,j,k,m,n string
	x,z int
	mw string
)

func Vigenère(w http.ResponseWriter,r *http.Request)  {
	t,_:=template.ParseFiles("view/Vigenère.html")
	t.Execute(w,nil)
	m1:=map[int]string{0:"A",1:"B",2:"C",3:"D",4:"E",5:"F",6:"G",7:"H",8:"I",9:"J",10:"K",11:"L",12:"M",13:"N",14:"O",15:"P",16:"Q",17:"R",18:"S",19:"T",20:"U",21:"V",22:"W",23:"X",24:"Y",25:"Z"}
	m2:=map[string]int{"A":0,"B":1,"C":2,"D":3,"E":4,"F":5,"G":6,"H":7,"I":8,"J":9,"K":10,"L":11,"M":12,"N":13,"O":14,"P":15,"Q":16,"R":17,"S":18,"T":19,"U":20,"V":21,"W":22,"X":23,"Y":24,"Z":25}
	s1:=strings.SplitN(r.FormValue("mingwen"),"",100)
	s2:=strings.SplitN(r.FormValue("miyao"),"",100)
	s3:=make([]string,1)
	for y:=0;y<len(s1);y++ {
		z=y
		k=s1[z]
		if z>=len(s2) {
			z=z%2
		}
		m=s2[z]
		x=(m2[k]+m2[m])%26
		n=m1[x]
		s3=append(s3,n)
	}
	for _,v:=range s3{
		mw=mw+v
	}
	io.WriteString(w,mw)
}

func main()  {
	server:=http.Server{Addr:":8090"}
	http.HandleFunc("/",Vigenère)
	server.ListenAndServe()
	//var i,j,k,m,n string
	//	//var x,z int
	//	//m1:=map[int]string{0:"A",1:"B",2:"C",3:"D",4:"E",5:"F",6:"G",7:"H",8:"I",9:"J",10:"K",11:"L",12:"M",13:"N",14:"O",15:"P",16:"Q",17:"R",18:"S",19:"T",20:"U",21:"V",22:"W",23:"X",24:"Y",25:"Z"}
	//	//m2:=map[string]int{"A":0,"B":1,"C":2,"D":3,"E":4,"F":5,"G":6,"H":7,"I":8,"J":9,"K":10,"L":11,"M":12,"N":13,"O":14,"P":15,"Q":16,"R":17,"S":18,"T":19,"U":20,"V":21,"W":22,"X":23,"Y":24,"Z":25}
	//	//fmt.Println("请输入明文：")
	//	//fmt.Scanln(&i)
	//	//s1:=strings.SplitN(i,"",20)
	//	//fmt.Println("请输入密钥：")
	//	//fmt.Scanln(&j)
	//	//s2:=strings.SplitN(j,"",10)
	//	//s3:=make([]string,1)
	//	//for y:=0;y<len(s1);y++ {
	//	//	z=y
	//	//	k=s1[z]
	//	//	if z>=len(s2) {
	//	//		z=z%2
	//	//	}
	//	//	m=s2[z]
	//	//	//fmt.Println("m:",m)
	//	//	//fmt.Println("m2[k],m2[m]:",m2[k],m2[m])
	//	//	x=(m2[k]+m2[m])%26
	//	//	//fmt.Println("x:",x)
	//	//	n=m1[x]
	//	//	//fmt.Println("n:",n)
	//	//	s3=append(s3,n)
	//	//}
	//	//fmt.Println(s3)
}