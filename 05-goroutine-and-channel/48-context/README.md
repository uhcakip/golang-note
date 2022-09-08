# Context

使用 context 一次性關閉一次性關閉**多個** goroutine

---

使用 channel + select 主動通知 goroutine 停止工作

```go
func main() {
	stop := make(chan bool)

	go func() {
		for {
			select {
			case <-stop:
				println("stop the goroutine")
				return
			default:
				println("still working")
				time.Sleep(time.Second)
			}
		}
	}()

	time.Sleep(5 * time.Second)
	stop <- true
}
```

使用 context 主動通知 goroutine 停止工作

```go
func main() {
	// context.Background()
	// 初始化 context 作為 top-level context
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		for {
			select {
			case <-ctx.Done():
				println("cancel func has been called")
				return
			default:
				println("still working")
				time.Sleep(time.Second)
			}
		}
	}()

	time.Sleep(5 * time.Second)
	cancel()
	time.Sleep(5 * time.Second)
}
``` 