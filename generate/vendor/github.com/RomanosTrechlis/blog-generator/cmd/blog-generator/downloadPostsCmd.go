package main

import (
	"flag"
	"github.com/RomanosTrechlis/blog-generator/datasource"
)

const downloadPostsShortHelp = `Downloads posts from given datasource`
const downloadPostsLongHelp = `
Downloads posts from a given datasource inside the config.json file.

The folowing part of config.json controls the behavior of "fetch-posts"
command.

"DataSource": {
    "Type": "git",
    "Repository": "https://github.com/RomanosTrechlis/blog.git"
},
"TempFolder": "./tmp"

The "Type" can also be "local" and the "Repository" local folder.
The "TempFolder" is were the posts will be cloned for generation.
`

type downloadPostCmd struct {
	sourceType, source, destination string
}

func (cmd *downloadPostCmd) Name() string      { return "fetch-posts" }
func (cmd *downloadPostCmd) Args() string      { return "" }
func (cmd *downloadPostCmd) ShortHelp() string { return downloadPostsShortHelp }
func (cmd *downloadPostCmd) LongHelp() string  { return downloadPostsLongHelp }
func (cmd *downloadPostCmd) Hidden() bool      { return false }

func (cmd *downloadPostCmd) Register(fs *flag.FlagSet) {
}

func (cmd *downloadPostCmd) Run(ctx *Ctx, args []string) error {
	var err error
	switch cmd.sourceType {
	case "git":
		ds := datasource.NewGitDataSource()
		_, err = ds.Fetch(cmd.source,
			cmd.destination)
	case "local":
		ds := datasource.NewLocalDataSource()
		_, err = ds.Fetch(cmd.source,
			cmd.destination)
	case "":
		ctx.Err.Fatal("please provide a datasource in the configuration file")
	}

	return err
}
