package plugin

import (
	"fmt"
	"github.com/bsthun/gut"
	"github.com/strawst/strawhouse-go"
	"os"
	"path/filepath"
	"plugin"
)

func (r *Service) Init() {
	entries, err := os.ReadDir(*r.config.PluginPath)
	if err != nil {
		gut.Fatal("unable to read plugin directory", err)
	}

	for _, entry := range entries {
		if !entry.IsDir() && filepath.Ext(entry.Name()) == ".so" {
			fmt.Printf("[plugin] Loaded %s\n", entry.Name())
			p, err := plugin.Open(filepath.Join(*r.config.PluginPath, entry.Name()))
			if err != nil {
				gut.Fatal("unable to open plugin", err)
			}
			initializer, err := p.Lookup("Plugin")
			if err != nil {
				gut.Fatal("unable to lookup plugin", err)
			}
			plugInitializer := initializer.(func() strawhouse.Plugin)
			plug := plugInitializer()
			plug.Load(r.Plugger())
			r.s.Plugins[entry.Name()] = plug
		}
	}
}
