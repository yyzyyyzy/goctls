package cmd

import (
	_ "embed"
	"fmt"
	"os"
	"runtime"
	"strings"
	"text/template"

	"github.com/gookit/color"
	"github.com/spf13/cobra"
	"github.com/withfig/autocomplete-tools/integrations/cobra"
	"github.com/yyzyyyzy/goctls/api"
	"github.com/yyzyyyzy/goctls/bug"
	"github.com/yyzyyyzy/goctls/config"
	"github.com/yyzyyyzy/goctls/docker"
	"github.com/yyzyyyzy/goctls/env"
	"github.com/yyzyyyzy/goctls/gateway"
	"github.com/yyzyyyzy/goctls/internal/cobrax"
	"github.com/yyzyyyzy/goctls/internal/version"
	"github.com/yyzyyyzy/goctls/kube"
	"github.com/yyzyyyzy/goctls/migrate"
	"github.com/yyzyyyzy/goctls/model"
	"github.com/yyzyyyzy/goctls/quickstart"
	"github.com/yyzyyyzy/goctls/rpc"
	"github.com/yyzyyyzy/goctls/tpl"
	"github.com/yyzyyyzy/goctls/upgrade"
)

const (
	codeFailure = 1
	dash        = "-"
	doubleDash  = "--"
	assign      = "="
)

var (
	//go:embed usage.tpl
	usageTpl string
	rootCmd  = cobrax.NewCommand("goctl")
)

// Execute executes the given command
func Execute() {
	os.Args = supportGoStdFlag(os.Args)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(color.Red.Render(err.Error()))
		os.Exit(codeFailure)
	}
}

func supportGoStdFlag(args []string) []string {
	copyArgs := append([]string(nil), args...)
	parentCmd, _, err := rootCmd.Traverse(args[:1])
	if err != nil { // ignore it to let cobra handle the error.
		return copyArgs
	}

	for idx, arg := range copyArgs[0:] {
		parentCmd, _, err = parentCmd.Traverse([]string{arg})
		if err != nil { // ignore it to let cobra handle the error.
			break
		}
		if !strings.HasPrefix(arg, dash) {
			continue
		}

		flagExpr := strings.TrimPrefix(arg, doubleDash)
		flagExpr = strings.TrimPrefix(flagExpr, dash)
		flagName, flagValue := flagExpr, ""
		assignIndex := strings.Index(flagExpr, assign)
		if assignIndex > 0 {
			flagName = flagExpr[:assignIndex]
			flagValue = flagExpr[assignIndex:]
		}

		if !isBuiltin(flagName) {
			// The method Flag can only match the user custom flags.
			f := parentCmd.Flag(flagName)
			if f == nil {
				continue
			}
			if f.Shorthand == flagName {
				continue
			}
		}

		goStyleFlag := doubleDash + flagName
		if assignIndex > 0 {
			goStyleFlag += flagValue
		}

		copyArgs[idx] = goStyleFlag
	}
	return copyArgs
}

func isBuiltin(name string) bool {
	return name == "version" || name == "help"
}

func init() {
	cobra.AddTemplateFuncs(template.FuncMap{
		"blue":    blue,
		"green":   green,
		"rpadx":   rpadx,
		"rainbow": rainbow,
	})

	rootCmd.Version = fmt.Sprintf(
		"%s %s/%s", version.BuildVersion,
		runtime.GOOS, runtime.GOARCH)

	rootCmd.SetUsageTemplate(usageTpl)
	rootCmd.AddCommand(api.Cmd, bug.Cmd, docker.Cmd, kube.Cmd, env.Cmd, gateway.Cmd, model.Cmd)
	rootCmd.AddCommand(migrate.Cmd, quickstart.Cmd, rpc.Cmd, tpl.Cmd, upgrade.Cmd, config.Cmd)
	rootCmd.Command.AddCommand(cobracompletefig.CreateCompletionSpecCommand())
	rootCmd.MustInit()
}
