# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables
from .. import _inputs as _root_inputs

__all__ = [
    'TrailDataResourceArgs',
    'TrailEventSelectorArgs',
    'TrailPropertiesArgs',
]

@pulumi.input_type
class TrailDataResourceArgs:
    def __init__(__self__, *,
                 type: pulumi.Input[str],
                 values: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudtrail-trail-dataresource.html
        :param pulumi.Input[str] type: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudtrail-trail-dataresource.html#cfn-cloudtrail-trail-dataresource-type
        :param pulumi.Input[Sequence[pulumi.Input[str]]] values: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudtrail-trail-dataresource.html#cfn-cloudtrail-trail-dataresource-values
        """
        pulumi.set(__self__, "type", type)
        if values is not None:
            pulumi.set(__self__, "values", values)

    @property
    @pulumi.getter(name="Type")
    def type(self) -> pulumi.Input[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudtrail-trail-dataresource.html#cfn-cloudtrail-trail-dataresource-type
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="Values")
    def values(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudtrail-trail-dataresource.html#cfn-cloudtrail-trail-dataresource-values
        """
        return pulumi.get(self, "values")

    @values.setter
    def values(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "values", value)


@pulumi.input_type
class TrailEventSelectorArgs:
    def __init__(__self__, *,
                 data_resources: Optional[pulumi.Input[Sequence[pulumi.Input['TrailDataResourceArgs']]]] = None,
                 include_management_events: Optional[pulumi.Input[bool]] = None,
                 read_write_type: Optional[pulumi.Input[str]] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudtrail-trail-eventselector.html
        :param pulumi.Input[Sequence[pulumi.Input['TrailDataResourceArgs']]] data_resources: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudtrail-trail-eventselector.html#cfn-cloudtrail-trail-eventselector-dataresources
        :param pulumi.Input[bool] include_management_events: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudtrail-trail-eventselector.html#cfn-cloudtrail-trail-eventselector-includemanagementevents
        :param pulumi.Input[str] read_write_type: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudtrail-trail-eventselector.html#cfn-cloudtrail-trail-eventselector-readwritetype
        """
        if data_resources is not None:
            pulumi.set(__self__, "data_resources", data_resources)
        if include_management_events is not None:
            pulumi.set(__self__, "include_management_events", include_management_events)
        if read_write_type is not None:
            pulumi.set(__self__, "read_write_type", read_write_type)

    @property
    @pulumi.getter(name="DataResources")
    def data_resources(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['TrailDataResourceArgs']]]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudtrail-trail-eventselector.html#cfn-cloudtrail-trail-eventselector-dataresources
        """
        return pulumi.get(self, "data_resources")

    @data_resources.setter
    def data_resources(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['TrailDataResourceArgs']]]]):
        pulumi.set(self, "data_resources", value)

    @property
    @pulumi.getter(name="IncludeManagementEvents")
    def include_management_events(self) -> Optional[pulumi.Input[bool]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudtrail-trail-eventselector.html#cfn-cloudtrail-trail-eventselector-includemanagementevents
        """
        return pulumi.get(self, "include_management_events")

    @include_management_events.setter
    def include_management_events(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "include_management_events", value)

    @property
    @pulumi.getter(name="ReadWriteType")
    def read_write_type(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudtrail-trail-eventselector.html#cfn-cloudtrail-trail-eventselector-readwritetype
        """
        return pulumi.get(self, "read_write_type")

    @read_write_type.setter
    def read_write_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "read_write_type", value)


