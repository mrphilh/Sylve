package iscsiModels

import (
	"fmt"
	"strings"
)

type Target struct {
	ID            int    `gorm:"primaryKey"`
	Name          string // iqn qualified string
	Alias         *string
	AuthGroupID   *int
	AuthGroup     AuthGroup
	PortalGroupID *int
	PortalGroup   PortalGroup
	Luns          []Lun
}

func (t *Target) AsUcl(indent int) string {
	builder := strings.Builder{}
	for i := 0; i < indent; i++ {
		builder.WriteString("\t")
	}
	tabs := builder.String()
	builder.Reset()
	builder.WriteString(fmt.Sprintf("%s\"%s\" {\n", tabs, t.Name))
	if t.Alias != nil {
		builder.WriteString(fmt.Sprintf("%s\talias = \"%s\"\n", tabs, *t.Alias))
	}
	builder.WriteString(fmt.Sprintf("%s\tauth-group = %s\n", tabs, t.AuthGroup.Name))
	builder.WriteString(fmt.Sprintf("%s\tportal-group = %s\n", tabs, t.PortalGroup.Name))
	builder.WriteString(fmt.Sprintf("%s\tlun = {\n", tabs))
	for i := 0; i < len(t.Luns); i++ {
		builder.WriteString(fmt.Sprintf("%s\t\t%d {\n", tabs, i))
		builder.WriteString(fmt.Sprintf("%s\t\t\tpath = %s\n", tabs, t.Luns[i].Path))
		builder.WriteString(fmt.Sprintf("%s\t\t\tsize = %s\n", tabs, t.Luns[i].Size))
		builder.WriteString(fmt.Sprintf("%s\t\t}\n", tabs))
	}
	builder.WriteString(fmt.Sprintf("%s\t}\n", tabs))
	builder.WriteString(fmt.Sprintf("%s}\n", tabs))
	return builder.String()
}

/* TODO: Complete supported configuration options
   target Context
       alias text
	       Assign a	human-readable description to the target.  There is no
	       default.

       auth-group name
	       Assign a	previously defined authentication group	to the target.
	       By default, targets that	do not specify	their  own  auth  set-
	       tings,  using  clauses  such as chap or initiator-name, are as-
	       signed predefined auth-group "default", which  denies  all  ac-
	       cess.   Another predefined auth-group, "no-authentication", may
	       be used to permit access	 without  authentication.   Note  that
	       this  clause  can  be overridden	using the second argument to a
	       portal-group clause.

       auth-type type
	       Sets the	authentication	type.	Type  can  be  either  "none",
	       "deny", "chap", or "chap-mutual".  In most cases	it is not nec-
	       essary to set the type using this clause; it is usually used to
	       disable	authentication for a given target.  This clause	is mu-
	       tually exclusive	with auth-group; one cannot use	both in	a sin-
	       gle target.

       chap user secret
	       A set of	CHAP authentication credentials.   Note	 that  targets
	       must  only use one of auth-group, chap, or chap-mutual; it is a
	       configuration error to mix multiple types in one	target.

       chap-mutual user	secret mutualuser mutualsecret
	       A set of	mutual CHAP  authentication  credentials.   Note  that
	       targets	must only use one of auth-group, chap, or chap-mutual;
	       it is a configuration error to mix multiple types in  one  tar-
	       get.

       initiator-name initiator-name
	       An  iSCSI initiator name.  Only initiators with a name matching
	       one of the defined names	will be	allowed	to  connect.   If  not
	       defined,	there will be no restrictions based on initiator name.
	       This  clause  is	mutually exclusive with	auth-group; one	cannot
	       use both	in a single target.

       initiator-portal	address[/prefixlen]
	       An iSCSI	initiator portal: an IPv4 or IPv6 address,  optionally
	       followed	 by a literal slash and	a prefix length.  Only initia-
	       tors with an address matching one of the	defined	addresses will
	       be allowed to connect.  If not defined, there will  be  no  re-
	       strictions based	on initiator address.  This clause is mutually
	       exclusive with auth-group; one cannot use both in a single tar-
	       get.

	       The   auth-type,	  chap,	  chap-mutual,	 initiator-name,   and
	       initiator-portal	clauses	in the target context provide  an  al-
	       ternative to assigning an auth-group defined separately,	useful
	       in  the	common	case  of authentication	settings specific to a
	       single target.

       portal-group name [ag-name]
	       Assign a	previously defined portal group	to  the	 target.   The
	       default	portal	group  is  "default",  which  makes the	target
	       available on TCP	port 3260 on all configured IPv4 and IPv6  ad-
	       dresses.	  Optional  second  argument  specifies	auth-group for
	       connections to this specific portal group.  If second  argument
	       is not specified, target	auth-group is used.

       port name

       port name/pp

       port name/pp/vp
	       Assign  specified  CTL port (such as "isp0" or "isp2/1")	to the
	       target.	This is	used to	export the target through  a  specific
	       physical	 -  eg	Fibre  Channel	- port,	in addition to portal-
	       groups configured for the target.  Use ctladm portlist  command
	       to  retrieve  the  list of available ports.  On startup ctld(8)
	       configures LUN mapping and enables all  assigned	 ports.	  Each
	       port can	be assigned to only one	target.

       redirect	address
	       IPv4  or	 IPv6 address to redirect initiators to.  When config-
	       ured, all initiators attempting to connect to this target  will
	       get redirected using "Target moved temporarily" login response.
	       Redirection happens after successful authentication.

       lun number name
	       Export previously defined lun by	the parent target.

       lun number
	       Create  a lun configuration context, defining a LUN exported by
	       the parent target.

	       This is an alternative to defining the LUN  separately,	useful
	       in the common case of a LUN being exported by a single target.

*/
