# Common commands

- `go env` 會列出 Go 各個變數
- `go get` 會把 source code 抓下來，若有 main 函式會 build 成一個執行檔放在 `$GOPATH/bin` 底下（需確認 `$GOPATH/bin` 有在 `$PATH` 底下）

  ```bash
  # .zshrc OR .bash_profile OR .bashrc
  export PATH=$GOPATH/bin:$PATH
  ```