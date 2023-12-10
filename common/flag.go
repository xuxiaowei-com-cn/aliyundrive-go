package common

import "github.com/urfave/cli/v2"

func AppIdFlag(required bool) cli.Flag {
	return &cli.StringFlag{
		Name:     AppId,
		Aliases:  []string{ClientId},
		Required: required,
		Usage:    "阿里云盘 APP ID、Client Id，网址：https://www.aliyundrive.com/developer/f",
	}
}

func AppSecretFlag(required bool) cli.Flag {
	return &cli.StringFlag{
		Name:     AppSecret,
		Aliases:  []string{ClientSecret},
		Required: required,
		Usage:    "阿里云盘 App Secret、Client Secret，网址：https://www.aliyundrive.com/developer/f",
	}
}

func ScopeFlag(required bool) cli.Flag {
	return &cli.StringSliceFlag{
		Name:     Scope,
		Required: required,
		Usage:    "权限列表，网址：https://www.yuque.com/aliyundrive/zpfszx/dspik0",
	}
}

func RedirectUriFlag(required bool) cli.Flag {
	return &cli.StringFlag{
		Name:     RedirectUri,
		Required: required,
		Usage:    "网址：https://www.yuque.com/aliyundrive/zpfszx/eam8ls1lmawwwksv",
	}
}

func AccessTokenFlag(required bool) cli.Flag {
	return &cli.StringFlag{
		Name:     AccessToken,
		Required: required,
		Usage:    "访问凭证，网址：https://www.yuque.com/aliyundrive/zpfszx/btw0tw",
	}
}
func WidthFlag() cli.Flag {
	return &cli.IntFlag{
		Name:  Width,
		Value: 430,
		Usage: "二维码宽度，网址：https://www.yuque.com/aliyundrive/zpfszx/ttfoy0xt2pza8lof",
	}
}

func HeightFlag() cli.Flag {
	return &cli.IntFlag{
		Name:  Height,
		Value: 430,
		Usage: "二维码高度，网址：https://www.yuque.com/aliyundrive/zpfszx/ttfoy0xt2pza8lof",
	}
}

func HttpListenAddrFlag() cli.Flag {
	return &cli.StringFlag{
		Name:  HttpListenAddr,
		Value: ":8080",
		Usage: "HTTP 监听地址",
	}
}
