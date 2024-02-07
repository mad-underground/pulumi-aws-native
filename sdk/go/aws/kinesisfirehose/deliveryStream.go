// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package kinesisfirehose

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::KinesisFirehose::DeliveryStream
type DeliveryStream struct {
	pulumi.CustomResourceState

	AmazonOpenSearchServerlessDestinationConfiguration DeliveryStreamAmazonOpenSearchServerlessDestinationConfigurationPtrOutput `pulumi:"amazonOpenSearchServerlessDestinationConfiguration"`
	AmazonopensearchserviceDestinationConfiguration    DeliveryStreamAmazonopensearchserviceDestinationConfigurationPtrOutput    `pulumi:"amazonopensearchserviceDestinationConfiguration"`
	Arn                                                pulumi.StringOutput                                                       `pulumi:"arn"`
	DeliveryStreamEncryptionConfigurationInput         DeliveryStreamEncryptionConfigurationInputTypePtrOutput                   `pulumi:"deliveryStreamEncryptionConfigurationInput"`
	DeliveryStreamName                                 pulumi.StringPtrOutput                                                    `pulumi:"deliveryStreamName"`
	DeliveryStreamType                                 DeliveryStreamTypePtrOutput                                               `pulumi:"deliveryStreamType"`
	ElasticsearchDestinationConfiguration              DeliveryStreamElasticsearchDestinationConfigurationPtrOutput              `pulumi:"elasticsearchDestinationConfiguration"`
	ExtendedS3DestinationConfiguration                 DeliveryStreamExtendedS3DestinationConfigurationPtrOutput                 `pulumi:"extendedS3DestinationConfiguration"`
	HttpEndpointDestinationConfiguration               DeliveryStreamHttpEndpointDestinationConfigurationPtrOutput               `pulumi:"httpEndpointDestinationConfiguration"`
	KinesisStreamSourceConfiguration                   DeliveryStreamKinesisStreamSourceConfigurationPtrOutput                   `pulumi:"kinesisStreamSourceConfiguration"`
	MskSourceConfiguration                             DeliveryStreamMskSourceConfigurationPtrOutput                             `pulumi:"mskSourceConfiguration"`
	RedshiftDestinationConfiguration                   DeliveryStreamRedshiftDestinationConfigurationPtrOutput                   `pulumi:"redshiftDestinationConfiguration"`
	S3DestinationConfiguration                         DeliveryStreamS3DestinationConfigurationPtrOutput                         `pulumi:"s3DestinationConfiguration"`
	SnowflakeDestinationConfiguration                  DeliveryStreamSnowflakeDestinationConfigurationPtrOutput                  `pulumi:"snowflakeDestinationConfiguration"`
	SplunkDestinationConfiguration                     DeliveryStreamSplunkDestinationConfigurationPtrOutput                     `pulumi:"splunkDestinationConfiguration"`
	Tags                                               DeliveryStreamTagArrayOutput                                              `pulumi:"tags"`
}

