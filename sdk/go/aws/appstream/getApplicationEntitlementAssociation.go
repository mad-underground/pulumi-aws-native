// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package appstream

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::AppStream::ApplicationEntitlementAssociation
func LookupApplicationEntitlementAssociation(ctx *pulumi.Context, args *LookupApplicationEntitlementAssociationArgs, opts ...pulumi.InvokeOption) (*LookupApplicationEntitlementAssociationResult, error) {
	var rv LookupApplicationEntitlementAssociationResult
	err := ctx.Invoke("aws-native:appstream:getApplicationEntitlementAssociation", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupApplicationEntitlementAssociationArgs struct {
	ApplicationIdentifier string `pulumi:"applicationIdentifier"`
	EntitlementName       string `pulumi:"entitlementName"`
	StackName             string `pulumi:"stackName"`
}

type LookupApplicationEntitlementAssociationResult struct {
}

func LookupApplicationEntitlementAssociationOutput(ctx *pulumi.Context, args LookupApplicationEntitlementAssociationOutputArgs, opts ...pulumi.InvokeOption) LookupApplicationEntitlementAssociationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupApplicationEntitlementAssociationResult, error) {
			args := v.(LookupApplicationEntitlementAssociationArgs)
			r, err := LookupApplicationEntitlementAssociation(ctx, &args, opts...)
			return *r, err
		}).(LookupApplicationEntitlementAssociationResultOutput)
}

type LookupApplicationEntitlementAssociationOutputArgs struct {
	ApplicationIdentifier pulumi.StringInput `pulumi:"applicationIdentifier"`
	EntitlementName       pulumi.StringInput `pulumi:"entitlementName"`
	StackName             pulumi.StringInput `pulumi:"stackName"`
}

func (LookupApplicationEntitlementAssociationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApplicationEntitlementAssociationArgs)(nil)).Elem()
}

type LookupApplicationEntitlementAssociationResultOutput struct{ *pulumi.OutputState }

func (LookupApplicationEntitlementAssociationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApplicationEntitlementAssociationResult)(nil)).Elem()
}

func (o LookupApplicationEntitlementAssociationResultOutput) ToLookupApplicationEntitlementAssociationResultOutput() LookupApplicationEntitlementAssociationResultOutput {
	return o
}

func (o LookupApplicationEntitlementAssociationResultOutput) ToLookupApplicationEntitlementAssociationResultOutputWithContext(ctx context.Context) LookupApplicationEntitlementAssociationResultOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(LookupApplicationEntitlementAssociationResultOutput{})
}
