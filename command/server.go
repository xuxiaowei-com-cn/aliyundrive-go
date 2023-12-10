package command

import (
	"github.com/urfave/cli/v2"
	"github.com/xuxiaowei-com-cn/aliyundrive-go/action"
	"github.com/xuxiaowei-com-cn/aliyundrive-go/common"
)

func ServerCommand() *cli.Command {
	return &cli.Command{
		Name:    "server",
		Aliases: []string{"s"},
		Usage:   "服务端命令",
		Flags: []cli.Flag{common.AppIdFlag(false), common.AppSecretFlag(false), common.ScopeFlag(false),
			common.WidthFlag(), common.HeightFlag(), common.HttpListenAddrFlag()},
		Subcommands: []*cli.Command{
			QrCodeCommand(),
		},
	}
}

func QrCodeCommand() *cli.Command {
	return &cli.Command{
		Name:  "qrcode",
		Usage: "手机扫码授权模式",
		Flags: []cli.Flag{common.AppIdFlag(true), common.AppSecretFlag(true), common.ScopeFlag(true),
			common.WidthFlag(), common.HeightFlag(), common.HttpListenAddrFlag()},
		Action: func(context *cli.Context) error {

			var appId = context.String(common.AppId)
			var appSecret = context.String(common.AppSecret)
			var scope = context.StringSlice(common.Scope)
			var width = context.Int(common.Width)
			var height = context.Int(common.Height)
			var httpListenAddr = context.String(common.HttpListenAddr)

			return action.QrCodeAction(appId, appSecret, scope, width, height, httpListenAddr)
		},
	}
}
