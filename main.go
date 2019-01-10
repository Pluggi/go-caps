// cmpout
// +build ignore

// Inspired from https://gist.github.com/sbz/1090868/0b190b8c222689f142242fdf92e56051d4afc6da

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Pluggi/go-caps"
)

// Generated using
// awk '$1 == "#define" &&                           \
//      $2 ~ /^CAP_\w+$/ &&                          \
//      $2 != "CAP_LAST_CAP"                         \
//      { print "\x22" tolower($2) "\x22," }' /usr/include/linux/capability.h
func main() {
	capNames := []string{
		"cap_chown",
		"cap_dac_override",
		"cap_dac_read_search",
		"cap_fowner",
		"cap_fsetid",
		"cap_kill",
		"cap_setgid",
		"cap_setuid",
		"cap_setpcap",
		"cap_linux_immutable",
		"cap_net_bind_service",
		"cap_net_broadcast",
		"cap_net_admin",
		"cap_net_raw",
		"cap_ipc_lock",
		"cap_ipc_owner",
		"cap_sys_module",
		"cap_sys_rawio",
		"cap_sys_chroot",
		"cap_sys_ptrace",
		"cap_sys_pacct",
		"cap_sys_admin",
		"cap_sys_boot",
		"cap_sys_nice",
		"cap_sys_resource",
		"cap_sys_time",
		"cap_sys_tty_config",
		"cap_mknod",
		"cap_lease",
		"cap_audit_write",
		"cap_audit_control",
		"cap_setfcap",
		"cap_mac_override",
		"cap_mac_admin",
		"cap_syslog",
		"cap_wake_alarm",
		"cap_block_suspend",
		"cap_audit_read",
	}

	pid := os.Getpid()
	c, err := caps.GetPid(pid)
	if err != nil {
		log.Fatal("GetPid: ", err)
	}

	capList := []caps.CapValue{caps.CAP_CHOWN}
	if err := c.SetFlag(caps.CAP_EFFECTIVE, capList, caps.CAP_SET); err != nil {
		log.Println("SetFlag CAP_EFFECTIVE CAP_CHOWN: ", err)
	}

	capList[0] = caps.CAP_MAC_ADMIN
	if err := c.SetFlag(caps.CAP_PERMITTED, capList, caps.CAP_SET); err != nil {
		log.Println("SetFlag CAP_PERMITTED CAP_MAC_ADMIN: ", err)
	}

	capList[0] = caps.CAP_SETFCAP
	if err := c.SetFlag(caps.CAP_INHERITABLE, capList, caps.CAP_SET); err != nil {
		log.Println("SetFlag CAP_INHERITABLE CAP_SETFCAP: ", err)
	}

	log.Println("Dumping capabilities")

	for _, capName := range capNames {
		value, err := caps.FromName(capName)
		if err != nil {
			fmt.Println("FromName ", capName, err)
		}

		fmt.Printf("%-20s\t\t", capName)
		fmt.Print("flags: ")

		cfv, err := c.GetFlag(value, caps.CAP_EFFECTIVE)
		if cfv == caps.CAP_SET {
			fmt.Printf("EFFECTIVE")
		}

		cfv, err = c.GetFlag(value, caps.CAP_PERMITTED)
		if cfv == caps.CAP_SET {
			fmt.Printf("PERMITTED")
		}

		cfv, err = c.GetFlag(value, caps.CAP_INHERITABLE)
		if cfv == caps.CAP_SET {
			fmt.Printf("INHERITABLE")
		}

		fmt.Println()
	}
}
