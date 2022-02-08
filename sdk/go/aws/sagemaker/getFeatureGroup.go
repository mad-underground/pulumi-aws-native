// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package sagemaker

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::SageMaker::FeatureGroup
func LookupFeatureGroup(ctx *pulumi.Context, args *LookupFeatureGroupArgs, opts ...pulumi.InvokeOption) (*LookupFeatureGroupResult, error) {
	var rv LookupFeatureGroupResult
	err := ctx.Invoke("aws-native:sagemaker:getFeatureGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupFeatureGroupArgs struct {
	// The Name of the FeatureGroup.
	FeatureGroupName string `pulumi:"featureGroupName"`
}

type LookupFeatureGroupResult struct {
}

func LookupFeatureGroupOutput(ctx *pulumi.Context, args LookupFeatureGroupOutputArgs, opts ...pulumi.InvokeOption) LookupFeatureGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFeatureGroupResult, error) {
			args := v.(LookupFeatureGroupArgs)
			r, err := LookupFeatureGroup(ctx, &args, opts...)
			return *r, err
		}).(LookupFeatureGroupResultOutput)
}

type LookupFeatureGroupOutputArgs struct {
	// The Name of the FeatureGroup.
	FeatureGroupName pulumi.StringInput `pulumi:"featureGroupName"`
}

func (LookupFeatureGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFeatureGroupArgs)(nil)).Elem()
}

type LookupFeatureGroupResultOutput struct{ *pulumi.OutputState }

func (LookupFeatureGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFeatureGroupResult)(nil)).Elem()
}

func (o LookupFeatureGroupResultOutput) ToLookupFeatureGroupResultOutput() LookupFeatureGroupResultOutput {
	return o
}

func (o LookupFeatureGroupResultOutput) ToLookupFeatureGroupResultOutputWithContext(ctx context.Context) LookupFeatureGroupResultOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(LookupFeatureGroupResultOutput{})
}
