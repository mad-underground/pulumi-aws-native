// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package events

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Events::Endpoint.
func LookupEndpoint(ctx *pulumi.Context, args *LookupEndpointArgs, opts ...pulumi.InvokeOption) (*LookupEndpointResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupEndpointResult
	err := ctx.Invoke("aws-native:events:getEndpoint", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupEndpointArgs struct {
	Name string `pulumi:"name"`
}

type LookupEndpointResult struct {
	Arn               *string                    `pulumi:"arn"`
	Description       *string                    `pulumi:"description"`
	EndpointId        *string                    `pulumi:"endpointId"`
	EndpointUrl       *string                    `pulumi:"endpointUrl"`
	EventBuses        []EndpointEventBus         `pulumi:"eventBuses"`
	ReplicationConfig *EndpointReplicationConfig `pulumi:"replicationConfig"`
	RoleArn           *string                    `pulumi:"roleArn"`
	RoutingConfig     *EndpointRoutingConfig     `pulumi:"routingConfig"`
	State             *EndpointStateEnum         `pulumi:"state"`
	StateReason       *string                    `pulumi:"stateReason"`
}

func LookupEndpointOutput(ctx *pulumi.Context, args LookupEndpointOutputArgs, opts ...pulumi.InvokeOption) LookupEndpointResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupEndpointResult, error) {
			args := v.(LookupEndpointArgs)
			r, err := LookupEndpoint(ctx, &args, opts...)
			var s LookupEndpointResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupEndpointResultOutput)
}

type LookupEndpointOutputArgs struct {
	Name pulumi.StringInput `pulumi:"name"`
}

func (LookupEndpointOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEndpointArgs)(nil)).Elem()
}

type LookupEndpointResultOutput struct{ *pulumi.OutputState }

func (LookupEndpointResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEndpointResult)(nil)).Elem()
}

func (o LookupEndpointResultOutput) ToLookupEndpointResultOutput() LookupEndpointResultOutput {
	return o
}

func (o LookupEndpointResultOutput) ToLookupEndpointResultOutputWithContext(ctx context.Context) LookupEndpointResultOutput {
	return o
}

func (o LookupEndpointResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEndpointResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

func (o LookupEndpointResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEndpointResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupEndpointResultOutput) EndpointId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEndpointResult) *string { return v.EndpointId }).(pulumi.StringPtrOutput)
}

func (o LookupEndpointResultOutput) EndpointUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEndpointResult) *string { return v.EndpointUrl }).(pulumi.StringPtrOutput)
}

func (o LookupEndpointResultOutput) EventBuses() EndpointEventBusArrayOutput {
	return o.ApplyT(func(v LookupEndpointResult) []EndpointEventBus { return v.EventBuses }).(EndpointEventBusArrayOutput)
}

func (o LookupEndpointResultOutput) ReplicationConfig() EndpointReplicationConfigPtrOutput {
	return o.ApplyT(func(v LookupEndpointResult) *EndpointReplicationConfig { return v.ReplicationConfig }).(EndpointReplicationConfigPtrOutput)
}

func (o LookupEndpointResultOutput) RoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEndpointResult) *string { return v.RoleArn }).(pulumi.StringPtrOutput)
}

func (o LookupEndpointResultOutput) RoutingConfig() EndpointRoutingConfigPtrOutput {
	return o.ApplyT(func(v LookupEndpointResult) *EndpointRoutingConfig { return v.RoutingConfig }).(EndpointRoutingConfigPtrOutput)
}

func (o LookupEndpointResultOutput) State() EndpointStateEnumPtrOutput {
	return o.ApplyT(func(v LookupEndpointResult) *EndpointStateEnum { return v.State }).(EndpointStateEnumPtrOutput)
}

func (o LookupEndpointResultOutput) StateReason() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEndpointResult) *string { return v.StateReason }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupEndpointResultOutput{})
}
