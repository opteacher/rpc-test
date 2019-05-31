package service

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"strings"

	"rpcserver/utils"
)

type baseService struct{}

type Discovery struct {
	sm map[string]reflect.Value
	mm map[string][]string
	bs *baseService
}

var skip = []string{
	"discovery.go",
}

func NewDiscovery() (*Discovery, error) {
	ds := &Discovery{
		sm: make(map[string]reflect.Value),
		mm: make(map[string][]string),
		bs: new(baseService),
	}
	// 扫描service/internal文件夹下的服务
	// NOTE：真实环境应该是扫描Docker的镜像
	svcDir := filepath.Join(filepath.Dir(os.Args[0]), "service")
	err := filepath.Walk(svcDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		} else if info.IsDir() || utils.Includes(skip, info.Name()) {
			return nil
		} else if content, err := utils.ReadFromFile(path); err != nil {
			return err
		} else if tpReg, err := regexp.Compile(`^type\s+(?P<class_name>\w+)\s+struct`); err != nil {
			return err
		} else {
			types := utils.Select(strings.Split(content, "\n"), tpReg)
			return combineServiceMap(tpReg, types, ds)
		}
	})
	return ds, err
}

func combineServiceMap(tpReg *regexp.Regexp, types []string, ds *Discovery) error {
	for _, typ := range types {
		if strs := tpReg.FindStringSubmatch(typ); len(strs) < 2 {
			continue
		} else {
			sname := strings.Title(strs[1])
			svc := reflect.ValueOf(ds.bs).MethodByName("New" + sname).Call(nil)[0].Interface()
			stype := reflect.TypeOf(svc)
			for i := 0; i < stype.NumMethod(); i++ {
				method := stype.Method(i)
				if _, exs := ds.mm[sname]; !exs {
					ds.sm[sname] = reflect.ValueOf(svc)
				}
				ds.mm[sname] = append(ds.mm[sname], method.Name)
			}
		}
	}
	return nil
}

func (ds *Discovery) wrapParams(params []interface{}) ([]reflect.Value, error) {
	var vs []reflect.Value
	for _, p := range params {
		vs = append(vs, reflect.ValueOf(p))
	}
	return vs, nil
}

func (ds *Discovery) CallMethod(method string, params []interface{}) ([]byte, error) {
	if strs := strings.Split(method, "."); len(strs) != 2 {
		return nil, fmt.Errorf("需要指定请求的服务和接口名：%s", method)
	} else if sname, mname := strs[0], strs[1]; false {
		return nil, fmt.Errorf("能进来有鬼了 @o@")
	} else if svc, exs := ds.sm[sname]; !exs {
		return nil, fmt.Errorf("未找到服务：%s", sname)
	} else if !utils.Includes(ds.mm[sname], mname) {
		return nil, fmt.Errorf("未找到接口：%s", mname)
	} else if ps, err := ds.wrapParams(params); err != nil {
		return nil, fmt.Errorf("打包参数错误：%v", err)
	} else if res := svc.MethodByName(mname).Call(ps); len(res) == 0 {
		return nil, fmt.Errorf("服务未返回任何数据")
	} else if resp, err := json.Marshal(res[0].Interface()); err != nil {
		return nil, fmt.Errorf("服务返回的JSON格式有误：%v", err)
	} else {
		return resp, nil
	}
}
