package api

import (
	"github.com/spf13/cobra"

	"github.com/yyzyyyzy/goctls/api/apigen"
	"github.com/yyzyyyzy/goctls/api/dartgen"
	"github.com/yyzyyyzy/goctls/api/docgen"
	"github.com/yyzyyyzy/goctls/api/format"
	"github.com/yyzyyyzy/goctls/api/gogen"
	"github.com/yyzyyyzy/goctls/api/javagen"
	"github.com/yyzyyyzy/goctls/api/ktgen"
	"github.com/yyzyyyzy/goctls/api/new"
	"github.com/yyzyyyzy/goctls/api/swagger"
	"github.com/yyzyyyzy/goctls/api/tsgen"
	"github.com/yyzyyyzy/goctls/api/validate"
	"github.com/yyzyyyzy/goctls/config"
	"github.com/yyzyyyzy/goctls/internal/cobrax"
	"github.com/yyzyyyzy/goctls/plugin"
)

var (
	// Cmd describes an api command.
	Cmd       = cobrax.NewCommand("api", cobrax.WithRunE(apigen.CreateApiTemplate))
	dartCmd   = cobrax.NewCommand("dart", cobrax.WithRunE(dartgen.DartCommand))
	docCmd    = cobrax.NewCommand("doc", cobrax.WithRunE(docgen.DocCommand))
	formatCmd = cobrax.NewCommand("format", cobrax.WithRunE(format.GoFormatApi))
	goCmd     = cobrax.NewCommand("go", cobrax.WithRunE(gogen.GoCommand))
	newCmd    = cobrax.NewCommand("new", cobrax.WithRunE(new.CreateServiceCommand),
		cobrax.WithArgs(cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs)))
	validateCmd = cobrax.NewCommand("validate", cobrax.WithRunE(validate.GoValidateApi))
	javaCmd     = cobrax.NewCommand("java", cobrax.WithRunE(javagen.JavaCommand), cobrax.WithHidden())
	ktCmd       = cobrax.NewCommand("kt", cobrax.WithRunE(ktgen.KtCommand))
	pluginCmd   = cobrax.NewCommand("plugin", cobrax.WithRunE(plugin.PluginCommand))
	tsCmd       = cobrax.NewCommand("ts", cobrax.WithRunE(tsgen.TsCommand))
	swaggerCmd  = cobrax.NewCommand("swagger", cobrax.WithRunE(swagger.Command))
)

func init() {
	var (
		apiCmdFlags      = Cmd.Flags()
		dartCmdFlags     = dartCmd.Flags()
		docCmdFlags      = docCmd.Flags()
		formatCmdFlags   = formatCmd.Flags()
		goCmdFlags       = goCmd.Flags()
		javaCmdFlags     = javaCmd.Flags()
		ktCmdFlags       = ktCmd.Flags()
		newCmdFlags      = newCmd.Flags()
		pluginCmdFlags   = pluginCmd.Flags()
		tsCmdFlags       = tsCmd.Flags()
		validateCmdFlags = validateCmd.Flags()
		swaggerCmdFlags  = swaggerCmd.Flags()
	)

	apiCmdFlags.StringVar(&apigen.VarStringOutput, "o")
	apiCmdFlags.StringVar(&apigen.VarStringHome, "home")
	apiCmdFlags.StringVar(&apigen.VarStringRemote, "remote")
	apiCmdFlags.StringVar(&apigen.VarStringBranch, "branch")

	dartCmdFlags.StringVar(&dartgen.VarStringDir, "dir")
	dartCmdFlags.StringVar(&dartgen.VarStringAPI, "api")
	dartCmdFlags.BoolVar(&dartgen.VarStringLegacy, "legacy")
	dartCmdFlags.StringVar(&dartgen.VarStringHostname, "hostname")
	dartCmdFlags.StringVar(&dartgen.VarStringScheme, "scheme")

	docCmdFlags.StringVar(&docgen.VarStringDir, "dir")
	docCmdFlags.StringVar(&docgen.VarStringOutput, "o")

	formatCmdFlags.StringVar(&format.VarStringDir, "dir")
	formatCmdFlags.BoolVar(&format.VarBoolIgnore, "iu")
	formatCmdFlags.BoolVar(&format.VarBoolUseStdin, "stdin")
	formatCmdFlags.BoolVar(&format.VarBoolSkipCheckDeclare, "declare")

	goCmdFlags.StringVar(&gogen.VarStringDir, "dir")
	goCmdFlags.StringVar(&gogen.VarStringAPI, "api")
	goCmdFlags.StringVar(&gogen.VarStringHome, "home")
	goCmdFlags.StringVar(&gogen.VarStringRemote, "remote")
	goCmdFlags.StringVar(&gogen.VarStringBranch, "branch")
	goCmdFlags.BoolVar(&gogen.VarBoolWithTest, "test")
	goCmdFlags.BoolVar(&gogen.VarBoolTypeGroup, "type-group")
	goCmdFlags.StringVarWithDefaultValue(&gogen.VarStringStyle, "style", config.DefaultFormat)

	javaCmdFlags.StringVar(&javagen.VarStringDir, "dir")
	javaCmdFlags.StringVar(&javagen.VarStringAPI, "api")

	ktCmdFlags.StringVar(&ktgen.VarStringDir, "dir")
	ktCmdFlags.StringVar(&ktgen.VarStringAPI, "api")
	ktCmdFlags.StringVar(&ktgen.VarStringPKG, "pkg")

	newCmdFlags.StringVar(&new.VarStringHome, "home")
	newCmdFlags.StringVar(&new.VarStringRemote, "remote")
	newCmdFlags.StringVar(&new.VarStringBranch, "branch")
	newCmdFlags.StringVarWithDefaultValue(&new.VarStringStyle, "style", config.DefaultFormat)

	pluginCmdFlags.StringVarP(&plugin.VarStringPlugin, "plugin", "p")
	pluginCmdFlags.StringVar(&plugin.VarStringDir, "dir")
	pluginCmdFlags.StringVar(&plugin.VarStringAPI, "api")
	pluginCmdFlags.StringVar(&plugin.VarStringStyle, "style")

	tsCmdFlags.StringVar(&tsgen.VarStringDir, "dir")
	tsCmdFlags.StringVar(&tsgen.VarStringAPI, "api")
	tsCmdFlags.StringVar(&tsgen.VarStringCaller, "caller")
	tsCmdFlags.BoolVar(&tsgen.VarBoolUnWrap, "unwrap")

	swaggerCmdFlags.StringVar(&swagger.VarStringAPI, "api")
	swaggerCmdFlags.StringVar(&swagger.VarStringDir, "dir")
	swaggerCmdFlags.StringVar(&swagger.VarStringFilename, "filename")
	swaggerCmdFlags.BoolVar(&swagger.VarBoolYaml, "yaml")

	validateCmdFlags.StringVar(&validate.VarStringAPI, "api")

	// Add sub-commands
	Cmd.AddCommand(dartCmd, docCmd, formatCmd, goCmd, javaCmd, ktCmd, newCmd, pluginCmd, tsCmd, validateCmd, swaggerCmd)
}
