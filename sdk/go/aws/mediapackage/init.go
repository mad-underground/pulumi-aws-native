// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mediapackage

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
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
	case "aws-native:mediapackage:Asset":
		r = &Asset{}
	case "aws-native:mediapackage:Channel":
		r = &Channel{}
	case "aws-native:mediapackage:OriginEndpoint":
		r = &OriginEndpoint{}
	case "aws-native:mediapackage:PackagingConfiguration":
		r = &PackagingConfiguration{}
	case "aws-native:mediapackage:PackagingGroup":
		r = &PackagingGroup{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := aws.PkgVersion()
	if err != nil {
		version = semver.Version{Major: 1}
	}
	pulumi.RegisterResourceModule(
		"aws-native",
		"mediapackage",
		&module{version},
	)
}
