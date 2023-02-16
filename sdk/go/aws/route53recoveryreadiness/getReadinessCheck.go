// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package route53recoveryreadiness

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Aws Route53 Recovery Readiness Check Schema and API specification.
func LookupReadinessCheck(ctx *pulumi.Context, args *LookupReadinessCheckArgs, opts ...pulumi.InvokeOption) (*LookupReadinessCheckResult, error) {
	var rv LookupReadinessCheckResult
	err := ctx.Invoke("aws-native:route53recoveryreadiness:getReadinessCheck", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupReadinessCheckArgs struct {
	// Name of the ReadinessCheck to create.
	ReadinessCheckName string `pulumi:"readinessCheckName"`
}

type LookupReadinessCheckResult struct {
	// The Amazon Resource Name (ARN) of the readiness check.
	ReadinessCheckArn *string `pulumi:"readinessCheckArn"`
	// The name of the resource set to check.
	ResourceSetName *string `pulumi:"resourceSetName"`
	// A collection of tags associated with a resource.
	Tags []ReadinessCheckTag `pulumi:"tags"`
}

func LookupReadinessCheckOutput(ctx *pulumi.Context, args LookupReadinessCheckOutputArgs, opts ...pulumi.InvokeOption) LookupReadinessCheckResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupReadinessCheckResult, error) {
			args := v.(LookupReadinessCheckArgs)
			r, err := LookupReadinessCheck(ctx, &args, opts...)
			var s LookupReadinessCheckResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupReadinessCheckResultOutput)
}

type LookupReadinessCheckOutputArgs struct {
	// Name of the ReadinessCheck to create.
	ReadinessCheckName pulumi.StringInput `pulumi:"readinessCheckName"`
}

func (LookupReadinessCheckOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupReadinessCheckArgs)(nil)).Elem()
}

type LookupReadinessCheckResultOutput struct{ *pulumi.OutputState }

func (LookupReadinessCheckResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupReadinessCheckResult)(nil)).Elem()
}

func (o LookupReadinessCheckResultOutput) ToLookupReadinessCheckResultOutput() LookupReadinessCheckResultOutput {
	return o
}

func (o LookupReadinessCheckResultOutput) ToLookupReadinessCheckResultOutputWithContext(ctx context.Context) LookupReadinessCheckResultOutput {
	return o
}

// The Amazon Resource Name (ARN) of the readiness check.
func (o LookupReadinessCheckResultOutput) ReadinessCheckArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupReadinessCheckResult) *string { return v.ReadinessCheckArn }).(pulumi.StringPtrOutput)
}

// The name of the resource set to check.
func (o LookupReadinessCheckResultOutput) ResourceSetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupReadinessCheckResult) *string { return v.ResourceSetName }).(pulumi.StringPtrOutput)
}

// A collection of tags associated with a resource.
func (o LookupReadinessCheckResultOutput) Tags() ReadinessCheckTagArrayOutput {
	return o.ApplyT(func(v LookupReadinessCheckResult) []ReadinessCheckTag { return v.Tags }).(ReadinessCheckTagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupReadinessCheckResultOutput{})
}
