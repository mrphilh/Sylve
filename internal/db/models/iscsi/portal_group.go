package iscsiModels

import (
	"fmt"
	"strings"
)

type PortalGroup struct {
	ID   int `gorm:"primaryKey"`
	Name string
	//TODO: Defaults don't seem to be working
	DiscoveryAuthGroup string `gorm:"default:'default'"`
	//TODO: listen shoud accept a list of addresses
	Listen *string
}

func (pg *PortalGroup) AsUcl(indent int) string {
	builder := strings.Builder{}
	for i := 0; i < indent; i++ {
		builder.WriteString("\t")
	}
	tabs := builder.String()
	builder.Reset()

	builder.WriteString(fmt.Sprintf("%s%s {\n", tabs, pg.Name))
	builder.WriteString(fmt.Sprintf("%s\tdiscovery-auth-group = %s\n", tabs, pg.DiscoveryAuthGroup))
	if pg.Listen != nil {
		builder.WriteString(fmt.Sprintf("%s\tlisten = %s\n", tabs, *pg.Listen))
	}
	builder.WriteString(fmt.Sprintf("%s}\n", tabs))

	return builder.String()
}

/* TODO: Add complete portalGroup configuration options
 portal-group	Context
       discovery-auth-group name
	       Assign a	previously defined authentication group	to the	portal
	       group,  to  be  used  for target	discovery.  By default,	portal
	       groups are assigned predefined auth-group "default", which  de-
	       nies	discovery.	Another	    predefined	   auth-group,
	       "no-authentication", may	be used	to  permit  discovery  without
	       authentication.

       discovery-filter	filter
	       Determines which	targets	are returned during discovery.	Filter
	       can    be    either   "none",   "portal",   "portal-name",   or
	       "portal-name-auth".  When set to	"none",	discovery will	return
	       all  targets  assigned  to  that	 portal	 group.	  When	set to
	       "portal", discovery will	not return targets that	cannot be  ac-
	       cessed  by  the	initiator  because  of their initiator-portal.
	       When  set  to  "portal-name",  the  check  will	include	  both
	       initiator-portal	   and	  initiator-name.     When    set   to
	       "portal-name-auth", the check  will  include  initiator-portal,
	       initiator-name,	and authentication credentials.	 The target is
	       returned	if it does not require CHAP authentication, or if  the
	       CHAP  user and secret used during discovery match those used by
	       the target.  Note that when using  "portal-name-auth",  targets
	       that  require  CHAP  authentication  will  only	be returned if
	       discovery-auth-group requires CHAP.  The	default	is "none".

       listen address
	       An IPv4 or IPv6 address and port	to listen on for incoming con-
	       nections.

       offload driver
	       Define  iSCSI  hardware	offload	 driver	 to   use   for	  this
	       portal-group.  The default is "none".

       option name value
	       The CTL-specific	port options passed to the kernel.

       redirect	address
	       IPv4  or	 IPv6 address to redirect initiators to.  When config-
	       ured, all initiators attempting to connect to portal  belonging
	       to  this	 portal-group  will get	redirected using "Target moved
	       temporarily" login response.  Redirection  happens  before  au-
	       thentication  and any initiator-name or initiator-portal	checks
	       are skipped.

       tag value
	       Unique 16-bit tag value of this portal-group.   If  not	speci-
	       fied, the value is generated automatically.

       foreign
	       Specifies  that	this  portal-group  is	listened by some other
	       host.  This host	will announce it on discovery stage, but won't
	       listen.

       dscp value
	       The DiffServ Codepoint used for sending data. The DSCP  can  be
	       set  to numeric,	or hexadecimal values directly,	as well	as the
	       well-defined "CSx" and "AFxx" codepoints.

       pcp value
	       The 802.1Q Priority CodePoint used for  sending	packets.   The
	       PCP  can	 be  set  to  a	value in the range between "0" to "7".
	       When omitted, the default for the outgoing interface is used.

*/
