# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from .. import _inputs as _root_inputs
from .. import outputs as _root_outputs
from ._enums import *

__all__ = ['TrackerArgs', 'Tracker']

@pulumi.input_type
class TrackerArgs:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 event_bridge_enabled: Optional[pulumi.Input[bool]] = None,
                 kms_key_enable_geospatial_queries: Optional[pulumi.Input[bool]] = None,
                 kms_key_id: Optional[pulumi.Input[str]] = None,
                 position_filtering: Optional[pulumi.Input['TrackerPositionFiltering']] = None,
                 pricing_plan: Optional[pulumi.Input['TrackerPricingPlan']] = None,
                 pricing_plan_data_source: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]] = None,
                 tracker_name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Tracker resource.
        :param pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]] tags: An array of key-value pairs to apply to this resource.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if event_bridge_enabled is not None:
            pulumi.set(__self__, "event_bridge_enabled", event_bridge_enabled)
        if kms_key_enable_geospatial_queries is not None:
            pulumi.set(__self__, "kms_key_enable_geospatial_queries", kms_key_enable_geospatial_queries)
        if kms_key_id is not None:
            pulumi.set(__self__, "kms_key_id", kms_key_id)
        if position_filtering is not None:
            pulumi.set(__self__, "position_filtering", position_filtering)
        if pricing_plan is not None:
            pulumi.set(__self__, "pricing_plan", pricing_plan)
        if pricing_plan_data_source is not None:
            pulumi.set(__self__, "pricing_plan_data_source", pricing_plan_data_source)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if tracker_name is not None:
            pulumi.set(__self__, "tracker_name", tracker_name)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="eventBridgeEnabled")
    def event_bridge_enabled(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "event_bridge_enabled")

    @event_bridge_enabled.setter
    def event_bridge_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "event_bridge_enabled", value)

    @property
    @pulumi.getter(name="kmsKeyEnableGeospatialQueries")
    def kms_key_enable_geospatial_queries(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "kms_key_enable_geospatial_queries")

    @kms_key_enable_geospatial_queries.setter
    def kms_key_enable_geospatial_queries(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "kms_key_enable_geospatial_queries", value)

    @property
    @pulumi.getter(name="kmsKeyId")
    def kms_key_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "kms_key_id")

    @kms_key_id.setter
    def kms_key_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "kms_key_id", value)

    @property
    @pulumi.getter(name="positionFiltering")
    def position_filtering(self) -> Optional[pulumi.Input['TrackerPositionFiltering']]:
        return pulumi.get(self, "position_filtering")

    @position_filtering.setter
    def position_filtering(self, value: Optional[pulumi.Input['TrackerPositionFiltering']]):
        pulumi.set(self, "position_filtering", value)

    @property
    @pulumi.getter(name="pricingPlan")
    def pricing_plan(self) -> Optional[pulumi.Input['TrackerPricingPlan']]:
        return pulumi.get(self, "pricing_plan")

    @pricing_plan.setter
    def pricing_plan(self, value: Optional[pulumi.Input['TrackerPricingPlan']]):
        pulumi.set(self, "pricing_plan", value)

    @property
    @pulumi.getter(name="pricingPlanDataSource")
    def pricing_plan_data_source(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "pricing_plan_data_source")

    @pricing_plan_data_source.setter
    def pricing_plan_data_source(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "pricing_plan_data_source", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]]:
        """
        An array of key-value pairs to apply to this resource.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="trackerName")
    def tracker_name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "tracker_name")

    @tracker_name.setter
    def tracker_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tracker_name", value)


class Tracker(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 event_bridge_enabled: Optional[pulumi.Input[bool]] = None,
                 kms_key_enable_geospatial_queries: Optional[pulumi.Input[bool]] = None,
                 kms_key_id: Optional[pulumi.Input[str]] = None,
                 position_filtering: Optional[pulumi.Input['TrackerPositionFiltering']] = None,
                 pricing_plan: Optional[pulumi.Input['TrackerPricingPlan']] = None,
                 pricing_plan_data_source: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['_root_inputs.TagArgs']]]]] = None,
                 tracker_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Definition of AWS::Location::Tracker Resource Type

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['_root_inputs.TagArgs']]]] tags: An array of key-value pairs to apply to this resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[TrackerArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Definition of AWS::Location::Tracker Resource Type

        :param str resource_name: The name of the resource.
        :param TrackerArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(TrackerArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 event_bridge_enabled: Optional[pulumi.Input[bool]] = None,
                 kms_key_enable_geospatial_queries: Optional[pulumi.Input[bool]] = None,
                 kms_key_id: Optional[pulumi.Input[str]] = None,
                 position_filtering: Optional[pulumi.Input['TrackerPositionFiltering']] = None,
                 pricing_plan: Optional[pulumi.Input['TrackerPricingPlan']] = None,
                 pricing_plan_data_source: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['_root_inputs.TagArgs']]]]] = None,
                 tracker_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = TrackerArgs.__new__(TrackerArgs)

            __props__.__dict__["description"] = description
            __props__.__dict__["event_bridge_enabled"] = event_bridge_enabled
            __props__.__dict__["kms_key_enable_geospatial_queries"] = kms_key_enable_geospatial_queries
            __props__.__dict__["kms_key_id"] = kms_key_id
            __props__.__dict__["position_filtering"] = position_filtering
            __props__.__dict__["pricing_plan"] = pricing_plan
            __props__.__dict__["pricing_plan_data_source"] = pricing_plan_data_source
            __props__.__dict__["tags"] = tags
            __props__.__dict__["tracker_name"] = tracker_name
            __props__.__dict__["arn"] = None
            __props__.__dict__["create_time"] = None
            __props__.__dict__["tracker_arn"] = None
            __props__.__dict__["update_time"] = None
        replace_on_changes = pulumi.ResourceOptions(replace_on_changes=["kms_key_id", "tracker_name"])
        opts = pulumi.ResourceOptions.merge(opts, replace_on_changes)
        super(Tracker, __self__).__init__(
            'aws-native:location:Tracker',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Tracker':
        """
        Get an existing Tracker resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = TrackerArgs.__new__(TrackerArgs)

        __props__.__dict__["arn"] = None
        __props__.__dict__["create_time"] = None
        __props__.__dict__["description"] = None
        __props__.__dict__["event_bridge_enabled"] = None
        __props__.__dict__["kms_key_enable_geospatial_queries"] = None
        __props__.__dict__["kms_key_id"] = None
        __props__.__dict__["position_filtering"] = None
        __props__.__dict__["pricing_plan"] = None
        __props__.__dict__["pricing_plan_data_source"] = None
        __props__.__dict__["tags"] = None
        __props__.__dict__["tracker_arn"] = None
        __props__.__dict__["tracker_name"] = None
        __props__.__dict__["update_time"] = None
        return Tracker(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> pulumi.Output[str]:
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="eventBridgeEnabled")
    def event_bridge_enabled(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "event_bridge_enabled")

    @property
    @pulumi.getter(name="kmsKeyEnableGeospatialQueries")
    def kms_key_enable_geospatial_queries(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "kms_key_enable_geospatial_queries")

    @property
    @pulumi.getter(name="kmsKeyId")
    def kms_key_id(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "kms_key_id")

    @property
    @pulumi.getter(name="positionFiltering")
    def position_filtering(self) -> pulumi.Output[Optional['TrackerPositionFiltering']]:
        return pulumi.get(self, "position_filtering")

    @property
    @pulumi.getter(name="pricingPlan")
    def pricing_plan(self) -> pulumi.Output[Optional['TrackerPricingPlan']]:
        return pulumi.get(self, "pricing_plan")

    @property
    @pulumi.getter(name="pricingPlanDataSource")
    def pricing_plan_data_source(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "pricing_plan_data_source")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence['_root_outputs.Tag']]]:
        """
        An array of key-value pairs to apply to this resource.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="trackerArn")
    def tracker_arn(self) -> pulumi.Output[str]:
        return pulumi.get(self, "tracker_arn")

    @property
    @pulumi.getter(name="trackerName")
    def tracker_name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "tracker_name")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> pulumi.Output[str]:
        return pulumi.get(self, "update_time")

