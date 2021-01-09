package Editor

import "github.com/txn2/txeh"

func HostEditor(domain string) {
	hosts, err := txeh.NewHostsDefault()
	if err != nil {
		panic(err)
	}
	hosts.AddHost("127.0.0.1", domain)
	hosts.Save()
}
