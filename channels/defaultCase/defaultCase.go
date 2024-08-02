package main

import (
  "time"
)


func saveBackUps (snapshotTicker, saveAfter <- chan time.Time, logChan chan string) {
  
  for {
    select {
    case  <- snapshotTicker:
      takeSnapshot(logChan)
    case <- saveAfter:
      saveSnapshot(logChan)
      return
    default: 
     waitForData(logChan)
     time.Sleep(500 * time.Millisecond)
    }
  }
  
}

func takeSnapshot(logChan chan string ){
  logChan <- "Taking a backup snapshot"
}

func saveSnapshot(logChan chan string){
  logChan <- "All backups saved"
  close(logChan)
}

func waitForData(logChan chan string){
  logChan <- "Nothing to do, waiting..."
}
