# Job queue

是否需要用到 job queue 取決於 server 的狀況及該 job 會消耗多少 CPU 及記憶體 (像是寄信這種就不太需要)

e.g. 短時間有大量的 job 需要處理時，可透過 job queue 限定同時能處理的個數

需注意 buffered channel 容量和實際處理 job 的速率不要差太多，否則 job 會一直被 block 住

