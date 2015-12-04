package main

import (
	"fmt"
	"strings"
)

type Server struct {
	Name string
}

type Servers []Server

func ListServers() Servers {
	return []Server{
		{Name: "Server1"},
		{Name: "Server2"},
		{Name: "Foo1"},
		{Name: "Foo2"},
	}
}

func (s Servers) Filter(name string) Servers {
	filtered := make(Servers, 0)
	for _, server := range s {
		if strings.Contains(server.Name, name) {
			filtered = append(filtered, server)
		}
	}

	return filtered
}

func main() {
	servers := ListServers()
	servers = servers.Filter("Foo")
	fmt.Printf("servers %+v\n", servers)
}
