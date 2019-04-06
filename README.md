# Go-Command

A command line tool that helps you define new commands on the fly.

## Installation

1. Select a release that that you want to use (can be found [here](https://gitlab.com/TimSatke/goc/tags)).
2. Click on "Download artifacts".
3. Download the correct executable based on your computer's OS and architecture.
4. Move the downloaded file to a folder that is in your $PATH, e.g. `~/Development/tools` (if you want to use another directory, see section `Configuration`).
5. Rename the file to a name you would like to use. Ideally, this would be `goc`.
6. Done.

## How to use

To display help, use `goc help` or `goc help <command>`.

Let's say, we want to remove all local git branches that don't have a corresponding remote anymore.
This can be done with the command `git fetch -p && for branch in ``git branch -vv | grep ': gone]' | awk '{print $1}'``; do git branch -D $branch; done`.
I personally don't want to have to type this every time.

With `goc`, we can now type `goc define git_remove_gone` and hit return.
By default, `goc` will try to open `vim`, where you can type the command from above.
Hit that `:wq` and return.
`goc` will confirm, that you now can use the command `git_remove_gone`.
When you type `git_remove_gone` (and (by default) `~/Development/tools` is in your path), the entered command will be executed.

## Configuration

If you want `goc` to place the generated commands somewhere else (not `~/Development/tools`), type
`goc prefs set cmd.output.directory <other directory>` and hit return.
The changes will apply instantly, but the already existing/created commands will **NOT** be moved.
You either have to move them manually or re-create them.

### Used keys
* `cmd.output.directory` is the directory, where generated commands will be placed. (default: `~/Development/tools`)
* `cmd.define.editor` is the editor process, that will be launched. It will be given a path to a generated command. E.g. if this is set to `vim` and you want to create a command `greeting`, the `goc` will launch `vim ~/Development/tools/greeting`. (default: `vim`)
* `cmd.undefine.prompt` will make `goc` prompt you, if you are about to delete a command with `goc undefine <command>`. (default: `true`)
* `cmd.prefs.clear.prompt` will make `goc` prompt you, if you are about to delete the configuration file with `goc prefs clear`. (default: `true`)

`goc` sub-commands sometimes have some aliases, e.g. `goc prefs clear` is the same as `goc prefs reset`. It does not matter which one you use, use the one that is most comfortable to type.
You can find out, what aliases a command has from the help text.
```
$> goc help prefs
The base command for preference commands.

Usage:
  goc prefs [command]

Aliases:
  prefs, preferences, settings, opts, options

Available Commands:
  clear       Resets all preferences to default.
  get         Prints a preference by key.
  list        List all defined preferences.
  set         Set a preference key to a given value.

Flags:
  -h, --help   help for prefs

Use "goc prefs [command] --help" for more information about a command.
```