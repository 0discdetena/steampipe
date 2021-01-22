package constants

import "fmt"

// Metaquery commands
const (
	CmdTableList   = ".tables"      // List all tables
	CmdOutput      = ".output"      // Set output mode
	CmdTiming      = ".timing"      // Toggle query timer
	CmdHeaders     = ".header"      // Toggle headers output
	CmdSeparator   = ".separator"   // Set the column separator
	CmdExit        = ".exit"        // Exit the interactive prompt
	CmdQuit        = ".quit"        // Alias for .exit
	CmdInspect     = ".inspect"     // inspect
	CmdConnections = ".connections" // list all connections
	CmdMulti       = ".multi"       // toggle multi line query
	CmdClear       = ".clear"       // clear the console
)

// ArgFromMetaquery :: convert metaquery of form '.header' into the config argument used to set the mode, i.e. 'header'
func ArgFromMetaquery(cmd string) string {
	if cmd[:1] != "." {
		panic(fmt.Sprintf("ArgFromMetaquery called for non-metyaquery: %s", cmd))
	}
	return cmd[1:]
}
