package mongo

import (
	"errors"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/yyzyyyzy/goctls/config"
	"github.com/yyzyyyzy/goctls/model/mongo/generate"
	file "github.com/yyzyyyzy/goctls/util"
	"github.com/yyzyyyzy/goctls/util/pathx"
)

var (
	// VarStringSliceType describes a golang data structure name for mongo.
	VarStringSliceType []string
	// VarStringDir describes an output directory.
	VarStringDir string
	// VarBoolCache describes whether cache is enabled.
	VarBoolCache bool
	// VarStringPrefix string describes the prefix for the cache key.
	VarStringPrefix string
	// VarBoolEasy  describes whether to generate Collection Name in the code for easy declare.
	VarBoolEasy bool
	// VarStringStyle describes the style.
	VarStringStyle string
	// VarStringHome describes the goctl home.
	VarStringHome string
	// VarStringRemote describes the remote git repository.
	VarStringRemote string
	// VarStringBranch describes the git branch.
	VarStringBranch string
)

// Action provides the entry for goctl mongo code generation.
func Action(_ *cobra.Command, _ []string) error {
	tp := VarStringSliceType
	c := VarBoolCache
	p := VarStringPrefix
	easy := VarBoolEasy
	o := strings.TrimSpace(VarStringDir)
	s := VarStringStyle
	home := VarStringHome
	remote := VarStringRemote
	branch := VarStringBranch

	if len(remote) > 0 {
		repo, _ := file.CloneIntoGitHome(remote, branch)
		if len(repo) > 0 {
			home = repo
		}
	}

	if len(home) > 0 {
		pathx.RegisterGoctlHome(home)
	}

	if len(tp) == 0 {
		return errors.New("missing type")
	}

	cfg, err := config.NewConfig(s)
	if err != nil {
		return err
	}

	a, err := filepath.Abs(o)
	if err != nil {
		return err
	}

	if err = pathx.MkdirIfNotExist(a); err != nil {
		return err
	}

	return generate.Do(&generate.Context{
		Types:  tp,
		Cache:  c,
		Prefix: p,
		Easy:   easy,
		Output: a,
		Cfg:    cfg,
	})
}
