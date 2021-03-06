/*
Sniperkit-Bot
- Status: analyzed
*/

// Copyright © 2015 Erik Brady <brady@dvln.org>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmds

import (
	cli "github.com/dvln/cobra"
	"github.com/dvln/lib"
	"github.com/dvln/out"
)

var versionCmd = &cli.Command{
	Use:   "version",
	Short: "get the current version of the 'dvln' tool",
	Long: `Get the current release version number of the 'dvln' tool:
  % dvln version
  Note: version can also be dumped directly from the dvln tool:
  % dvln --version
  % dvln -V`,
	Run: version,
}

// init bootstraps the options used for the version subcommand and descriptions
// and initial defaults for those options and such.
func init() {
	reloadCLIFlags := false
	setupVersionCmdCLIArgs(versionCmd, reloadCLIFlags)
}

// setupVersionCmdCLIArgs is used from init() to set up the 'globs' (viper) pkg
// CLI options available to this subcommand (other options were already set up
// in the "parent" dvln subcommand in a like-named method, every subcommand has
// a like named method "setup<subcmd>CmdCLIArgs()"
func setupVersionCmdCLIArgs(c *cli.Command, reloadCLIFlags bool) {
	if reloadCLIFlags {
		c.Flags().SetDefValueReparseOK(true)
	}

	// NewCLIOpts: if there were opts for this subcmd set them here, see
	// cmds/get.go for an example.  Note that "persistent" opts are set in
	// cmds/dvln.go, only opts specific to the 'dvln version' subcommand
	// would go here
	// Note that you'll need to modify cmds/global.go as well otherwise your
	// globs.Desc() call and globs.GetBool("myopt") will not work.

	c.Run = version
	if reloadCLIFlags {
		c.Flags().SetDefValueReparseOK(false)
	}
}

// version is the function executed by 'dvln version' assuming all opts are
// validated as OK and such, currently it prints out the tool version and
// relies upon the library to format the version based on the options
// selected (eg: terse, regular, verbose with text or json as the "look")
func version(cmd *cli.Command, args []string) {
	dvlnVerStr := lib.DvlnVerStr(commitSHA1)
	out.Print(dvlnVerStr)
}
