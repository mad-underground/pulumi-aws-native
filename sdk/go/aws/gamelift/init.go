// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package gamelift

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
	case "aws-native:gamelift:Alias":
		r = &Alias{}
	case "aws-native:gamelift:Build":
		r = &Build{}
	case "aws-native:gamelift:Fleet":
		r = &Fleet{}
	case "aws-native:gamelift:GameServerGroup":
		r = &GameServerGroup{}
	case "aws-native:gamelift:GameSessionQueue":
		r = &GameSessionQueue{}
	case "aws-native:gamelift:Location":
		r = &Location{}
	case "aws-native:gamelift:MatchmakingConfiguration":
		r = &MatchmakingConfiguration{}
	case "aws-native:gamelift:MatchmakingRuleSet":
		r = &MatchmakingRuleSet{}
	case "aws-native:gamelift:Script":
		r = &Script{}
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
		"gamelift",
		&module{version},
	)
}