@pulumi.input_type
class TrailPropertiesArgs:
    def __init__(__self__, *,
                 is_logging: pulumi.Input[bool],
                 s3_bucket_name: pulumi.Input[str],
                 cloud_watch_logs_log_group_arn: Optional[pulumi.Input[str]] = None,
                 cloud_watch_logs_role_arn: Optional[pulumi.Input[str]] = None,
                 enable_log_file_validation: Optional[pulumi.Input[bool]] = None,
                 event_selectors: Optional[pulumi.Input[Sequence[pulumi.Input['TrailEventSelectorArgs']]]] = None,
                 include_global_service_events: Optional[pulumi.Input[bool]] = None,
                 is_multi_region_trail: Optional[pulumi.Input[bool]] = None,
                 kms_key_id: Optional[pulumi.Input[str]] = None,
                 s3_key_prefix: Optional[pulumi.Input[str]] = None,
                 sns_topic_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]] = None,
                 trail_name: Optional[pulumi.Input[str]] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudtrail-trail.html
        :param pulumi.Input[bool] is_logging: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudtrail-trail.html#cfn-cloudtrail-trail-islogging
        :param pulumi.Input[str] s3_bucket_name: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudtrail-trail.html#cfn-cloudtrail-trail-s3bucketname
        :param pulumi.Input[str] cloud_watch_logs_log_group_arn: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudtrail-trail.html#cfn-cloudtrail-trail-cloudwatchlogsloggrouparn
        :param pulumi.Input[str] cloud_watch_logs_role_arn: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudtrail-trail.html#cfn-cloudtrail-trail-cloudwatchlogsrolearn
        :param pulumi.Input[bool] enable_log_file_validation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudtrail-trail.html#cfn-cloudtrail-trail-enablelogfilevalidation
        :param pulumi.Input[Sequence[pulumi.Input['TrailEventSelectorArgs']]] event_selectors: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudtrail-trail.html#cfn-cloudtrail-trail-eventselectors
        :param pulumi.Input[bool] include_global_service_events: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudtrail-trail.html#cfn-cloudtrail-trail-includeglobalserviceevents
        :param pulumi.Input[bool] is_multi_region_trail: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudtrail-trail.html#cfn-cloudtrail-trail-ismultiregiontrail
        :param pulumi.Input[str] kms_key_id: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudtrail-trail.html#cfn-cloudtrail-trail-kmskeyid
        :param pulumi.Input[str] s3_key_prefix: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudtrail-trail.html#cfn-cloudtrail-trail-s3keyprefix
        :param pulumi.Input[str] sns_topic_name: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudtrail-trail.html#cfn-cloudtrail-trail-snstopicname
        :param pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]] tags: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudtrail-trail.html#cfn-cloudtrail-trail-tags
        :param pulumi.Input[str] trail_name: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudtrail-trail.html#cfn-cloudtrail-trail-trailname
        """
        pulumi.set(__self__, "is_logging", is_logging)
        pulumi.set(__self__, "s3_bucket_name", s3_bucket_name)
        if cloud_watch_logs_log_group_arn is not None:
            pulumi.set(__self__, "cloud_watch_logs_log_group_arn", cloud_watch_logs_log_group_arn)
        if cloud_watch_logs_role_arn is not None:
            pulumi.set(__self__, "cloud_watch_logs_role_arn", cloud_watch_logs_role_arn)
        if enable_log_file_validation is not None:
            pulumi.set(__self__, "enable_log_file_validation", enable_log_file_validation)
        if event_selectors is not None:
            pulumi.set(__self__, "event_selectors", event_selectors)
        if include_global_service_events is not None:
            pulumi.set(__self__, "include_global_service_events", include_global_service_events)
        if is_multi_region_trail is not None:
            pulumi.set(__self__, "is_multi_region_trail", is_multi_region_trail)
        if kms_key_id is not None:
            pulumi.set(__self__, "kms_key_id", kms_key_id)
        if s3_key_prefix is not None:
            pulumi.set(__self__, "s3_key_prefix", s3_key_prefix)
        if sns_topic_name is not None:
            pulumi.set(__self__, "sns_topic_name", sns_topic_name)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if trail_name is not None:
            pulumi.set(__self__, "trail_name", trail_name)

    @property
    @pulumi.getter(name="IsLogging")
    def is_logging(self) -> pulumi.Input[bool]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudtrail-trail.html#cfn-cloudtrail-trail-islogging
        """
        return pulumi.get(self, "is_logging")

    @is_logging.setter
    def is_logging(self, value: pulumi.Input[bool]):
        pulumi.set(self, "is_logging", value)

    @property
    @pulumi.getter(name="S3BucketName")
    def s3_bucket_name(self) -> pulumi.Input[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudtrail-trail.html#cfn-cloudtrail-trail-s3bucketname
        """
        return pulumi.get(self, "s3_bucket_name")

    @s3_bucket_name.setter
    def s3_bucket_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "s3_bucket_name", value)

    @property
    @pulumi.getter(name="CloudWatchLogsLogGroupArn")
    def cloud_watch_logs_log_group_arn(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudtrail-trail.html#cfn-cloudtrail-trail-cloudwatchlogsloggrouparn
        """
        return pulumi.get(self, "cloud_watch_logs_log_group_arn")

    @cloud_watch_logs_log_group_arn.setter
    def cloud_watch_logs_log_group_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cloud_watch_logs_log_group_arn", value)

    @property
    @pulumi.getter(name="CloudWatchLogsRoleArn")
    def cloud_watch_logs_role_arn(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudtrail-trail.html#cfn-cloudtrail-trail-cloudwatchlogsrolearn
        """
        return pulumi.get(self, "cloud_watch_logs_role_arn")

    @cloud_watch_logs_role_arn.setter
    def cloud_watch_logs_role_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cloud_watch_logs_role_arn", value)

    @property
    @pulumi.getter(name="EnableLogFileValidation")
    def enable_log_file_validation(self) -> Optional[pulumi.Input[bool]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudtrail-trail.html#cfn-cloudtrail-trail-enablelogfilevalidation
        """
        return pulumi.get(self, "enable_log_file_validation")

    @enable_log_file_validation.setter
    def enable_log_file_validation(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_log_file_validation", value)

    @property
    @pulumi.getter(name="EventSelectors")
    def event_selectors(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['TrailEventSelectorArgs']]]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudtrail-trail.html#cfn-cloudtrail-trail-eventselectors
        """
        return pulumi.get(self, "event_selectors")

    @event_selectors.setter
    def event_selectors(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['TrailEventSelectorArgs']]]]):
        pulumi.set(self, "event_selectors", value)

    @property
    @pulumi.getter(name="IncludeGlobalServiceEvents")
    def include_global_service_events(self) -> Optional[pulumi.Input[bool]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudtrail-trail.html#cfn-cloudtrail-trail-includeglobalserviceevents
        """
        return pulumi.get(self, "include_global_service_events")

    @include_global_service_events.setter
    def include_global_service_events(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "include_global_service_events", value)

    @property
    @pulumi.getter(name="IsMultiRegionTrail")
    def is_multi_region_trail(self) -> Optional[pulumi.Input[bool]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudtrail-trail.html#cfn-cloudtrail-trail-ismultiregiontrail
        """
        return pulumi.get(self, "is_multi_region_trail")

    @is_multi_region_trail.setter
    def is_multi_region_trail(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "is_multi_region_trail", value)

    @property
    @pulumi.getter(name="KMSKeyId")
    def kms_key_id(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudtrail-trail.html#cfn-cloudtrail-trail-kmskeyid
        """
        return pulumi.get(self, "kms_key_id")

    @kms_key_id.setter
    def kms_key_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "kms_key_id", value)

    @property
    @pulumi.getter(name="S3KeyPrefix")
    def s3_key_prefix(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudtrail-trail.html#cfn-cloudtrail-trail-s3keyprefix
        """
        return pulumi.get(self, "s3_key_prefix")

    @s3_key_prefix.setter
    def s3_key_prefix(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "s3_key_prefix", value)

    @property
    @pulumi.getter(name="SnsTopicName")
    def sns_topic_name(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudtrail-trail.html#cfn-cloudtrail-trail-snstopicname
        """
        return pulumi.get(self, "sns_topic_name")

    @sns_topic_name.setter
    def sns_topic_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sns_topic_name", value)

    @property
    @pulumi.getter(name="Tags")
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudtrail-trail.html#cfn-cloudtrail-trail-tags
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="TrailName")
    def trail_name(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudtrail-trail.html#cfn-cloudtrail-trail-trailname
        """
        return pulumi.get(self, "trail_name")

    @trail_name.setter
    def trail_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "trail_name", value)

