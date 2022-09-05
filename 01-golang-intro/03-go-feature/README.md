# Features of Golang

- import 的 package 若沒用到需移除，進而限制檔案 binary 檔案大小
- 非物件導向（無繼承特性）
- package 無法繼承其他 package（速度快的原因）
- 強型別
- 沒有錯誤處理
- 用字首大小寫來區別是否可存取（public / private）
    - 大寫 = 可以被其他 package 調用
    - 小寫 = 僅當前的 package 可以調用
- 沒用到的 import 或變數會引起編譯錯誤
- 完整的標準函式
- 支援 UTF-8 格式
- 所有的 package 只有 function 和 method

  ```go
  // method
  // 可以寫一個 structor，有很多個 member 或 method
  func (b *Task) loadAttributes(e Engine) (err error) {
      return b.loadUser(e)
  }
  
  // function
  func getMaxSorter(blueID int64) (int, error) {
      var maxSorter int
      if _, err := x.Select("MAX(sroter)").Table("task").Where("blue_id = ?", blueID).Get(&maxSorter); err != nil {
          return 0, err
      }
      return maxSorter, nil
  }
  ```

# Benefits of Golang

- 開發及執行效率
- 由 Google 維護
- 部署方便
- 跨平台編譯（直接把 binary 檔丟給別人測試）
- 內建 Coding Style, Testing 等工具
- 多核心處理
- 團隊開發工具整合
    - Coding Style
    - Testing Tool
    - Benchmark Tool
- 系統效能（記憶體用量、CPU 使用率 ...）
    - 重啟時間非常快，Load-Balancer 不需要 Pre-warning
    - EC2 使用量降低（降低 80 - 85%）
    - Response time 100ms —> 10ms