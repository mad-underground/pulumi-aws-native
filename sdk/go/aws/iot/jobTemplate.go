// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iot

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Job templates enable you to preconfigure jobs so that you can deploy them to multiple sets of target devices.
type JobTemplate struct {
	pulumi.CustomResourceState

	// The criteria that determine when and how a job abort takes place.
	AbortConfig AbortConfigPropertiesPtrOutput `pulumi:"abortConfig"`
	Arn         pulumi.StringOutput            `pulumi:"arn"`
	// A description of the Job Template.
	Description                pulumi.StringOutput      `pulumi:"description"`
	DestinationPackageVersions pulumi.StringArrayOutput `pulumi:"destinationPackageVersions"`
	// The job document. Required if you don't specify a value for documentSource.
	Document pulumi.StringPtrOutput `pulumi:"document"`
	// An S3 link to the job document to use in the template. Required if you don't specify a value for document.
	DocumentSource pulumi.StringPtrOutput `pulumi:"documentSource"`
	// Optional for copying a JobTemplate from a pre-existing Job configuration.
	JobArn                   pulumi.StringPtrOutput                      `pulumi:"jobArn"`
	JobExecutionsRetryConfig JobExecutionsRetryConfigPropertiesPtrOutput `pulumi:"jobExecutionsRetryConfig"`
	// Allows you to create a staged rollout of a job.
	JobExecutionsRolloutConfig JobExecutionsRolloutConfigPropertiesPtrOutput `pulumi:"jobExecutionsRolloutConfig"`
	JobTemplateId              pulumi.StringOutput                           `pulumi:"jobTemplateId"`
	MaintenanceWindows         JobTemplateMaintenanceWindowArrayOutput       `pulumi:"maintenanceWindows"`
	// Configuration for pre-signed S3 URLs.
	PresignedUrlConfig PresignedUrlConfigPropertiesPtrOutput `pulumi:"presignedUrlConfig"`
	// Metadata that can be used to manage the JobTemplate.
	Tags JobTemplateTagArrayOutput `pulumi:"tags"`
	// Specifies the amount of time each device has to finish its execution of the job.
	TimeoutConfig TimeoutConfigPropertiesPtrOutput `pulumi:"timeoutConfig"`
}

