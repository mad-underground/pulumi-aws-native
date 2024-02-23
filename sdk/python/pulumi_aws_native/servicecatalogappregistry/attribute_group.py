# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['AttributeGroupArgs', 'AttributeGroup']

@pulumi.input_type
class AttributeGroupArgs:
    def __init__(__self__, *,
                 attributes: Any,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a AttributeGroup resource.
        :param Any attributes: Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::ServiceCatalogAppRegistry::AttributeGroup` for more information about the expected schema for this property.
        :param pulumi.Input[str] description: The description of the attribute group. 
        :param pulumi.Input[str] name: The name of the attribute group. 
        """
        pulumi.set(__self__, "attributes", attributes)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def attributes(self) -> Any:
        """
        Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::ServiceCatalogAppRegistry::AttributeGroup` for more information about the expected schema for this property.
        """
        return pulumi.get(self, "attributes")

    @attributes.setter
    def attributes(self, value: Any):
        pulumi.set(self, "attributes", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the attribute group. 
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the attribute group. 
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)


class AttributeGroup(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 attributes: Optional[Any] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        Resource Schema for AWS::ServiceCatalogAppRegistry::AttributeGroup.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param Any attributes: Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::ServiceCatalogAppRegistry::AttributeGroup` for more information about the expected schema for this property.
        :param pulumi.Input[str] description: The description of the attribute group. 
        :param pulumi.Input[str] name: The name of the attribute group. 
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AttributeGroupArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource Schema for AWS::ServiceCatalogAppRegistry::AttributeGroup.

        :param str resource_name: The name of the resource.
        :param AttributeGroupArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AttributeGroupArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 attributes: Optional[Any] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AttributeGroupArgs.__new__(AttributeGroupArgs)

            if attributes is None and not opts.urn:
                raise TypeError("Missing required property 'attributes'")
            __props__.__dict__["attributes"] = attributes
            __props__.__dict__["description"] = description
            __props__.__dict__["name"] = name
            __props__.__dict__["tags"] = tags
            __props__.__dict__["arn"] = None
        super(AttributeGroup, __self__).__init__(
            'aws-native:servicecatalogappregistry:AttributeGroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'AttributeGroup':
        """
        Get an existing AttributeGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = AttributeGroupArgs.__new__(AttributeGroupArgs)

        __props__.__dict__["arn"] = None
        __props__.__dict__["attributes"] = None
        __props__.__dict__["description"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["tags"] = None
        return AttributeGroup(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter
    def attributes(self) -> pulumi.Output[Any]:
        """
        Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::ServiceCatalogAppRegistry::AttributeGroup` for more information about the expected schema for this property.
        """
        return pulumi.get(self, "attributes")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The description of the attribute group. 
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the attribute group. 
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        return pulumi.get(self, "tags")

