package mystrings 

// in go to export a function it must be capitalised otherwise you cant export it
func Reverse (s string) string {
  result := ""
  for _ , v := range s {
    result = string(v) + result
  }
  return result 
}
