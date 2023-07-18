// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource schema for AWS::EC2::NetworkInsightsAccessScopeAnalysis
type NetworkInsightsAccessScopeAnalysis struct {
	pulumi.CustomResourceState

	AnalyzedEniCount                      pulumi.IntOutput                                      `pulumi:"analyzedEniCount"`
	EndDate                               pulumi.StringOutput                                   `pulumi:"endDate"`
	FindingsFound                         NetworkInsightsAccessScopeAnalysisFindingsFoundOutput `pulumi:"findingsFound"`
	NetworkInsightsAccessScopeAnalysisArn pulumi.StringOutput                                   `pulumi:"networkInsightsAccessScopeAnalysisArn"`
	NetworkInsightsAccessScopeAnalysisId  pulumi.StringOutput                                   `pulumi:"networkInsightsAccessScopeAnalysisId"`
	NetworkInsightsAccessScopeId          pulumi.StringOutput                                   `pulumi:"networkInsightsAccessScopeId"`
	StartDate                             pulumi.StringOutput                                   `pulumi:"startDate"`
	Status                                NetworkInsightsAccessScopeAnalysisStatusOutput        `pulumi:"status"`
	StatusMessage                         pulumi.StringOutput                                   `pulumi:"statusMessage"`
	Tags                                  NetworkInsightsAccessScopeAnalysisTagArrayOutput      `pulumi:"tags"`
}

// NewNetworkInsightsAccessScopeAnalysis registers a new resource with the given unique name, arguments, and options.
func NewNetworkInsightsAccessScopeAnalysis(ctx *pulumi.Context,
	name string, args *NetworkInsightsAccessScopeAnalysisArgs, opts ...pulumi.ResourceOption) (*NetworkInsightsAccessScopeAnalysis, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NetworkInsightsAccessScopeId == nil {
		return nil, errors.New("invalid value for required argument 'NetworkInsightsAccessScopeId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource NetworkInsightsAccessScopeAnalysis
	err := ctx.RegisterResource("aws-native:ec2:NetworkInsightsAccessScopeAnalysis", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNetworkInsightsAccessScopeAnalysis gets an existing NetworkInsightsAccessScopeAnalysis resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetworkInsightsAccessScopeAnalysis(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkInsightsAccessScopeAnalysisState, opts ...pulumi.ResourceOption) (*NetworkInsightsAccessScopeAnalysis, error) {
	var resource NetworkInsightsAccessScopeAnalysis
	err := ctx.ReadResource("aws-native:ec2:NetworkInsightsAccessScopeAnalysis", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NetworkInsightsAccessScopeAnalysis resources.
type networkInsightsAccessScopeAnalysisState struct {
}

type NetworkInsightsAccessScopeAnalysisState struct {
}

func (NetworkInsightsAccessScopeAnalysisState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkInsightsAccessScopeAnalysisState)(nil)).Elem()
}

type networkInsightsAccessScopeAnalysisArgs struct {
	NetworkInsightsAccessScopeId string                                  `pulumi:"networkInsightsAccessScopeId"`
	Tags                         []NetworkInsightsAccessScopeAnalysisTag `pulumi:"tags"`
}

// The set of arguments for constructing a NetworkInsightsAccessScopeAnalysis resource.
type NetworkInsightsAccessScopeAnalysisArgs struct {
	NetworkInsightsAccessScopeId pulumi.StringInput
	Tags                         NetworkInsightsAccessScopeAnalysisTagArrayInput
}

func (NetworkInsightsAccessScopeAnalysisArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkInsightsAccessScopeAnalysisArgs)(nil)).Elem()
}

type NetworkInsightsAccessScopeAnalysisInput interface {
	pulumi.Input

	ToNetworkInsightsAccessScopeAnalysisOutput() NetworkInsightsAccessScopeAnalysisOutput
	ToNetworkInsightsAccessScopeAnalysisOutputWithContext(ctx context.Context) NetworkInsightsAccessScopeAnalysisOutput
}

func (*NetworkInsightsAccessScopeAnalysis) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkInsightsAccessScopeAnalysis)(nil)).Elem()
}

func (i *NetworkInsightsAccessScopeAnalysis) ToNetworkInsightsAccessScopeAnalysisOutput() NetworkInsightsAccessScopeAnalysisOutput {
	return i.ToNetworkInsightsAccessScopeAnalysisOutputWithContext(context.Background())
}

