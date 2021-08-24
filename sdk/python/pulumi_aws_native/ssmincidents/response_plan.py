# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from .. import _inputs as _root_inputs
from .. import outputs as _root_outputs
from ._inputs import *

__all__ = ['ResponsePlanArgs', 'ResponsePlan']

@pulumi.input_type
class ResponsePlanArgs:
    def __init__(__self__, *,
                 incident_template: pulumi.Input['ResponsePlanIncidentTemplateArgs'],
                 name: pulumi.Input[str],
                 actions: Optional[pulumi.Input[Sequence[pulumi.Input['ResponsePlanActionArgs']]]] = None,
                 chat_channel: Optional[pulumi.Input['ResponsePlanChatChannelArgs']] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 engagements: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]] = None):
        """
        The set of arguments for constructing a ResponsePlan resource.
        :param pulumi.Input['ResponsePlanIncidentTemplateArgs'] incident_template: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssmincidents-responseplan.html#cfn-ssmincidents-responseplan-incidenttemplate
        :param pulumi.Input[str] name: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssmincidents-responseplan.html#cfn-ssmincidents-responseplan-name
        :param pulumi.Input[Sequence[pulumi.Input['ResponsePlanActionArgs']]] actions: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssmincidents-responseplan.html#cfn-ssmincidents-responseplan-actions
        :param pulumi.Input['ResponsePlanChatChannelArgs'] chat_channel: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssmincidents-responseplan.html#cfn-ssmincidents-responseplan-chatchannel
        :param pulumi.Input[str] display_name: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssmincidents-responseplan.html#cfn-ssmincidents-responseplan-displayname
        :param pulumi.Input[Sequence[pulumi.Input[str]]] engagements: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssmincidents-responseplan.html#cfn-ssmincidents-responseplan-engagements
        :param pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]] tags: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssmincidents-responseplan.html#cfn-ssmincidents-responseplan-tags
        """
        pulumi.set(__self__, "incident_template", incident_template)
        pulumi.set(__self__, "name", name)
        if actions is not None:
            pulumi.set(__self__, "actions", actions)
        if chat_channel is not None:
            pulumi.set(__self__, "chat_channel", chat_channel)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if engagements is not None:
            pulumi.set(__self__, "engagements", engagements)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="incidentTemplate")
    def incident_template(self) -> pulumi.Input['ResponsePlanIncidentTemplateArgs']:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssmincidents-responseplan.html#cfn-ssmincidents-responseplan-incidenttemplate
        """
        return pulumi.get(self, "incident_template")

    @incident_template.setter
    def incident_template(self, value: pulumi.Input['ResponsePlanIncidentTemplateArgs']):
        pulumi.set(self, "incident_template", value)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssmincidents-responseplan.html#cfn-ssmincidents-responseplan-name
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def actions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ResponsePlanActionArgs']]]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssmincidents-responseplan.html#cfn-ssmincidents-responseplan-actions
        """
        return pulumi.get(self, "actions")

    @actions.setter
    def actions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ResponsePlanActionArgs']]]]):
        pulumi.set(self, "actions", value)

    @property
    @pulumi.getter(name="chatChannel")
    def chat_channel(self) -> Optional[pulumi.Input['ResponsePlanChatChannelArgs']]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssmincidents-responseplan.html#cfn-ssmincidents-responseplan-chatchannel
        """
        return pulumi.get(self, "chat_channel")

    @chat_channel.setter
    def chat_channel(self, value: Optional[pulumi.Input['ResponsePlanChatChannelArgs']]):
        pulumi.set(self, "chat_channel", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssmincidents-responseplan.html#cfn-ssmincidents-responseplan-displayname
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter
    def engagements(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssmincidents-responseplan.html#cfn-ssmincidents-responseplan-engagements
        """
        return pulumi.get(self, "engagements")

    @engagements.setter
    def engagements(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "engagements", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssmincidents-responseplan.html#cfn-ssmincidents-responseplan-tags
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]]):
        pulumi.set(self, "tags", value)


