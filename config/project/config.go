package project

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
    p.AddResourceConfigurator("gitlab_project", func(r *config.Resource) {
        // We need to override the default group that upjet generated for
        // this resource, which would be "gitlab"
        r.ShortGroup = "project"
    })
}
