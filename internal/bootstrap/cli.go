package bootstrap

import (
	"strings"

	"github.com/urfave/cli/v3"

	"github.com/tnb-labs/panel/internal/app"
	"github.com/tnb-labs/panel/internal/route"
)

func NewCli(cmd *route.Cli) *cli.Command {
	cli.RootCommandHelpTemplate = strings.ReplaceAll(cli.RootCommandHelpTemplate, "NAME", "名称")
	cli.RootCommandHelpTemplate = strings.ReplaceAll(cli.RootCommandHelpTemplate, "USAGE", "用法")
	cli.RootCommandHelpTemplate = strings.ReplaceAll(cli.RootCommandHelpTemplate, "VERSION", "版本")
	cli.RootCommandHelpTemplate = strings.ReplaceAll(cli.RootCommandHelpTemplate, "DESCRIPTION", "描述")
	cli.RootCommandHelpTemplate = strings.ReplaceAll(cli.RootCommandHelpTemplate, "AUTHOR", "作者")
	cli.RootCommandHelpTemplate = strings.ReplaceAll(cli.RootCommandHelpTemplate, "COMMANDS", "命令")
	cli.RootCommandHelpTemplate = strings.ReplaceAll(cli.RootCommandHelpTemplate, "GLOBAL OPTIONS", "全局选项")
	cli.RootCommandHelpTemplate = strings.ReplaceAll(cli.RootCommandHelpTemplate, "COPYRIGHT", "版权")
	cli.CommandHelpTemplate = strings.ReplaceAll(cli.CommandHelpTemplate, "NAME", "名称")
	cli.CommandHelpTemplate = strings.ReplaceAll(cli.CommandHelpTemplate, "USAGE", "用法")
	cli.CommandHelpTemplate = strings.ReplaceAll(cli.CommandHelpTemplate, "CATEGORY", "分类")
	cli.CommandHelpTemplate = strings.ReplaceAll(cli.CommandHelpTemplate, "DESCRIPTION", "描述")
	cli.CommandHelpTemplate = strings.ReplaceAll(cli.CommandHelpTemplate, "OPTIONS", "选项")
	cli.SubcommandHelpTemplate = strings.ReplaceAll(cli.SubcommandHelpTemplate, "NAME", "名称")
	cli.SubcommandHelpTemplate = strings.ReplaceAll(cli.SubcommandHelpTemplate, "USAGE", "用法")
	cli.SubcommandHelpTemplate = strings.ReplaceAll(cli.SubcommandHelpTemplate, "DESCRIPTION", "描述")
	cli.SubcommandHelpTemplate = strings.ReplaceAll(cli.SubcommandHelpTemplate, "COMMANDS", "命令")
	cli.SubcommandHelpTemplate = strings.ReplaceAll(cli.SubcommandHelpTemplate, "OPTIONS", "选项")

	cli.RootCommandHelpTemplate += "\n官网：https://panel.haozi.net"
	cli.RootCommandHelpTemplate += "\n论坛：https://bbs.haozi.net"
	cli.RootCommandHelpTemplate += "\nQ群：12370907\n"

	return &cli.Command{
		Name:     "panel-cli",
		Usage:    "耗子面板命令行工具",
		Version:  app.Version,
		Commands: cmd.Commands(),
	}
}