class ResponsePlan(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 actions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ResponsePlanActionArgs']]]]] = None,
                 chat_channel: Optional[pulumi.Input[pulumi.InputType['ResponsePlanChatChannelArgs']]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 engagements: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 incident_template: Optional[pulumi.Input[pulumi.InputType['ResponsePlanIncidentTemplateArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['_root_inputs.TagArgs']]]]] = None,
                 __props__=None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssmincidents-responseplan.html

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ResponsePlanActionArgs']]]] actions: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssmincidents-responseplan.html#cfn-ssmincidents-responseplan-actions
        :param pulumi.Input[pulumi.InputType['ResponsePlanChatChannelArgs']] chat_channel: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssmincidents-responseplan.html#cfn-ssmincidents-responseplan-chatchannel
        :param pulumi.Input[str] display_name: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssmincidents-responseplan.html#cfn-ssmincidents-responseplan-displayname
        :param pulumi.Input[Sequence[pulumi.Input[str]]] engagements: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssmincidents-responseplan.html#cfn-ssmincidents-responseplan-engagements
        :param pulumi.Input[pulumi.InputType['ResponsePlanIncidentTemplateArgs']] incident_template: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssmincidents-responseplan.html#cfn-ssmincidents-responseplan-incidenttemplate
        :param pulumi.Input[str] name: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssmincidents-responseplan.html#cfn-ssmincidents-responseplan-name
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['_root_inputs.TagArgs']]]] tags: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssmincidents-responseplan.html#cfn-ssmincidents-responseplan-tags
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ResponsePlanArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssmincidents-responseplan.html

        :param str resource_name: The name of the resource.
        :param ResponsePlanArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ResponsePlanArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 actions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ResponsePlanActionArgs']]]]] = None,
                 chat_channel: Optional[pulumi.Input[pulumi.InputType['ResponsePlanChatChannelArgs']]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 engagements: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 incident_template: Optional[pulumi.Input[pulumi.InputType['ResponsePlanIncidentTemplateArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['_root_inputs.TagArgs']]]]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ResponsePlanArgs.__new__(ResponsePlanArgs)

            __props__.__dict__["actions"] = actions
            __props__.__dict__["chat_channel"] = chat_channel
            __props__.__dict__["display_name"] = display_name
            __props__.__dict__["engagements"] = engagements
            if incident_template is None and not opts.urn:
                raise TypeError("Missing required property 'incident_template'")
            __props__.__dict__["incident_template"] = incident_template
            if name is None and not opts.urn:
                raise TypeError("Missing required property 'name'")
            __props__.__dict__["name"] = name
            __props__.__dict__["tags"] = tags
            __props__.__dict__["arn"] = None
        super(ResponsePlan, __self__).__init__(
            'aws-native:ssmincidents:ResponsePlan',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'ResponsePlan':
        """
        Get an existing ResponsePlan resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = ResponsePlanArgs.__new__(ResponsePlanArgs)

        __props__.__dict__["actions"] = None
        __props__.__dict__["arn"] = None
        __props__.__dict__["chat_channel"] = None
        __props__.__dict__["display_name"] = None
        __props__.__dict__["engagements"] = None
        __props__.__dict__["incident_template"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["tags"] = None
        return ResponsePlan(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def actions(self) -> pulumi.Output[Optional[Sequence['outputs.ResponsePlanAction']]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssmincidents-responseplan.html#cfn-ssmincidents-responseplan-actions
        """
        return pulumi.get(self, "actions")

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="chatChannel")
    def chat_channel(self) -> pulumi.Output[Optional['outputs.ResponsePlanChatChannel']]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssmincidents-responseplan.html#cfn-ssmincidents-responseplan-chatchannel
        """
        return pulumi.get(self, "chat_channel")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[Optional[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssmincidents-responseplan.html#cfn-ssmincidents-responseplan-displayname
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def engagements(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssmincidents-responseplan.html#cfn-ssmincidents-responseplan-engagements
        """
        return pulumi.get(self, "engagements")

    @property
    @pulumi.getter(name="incidentTemplate")
    def incident_template(self) -> pulumi.Output['outputs.ResponsePlanIncidentTemplate']:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssmincidents-responseplan.html#cfn-ssmincidents-responseplan-incidenttemplate
        """
        return pulumi.get(self, "incident_template")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssmincidents-responseplan.html#cfn-ssmincidents-responseplan-name
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence['_root_outputs.Tag']]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssmincidents-responseplan.html#cfn-ssmincidents-responseplan-tags
        """
        return pulumi.get(self, "tags")