// NewDeliveryStream registers a new resource with the given unique name, arguments, and options.
func NewDeliveryStream(ctx *pulumi.Context,
	name string, args *DeliveryStreamArgs, opts ...pulumi.ResourceOption) (*DeliveryStream, error) {
	if args == nil {
		args = &DeliveryStreamArgs{}
	}

	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"amazonOpenSearchServerlessDestinationConfiguration.vpcConfiguration",
		"amazonopensearchserviceDestinationConfiguration.vpcConfiguration",
		"deliveryStreamName",
		"deliveryStreamType",
		"elasticsearchDestinationConfiguration.vpcConfiguration",
		"kinesisStreamSourceConfiguration",
		"mskSourceConfiguration",
		"snowflakeDestinationConfiguration.snowflakeVpcConfiguration",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DeliveryStream
	err := ctx.RegisterResource("aws-native:kinesisfirehose:DeliveryStream", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDeliveryStream gets an existing DeliveryStream resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDeliveryStream(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DeliveryStreamState, opts ...pulumi.ResourceOption) (*DeliveryStream, error) {
	var resource DeliveryStream
	err := ctx.ReadResource("aws-native:kinesisfirehose:DeliveryStream", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DeliveryStream resources.
type deliveryStreamState struct {
}

type DeliveryStreamState struct {
}

func (DeliveryStreamState) ElementType() reflect.Type {
	return reflect.TypeOf((*deliveryStreamState)(nil)).Elem()
}

type deliveryStreamArgs struct {
	AmazonOpenSearchServerlessDestinationConfiguration *DeliveryStreamAmazonOpenSearchServerlessDestinationConfiguration `pulumi:"amazonOpenSearchServerlessDestinationConfiguration"`
	AmazonopensearchserviceDestinationConfiguration    *DeliveryStreamAmazonopensearchserviceDestinationConfiguration    `pulumi:"amazonopensearchserviceDestinationConfiguration"`
	DeliveryStreamEncryptionConfigurationInput         *DeliveryStreamEncryptionConfigurationInputType                   `pulumi:"deliveryStreamEncryptionConfigurationInput"`
	DeliveryStreamName                                 *string                                                           `pulumi:"deliveryStreamName"`
	DeliveryStreamType                                 *DeliveryStreamType                                               `pulumi:"deliveryStreamType"`
	ElasticsearchDestinationConfiguration              *DeliveryStreamElasticsearchDestinationConfiguration              `pulumi:"elasticsearchDestinationConfiguration"`
	ExtendedS3DestinationConfiguration                 *DeliveryStreamExtendedS3DestinationConfiguration                 `pulumi:"extendedS3DestinationConfiguration"`
	HttpEndpointDestinationConfiguration               *DeliveryStreamHttpEndpointDestinationConfiguration               `pulumi:"httpEndpointDestinationConfiguration"`
	KinesisStreamSourceConfiguration                   *DeliveryStreamKinesisStreamSourceConfiguration                   `pulumi:"kinesisStreamSourceConfiguration"`
	MskSourceConfiguration                             *DeliveryStreamMskSourceConfiguration                             `pulumi:"mskSourceConfiguration"`
	RedshiftDestinationConfiguration                   *DeliveryStreamRedshiftDestinationConfiguration                   `pulumi:"redshiftDestinationConfiguration"`
	S3DestinationConfiguration                         *DeliveryStreamS3DestinationConfiguration                         `pulumi:"s3DestinationConfiguration"`
	SnowflakeDestinationConfiguration                  *DeliveryStreamSnowflakeDestinationConfiguration                  `pulumi:"snowflakeDestinationConfiguration"`
	SplunkDestinationConfiguration                     *DeliveryStreamSplunkDestinationConfiguration                     `pulumi:"splunkDestinationConfiguration"`
	Tags                                               []DeliveryStreamTag                                               `pulumi:"tags"`
}

// The set of arguments for constructing a DeliveryStream resource.
type DeliveryStreamArgs struct {
	AmazonOpenSearchServerlessDestinationConfiguration DeliveryStreamAmazonOpenSearchServerlessDestinationConfigurationPtrInput
	AmazonopensearchserviceDestinationConfiguration    DeliveryStreamAmazonopensearchserviceDestinationConfigurationPtrInput
	DeliveryStreamEncryptionConfigurationInput         DeliveryStreamEncryptionConfigurationInputTypePtrInput
	DeliveryStreamName                                 pulumi.StringPtrInput
	DeliveryStreamType                                 DeliveryStreamTypePtrInput
	ElasticsearchDestinationConfiguration              DeliveryStreamElasticsearchDestinationConfigurationPtrInput
	ExtendedS3DestinationConfiguration                 DeliveryStreamExtendedS3DestinationConfigurationPtrInput
	HttpEndpointDestinationConfiguration               DeliveryStreamHttpEndpointDestinationConfigurationPtrInput
	KinesisStreamSourceConfiguration                   DeliveryStreamKinesisStreamSourceConfigurationPtrInput
	MskSourceConfiguration                             DeliveryStreamMskSourceConfigurationPtrInput
	RedshiftDestinationConfiguration                   DeliveryStreamRedshiftDestinationConfigurationPtrInput
	S3DestinationConfiguration                         DeliveryStreamS3DestinationConfigurationPtrInput
	SnowflakeDestinationConfiguration                  DeliveryStreamSnowflakeDestinationConfigurationPtrInput
	SplunkDestinationConfiguration                     DeliveryStreamSplunkDestinationConfigurationPtrInput
	Tags                                               DeliveryStreamTagArrayInput
}

func (DeliveryStreamArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*deliveryStreamArgs)(nil)).Elem()
}

type DeliveryStreamInput interface {
	pulumi.Input

	ToDeliveryStreamOutput() DeliveryStreamOutput
	ToDeliveryStreamOutputWithContext(ctx context.Context) DeliveryStreamOutput
}

func (*DeliveryStream) ElementType() reflect.Type {
	return reflect.TypeOf((**DeliveryStream)(nil)).Elem()
}

func (i *DeliveryStream) ToDeliveryStreamOutput() DeliveryStreamOutput {
	return i.ToDeliveryStreamOutputWithContext(context.Background())
}

func (i *DeliveryStream) ToDeliveryStreamOutputWithContext(ctx context.Context) DeliveryStreamOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeliveryStreamOutput)
}

