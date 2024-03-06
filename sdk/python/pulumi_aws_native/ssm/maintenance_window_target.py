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
from ._inputs import *

__all__ = ['MaintenanceWindowTargetArgs', 'MaintenanceWindowTarget']

@pulumi.input_type
class MaintenanceWindowTargetArgs:
    def __init__(__self__, *,
                 resource_type: pulumi.Input[str],
                 targets: pulumi.Input[Sequence[pulumi.Input['MaintenanceWindowTargetTargetsArgs']]],
                 window_id: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 owner_information: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a MaintenanceWindowTarget resource.
        """
        pulumi.set(__self__, "resource_type", resource_type)
        pulumi.set(__self__, "targets", targets)
        pulumi.set(__self__, "window_id", window_id)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if owner_information is not None:
            pulumi.set(__self__, "owner_information", owner_information)

    @property
    @pulumi.getter(name="resourceType")
    def resource_type(self) -> pulumi.Input[str]:
        return pulumi.get(self, "resource_type")

    @resource_type.setter
    def resource_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "resource_type", value)

    @property
    @pulumi.getter
    def targets(self) -> pulumi.Input[Sequence[pulumi.Input['MaintenanceWindowTargetTargetsArgs']]]:
        return pulumi.get(self, "targets")

    @targets.setter
    def targets(self, value: pulumi.Input[Sequence[pulumi.Input['MaintenanceWindowTargetTargetsArgs']]]):
        pulumi.set(self, "targets", value)

    @property
    @pulumi.getter(name="windowId")
    def window_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "window_id")

    @window_id.setter
    def window_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "window_id", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="ownerInformation")
    def owner_information(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "owner_information")

    @owner_information.setter
    def owner_information(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "owner_information", value)


warnings.warn("""MaintenanceWindowTarget is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.""", DeprecationWarning)


class MaintenanceWindowTarget(pulumi.CustomResource):
    warnings.warn("""MaintenanceWindowTarget is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.""", DeprecationWarning)

    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 owner_information: Optional[pulumi.Input[str]] = None,
                 resource_type: Optional[pulumi.Input[str]] = None,
                 targets: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['MaintenanceWindowTargetTargetsArgs']]]]] = None,
                 window_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Resource Type definition for AWS::SSM::MaintenanceWindowTarget

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: MaintenanceWindowTargetArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource Type definition for AWS::SSM::MaintenanceWindowTarget

        :param str resource_name: The name of the resource.
        :param MaintenanceWindowTargetArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(MaintenanceWindowTargetArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 owner_information: Optional[pulumi.Input[str]] = None,
                 resource_type: Optional[pulumi.Input[str]] = None,
                 targets: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['MaintenanceWindowTargetTargetsArgs']]]]] = None,
                 window_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        pulumi.log.warn("""MaintenanceWindowTarget is deprecated: MaintenanceWindowTarget is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.""")
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = MaintenanceWindowTargetArgs.__new__(MaintenanceWindowTargetArgs)

            __props__.__dict__["description"] = description
            __props__.__dict__["name"] = name
            __props__.__dict__["owner_information"] = owner_information
            if resource_type is None and not opts.urn:
                raise TypeError("Missing required property 'resource_type'")
            __props__.__dict__["resource_type"] = resource_type
            if targets is None and not opts.urn:
                raise TypeError("Missing required property 'targets'")
            __props__.__dict__["targets"] = targets
            if window_id is None and not opts.urn:
                raise TypeError("Missing required property 'window_id'")
            __props__.__dict__["window_id"] = window_id
            __props__.__dict__["aws_id"] = None
        replace_on_changes = pulumi.ResourceOptions(replace_on_changes=["window_id"])
        opts = pulumi.ResourceOptions.merge(opts, replace_on_changes)
        super(MaintenanceWindowTarget, __self__).__init__(
            'aws-native:ssm:MaintenanceWindowTarget',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'MaintenanceWindowTarget':
        """
        Get an existing MaintenanceWindowTarget resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = MaintenanceWindowTargetArgs.__new__(MaintenanceWindowTargetArgs)

        __props__.__dict__["aws_id"] = None
        __props__.__dict__["description"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["owner_information"] = None
        __props__.__dict__["resource_type"] = None
        __props__.__dict__["targets"] = None
        __props__.__dict__["window_id"] = None
        return MaintenanceWindowTarget(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="awsId")
    def aws_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "aws_id")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="ownerInformation")
    def owner_information(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "owner_information")

    @property
    @pulumi.getter(name="resourceType")
    def resource_type(self) -> pulumi.Output[str]:
        return pulumi.get(self, "resource_type")

    @property
    @pulumi.getter
    def targets(self) -> pulumi.Output[Sequence['outputs.MaintenanceWindowTargetTargets']]:
        return pulumi.get(self, "targets")

    @property
    @pulumi.getter(name="windowId")
    def window_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "window_id")

