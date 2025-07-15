// Copyright 2025 JuZX <wo_sakura@163.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/ketitongxue/miniblog.

package app

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/ketitongxue/miniblog/cmd/mb-apiserver/app/options"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	configFile string
	opts       *options.ServerOptions
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "mb-apiserver",
	Short: "A brief description of mb-apiserver",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using mb-apiserver. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
	// 命令出错时，不打印帮助信息。设置为 true 可以确保命令出错时一眼就能看到错误信息
	SilenceUsage: true,
	// 指定调用 cmd.Execute() 时，执行的 Run 函数
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := viper.Unmarshal(opts); err != nil {
			return err
		}
		fmt.Printf("All Viper settings: %v\n", viper.AllSettings())

		// 输出 opts 结构体内容
		jsonData, _ := json.MarshalIndent(opts, "", "  ")
		fmt.Println(string(jsonData))

		return nil
	},
	// 设置命令运行时的参数检查，不需要指定命令行参数。例如：。/miniblog param1 param2
	Args: cobra.NoArgs,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.mb-apiserver.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	// 初始化配置函数，在每个命令运行时调用
	cobra.OnInitialize(onInitialize)

	// cobra 支持持久性标志（PersistentFlag），该标志可用于它所分配的命令以及该命令下的每个子命令
	// 推荐使用配置文件来配置应用，便于管理配置项
	rootCmd.PersistentFlags().StringVarP(&configFile, "config", "c", filePath(), "Path to the miniblog configuration file.")
	opts = options.NewServerOptions()
	// 将 ServerOptions 中的选项绑定到命令标志
	opts.AddFlags(rootCmd.PersistentFlags())
}
