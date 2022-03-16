// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package greengrass

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
	case "aws-native:greengrass:ConnectorDefinition":
		r = &ConnectorDefinition{}
	case "aws-native:greengrass:ConnectorDefinitionVersion":
		r = &ConnectorDefinitionVersion{}
	case "aws-native:greengrass:CoreDefinition":
		r = &CoreDefinition{}
	case "aws-native:greengrass:CoreDefinitionVersion":
		r = &CoreDefinitionVersion{}
	case "aws-native:greengrass:DeviceDefinition":
		r = &DeviceDefinition{}
	case "aws-native:greengrass:DeviceDefinitionVersion":
		r = &DeviceDefinitionVersion{}
	case "aws-native:greengrass:FunctionDefinition":
		r = &FunctionDefinition{}
	case "aws-native:greengrass:FunctionDefinitionVersion":
		r = &FunctionDefinitionVersion{}
	case "aws-native:greengrass:Group":
		r = &Group{}
	case "aws-native:greengrass:GroupVersion":
		r = &GroupVersion{}
	case "aws-native:greengrass:LoggerDefinition":
		r = &LoggerDefinition{}
	case "aws-native:greengrass:LoggerDefinitionVersion":
		r = &LoggerDefinitionVersion{}
	case "aws-native:greengrass:ResourceDefinition":
		r = &ResourceDefinition{}
	case "aws-native:greengrass:ResourceDefinitionVersion":
		r = &ResourceDefinitionVersion{}
	case "aws-native:greengrass:SubscriptionDefinition":
		r = &SubscriptionDefinition{}
	case "aws-native:greengrass:SubscriptionDefinitionVersion":
		r = &SubscriptionDefinitionVersion{}
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
		"greengrass",
		&module{version},
	)
}
