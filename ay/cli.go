package ay

import (
	"sync"
)

var once sync.Once

func Init(fn ...func()) {

	fn = append(fn, LoggerInit)
	fn = append(fn, InitConfig)
	fn = append(fn, InitDB)

	once.Do(func() {
		for _, f := range fn {
			f()
		}
	})
	go WatchConf()
}
