# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['MacroArgs', 'Macro']

@pulumi.input_type
class MacroArgs:
    def __init__(__self__, *,
                 function_name: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 log_group_name: Optional[pulumi.Input[str]] = None,
                 log_role_arn: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Macro resource.
        """
        pulumi.set(__self__, "function_name", function_name)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if log_group_name is not None:
            pulumi.set(__self__, "log_group_name", log_group_name)
        if log_role_arn is not None:
            pulumi.set(__self__, "log_role_arn", log_role_arn)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="functionName")
    def function_name(self) -> pulumi.Input[str]:
        return pulumi.get(self, "function_name")

    @function_name.setter
    def function_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "function_name", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="logGroupName")
    def log_group_name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "log_group_name")

    @log_group_name.setter
    def log_group_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "log_group_name", value)

    @property
    @pulumi.getter(name="logRoleArn")
    def log_role_arn(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "log_role_arn")

    @log_role_arn.setter
    def log_role_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "log_role_arn", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


warnings.warn("""Macro is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.""", DeprecationWarning)


class Macro(pulumi.CustomResource):
    warnings.warn("""Macro is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.""", DeprecationWarning)

    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 function_name: Optional[pulumi.Input[str]] = None,
                 log_group_name: Optional[pulumi.Input[str]] = None,
                 log_role_arn: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Resource Type definition for AWS::CloudFormation::Macro

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: MacroArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource Type definition for AWS::CloudFormation::Macro

        :param str resource_name: The name of the resource.
        :param MacroArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(MacroArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 function_name: Optional[pulumi.Input[str]] = None,
                 log_group_name: Optional[pulumi.Input[str]] = None,
                 log_role_arn: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        pulumi.log.warn("""Macro is deprecated: Macro is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.""")
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = MacroArgs.__new__(MacroArgs)

            __props__.__dict__["description"] = description
            if function_name is None and not opts.urn:
                raise TypeError("Missing required property 'function_name'")
            __props__.__dict__["function_name"] = function_name
            __props__.__dict__["log_group_name"] = log_group_name
            __props__.__dict__["log_role_arn"] = log_role_arn
            __props__.__dict__["name"] = name
            __props__.__dict__["aws_id"] = None
        replace_on_changes = pulumi.ResourceOptions(replace_on_changes=["name"])
        opts = pulumi.ResourceOptions.merge(opts, replace_on_changes)
        super(Macro, __self__).__init__(
            'aws-native:cloudformation:Macro',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Macro':
        """
        Get an existing Macro resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = MacroArgs.__new__(MacroArgs)

        __props__.__dict__["aws_id"] = None
        __props__.__dict__["description"] = None
        __props__.__dict__["function_name"] = None
        __props__.__dict__["log_group_name"] = None
        __props__.__dict__["log_role_arn"] = None
        __props__.__dict__["name"] = None
        return Macro(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="awsId")
    def aws_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "aws_id")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="functionName")
    def function_name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "function_name")

    @property
    @pulumi.getter(name="logGroupName")
    def log_group_name(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "log_group_name")

    @property
    @pulumi.getter(name="logRoleArn")
    def log_role_arn(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "log_role_arn")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "name")

