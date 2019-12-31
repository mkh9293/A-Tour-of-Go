# 메소드(Methods) 



- ### 메소드(Methods))

```go
type Vertex2 struct {
	X, Y float64
}

func (v Vertex2) Abs() float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

func methods() {
	v := Vertex2{3, 4}
	fmt.Println(v.Abs())
}

```

golang에는 클래스가 없다.  
대신, 메소드를 구조체(struct)에 붙여서 사용할 수 있다.  
메소드 리시버는 func 키워드와 메소드의 이름 사이에 인자로 들어간다.

<br>