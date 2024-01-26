// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package kinesisfirehose

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Resource Type definition for AWS::KinesisFirehose::DeliveryStream
func LookupDeliveryStream(ctx *pulumi.Context, args *LookupDeliveryStreamArgs, opts ...pulumi.InvokeOption) (*LookupDeliveryStreamResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupDeliveryStreamResult
	err := ctx.Invoke("aws-native:kinesisfirehose:getDeliveryStream", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDeliveryStreamArgs struct {
	DeliveryStreamName string `pulumi:"deliveryStreamName"`
}

type LookupDeliveryStreamResult struct {
	AmazonOpenSearchServerlessDestinationConfiguration *DeliveryStreamAmazonOpenSearchServerlessDestinationConfiguration `pulumi:"amazonOpenSearchServerlessDestinationConfiguration"`
	AmazonopensearchserviceDestinationConfiguration    *DeliveryStreamAmazonopensearchserviceDestinationConfiguration    `pulumi:"amazonopensearchserviceDestinationConfiguration"`
	Arn                                                *string                                                           `pulumi:"arn"`
	DeliveryStreamEncryptionConfigurationInput         *DeliveryStreamEncryptionConfigurationInputType                   `pulumi:"deliveryStreamEncryptionConfigurationInput"`
	ElasticsearchDestinationConfiguration              *DeliveryStreamElasticsearchDestinationConfiguration              `pulumi:"elasticsearchDestinationConfiguration"`
	ExtendedS3DestinationConfiguration                 *DeliveryStreamExtendedS3DestinationConfiguration                 `pulumi:"extendedS3DestinationConfiguration"`
	HttpEndpointDestinationConfiguration               *DeliveryStreamHttpEndpointDestinationConfiguration               `pulumi:"httpEndpointDestinationConfiguration"`
	RedshiftDestinationConfiguration                   *DeliveryStreamRedshiftDestinationConfiguration                   `pulumi:"redshiftDestinationConfiguration"`
	S3DestinationConfiguration                         *DeliveryStreamS3DestinationConfiguration                         `pulumi:"s3DestinationConfiguration"`
	SnowflakeDestinationConfiguration                  *DeliveryStreamSnowflakeDestinationConfiguration                  `pulumi:"snowflakeDestinationConfiguration"`
	SplunkDestinationConfiguration                     *DeliveryStreamSplunkDestinationConfiguration                     `pulumi:"splunkDestinationConfiguration"`
	Tags                                               []DeliveryStreamTag                                               `pulumi:"tags"`
}

func LookupDeliveryStreamOutput(ctx *pulumi.Context, args LookupDeliveryStreamOutputArgs, opts ...pulumi.InvokeOption) LookupDeliveryStreamResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDeliveryStreamResult, error) {
			args := v.(LookupDeliveryStreamArgs)
			r, err := LookupDeliveryStream(ctx, &args, opts...)
			var s LookupDeliveryStreamResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDeliveryStreamResultOutput)
}

type LookupDeliveryStreamOutputArgs struct {
	DeliveryStreamName pulumi.StringInput `pulumi:"deliveryStreamName"`
}

func (LookupDeliveryStreamOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDeliveryStreamArgs)(nil)).Elem()
}

type LookupDeliveryStreamResultOutput struct{ *pulumi.OutputState }

func (LookupDeliveryStreamResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDeliveryStreamResult)(nil)).Elem()
}

func (o LookupDeliveryStreamResultOutput) ToLookupDeliveryStreamResultOutput() LookupDeliveryStreamResultOutput {
	return o
}

func (o LookupDeliveryStreamResultOutput) ToLookupDeliveryStreamResultOutputWithContext(ctx context.Context) LookupDeliveryStreamResultOutput {
	return o
}

func (o LookupDeliveryStreamResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupDeliveryStreamResult] {
	return pulumix.Output[LookupDeliveryStreamResult]{
		OutputState: o.OutputState,
	}
}

func (o LookupDeliveryStreamResultOutput) AmazonOpenSearchServerlessDestinationConfiguration() DeliveryStreamAmazonOpenSearchServerlessDestinationConfigurationPtrOutput {
	return o.ApplyT(func(v LookupDeliveryStreamResult) *DeliveryStreamAmazonOpenSearchServerlessDestinationConfiguration {
		return v.AmazonOpenSearchServerlessDestinationConfiguration
	}).(DeliveryStreamAmazonOpenSearchServerlessDestinationConfigurationPtrOutput)
}

