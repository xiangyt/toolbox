/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/xiangyt/tools/file"

	"github.com/spf13/cobra"
)

// fileCmd represents the file command
var fileCmd = &cobra.Command{
	Use:   "file",
	Short: "文件操作",
	Long:  `文件相关操作`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("file called")
	},
}

var fnRenameCmd = &cobra.Command{
	Use:   "rename",
	Short: "批量修改文件名",
	Long:  "批量修改某个文件夹下所有符合条件的文件名,默认将`[资源来源]动漫名 集数.mkv`替换为`[资源来源]动漫名 SxxExx.mkv`",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("file rename called. regexpStr:%s, pathStr:%s, fileNameSuffix:%s\r\n", regexpStr, pathStr, fileNameSuffix)
		file.BatchRename(regexpStr, pathStr, fileNameSuffix, season, recursion)
	},
}

var fileNameSuffix string
var season string
var regexpStr string
var pathStr string
var recursion bool

func init() {
	rootCmd.AddCommand(fileCmd)
	fileCmd.AddCommand(fnRenameCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// fileCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// fileCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	fnRenameCmd.Flags().StringVarP(&fileNameSuffix, "ext", "e", "mkv", "文件后缀名")
	fnRenameCmd.Flags().StringVarP(&season, "season", "s", "01", "xx季")
	fnRenameCmd.Flags().StringVarP(&regexpStr, "regexp", "r", `^\[(.*)\](.*)(\d\d).*$`, "正则表达式")
	fnRenameCmd.Flags().StringVarP(&pathStr, "path", "p", "./", "目标文件夹路径")
	fnRenameCmd.Flags().BoolVarP(&recursion, "recursion", "v", false, "是否递归执行")
}
