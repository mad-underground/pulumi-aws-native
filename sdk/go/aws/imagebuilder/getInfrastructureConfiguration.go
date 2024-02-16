// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package imagebuilder

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource schema for AWS::ImageBuilder::InfrastructureConfiguration
func LookupInfrastructureConfiguration(ctx *pulumi.Context, args *LookupInfrastructureConfigurationArgs, opts ...pulumi.InvokeOption) (*LookupInfrastructureConfigurationResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupInfrastructureConfigurationResult
	err := ctx.Invoke("aws-native:imagebuilder:getInfrastructureConfiguration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupInfrastructureConfigurationArgs struct {
	// The Amazon Resource Name (ARN) of the infrastructure configuration.
	Arn string `pulumi:"arn"`
}

type LookupInfrastructureConfigurationResult struct {
	// The Amazon Resource Name (ARN) of the infrastructure configuration.
	Arn *string `pulumi:"arn"`
	// The description of the infrastructure configuration.
	Description *string `pulumi:"description"`
	// The instance metadata option settings for the infrastructure configuration.
	InstanceMetadataOptions *InfrastructureConfigurationInstanceMetadataOptions `pulumi:"instanceMetadataOptions"`
	// The instance profile of the infrastructure configuration.
	InstanceProfileName *string `pulumi:"instanceProfileName"`
	// The instance types of the infrastructure configuration.
	InstanceTypes []string `pulumi:"instanceTypes"`
	// The EC2 key pair of the infrastructure configuration..
	KeyPair *string `pulumi:"keyPair"`
	// The logging configuration of the infrastructure configuration.
	Logging *InfrastructureConfigurationLogging `pulumi:"logging"`
	// The tags attached to the resource created by Image Builder.
	ResourceTags map[string]string `pulumi:"resourceTags"`
	// The security group IDs of the infrastructure configuration.
	SecurityGroupIds []string `pulumi:"securityGroupIds"`
	// The SNS Topic Amazon Resource Name (ARN) of the infrastructure configuration.
	SnsTopicArn *string `pulumi:"snsTopicArn"`
	// The subnet ID of the infrastructure configuration.
	SubnetId *string `pulumi:"subnetId"`
	// The tags associated with the component.
	Tags map[string]string `pulumi:"tags"`
	// The terminate instance on failure configuration of the infrastructure configuration.
	TerminateInstanceOnFailure *bool `pulumi:"terminateInstanceOnFailure"`
}

func LookupInfrastructureConfigurationOutput(ctx *pulumi.Context, args LookupInfrastructureConfigurationOutputArgs, opts ...pulumi.InvokeOption) LookupInfrastructureConfigurationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupInfrastructureConfigurationResult, error) {
			args := v.(LookupInfrastructureConfigurationArgs)
			r, err := LookupInfrastructureConfiguration(ctx, &args, opts...)
			var s LookupInfrastructureConfigurationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupInfrastructureConfigurationResultOutput)
}

type LookupInfrastructureConfigurationOutputArgs struct {
	// The Amazon Resource Name (ARN) of the infrastructure configuration.
	Arn pulumi.StringInput `pulumi:"arn"`
}

func (LookupInfrastructureConfigurationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInfrastructureConfigurationArgs)(nil)).Elem()
}

type LookupInfrastructureConfigurationResultOutput struct{ *pulumi.OutputState }

func (LookupInfrastructureConfigurationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInfrastructureConfigurationResult)(nil)).Elem()
}

func (o LookupInfrastructureConfigurationResultOutput) ToLookupInfrastructureConfigurationResultOutput() LookupInfrastructureConfigurationResultOutput {
	return o
}

func (o LookupInfrastructureConfigurationResultOutput) ToLookupInfrastructureConfigurationResultOutputWithContext(ctx context.Context) LookupInfrastructureConfigurationResultOutput {
	return o
}

// The Amazon Resource Name (ARN) of the infrastructure configuration.
func (o LookupInfrastructureConfigurationResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInfrastructureConfigurationResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

// The description of the infrastructure configuration.
func (o LookupInfrastructureConfigurationResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInfrastructureConfigurationResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// The instance metadata option settings for the infrastructure configuration.
func (o LookupInfrastructureConfigurationResultOutput) InstanceMetadataOptions() InfrastructureConfigurationInstanceMetadataOptionsPtrOutput {
	return o.ApplyT(func(v LookupInfrastructureConfigurationResult) *InfrastructureConfigurationInstanceMetadataOptions {
		return v.InstanceMetadataOptions
	}).(InfrastructureConfigurationInstanceMetadataOptionsPtrOutput)
}

// The instance profile of the infrastructure configuration.
func (o LookupInfrastructureConfigurationResultOutput) InstanceProfileName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInfrastructureConfigurationResult) *string { return v.InstanceProfileName }).(pulumi.StringPtrOutput)
}

// The instance types of the infrastructure configuration.
func (o LookupInfrastructureConfigurationResultOutput) InstanceTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupInfrastructureConfigurationResult) []string { return v.InstanceTypes }).(pulumi.StringArrayOutput)
}

// The EC2 key pair of the infrastructure configuration..
func (o LookupInfrastructureConfigurationResultOutput) KeyPair() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInfrastructureConfigurationResult) *string { return v.KeyPair }).(pulumi.StringPtrOutput)
}

// The logging configuration of the infrastructure configuration.
func (o LookupInfrastructureConfigurationResultOutput) Logging() InfrastructureConfigurationLoggingPtrOutput {
	return o.ApplyT(func(v LookupInfrastructureConfigurationResult) *InfrastructureConfigurationLogging { return v.Logging }).(InfrastructureConfigurationLoggingPtrOutput)
}

// The tags attached to the resource created by Image Builder.
func (o LookupInfrastructureConfigurationResultOutput) ResourceTags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupInfrastructureConfigurationResult) map[string]string { return v.ResourceTags }).(pulumi.StringMapOutput)
}

// The security group IDs of the infrastructure configuration.
func (o LookupInfrastructureConfigurationResultOutput) SecurityGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupInfrastructureConfigurationResult) []string { return v.SecurityGroupIds }).(pulumi.StringArrayOutput)
}

// The SNS Topic Amazon Resource Name (ARN) of the infrastructure configuration.
func (o LookupInfrastructureConfigurationResultOutput) SnsTopicArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInfrastructureConfigurationResult) *string { return v.SnsTopicArn }).(pulumi.StringPtrOutput)
}

// The subnet ID of the infrastructure configuration.
func (o LookupInfrastructureConfigurationResultOutput) SubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInfrastructureConfigurationResult) *string { return v.SubnetId }).(pulumi.StringPtrOutput)
}

// The tags associated with the component.
func (o LookupInfrastructureConfigurationResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupInfrastructureConfigurationResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The terminate instance on failure configuration of the infrastructure configuration.
func (o LookupInfrastructureConfigurationResultOutput) TerminateInstanceOnFailure() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupInfrastructureConfigurationResult) *bool { return v.TerminateInstanceOnFailure }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupInfrastructureConfigurationResultOutput{})
}
