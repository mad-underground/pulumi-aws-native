# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = [
    'HealthCheckHealthCheckTagArgs',
    'HealthCheckPropertiesArgs',
    'HostedZoneHostedZoneConfigArgs',
    'HostedZoneHostedZoneTagArgs',
    'HostedZonePropertiesArgs',
    'HostedZoneQueryLoggingConfigArgs',
    'HostedZoneVPCArgs',
]

@pulumi.input_type
class HealthCheckHealthCheckTagArgs:
    def __init__(__self__, *,
                 key: pulumi.Input[str],
                 value: pulumi.Input[str]):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-healthcheck-healthchecktag.html
        :param pulumi.Input[str] key: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-healthcheck-healthchecktag.html#cfn-route53-healthcheck-healthchecktag-key
        :param pulumi.Input[str] value: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-healthcheck-healthchecktag.html#cfn-route53-healthcheck-healthchecktag-value
        """
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter(name="Key")
    def key(self) -> pulumi.Input[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-healthcheck-healthchecktag.html#cfn-route53-healthcheck-healthchecktag-key
        """
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: pulumi.Input[str]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter(name="Value")
    def value(self) -> pulumi.Input[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-healthcheck-healthchecktag.html#cfn-route53-healthcheck-healthchecktag-value
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: pulumi.Input[str]):
        pulumi.set(self, "value", value)


@pulumi.input_type
class HealthCheckPropertiesArgs:
    def __init__(__self__, *,
                 health_check_config: pulumi.Input[Union[Any, str]],
                 health_check_tags: Optional[pulumi.Input[Sequence[pulumi.Input['HealthCheckHealthCheckTagArgs']]]] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-healthcheck.html
        :param pulumi.Input[Union[Any, str]] health_check_config: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-healthcheck.html#cfn-route53-healthcheck-healthcheckconfig
        :param pulumi.Input[Sequence[pulumi.Input['HealthCheckHealthCheckTagArgs']]] health_check_tags: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-healthcheck.html#cfn-route53-healthcheck-healthchecktags
        """
        pulumi.set(__self__, "health_check_config", health_check_config)
        if health_check_tags is not None:
            pulumi.set(__self__, "health_check_tags", health_check_tags)

    @property
    @pulumi.getter(name="HealthCheckConfig")
    def health_check_config(self) -> pulumi.Input[Union[Any, str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-healthcheck.html#cfn-route53-healthcheck-healthcheckconfig
        """
        return pulumi.get(self, "health_check_config")

    @health_check_config.setter
    def health_check_config(self, value: pulumi.Input[Union[Any, str]]):
        pulumi.set(self, "health_check_config", value)

    @property
    @pulumi.getter(name="HealthCheckTags")
    def health_check_tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['HealthCheckHealthCheckTagArgs']]]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-healthcheck.html#cfn-route53-healthcheck-healthchecktags
        """
        return pulumi.get(self, "health_check_tags")

    @health_check_tags.setter
    def health_check_tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['HealthCheckHealthCheckTagArgs']]]]):
        pulumi.set(self, "health_check_tags", value)


@pulumi.input_type
class HostedZoneHostedZoneConfigArgs:
    def __init__(__self__, *,
                 comment: Optional[pulumi.Input[str]] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-hostedzone-hostedzoneconfig.html
        :param pulumi.Input[str] comment: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-hostedzone-hostedzoneconfig.html#cfn-route53-hostedzone-hostedzoneconfig-comment
        """
        if comment is not None:
            pulumi.set(__self__, "comment", comment)

    @property
    @pulumi.getter(name="Comment")
    def comment(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-hostedzone-hostedzoneconfig.html#cfn-route53-hostedzone-hostedzoneconfig-comment
        """
        return pulumi.get(self, "comment")

    @comment.setter
    def comment(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "comment", value)


@pulumi.input_type
class HostedZoneHostedZoneTagArgs:
    def __init__(__self__, *,
                 key: pulumi.Input[str],
                 value: pulumi.Input[str]):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-hostedzone-hostedzonetag.html
        :param pulumi.Input[str] key: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-hostedzone-hostedzonetag.html#cfn-route53-hostedzone-hostedzonetag-key
        :param pulumi.Input[str] value: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-hostedzone-hostedzonetag.html#cfn-route53-hostedzone-hostedzonetag-value
        """
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter(name="Key")
    def key(self) -> pulumi.Input[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-hostedzone-hostedzonetag.html#cfn-route53-hostedzone-hostedzonetag-key
        """
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: pulumi.Input[str]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter(name="Value")
    def value(self) -> pulumi.Input[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-hostedzone-hostedzonetag.html#cfn-route53-hostedzone-hostedzonetag-value
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: pulumi.Input[str]):
        pulumi.set(self, "value", value)


@pulumi.input_type
class HostedZonePropertiesArgs:
    def __init__(__self__, *,
                 name: pulumi.Input[str],
                 hosted_zone_config: Optional[pulumi.Input['HostedZoneHostedZoneConfigArgs']] = None,
                 hosted_zone_tags: Optional[pulumi.Input[Sequence[pulumi.Input['HostedZoneHostedZoneTagArgs']]]] = None,
                 query_logging_config: Optional[pulumi.Input['HostedZoneQueryLoggingConfigArgs']] = None,
                 vpcs: Optional[pulumi.Input[Sequence[pulumi.Input['HostedZoneVPCArgs']]]] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-hostedzone.html
        :param pulumi.Input[str] name: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-hostedzone.html#cfn-route53-hostedzone-name
        :param pulumi.Input['HostedZoneHostedZoneConfigArgs'] hosted_zone_config: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-hostedzone.html#cfn-route53-hostedzone-hostedzoneconfig
        :param pulumi.Input[Sequence[pulumi.Input['HostedZoneHostedZoneTagArgs']]] hosted_zone_tags: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-hostedzone.html#cfn-route53-hostedzone-hostedzonetags
        :param pulumi.Input['HostedZoneQueryLoggingConfigArgs'] query_logging_config: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-hostedzone.html#cfn-route53-hostedzone-queryloggingconfig
        :param pulumi.Input[Sequence[pulumi.Input['HostedZoneVPCArgs']]] vpcs: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-hostedzone.html#cfn-route53-hostedzone-vpcs
        """
        pulumi.set(__self__, "name", name)
        if hosted_zone_config is not None:
            pulumi.set(__self__, "hosted_zone_config", hosted_zone_config)
        if hosted_zone_tags is not None:
            pulumi.set(__self__, "hosted_zone_tags", hosted_zone_tags)
        if query_logging_config is not None:
            pulumi.set(__self__, "query_logging_config", query_logging_config)
        if vpcs is not None:
            pulumi.set(__self__, "vpcs", vpcs)

    @property
    @pulumi.getter(name="Name")
    def name(self) -> pulumi.Input[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-hostedzone.html#cfn-route53-hostedzone-name
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="HostedZoneConfig")
    def hosted_zone_config(self) -> Optional[pulumi.Input['HostedZoneHostedZoneConfigArgs']]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-hostedzone.html#cfn-route53-hostedzone-hostedzoneconfig
        """
        return pulumi.get(self, "hosted_zone_config")

    @hosted_zone_config.setter
    def hosted_zone_config(self, value: Optional[pulumi.Input['HostedZoneHostedZoneConfigArgs']]):
        pulumi.set(self, "hosted_zone_config", value)

    @property
    @pulumi.getter(name="HostedZoneTags")
    def hosted_zone_tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['HostedZoneHostedZoneTagArgs']]]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-hostedzone.html#cfn-route53-hostedzone-hostedzonetags
        """
        return pulumi.get(self, "hosted_zone_tags")

    @hosted_zone_tags.setter
    def hosted_zone_tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['HostedZoneHostedZoneTagArgs']]]]):
        pulumi.set(self, "hosted_zone_tags", value)

    @property
    @pulumi.getter(name="QueryLoggingConfig")
    def query_logging_config(self) -> Optional[pulumi.Input['HostedZoneQueryLoggingConfigArgs']]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-hostedzone.html#cfn-route53-hostedzone-queryloggingconfig
        """
        return pulumi.get(self, "query_logging_config")

    @query_logging_config.setter
    def query_logging_config(self, value: Optional[pulumi.Input['HostedZoneQueryLoggingConfigArgs']]):
        pulumi.set(self, "query_logging_config", value)

    @property
    @pulumi.getter(name="VPCs")
    def vpcs(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['HostedZoneVPCArgs']]]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-hostedzone.html#cfn-route53-hostedzone-vpcs
        """
        return pulumi.get(self, "vpcs")

    @vpcs.setter
    def vpcs(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['HostedZoneVPCArgs']]]]):
        pulumi.set(self, "vpcs", value)