// NewJobTemplate registers a new resource with the given unique name, arguments, and options.
func NewJobTemplate(ctx *pulumi.Context,
	name string, args *JobTemplateArgs, opts ...pulumi.ResourceOption) (*JobTemplate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Description == nil {
		return nil, errors.New("invalid value for required argument 'Description'")
	}
	if args.JobTemplateId == nil {
		return nil, errors.New("invalid value for required argument 'JobTemplateId'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"abortConfig",
		"description",
		"destinationPackageVersions[*]",
		"document",
		"documentSource",
		"jobArn",
		"jobExecutionsRetryConfig",
		"jobExecutionsRolloutConfig",
		"jobTemplateId",
		"maintenanceWindows[*]",
		"presignedUrlConfig",
		"tags[*]",
		"timeoutConfig",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource JobTemplate
	err := ctx.RegisterResource("aws-native:iot:JobTemplate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetJobTemplate gets an existing JobTemplate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetJobTemplate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *JobTemplateState, opts ...pulumi.ResourceOption) (*JobTemplate, error) {
	var resource JobTemplate
	err := ctx.ReadResource("aws-native:iot:JobTemplate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering JobTemplate resources.
type jobTemplateState struct {
}

type JobTemplateState struct {
}

func (JobTemplateState) ElementType() reflect.Type {
	return reflect.TypeOf((*jobTemplateState)(nil)).Elem()
}

type jobTemplateArgs struct {
	// The criteria that determine when and how a job abort takes place.
	AbortConfig *AbortConfigProperties `pulumi:"abortConfig"`
	// A description of the Job Template.
	Description                string   `pulumi:"description"`
	DestinationPackageVersions []string `pulumi:"destinationPackageVersions"`
	// The job document. Required if you don't specify a value for documentSource.
	Document *string `pulumi:"document"`
	// An S3 link to the job document to use in the template. Required if you don't specify a value for document.
	DocumentSource *string `pulumi:"documentSource"`
	// Optional for copying a JobTemplate from a pre-existing Job configuration.
	JobArn                   *string                             `pulumi:"jobArn"`
	JobExecutionsRetryConfig *JobExecutionsRetryConfigProperties `pulumi:"jobExecutionsRetryConfig"`
	// Allows you to create a staged rollout of a job.
	JobExecutionsRolloutConfig *JobExecutionsRolloutConfigProperties `pulumi:"jobExecutionsRolloutConfig"`
	JobTemplateId              string                                `pulumi:"jobTemplateId"`
	MaintenanceWindows         []JobTemplateMaintenanceWindow        `pulumi:"maintenanceWindows"`
	// Configuration for pre-signed S3 URLs.
	PresignedUrlConfig *PresignedUrlConfigProperties `pulumi:"presignedUrlConfig"`
	// Metadata that can be used to manage the JobTemplate.
	Tags []JobTemplateTag `pulumi:"tags"`
	// Specifies the amount of time each device has to finish its execution of the job.
	TimeoutConfig *TimeoutConfigProperties `pulumi:"timeoutConfig"`
}

// The set of arguments for constructing a JobTemplate resource.
type JobTemplateArgs struct {
	// The criteria that determine when and how a job abort takes place.
	AbortConfig AbortConfigPropertiesPtrInput
	// A description of the Job Template.
	Description                pulumi.StringInput
	DestinationPackageVersions pulumi.StringArrayInput
	// The job document. Required if you don't specify a value for documentSource.
	Document pulumi.StringPtrInput
	// An S3 link to the job document to use in the template. Required if you don't specify a value for document.
	DocumentSource pulumi.StringPtrInput
	// Optional for copying a JobTemplate from a pre-existing Job configuration.
	JobArn                   pulumi.StringPtrInput
	JobExecutionsRetryConfig JobExecutionsRetryConfigPropertiesPtrInput
	// Allows you to create a staged rollout of a job.
	JobExecutionsRolloutConfig JobExecutionsRolloutConfigPropertiesPtrInput
	JobTemplateId              pulumi.StringInput
	MaintenanceWindows         JobTemplateMaintenanceWindowArrayInput
	// Configuration for pre-signed S3 URLs.
	PresignedUrlConfig PresignedUrlConfigPropertiesPtrInput
	// Metadata that can be used to manage the JobTemplate.
	Tags JobTemplateTagArrayInput
	// Specifies the amount of time each device has to finish its execution of the job.
	TimeoutConfig TimeoutConfigPropertiesPtrInput
}

func (JobTemplateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*jobTemplateArgs)(nil)).Elem()
}

type JobTemplateInput interface {
	pulumi.Input

	ToJobTemplateOutput() JobTemplateOutput
	ToJobTemplateOutputWithContext(ctx context.Context) JobTemplateOutput
}

func (*JobTemplate) ElementType() reflect.Type {
	return reflect.TypeOf((**JobTemplate)(nil)).Elem()
}

func (i *JobTemplate) ToJobTemplateOutput() JobTemplateOutput {
	return i.ToJobTemplateOutputWithContext(context.Background())
}

func (i *JobTemplate) ToJobTemplateOutputWithContext(ctx context.Context) JobTemplateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobTemplateOutput)
}

func (i *JobTemplate) ToOutput(ctx context.Context) pulumix.Output[*JobTemplate] {
	return pulumix.Output[*JobTemplate]{
		OutputState: i.ToJobTemplateOutputWithContext(ctx).OutputState,
	}
}

type JobTemplateOutput struct{ *pulumi.OutputState }

func (JobTemplateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**JobTemplate)(nil)).Elem()
}

func (o JobTemplateOutput) ToJobTemplateOutput() JobTemplateOutput {
	return o
}

func (o JobTemplateOutput) ToJobTemplateOutputWithContext(ctx context.Context) JobTemplateOutput {
	return o
}

