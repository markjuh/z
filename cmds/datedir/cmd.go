package datedir

import (
  "fmt"
  "time"
  "os"

  "github.com/rwxrob/bonzai"
  "github.com/rwxrob/bonzai/cmds/help"
  "github.com/rwxrob/bonzai/comp"
  "github.com/rwxrob/term"
)

var Cmd = &bonzai.Cmd{
  Name: `datedir`,
  Vers: `v0.0.1`,
  Short: `Make, remove and switch to directory for current date`,
  Alias: `d`,
  Long: `
    The {{aka .}} command can create and remove a directory for the
    current date. Future functionality should include switching to the
    directory. This depends on implementing shell-hooks.`,
  Comp: comp.Cmds,
  Cmds: []*bonzai.Cmd{
    cdCmd,
    mkdirCmd,
    rmdirCmd,
    help.Cmd,
  },
  Def: help.Cmd,
}

var cdCmd = &bonzai.Cmd{
  Name: `cd`,
  Comp: comp.Cmds,
  Do: func(cmd *bonzai.Cmd, _ ...string) error {
    term.Print(fmt.Sprintf("%s %s", cmd.Name, currentDate()))

    return os.Chdir(currentDate())

    return nil
  },
}

var mkdirCmd = &bonzai.Cmd{
  Name: `mkdir`,
  Alias: `mk`,
  Comp: comp.Cmds,
  Short: `Create directory for the current date`,
  Do: func(cmd *bonzai.Cmd, _ ...string) error {
    term.Print(fmt.Sprintf("%s %s", cmd.Name, currentDate()))

    return os.Mkdir(currentDate(), 0755)
  },
}

var rmdirCmd = &bonzai.Cmd{
  Name: `rmdir`,
  Alias: `rm`,
  Short: `Remove directory for the current date`,
  Comp: comp.Cmds,
  Do: func(cmd *bonzai.Cmd, _ ...string) error {
    term.Print(fmt.Sprintf("%s %s", cmd.Name, currentDate()))

    return os.Remove(currentDate())
  },
}

func currentDate() string {
  currentTime := time.Now()
  formattedDate := currentTime.Format("20060102")

  return formattedDate
}

