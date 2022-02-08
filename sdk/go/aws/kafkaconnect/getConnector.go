// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package kafkaconnect

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::KafkaConnect::Connector
func LookupConnector(ctx *pulumi.Context, args *LookupConnectorArgs, opts ...pulumi.InvokeOption) (*LookupConnectorResult, error) {
	var rv LookupConnectorResult
	err := ctx.Invoke("aws-native:kafkaconnect:getConnector", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupConnectorArgs struct {
	// Amazon Resource Name for the created Connector.
	ConnectorArn string `pulumi:"connectorArn"`
}

type LookupConnectorResult struct {
	Capacity *ConnectorCapacity `pulumi:"capacity"`
	// Amazon Resource Name for the created Connector.
	ConnectorArn *string `pulumi:"connectorArn"`
}

func LookupConnectorOutput(ctx *pulumi.Context, args LookupConnectorOutputArgs, opts ...pulumi.InvokeOption) LookupConnectorResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupConnectorResult, error) {
			args := v.(LookupConnectorArgs)
			r, err := LookupConnector(ctx, &args, opts...)
			return *r, err
		}).(LookupConnectorResultOutput)
}

type LookupConnectorOutputArgs struct {
	// Amazon Resource Name for the created Connector.
	ConnectorArn pulumi.StringInput `pulumi:"connectorArn"`
}

func (LookupConnectorOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConnectorArgs)(nil)).Elem()
}

type LookupConnectorResultOutput struct{ *pulumi.OutputState }

func (LookupConnectorResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConnectorResult)(nil)).Elem()
}

func (o LookupConnectorResultOutput) ToLookupConnectorResultOutput() LookupConnectorResultOutput {
	return o
}

func (o LookupConnectorResultOutput) ToLookupConnectorResultOutputWithContext(ctx context.Context) LookupConnectorResultOutput {
	return o
}

func (o LookupConnectorResultOutput) Capacity() ConnectorCapacityPtrOutput {
	return o.ApplyT(func(v LookupConnectorResult) *ConnectorCapacity { return v.Capacity }).(ConnectorCapacityPtrOutput)
}

// Amazon Resource Name for the created Connector.
func (o LookupConnectorResultOutput) ConnectorArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConnectorResult) *string { return v.ConnectorArn }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupConnectorResultOutput{})
}
