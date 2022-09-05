# GOROOT

Go 原始碼的路徑

# GOPATH

所有東西都必須放在 GOPATH 底下

Go 1.8 之後預設 GOPATH 在 `$HOME/go`

## Folders in GOPATH

- `bin` 編譯後產生可執行檔案
- `src` 存放程式碼（工作目錄）
    - 資料夾結構 `$GOPATH/src/網域名稱/帳戶名稱/專案名稱`
        - e.g. `src/github.com/appleboy/helloworld`
- `pkg` 放置編譯後的 .a 檔案