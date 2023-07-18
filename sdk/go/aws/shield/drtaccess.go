// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package shield

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Config the role and list of Amazon S3 log buckets used by the Shield Response Team (SRT) to access your AWS account while assisting with attack mitigation.
type DRTAccess struct {
	pulumi.CustomResourceState

	AccountId pulumi.StringOutput `pulumi:"accountId"`
	// Authorizes the Shield Response Team (SRT) to access the specified Amazon S3 bucket containing log data such as Application Load Balancer access logs, CloudFront logs, or logs from third party sources. You can associate up to 10 Amazon S3 buckets with your subscription.
	LogBucketList pulumi.StringArrayOutput `pulumi:"logBucketList"`
	// Authorizes the Shield Response Team (SRT) using the specified role, to access your AWS account to assist with DDoS attack mitigation during potential attacks. This enables the SRT to inspect your AWS WAF configuration and create or update AWS WAF rules and web ACLs.
	RoleArn pulumi.StringOutput `pulumi:"roleArn"`
}

// NewDRTAccess registers a new resource with the given unique name, arguments, and options.
func NewDRTAccess(ctx *pulumi.Context,
	name string, args *DRTAccessArgs, opts ...pulumi.ResourceOption) (*DRTAccess, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RoleArn == nil {
		return nil, errors.New("invalid value for required argument 'RoleArn'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DRTAccess
	err := ctx.RegisterResource("aws-native:shield:DRTAccess", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDRTAccess gets an existing DRTAccess resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDRTAccess(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DRTAccessState, opts ...pulumi.ResourceOption) (*DRTAccess, error) {
	var resource DRTAccess
	err := ctx.ReadResource("aws-native:shield:DRTAccess", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DRTAccess resources.
type drtaccessState struct {
}

type DRTAccessState struct {
}

func (DRTAccessState) ElementType() reflect.Type {
	return reflect.TypeOf((*drtaccessState)(nil)).Elem()
}

type drtaccessArgs struct {
	// Authorizes the Shield Response Team (SRT) to access the specified Amazon S3 bucket containing log data such as Application Load Balancer access logs, CloudFront logs, or logs from third party sources. You can associate up to 10 Amazon S3 buckets with your subscription.
	LogBucketList []string `pulumi:"logBucketList"`
	// Authorizes the Shield Response Team (SRT) using the specified role, to access your AWS account to assist with DDoS attack mitigation during potential attacks. This enables the SRT to inspect your AWS WAF configuration and create or update AWS WAF rules and web ACLs.
	RoleArn string `pulumi:"roleArn"`
}

// The set of arguments for constructing a DRTAccess resource.
type DRTAccessArgs struct {
	// Authorizes the Shield Response Team (SRT) to access the specified Amazon S3 bucket containing log data such as Application Load Balancer access logs, CloudFront logs, or logs from third party sources. You can associate up to 10 Amazon S3 buckets with your subscription.
	LogBucketList pulumi.StringArrayInput
	// Authorizes the Shield Response Team (SRT) using the specified role, to access your AWS account to assist with DDoS attack mitigation during potential attacks. This enables the SRT to inspect your AWS WAF configuration and create or update AWS WAF rules and web ACLs.
	RoleArn pulumi.StringInput
}

func (DRTAccessArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*drtaccessArgs)(nil)).Elem()
}

type DRTAccessInput interface {
	pulumi.Input

	ToDRTAccessOutput() DRTAccessOutput
	ToDRTAccessOutputWithContext(ctx context.Context) DRTAccessOutput
}

func (*DRTAccess) ElementType() reflect.Type {
	return reflect.TypeOf((**DRTAccess)(nil)).Elem()
}

func (i *DRTAccess) ToDRTAccessOutput() DRTAccessOutput {
	return i.ToDRTAccessOutputWithContext(context.Background())
}

func (i *DRTAccess) ToDRTAccessOutputWithContext(ctx context.Context) DRTAccessOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DRTAccessOutput)
}

type DRTAccessOutput struct{ *pulumi.OutputState }

func (DRTAccessOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DRTAccess)(nil)).Elem()
}

func (o DRTAccessOutput) ToDRTAccessOutput() DRTAccessOutput {
	return o
}

func (o DRTAccessOutput) ToDRTAccessOutputWithContext(ctx context.Context) DRTAccessOutput {
	return o
}

func (o DRTAccessOutput) AccountId() pulumi.StringOutput {
	return o.ApplyT(func(v *DRTAccess) pulumi.StringOutput { return v.AccountId }).(pulumi.StringOutput)
}

// Authorizes the Shield Response Team (SRT) to access the specified Amazon S3 bucket containing log data such as Application Load Balancer access logs, CloudFront logs, or logs from third party sources. You can associate up to 10 Amazon S3 buckets with your subscription.
func (o DRTAccessOutput) LogBucketList() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DRTAccess) pulumi.StringArrayOutput { return v.LogBucketList }).(pulumi.StringArrayOutput)
}

// Authorizes the Shield Response Team (SRT) using the specified role, to access your AWS account to assist with DDoS attack mitigation during potential attacks. This enables the SRT to inspect your AWS WAF configuration and create or update AWS WAF rules and web ACLs.
func (o DRTAccessOutput) RoleArn() pulumi.StringOutput {
	return o.ApplyT(func(v *DRTAccess) pulumi.StringOutput { return v.RoleArn }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DRTAccessInput)(nil)).Elem(), &DRTAccess{})
	pulumi.RegisterOutputType(DRTAccessOutput{})
}
