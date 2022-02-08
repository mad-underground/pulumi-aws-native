// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package s3objectlambda

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// AWS::S3ObjectLambda::AccessPointPolicy resource is an Amazon S3ObjectLambda policy type that you can use to control permissions for your S3ObjectLambda
func LookupAccessPointPolicy(ctx *pulumi.Context, args *LookupAccessPointPolicyArgs, opts ...pulumi.InvokeOption) (*LookupAccessPointPolicyResult, error) {
	var rv LookupAccessPointPolicyResult
	err := ctx.Invoke("aws-native:s3objectlambda:getAccessPointPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAccessPointPolicyArgs struct {
	// The name of the Amazon S3 ObjectLambdaAccessPoint to which the policy applies.
	ObjectLambdaAccessPoint string `pulumi:"objectLambdaAccessPoint"`
}

type LookupAccessPointPolicyResult struct {
	// A policy document containing permissions to add to the specified ObjectLambdaAccessPoint. For more information, see Access Policy Language Overview (https://docs.aws.amazon.com/AmazonS3/latest/dev/access-policy-language-overview.html) in the Amazon Simple Storage Service Developer Guide.
	PolicyDocument interface{} `pulumi:"policyDocument"`
}

func LookupAccessPointPolicyOutput(ctx *pulumi.Context, args LookupAccessPointPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupAccessPointPolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAccessPointPolicyResult, error) {
			args := v.(LookupAccessPointPolicyArgs)
			r, err := LookupAccessPointPolicy(ctx, &args, opts...)
			return *r, err
		}).(LookupAccessPointPolicyResultOutput)
}

type LookupAccessPointPolicyOutputArgs struct {
	// The name of the Amazon S3 ObjectLambdaAccessPoint to which the policy applies.
	ObjectLambdaAccessPoint pulumi.StringInput `pulumi:"objectLambdaAccessPoint"`
}

func (LookupAccessPointPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAccessPointPolicyArgs)(nil)).Elem()
}

type LookupAccessPointPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupAccessPointPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAccessPointPolicyResult)(nil)).Elem()
}

func (o LookupAccessPointPolicyResultOutput) ToLookupAccessPointPolicyResultOutput() LookupAccessPointPolicyResultOutput {
	return o
}

func (o LookupAccessPointPolicyResultOutput) ToLookupAccessPointPolicyResultOutputWithContext(ctx context.Context) LookupAccessPointPolicyResultOutput {
	return o
}

// A policy document containing permissions to add to the specified ObjectLambdaAccessPoint. For more information, see Access Policy Language Overview (https://docs.aws.amazon.com/AmazonS3/latest/dev/access-policy-language-overview.html) in the Amazon Simple Storage Service Developer Guide.
func (o LookupAccessPointPolicyResultOutput) PolicyDocument() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupAccessPointPolicyResult) interface{} { return v.PolicyDocument }).(pulumi.AnyOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAccessPointPolicyResultOutput{})
}