type DeliveryStreamOutput struct{ *pulumi.OutputState }

func (DeliveryStreamOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DeliveryStream)(nil)).Elem()
}

func (o DeliveryStreamOutput) ToDeliveryStreamOutput() DeliveryStreamOutput {
	return o
}

func (o DeliveryStreamOutput) ToDeliveryStreamOutputWithContext(ctx context.Context) DeliveryStreamOutput {
	return o
}

func (o DeliveryStreamOutput) AmazonOpenSearchServerlessDestinationConfiguration() DeliveryStreamAmazonOpenSearchServerlessDestinationConfigurationPtrOutput {
	return o.ApplyT(func(v *DeliveryStream) DeliveryStreamAmazonOpenSearchServerlessDestinationConfigurationPtrOutput {
		return v.AmazonOpenSearchServerlessDestinationConfiguration
	}).(DeliveryStreamAmazonOpenSearchServerlessDestinationConfigurationPtrOutput)
}

func (o DeliveryStreamOutput) AmazonopensearchserviceDestinationConfiguration() DeliveryStreamAmazonopensearchserviceDestinationConfigurationPtrOutput {
	return o.ApplyT(func(v *DeliveryStream) DeliveryStreamAmazonopensearchserviceDestinationConfigurationPtrOutput {
		return v.AmazonopensearchserviceDestinationConfiguration
	}).(DeliveryStreamAmazonopensearchserviceDestinationConfigurationPtrOutput)
}

func (o DeliveryStreamOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *DeliveryStream) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

func (o DeliveryStreamOutput) DeliveryStreamEncryptionConfigurationInput() DeliveryStreamEncryptionConfigurationInputTypePtrOutput {
	return o.ApplyT(func(v *DeliveryStream) DeliveryStreamEncryptionConfigurationInputTypePtrOutput {
		return v.DeliveryStreamEncryptionConfigurationInput
	}).(DeliveryStreamEncryptionConfigurationInputTypePtrOutput)
}

func (o DeliveryStreamOutput) DeliveryStreamName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DeliveryStream) pulumi.StringPtrOutput { return v.DeliveryStreamName }).(pulumi.StringPtrOutput)
}

func (o DeliveryStreamOutput) DeliveryStreamType() DeliveryStreamTypePtrOutput {
	return o.ApplyT(func(v *DeliveryStream) DeliveryStreamTypePtrOutput { return v.DeliveryStreamType }).(DeliveryStreamTypePtrOutput)
}

