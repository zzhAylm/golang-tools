package main

import "fmt"

// map
func main() {
	maps := make(map[string]string, 10)
	maps["zzh"] = "zzh"
	maps["ylm"] = "ylm"
	fmt.Println(maps)
	delete(maps, "zzh")
	fmt.Println(maps)
	s, ok := maps["ylm"]
	fmt.Println(s, ok)
}
