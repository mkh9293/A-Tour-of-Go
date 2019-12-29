# 함수(Function) 



- ### 함수 값(Function values)

```go
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func functionValues() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt((x * x) + (y * y))
	}

	fmt.Println(hypot(3, 4))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}

```

golang에서는 함수도 값과 똑같다.  
함수 자체가 다른 함수의 인수로 전달될 수 있다. (위 소스에서는 Println의 인수로 전달됨)

<br>

- ### 함수 클로저

```go
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func functionClosures() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}
```

golang에서 함수 클로저를 사용할 경우 각 클로저는 자신만의 값을 저장할 수 있다.  



> 클로저는 내부함수에서 외부함수 블록에 접근할 수 있는 것 뜻한다. 