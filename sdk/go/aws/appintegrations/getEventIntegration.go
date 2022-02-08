// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package appintegrations

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::AppIntegrations::EventIntegration
func LookupEventIntegration(ctx *pulumi.Context, args *LookupEventIntegrationArgs, opts ...pulumi.InvokeOption) (*LookupEventIntegrationResult, error) {
	var rv LookupEventIntegrationResult
	err := ctx.Invoke("aws-native:appintegrations:getEventIntegration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupEventIntegrationArgs struct {
	// The name of the event integration.
	Name string `pulumi:"name"`
}

type LookupEventIntegrationResult struct {
	// The associations with the event integration.
	Associations []EventIntegrationAssociation `pulumi:"associations"`
	// The event integration description.
	Description *string `pulumi:"description"`
	// The Amazon Resource Name (ARN) of the event integration.
	EventIntegrationArn *string `pulumi:"eventIntegrationArn"`
	// The tags (keys and values) associated with the event integration.
	Tags []EventIntegrationTag `pulumi:"tags"`
}

func LookupEventIntegrationOutput(ctx *pulumi.Context, args LookupEventIntegrationOutputArgs, opts ...pulumi.InvokeOption) LookupEventIntegrationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupEventIntegrationResult, error) {
			args := v.(LookupEventIntegrationArgs)
			r, err := LookupEventIntegration(ctx, &args, opts...)
			return *r, err
		}).(LookupEventIntegrationResultOutput)
}

type LookupEventIntegrationOutputArgs struct {
	// The name of the event integration.
	Name pulumi.StringInput `pulumi:"name"`
}

func (LookupEventIntegrationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEventIntegrationArgs)(nil)).Elem()
}

type LookupEventIntegrationResultOutput struct{ *pulumi.OutputState }

func (LookupEventIntegrationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEventIntegrationResult)(nil)).Elem()
}

func (o LookupEventIntegrationResultOutput) ToLookupEventIntegrationResultOutput() LookupEventIntegrationResultOutput {
	return o
}

func (o LookupEventIntegrationResultOutput) ToLookupEventIntegrationResultOutputWithContext(ctx context.Context) LookupEventIntegrationResultOutput {
	return o
}

// The associations with the event integration.
func (o LookupEventIntegrationResultOutput) Associations() EventIntegrationAssociationArrayOutput {
	return o.ApplyT(func(v LookupEventIntegrationResult) []EventIntegrationAssociation { return v.Associations }).(EventIntegrationAssociationArrayOutput)
}

// The event integration description.
func (o LookupEventIntegrationResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEventIntegrationResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// The Amazon Resource Name (ARN) of the event integration.
func (o LookupEventIntegrationResultOutput) EventIntegrationArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEventIntegrationResult) *string { return v.EventIntegrationArn }).(pulumi.StringPtrOutput)
}

// The tags (keys and values) associated with the event integration.
func (o LookupEventIntegrationResultOutput) Tags() EventIntegrationTagArrayOutput {
	return o.ApplyT(func(v LookupEventIntegrationResult) []EventIntegrationTag { return v.Tags }).(EventIntegrationTagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupEventIntegrationResultOutput{})
}
