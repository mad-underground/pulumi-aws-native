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
	FilterInArns          []string                     `pulumi:"filterInArns"`
	NetworkInsightsPathId string                       `pulumi:"networkInsightsPathId"`
	Tags                  []NetworkInsightsAnalysisTag `pulumi:"tags"`
}

// The set of arguments for constructing a NetworkInsightsAnalysis resource.
type NetworkInsightsAnalysisArgs struct {
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

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkInsightsAnalysisInput)(nil)).Elem(), &NetworkInsightsAnalysis{})
	pulumi.RegisterOutputType(NetworkInsightsAnalysisOutput{})
}
