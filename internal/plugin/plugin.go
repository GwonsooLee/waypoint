package plugin

import (
	"github.com/hashicorp/waypoint/internal/factory"
	"github.com/hashicorp/waypoint/sdk"
	"github.com/hashicorp/waypoint/sdk/component"

	"github.com/hashicorp/waypoint/builtin/aws/alb"
	"github.com/hashicorp/waypoint/builtin/aws/ami"
	"github.com/hashicorp/waypoint/builtin/aws/ec2"
	"github.com/hashicorp/waypoint/builtin/aws/ecr"
	"github.com/hashicorp/waypoint/builtin/aws/ecs"
	"github.com/hashicorp/waypoint/builtin/azure/aci"
	"github.com/hashicorp/waypoint/builtin/docker"
	"github.com/hashicorp/waypoint/builtin/files"
	"github.com/hashicorp/waypoint/builtin/google/cloudrun"
	"github.com/hashicorp/waypoint/builtin/guide/golang"
	"github.com/hashicorp/waypoint/builtin/k8s"
	"github.com/hashicorp/waypoint/builtin/netlify"
	"github.com/hashicorp/waypoint/builtin/nomad"
	"github.com/hashicorp/waypoint/builtin/pack"
)

var (
	Builders   = mustFactory(factory.New((*component.Builder)(nil)))
	Registries = mustFactory(factory.New((*component.Registry)(nil)))
	Platforms  = mustFactory(factory.New((*component.Platform)(nil)))
	Releasers  = mustFactory(factory.New((*component.ReleaseManager)(nil)))

	// Builtins is the map of all available builtin plugins and their
	// options for launching them.
	Builtins = map[string][]sdk.Option{
		"files":                    files.Options,
		"pack":                     pack.Options,
		"docker":                   docker.Options,
		"google-cloud-run":         cloudrun.Options,
		"azure-container-instance": aci.Options,
		"kubernetes":               k8s.Options,
		"netlify":                  netlify.Options,
		"aws-ecs":                  ecs.Options,
		"aws-ecr":                  ecr.Options,
		"nomad":                    nomad.Options,
		"aws-ami":                  ami.Options,
		"aws-ec2":                  ec2.Options,
		"aws-alb":                  alb.Options,
		"go":                       golang.Options,
	}
)

func init() {
	Builders.Register("docker", BuiltinFactory("docker", component.BuilderType))
	Builders.Register("files", BuiltinFactory("files", component.BuilderType))
	Builders.Register("pack", BuiltinFactory("pack", component.BuilderType))
	Builders.Register("aws-ami", BuiltinFactory("aws-ami", component.BuilderType))
	Builders.Register("go", BuiltinFactory("go", component.BuilderType))

	Registries.Register("docker", BuiltinFactory("docker", component.RegistryType))
	Registries.Register("files", BuiltinFactory("files", component.RegistryType))
	Registries.Register("aws-ecr", BuiltinFactory("aws-ecr", component.RegistryType))

	Platforms.Register("google-cloud-run", BuiltinFactory("google-cloud-run", component.PlatformType))
	Platforms.Register("kubernetes", BuiltinFactory("kubernetes", component.PlatformType))
	Platforms.Register("azure-container-instance", BuiltinFactory("azure-container-instance", component.PlatformType))
	Platforms.Register("netlify", BuiltinFactory("netlify", component.PlatformType))
	Platforms.Register("docker", BuiltinFactory("docker", component.PlatformType))
	Platforms.Register("aws-ecs", BuiltinFactory("aws-ecs", component.PlatformType))
	Platforms.Register("nomad", BuiltinFactory("nomad", component.PlatformType))
	Platforms.Register("aws-ec2", BuiltinFactory("aws-ec2", component.PlatformType))

	Releasers.Register("google-cloud-run", BuiltinFactory("google-cloud-run", component.ReleaseManagerType))
	Releasers.Register("kubernetes", BuiltinFactory("kubernetes", component.ReleaseManagerType))
	Releasers.Register("aws-alb", BuiltinFactory("aws-alb", component.ReleaseManagerType))
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func mustFactory(f *factory.Factory, err error) *factory.Factory {
	must(err)
	return f
}
