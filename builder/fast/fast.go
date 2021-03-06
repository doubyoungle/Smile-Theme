package fast

import (
	"io/ioutil"
	"main/builder/log"
	"os"
	"path/filepath"
)

// Each 遍历路径
func Each(dir string, includeChild bool, action func(string, bool) bool) bool {
	info, err := os.Stat(dir)
	if err != nil {
		log.Error(err)
	}
	if !info.IsDir() {
		log.Error("获取到的路径不是目录")
	}
	infos, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Error(err)
	}
	for _, fi := range infos {
		p := filepath.Join(dir, fi.Name())
		if action(p, !fi.IsDir()) {
			return true
		}

		if fi.IsDir() && includeChild {
			if Each(p, true, action) {
				return true
			}
		}
	}
	return false
}
