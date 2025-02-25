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

__all__ = ['DashboardArgs', 'Dashboard']

@pulumi.input_type
class DashboardArgs:
    def __init__(__self__, *,
                 aws_account_id: pulumi.Input[str],
                 dashboard_id: pulumi.Input[str],
                 dashboard_publish_options: Optional[pulumi.Input['DashboardPublishOptionsArgs']] = None,
                 definition: Optional[pulumi.Input['DashboardVersionDefinitionArgs']] = None,
                 link_entities: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 link_sharing_configuration: Optional[pulumi.Input['DashboardLinkSharingConfigurationArgs']] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 parameters: Optional[pulumi.Input['DashboardParametersArgs']] = None,
                 permissions: Optional[pulumi.Input[Sequence[pulumi.Input['DashboardResourcePermissionArgs']]]] = None,
                 source_entity: Optional[pulumi.Input['DashboardSourceEntityArgs']] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]] = None,
                 theme_arn: Optional[pulumi.Input[str]] = None,
                 validation_strategy: Optional[pulumi.Input['DashboardValidationStrategyArgs']] = None,
                 version_description: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Dashboard resource.
        """
        pulumi.set(__self__, "aws_account_id", aws_account_id)
        pulumi.set(__self__, "dashboard_id", dashboard_id)
        if dashboard_publish_options is not None:
            pulumi.set(__self__, "dashboard_publish_options", dashboard_publish_options)
        if definition is not None:
            pulumi.set(__self__, "definition", definition)
        if link_entities is not None:
            pulumi.set(__self__, "link_entities", link_entities)
        if link_sharing_configuration is not None:
            pulumi.set(__self__, "link_sharing_configuration", link_sharing_configuration)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if parameters is not None:
            pulumi.set(__self__, "parameters", parameters)
        if permissions is not None:
            pulumi.set(__self__, "permissions", permissions)
        if source_entity is not None:
            pulumi.set(__self__, "source_entity", source_entity)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if theme_arn is not None:
            pulumi.set(__self__, "theme_arn", theme_arn)
        if validation_strategy is not None:
            pulumi.set(__self__, "validation_strategy", validation_strategy)
        if version_description is not None:
            pulumi.set(__self__, "version_description", version_description)

    @property
    @pulumi.getter(name="awsAccountId")
    def aws_account_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "aws_account_id")

    @aws_account_id.setter
    def aws_account_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "aws_account_id", value)

    @property
    @pulumi.getter(name="dashboardId")
    def dashboard_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "dashboard_id")

    @dashboard_id.setter
    def dashboard_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "dashboard_id", value)

    @property
    @pulumi.getter(name="dashboardPublishOptions")
    def dashboard_publish_options(self) -> Optional[pulumi.Input['DashboardPublishOptionsArgs']]:
        return pulumi.get(self, "dashboard_publish_options")

    @dashboard_publish_options.setter
    def dashboard_publish_options(self, value: Optional[pulumi.Input['DashboardPublishOptionsArgs']]):
        pulumi.set(self, "dashboard_publish_options", value)

    @property
    @pulumi.getter
    def definition(self) -> Optional[pulumi.Input['DashboardVersionDefinitionArgs']]:
        return pulumi.get(self, "definition")

    @definition.setter
    def definition(self, value: Optional[pulumi.Input['DashboardVersionDefinitionArgs']]):
        pulumi.set(self, "definition", value)

    @property
    @pulumi.getter(name="linkEntities")
    def link_entities(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "link_entities")

    @link_entities.setter
    def link_entities(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "link_entities", value)

    @property
    @pulumi.getter(name="linkSharingConfiguration")
    def link_sharing_configuration(self) -> Optional[pulumi.Input['DashboardLinkSharingConfigurationArgs']]:
        return pulumi.get(self, "link_sharing_configuration")

    @link_sharing_configuration.setter
    def link_sharing_configuration(self, value: Optional[pulumi.Input['DashboardLinkSharingConfigurationArgs']]):
        pulumi.set(self, "link_sharing_configuration", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def parameters(self) -> Optional[pulumi.Input['DashboardParametersArgs']]:
        return pulumi.get(self, "parameters")

    @parameters.setter
    def parameters(self, value: Optional[pulumi.Input['DashboardParametersArgs']]):
        pulumi.set(self, "parameters", value)

    @property
    @pulumi.getter
    def permissions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['DashboardResourcePermissionArgs']]]]:
        return pulumi.get(self, "permissions")

    @permissions.setter
    def permissions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['DashboardResourcePermissionArgs']]]]):
        pulumi.set(self, "permissions", value)

    @property
    @pulumi.getter(name="sourceEntity")
    def source_entity(self) -> Optional[pulumi.Input['DashboardSourceEntityArgs']]:
        return pulumi.get(self, "source_entity")

    @source_entity.setter
    def source_entity(self, value: Optional[pulumi.Input['DashboardSourceEntityArgs']]):
        pulumi.set(self, "source_entity", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]]:
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="themeArn")
    def theme_arn(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "theme_arn")

    @theme_arn.setter
    def theme_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "theme_arn", value)

    @property
    @pulumi.getter(name="validationStrategy")
    def validation_strategy(self) -> Optional[pulumi.Input['DashboardValidationStrategyArgs']]:
        return pulumi.get(self, "validation_strategy")

    @validation_strategy.setter
    def validation_strategy(self, value: Optional[pulumi.Input['DashboardValidationStrategyArgs']]):
        pulumi.set(self, "validation_strategy", value)

    @property
    @pulumi.getter(name="versionDescription")
    def version_description(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "version_description")

    @version_description.setter
    def version_description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "version_description", value)


class Dashboard(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 aws_account_id: Optional[pulumi.Input[str]] = None,
                 dashboard_id: Optional[pulumi.Input[str]] = None,
                 dashboard_publish_options: Optional[pulumi.Input[pulumi.InputType['DashboardPublishOptionsArgs']]] = None,
                 definition: Optional[pulumi.Input[pulumi.InputType['DashboardVersionDefinitionArgs']]] = None,
                 link_entities: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 link_sharing_configuration: Optional[pulumi.Input[pulumi.InputType['DashboardLinkSharingConfigurationArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 parameters: Optional[pulumi.Input[pulumi.InputType['DashboardParametersArgs']]] = None,
                 permissions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DashboardResourcePermissionArgs']]]]] = None,
                 source_entity: Optional[pulumi.Input[pulumi.InputType['DashboardSourceEntityArgs']]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['_root_inputs.TagArgs']]]]] = None,
                 theme_arn: Optional[pulumi.Input[str]] = None,
                 validation_strategy: Optional[pulumi.Input[pulumi.InputType['DashboardValidationStrategyArgs']]] = None,
                 version_description: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Definition of the AWS::QuickSight::Dashboard Resource Type.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DashboardArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Definition of the AWS::QuickSight::Dashboard Resource Type.

        :param str resource_name: The name of the resource.
        :param DashboardArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DashboardArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 aws_account_id: Optional[pulumi.Input[str]] = None,
                 dashboard_id: Optional[pulumi.Input[str]] = None,
                 dashboard_publish_options: Optional[pulumi.Input[pulumi.InputType['DashboardPublishOptionsArgs']]] = None,
                 definition: Optional[pulumi.Input[pulumi.InputType['DashboardVersionDefinitionArgs']]] = None,
                 link_entities: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 link_sharing_configuration: Optional[pulumi.Input[pulumi.InputType['DashboardLinkSharingConfigurationArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 parameters: Optional[pulumi.Input[pulumi.InputType['DashboardParametersArgs']]] = None,
                 permissions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DashboardResourcePermissionArgs']]]]] = None,
                 source_entity: Optional[pulumi.Input[pulumi.InputType['DashboardSourceEntityArgs']]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['_root_inputs.TagArgs']]]]] = None,
                 theme_arn: Optional[pulumi.Input[str]] = None,
                 validation_strategy: Optional[pulumi.Input[pulumi.InputType['DashboardValidationStrategyArgs']]] = None,
                 version_description: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DashboardArgs.__new__(DashboardArgs)

            if aws_account_id is None and not opts.urn:
                raise TypeError("Missing required property 'aws_account_id'")
            __props__.__dict__["aws_account_id"] = aws_account_id
            if dashboard_id is None and not opts.urn:
                raise TypeError("Missing required property 'dashboard_id'")
            __props__.__dict__["dashboard_id"] = dashboard_id
            __props__.__dict__["dashboard_publish_options"] = dashboard_publish_options
            __props__.__dict__["definition"] = definition
            __props__.__dict__["link_entities"] = link_entities
            __props__.__dict__["link_sharing_configuration"] = link_sharing_configuration
            __props__.__dict__["name"] = name
            __props__.__dict__["parameters"] = parameters
            __props__.__dict__["permissions"] = permissions
            __props__.__dict__["source_entity"] = source_entity
            __props__.__dict__["tags"] = tags
            __props__.__dict__["theme_arn"] = theme_arn
            __props__.__dict__["validation_strategy"] = validation_strategy
            __props__.__dict__["version_description"] = version_description
            __props__.__dict__["arn"] = None
            __props__.__dict__["created_time"] = None
            __props__.__dict__["last_published_time"] = None
            __props__.__dict__["last_updated_time"] = None
            __props__.__dict__["version"] = None
        replace_on_changes = pulumi.ResourceOptions(replace_on_changes=["aws_account_id", "dashboard_id"])
        opts = pulumi.ResourceOptions.merge(opts, replace_on_changes)
        super(Dashboard, __self__).__init__(
            'aws-native:quicksight:Dashboard',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Dashboard':
        """
        Get an existing Dashboard resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = DashboardArgs.__new__(DashboardArgs)

        __props__.__dict__["arn"] = None
        __props__.__dict__["aws_account_id"] = None
        __props__.__dict__["created_time"] = None
        __props__.__dict__["dashboard_id"] = None
        __props__.__dict__["dashboard_publish_options"] = None
        __props__.__dict__["definition"] = None
        __props__.__dict__["last_published_time"] = None
        __props__.__dict__["last_updated_time"] = None
        __props__.__dict__["link_entities"] = None
        __props__.__dict__["link_sharing_configuration"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["parameters"] = None
        __props__.__dict__["permissions"] = None
        __props__.__dict__["source_entity"] = None
        __props__.__dict__["tags"] = None
        __props__.__dict__["theme_arn"] = None
        __props__.__dict__["validation_strategy"] = None
        __props__.__dict__["version"] = None
        __props__.__dict__["version_description"] = None
        return Dashboard(resource_name, opts=opts, __props__=__props__)

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
    @pulumi.getter(name="dashboardId")
    def dashboard_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "dashboard_id")

    @property
    @pulumi.getter(name="dashboardPublishOptions")
    def dashboard_publish_options(self) -> pulumi.Output[Optional['outputs.DashboardPublishOptions']]:
        return pulumi.get(self, "dashboard_publish_options")

    @property
    @pulumi.getter
    def definition(self) -> pulumi.Output[Optional['outputs.DashboardVersionDefinition']]:
        return pulumi.get(self, "definition")

    @property
    @pulumi.getter(name="lastPublishedTime")
    def last_published_time(self) -> pulumi.Output[str]:
        return pulumi.get(self, "last_published_time")

    @property
    @pulumi.getter(name="lastUpdatedTime")
    def last_updated_time(self) -> pulumi.Output[str]:
        return pulumi.get(self, "last_updated_time")

    @property
    @pulumi.getter(name="linkEntities")
    def link_entities(self) -> pulumi.Output[Optional[Sequence[str]]]:
        return pulumi.get(self, "link_entities")

    @property
    @pulumi.getter(name="linkSharingConfiguration")
    def link_sharing_configuration(self) -> pulumi.Output[Optional['outputs.DashboardLinkSharingConfiguration']]:
        return pulumi.get(self, "link_sharing_configuration")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def parameters(self) -> pulumi.Output[Optional['outputs.DashboardParameters']]:
        return pulumi.get(self, "parameters")

    @property
    @pulumi.getter
    def permissions(self) -> pulumi.Output[Optional[Sequence['outputs.DashboardResourcePermission']]]:
        return pulumi.get(self, "permissions")

    @property
    @pulumi.getter(name="sourceEntity")
    def source_entity(self) -> pulumi.Output[Optional['outputs.DashboardSourceEntity']]:
        return pulumi.get(self, "source_entity")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence['_root_outputs.Tag']]]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="themeArn")
    def theme_arn(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "theme_arn")

    @property
    @pulumi.getter(name="validationStrategy")
    def validation_strategy(self) -> pulumi.Output[Optional['outputs.DashboardValidationStrategy']]:
        return pulumi.get(self, "validation_strategy")

    @property
    @pulumi.getter
    def version(self) -> pulumi.Output['outputs.DashboardVersion']:
        return pulumi.get(self, "version")

    @property
    @pulumi.getter(name="versionDescription")
    def version_description(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "version_description")

