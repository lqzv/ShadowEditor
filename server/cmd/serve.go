// Copyright 2017-2020 The ShadowEditor Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
//
// For more information, please visit: https://github.com/tengge1/ShadowEditor
// You can also visit: https://gitee.com/tengge1/ShadowEditor

package cmd

import (
	"bufio"
	"log"
	"os"

	"github.com/inconshreveable/mousetrap"
	"github.com/spf13/cobra"

	"github.com/tengge1/shadoweditor/server"

	// Register the server sub packages.
	_ "github.com/tengge1/shadoweditor/server/assets"   // assets api
	_ "github.com/tengge1/shadoweditor/server/category" // category api
	_ "github.com/tengge1/shadoweditor/server/export"   // export api
	_ "github.com/tengge1/shadoweditor/server/system"   // system api
	_ "github.com/tengge1/shadoweditor/server/tools"    // tools api
	_ "github.com/tengge1/shadoweditor/server/upload"   // upload api
)

// serveCmd launch the shadow editor server.
var serveCmd = &cobra.Command{
	Use:     "serve",
	Short:   "Start shadow editor server",
	Aliases: []string{"server"},
	Long:    `Use shadow editor server to provider scene and model data.`,
	Run: func(cmd *cobra.Command, args []string) {
		runServe()
	},
}

func runServe() {
	if _, err := os.Stat(cfgFile); os.IsNotExist(err) {
		log.Printf("cannot find config file: %v", cfgFile)
		wait()
		return
	}

	err := server.Create(cfgFile)
	if err != nil {
		log.Printf(err.Error())
		wait()
		return
	}

	server.Start()
	wait()
}

func wait() {
	// When you double click ShadowEditor.exe in the Windows explorer,
	// wait to see errors in order not to crash immediately.
	if mousetrap.StartedByExplorer() {
		reader := bufio.NewReader(os.Stdin)
		reader.ReadString('\n')
	}
}
