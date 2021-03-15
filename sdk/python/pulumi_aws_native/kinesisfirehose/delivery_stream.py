# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables
from . import outputs
from .. import _inputs as _root_inputs
from .. import outputs as _root_outputs
from ._inputs import *

__all__ = ['DeliveryStream']


class DeliveryStream(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 delivery_stream_encryption_configuration_input: Optional[pulumi.Input[pulumi.InputType['DeliveryStreamDeliveryStreamEncryptionConfigurationInputArgs']]] = None,
                 delivery_stream_name: Optional[pulumi.Input[str]] = None,
                 delivery_stream_type: Optional[pulumi.Input[str]] = None,
                 elasticsearch_destination_configuration: Optional[pulumi.Input[pulumi.InputType['DeliveryStreamElasticsearchDestinationConfigurationArgs']]] = None,
                 extended_s3_destination_configuration: Optional[pulumi.Input[pulumi.InputType['DeliveryStreamExtendedS3DestinationConfigurationArgs']]] = None,
                 http_endpoint_destination_configuration: Optional[pulumi.Input[pulumi.InputType['DeliveryStreamHttpEndpointDestinationConfigurationArgs']]] = None,
                 kinesis_stream_source_configuration: Optional[pulumi.Input[pulumi.InputType['DeliveryStreamKinesisStreamSourceConfigurationArgs']]] = None,
                 redshift_destination_configuration: Optional[pulumi.Input[pulumi.InputType['DeliveryStreamRedshiftDestinationConfigurationArgs']]] = None,
                 s3_destination_configuration: Optional[pulumi.Input[pulumi.InputType['DeliveryStreamS3DestinationConfigurationArgs']]] = None,
                 splunk_destination_configuration: Optional[pulumi.Input[pulumi.InputType['DeliveryStreamSplunkDestinationConfigurationArgs']]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['_root_inputs.TagArgs']]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisfirehose-deliverystream.html

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['DeliveryStreamDeliveryStreamEncryptionConfigurationInputArgs']] delivery_stream_encryption_configuration_input: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisfirehose-deliverystream.html#cfn-kinesisfirehose-deliverystream-deliverystreamencryptionconfigurationinput
        :param pulumi.Input[str] delivery_stream_name: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisfirehose-deliverystream.html#cfn-kinesisfirehose-deliverystream-deliverystreamname
        :param pulumi.Input[str] delivery_stream_type: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisfirehose-deliverystream.html#cfn-kinesisfirehose-deliverystream-deliverystreamtype
        :param pulumi.Input[pulumi.InputType['DeliveryStreamElasticsearchDestinationConfigurationArgs']] elasticsearch_destination_configuration: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisfirehose-deliverystream.html#cfn-kinesisfirehose-deliverystream-elasticsearchdestinationconfiguration
        :param pulumi.Input[pulumi.InputType['DeliveryStreamExtendedS3DestinationConfigurationArgs']] extended_s3_destination_configuration: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisfirehose-deliverystream.html#cfn-kinesisfirehose-deliverystream-extendeds3destinationconfiguration
        :param pulumi.Input[pulumi.InputType['DeliveryStreamHttpEndpointDestinationConfigurationArgs']] http_endpoint_destination_configuration: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisfirehose-deliverystream.html#cfn-kinesisfirehose-deliverystream-httpendpointdestinationconfiguration
        :param pulumi.Input[pulumi.InputType['DeliveryStreamKinesisStreamSourceConfigurationArgs']] kinesis_stream_source_configuration: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisfirehose-deliverystream.html#cfn-kinesisfirehose-deliverystream-kinesisstreamsourceconfiguration
        :param pulumi.Input[pulumi.InputType['DeliveryStreamRedshiftDestinationConfigurationArgs']] redshift_destination_configuration: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisfirehose-deliverystream.html#cfn-kinesisfirehose-deliverystream-redshiftdestinationconfiguration
        :param pulumi.Input[pulumi.InputType['DeliveryStreamS3DestinationConfigurationArgs']] s3_destination_configuration: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisfirehose-deliverystream.html#cfn-kinesisfirehose-deliverystream-s3destinationconfiguration
        :param pulumi.Input[pulumi.InputType['DeliveryStreamSplunkDestinationConfigurationArgs']] splunk_destination_configuration: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisfirehose-deliverystream.html#cfn-kinesisfirehose-deliverystream-splunkdestinationconfiguration
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['_root_inputs.TagArgs']]]] tags: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisfirehose-deliverystream.html#cfn-kinesisfirehose-deliverystream-tags
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['delivery_stream_encryption_configuration_input'] = delivery_stream_encryption_configuration_input
            __props__['delivery_stream_name'] = delivery_stream_name
            __props__['delivery_stream_type'] = delivery_stream_type
            __props__['elasticsearch_destination_configuration'] = elasticsearch_destination_configuration
            __props__['extended_s3_destination_configuration'] = extended_s3_destination_configuration
            __props__['http_endpoint_destination_configuration'] = http_endpoint_destination_configuration
            __props__['kinesis_stream_source_configuration'] = kinesis_stream_source_configuration
            __props__['redshift_destination_configuration'] = redshift_destination_configuration
            __props__['s3_destination_configuration'] = s3_destination_configuration
            __props__['splunk_destination_configuration'] = splunk_destination_configuration
            __props__['tags'] = tags
            __props__['arn'] = None
        super(DeliveryStream, __self__).__init__(
            'aws-native:KinesisFirehose:DeliveryStream',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'DeliveryStream':
        """
        Get an existing DeliveryStream resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return DeliveryStream(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="deliveryStreamEncryptionConfigurationInput")
    def delivery_stream_encryption_configuration_input(self) -> pulumi.Output[Optional['outputs.DeliveryStreamDeliveryStreamEncryptionConfigurationInput']]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisfirehose-deliverystream.html#cfn-kinesisfirehose-deliverystream-deliverystreamencryptionconfigurationinput
        """
        return pulumi.get(self, "delivery_stream_encryption_configuration_input")

    @property
    @pulumi.getter(name="deliveryStreamName")
    def delivery_stream_name(self) -> pulumi.Output[Optional[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisfirehose-deliverystream.html#cfn-kinesisfirehose-deliverystream-deliverystreamname
        """
        return pulumi.get(self, "delivery_stream_name")

    @property
    @pulumi.getter(name="deliveryStreamType")
    def delivery_stream_type(self) -> pulumi.Output[Optional[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisfirehose-deliverystream.html#cfn-kinesisfirehose-deliverystream-deliverystreamtype
        """
        return pulumi.get(self, "delivery_stream_type")

    @property
    @pulumi.getter(name="elasticsearchDestinationConfiguration")
    def elasticsearch_destination_configuration(self) -> pulumi.Output[Optional['outputs.DeliveryStreamElasticsearchDestinationConfiguration']]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisfirehose-deliverystream.html#cfn-kinesisfirehose-deliverystream-elasticsearchdestinationconfiguration
        """
        return pulumi.get(self, "elasticsearch_destination_configuration")

    @property
    @pulumi.getter(name="extendedS3DestinationConfiguration")
    def extended_s3_destination_configuration(self) -> pulumi.Output[Optional['outputs.DeliveryStreamExtendedS3DestinationConfiguration']]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisfirehose-deliverystream.html#cfn-kinesisfirehose-deliverystream-extendeds3destinationconfiguration
        """
        return pulumi.get(self, "extended_s3_destination_configuration")

    @property
    @pulumi.getter(name="httpEndpointDestinationConfiguration")
    def http_endpoint_destination_configuration(self) -> pulumi.Output[Optional['outputs.DeliveryStreamHttpEndpointDestinationConfiguration']]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisfirehose-deliverystream.html#cfn-kinesisfirehose-deliverystream-httpendpointdestinationconfiguration
        """
        return pulumi.get(self, "http_endpoint_destination_configuration")

    @property
    @pulumi.getter(name="kinesisStreamSourceConfiguration")
    def kinesis_stream_source_configuration(self) -> pulumi.Output[Optional['outputs.DeliveryStreamKinesisStreamSourceConfiguration']]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisfirehose-deliverystream.html#cfn-kinesisfirehose-deliverystream-kinesisstreamsourceconfiguration
        """
        return pulumi.get(self, "kinesis_stream_source_configuration")

    @property
    @pulumi.getter(name="redshiftDestinationConfiguration")
    def redshift_destination_configuration(self) -> pulumi.Output[Optional['outputs.DeliveryStreamRedshiftDestinationConfiguration']]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisfirehose-deliverystream.html#cfn-kinesisfirehose-deliverystream-redshiftdestinationconfiguration
        """
        return pulumi.get(self, "redshift_destination_configuration")

    @property
    @pulumi.getter(name="s3DestinationConfiguration")
    def s3_destination_configuration(self) -> pulumi.Output[Optional['outputs.DeliveryStreamS3DestinationConfiguration']]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisfirehose-deliverystream.html#cfn-kinesisfirehose-deliverystream-s3destinationconfiguration
        """
        return pulumi.get(self, "s3_destination_configuration")

    @property
    @pulumi.getter(name="splunkDestinationConfiguration")
    def splunk_destination_configuration(self) -> pulumi.Output[Optional['outputs.DeliveryStreamSplunkDestinationConfiguration']]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisfirehose-deliverystream.html#cfn-kinesisfirehose-deliverystream-splunkdestinationconfiguration
        """
        return pulumi.get(self, "splunk_destination_configuration")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence['_root_outputs.Tag']]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisfirehose-deliverystream.html#cfn-kinesisfirehose-deliverystream-tags
        """
        return pulumi.get(self, "tags")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

