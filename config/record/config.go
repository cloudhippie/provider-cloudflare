package record

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudflare_record", func(r *config.Resource) {
		r.ShortGroup = "record"
	})
}
