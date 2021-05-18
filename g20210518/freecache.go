package g20210518

import (
	"fmt"
	"github.com/coocood/freecache"
	"runtime/debug"
)
// Go缓冲系列之-free-cache
// https://mp.weixin.qq.com/s/SwGTdFe9AgHPGqkDbZoUjQ
func main() {
	// 缓存大小，100M
	cacheSize := 100 * 1024 * 1024
	cache := freecache.NewCache(cacheSize)
	debug.SetGCPercent(20)
	key := []byte("abc")
	val := []byte("def")

	expire := 60 // expire in 60 seconds
	// 设置 key
	cache.Set(key, val, expire)
	got, err := cache.Get(key)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("$s\n", got)
	}
	fmt.Println("entry count ", cache.EntryCount())
	affected := cache.Del(key)

	fmt.Println("deleted key ", affected)
	fmt.Println("entry count ", cache.EntryCount())

}
