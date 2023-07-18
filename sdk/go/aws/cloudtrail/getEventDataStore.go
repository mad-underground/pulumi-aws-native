// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudtrail

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A storage lake of event data against which you can run complex SQL-based queries. An event data store can include events that you have logged on your account from the last 90 to 2555 days (about three months to up to seven years).
func LookupEventDataStore(ctx *pulumi.Context, args *LookupEventDataStoreArgs, opts ...pulumi.InvokeOption) (*LookupEventDataStoreResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupEventDataStoreResult
	err := ctx.Invoke("aws-native:cloudtrail:getEventDataStore", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupEventDataStoreArgs struct {
	// The ARN of the event data store.
	EventDataStoreArn string `pulumi:"eventDataStoreArn"`
}

type LookupEventDataStoreResult struct {
	// The advanced event selectors that were used to select events for the data store.
	AdvancedEventSelectors []EventDataStoreAdvancedEventSelector `pulumi:"advancedEventSelectors"`
	// The timestamp of the event data store's creation.
	CreatedTimestamp *string `pulumi:"createdTimestamp"`
	// The ARN of the event data store.
	EventDataStoreArn *string `pulumi:"eventDataStoreArn"`
	// Indicates whether the event data store is ingesting events.
	IngestionEnabled *bool `pulumi:"ingestionEnabled"`
	// Specifies the KMS key ID to use to encrypt the events delivered by CloudTrail. The value can be an alias name prefixed by 'alias/', a fully specified ARN to an alias, a fully specified ARN to a key, or a globally unique identifier.
	KmsKeyId *string `pulumi:"kmsKeyId"`
	// Indicates whether the event data store includes events from all regions, or only from the region in which it was created.
	MultiRegionEnabled *bool `pulumi:"multiRegionEnabled"`
	// The name of the event data store.
	Name *string `pulumi:"name"`
	// Indicates that an event data store is collecting logged events for an organization.
	OrganizationEnabled *bool `pulumi:"organizationEnabled"`
	// The retention period, in days.
	RetentionPeriod *int `pulumi:"retentionPeriod"`
	// The status of an event data store. Values are STARTING_INGESTION, ENABLED, STOPPING_INGESTION, STOPPED_INGESTION and PENDING_DELETION.
	Status *string             `pulumi:"status"`
	Tags   []EventDataStoreTag `pulumi:"tags"`
	// Indicates whether the event data store is protected from termination.
	TerminationProtectionEnabled *bool `pulumi:"terminationProtectionEnabled"`
	// The timestamp showing when an event data store was updated, if applicable. UpdatedTimestamp is always either the same or newer than the time shown in CreatedTimestamp.
	UpdatedTimestamp *string `pulumi:"updatedTimestamp"`
}

func LookupEventDataStoreOutput(ctx *pulumi.Context, args LookupEventDataStoreOutputArgs, opts ...pulumi.InvokeOption) LookupEventDataStoreResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupEventDataStoreResult, error) {
			args := v.(LookupEventDataStoreArgs)
			r, err := LookupEventDataStore(ctx, &args, opts...)
			var s LookupEventDataStoreResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupEventDataStoreResultOutput)
}

type LookupEventDataStoreOutputArgs struct {
	// The ARN of the event data store.
	EventDataStoreArn pulumi.StringInput `pulumi:"eventDataStoreArn"`
}

func (LookupEventDataStoreOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEventDataStoreArgs)(nil)).Elem()
}

type LookupEventDataStoreResultOutput struct{ *pulumi.OutputState }

func (LookupEventDataStoreResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEventDataStoreResult)(nil)).Elem()
}

func (o LookupEventDataStoreResultOutput) ToLookupEventDataStoreResultOutput() LookupEventDataStoreResultOutput {
	return o
}

func (o LookupEventDataStoreResultOutput) ToLookupEventDataStoreResultOutputWithContext(ctx context.Context) LookupEventDataStoreResultOutput {
	return o
}

// The advanced event selectors that were used to select events for the data store.
func (o LookupEventDataStoreResultOutput) AdvancedEventSelectors() EventDataStoreAdvancedEventSelectorArrayOutput {
	return o.ApplyT(func(v LookupEventDataStoreResult) []EventDataStoreAdvancedEventSelector {
		return v.AdvancedEventSelectors
	}).(EventDataStoreAdvancedEventSelectorArrayOutput)
}

// The timestamp of the event data store's creation.
func (o LookupEventDataStoreResultOutput) CreatedTimestamp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEventDataStoreResult) *string { return v.CreatedTimestamp }).(pulumi.StringPtrOutput)
}

// The ARN of the event data store.
func (o LookupEventDataStoreResultOutput) EventDataStoreArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEventDataStoreResult) *string { return v.EventDataStoreArn }).(pulumi.StringPtrOutput)
}

// Indicates whether the event data store is ingesting events.
func (o LookupEventDataStoreResultOutput) IngestionEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupEventDataStoreResult) *bool { return v.IngestionEnabled }).(pulumi.BoolPtrOutput)
}

// Specifies the KMS key ID to use to encrypt the events delivered by CloudTrail. The value can be an alias name prefixed by 'alias/', a fully specified ARN to an alias, a fully specified ARN to a key, or a globally unique identifier.
func (o LookupEventDataStoreResultOutput) KmsKeyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEventDataStoreResult) *string { return v.KmsKeyId }).(pulumi.StringPtrOutput)
}

// Indicates whether the event data store includes events from all regions, or only from the region in which it was created.
func (o LookupEventDataStoreResultOutput) MultiRegionEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupEventDataStoreResult) *bool { return v.MultiRegionEnabled }).(pulumi.BoolPtrOutput)
}

// The name of the event data store.
func (o LookupEventDataStoreResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEventDataStoreResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// Indicates that an event data store is collecting logged events for an organization.
func (o LookupEventDataStoreResultOutput) OrganizationEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupEventDataStoreResult) *bool { return v.OrganizationEnabled }).(pulumi.BoolPtrOutput)
}

// The retention period, in days.
func (o LookupEventDataStoreResultOutput) RetentionPeriod() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupEventDataStoreResult) *int { return v.RetentionPeriod }).(pulumi.IntPtrOutput)
}

// The status of an event data store. Values are STARTING_INGESTION, ENABLED, STOPPING_INGESTION, STOPPED_INGESTION and PENDING_DELETION.
func (o LookupEventDataStoreResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEventDataStoreResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func (o LookupEventDataStoreResultOutput) Tags() EventDataStoreTagArrayOutput {
	return o.ApplyT(func(v LookupEventDataStoreResult) []EventDataStoreTag { return v.Tags }).(EventDataStoreTagArrayOutput)
}

// Indicates whether the event data store is protected from termination.
func (o LookupEventDataStoreResultOutput) TerminationProtectionEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupEventDataStoreResult) *bool { return v.TerminationProtectionEnabled }).(pulumi.BoolPtrOutput)
}

// The timestamp showing when an event data store was updated, if applicable. UpdatedTimestamp is always either the same or newer than the time shown in CreatedTimestamp.
func (o LookupEventDataStoreResultOutput) UpdatedTimestamp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEventDataStoreResult) *string { return v.UpdatedTimestamp }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupEventDataStoreResultOutput{})
}
