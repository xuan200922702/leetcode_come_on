package main


import (
	"errors"
	"fmt"
	"strconv"
)

type  stack struct {
	MaxTop int
	Top   int
	nums  [5]int
}
//入栈
func (this *stack)Push(val  int)(err error)  {
	//先判断是否满了
	if  this.Top ==this.MaxTop{
		fmt.Println("stack full")
		return errors.New("stack full")
	}
	this.Top++
	//放入数据
	this.nums[this.Top]=val
	return
}

//出栈
func(this *stack)Pop()(val int,err  error){
	if  this.Top==-1{
		fmt.Println("stack is  empty")
	}
	val=this.nums[this.Top]
	this.Top--
	return val,nil
}
//判断一个字符是不是一个运算符
func (this *stack)IsOper(val  int)bool{
	if val ==42 ||  val==43 ||val ==45  ||val==47{
		return true
	}else {
		return false
	}
}
//运算的方法
func (this *stack)Cal(num1 int,num2 int,oper  int)int{
	res:=0
	switch oper {
	case 42:
		res=num2  *  num1
	case 43:
		res=num2  +  num1
	case 45:
		res=num2  -  num1
	case 47:
		res=num2  /  num1
	default:
		fmt.Println("运算符错误")
	}
	return res

}



//判断优先级
func (this *stack)Priorty(oper int)int{
	res:=0
	if oper==42 ||oper==47{
		res=1
	}else if oper==43 ||oper  ==47{
		res=0
	}
	return res
}

//遍历栈
func(this  *stack)List(){
	if this.Top==-1{
		fmt.Println("stack empty")
		return
	}
	for i:=this.Top;i>=0;i--{
		fmt.Printf("nums[%d]=%d\n",i,this.nums[i])
	}
}

func main()  {
	//数栈
	numstack:=&stack{
		MaxTop: 20,
		Top:   -1,
	}
	//符号栈
	operstack:=&stack{
		MaxTop: 20,
		Top:   -1,
	}
	exp:="3+2*6-2"
	index:=0
	num1:=0
	num2:=0
	oper:=0
	result:=0
	keepNum :=""
	for{
		ch:=exp[index:index+1]
		temp:=int([]byte(ch)[0])
		if operstack.IsOper(temp){
			if operstack.Top==-1{
				operstack.Push(temp)
			}else {
				if operstack.Priorty(operstack.nums[operstack.Top]) >=operstack.Priorty(temp){
					num1,_=numstack.Pop()
					num2,_=numstack.Pop()
					oper,_=operstack.Pop()
					result=operstack.Cal(num1,num2,oper)
					numstack.Push(result)
					operstack.Push(temp)

				}else {
					operstack.Push(temp)
				}
			}

		}else { //说明是数
			keepNum+=ch
			if  index ==  len(exp) -1{
				val,_ :=strconv.ParseInt(ch,10,64)
				numstack.Push(int(val))
			}else {
				if  operstack.IsOper(int([]byte(exp[index+1:index+2])[0])){
					val,_ :=strconv.ParseInt(ch,10,64)
					numstack.Push(int(val))
					keepNum=""
				}

			}

		}

		//继续 扫描
		if index+1==len(exp){
			break
		}
		index++
	}

	for{
		if  operstack.Top==-1{
			break
		}
		num1,_=numstack.Pop()
		num2,_=numstack.Pop()
		oper,_=operstack.Pop()
		result=operstack.Cal(num1,num2,oper)
		numstack.Push(result)


	}
	res,_:=numstack.Pop()
	fmt.Printf("表达式%s = %v",exp,res)
}
