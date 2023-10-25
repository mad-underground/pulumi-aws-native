# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._enums import *
from ._inputs import *

__all__ = ['TemplateArgs', 'Template']

@pulumi.input_type
class TemplateArgs:
    def __init__(__self__, *,
                 aws_account_id: pulumi.Input[str],
                 template_id: pulumi.Input[str],
                 definition: Optional[pulumi.Input['TemplateVersionDefinitionArgs']] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 permissions: Optional[pulumi.Input[Sequence[pulumi.Input['TemplateResourcePermissionArgs']]]] = None,
                 source_entity: Optional[pulumi.Input['TemplateSourceEntityArgs']] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['TemplateTagArgs']]]] = None,
                 validation_strategy: Optional[pulumi.Input['TemplateValidationStrategyArgs']] = None,
                 version_description: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Template resource.
        """
        TemplateArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            aws_account_id=aws_account_id,
            template_id=template_id,
            definition=definition,
            name=name,
            permissions=permissions,
            source_entity=source_entity,
            tags=tags,
            validation_strategy=validation_strategy,
            version_description=version_description,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             aws_account_id: pulumi.Input[str],
             template_id: pulumi.Input[str],
             definition: Optional[pulumi.Input['TemplateVersionDefinitionArgs']] = None,
             name: Optional[pulumi.Input[str]] = None,
             permissions: Optional[pulumi.Input[Sequence[pulumi.Input['TemplateResourcePermissionArgs']]]] = None,
             source_entity: Optional[pulumi.Input['TemplateSourceEntityArgs']] = None,
             tags: Optional[pulumi.Input[Sequence[pulumi.Input['TemplateTagArgs']]]] = None,
             validation_strategy: Optional[pulumi.Input['TemplateValidationStrategyArgs']] = None,
             version_description: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions]=None):
        _setter("aws_account_id", aws_account_id)
        _setter("template_id", template_id)
        if definition is not None:
            _setter("definition", definition)
        if name is not None:
            _setter("name", name)
        if permissions is not None:
            _setter("permissions", permissions)
        if source_entity is not None:
            _setter("source_entity", source_entity)
        if tags is not None:
            _setter("tags", tags)
        if validation_strategy is not None:
            _setter("validation_strategy", validation_strategy)
        if version_description is not None:
            _setter("version_description", version_description)

    @property
    @pulumi.getter(name="awsAccountId")
    def aws_account_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "aws_account_id")

    @aws_account_id.setter
    def aws_account_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "aws_account_id", value)

    @property
    @pulumi.getter(name="templateId")
    def template_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "template_id")

    @template_id.setter
    def template_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "template_id", value)

    @property
    @pulumi.getter
    def definition(self) -> Optional[pulumi.Input['TemplateVersionDefinitionArgs']]:
        return pulumi.get(self, "definition")

    @definition.setter
    def definition(self, value: Optional[pulumi.Input['TemplateVersionDefinitionArgs']]):
        pulumi.set(self, "definition", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def permissions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['TemplateResourcePermissionArgs']]]]:
        return pulumi.get(self, "permissions")

    @permissions.setter
    def permissions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['TemplateResourcePermissionArgs']]]]):
        pulumi.set(self, "permissions", value)

    @property
    @pulumi.getter(name="sourceEntity")
    def source_entity(self) -> Optional[pulumi.Input['TemplateSourceEntityArgs']]:
        return pulumi.get(self, "source_entity")

    @source_entity.setter
    def source_entity(self, value: Optional[pulumi.Input['TemplateSourceEntityArgs']]):
        pulumi.set(self, "source_entity", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['TemplateTagArgs']]]]:
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['TemplateTagArgs']]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="validationStrategy")
    def validation_strategy(self) -> Optional[pulumi.Input['TemplateValidationStrategyArgs']]:
        return pulumi.get(self, "validation_strategy")

    @validation_strategy.setter
    def validation_strategy(self, value: Optional[pulumi.Input['TemplateValidationStrategyArgs']]):
        pulumi.set(self, "validation_strategy", value)

    @property
    @pulumi.getter(name="versionDescription")
    def version_description(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "version_description")

    @version_description.setter
    def version_description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "version_description", value)


class Template(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 aws_account_id: Optional[pulumi.Input[str]] = None,
                 definition: Optional[pulumi.Input[pulumi.InputType['TemplateVersionDefinitionArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 permissions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TemplateResourcePermissionArgs']]]]] = None,
                 source_entity: Optional[pulumi.Input[pulumi.InputType['TemplateSourceEntityArgs']]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TemplateTagArgs']]]]] = None,
                 template_id: Optional[pulumi.Input[str]] = None,
                 validation_strategy: Optional[pulumi.Input[pulumi.InputType['TemplateValidationStrategyArgs']]] = None,
                 version_description: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Definition of the AWS::QuickSight::Template Resource Type.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: TemplateArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Definition of the AWS::QuickSight::Template Resource Type.

        :param str resource_name: The name of the resource.
        :param TemplateArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(TemplateArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            TemplateArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 aws_account_id: Optional[pulumi.Input[str]] = None,
                 definition: Optional[pulumi.Input[pulumi.InputType['TemplateVersionDefinitionArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 permissions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TemplateResourcePermissionArgs']]]]] = None,
                 source_entity: Optional[pulumi.Input[pulumi.InputType['TemplateSourceEntityArgs']]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TemplateTagArgs']]]]] = None,
                 template_id: Optional[pulumi.Input[str]] = None,
                 validation_strategy: Optional[pulumi.Input[pulumi.InputType['TemplateValidationStrategyArgs']]] = None,
                 version_description: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = TemplateArgs.__new__(TemplateArgs)

            if aws_account_id is None and not opts.urn:
                raise TypeError("Missing required property 'aws_account_id'")
            __props__.__dict__["aws_account_id"] = aws_account_id
            if definition is not None and not isinstance(definition, TemplateVersionDefinitionArgs):
                definition = definition or {}
                def _setter(key, value):
                    definition[key] = value
                TemplateVersionDefinitionArgs._configure(_setter, **definition)
            __props__.__dict__["definition"] = definition
            __props__.__dict__["name"] = name
            __props__.__dict__["permissions"] = permissions
            if source_entity is not None and not isinstance(source_entity, TemplateSourceEntityArgs):
                source_entity = source_entity or {}
                def _setter(key, value):
                    source_entity[key] = value
                TemplateSourceEntityArgs._configure(_setter, **source_entity)
            __props__.__dict__["source_entity"] = source_entity
            __props__.__dict__["tags"] = tags
            if template_id is None and not opts.urn:
                raise TypeError("Missing required property 'template_id'")
            __props__.__dict__["template_id"] = template_id
            if validation_strategy is not None and not isinstance(validation_strategy, TemplateValidationStrategyArgs):
                validation_strategy = validation_strategy or {}
                def _setter(key, value):
                    validation_strategy[key] = value
                TemplateValidationStrategyArgs._configure(_setter, **validation_strategy)
            __props__.__dict__["validation_strategy"] = validation_strategy
            __props__.__dict__["version_description"] = version_description
            __props__.__dict__["arn"] = None
            __props__.__dict__["created_time"] = None
            __props__.__dict__["last_updated_time"] = None
            __props__.__dict__["version"] = None
        replace_on_changes = pulumi.ResourceOptions(replace_on_changes=["aws_account_id", "template_id"])
        opts = pulumi.ResourceOptions.merge(opts, replace_on_changes)
        super(Template, __self__).__init__(
            'aws-native:quicksight:Template',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Template':
        """
        Get an existing Template resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = TemplateArgs.__new__(TemplateArgs)

        __props__.__dict__["arn"] = None
        __props__.__dict__["aws_account_id"] = None
        __props__.__dict__["created_time"] = None
        __props__.__dict__["definition"] = None
        __props__.__dict__["last_updated_time"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["permissions"] = None
        __props__.__dict__["source_entity"] = None
        __props__.__dict__["tags"] = None
        __props__.__dict__["template_id"] = None
        __props__.__dict__["validation_strategy"] = None
        __props__.__dict__["version"] = None
        __props__.__dict__["version_description"] = None
        return Template(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="awsAccountId")
    def aws_account_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "aws_account_id")

    @property
    @pulumi.getter(name="createdTime")
    def created_time(self) -> pulumi.Output[str]:
        return pulumi.get(self, "created_time")

    @property
    @pulumi.getter
    def definition(self) -> pulumi.Output[Optional['outputs.TemplateVersionDefinition']]:
        return pulumi.get(self, "definition")

    @property
    @pulumi.getter(name="lastUpdatedTime")
    def last_updated_time(self) -> pulumi.Output[str]:
        return pulumi.get(self, "last_updated_time")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def permissions(self) -> pulumi.Output[Optional[Sequence['outputs.TemplateResourcePermission']]]:
        return pulumi.get(self, "permissions")

    @property
    @pulumi.getter(name="sourceEntity")
    def source_entity(self) -> pulumi.Output[Optional['outputs.TemplateSourceEntity']]:
        return pulumi.get(self, "source_entity")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence['outputs.TemplateTag']]]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="templateId")
    def template_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "template_id")

    @property
    @pulumi.getter(name="validationStrategy")
    def validation_strategy(self) -> pulumi.Output[Optional['outputs.TemplateValidationStrategy']]:
        return pulumi.get(self, "validation_strategy")

    @property
    @pulumi.getter
    def version(self) -> pulumi.Output['outputs.TemplateVersion']:
        return pulumi.get(self, "version")

    @property
    @pulumi.getter(name="versionDescription")
    def version_description(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "version_description")