func (o LookupDeliveryStreamResultOutput) AmazonopensearchserviceDestinationConfiguration() DeliveryStreamAmazonopensearchserviceDestinationConfigurationPtrOutput {
	return o.ApplyT(func(v LookupDeliveryStreamResult) *DeliveryStreamAmazonopensearchserviceDestinationConfiguration {
		return v.AmazonopensearchserviceDestinationConfiguration
	}).(DeliveryStreamAmazonopensearchserviceDestinationConfigurationPtrOutput)
}

func (o LookupDeliveryStreamResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDeliveryStreamResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

func (o LookupDeliveryStreamResultOutput) DeliveryStreamEncryptionConfigurationInput() DeliveryStreamEncryptionConfigurationInputTypePtrOutput {
	return o.ApplyT(func(v LookupDeliveryStreamResult) *DeliveryStreamEncryptionConfigurationInputType {
		return v.DeliveryStreamEncryptionConfigurationInput
	}).(DeliveryStreamEncryptionConfigurationInputTypePtrOutput)
}

func (o LookupDeliveryStreamResultOutput) ElasticsearchDestinationConfiguration() DeliveryStreamElasticsearchDestinationConfigurationPtrOutput {
	return o.ApplyT(func(v LookupDeliveryStreamResult) *DeliveryStreamElasticsearchDestinationConfiguration {
		return v.ElasticsearchDestinationConfiguration
	}).(DeliveryStreamElasticsearchDestinationConfigurationPtrOutput)
}

func (o LookupDeliveryStreamResultOutput) ExtendedS3DestinationConfiguration() DeliveryStreamExtendedS3DestinationConfigurationPtrOutput {
	return o.ApplyT(func(v LookupDeliveryStreamResult) *DeliveryStreamExtendedS3DestinationConfiguration {
		return v.ExtendedS3DestinationConfiguration
	}).(DeliveryStreamExtendedS3DestinationConfigurationPtrOutput)
}

func (o LookupDeliveryStreamResultOutput) HttpEndpointDestinationConfiguration() DeliveryStreamHttpEndpointDestinationConfigurationPtrOutput {
	return o.ApplyT(func(v LookupDeliveryStreamResult) *DeliveryStreamHttpEndpointDestinationConfiguration {
		return v.HttpEndpointDestinationConfiguration
	}).(DeliveryStreamHttpEndpointDestinationConfigurationPtrOutput)
}

func (o LookupDeliveryStreamResultOutput) RedshiftDestinationConfiguration() DeliveryStreamRedshiftDestinationConfigurationPtrOutput {
	return o.ApplyT(func(v LookupDeliveryStreamResult) *DeliveryStreamRedshiftDestinationConfiguration {
		return v.RedshiftDestinationConfiguration
	}).(DeliveryStreamRedshiftDestinationConfigurationPtrOutput)
}

func (o LookupDeliveryStreamResultOutput) S3DestinationConfiguration() DeliveryStreamS3DestinationConfigurationPtrOutput {
	return o.ApplyT(func(v LookupDeliveryStreamResult) *DeliveryStreamS3DestinationConfiguration {
		return v.S3DestinationConfiguration
	}).(DeliveryStreamS3DestinationConfigurationPtrOutput)
}

func (o LookupDeliveryStreamResultOutput) SnowflakeDestinationConfiguration() DeliveryStreamSnowflakeDestinationConfigurationPtrOutput {
	return o.ApplyT(func(v LookupDeliveryStreamResult) *DeliveryStreamSnowflakeDestinationConfiguration {
		return v.SnowflakeDestinationConfiguration
	}).(DeliveryStreamSnowflakeDestinationConfigurationPtrOutput)
}

func (o LookupDeliveryStreamResultOutput) SplunkDestinationConfiguration() DeliveryStreamSplunkDestinationConfigurationPtrOutput {
	return o.ApplyT(func(v LookupDeliveryStreamResult) *DeliveryStreamSplunkDestinationConfiguration {
		return v.SplunkDestinationConfiguration
	}).(DeliveryStreamSplunkDestinationConfigurationPtrOutput)
}

func (o LookupDeliveryStreamResultOutput) Tags() DeliveryStreamTagArrayOutput {
	return o.ApplyT(func(v LookupDeliveryStreamResult) []DeliveryStreamTag { return v.Tags }).(DeliveryStreamTagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDeliveryStreamResultOutput{})
}
