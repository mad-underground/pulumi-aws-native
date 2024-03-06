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

__all__ = ['EventBusPolicyArgs', 'EventBusPolicy']

@pulumi.input_type
class EventBusPolicyArgs:
    def __init__(__self__, *,
                 statement_id: pulumi.Input[str],
                 action: Optional[pulumi.Input[str]] = None,
                 condition: Optional[pulumi.Input['EventBusPolicyConditionArgs']] = None,
                 event_bus_name: Optional[pulumi.Input[str]] = None,
                 principal: Optional[pulumi.Input[str]] = None,
                 statement: Optional[Any] = None):
        """
        The set of arguments for constructing a EventBusPolicy resource.
        :param Any statement: Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::Events::EventBusPolicy` for more information about the expected schema for this property.
        """
        pulumi.set(__self__, "statement_id", statement_id)
        if action is not None:
            pulumi.set(__self__, "action", action)
        if condition is not None:
            pulumi.set(__self__, "condition", condition)
        if event_bus_name is not None:
            pulumi.set(__self__, "event_bus_name", event_bus_name)
        if principal is not None:
            pulumi.set(__self__, "principal", principal)
        if statement is not None:
            pulumi.set(__self__, "statement", statement)

    @property
    @pulumi.getter(name="statementId")
    def statement_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "statement_id")

    @statement_id.setter
    def statement_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "statement_id", value)

    @property
    @pulumi.getter
    def action(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "action")

    @action.setter
    def action(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "action", value)

    @property
    @pulumi.getter
    def condition(self) -> Optional[pulumi.Input['EventBusPolicyConditionArgs']]:
        return pulumi.get(self, "condition")

    @condition.setter
    def condition(self, value: Optional[pulumi.Input['EventBusPolicyConditionArgs']]):
        pulumi.set(self, "condition", value)

    @property
    @pulumi.getter(name="eventBusName")
    def event_bus_name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "event_bus_name")

    @event_bus_name.setter
    def event_bus_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "event_bus_name", value)

    @property
    @pulumi.getter
    def principal(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "principal")

    @principal.setter
    def principal(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "principal", value)

    @property
    @pulumi.getter
    def statement(self) -> Optional[Any]:
        """
        Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::Events::EventBusPolicy` for more information about the expected schema for this property.
        """
        return pulumi.get(self, "statement")

    @statement.setter
    def statement(self, value: Optional[Any]):
        pulumi.set(self, "statement", value)


warnings.warn("""EventBusPolicy is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.""", DeprecationWarning)


class EventBusPolicy(pulumi.CustomResource):
    warnings.warn("""EventBusPolicy is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.""", DeprecationWarning)

    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 action: Optional[pulumi.Input[str]] = None,
                 condition: Optional[pulumi.Input[pulumi.InputType['EventBusPolicyConditionArgs']]] = None,
                 event_bus_name: Optional[pulumi.Input[str]] = None,
                 principal: Optional[pulumi.Input[str]] = None,
                 statement: Optional[Any] = None,
                 statement_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Resource Type definition for AWS::Events::EventBusPolicy

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param Any statement: Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::Events::EventBusPolicy` for more information about the expected schema for this property.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: EventBusPolicyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource Type definition for AWS::Events::EventBusPolicy

        :param str resource_name: The name of the resource.
        :param EventBusPolicyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(EventBusPolicyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 action: Optional[pulumi.Input[str]] = None,
                 condition: Optional[pulumi.Input[pulumi.InputType['EventBusPolicyConditionArgs']]] = None,
                 event_bus_name: Optional[pulumi.Input[str]] = None,
                 principal: Optional[pulumi.Input[str]] = None,
                 statement: Optional[Any] = None,
                 statement_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        pulumi.log.warn("""EventBusPolicy is deprecated: EventBusPolicy is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.""")
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = EventBusPolicyArgs.__new__(EventBusPolicyArgs)

            __props__.__dict__["action"] = action
            __props__.__dict__["condition"] = condition
            __props__.__dict__["event_bus_name"] = event_bus_name
            __props__.__dict__["principal"] = principal
            __props__.__dict__["statement"] = statement
            if statement_id is None and not opts.urn:
                raise TypeError("Missing required property 'statement_id'")
            __props__.__dict__["statement_id"] = statement_id
            __props__.__dict__["aws_id"] = None
        replace_on_changes = pulumi.ResourceOptions(replace_on_changes=["event_bus_name", "statement_id"])
        opts = pulumi.ResourceOptions.merge(opts, replace_on_changes)
        super(EventBusPolicy, __self__).__init__(
            'aws-native:events:EventBusPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'EventBusPolicy':
        """
        Get an existing EventBusPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = EventBusPolicyArgs.__new__(EventBusPolicyArgs)

        __props__.__dict__["action"] = None
        __props__.__dict__["aws_id"] = None
        __props__.__dict__["condition"] = None
        __props__.__dict__["event_bus_name"] = None
        __props__.__dict__["principal"] = None
        __props__.__dict__["statement"] = None
        __props__.__dict__["statement_id"] = None
        return EventBusPolicy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def action(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "action")

    @property
    @pulumi.getter(name="awsId")
    def aws_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "aws_id")

    @property
    @pulumi.getter
    def condition(self) -> pulumi.Output[Optional['outputs.EventBusPolicyCondition']]:
        return pulumi.get(self, "condition")

    @property
    @pulumi.getter(name="eventBusName")
    def event_bus_name(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "event_bus_name")

    @property
    @pulumi.getter
    def principal(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "principal")

    @property
    @pulumi.getter
    def statement(self) -> pulumi.Output[Optional[Any]]:
        """
        Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::Events::EventBusPolicy` for more information about the expected schema for this property.
        """
        return pulumi.get(self, "statement")

    @property
    @pulumi.getter(name="statementId")
    def statement_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "statement_id")

