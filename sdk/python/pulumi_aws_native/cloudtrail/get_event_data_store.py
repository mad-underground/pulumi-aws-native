# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs

__all__ = [
    'GetEventDataStoreResult',
    'AwaitableGetEventDataStoreResult',
    'get_event_data_store',
    'get_event_data_store_output',
]

@pulumi.output_type
class GetEventDataStoreResult:
    def __init__(__self__, advanced_event_selectors=None, created_timestamp=None, event_data_store_arn=None, ingestion_enabled=None, kms_key_id=None, multi_region_enabled=None, name=None, organization_enabled=None, retention_period=None, status=None, tags=None, termination_protection_enabled=None, updated_timestamp=None):
        if advanced_event_selectors and not isinstance(advanced_event_selectors, list):
            raise TypeError("Expected argument 'advanced_event_selectors' to be a list")
        pulumi.set(__self__, "advanced_event_selectors", advanced_event_selectors)
        if created_timestamp and not isinstance(created_timestamp, str):
            raise TypeError("Expected argument 'created_timestamp' to be a str")
        pulumi.set(__self__, "created_timestamp", created_timestamp)
        if event_data_store_arn and not isinstance(event_data_store_arn, str):
            raise TypeError("Expected argument 'event_data_store_arn' to be a str")
        pulumi.set(__self__, "event_data_store_arn", event_data_store_arn)
        if ingestion_enabled and not isinstance(ingestion_enabled, bool):
            raise TypeError("Expected argument 'ingestion_enabled' to be a bool")
        pulumi.set(__self__, "ingestion_enabled", ingestion_enabled)
        if kms_key_id and not isinstance(kms_key_id, str):
            raise TypeError("Expected argument 'kms_key_id' to be a str")
        pulumi.set(__self__, "kms_key_id", kms_key_id)
        if multi_region_enabled and not isinstance(multi_region_enabled, bool):
            raise TypeError("Expected argument 'multi_region_enabled' to be a bool")
        pulumi.set(__self__, "multi_region_enabled", multi_region_enabled)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if organization_enabled and not isinstance(organization_enabled, bool):
            raise TypeError("Expected argument 'organization_enabled' to be a bool")
        pulumi.set(__self__, "organization_enabled", organization_enabled)
        if retention_period and not isinstance(retention_period, int):
            raise TypeError("Expected argument 'retention_period' to be a int")
        pulumi.set(__self__, "retention_period", retention_period)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)
        if termination_protection_enabled and not isinstance(termination_protection_enabled, bool):
            raise TypeError("Expected argument 'termination_protection_enabled' to be a bool")
        pulumi.set(__self__, "termination_protection_enabled", termination_protection_enabled)
        if updated_timestamp and not isinstance(updated_timestamp, str):
            raise TypeError("Expected argument 'updated_timestamp' to be a str")
        pulumi.set(__self__, "updated_timestamp", updated_timestamp)

    @property
    @pulumi.getter(name="advancedEventSelectors")
    def advanced_event_selectors(self) -> Optional[Sequence['outputs.EventDataStoreAdvancedEventSelector']]:
        """
        The advanced event selectors that were used to select events for the data store.
        """
        return pulumi.get(self, "advanced_event_selectors")

    @property
    @pulumi.getter(name="createdTimestamp")
    def created_timestamp(self) -> Optional[str]:
        """
        The timestamp of the event data store's creation.
        """
        return pulumi.get(self, "created_timestamp")

    @property
    @pulumi.getter(name="eventDataStoreArn")
    def event_data_store_arn(self) -> Optional[str]:
        """
        The ARN of the event data store.
        """
        return pulumi.get(self, "event_data_store_arn")

    @property
    @pulumi.getter(name="ingestionEnabled")
    def ingestion_enabled(self) -> Optional[bool]:
        """
        Indicates whether the event data store is ingesting events.
        """
        return pulumi.get(self, "ingestion_enabled")

    @property
    @pulumi.getter(name="kmsKeyId")
    def kms_key_id(self) -> Optional[str]:
        """
        Specifies the KMS key ID to use to encrypt the events delivered by CloudTrail. The value can be an alias name prefixed by 'alias/', a fully specified ARN to an alias, a fully specified ARN to a key, or a globally unique identifier.
        """
        return pulumi.get(self, "kms_key_id")

    @property
    @pulumi.getter(name="multiRegionEnabled")
    def multi_region_enabled(self) -> Optional[bool]:
        """
        Indicates whether the event data store includes events from all regions, or only from the region in which it was created.
        """
        return pulumi.get(self, "multi_region_enabled")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        The name of the event data store.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="organizationEnabled")
    def organization_enabled(self) -> Optional[bool]:
        """
        Indicates that an event data store is collecting logged events for an organization.
        """
        return pulumi.get(self, "organization_enabled")

    @property
    @pulumi.getter(name="retentionPeriod")
    def retention_period(self) -> Optional[int]:
        """
        The retention period, in days.
        """
        return pulumi.get(self, "retention_period")

    @property
    @pulumi.getter
    def status(self) -> Optional[str]:
        """
        The status of an event data store. Values are STARTING_INGESTION, ENABLED, STOPPING_INGESTION, STOPPED_INGESTION and PENDING_DELETION.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['outputs.EventDataStoreTag']]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="terminationProtectionEnabled")
    def termination_protection_enabled(self) -> Optional[bool]:
        """
        Indicates whether the event data store is protected from termination.
        """
        return pulumi.get(self, "termination_protection_enabled")

    @property
    @pulumi.getter(name="updatedTimestamp")
    def updated_timestamp(self) -> Optional[str]:
        """
        The timestamp showing when an event data store was updated, if applicable. UpdatedTimestamp is always either the same or newer than the time shown in CreatedTimestamp.
        """
        return pulumi.get(self, "updated_timestamp")


class AwaitableGetEventDataStoreResult(GetEventDataStoreResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetEventDataStoreResult(
            advanced_event_selectors=self.advanced_event_selectors,
            created_timestamp=self.created_timestamp,
            event_data_store_arn=self.event_data_store_arn,
            ingestion_enabled=self.ingestion_enabled,
            kms_key_id=self.kms_key_id,
            multi_region_enabled=self.multi_region_enabled,
            name=self.name,
            organization_enabled=self.organization_enabled,
            retention_period=self.retention_period,
            status=self.status,
            tags=self.tags,
            termination_protection_enabled=self.termination_protection_enabled,
            updated_timestamp=self.updated_timestamp)


def get_event_data_store(event_data_store_arn: Optional[str] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetEventDataStoreResult:
    """
    A storage lake of event data against which you can run complex SQL-based queries. An event data store can include events that you have logged on your account from the last 90 to 2555 days (about three months to up to seven years).


    :param str event_data_store_arn: The ARN of the event data store.
    """
    __args__ = dict()
    __args__['eventDataStoreArn'] = event_data_store_arn
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:cloudtrail:getEventDataStore', __args__, opts=opts, typ=GetEventDataStoreResult).value

    return AwaitableGetEventDataStoreResult(
        advanced_event_selectors=pulumi.get(__ret__, 'advanced_event_selectors'),
        created_timestamp=pulumi.get(__ret__, 'created_timestamp'),
        event_data_store_arn=pulumi.get(__ret__, 'event_data_store_arn'),
        ingestion_enabled=pulumi.get(__ret__, 'ingestion_enabled'),
        kms_key_id=pulumi.get(__ret__, 'kms_key_id'),
        multi_region_enabled=pulumi.get(__ret__, 'multi_region_enabled'),
        name=pulumi.get(__ret__, 'name'),
        organization_enabled=pulumi.get(__ret__, 'organization_enabled'),
        retention_period=pulumi.get(__ret__, 'retention_period'),
        status=pulumi.get(__ret__, 'status'),
        tags=pulumi.get(__ret__, 'tags'),
        termination_protection_enabled=pulumi.get(__ret__, 'termination_protection_enabled'),
        updated_timestamp=pulumi.get(__ret__, 'updated_timestamp'))


@_utilities.lift_output_func(get_event_data_store)
def get_event_data_store_output(event_data_store_arn: Optional[pulumi.Input[str]] = None,
                                opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetEventDataStoreResult]:
    """
    A storage lake of event data against which you can run complex SQL-based queries. An event data store can include events that you have logged on your account from the last 90 to 2555 days (about three months to up to seven years).


    :param str event_data_store_arn: The ARN of the event data store.
    """
    ...
