// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package logs

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "aws-native:logs:AccountPolicy":
		r = &AccountPolicy{}
	case "aws-native:logs:Destination":
		r = &Destination{}
	case "aws-native:logs:LogGroup":
		r = &LogGroup{}
	case "aws-native:logs:LogStream":
		r = &LogStream{}
	case "aws-native:logs:MetricFilter":
		r = &MetricFilter{}
	case "aws-native:logs:QueryDefinition":
		r = &QueryDefinition{}
	case "aws-native:logs:ResourcePolicy":
		r = &ResourcePolicy{}
	case "aws-native:logs:SubscriptionFilter":
		r = &SubscriptionFilter{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := internal.PkgVersion()
	if err != nil {
		version = semver.Version{Major: 1}
	}
	pulumi.RegisterResourceModule(
		"aws-native",
		"logs",
		&module{version},
	)
}
