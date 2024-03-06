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
from .. import _inputs as _root_inputs
from .. import outputs as _root_outputs
from ._enums import *
from ._inputs import *

__all__ = ['FindingsFilterArgs', 'FindingsFilter']

@pulumi.input_type
class FindingsFilterArgs:
    def __init__(__self__, *,
                 finding_criteria: pulumi.Input['FindingsFilterFindingCriteriaArgs'],
                 action: Optional[pulumi.Input['FindingsFilterFindingFilterAction']] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 position: Optional[pulumi.Input[int]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]] = None):
        """
        The set of arguments for constructing a FindingsFilter resource.
        :param pulumi.Input['FindingsFilterFindingCriteriaArgs'] finding_criteria: Findings filter criteria.
        :param pulumi.Input['FindingsFilterFindingFilterAction'] action: Findings filter action.
        :param pulumi.Input[str] description: Findings filter description
        :param pulumi.Input[str] name: Findings filter name
        :param pulumi.Input[int] position: Findings filter position.
        :param pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]] tags: A collection of tags associated with a resource
        """
        pulumi.set(__self__, "finding_criteria", finding_criteria)
        if action is not None:
            pulumi.set(__self__, "action", action)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if position is not None:
            pulumi.set(__self__, "position", position)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="findingCriteria")
    def finding_criteria(self) -> pulumi.Input['FindingsFilterFindingCriteriaArgs']:
        """
        Findings filter criteria.
        """
        return pulumi.get(self, "finding_criteria")

    @finding_criteria.setter
    def finding_criteria(self, value: pulumi.Input['FindingsFilterFindingCriteriaArgs']):
        pulumi.set(self, "finding_criteria", value)

    @property
    @pulumi.getter
    def action(self) -> Optional[pulumi.Input['FindingsFilterFindingFilterAction']]:
        """
        Findings filter action.
        """
        return pulumi.get(self, "action")

    @action.setter
    def action(self, value: Optional[pulumi.Input['FindingsFilterFindingFilterAction']]):
        pulumi.set(self, "action", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Findings filter description
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Findings filter name
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def position(self) -> Optional[pulumi.Input[int]]:
        """
        Findings filter position.
        """
        return pulumi.get(self, "position")

    @position.setter
    def position(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "position", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]]:
        """
        A collection of tags associated with a resource
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]]):
        pulumi.set(self, "tags", value)


class FindingsFilter(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 action: Optional[pulumi.Input['FindingsFilterFindingFilterAction']] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 finding_criteria: Optional[pulumi.Input[pulumi.InputType['FindingsFilterFindingCriteriaArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 position: Optional[pulumi.Input[int]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['_root_inputs.TagArgs']]]]] = None,
                 __props__=None):
        """
        Macie FindingsFilter resource schema.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input['FindingsFilterFindingFilterAction'] action: Findings filter action.
        :param pulumi.Input[str] description: Findings filter description
        :param pulumi.Input[pulumi.InputType['FindingsFilterFindingCriteriaArgs']] finding_criteria: Findings filter criteria.
        :param pulumi.Input[str] name: Findings filter name
        :param pulumi.Input[int] position: Findings filter position.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['_root_inputs.TagArgs']]]] tags: A collection of tags associated with a resource
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: FindingsFilterArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Macie FindingsFilter resource schema.

        :param str resource_name: The name of the resource.
        :param FindingsFilterArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FindingsFilterArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 action: Optional[pulumi.Input['FindingsFilterFindingFilterAction']] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 finding_criteria: Optional[pulumi.Input[pulumi.InputType['FindingsFilterFindingCriteriaArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 position: Optional[pulumi.Input[int]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['_root_inputs.TagArgs']]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = FindingsFilterArgs.__new__(FindingsFilterArgs)

            __props__.__dict__["action"] = action
            __props__.__dict__["description"] = description
            if finding_criteria is None and not opts.urn:
                raise TypeError("Missing required property 'finding_criteria'")
            __props__.__dict__["finding_criteria"] = finding_criteria
            __props__.__dict__["name"] = name
            __props__.__dict__["position"] = position
            __props__.__dict__["tags"] = tags
            __props__.__dict__["arn"] = None
            __props__.__dict__["aws_id"] = None
        super(FindingsFilter, __self__).__init__(
            'aws-native:macie:FindingsFilter',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'FindingsFilter':
        """
        Get an existing FindingsFilter resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = FindingsFilterArgs.__new__(FindingsFilterArgs)

        __props__.__dict__["action"] = None
        __props__.__dict__["arn"] = None
        __props__.__dict__["aws_id"] = None
        __props__.__dict__["description"] = None
        __props__.__dict__["finding_criteria"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["position"] = None
        __props__.__dict__["tags"] = None
        return FindingsFilter(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def action(self) -> pulumi.Output[Optional['FindingsFilterFindingFilterAction']]:
        """
        Findings filter action.
        """
        return pulumi.get(self, "action")

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        Findings filter ARN.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="awsId")
    def aws_id(self) -> pulumi.Output[str]:
        """
        Findings filter ID.
        """
        return pulumi.get(self, "aws_id")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        Findings filter description
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="findingCriteria")
    def finding_criteria(self) -> pulumi.Output['outputs.FindingsFilterFindingCriteria']:
        """
        Findings filter criteria.
        """
        return pulumi.get(self, "finding_criteria")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Findings filter name
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def position(self) -> pulumi.Output[Optional[int]]:
        """
        Findings filter position.
        """
        return pulumi.get(self, "position")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence['_root_outputs.Tag']]]:
        """
        A collection of tags associated with a resource
        """
        return pulumi.get(self, "tags")

