# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['FleetArgs', 'Fleet']

@pulumi.input_type
class FleetArgs:
    def __init__(__self__, *,
                 compute_capacity: pulumi.Input['FleetComputeCapacityArgs'],
                 instance_type: pulumi.Input[str],
                 name: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 disconnect_timeout_in_seconds: Optional[pulumi.Input[int]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 domain_join_info: Optional[pulumi.Input['FleetDomainJoinInfoArgs']] = None,
                 enable_default_internet_access: Optional[pulumi.Input[bool]] = None,
                 fleet_type: Optional[pulumi.Input[str]] = None,
                 iam_role_arn: Optional[pulumi.Input[str]] = None,
                 idle_disconnect_timeout_in_seconds: Optional[pulumi.Input[int]] = None,
                 image_arn: Optional[pulumi.Input[str]] = None,
                 image_name: Optional[pulumi.Input[str]] = None,
                 max_user_duration_in_seconds: Optional[pulumi.Input[int]] = None,
                 stream_view: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['FleetTagArgs']]]] = None,
                 vpc_config: Optional[pulumi.Input['FleetVpcConfigArgs']] = None):
        """
        The set of arguments for constructing a Fleet resource.
        """
        pulumi.set(__self__, "compute_capacity", compute_capacity)
        pulumi.set(__self__, "instance_type", instance_type)
        pulumi.set(__self__, "name", name)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if disconnect_timeout_in_seconds is not None:
            pulumi.set(__self__, "disconnect_timeout_in_seconds", disconnect_timeout_in_seconds)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if domain_join_info is not None:
            pulumi.set(__self__, "domain_join_info", domain_join_info)
        if enable_default_internet_access is not None:
            pulumi.set(__self__, "enable_default_internet_access", enable_default_internet_access)
        if fleet_type is not None:
            pulumi.set(__self__, "fleet_type", fleet_type)
        if iam_role_arn is not None:
            pulumi.set(__self__, "iam_role_arn", iam_role_arn)
        if idle_disconnect_timeout_in_seconds is not None:
            pulumi.set(__self__, "idle_disconnect_timeout_in_seconds", idle_disconnect_timeout_in_seconds)
        if image_arn is not None:
            pulumi.set(__self__, "image_arn", image_arn)
        if image_name is not None:
            pulumi.set(__self__, "image_name", image_name)
        if max_user_duration_in_seconds is not None:
            pulumi.set(__self__, "max_user_duration_in_seconds", max_user_duration_in_seconds)
        if stream_view is not None:
            pulumi.set(__self__, "stream_view", stream_view)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if vpc_config is not None:
            pulumi.set(__self__, "vpc_config", vpc_config)

    @property
    @pulumi.getter(name="computeCapacity")
    def compute_capacity(self) -> pulumi.Input['FleetComputeCapacityArgs']:
        return pulumi.get(self, "compute_capacity")

    @compute_capacity.setter
    def compute_capacity(self, value: pulumi.Input['FleetComputeCapacityArgs']):
        pulumi.set(self, "compute_capacity", value)

    @property
    @pulumi.getter(name="instanceType")
    def instance_type(self) -> pulumi.Input[str]:
        return pulumi.get(self, "instance_type")

    @instance_type.setter
    def instance_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_type", value)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="disconnectTimeoutInSeconds")
    def disconnect_timeout_in_seconds(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "disconnect_timeout_in_seconds")

    @disconnect_timeout_in_seconds.setter
    def disconnect_timeout_in_seconds(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "disconnect_timeout_in_seconds", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter(name="domainJoinInfo")
    def domain_join_info(self) -> Optional[pulumi.Input['FleetDomainJoinInfoArgs']]:
        return pulumi.get(self, "domain_join_info")

    @domain_join_info.setter
    def domain_join_info(self, value: Optional[pulumi.Input['FleetDomainJoinInfoArgs']]):
        pulumi.set(self, "domain_join_info", value)

    @property
    @pulumi.getter(name="enableDefaultInternetAccess")
    def enable_default_internet_access(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "enable_default_internet_access")

    @enable_default_internet_access.setter
    def enable_default_internet_access(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_default_internet_access", value)

    @property
    @pulumi.getter(name="fleetType")
    def fleet_type(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "fleet_type")

    @fleet_type.setter
    def fleet_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "fleet_type", value)

    @property
    @pulumi.getter(name="iamRoleArn")
    def iam_role_arn(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "iam_role_arn")

    @iam_role_arn.setter
    def iam_role_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "iam_role_arn", value)

    @property
    @pulumi.getter(name="idleDisconnectTimeoutInSeconds")
    def idle_disconnect_timeout_in_seconds(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "idle_disconnect_timeout_in_seconds")

    @idle_disconnect_timeout_in_seconds.setter
    def idle_disconnect_timeout_in_seconds(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "idle_disconnect_timeout_in_seconds", value)

    @property
    @pulumi.getter(name="imageArn")
    def image_arn(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "image_arn")

    @image_arn.setter
    def image_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "image_arn", value)

    @property
    @pulumi.getter(name="imageName")
    def image_name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "image_name")

    @image_name.setter
    def image_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "image_name", value)

    @property
    @pulumi.getter(name="maxUserDurationInSeconds")
    def max_user_duration_in_seconds(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "max_user_duration_in_seconds")

    @max_user_duration_in_seconds.setter
    def max_user_duration_in_seconds(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_user_duration_in_seconds", value)

    @property
    @pulumi.getter(name="streamView")
    def stream_view(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "stream_view")

    @stream_view.setter
    def stream_view(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "stream_view", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['FleetTagArgs']]]]:
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['FleetTagArgs']]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="vpcConfig")
    def vpc_config(self) -> Optional[pulumi.Input['FleetVpcConfigArgs']]:
        return pulumi.get(self, "vpc_config")

    @vpc_config.setter
    def vpc_config(self, value: Optional[pulumi.Input['FleetVpcConfigArgs']]):
        pulumi.set(self, "vpc_config", value)


