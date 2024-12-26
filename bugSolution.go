func fastAppend(s []string, elems ...string) []string {
  s = append(s, elems...)
  return s
}

func main() {
  var strs []string
  strs = fastAppend(strs, "a", "b", "c")
  fmt.Println(strs)

  strs = fastAppend(strs, "d", "e", "f")
  fmt.Println(strs)
} 