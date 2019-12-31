# 메소드(Methods)



- ### 메소드(Methods)

메소드를 구조체(struct) 에 붙인 경우
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

- ### 값 리시버 (Value receiver)

func 다음 (f MyFloat) 에 * 를 붙이지 않은 경우  
ex. (f *MyFloat)  

```go
type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}

	return float64(f)
}

func methodsFuncs() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}
```
메소드는 구조체 뿐만 아니라 아무 타입에나 붙일 수도 있다.


<br>

- ### 포인터 리시버 (Pointer receiver)

```go
type Vertex3 struct {
	X, Y float64
}

func (v Vertex3) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y * v.Y)
}

func (v *Vertex3) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func methodsPointers() {
	v := Vertex3{3, 4}
	v.Scale(10)
	fmt.Println(v, v.Abs())
}
```

포인터 리시버는 메소드 리시버에 * 를 붙인 것이다.
*를 붙인 메소드 리시버는 복사본 객체가 아닌 원본 객체의 값을 바꿀 수 있게 된다.

Scale() 메소드는 메소드 리시버로 *Vertex3 가 정의되어 있으므로 기존 Vertex3
구조체의 값을 복사한 것이 아닌 직접 해당 구조체의 값을 바라보고 있다.

참고로. 보통 객체의 복사본이 아닌 원본을 바라보려면 &로 메모리 주소를 알아야하지만, methodsPointers() 메소드에 있는 Vertex3은 &가 붙어있지 않다.
메소드 리시버에서 * 를 붙인 경우 &가 필요 없다.  
그 이유는! 포인터 리시버인 경우 Go 인터프리터가 자동으로 &를 붙여준다고 한다.  
ex. v.Scale(10) -> (&v).Scale(10)  

반대의 경우도 똑같다. (만약 v := &Vertex3{3, 4} 인 경우)  
ex. v.Abs() -> (*v).Abs()

<br>

### 참고1
> 메소드(method)에 붙일 수 있는 리시버는 포인터 타입의 리시버(pointer receiver) 와 값 타입의 리시버(value receiver)가 있다.  
값 타입의 리시버는 메소드 리시버에 * 를 붙이지 않은 경우인데 이 경우
값을 단지 복사하여 사용한 것이고
포인터 타입의 리시버는 * 가 붙어 있는 경우이고 원본 객체를 바라보고 있다.  
참고로. 큰 구조체의 값을 복사하면 효율이 떨어지므로 복사하기 싫은 경우 포인터 리시버를 사용

### 참고2
> \* 와 & 의 차이는
> 1. &는 메모리 주소를 가리키는 것이고
> 2. \* 는 변수 선언 시 변수 앞에 * 를 붙인 경우 포인터형 변수가 되어 메모리 주소만 저장할 수 있다.
또한, 메모리 주소를 가진 변수에 값을 대입하고 싶은 경우
(포인터형 변수이거나 &가 붙어 있어 메모리 주소를 아는 경우)  
\* 를 붙이면 직접 값을 저장하거나 불러올 수 있다. (역참조 라고 함)

<br>

참고 2 예시코드
```go
func methodsPointers() {
	var num *int = new(int) // 포인터형 변수로 선언
	*num = 1          		// 역참조로 포인터형 변수에 값 저장
	fmt.Println(*num) 		// 포인터형 변수에서 값 호출

	var num int = 1
	var num2 *int = &num // & 로 num 변수의 메모리 주소를 구하여 num2 변수에 대입

	fmt.Println(num2) // 0xc0...메모리 주소
	fmt.Println(&num) // 0xc0...메모리 주소 (위와 동일한 주소)
}
```