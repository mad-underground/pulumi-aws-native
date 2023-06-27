// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cleanrooms

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Represents a collaboration between AWS accounts that allows for secure data collaboration
func LookupCollaboration(ctx *pulumi.Context, args *LookupCollaborationArgs, opts ...pulumi.InvokeOption) (*LookupCollaborationResult, error) {
	var rv LookupCollaborationResult
	err := ctx.Invoke("aws-native:cleanrooms:getCollaboration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCollaborationArgs struct {
	CollaborationIdentifier string `pulumi:"collaborationIdentifier"`
}

type LookupCollaborationResult struct {
	Arn                     *string `pulumi:"arn"`
	CollaborationIdentifier *string `pulumi:"collaborationIdentifier"`
	Description             *string `pulumi:"description"`
	Name                    *string `pulumi:"name"`
	// An arbitrary set of tags (key-value pairs) for this cleanrooms collaboration.
	Tags []CollaborationTag `pulumi:"tags"`
}

func LookupCollaborationOutput(ctx *pulumi.Context, args LookupCollaborationOutputArgs, opts ...pulumi.InvokeOption) LookupCollaborationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCollaborationResult, error) {
			args := v.(LookupCollaborationArgs)
			r, err := LookupCollaboration(ctx, &args, opts...)
			var s LookupCollaborationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupCollaborationResultOutput)
}

type LookupCollaborationOutputArgs struct {
	CollaborationIdentifier pulumi.StringInput `pulumi:"collaborationIdentifier"`
}

func (LookupCollaborationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCollaborationArgs)(nil)).Elem()
}

type LookupCollaborationResultOutput struct{ *pulumi.OutputState }

func (LookupCollaborationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCollaborationResult)(nil)).Elem()
}

func (o LookupCollaborationResultOutput) ToLookupCollaborationResultOutput() LookupCollaborationResultOutput {
	return o
}

func (o LookupCollaborationResultOutput) ToLookupCollaborationResultOutputWithContext(ctx context.Context) LookupCollaborationResultOutput {
	return o
}

func (o LookupCollaborationResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCollaborationResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

func (o LookupCollaborationResultOutput) CollaborationIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCollaborationResult) *string { return v.CollaborationIdentifier }).(pulumi.StringPtrOutput)
}

func (o LookupCollaborationResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCollaborationResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupCollaborationResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCollaborationResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// An arbitrary set of tags (key-value pairs) for this cleanrooms collaboration.
func (o LookupCollaborationResultOutput) Tags() CollaborationTagArrayOutput {
	return o.ApplyT(func(v LookupCollaborationResult) []CollaborationTag { return v.Tags }).(CollaborationTagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCollaborationResultOutput{})
}