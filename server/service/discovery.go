package service

import (
	"fmt"
	"os"
	"path/filepath"

	"rpcserver/utils"
)

type ServiceMap map[string]map[string][]string

type Discovery struct {
	sm ServiceMap
}

var skip []string

func NewDiscovery() (*Discovery, error) {
	// 扫描service/internal文件夹下的服务
	// NOTE：真实环境应该是扫描Docker的镜像
	svcDir := filepath.Join(filepath.Dir(os.Args[0]), "service", "internal")
	err := filepath.Walk(svcDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		} else if info.IsDir() || utils.Includes(skip, info.Name()) {
			return nil
		} else {
			fmt.Println(filepath.Join(path, info.Name()))
			return nil
		}
	})
	return nil, err
}