# 에러 (Error) 

golang 에서 에러는 하나의 값이다.  
그래서 여러가지 처리 방법이 있는데 이 예제에서는 단순 출력을 한다.  


```go
type MyError struct {
	When time.Time
	What string
}

func run() error {
	return &MyError {
		time.Now(),
		"it didn't work",
	}
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func errors() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
```

golang은 error 타입이 interface로 정의되어 있으며 단순히 Error() 메소드 하나만을 가지고 있다.  
그래서 Error() 메소드를 구현하기만 하면 error 로 사용할 수 있다.

<br>

error interface

```go
type error interface {
    Error() string
}
```

- ### 연습코드

```go
type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}

	z := float64(1)
	for i:= 1; i<=10; i++ {
		z -= (z*z - x) / (2*z)
	}

	return z, nil
}

func exerciseErrors(){
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
```

fmt 패키지의 출력함수들은 error에 대한 출력을 기본으로 하기 때문에, (val, error)  
리턴 값에 error 값이 있을 경우, 해당 error 구현체를 찾아서 출력한다.  

Sqrt에서 ErrNegativeSqrt(x) 을 리턴하고 있고 ErrNegativeSqrt 타입을 구현한 Error() 메소드가 있으므로 해당 Error() 메소드의 string 값을 리턴하여 출력한다.
