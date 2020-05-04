// Copyright 2017-2020 The ShadowEditor Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
//
// For more information, please visit: https://github.com/tengge1/ShadowEditor
// You can also visit: https://gitee.com/tengge1/ShadowEditor

package data

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/tengge1/shadoweditor/helper"
)

// CreateAnimationDataFile 复制动画数据文件
func CreateAnimationDataFile(path string) {
	dirName := filepath.Join(path, "api", "Animation")

	if _, err := os.Stat(dirName); os.IsNotExist(err) {
		os.MkdirAll(dirName, 0755)
	}

	// 获取列表
	fileName := filepath.Join(path, "api", "Animation", "List")
	data, _ := helper.ToJSON(map[string]interface{}{
		"Code": 200,
		"Msg":  "获取成功！",
		"Data": []string{},
	})
	ioutil.WriteFile(fileName, []byte(data), 0755)

	// 其他接口
	apiList := []string{
		"/api/Animation/Add",
		"/api/Animation/Edit",
		"/api/Animation/Delete",
	}

	data, _ = helper.ToJSON(map[string]interface{}{
		"Code": 300,
		"Msg":  "演示程序，操作失败！",
	})

	for _, i := range apiList {
		fileName = filepath.Join(path, i)
		ioutil.WriteFile(fileName, []byte(data), 0755)
	}
}