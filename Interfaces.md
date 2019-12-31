# 인터페이스 (Interface) 


인터페이스는 메소드들의 집합을 정의한 것이다. (메소드 정의서)

```go
type Abser interface {
	Abs() float64
}

func interfaces() {
	var a Abser
	f := MyFloat2(-math.Sqrt2)
	v := Vertex4{3, 4}

	a = f
	a = &v

	fmt.Println(a.Abs())
}

type MyFloat2 float64

func (v MyFloat2) Abs() float64 {
	if(v < 0) {
		return float64(-v)
	}

	return float64(v)
}

type Vertex4 struct {
	X, Y float64
}

func (v *Vertex4) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
```

Abser 인터페이스는 Abs() 메소드를 정의하고 있으며  
Vertex4 타입의 값으로 구현된 메소드와 MyFloat 타입의 값으로 구현된 메소드 모두  
Abser 인터페이스 타입을 정의한 값(a) 이 호출하여 사용 가능하다.
<br>

예시코드2

```go
type I interface {
	M()
}

func interfaceValues() {
	var i I

	// 1 
	i = &T{"Hello"}
	describe(i)
	i.M()

	// 2
	i = F(math.Pi)
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

// 1
type T struct {
	S string
}

func (t *T) M() {
	fmt.Println(t.S)
}

// 2
type F float64

func (f F) M() {
	fmt.Println(f)
}
```
기존에 포인터 리시버로 구현된 메소드를 호출하여 사용할 때 자동으로 &를 붙여 줬었지만,  
인터페이스로 선언된 변수가 포인터 리시버로 구현된 메소드를 호출할 때는 자동으로 붙여주지 않는다.  

그 이유는 i = &T{"Hello"} 처럼 처음 인터페이스 타입 변수에 값을 할당할 때,  
해당 타입 값으로 구현된 메소드가 존재하는지 찾는데, 찾을 때는 &를 인터프리터가 자동으로 붙여주지 못한다.

<br>

참고.
> golang은 처음 인터페이스 타입 변수에 값을 할당할 때,   
구현체(메소드)를 찾기 때문에 구현체에 인터페이스 이름을 달 필요가 없다.  
(ex. 자바의 구현 클래스에 "implements F" 처럼 안 붙여도 됨) 