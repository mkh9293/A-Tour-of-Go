# 맵(Map) 



- ### 맵 리터럴(Map literals)

```go
var m = map[string]Vertex {
	"Bell Labs" : {40.6, -74.3},
	"Google"	: {37.42, -122.08},
}
```

맵 리터럴의 가장 상위 타입이 타입명(Vertex) 이면, 하위 리터럴 표현 시 타입명 생략 가능하다.   

<br>

- ### 맵 다루기

```go
m := make(map[string]int)

m["Answer"] = 42
fmt.Println("The value:", m["Answer"])

m["Answer"] = 48
fmt.Println("The value:", m["Answer"])

delete(m, "Answer")
fmt.Println("The value:", m["Answer"])

v, ok := m["Answer"]
fmt.Println("The value:", v, "Present?", ok)
```

golang은 delete() 메소드로 요소를 제거한다.  
키 존재 여부 확인을 위해 두번째 변수(ok)를 추가하여 확인 가능.  
map은 값이 없을 경우 0을 리턴.

<br>

- ### strings 함수

```go
func WordCount(s string) map[string]int {
	arr := strings.Fields(s)
	result := make(map[string]int)

	for _, v := range arr {
		result[v]++
	}

	return result
}
```

strings 패키지는 간단한 문자 조작 함수를 제공.  
Fields() 메소드는 인수로 받은 문자열에 있는 문자를 배열로 리턴해줌.