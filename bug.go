func slowAppend(s []string, elems ...string) []string {
  for _, elem := range elems {
    s = append(s, elem)
  }
  return s
}

func main() {
  var strs []string
  strs = slowAppend(strs, "a", "b", "c")
  fmt.Println(strs)

  strs = slowAppend(strs, "d", "e", "f")
  fmt.Println(strs)
}