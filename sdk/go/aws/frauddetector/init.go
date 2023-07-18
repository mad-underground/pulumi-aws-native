// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package frauddetector

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
	case "aws-native:frauddetector:Detector":
		r = &Detector{}
	case "aws-native:frauddetector:EntityType":
		r = &EntityType{}
	case "aws-native:frauddetector:EventType":
		r = &EventType{}
	case "aws-native:frauddetector:Label":
		r = &Label{}
	case "aws-native:frauddetector:List":
		r = &List{}
	case "aws-native:frauddetector:Outcome":
		r = &Outcome{}
	case "aws-native:frauddetector:Variable":
		r = &Variable{}
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
		"frauddetector",
		&module{version},
	)
}
