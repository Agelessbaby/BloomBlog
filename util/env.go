package env

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"path"
	"runtime"
	"time"
)

var (
	ProjectRootPath = path.Dir(GetConcurrentPath()) + "/"
	Loglevelmap     = map[string]klog.Level{"info": klog.LevelInfo, "debug": klog.LevelDebug, "error": klog.LevelError, "fatal": klog.LevelFatal}
)

func GetConcurrentPath() string {
	_, filename, _, _ := runtime.Caller(0)
	return path.Dir(filename) //return the dir that the code currently runs on
}

const (
	TOKEN_EXPIRE = time.Hour * 24 * 7
	JWT_SECRET   = "BloomBlog"
)
