# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['EmailTemplateArgs', 'EmailTemplate']

@pulumi.input_type
class EmailTemplateArgs:
    def __init__(__self__, *,
                 subject: pulumi.Input[str],
                 template_name: pulumi.Input[str],
                 default_substitutions: Optional[pulumi.Input[str]] = None,
                 html_part: Optional[pulumi.Input[str]] = None,
                 tags: Optional[Any] = None,
                 template_description: Optional[pulumi.Input[str]] = None,
                 text_part: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a EmailTemplate resource.
        :param Any tags: Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::Pinpoint::EmailTemplate` for more information about the expected schema for this property.
        """
        pulumi.set(__self__, "subject", subject)
        pulumi.set(__self__, "template_name", template_name)
        if default_substitutions is not None:
            pulumi.set(__self__, "default_substitutions", default_substitutions)
        if html_part is not None:
            pulumi.set(__self__, "html_part", html_part)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if template_description is not None:
            pulumi.set(__self__, "template_description", template_description)
        if text_part is not None:
            pulumi.set(__self__, "text_part", text_part)

    @property
    @pulumi.getter
    def subject(self) -> pulumi.Input[str]:
        return pulumi.get(self, "subject")

    @subject.setter
    def subject(self, value: pulumi.Input[str]):
        pulumi.set(self, "subject", value)

    @property
    @pulumi.getter(name="templateName")
    def template_name(self) -> pulumi.Input[str]:
        return pulumi.get(self, "template_name")

    @template_name.setter
    def template_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "template_name", value)

    @property
    @pulumi.getter(name="defaultSubstitutions")
    def default_substitutions(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "default_substitutions")

    @default_substitutions.setter
    def default_substitutions(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "default_substitutions", value)

    @property
    @pulumi.getter(name="htmlPart")
    def html_part(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "html_part")

    @html_part.setter
    def html_part(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "html_part", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[Any]:
        """
        Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::Pinpoint::EmailTemplate` for more information about the expected schema for this property.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[Any]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="templateDescription")
    def template_description(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "template_description")

    @template_description.setter
    def template_description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "template_description", value)

    @property
    @pulumi.getter(name="textPart")
    def text_part(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "text_part")

    @text_part.setter
    def text_part(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "text_part", value)


warnings.warn("""EmailTemplate is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.""", DeprecationWarning)


class EmailTemplate(pulumi.CustomResource):
    warnings.warn("""EmailTemplate is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.""", DeprecationWarning)

    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 default_substitutions: Optional[pulumi.Input[str]] = None,
                 html_part: Optional[pulumi.Input[str]] = None,
                 subject: Optional[pulumi.Input[str]] = None,
                 tags: Optional[Any] = None,
                 template_description: Optional[pulumi.Input[str]] = None,
                 template_name: Optional[pulumi.Input[str]] = None,
                 text_part: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Resource Type definition for AWS::Pinpoint::EmailTemplate

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param Any tags: Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::Pinpoint::EmailTemplate` for more information about the expected schema for this property.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: EmailTemplateArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource Type definition for AWS::Pinpoint::EmailTemplate

        :param str resource_name: The name of the resource.
        :param EmailTemplateArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(EmailTemplateArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 default_substitutions: Optional[pulumi.Input[str]] = None,
                 html_part: Optional[pulumi.Input[str]] = None,
                 subject: Optional[pulumi.Input[str]] = None,
                 tags: Optional[Any] = None,
                 template_description: Optional[pulumi.Input[str]] = None,
                 template_name: Optional[pulumi.Input[str]] = None,
                 text_part: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        pulumi.log.warn("""EmailTemplate is deprecated: EmailTemplate is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.""")
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = EmailTemplateArgs.__new__(EmailTemplateArgs)

            __props__.__dict__["default_substitutions"] = default_substitutions
            __props__.__dict__["html_part"] = html_part
            if subject is None and not opts.urn:
                raise TypeError("Missing required property 'subject'")
            __props__.__dict__["subject"] = subject
            __props__.__dict__["tags"] = tags
            __props__.__dict__["template_description"] = template_description
            if template_name is None and not opts.urn:
                raise TypeError("Missing required property 'template_name'")
            __props__.__dict__["template_name"] = template_name
            __props__.__dict__["text_part"] = text_part
            __props__.__dict__["arn"] = None
            __props__.__dict__["aws_id"] = None
        replace_on_changes = pulumi.ResourceOptions(replace_on_changes=["template_name"])
        opts = pulumi.ResourceOptions.merge(opts, replace_on_changes)
        super(EmailTemplate, __self__).__init__(
            'aws-native:pinpoint:EmailTemplate',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'EmailTemplate':
        """
        Get an existing EmailTemplate resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = EmailTemplateArgs.__new__(EmailTemplateArgs)

        __props__.__dict__["arn"] = None
        __props__.__dict__["aws_id"] = None
        __props__.__dict__["default_substitutions"] = None
        __props__.__dict__["html_part"] = None
        __props__.__dict__["subject"] = None
        __props__.__dict__["tags"] = None
        __props__.__dict__["template_description"] = None
        __props__.__dict__["template_name"] = None
        __props__.__dict__["text_part"] = None
        return EmailTemplate(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="awsId")
    def aws_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "aws_id")

    @property
    @pulumi.getter(name="defaultSubstitutions")
    def default_substitutions(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "default_substitutions")

    @property
    @pulumi.getter(name="htmlPart")
    def html_part(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "html_part")

    @property
    @pulumi.getter
    def subject(self) -> pulumi.Output[str]:
        return pulumi.get(self, "subject")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Any]]:
        """
        Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::Pinpoint::EmailTemplate` for more information about the expected schema for this property.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="templateDescription")
    def template_description(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "template_description")

    @property
    @pulumi.getter(name="templateName")
    def template_name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "template_name")

    @property
    @pulumi.getter(name="textPart")
    def text_part(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "text_part")

