// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource schema for AWS::EC2::NetworkInsightsAnalysis
type NetworkInsightsAnalysis struct {
	pulumi.CustomResourceState

	AdditionalAccounts         pulumi.StringArrayOutput                            `pulumi:"additionalAccounts"`
	AlternatePathHints         NetworkInsightsAnalysisAlternatePathHintArrayOutput `pulumi:"alternatePathHints"`
	Explanations               NetworkInsightsAnalysisExplanationArrayOutput       `pulumi:"explanations"`
	FilterInArns               pulumi.StringArrayOutput                            `pulumi:"filterInArns"`
	ForwardPathComponents      NetworkInsightsAnalysisPathComponentArrayOutput     `pulumi:"forwardPathComponents"`
	NetworkInsightsAnalysisArn pulumi.StringOutput                                 `pulumi:"networkInsightsAnalysisArn"`
	NetworkInsightsAnalysisId  pulumi.StringOutput                                 `pulumi:"networkInsightsAnalysisId"`
	NetworkInsightsPathId      pulumi.StringOutput                                 `pulumi:"networkInsightsPathId"`
	NetworkPathFound           pulumi.BoolOutput                                   `pulumi:"networkPathFound"`
	ReturnPathComponents       NetworkInsightsAnalysisPathComponentArrayOutput     `pulumi:"returnPathComponents"`
	StartDate                  pulumi.StringOutput                                 `pulumi:"startDate"`
	Status                     NetworkInsightsAnalysisStatusOutput                 `pulumi:"status"`
	StatusMessage              pulumi.StringOutput                                 `pulumi:"statusMessage"`
	SuggestedAccounts          pulumi.StringArrayOutput                            `pulumi:"suggestedAccounts"`
	Tags                       NetworkInsightsAnalysisTagArrayOutput               `pulumi:"tags"`
}

// NewNetworkInsightsAnalysis registers a new resource with the given unique name, arguments, and options.
func NewNetworkInsightsAnalysis(ctx *pulumi.Context,
	name string, args *NetworkInsightsAnalysisArgs, opts ...pulumi.ResourceOption) (*NetworkInsightsAnalysis, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NetworkInsightsPathId == nil {
		return nil, errors.New("invalid value for required argument 'NetworkInsightsPathId'")
	}
	var resource NetworkInsightsAnalysis
	err := ctx.RegisterResource("aws-native:ec2:NetworkInsightsAnalysis", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNetworkInsightsAnalysis gets an existing NetworkInsightsAnalysis resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetworkInsightsAnalysis(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkInsightsAnalysisState, opts ...pulumi.ResourceOption) (*NetworkInsightsAnalysis, error) {
	var resource NetworkInsightsAnalysis
	err := ctx.ReadResource("aws-native:ec2:NetworkInsightsAnalysis", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NetworkInsightsAnalysis resources.
type networkInsightsAnalysisState struct {
}

type NetworkInsightsAnalysisState struct {
}

func (NetworkInsightsAnalysisState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkInsightsAnalysisState)(nil)).Elem()
}

type networkInsightsAnalysisArgs struct {
	AdditionalAccounts    []string                     `pulumi:"additionalAccounts"`
	FilterInArns          []string                     `pulumi:"filterInArns"`
	NetworkInsightsPathId string                       `pulumi:"networkInsightsPathId"`
	Tags                  []NetworkInsightsAnalysisTag `pulumi:"tags"`
}

// The set of arguments for constructing a NetworkInsightsAnalysis resource.
type NetworkInsightsAnalysisArgs struct {
	AdditionalAccounts    pulumi.StringArrayInput
	FilterInArns          pulumi.StringArrayInput
	NetworkInsightsPathId pulumi.StringInput
	Tags                  NetworkInsightsAnalysisTagArrayInput
}

func (NetworkInsightsAnalysisArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkInsightsAnalysisArgs)(nil)).Elem()
}

type NetworkInsightsAnalysisInput interface {
	pulumi.Input

	ToNetworkInsightsAnalysisOutput() NetworkInsightsAnalysisOutput
	ToNetworkInsightsAnalysisOutputWithContext(ctx context.Context) NetworkInsightsAnalysisOutput
}

func (*NetworkInsightsAnalysis) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkInsightsAnalysis)(nil)).Elem()
}

func (i *NetworkInsightsAnalysis) ToNetworkInsightsAnalysisOutput() NetworkInsightsAnalysisOutput {
	return i.ToNetworkInsightsAnalysisOutputWithContext(context.Background())
}

func (i *NetworkInsightsAnalysis) ToNetworkInsightsAnalysisOutputWithContext(ctx context.Context) NetworkInsightsAnalysisOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkInsightsAnalysisOutput)
}

type NetworkInsightsAnalysisOutput struct{ *pulumi.OutputState }

func (NetworkInsightsAnalysisOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkInsightsAnalysis)(nil)).Elem()
}

func (o NetworkInsightsAnalysisOutput) ToNetworkInsightsAnalysisOutput() NetworkInsightsAnalysisOutput {
	return o
}

func (o NetworkInsightsAnalysisOutput) ToNetworkInsightsAnalysisOutputWithContext(ctx context.Context) NetworkInsightsAnalysisOutput {
	return o
}

func (o NetworkInsightsAnalysisOutput) AdditionalAccounts() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NetworkInsightsAnalysis) pulumi.StringArrayOutput { return v.AdditionalAccounts }).(pulumi.StringArrayOutput)
}

func (o NetworkInsightsAnalysisOutput) AlternatePathHints() NetworkInsightsAnalysisAlternatePathHintArrayOutput {
	return o.ApplyT(func(v *NetworkInsightsAnalysis) NetworkInsightsAnalysisAlternatePathHintArrayOutput {
		return v.AlternatePathHints
	}).(NetworkInsightsAnalysisAlternatePathHintArrayOutput)
}

func (o NetworkInsightsAnalysisOutput) Explanations() NetworkInsightsAnalysisExplanationArrayOutput {
	return o.ApplyT(func(v *NetworkInsightsAnalysis) NetworkInsightsAnalysisExplanationArrayOutput { return v.Explanations }).(NetworkInsightsAnalysisExplanationArrayOutput)
}

func (o NetworkInsightsAnalysisOutput) FilterInArns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NetworkInsightsAnalysis) pulumi.StringArrayOutput { return v.FilterInArns }).(pulumi.StringArrayOutput)
}

func (o NetworkInsightsAnalysisOutput) ForwardPathComponents() NetworkInsightsAnalysisPathComponentArrayOutput {
	return o.ApplyT(func(v *NetworkInsightsAnalysis) NetworkInsightsAnalysisPathComponentArrayOutput {
		return v.ForwardPathComponents
	}).(NetworkInsightsAnalysisPathComponentArrayOutput)
}

func (o NetworkInsightsAnalysisOutput) NetworkInsightsAnalysisArn() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkInsightsAnalysis) pulumi.StringOutput { return v.NetworkInsightsAnalysisArn }).(pulumi.StringOutput)
}

func (o NetworkInsightsAnalysisOutput) NetworkInsightsAnalysisId() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkInsightsAnalysis) pulumi.StringOutput { return v.NetworkInsightsAnalysisId }).(pulumi.StringOutput)
}

func (o NetworkInsightsAnalysisOutput) NetworkInsightsPathId() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkInsightsAnalysis) pulumi.StringOutput { return v.NetworkInsightsPathId }).(pulumi.StringOutput)
}

func (o NetworkInsightsAnalysisOutput) NetworkPathFound() pulumi.BoolOutput {
	return o.ApplyT(func(v *NetworkInsightsAnalysis) pulumi.BoolOutput { return v.NetworkPathFound }).(pulumi.BoolOutput)
}

func (o NetworkInsightsAnalysisOutput) ReturnPathComponents() NetworkInsightsAnalysisPathComponentArrayOutput {
	return o.ApplyT(func(v *NetworkInsightsAnalysis) NetworkInsightsAnalysisPathComponentArrayOutput {
		return v.ReturnPathComponents
	}).(NetworkInsightsAnalysisPathComponentArrayOutput)
}

func (o NetworkInsightsAnalysisOutput) StartDate() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkInsightsAnalysis) pulumi.StringOutput { return v.StartDate }).(pulumi.StringOutput)
}

func (o NetworkInsightsAnalysisOutput) Status() NetworkInsightsAnalysisStatusOutput {
	return o.ApplyT(func(v *NetworkInsightsAnalysis) NetworkInsightsAnalysisStatusOutput { return v.Status }).(NetworkInsightsAnalysisStatusOutput)
}

func (o NetworkInsightsAnalysisOutput) StatusMessage() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkInsightsAnalysis) pulumi.StringOutput { return v.StatusMessage }).(pulumi.StringOutput)
}

func (o NetworkInsightsAnalysisOutput) SuggestedAccounts() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NetworkInsightsAnalysis) pulumi.StringArrayOutput { return v.SuggestedAccounts }).(pulumi.StringArrayOutput)
}

func (o NetworkInsightsAnalysisOutput) Tags() NetworkInsightsAnalysisTagArrayOutput {
	return o.ApplyT(func(v *NetworkInsightsAnalysis) NetworkInsightsAnalysisTagArrayOutput { return v.Tags }).(NetworkInsightsAnalysisTagArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkInsightsAnalysisInput)(nil)).Elem(), &NetworkInsightsAnalysis{})
	pulumi.RegisterOutputType(NetworkInsightsAnalysisOutput{})
}
