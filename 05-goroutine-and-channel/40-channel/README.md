# Channel

goroutine 之間透過 channel 傳遞資料

- 01 (channel 基礎使用方式)
- 02 (channel 分離讀出與寫入的邏輯)
- 03 (deadlock)
- 04 (Closing channel & For-range loop with channel)

## Block vs Deadlock

channel 在等待值被讀出或寫入時，會進入阻塞狀態 (block)

此時 runtime 會切換到其他 goroutine (或 main) 繼續執行

當 block 永遠無法被解開時，就是 deadlock