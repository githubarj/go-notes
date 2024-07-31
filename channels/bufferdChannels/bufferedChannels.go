package main

func addEmailsToQueue(emails []string) chan string {
    batch := make(chan string, len(emails)) 
  for _ , email := range emails {
    batch <- email 
  }
  return batch
}
