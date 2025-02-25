// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vpclattice

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Enables access logs to be sent to Amazon CloudWatch, Amazon S3, and Amazon Kinesis Data Firehose. The service network owner can use the access logs to audit the services in the network. The service network owner will only see access logs from clients and services that are associated with their service network. Access log entries represent traffic originated from VPCs associated with that network.
type AccessLogSubscription struct {
	pulumi.CustomResourceState

	Arn                pulumi.StringOutput    `pulumi:"arn"`
	AwsId              pulumi.StringOutput    `pulumi:"awsId"`
	DestinationArn     pulumi.StringOutput    `pulumi:"destinationArn"`
	ResourceArn        pulumi.StringOutput    `pulumi:"resourceArn"`
	ResourceId         pulumi.StringOutput    `pulumi:"resourceId"`
	ResourceIdentifier pulumi.StringPtrOutput `pulumi:"resourceIdentifier"`
	Tags               aws.TagArrayOutput     `pulumi:"tags"`
}

// NewAccessLogSubscription registers a new resource with the given unique name, arguments, and options.
func NewAccessLogSubscription(ctx *pulumi.Context,
	name string, args *AccessLogSubscriptionArgs, opts ...pulumi.ResourceOption) (*AccessLogSubscription, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DestinationArn == nil {
		return nil, errors.New("invalid value for required argument 'DestinationArn'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"resourceIdentifier",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource AccessLogSubscription
	err := ctx.RegisterResource("aws-native:vpclattice:AccessLogSubscription", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAccessLogSubscription gets an existing AccessLogSubscription resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAccessLogSubscription(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AccessLogSubscriptionState, opts ...pulumi.ResourceOption) (*AccessLogSubscription, error) {
	var resource AccessLogSubscription
	err := ctx.ReadResource("aws-native:vpclattice:AccessLogSubscription", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AccessLogSubscription resources.
type accessLogSubscriptionState struct {
}

type AccessLogSubscriptionState struct {
}

func (AccessLogSubscriptionState) ElementType() reflect.Type {
	return reflect.TypeOf((*accessLogSubscriptionState)(nil)).Elem()
}

type accessLogSubscriptionArgs struct {
	DestinationArn     string    `pulumi:"destinationArn"`
	ResourceIdentifier *string   `pulumi:"resourceIdentifier"`
	Tags               []aws.Tag `pulumi:"tags"`
}

// The set of arguments for constructing a AccessLogSubscription resource.
type AccessLogSubscriptionArgs struct {
	DestinationArn     pulumi.StringInput
	ResourceIdentifier pulumi.StringPtrInput
	Tags               aws.TagArrayInput
}

func (AccessLogSubscriptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*accessLogSubscriptionArgs)(nil)).Elem()
}

type AccessLogSubscriptionInput interface {
	pulumi.Input

	ToAccessLogSubscriptionOutput() AccessLogSubscriptionOutput
	ToAccessLogSubscriptionOutputWithContext(ctx context.Context) AccessLogSubscriptionOutput
}

func (*AccessLogSubscription) ElementType() reflect.Type {
	return reflect.TypeOf((**AccessLogSubscription)(nil)).Elem()
}

func (i *AccessLogSubscription) ToAccessLogSubscriptionOutput() AccessLogSubscriptionOutput {
	return i.ToAccessLogSubscriptionOutputWithContext(context.Background())
}

func (i *AccessLogSubscription) ToAccessLogSubscriptionOutputWithContext(ctx context.Context) AccessLogSubscriptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessLogSubscriptionOutput)
}

type AccessLogSubscriptionOutput struct{ *pulumi.OutputState }

func (AccessLogSubscriptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AccessLogSubscription)(nil)).Elem()
}

func (o AccessLogSubscriptionOutput) ToAccessLogSubscriptionOutput() AccessLogSubscriptionOutput {
	return o
}

func (o AccessLogSubscriptionOutput) ToAccessLogSubscriptionOutputWithContext(ctx context.Context) AccessLogSubscriptionOutput {
	return o
}

func (o AccessLogSubscriptionOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessLogSubscription) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

func (o AccessLogSubscriptionOutput) AwsId() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessLogSubscription) pulumi.StringOutput { return v.AwsId }).(pulumi.StringOutput)
}

func (o AccessLogSubscriptionOutput) DestinationArn() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessLogSubscription) pulumi.StringOutput { return v.DestinationArn }).(pulumi.StringOutput)
}

func (o AccessLogSubscriptionOutput) ResourceArn() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessLogSubscription) pulumi.StringOutput { return v.ResourceArn }).(pulumi.StringOutput)
}

func (o AccessLogSubscriptionOutput) ResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessLogSubscription) pulumi.StringOutput { return v.ResourceId }).(pulumi.StringOutput)
}

func (o AccessLogSubscriptionOutput) ResourceIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AccessLogSubscription) pulumi.StringPtrOutput { return v.ResourceIdentifier }).(pulumi.StringPtrOutput)
}

func (o AccessLogSubscriptionOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v *AccessLogSubscription) aws.TagArrayOutput { return v.Tags }).(aws.TagArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AccessLogSubscriptionInput)(nil)).Elem(), &AccessLogSubscription{})
	pulumi.RegisterOutputType(AccessLogSubscriptionOutput{})
}
