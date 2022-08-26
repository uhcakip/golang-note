# Multiple channels

透過 goroutine 跑 20 個 job

- 使用 outChan 顯示 job 完成狀況
- 使用 errChan 顯示 job 發生的錯誤，並跳出 main func
- 使用 finishChan 通知所有 job 皆已完成
- 設定 timeout (一秒內要完成所有 job)