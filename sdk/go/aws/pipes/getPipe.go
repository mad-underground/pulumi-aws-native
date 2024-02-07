// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package pipes

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Definition of AWS::Pipes::Pipe Resource Type
func LookupPipe(ctx *pulumi.Context, args *LookupPipeArgs, opts ...pulumi.InvokeOption) (*LookupPipeResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupPipeResult
	err := ctx.Invoke("aws-native:pipes:getPipe", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPipeArgs struct {
	Name string `pulumi:"name"`
}

type LookupPipeResult struct {
	Arn                  *string                   `pulumi:"arn"`
	CreationTime         *string                   `pulumi:"creationTime"`
	CurrentState         *PipeStateEnum            `pulumi:"currentState"`
	Description          *string                   `pulumi:"description"`
	DesiredState         *PipeRequestedPipeState   `pulumi:"desiredState"`
	Enrichment           *string                   `pulumi:"enrichment"`
	EnrichmentParameters *PipeEnrichmentParameters `pulumi:"enrichmentParameters"`
	LastModifiedTime     *string                   `pulumi:"lastModifiedTime"`
	LogConfiguration     *PipeLogConfiguration     `pulumi:"logConfiguration"`
	RoleArn              *string                   `pulumi:"roleArn"`
	StateReason          *string                   `pulumi:"stateReason"`
	Tags                 *PipeTagMap               `pulumi:"tags"`
	Target               *string                   `pulumi:"target"`
}

func LookupPipeOutput(ctx *pulumi.Context, args LookupPipeOutputArgs, opts ...pulumi.InvokeOption) LookupPipeResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPipeResult, error) {
			args := v.(LookupPipeArgs)
			r, err := LookupPipe(ctx, &args, opts...)
			var s LookupPipeResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPipeResultOutput)
}

type LookupPipeOutputArgs struct {
	Name pulumi.StringInput `pulumi:"name"`
}

func (LookupPipeOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPipeArgs)(nil)).Elem()
}

type LookupPipeResultOutput struct{ *pulumi.OutputState }

func (LookupPipeResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPipeResult)(nil)).Elem()
}

func (o LookupPipeResultOutput) ToLookupPipeResultOutput() LookupPipeResultOutput {
	return o
}

func (o LookupPipeResultOutput) ToLookupPipeResultOutputWithContext(ctx context.Context) LookupPipeResultOutput {
	return o
}

func (o LookupPipeResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPipeResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

func (o LookupPipeResultOutput) CreationTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPipeResult) *string { return v.CreationTime }).(pulumi.StringPtrOutput)
}

func (o LookupPipeResultOutput) CurrentState() PipeStateEnumPtrOutput {
	return o.ApplyT(func(v LookupPipeResult) *PipeStateEnum { return v.CurrentState }).(PipeStateEnumPtrOutput)
}

func (o LookupPipeResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPipeResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupPipeResultOutput) DesiredState() PipeRequestedPipeStatePtrOutput {
	return o.ApplyT(func(v LookupPipeResult) *PipeRequestedPipeState { return v.DesiredState }).(PipeRequestedPipeStatePtrOutput)
}

func (o LookupPipeResultOutput) Enrichment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPipeResult) *string { return v.Enrichment }).(pulumi.StringPtrOutput)
}

func (o LookupPipeResultOutput) EnrichmentParameters() PipeEnrichmentParametersPtrOutput {
	return o.ApplyT(func(v LookupPipeResult) *PipeEnrichmentParameters { return v.EnrichmentParameters }).(PipeEnrichmentParametersPtrOutput)
}

func (o LookupPipeResultOutput) LastModifiedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPipeResult) *string { return v.LastModifiedTime }).(pulumi.StringPtrOutput)
}

func (o LookupPipeResultOutput) LogConfiguration() PipeLogConfigurationPtrOutput {
	return o.ApplyT(func(v LookupPipeResult) *PipeLogConfiguration { return v.LogConfiguration }).(PipeLogConfigurationPtrOutput)
}

func (o LookupPipeResultOutput) RoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPipeResult) *string { return v.RoleArn }).(pulumi.StringPtrOutput)
}

func (o LookupPipeResultOutput) StateReason() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPipeResult) *string { return v.StateReason }).(pulumi.StringPtrOutput)
}

func (o LookupPipeResultOutput) Tags() PipeTagMapPtrOutput {
	return o.ApplyT(func(v LookupPipeResult) *PipeTagMap { return v.Tags }).(PipeTagMapPtrOutput)
}

func (o LookupPipeResultOutput) Target() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPipeResult) *string { return v.Target }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPipeResultOutput{})
}
