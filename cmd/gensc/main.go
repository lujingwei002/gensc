package main

import (
	"bufio"
	"fmt"
	"log"
	_ "net/http/pprof"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"syscall"

	"github.com/urfave/cli/v2"

	"github.com/lujingwei002/gensc"
	"github.com/lujingwei002/gensc/cmd/gensc/cmd"
	"github.com/lujingwei002/gensc/gen/gen_application"
	"github.com/lujingwei002/gensc/gen/gen_behavior"
	"github.com/lujingwei002/gensc/gen/gen_const"
	"github.com/lujingwei002/gensc/gen/gen_macro"
	"github.com/lujingwei002/gensc/gen/gen_model"
	"github.com/lujingwei002/gensc/gen/gen_protocol"
	"github.com/lujingwei002/gensc/gen/gen_resource"
)

func main() {
	cmd.Execute()

	// app := &cli.App{
	// 	Name: "gira-cli",
	// 	Authors: []*cli.Author{{
	// 		Name:  "lujingwei",
	// 		Email: "lujingwei@xx.org",
	// 	}},
	// 	Description: "gira-cli",
	// 	Flags:       []cli.Flag{},
	// 	//log.SetFlags(log.LstdFlags | log.Lshortfile)
	// 	Before: beforeAction,
	// 	Commands: []*cli.Command{
	// 		{
	// 			Name:   "gen",
	// 			Usage:  "gen [resource|model|protocol|const]",
	// 			Before: beforeAction1,
	// 			Subcommands: []*cli.Command{
	// 				{
	// 					Name:   "all",
	// 					Usage:  "gen all",
	// 					Action: genAllAction,
	// 				},
	// 				{
	// 					Name:   "resource",
	// 					Usage:  "gen resource",
	// 					Action: genResourceAction,
	// 				},
	// 				{
	// 					Name:   "model",
	// 					Usage:  "gen model",
	// 					Action: genModelAction,
	// 				},
	// 				{
	// 					Name:   "protocol",
	// 					Usage:  "gen protocol",
	// 					Action: genProtocolAction,
	// 				},
	// 				{
	// 					Name:   "const",
	// 					Usage:  "gen const",
	// 					Action: genConstAction,
	// 				},
	// 				{
	// 					Name:   "application",
	// 					Usage:  "gen application",
	// 					Action: genApplicationAction,
	// 				},
	// 				{
	// 					Name:   "behavior",
	// 					Usage:  "gen behavior",
	// 					Action: genBehaviorAction,
	// 				},
	// 			},
	// 		},
	// 		{
	// 			Name:   "macro",
	// 			Usage:  "macro code",
	// 			Action: macroAction,
	// 			Before: beforeAction1,
	// 		},
	// 	},
	// }
	// if err := app.Run(os.Args); err != nil {
	// 	log.Println(err)
	// }
}

func beforeAction1(args *cli.Context) error {
	return nil
}

func command(name string, argv []string) error {
	cmd := exec.Command(name, argv...)
	// 获取命令的标准输出管道
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return err
	}
	// 启动命令
	if err := cmd.Start(); err != nil {
		return err
	}
	// 创建一个channel，用于接收信号
	c := make(chan os.Signal, 1)
	// 监听SIGINT信号
	signal.Notify(c, os.Interrupt, syscall.SIGINT)
	defer func() {
		signal.Reset(os.Interrupt, syscall.SIGINT)
	}()
	// 创建一个 Scanner 对象，对命令的标准输出和标准错误输出进行扫描
	scanner1 := bufio.NewScanner(stdout)
	go func() {
		for scanner1.Scan() {
			// 输出命令的标准输出
			log.Println(scanner1.Text())
		}
	}()
	scanner2 := bufio.NewScanner(stderr)
	go func() {
		for scanner2.Scan() {
			// 输出命令的标准错误输出
			fmt.Fprintln(os.Stderr, scanner2.Text())
		}
	}()
	go func() {
		// 等待信号
		<-c
	}()
	// 等待命令执行完成
	if err := cmd.Wait(); err != nil {
		return err
	}
	return nil
}

func beforeAction(args *cli.Context) error {
	log.Println("*************", strings.Join(os.Args, " "), "***************")
	return nil
}

func genResourceAction(args *cli.Context) error {
	if err := gen_resource.Gen(gensc.GenResourceConfig{
		Force: true,
	}); err != nil {
		return err
	}
	return nil
}

func genProtocolAction(args *cli.Context) error {
	if err := gen_protocol.Gen(gensc.GenProtocolConfig{}); err != nil {
		return err
	}
	return nil
}

func genModelAction(args *cli.Context) error {
	if err := gen_model.Gen(gensc.GenModelConfig{}); err != nil {
		return err
	}
	return nil
}

func genApplicationAction(args *cli.Context) error {
	if err := gen_application.Gen(gensc.GenApplicationConfig{}); err != nil {
		return err
	}
	return nil
}

func genConstAction(args *cli.Context) error {
	if err := gen_const.Gen(gensc.GenConstConfig{}); err != nil {
		return err
	}
	return nil
}

func genBehaviorAction(args *cli.Context) error {
	if err := gen_behavior.Gen(gensc.GenBehaviorConfig{}); err != nil {
		return err
	}
	return nil
}

func genAllAction(args *cli.Context) error {
	if err := gen_protocol.Gen(gensc.GenProtocolConfig{}); err != nil {
		return err
	}
	if err := gen_application.Gen(gensc.GenApplicationConfig{}); err != nil {
		return err
	}
	if err := gen_resource.Gen(gensc.GenResourceConfig{
		Force: true,
	}); err != nil {
		return err
	}
	if err := gen_const.Gen(gensc.GenConstConfig{}); err != nil {
		return err
	}
	if err := gen_model.Gen(gensc.GenModelConfig{}); err != nil {
		return err
	}
	return nil
}

func resourceAction(args *cli.Context) error {
	log.Println("resourceAction")
	return nil
}

func macroAction(args *cli.Context) error {
	log.Println("macroAction")
	if err := gen_macro.Gen(gensc.MacroConfig{}); err != nil {
		return err
	}
	return nil
}