func (o JobTemplateOutput) ToOutput(ctx context.Context) pulumix.Output[*JobTemplate] {
	return pulumix.Output[*JobTemplate]{
		OutputState: o.OutputState,
	}
}

// The criteria that determine when and how a job abort takes place.
func (o JobTemplateOutput) AbortConfig() AbortConfigPropertiesPtrOutput {
	return o.ApplyT(func(v *JobTemplate) AbortConfigPropertiesPtrOutput { return v.AbortConfig }).(AbortConfigPropertiesPtrOutput)
}

func (o JobTemplateOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *JobTemplate) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// A description of the Job Template.
func (o JobTemplateOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *JobTemplate) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

func (o JobTemplateOutput) DestinationPackageVersions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *JobTemplate) pulumi.StringArrayOutput { return v.DestinationPackageVersions }).(pulumi.StringArrayOutput)
}

// The job document. Required if you don't specify a value for documentSource.
func (o JobTemplateOutput) Document() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobTemplate) pulumi.StringPtrOutput { return v.Document }).(pulumi.StringPtrOutput)
}

// An S3 link to the job document to use in the template. Required if you don't specify a value for document.
func (o JobTemplateOutput) DocumentSource() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobTemplate) pulumi.StringPtrOutput { return v.DocumentSource }).(pulumi.StringPtrOutput)
}

// Optional for copying a JobTemplate from a pre-existing Job configuration.
func (o JobTemplateOutput) JobArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobTemplate) pulumi.StringPtrOutput { return v.JobArn }).(pulumi.StringPtrOutput)
}

func (o JobTemplateOutput) JobExecutionsRetryConfig() JobExecutionsRetryConfigPropertiesPtrOutput {
	return o.ApplyT(func(v *JobTemplate) JobExecutionsRetryConfigPropertiesPtrOutput { return v.JobExecutionsRetryConfig }).(JobExecutionsRetryConfigPropertiesPtrOutput)
}

// Allows you to create a staged rollout of a job.
func (o JobTemplateOutput) JobExecutionsRolloutConfig() JobExecutionsRolloutConfigPropertiesPtrOutput {
	return o.ApplyT(func(v *JobTemplate) JobExecutionsRolloutConfigPropertiesPtrOutput {
		return v.JobExecutionsRolloutConfig
	}).(JobExecutionsRolloutConfigPropertiesPtrOutput)
}

func (o JobTemplateOutput) JobTemplateId() pulumi.StringOutput {
	return o.ApplyT(func(v *JobTemplate) pulumi.StringOutput { return v.JobTemplateId }).(pulumi.StringOutput)
}

func (o JobTemplateOutput) MaintenanceWindows() JobTemplateMaintenanceWindowArrayOutput {
	return o.ApplyT(func(v *JobTemplate) JobTemplateMaintenanceWindowArrayOutput { return v.MaintenanceWindows }).(JobTemplateMaintenanceWindowArrayOutput)
}

// Configuration for pre-signed S3 URLs.
func (o JobTemplateOutput) PresignedUrlConfig() PresignedUrlConfigPropertiesPtrOutput {
	return o.ApplyT(func(v *JobTemplate) PresignedUrlConfigPropertiesPtrOutput { return v.PresignedUrlConfig }).(PresignedUrlConfigPropertiesPtrOutput)
}

// Metadata that can be used to manage the JobTemplate.
func (o JobTemplateOutput) Tags() JobTemplateTagArrayOutput {
	return o.ApplyT(func(v *JobTemplate) JobTemplateTagArrayOutput { return v.Tags }).(JobTemplateTagArrayOutput)
}

// Specifies the amount of time each device has to finish its execution of the job.
func (o JobTemplateOutput) TimeoutConfig() TimeoutConfigPropertiesPtrOutput {
	return o.ApplyT(func(v *JobTemplate) TimeoutConfigPropertiesPtrOutput { return v.TimeoutConfig }).(TimeoutConfigPropertiesPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*JobTemplateInput)(nil)).Elem(), &JobTemplate{})
	pulumi.RegisterOutputType(JobTemplateOutput{})
}
