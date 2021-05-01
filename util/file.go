package util

import (
	"github.com/pkg/errors"
	"os"
	"path/filepath"
)

// 获取当前运行环境的路径
func RuntimeDir() (string, error) {
	ex, err := os.Executable()
	if err != nil {
		return "", errors.Wrap(err, "Get the dir path of current process runtime failure. ")
	}

	return filepath.Dir(ex), nil
}