func (o DeliveryStreamOutput) ElasticsearchDestinationConfiguration() DeliveryStreamElasticsearchDestinationConfigurationPtrOutput {
	return o.ApplyT(func(v *DeliveryStream) DeliveryStreamElasticsearchDestinationConfigurationPtrOutput {
		return v.ElasticsearchDestinationConfiguration
	}).(DeliveryStreamElasticsearchDestinationConfigurationPtrOutput)
}

func (o DeliveryStreamOutput) ExtendedS3DestinationConfiguration() DeliveryStreamExtendedS3DestinationConfigurationPtrOutput {
	return o.ApplyT(func(v *DeliveryStream) DeliveryStreamExtendedS3DestinationConfigurationPtrOutput {
		return v.ExtendedS3DestinationConfiguration
	}).(DeliveryStreamExtendedS3DestinationConfigurationPtrOutput)
}

func (o DeliveryStreamOutput) HttpEndpointDestinationConfiguration() DeliveryStreamHttpEndpointDestinationConfigurationPtrOutput {
	return o.ApplyT(func(v *DeliveryStream) DeliveryStreamHttpEndpointDestinationConfigurationPtrOutput {
		return v.HttpEndpointDestinationConfiguration
	}).(DeliveryStreamHttpEndpointDestinationConfigurationPtrOutput)
}

func (o DeliveryStreamOutput) KinesisStreamSourceConfiguration() DeliveryStreamKinesisStreamSourceConfigurationPtrOutput {
	return o.ApplyT(func(v *DeliveryStream) DeliveryStreamKinesisStreamSourceConfigurationPtrOutput {
		return v.KinesisStreamSourceConfiguration
	}).(DeliveryStreamKinesisStreamSourceConfigurationPtrOutput)
}

func (o DeliveryStreamOutput) MskSourceConfiguration() DeliveryStreamMskSourceConfigurationPtrOutput {
	return o.ApplyT(func(v *DeliveryStream) DeliveryStreamMskSourceConfigurationPtrOutput { return v.MskSourceConfiguration }).(DeliveryStreamMskSourceConfigurationPtrOutput)
}

func (o DeliveryStreamOutput) RedshiftDestinationConfiguration() DeliveryStreamRedshiftDestinationConfigurationPtrOutput {
	return o.ApplyT(func(v *DeliveryStream) DeliveryStreamRedshiftDestinationConfigurationPtrOutput {
		return v.RedshiftDestinationConfiguration
	}).(DeliveryStreamRedshiftDestinationConfigurationPtrOutput)
}

func (o DeliveryStreamOutput) S3DestinationConfiguration() DeliveryStreamS3DestinationConfigurationPtrOutput {
	return o.ApplyT(func(v *DeliveryStream) DeliveryStreamS3DestinationConfigurationPtrOutput {
		return v.S3DestinationConfiguration
	}).(DeliveryStreamS3DestinationConfigurationPtrOutput)
}

func (o DeliveryStreamOutput) SnowflakeDestinationConfiguration() DeliveryStreamSnowflakeDestinationConfigurationPtrOutput {
	return o.ApplyT(func(v *DeliveryStream) DeliveryStreamSnowflakeDestinationConfigurationPtrOutput {
		return v.SnowflakeDestinationConfiguration
	}).(DeliveryStreamSnowflakeDestinationConfigurationPtrOutput)
}

func (o DeliveryStreamOutput) SplunkDestinationConfiguration() DeliveryStreamSplunkDestinationConfigurationPtrOutput {
	return o.ApplyT(func(v *DeliveryStream) DeliveryStreamSplunkDestinationConfigurationPtrOutput {
		return v.SplunkDestinationConfiguration
	}).(DeliveryStreamSplunkDestinationConfigurationPtrOutput)
}

func (o DeliveryStreamOutput) Tags() DeliveryStreamTagArrayOutput {
	return o.ApplyT(func(v *DeliveryStream) DeliveryStreamTagArrayOutput { return v.Tags }).(DeliveryStreamTagArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DeliveryStreamInput)(nil)).Elem(), &DeliveryStream{})
	pulumi.RegisterOutputType(DeliveryStreamOutput{})
}
