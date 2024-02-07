// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package healthimaging

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Definition of AWS::HealthImaging::Datastore Resource Type
func LookupDatastore(ctx *pulumi.Context, args *LookupDatastoreArgs, opts ...pulumi.InvokeOption) (*LookupDatastoreResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupDatastoreResult
	err := ctx.Invoke("aws-native:healthimaging:getDatastore", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDatastoreArgs struct {
	DatastoreId string `pulumi:"datastoreId"`
}

type LookupDatastoreResult struct {
	CreatedAt       *string          `pulumi:"createdAt"`
	DatastoreArn    *string          `pulumi:"datastoreArn"`
	DatastoreId     *string          `pulumi:"datastoreId"`
	DatastoreStatus *DatastoreStatus `pulumi:"datastoreStatus"`
	UpdatedAt       *string          `pulumi:"updatedAt"`
}

func LookupDatastoreOutput(ctx *pulumi.Context, args LookupDatastoreOutputArgs, opts ...pulumi.InvokeOption) LookupDatastoreResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDatastoreResult, error) {
			args := v.(LookupDatastoreArgs)
			r, err := LookupDatastore(ctx, &args, opts...)
			var s LookupDatastoreResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDatastoreResultOutput)
}

type LookupDatastoreOutputArgs struct {
	DatastoreId pulumi.StringInput `pulumi:"datastoreId"`
}

func (LookupDatastoreOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDatastoreArgs)(nil)).Elem()
}

type LookupDatastoreResultOutput struct{ *pulumi.OutputState }

func (LookupDatastoreResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDatastoreResult)(nil)).Elem()
}

func (o LookupDatastoreResultOutput) ToLookupDatastoreResultOutput() LookupDatastoreResultOutput {
	return o
}

func (o LookupDatastoreResultOutput) ToLookupDatastoreResultOutputWithContext(ctx context.Context) LookupDatastoreResultOutput {
	return o
}

func (o LookupDatastoreResultOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDatastoreResult) *string { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

func (o LookupDatastoreResultOutput) DatastoreArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDatastoreResult) *string { return v.DatastoreArn }).(pulumi.StringPtrOutput)
}

func (o LookupDatastoreResultOutput) DatastoreId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDatastoreResult) *string { return v.DatastoreId }).(pulumi.StringPtrOutput)
}

func (o LookupDatastoreResultOutput) DatastoreStatus() DatastoreStatusPtrOutput {
	return o.ApplyT(func(v LookupDatastoreResult) *DatastoreStatus { return v.DatastoreStatus }).(DatastoreStatusPtrOutput)
}

func (o LookupDatastoreResultOutput) UpdatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDatastoreResult) *string { return v.UpdatedAt }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDatastoreResultOutput{})
}
