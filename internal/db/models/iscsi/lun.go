package iscsiModels

import "fmt"
import "strings"

type Lun struct {
	ID   int `gorm:"primaryKey"`
	Name string	// TODO: Name should be something automatic
	Path string
	Size string
}

func (l *Lun) AsUcl(indent int) string {
	builder := strings.Builder{}
	for i := 0; i < indent; i++ {
		builder.WriteString("\t")
	}
	tabs := builder.String()
	builder.Reset()
	builder.WriteString(fmt.Sprintf("%s%s {\n", tabs, l.Name))
	builder.WriteString(fmt.Sprintf("%s\tpath = %s\n", tabs, l.Path))
	builder.WriteString(fmt.Sprintf("%s\tsize = %s\n", tabs, l.Size))
	builder.WriteString(fmt.Sprintf("%s}\n", tabs))
	return builder.String()
}
/* TODO: Add complete LUN configuration
   lun Context
       backend block | ramdisk
	       The  CTL	 backend  to  use  for a given LUN.  Valid choices are
	       "block" and "ramdisk"; block is used for	LUNs backed  by	 files
	       or  disk	device nodes; ramdisk is a bitsink device, used	mostly
	       for testing.  The default backend is block.

       blocksize size
	       The blocksize visible to	the initiator.	The default  blocksize
	       is 512 for disks, and 2048 for CD/DVDs.

       ctl-lun lun_id
	       Global  numeric	identifier  to use for a given LUN inside CTL.
	       By default CTL allocates	those IDs  dynamically,	 but  explicit
	       specification  may  be  needed for consistency in HA configura-
	       tions.

       device-id string
	       The SCSI	Device Identification string presented to  iSCSI  ini-
	       tiators.

       device-type type
	       Specify	the  SCSI  device  type	 to use	when creating the LUN.
	       Currently CTL supports Direct Access (type 0), Processor	 (type
	       3) and CD/DVD (type 5) LUNs.

       option name value
	       The  CTL-specific  options  passed to the kernel.  All CTL-spe-
	       cific options  are  documented  in  the	"OPTIONS"  section  of
	       ctladm(8).

       path path
	       The  path  to  the  file, device	node, or zfs(8)	volume used to
	       back the	LUN.  For optimal performance, create ZFS volumes with
	       the "volmode=dev" property set.

       serial string
	       The SCSI	serial number presented	to iSCSI initiators.

       size size
	       The LUN size, in	bytes or by number with	a suffix of K, M, G, T
	       (for kilobytes, megabytes, gigabytes, or	terabytes).  When  the
	       configuration   is   in	UCL  format,  use  the	suffix	format
	       kKmMgG|bB, (i.e., 4GB, 4gb, and 4Gb are all equivalent).
*/