@pulumi.input_type
class HostedZoneQueryLoggingConfigArgs:
    def __init__(__self__, *,
                 cloud_watch_logs_log_group_arn: pulumi.Input[str]):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-hostedzone-queryloggingconfig.html
        :param pulumi.Input[str] cloud_watch_logs_log_group_arn: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-hostedzone-queryloggingconfig.html#cfn-route53-hostedzone-queryloggingconfig-cloudwatchlogsloggrouparn
        """
        pulumi.set(__self__, "cloud_watch_logs_log_group_arn", cloud_watch_logs_log_group_arn)

    @property
    @pulumi.getter(name="CloudWatchLogsLogGroupArn")
    def cloud_watch_logs_log_group_arn(self) -> pulumi.Input[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-hostedzone-queryloggingconfig.html#cfn-route53-hostedzone-queryloggingconfig-cloudwatchlogsloggrouparn
        """
        return pulumi.get(self, "cloud_watch_logs_log_group_arn")

    @cloud_watch_logs_log_group_arn.setter
    def cloud_watch_logs_log_group_arn(self, value: pulumi.Input[str]):
        pulumi.set(self, "cloud_watch_logs_log_group_arn", value)


@pulumi.input_type
class HostedZoneVPCArgs:
    def __init__(__self__, *,
                 vpc_id: pulumi.Input[str],
                 vpc_region: pulumi.Input[str]):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-hostedzone-vpc.html
        :param pulumi.Input[str] vpc_id: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-hostedzone-vpc.html#cfn-route53-hostedzone-vpc-vpcid
        :param pulumi.Input[str] vpc_region: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-hostedzone-vpc.html#cfn-route53-hostedzone-vpc-vpcregion
        """
        pulumi.set(__self__, "vpc_id", vpc_id)
        pulumi.set(__self__, "vpc_region", vpc_region)

    @property
    @pulumi.getter(name="VPCId")
    def vpc_id(self) -> pulumi.Input[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-hostedzone-vpc.html#cfn-route53-hostedzone-vpc-vpcid
        """
        return pulumi.get(self, "vpc_id")

    @vpc_id.setter
    def vpc_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "vpc_id", value)

    @property
    @pulumi.getter(name="VPCRegion")
    def vpc_region(self) -> pulumi.Input[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-hostedzone-vpc.html#cfn-route53-hostedzone-vpc-vpcregion
        """
        return pulumi.get(self, "vpc_region")

    @vpc_region.setter
    def vpc_region(self, value: pulumi.Input[str]):
        pulumi.set(self, "vpc_region", value)


