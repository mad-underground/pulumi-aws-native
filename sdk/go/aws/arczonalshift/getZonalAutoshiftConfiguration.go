// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package arczonalshift

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Definition of AWS::ARCZonalShift::ZonalAutoshiftConfiguration Resource Type
func LookupZonalAutoshiftConfiguration(ctx *pulumi.Context, args *LookupZonalAutoshiftConfigurationArgs, opts ...pulumi.InvokeOption) (*LookupZonalAutoshiftConfigurationResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupZonalAutoshiftConfigurationResult
	err := ctx.Invoke("aws-native:arczonalshift:getZonalAutoshiftConfiguration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupZonalAutoshiftConfigurationArgs struct {
	ResourceIdentifier string `pulumi:"resourceIdentifier"`
}

type LookupZonalAutoshiftConfigurationResult struct {
	PracticeRunConfiguration *ZonalAutoshiftConfigurationPracticeRunConfiguration `pulumi:"practiceRunConfiguration"`
	ZonalAutoshiftStatus     *ZonalAutoshiftConfigurationZonalAutoshiftStatus     `pulumi:"zonalAutoshiftStatus"`
}

func LookupZonalAutoshiftConfigurationOutput(ctx *pulumi.Context, args LookupZonalAutoshiftConfigurationOutputArgs, opts ...pulumi.InvokeOption) LookupZonalAutoshiftConfigurationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupZonalAutoshiftConfigurationResult, error) {
			args := v.(LookupZonalAutoshiftConfigurationArgs)
			r, err := LookupZonalAutoshiftConfiguration(ctx, &args, opts...)
			var s LookupZonalAutoshiftConfigurationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupZonalAutoshiftConfigurationResultOutput)
}

type LookupZonalAutoshiftConfigurationOutputArgs struct {
	ResourceIdentifier pulumi.StringInput `pulumi:"resourceIdentifier"`
}

func (LookupZonalAutoshiftConfigurationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupZonalAutoshiftConfigurationArgs)(nil)).Elem()
}

type LookupZonalAutoshiftConfigurationResultOutput struct{ *pulumi.OutputState }

func (LookupZonalAutoshiftConfigurationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupZonalAutoshiftConfigurationResult)(nil)).Elem()
}

func (o LookupZonalAutoshiftConfigurationResultOutput) ToLookupZonalAutoshiftConfigurationResultOutput() LookupZonalAutoshiftConfigurationResultOutput {
	return o
}

func (o LookupZonalAutoshiftConfigurationResultOutput) ToLookupZonalAutoshiftConfigurationResultOutputWithContext(ctx context.Context) LookupZonalAutoshiftConfigurationResultOutput {
	return o
}

func (o LookupZonalAutoshiftConfigurationResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupZonalAutoshiftConfigurationResult] {
	return pulumix.Output[LookupZonalAutoshiftConfigurationResult]{
		OutputState: o.OutputState,
	}
}

func (o LookupZonalAutoshiftConfigurationResultOutput) PracticeRunConfiguration() ZonalAutoshiftConfigurationPracticeRunConfigurationPtrOutput {
	return o.ApplyT(func(v LookupZonalAutoshiftConfigurationResult) *ZonalAutoshiftConfigurationPracticeRunConfiguration {
		return v.PracticeRunConfiguration
	}).(ZonalAutoshiftConfigurationPracticeRunConfigurationPtrOutput)
}

func (o LookupZonalAutoshiftConfigurationResultOutput) ZonalAutoshiftStatus() ZonalAutoshiftConfigurationZonalAutoshiftStatusPtrOutput {
	return o.ApplyT(func(v LookupZonalAutoshiftConfigurationResult) *ZonalAutoshiftConfigurationZonalAutoshiftStatus {
		return v.ZonalAutoshiftStatus
	}).(ZonalAutoshiftConfigurationZonalAutoshiftStatusPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupZonalAutoshiftConfigurationResultOutput{})
}
