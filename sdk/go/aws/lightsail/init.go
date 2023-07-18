// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package lightsail

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
	case "aws-native:lightsail:Alarm":
		r = &Alarm{}
	case "aws-native:lightsail:Bucket":
		r = &Bucket{}
	case "aws-native:lightsail:Certificate":
		r = &Certificate{}
	case "aws-native:lightsail:Container":
		r = &Container{}
	case "aws-native:lightsail:Database":
		r = &Database{}
	case "aws-native:lightsail:Disk":
		r = &Disk{}
	case "aws-native:lightsail:Distribution":
		r = &Distribution{}
	case "aws-native:lightsail:Instance":
		r = &Instance{}
	case "aws-native:lightsail:LoadBalancer":
		r = &LoadBalancer{}
	case "aws-native:lightsail:LoadBalancerTlsCertificate":
		r = &LoadBalancerTlsCertificate{}
	case "aws-native:lightsail:StaticIp":
		r = &StaticIp{}
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
		"lightsail",
		&module{version},
	)
}