warnings.warn("""Fleet is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.""", DeprecationWarning)


class Fleet(pulumi.CustomResource):
    warnings.warn("""Fleet is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.""", DeprecationWarning)

    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 compute_capacity: Optional[pulumi.Input[pulumi.InputType['FleetComputeCapacityArgs']]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 disconnect_timeout_in_seconds: Optional[pulumi.Input[int]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 domain_join_info: Optional[pulumi.Input[pulumi.InputType['FleetDomainJoinInfoArgs']]] = None,
                 enable_default_internet_access: Optional[pulumi.Input[bool]] = None,
                 fleet_type: Optional[pulumi.Input[str]] = None,
                 iam_role_arn: Optional[pulumi.Input[str]] = None,
                 idle_disconnect_timeout_in_seconds: Optional[pulumi.Input[int]] = None,
                 image_arn: Optional[pulumi.Input[str]] = None,
                 image_name: Optional[pulumi.Input[str]] = None,
                 instance_type: Optional[pulumi.Input[str]] = None,
                 max_user_duration_in_seconds: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 stream_view: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FleetTagArgs']]]]] = None,
                 vpc_config: Optional[pulumi.Input[pulumi.InputType['FleetVpcConfigArgs']]] = None,
                 __props__=None):
        """
        Resource Type definition for AWS::AppStream::Fleet

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: FleetArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource Type definition for AWS::AppStream::Fleet

        :param str resource_name: The name of the resource.
        :param FleetArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FleetArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 compute_capacity: Optional[pulumi.Input[pulumi.InputType['FleetComputeCapacityArgs']]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 disconnect_timeout_in_seconds: Optional[pulumi.Input[int]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 domain_join_info: Optional[pulumi.Input[pulumi.InputType['FleetDomainJoinInfoArgs']]] = None,
                 enable_default_internet_access: Optional[pulumi.Input[bool]] = None,
                 fleet_type: Optional[pulumi.Input[str]] = None,
                 iam_role_arn: Optional[pulumi.Input[str]] = None,
                 idle_disconnect_timeout_in_seconds: Optional[pulumi.Input[int]] = None,
                 image_arn: Optional[pulumi.Input[str]] = None,
                 image_name: Optional[pulumi.Input[str]] = None,
                 instance_type: Optional[pulumi.Input[str]] = None,
                 max_user_duration_in_seconds: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 stream_view: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FleetTagArgs']]]]] = None,
                 vpc_config: Optional[pulumi.Input[pulumi.InputType['FleetVpcConfigArgs']]] = None,
                 __props__=None):
        pulumi.log.warn("""Fleet is deprecated: Fleet is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.""")
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = FleetArgs.__new__(FleetArgs)

            if compute_capacity is None and not opts.urn:
                raise TypeError("Missing required property 'compute_capacity'")
            __props__.__dict__["compute_capacity"] = compute_capacity
            __props__.__dict__["description"] = description
            __props__.__dict__["disconnect_timeout_in_seconds"] = disconnect_timeout_in_seconds
            __props__.__dict__["display_name"] = display_name
            __props__.__dict__["domain_join_info"] = domain_join_info
            __props__.__dict__["enable_default_internet_access"] = enable_default_internet_access
            __props__.__dict__["fleet_type"] = fleet_type
            __props__.__dict__["iam_role_arn"] = iam_role_arn
            __props__.__dict__["idle_disconnect_timeout_in_seconds"] = idle_disconnect_timeout_in_seconds
            __props__.__dict__["image_arn"] = image_arn
            __props__.__dict__["image_name"] = image_name
            if instance_type is None and not opts.urn:
                raise TypeError("Missing required property 'instance_type'")
            __props__.__dict__["instance_type"] = instance_type
            __props__.__dict__["max_user_duration_in_seconds"] = max_user_duration_in_seconds
            if name is None and not opts.urn:
                raise TypeError("Missing required property 'name'")
            __props__.__dict__["name"] = name
            __props__.__dict__["stream_view"] = stream_view
            __props__.__dict__["tags"] = tags
            __props__.__dict__["vpc_config"] = vpc_config
        super(Fleet, __self__).__init__(
            'aws-native:appstream:Fleet',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Fleet':
        """
        Get an existing Fleet resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = FleetArgs.__new__(FleetArgs)

        __props__.__dict__["compute_capacity"] = None
        __props__.__dict__["description"] = None
        __props__.__dict__["disconnect_timeout_in_seconds"] = None
        __props__.__dict__["display_name"] = None
        __props__.__dict__["domain_join_info"] = None
        __props__.__dict__["enable_default_internet_access"] = None
        __props__.__dict__["fleet_type"] = None
        __props__.__dict__["iam_role_arn"] = None
        __props__.__dict__["idle_disconnect_timeout_in_seconds"] = None
        __props__.__dict__["image_arn"] = None
        __props__.__dict__["image_name"] = None
        __props__.__dict__["instance_type"] = None
        __props__.__dict__["max_user_duration_in_seconds"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["stream_view"] = None
        __props__.__dict__["tags"] = None
        __props__.__dict__["vpc_config"] = None
        return Fleet(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="computeCapacity")
    def compute_capacity(self) -> pulumi.Output['outputs.FleetComputeCapacity']:
        return pulumi.get(self, "compute_capacity")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="disconnectTimeoutInSeconds")
    def disconnect_timeout_in_seconds(self) -> pulumi.Output[Optional[int]]:
        return pulumi.get(self, "disconnect_timeout_in_seconds")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="domainJoinInfo")
    def domain_join_info(self) -> pulumi.Output[Optional['outputs.FleetDomainJoinInfo']]:
        return pulumi.get(self, "domain_join_info")

    @property
    @pulumi.getter(name="enableDefaultInternetAccess")
    def enable_default_internet_access(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "enable_default_internet_access")

    @property
    @pulumi.getter(name="fleetType")
    def fleet_type(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "fleet_type")

    @property
    @pulumi.getter(name="iamRoleArn")
    def iam_role_arn(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "iam_role_arn")

    @property
    @pulumi.getter(name="idleDisconnectTimeoutInSeconds")
    def idle_disconnect_timeout_in_seconds(self) -> pulumi.Output[Optional[int]]:
        return pulumi.get(self, "idle_disconnect_timeout_in_seconds")

    @property
    @pulumi.getter(name="imageArn")
    def image_arn(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "image_arn")

    @property
    @pulumi.getter(name="imageName")
    def image_name(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "image_name")

    @property
    @pulumi.getter(name="instanceType")
    def instance_type(self) -> pulumi.Output[str]:
        return pulumi.get(self, "instance_type")

    @property
    @pulumi.getter(name="maxUserDurationInSeconds")
    def max_user_duration_in_seconds(self) -> pulumi.Output[Optional[int]]:
        return pulumi.get(self, "max_user_duration_in_seconds")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="streamView")
    def stream_view(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "stream_view")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence['outputs.FleetTag']]]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="vpcConfig")
    def vpc_config(self) -> pulumi.Output[Optional['outputs.FleetVpcConfig']]:
        return pulumi.get(self, "vpc_config")
