# 맵(Map) 

## 맵 리터럴(Map literals)

```go
var m = map[string]Vertex {
	"Bell Labs" : {40.6, -74.3},
	"Google"	: {37.42, -122.08},
}
```

맵 리터럴의 가장 상위 타입이 타입명(Vertex) 이면, 하위 리터럴 표현 시 타입명 생략 가능하다. 