package utility

//日志的各种策略

import (
	"encoding/gob"
	file_cache "github.com/patrickmn/go-cache"
	"os"
	"time"
	log "github.com/sirupsen/logrus"
)

var cache *file_cache.Cache

func init()  {
	cache_file := "cache.gob"
	_, err := os.Lstat(cache_file)
	var M map[string]file_cache.Item
	if !os.IsNotExist(err) {
		File, _ := os.Open(cache_file)
		D := gob.NewDecoder(File)
		D.Decode(&M)
	}
	if len(M) > 0 {
		cache = file_cache.NewFrom(file_cache.NoExpiration, 10*time.Minute, M)
	} else {
		cache = file_cache.New(file_cache.NoExpiration, 10*time.Minute)
	}
	go func() {
		for {
			time.Sleep(time.Duration(60) * time.Second)
			File, _ := os.OpenFile(cache_file, os.O_RDWR|os.O_CREATE, 0777)
			defer File.Close()
			enc := gob.NewEncoder(File)
			if err := enc.Encode(cache.Items()); err != nil {
				log.Error(err)
			}
		}
	}()
}