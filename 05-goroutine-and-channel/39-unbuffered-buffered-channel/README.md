# Unbuffered and Buffered Channel

- 01 (Unbuffered channel)
- 02 (Buffered channel)

## Differences

Unbuffered channel

- 同步
- 需要等到讀或寫都完成 (同時有一邊寫入，另一邊讀出)，main 才會結束
- 讀跟寫需要在不同的 goroutine (或 main) 才不會被 block

Buffered channel

- 異步
- 事先給定 channel 的容量
- 在寫入的數量沒有超出容量的情況下，一直寫入不讀出來 main 也會繼續執行
