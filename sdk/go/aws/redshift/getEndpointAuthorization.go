// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package redshift

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Describes an endpoint authorization for authorizing Redshift-managed VPC endpoint access to a cluster across AWS accounts.
func LookupEndpointAuthorization(ctx *pulumi.Context, args *LookupEndpointAuthorizationArgs, opts ...pulumi.InvokeOption) (*LookupEndpointAuthorizationResult, error) {
	var rv LookupEndpointAuthorizationResult
	err := ctx.Invoke("aws-native:redshift:getEndpointAuthorization", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupEndpointAuthorizationArgs struct {
	// The target AWS account ID to grant or revoke access for.
	Account string `pulumi:"account"`
	// The cluster identifier.
	ClusterIdentifier string `pulumi:"clusterIdentifier"`
}

type LookupEndpointAuthorizationResult struct {
	// Indicates whether all VPCs in the grantee account are allowed access to the cluster.
	AllowedAllVPCs *bool `pulumi:"allowedAllVPCs"`
	// The VPCs allowed access to the cluster.
	AllowedVPCs []string `pulumi:"allowedVPCs"`
	// The time (UTC) when the authorization was created.
	AuthorizeTime *string `pulumi:"authorizeTime"`
	// The status of the cluster.
	ClusterStatus *string `pulumi:"clusterStatus"`
	// The number of Redshift-managed VPC endpoints created for the authorization.
	EndpointCount *int `pulumi:"endpointCount"`
	// The AWS account ID of the grantee of the cluster.
	Grantee *string `pulumi:"grantee"`
	// The AWS account ID of the cluster owner.
	Grantor *string `pulumi:"grantor"`
	// The status of the authorization action.
	Status *string `pulumi:"status"`
	// The virtual private cloud (VPC) identifiers to grant or revoke access to.
	VpcIds []string `pulumi:"vpcIds"`
}

func LookupEndpointAuthorizationOutput(ctx *pulumi.Context, args LookupEndpointAuthorizationOutputArgs, opts ...pulumi.InvokeOption) LookupEndpointAuthorizationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupEndpointAuthorizationResult, error) {
			args := v.(LookupEndpointAuthorizationArgs)
			r, err := LookupEndpointAuthorization(ctx, &args, opts...)
			return *r, err
		}).(LookupEndpointAuthorizationResultOutput)
}

type LookupEndpointAuthorizationOutputArgs struct {
	// The target AWS account ID to grant or revoke access for.
	Account pulumi.StringInput `pulumi:"account"`
	// The cluster identifier.
	ClusterIdentifier pulumi.StringInput `pulumi:"clusterIdentifier"`
}

func (LookupEndpointAuthorizationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEndpointAuthorizationArgs)(nil)).Elem()
}

type LookupEndpointAuthorizationResultOutput struct{ *pulumi.OutputState }

func (LookupEndpointAuthorizationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEndpointAuthorizationResult)(nil)).Elem()
}

func (o LookupEndpointAuthorizationResultOutput) ToLookupEndpointAuthorizationResultOutput() LookupEndpointAuthorizationResultOutput {
	return o
}

func (o LookupEndpointAuthorizationResultOutput) ToLookupEndpointAuthorizationResultOutputWithContext(ctx context.Context) LookupEndpointAuthorizationResultOutput {
	return o
}

// Indicates whether all VPCs in the grantee account are allowed access to the cluster.
func (o LookupEndpointAuthorizationResultOutput) AllowedAllVPCs() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupEndpointAuthorizationResult) *bool { return v.AllowedAllVPCs }).(pulumi.BoolPtrOutput)
}

// The VPCs allowed access to the cluster.
func (o LookupEndpointAuthorizationResultOutput) AllowedVPCs() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupEndpointAuthorizationResult) []string { return v.AllowedVPCs }).(pulumi.StringArrayOutput)
}

// The time (UTC) when the authorization was created.
func (o LookupEndpointAuthorizationResultOutput) AuthorizeTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEndpointAuthorizationResult) *string { return v.AuthorizeTime }).(pulumi.StringPtrOutput)
}

// The status of the cluster.
func (o LookupEndpointAuthorizationResultOutput) ClusterStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEndpointAuthorizationResult) *string { return v.ClusterStatus }).(pulumi.StringPtrOutput)
}

// The number of Redshift-managed VPC endpoints created for the authorization.
func (o LookupEndpointAuthorizationResultOutput) EndpointCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupEndpointAuthorizationResult) *int { return v.EndpointCount }).(pulumi.IntPtrOutput)
}

// The AWS account ID of the grantee of the cluster.
func (o LookupEndpointAuthorizationResultOutput) Grantee() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEndpointAuthorizationResult) *string { return v.Grantee }).(pulumi.StringPtrOutput)
}

// The AWS account ID of the cluster owner.
func (o LookupEndpointAuthorizationResultOutput) Grantor() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEndpointAuthorizationResult) *string { return v.Grantor }).(pulumi.StringPtrOutput)
}

// The status of the authorization action.
func (o LookupEndpointAuthorizationResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEndpointAuthorizationResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

// The virtual private cloud (VPC) identifiers to grant or revoke access to.
func (o LookupEndpointAuthorizationResultOutput) VpcIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupEndpointAuthorizationResult) []string { return v.VpcIds }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupEndpointAuthorizationResultOutput{})
}