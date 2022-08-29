# 使用 Goroutine 及 Channel 的時機

- 01 (goroutine 之間共享同個變數造成每次執行的結果都不同)
  - 避免 goroutine 之間共享同個變數
- 02 (使用 WaitGroup + Lock) 
- 03 (使用 Channel)
  - 用於多個 goroutine 之間需要交換資料

