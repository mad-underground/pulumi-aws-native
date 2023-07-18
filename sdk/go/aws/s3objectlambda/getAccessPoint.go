// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package s3objectlambda

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The AWS::S3ObjectLambda::AccessPoint resource is an Amazon S3ObjectLambda resource type that you can use to add computation to S3 actions
func LookupAccessPoint(ctx *pulumi.Context, args *LookupAccessPointArgs, opts ...pulumi.InvokeOption) (*LookupAccessPointResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupAccessPointResult
	err := ctx.Invoke("aws-native:s3objectlambda:getAccessPoint", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAccessPointArgs struct {
	// The name you want to assign to this Object lambda Access Point.
	Name string `pulumi:"name"`
}

type LookupAccessPointResult struct {
	Alias *AccessPointAlias `pulumi:"alias"`
	Arn   *string           `pulumi:"arn"`
	// The date and time when the Object lambda Access Point was created.
	CreationDate *string `pulumi:"creationDate"`
	// The Object lambda Access Point Configuration that configures transformations to be applied on the objects on specified S3 Actions
	ObjectLambdaConfiguration *AccessPointObjectLambdaConfiguration `pulumi:"objectLambdaConfiguration"`
	PolicyStatus              *AccessPointPolicyStatus              `pulumi:"policyStatus"`
	// The PublicAccessBlock configuration that you want to apply to this Access Point. You can enable the configuration options in any combination. For more information about when Amazon S3 considers a bucket or object public, see https://docs.aws.amazon.com/AmazonS3/latest/dev/access-control-block-public-access.html#access-control-block-public-access-policy-status 'The Meaning of Public' in the Amazon Simple Storage Service Developer Guide.
	PublicAccessBlockConfiguration *AccessPointPublicAccessBlockConfiguration `pulumi:"publicAccessBlockConfiguration"`
}

func LookupAccessPointOutput(ctx *pulumi.Context, args LookupAccessPointOutputArgs, opts ...pulumi.InvokeOption) LookupAccessPointResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAccessPointResult, error) {
			args := v.(LookupAccessPointArgs)
			r, err := LookupAccessPoint(ctx, &args, opts...)
			var s LookupAccessPointResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAccessPointResultOutput)
}

type LookupAccessPointOutputArgs struct {
	// The name you want to assign to this Object lambda Access Point.
	Name pulumi.StringInput `pulumi:"name"`
}

func (LookupAccessPointOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAccessPointArgs)(nil)).Elem()
}

type LookupAccessPointResultOutput struct{ *pulumi.OutputState }

func (LookupAccessPointResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAccessPointResult)(nil)).Elem()
}

func (o LookupAccessPointResultOutput) ToLookupAccessPointResultOutput() LookupAccessPointResultOutput {
	return o
}

func (o LookupAccessPointResultOutput) ToLookupAccessPointResultOutputWithContext(ctx context.Context) LookupAccessPointResultOutput {
	return o
}

func (o LookupAccessPointResultOutput) Alias() AccessPointAliasPtrOutput {
	return o.ApplyT(func(v LookupAccessPointResult) *AccessPointAlias { return v.Alias }).(AccessPointAliasPtrOutput)
}

func (o LookupAccessPointResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAccessPointResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

// The date and time when the Object lambda Access Point was created.
func (o LookupAccessPointResultOutput) CreationDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAccessPointResult) *string { return v.CreationDate }).(pulumi.StringPtrOutput)
}

// The Object lambda Access Point Configuration that configures transformations to be applied on the objects on specified S3 Actions
func (o LookupAccessPointResultOutput) ObjectLambdaConfiguration() AccessPointObjectLambdaConfigurationPtrOutput {
	return o.ApplyT(func(v LookupAccessPointResult) *AccessPointObjectLambdaConfiguration {
		return v.ObjectLambdaConfiguration
	}).(AccessPointObjectLambdaConfigurationPtrOutput)
}

func (o LookupAccessPointResultOutput) PolicyStatus() AccessPointPolicyStatusPtrOutput {
	return o.ApplyT(func(v LookupAccessPointResult) *AccessPointPolicyStatus { return v.PolicyStatus }).(AccessPointPolicyStatusPtrOutput)
}

// The PublicAccessBlock configuration that you want to apply to this Access Point. You can enable the configuration options in any combination. For more information about when Amazon S3 considers a bucket or object public, see https://docs.aws.amazon.com/AmazonS3/latest/dev/access-control-block-public-access.html#access-control-block-public-access-policy-status 'The Meaning of Public' in the Amazon Simple Storage Service Developer Guide.
func (o LookupAccessPointResultOutput) PublicAccessBlockConfiguration() AccessPointPublicAccessBlockConfigurationPtrOutput {
	return o.ApplyT(func(v LookupAccessPointResult) *AccessPointPublicAccessBlockConfiguration {
		return v.PublicAccessBlockConfiguration
	}).(AccessPointPublicAccessBlockConfigurationPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAccessPointResultOutput{})
}
