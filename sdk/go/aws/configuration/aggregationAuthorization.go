// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package configuration

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Config::AggregationAuthorization
type AggregationAuthorization struct {
	pulumi.CustomResourceState

	// The ARN of the AggregationAuthorization.
	AggregationAuthorizationArn pulumi.StringOutput `pulumi:"aggregationAuthorizationArn"`
	// The 12-digit account ID of the account authorized to aggregate data.
	AuthorizedAccountId pulumi.StringOutput `pulumi:"authorizedAccountId"`
	// The region authorized to collect aggregated data.
	AuthorizedAwsRegion pulumi.StringOutput `pulumi:"authorizedAwsRegion"`
	// The tags for the AggregationAuthorization.
	Tags AggregationAuthorizationTagArrayOutput `pulumi:"tags"`
}

// NewAggregationAuthorization registers a new resource with the given unique name, arguments, and options.
func NewAggregationAuthorization(ctx *pulumi.Context,
	name string, args *AggregationAuthorizationArgs, opts ...pulumi.ResourceOption) (*AggregationAuthorization, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AuthorizedAccountId == nil {
		return nil, errors.New("invalid value for required argument 'AuthorizedAccountId'")
	}
	if args.AuthorizedAwsRegion == nil {
		return nil, errors.New("invalid value for required argument 'AuthorizedAwsRegion'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource AggregationAuthorization
	err := ctx.RegisterResource("aws-native:configuration:AggregationAuthorization", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAggregationAuthorization gets an existing AggregationAuthorization resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAggregationAuthorization(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AggregationAuthorizationState, opts ...pulumi.ResourceOption) (*AggregationAuthorization, error) {
	var resource AggregationAuthorization
	err := ctx.ReadResource("aws-native:configuration:AggregationAuthorization", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AggregationAuthorization resources.
type aggregationAuthorizationState struct {
}

type AggregationAuthorizationState struct {
}

func (AggregationAuthorizationState) ElementType() reflect.Type {
	return reflect.TypeOf((*aggregationAuthorizationState)(nil)).Elem()
}

type aggregationAuthorizationArgs struct {
	// The 12-digit account ID of the account authorized to aggregate data.
	AuthorizedAccountId string `pulumi:"authorizedAccountId"`
	// The region authorized to collect aggregated data.
	AuthorizedAwsRegion string `pulumi:"authorizedAwsRegion"`
	// The tags for the AggregationAuthorization.
	Tags []AggregationAuthorizationTag `pulumi:"tags"`
}

// The set of arguments for constructing a AggregationAuthorization resource.
type AggregationAuthorizationArgs struct {
	// The 12-digit account ID of the account authorized to aggregate data.
	AuthorizedAccountId pulumi.StringInput
	// The region authorized to collect aggregated data.
	AuthorizedAwsRegion pulumi.StringInput
	// The tags for the AggregationAuthorization.
	Tags AggregationAuthorizationTagArrayInput
}

func (AggregationAuthorizationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*aggregationAuthorizationArgs)(nil)).Elem()
}

type AggregationAuthorizationInput interface {
	pulumi.Input

	ToAggregationAuthorizationOutput() AggregationAuthorizationOutput
	ToAggregationAuthorizationOutputWithContext(ctx context.Context) AggregationAuthorizationOutput
}

func (*AggregationAuthorization) ElementType() reflect.Type {
	return reflect.TypeOf((**AggregationAuthorization)(nil)).Elem()
}

func (i *AggregationAuthorization) ToAggregationAuthorizationOutput() AggregationAuthorizationOutput {
	return i.ToAggregationAuthorizationOutputWithContext(context.Background())
}

func (i *AggregationAuthorization) ToAggregationAuthorizationOutputWithContext(ctx context.Context) AggregationAuthorizationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AggregationAuthorizationOutput)
}

type AggregationAuthorizationOutput struct{ *pulumi.OutputState }

func (AggregationAuthorizationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AggregationAuthorization)(nil)).Elem()
}

func (o AggregationAuthorizationOutput) ToAggregationAuthorizationOutput() AggregationAuthorizationOutput {
	return o
}

func (o AggregationAuthorizationOutput) ToAggregationAuthorizationOutputWithContext(ctx context.Context) AggregationAuthorizationOutput {
	return o
}

// The ARN of the AggregationAuthorization.
func (o AggregationAuthorizationOutput) AggregationAuthorizationArn() pulumi.StringOutput {
	return o.ApplyT(func(v *AggregationAuthorization) pulumi.StringOutput { return v.AggregationAuthorizationArn }).(pulumi.StringOutput)
}

// The 12-digit account ID of the account authorized to aggregate data.
func (o AggregationAuthorizationOutput) AuthorizedAccountId() pulumi.StringOutput {
	return o.ApplyT(func(v *AggregationAuthorization) pulumi.StringOutput { return v.AuthorizedAccountId }).(pulumi.StringOutput)
}

// The region authorized to collect aggregated data.
func (o AggregationAuthorizationOutput) AuthorizedAwsRegion() pulumi.StringOutput {
	return o.ApplyT(func(v *AggregationAuthorization) pulumi.StringOutput { return v.AuthorizedAwsRegion }).(pulumi.StringOutput)
}

// The tags for the AggregationAuthorization.
func (o AggregationAuthorizationOutput) Tags() AggregationAuthorizationTagArrayOutput {
	return o.ApplyT(func(v *AggregationAuthorization) AggregationAuthorizationTagArrayOutput { return v.Tags }).(AggregationAuthorizationTagArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AggregationAuthorizationInput)(nil)).Elem(), &AggregationAuthorization{})
	pulumi.RegisterOutputType(AggregationAuthorizationOutput{})
}
