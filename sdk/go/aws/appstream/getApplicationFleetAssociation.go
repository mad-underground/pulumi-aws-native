// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package appstream

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::AppStream::ApplicationFleetAssociation
func LookupApplicationFleetAssociation(ctx *pulumi.Context, args *LookupApplicationFleetAssociationArgs, opts ...pulumi.InvokeOption) (*LookupApplicationFleetAssociationResult, error) {
	var rv LookupApplicationFleetAssociationResult
	err := ctx.Invoke("aws-native:appstream:getApplicationFleetAssociation", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupApplicationFleetAssociationArgs struct {
	ApplicationArn string `pulumi:"applicationArn"`
	FleetName      string `pulumi:"fleetName"`
}

type LookupApplicationFleetAssociationResult struct {
}

func LookupApplicationFleetAssociationOutput(ctx *pulumi.Context, args LookupApplicationFleetAssociationOutputArgs, opts ...pulumi.InvokeOption) LookupApplicationFleetAssociationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupApplicationFleetAssociationResult, error) {
			args := v.(LookupApplicationFleetAssociationArgs)
			r, err := LookupApplicationFleetAssociation(ctx, &args, opts...)
			return *r, err
		}).(LookupApplicationFleetAssociationResultOutput)
}

type LookupApplicationFleetAssociationOutputArgs struct {
	ApplicationArn pulumi.StringInput `pulumi:"applicationArn"`
	FleetName      pulumi.StringInput `pulumi:"fleetName"`
}

func (LookupApplicationFleetAssociationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApplicationFleetAssociationArgs)(nil)).Elem()
}

type LookupApplicationFleetAssociationResultOutput struct{ *pulumi.OutputState }

func (LookupApplicationFleetAssociationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApplicationFleetAssociationResult)(nil)).Elem()
}

func (o LookupApplicationFleetAssociationResultOutput) ToLookupApplicationFleetAssociationResultOutput() LookupApplicationFleetAssociationResultOutput {
	return o
}

func (o LookupApplicationFleetAssociationResultOutput) ToLookupApplicationFleetAssociationResultOutputWithContext(ctx context.Context) LookupApplicationFleetAssociationResultOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(LookupApplicationFleetAssociationResultOutput{})
}
