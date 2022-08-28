# Select statement with channels

## Switch vs Select

Switch

- 照順序執行
- 常用於接口 interface{} 型別判斷 `i.type()

```go
func convert(i interface{}) {
	switch t := i.(type) { 
		case int:
			// ...
		case string:
			// ...
		case float64:
			// ..
	}
}
```

Select

- 只能用於 channel
- 會隨機選取 case (random value 特性)
- 沒設定 default 的話會被 block 住

```go
func main() {
	ch := make(chan int, 1)
	ch <- 1 // 因為沒有設定 default，註解掉這句會產生 deadlock
	
	select {
		case <-ch:
			println("01")
		case <-ch:
			println("02")
	}

    /*
       output:
           01 或 02 都有可能
    */
}
```

```go
func main() {
	ch := make(chan int, 1)
	ch <- 1
	
	select {
		case ch <- 2: // 超過限定容量，塞不進去，所以不會執行這個 case
			println("Channel value is", <-ch)
		default:
			println("Channel blocking")
	}

    /*
	output:
	    Channel blocking (容量改 2 則會印出 Channel value is 1)
    */
}
```

- 01 (Timeout)
- 02 (Breaking for loop channel)
  - 透過 for-select 持續讀出 channel 裡的資料