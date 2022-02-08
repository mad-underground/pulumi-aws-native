// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package autoscaling

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The AWS::AutoScaling::LaunchConfiguration resource specifies the launch configuration that can be used by an Auto Scaling group to configure Amazon EC2 instances.
func LookupLaunchConfiguration(ctx *pulumi.Context, args *LookupLaunchConfigurationArgs, opts ...pulumi.InvokeOption) (*LookupLaunchConfigurationResult, error) {
	var rv LookupLaunchConfigurationResult
	err := ctx.Invoke("aws-native:autoscaling:getLaunchConfiguration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupLaunchConfigurationArgs struct {
	// The name of the launch configuration. This name must be unique per Region per account.
	LaunchConfigurationName string `pulumi:"launchConfigurationName"`
}

type LookupLaunchConfigurationResult struct {
}

func LookupLaunchConfigurationOutput(ctx *pulumi.Context, args LookupLaunchConfigurationOutputArgs, opts ...pulumi.InvokeOption) LookupLaunchConfigurationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupLaunchConfigurationResult, error) {
			args := v.(LookupLaunchConfigurationArgs)
			r, err := LookupLaunchConfiguration(ctx, &args, opts...)
			return *r, err
		}).(LookupLaunchConfigurationResultOutput)
}

type LookupLaunchConfigurationOutputArgs struct {
	// The name of the launch configuration. This name must be unique per Region per account.
	LaunchConfigurationName pulumi.StringInput `pulumi:"launchConfigurationName"`
}

func (LookupLaunchConfigurationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLaunchConfigurationArgs)(nil)).Elem()
}

type LookupLaunchConfigurationResultOutput struct{ *pulumi.OutputState }

func (LookupLaunchConfigurationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLaunchConfigurationResult)(nil)).Elem()
}

func (o LookupLaunchConfigurationResultOutput) ToLookupLaunchConfigurationResultOutput() LookupLaunchConfigurationResultOutput {
	return o
}

func (o LookupLaunchConfigurationResultOutput) ToLookupLaunchConfigurationResultOutputWithContext(ctx context.Context) LookupLaunchConfigurationResultOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(LookupLaunchConfigurationResultOutput{})
}