func (i *NetworkInsightsAccessScopeAnalysis) ToNetworkInsightsAccessScopeAnalysisOutputWithContext(ctx context.Context) NetworkInsightsAccessScopeAnalysisOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkInsightsAccessScopeAnalysisOutput)
}

type NetworkInsightsAccessScopeAnalysisOutput struct{ *pulumi.OutputState }

func (NetworkInsightsAccessScopeAnalysisOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkInsightsAccessScopeAnalysis)(nil)).Elem()
}

func (o NetworkInsightsAccessScopeAnalysisOutput) ToNetworkInsightsAccessScopeAnalysisOutput() NetworkInsightsAccessScopeAnalysisOutput {
	return o
}

func (o NetworkInsightsAccessScopeAnalysisOutput) ToNetworkInsightsAccessScopeAnalysisOutputWithContext(ctx context.Context) NetworkInsightsAccessScopeAnalysisOutput {
	return o
}

func (o NetworkInsightsAccessScopeAnalysisOutput) AnalyzedEniCount() pulumi.IntOutput {
	return o.ApplyT(func(v *NetworkInsightsAccessScopeAnalysis) pulumi.IntOutput { return v.AnalyzedEniCount }).(pulumi.IntOutput)
}

func (o NetworkInsightsAccessScopeAnalysisOutput) EndDate() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkInsightsAccessScopeAnalysis) pulumi.StringOutput { return v.EndDate }).(pulumi.StringOutput)
}

func (o NetworkInsightsAccessScopeAnalysisOutput) FindingsFound() NetworkInsightsAccessScopeAnalysisFindingsFoundOutput {
	return o.ApplyT(func(v *NetworkInsightsAccessScopeAnalysis) NetworkInsightsAccessScopeAnalysisFindingsFoundOutput {
		return v.FindingsFound
	}).(NetworkInsightsAccessScopeAnalysisFindingsFoundOutput)
}

func (o NetworkInsightsAccessScopeAnalysisOutput) NetworkInsightsAccessScopeAnalysisArn() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkInsightsAccessScopeAnalysis) pulumi.StringOutput {
		return v.NetworkInsightsAccessScopeAnalysisArn
	}).(pulumi.StringOutput)
}

func (o NetworkInsightsAccessScopeAnalysisOutput) NetworkInsightsAccessScopeAnalysisId() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkInsightsAccessScopeAnalysis) pulumi.StringOutput {
		return v.NetworkInsightsAccessScopeAnalysisId
	}).(pulumi.StringOutput)
}

func (o NetworkInsightsAccessScopeAnalysisOutput) NetworkInsightsAccessScopeId() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkInsightsAccessScopeAnalysis) pulumi.StringOutput { return v.NetworkInsightsAccessScopeId }).(pulumi.StringOutput)
}

func (o NetworkInsightsAccessScopeAnalysisOutput) StartDate() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkInsightsAccessScopeAnalysis) pulumi.StringOutput { return v.StartDate }).(pulumi.StringOutput)
}

func (o NetworkInsightsAccessScopeAnalysisOutput) Status() NetworkInsightsAccessScopeAnalysisStatusOutput {
	return o.ApplyT(func(v *NetworkInsightsAccessScopeAnalysis) NetworkInsightsAccessScopeAnalysisStatusOutput {
		return v.Status
	}).(NetworkInsightsAccessScopeAnalysisStatusOutput)
}

func (o NetworkInsightsAccessScopeAnalysisOutput) StatusMessage() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkInsightsAccessScopeAnalysis) pulumi.StringOutput { return v.StatusMessage }).(pulumi.StringOutput)
}

func (o NetworkInsightsAccessScopeAnalysisOutput) Tags() NetworkInsightsAccessScopeAnalysisTagArrayOutput {
	return o.ApplyT(func(v *NetworkInsightsAccessScopeAnalysis) NetworkInsightsAccessScopeAnalysisTagArrayOutput {
		return v.Tags
	}).(NetworkInsightsAccessScopeAnalysisTagArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkInsightsAccessScopeAnalysisInput)(nil)).Elem(), &NetworkInsightsAccessScopeAnalysis{})
	pulumi.RegisterOutputType(NetworkInsightsAccessScopeAnalysisOutput{})
}
