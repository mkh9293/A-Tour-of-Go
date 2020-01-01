# 고루틴 (Goroutine) 

고루틴은 go 런타임에 의해 관리되는 경량 쓰레드이며  
실행된 스레드는 자원이 공유되므로 동기화에 신경써야 한다.

```go
func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func goroutines() {
	go say("world")
	say("hello")
}

```

go say("~") 명령에 의해 새로운 스레드에서 코드가 실행됨.

<br>

- ### 채널(channels)

```go
func sum(a []int, c chan int) {
	sum := 0

	for _, v := range a {
		sum += v
	}

	c <- sum
}

func channels() {
	a := []int{7,2,8,-9,4,0}

	c := make(chan int)
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)
	x, y := <-c, <-c

	fmt.Println(x, y, x+y)
}
```

채널은 연산자로 "<-" 를 사용하며 타입이 존재하며  
채널은 사용되기 전에 생성되어야 한다.  

예시코드.

```go
func main() {
    ch <-v      // v를 ch로 보냄
    v := <-ch   // ch로부터 값을 받아서 v로 넘김

    // 채널 생성
    ch := make(chan int)
}
```

채널 송/수신은 상대편이 준비될때까지 블록되며,  
이러한 특성 덕분에 고루틴이 명시적인 락 or 어떠한 조건 없이도 동기화 될 수 있도록 한다.

<br>

- ### 버퍼링 채널 (buffered channels)

채널은 버퍼링 될 수 있다.  
make 로 채널 생성 시, 두번째 인자에 버퍼의 크기를 지정할 수 있다.  
해당 크기 만큼 버퍼가 찰 때까지 블록되며  
수신 측은 버퍼가 빌 때 블록된다.

```go
func befferedChannels() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	
	fmt.Println(<-ch)
	fmt.Println(<-ch)	
}
```
참고로 블록된다는 뜻은 동기화가 유지된다는 뜻이며,  
스레드가 멈추는 것은 아니다.

<br>

- ### range 와 close

송신 측에서 더이상 보낼 데이터가 없을 때 명시적으로 close() 명령어를 사용하여 채널을 닫을 수 있다.   
수신 측은 채널이 닫혔는지 확인할 수 있다.

```go
func main() {
    v, ok := <-ch // 채널이 닫혔는지 ok 로 확인가능하며 닫혔다면 false.
}
```

또한 수신측은 for i := range c 로 채널이 닫힐 때까지 계속해서 값을 받을 수 있다.

```go
func fibonacci2(n int, c chan int) {
	x, y := 0, 1

	for i := 0; i<n; i++ {
		c <- x
		x, y = y, x+y
	}

	close(c)
}

func rangeAndClose() {
	c := make(chan int, 10)

	go fibonacci2(cap(c), c)
	
	for i := range c {
		fmt.Println(i)
	}
}
```

수신측은 채널을 닫을 수 없다.  
만약 닫힏 채널에 데이터를 보내려 할 경우 패닉이 발생한다.  

채널은 파일과 다르기 때문에 꼭 닫을 필요는 없다.

<br>

- ### 셀렉트(select)

셀렉트는 다수의 통신 동작 중 하나라도 수행할 수 있을 때까지 기다릴 수 있게 한다.

```go
func fibonacci3(c, quit chan int) {
	x, y := 0, 1

	for {
		select {
		case c <- x :
			x, y = y, x+y
		case <-quit:
			fmt.Println("Quit!")
			return
		}
	}
}

func selects() {
	c := make(chan int)
	quit := make(chan int)
	
	go func() {
		for i:=0; i<10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()

	fibonacci3(c, quit)
}
```

go func() 에서 <-c 에서 블록되고 fibonicci3가 실행되는데,  
select 구문에서 case c <- x 때문에 go func() 의 <-c 값이 출력되고  
해당 for문이 끝나면, quit <- 0 때문에 fibonicci3의 case <-quit: 의 내용이 실행되어 Quit1 을 출력하고 끝낸다.

위 코드처럼 select를 이용하면 스레드의 수행동작을 명시할 수 있다.

<br>

- ### 디폴트 셀렉트 (default select)

수행 준비가 완료된 케이스가 없는 경우 default 케이스를 사용하는 방법도 있다.

```go
func defaultSelection() {
	tick := time.Tick(1e8)
	boom := time.After(5e8)

	for {
		select {
		case <- tick :
			fmt.Println("tick.")
		case <- boom :
			fmt.Println("boom.")
			return
		default:
			fmt.Println("    .")
			time.Sleep(5e7)
		}
	}
}
```