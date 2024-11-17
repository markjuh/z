package main

import "github.com/rwxrob/bonzai"
import "github.com/rwxrob/bonzai/comp"

// Commands imported from Bonzai.
import "github.com/rwxrob/bonzai/cmds/help"
import "github.com/rwxrob/bonzai/cmds/kimono"
import "github.com/rwxrob/bonzai/cmds/rayfish"
import "github.com/rwxrob/bonzai/cmds/sunrise"

// My own commands.
import "github.com/markjuh/z/cmds/datedir"

var Cmd = &bonzai.Cmd{
  Name: `z`,
  Short: `command tree`,
  Comp: comp.Cmds,
  Def: help.Cmd,
  Cmds: []*bonzai.Cmd{
    datedir.Cmd,
    help.Cmd,
    kimono.Cmd,
    rayfish.Cmd,
    sunrise.Cmd,
  },
}
