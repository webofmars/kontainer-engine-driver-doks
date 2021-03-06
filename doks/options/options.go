package options

import "github.com/rancher/kontainer-engine/types"

type Builder interface {
	BuildCreateOptions() *types.DriverFlags
	BuildUpdateOptions() *types.DriverFlags
}

type builderImpl struct {
}

func NewBuilder() Builder {
	return builderImpl{}
}

func (builderImpl) BuildCreateOptions() *types.DriverFlags {

	builder := flagBuilder()

	builder(
			"token",
			types.StringType,
			"Personal access token",
			nil,
		)

	builder(
		"display-name",
		types.StringType,
		"The displayed name of the cluster in the Rancher UI",
		nil,
	)

	builder(
		"name",
		types.StringType,
		"Default name to something meaningful to you",
		nil,
	)

	builder(
		"auto-upgraded",
		types.BoolPointerType,
		"Automatically updates Kubernetes version",
		&types.Default{
			DefaultBool: false,
		},
	)

	builder(
		"region-slug",
		types.StringType,
		"Your Kubernetes cluster will be located in a single datacenter.",
		&types.Default{
			DefaultString: "nyc3",
		},
	)

	builder(
		"version-slug",
		types.StringType,
		"Kubernetes version",
		nil,
	)

	builder(
		"node-pool-name",
		types.StringType,
		"Name of pool of worker nodes",
		nil,
	)

	builder(
		"node-pool-autoscale",
		types.BoolPointerType,
		"Enables auto scaling group",
		&types.Default{
			DefaultBool: false,
		},
	)

	builder(
		"node-pool-count",
		types.IntType,
		"The desired number of worker nodes",
		&types.Default{
			DefaultInt: 2,
		},
	)

	builder(
		"node-pool-min",
		types.IntType,
		"The minimum number of worker nodes",
		nil,
	)

	builder(
		"node-pool-max",
		types.IntType,
		"The maximum number of worker nodes",
		nil,
	)

	builder(
		"node-pool-size",
		types.StringType,
		"The type of machine to use for worker nodes",
		&types.Default{
			DefaultString: "s-2vcpu-2gb",
		},
	)

	builder(
		"tags",
		types.StringSliceType,
		"Optional tags to your cluster",
		nil,
	)

	builder(
		"node-pool-labels",
		types.StringSliceType,
		"Labels for the node pool",
		nil,
	)

	return builder(
		"vpc-id",
		types.StringType,
		"Cluster VPC",
		nil,
	)
}

func (builderImpl) BuildUpdateOptions() *types.DriverFlags{

	builder := flagBuilder()

	builder(
		"auto-upgraded",
		types.BoolPointerType,
		"Automatically updates Kubernetes version",
		nil,
	)

	builder(
		"token",
		types.StringType,
		"Personal access token",
		nil,
	)

	builder(
		"tags",
		types.StringSliceType,
		"Optional tags to your cluster",
		nil,
	)

	builder(
		"node-pool-labels",
		types.StringSliceType,
		"Labels for the node pool",
		nil,
	)

	builder(
		"node-pool-min",
		types.IntType,
		"The minimum number of worker nodes",
		nil,
	)

	builder(
		"node-pool-max",
		types.IntType,
		"The maximum number of worker nodes",
		nil,
	)

	builder(
		"node-pool-name",
		types.StringType,
		"Name of pool of worker nodes",
		nil,
	)

	builder(
		"node-pool-autoscale",
		types.BoolPointerType,
		"Enables auto scaling group",
		nil,
	)

	return builder(
		"node-pool-count",
		types.IntType,
		"The desired number of worker nodes",
		nil,
	)
}

func flagBuilder() func(name, typ, usage string, def *types.Default) *types.DriverFlags {
	flags := &types.DriverFlags{
		Options: make(map[string]*types.Flag),
	}

	return func(name, typ, usage string, def *types.Default) *types.DriverFlags {
		flags.Options[name] = &types.Flag{
			Type:    typ,
			Usage:   usage,
			Default: def,
		}

		return flags
	}
}
