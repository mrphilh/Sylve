package iscsiModels

import (
	"fmt"
	"strings"
)

type GlobalSetting struct {
	ID    int `gorm:"primaryKey"`
	Debug int `gorm:"default:1"`
}

func (gs *GlobalSetting) AsUcl(indent int) string {
	builder := strings.Builder{}
	for i := 0; i < indent; i++ {
		builder.WriteString("\t")
	}
	tabs := builder.String()
	builder.Reset()
	builder.WriteString(fmt.Sprintf("%sdebug = %d\n", tabs, gs.Debug))
	return builder.String()
}

/*

       debug level
	       The debug verbosity level.  The default is 0.

       maxproc number
	       The limit for concurrently running child	processes handling in-
	       coming connections.  The	default	is 30.	A setting  of  0  dis-
	       ables the limit.

       pidfile path
	       The path	to the pidfile.	 The default is	/var/run/ctld.pid.



       timeout seconds
	       The timeout for login sessions, after which the connection will
	       be  forcibly  terminated.   The	default	is 60.	A setting of 0
	       disables	the timeout.

       isns-server address
	       An IPv4 or IPv6 address and optionally port of iSNS  server  to
	       register	on.

       isns-period seconds
	       iSNS  registration  period.   Registered	Network	Entity not up-
	       dated during this period	will be	unregistered.  The default  is
	       900.

       isns-timeout seconds
	       Timeout for iSNS	requests.  The default is 5.
*/
