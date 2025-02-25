// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package kinesisvideo

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type Definition for AWS::KinesisVideo::Stream
type Stream struct {
	pulumi.CustomResourceState

	// The Amazon Resource Name (ARN) of the Kinesis Video stream.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The number of hours till which Kinesis Video will retain the data in the stream
	DataRetentionInHours pulumi.IntPtrOutput `pulumi:"dataRetentionInHours"`
	// The name of the device that is writing to the stream.
	DeviceName pulumi.StringPtrOutput `pulumi:"deviceName"`
	// AWS KMS key ID that Kinesis Video Streams uses to encrypt stream data.
	KmsKeyId pulumi.StringPtrOutput `pulumi:"kmsKeyId"`
	// The media type of the stream. Consumers of the stream can use this information when processing the stream.
	MediaType pulumi.StringPtrOutput `pulumi:"mediaType"`
	// The name of the Kinesis Video stream.
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// An array of key-value pairs associated with the Kinesis Video Stream.
	Tags aws.TagArrayOutput `pulumi:"tags"`
}

// NewStream registers a new resource with the given unique name, arguments, and options.
func NewStream(ctx *pulumi.Context,
	name string, args *StreamArgs, opts ...pulumi.ResourceOption) (*Stream, error) {
	if args == nil {
		args = &StreamArgs{}
	}

	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"name",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Stream
	err := ctx.RegisterResource("aws-native:kinesisvideo:Stream", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStream gets an existing Stream resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStream(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StreamState, opts ...pulumi.ResourceOption) (*Stream, error) {
	var resource Stream
	err := ctx.ReadResource("aws-native:kinesisvideo:Stream", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Stream resources.
type streamState struct {
}

type StreamState struct {
}

func (StreamState) ElementType() reflect.Type {
	return reflect.TypeOf((*streamState)(nil)).Elem()
}

type streamArgs struct {
	// The number of hours till which Kinesis Video will retain the data in the stream
	DataRetentionInHours *int `pulumi:"dataRetentionInHours"`
	// The name of the device that is writing to the stream.
	DeviceName *string `pulumi:"deviceName"`
	// AWS KMS key ID that Kinesis Video Streams uses to encrypt stream data.
	KmsKeyId *string `pulumi:"kmsKeyId"`
	// The media type of the stream. Consumers of the stream can use this information when processing the stream.
	MediaType *string `pulumi:"mediaType"`
	// The name of the Kinesis Video stream.
	Name *string `pulumi:"name"`
	// An array of key-value pairs associated with the Kinesis Video Stream.
	Tags []aws.Tag `pulumi:"tags"`
}

// The set of arguments for constructing a Stream resource.
type StreamArgs struct {
	// The number of hours till which Kinesis Video will retain the data in the stream
	DataRetentionInHours pulumi.IntPtrInput
	// The name of the device that is writing to the stream.
	DeviceName pulumi.StringPtrInput
	// AWS KMS key ID that Kinesis Video Streams uses to encrypt stream data.
	KmsKeyId pulumi.StringPtrInput
	// The media type of the stream. Consumers of the stream can use this information when processing the stream.
	MediaType pulumi.StringPtrInput
	// The name of the Kinesis Video stream.
	Name pulumi.StringPtrInput
	// An array of key-value pairs associated with the Kinesis Video Stream.
	Tags aws.TagArrayInput
}

func (StreamArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*streamArgs)(nil)).Elem()
}

type StreamInput interface {
	pulumi.Input

	ToStreamOutput() StreamOutput
	ToStreamOutputWithContext(ctx context.Context) StreamOutput
}

func (*Stream) ElementType() reflect.Type {
	return reflect.TypeOf((**Stream)(nil)).Elem()
}

func (i *Stream) ToStreamOutput() StreamOutput {
	return i.ToStreamOutputWithContext(context.Background())
}

func (i *Stream) ToStreamOutputWithContext(ctx context.Context) StreamOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StreamOutput)
}

type StreamOutput struct{ *pulumi.OutputState }

func (StreamOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Stream)(nil)).Elem()
}

func (o StreamOutput) ToStreamOutput() StreamOutput {
	return o
}

func (o StreamOutput) ToStreamOutputWithContext(ctx context.Context) StreamOutput {
	return o
}

// The Amazon Resource Name (ARN) of the Kinesis Video stream.
func (o StreamOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *Stream) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// The number of hours till which Kinesis Video will retain the data in the stream
func (o StreamOutput) DataRetentionInHours() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Stream) pulumi.IntPtrOutput { return v.DataRetentionInHours }).(pulumi.IntPtrOutput)
}

// The name of the device that is writing to the stream.
func (o StreamOutput) DeviceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Stream) pulumi.StringPtrOutput { return v.DeviceName }).(pulumi.StringPtrOutput)
}

// AWS KMS key ID that Kinesis Video Streams uses to encrypt stream data.
func (o StreamOutput) KmsKeyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Stream) pulumi.StringPtrOutput { return v.KmsKeyId }).(pulumi.StringPtrOutput)
}

// The media type of the stream. Consumers of the stream can use this information when processing the stream.
func (o StreamOutput) MediaType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Stream) pulumi.StringPtrOutput { return v.MediaType }).(pulumi.StringPtrOutput)
}

// The name of the Kinesis Video stream.
func (o StreamOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Stream) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

// An array of key-value pairs associated with the Kinesis Video Stream.
func (o StreamOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v *Stream) aws.TagArrayOutput { return v.Tags }).(aws.TagArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*StreamInput)(nil)).Elem(), &Stream{})
	pulumi.RegisterOutputType(StreamOutput{})
}
