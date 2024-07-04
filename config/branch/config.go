package branch

import "github.com/crossplane/upjet/pkg/config"

func Configure(p *config.Provider) {
    p.AddResourceConfigurator("gitlab_branch", func(r *config.Resource) {
        // We need to override the default group that upjet generated for
        // this resource, which would be "github"
        r.ShortGroup = "branch"

        // This resource need the repository in which branch would be created
        // as an input. And by defining it as a reference to Repository
        // object, we can build cross resource referencing. See
        // repositoryRef in the example in the Testing section below.
        r.References["project"] = config.Reference{
            Type: "github.com/eperdeme/provider-github/apis/project/v1alpha1.Project",
        }
    })
}
