# Concurrency management 

- WaitGroup
- Channel
- Context

# 適用情境

- WaitGroup: 需要將 job 拆分成多個子任務，待全部子任務完成後，才進行下一步
- Channel + select: 需要透過 channel 通知在背景執行的 goroutine 進行關閉等 (只能用在比較單純的 goroutine 情況)
- Context: 需要一次性控管多個 goroutine 則是用 context 比較方便