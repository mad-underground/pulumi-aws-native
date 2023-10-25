# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['ResourceArgs', 'Resource']

@pulumi.input_type
class ResourceArgs:
    def __init__(__self__, *,
                 parent_id: pulumi.Input[str],
                 path_part: pulumi.Input[str],
                 rest_api_id: pulumi.Input[str]):
        """
        The set of arguments for constructing a Resource resource.
        :param pulumi.Input[str] parent_id: The parent resource's identifier.
        :param pulumi.Input[str] path_part: The last path segment for this resource.
        :param pulumi.Input[str] rest_api_id: The string identifier of the associated RestApi.
        """
        ResourceArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            parent_id=parent_id,
            path_part=path_part,
            rest_api_id=rest_api_id,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             parent_id: pulumi.Input[str],
             path_part: pulumi.Input[str],
             rest_api_id: pulumi.Input[str],
             opts: Optional[pulumi.ResourceOptions]=None):
        _setter("parent_id", parent_id)
        _setter("path_part", path_part)
        _setter("rest_api_id", rest_api_id)

    @property
    @pulumi.getter(name="parentId")
    def parent_id(self) -> pulumi.Input[str]:
        """
        The parent resource's identifier.
        """
        return pulumi.get(self, "parent_id")

    @parent_id.setter
    def parent_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "parent_id", value)

    @property
    @pulumi.getter(name="pathPart")
    def path_part(self) -> pulumi.Input[str]:
        """
        The last path segment for this resource.
        """
        return pulumi.get(self, "path_part")

    @path_part.setter
    def path_part(self, value: pulumi.Input[str]):
        pulumi.set(self, "path_part", value)

    @property
    @pulumi.getter(name="restApiId")
    def rest_api_id(self) -> pulumi.Input[str]:
        """
        The string identifier of the associated RestApi.
        """
        return pulumi.get(self, "rest_api_id")

    @rest_api_id.setter
    def rest_api_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "rest_api_id", value)


class Resource(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 parent_id: Optional[pulumi.Input[str]] = None,
                 path_part: Optional[pulumi.Input[str]] = None,
                 rest_api_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        The ``AWS::ApiGateway::Resource`` resource creates a resource in an API.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] parent_id: The parent resource's identifier.
        :param pulumi.Input[str] path_part: The last path segment for this resource.
        :param pulumi.Input[str] rest_api_id: The string identifier of the associated RestApi.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ResourceArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The ``AWS::ApiGateway::Resource`` resource creates a resource in an API.

        :param str resource_name: The name of the resource.
        :param ResourceArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ResourceArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            ResourceArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 parent_id: Optional[pulumi.Input[str]] = None,
                 path_part: Optional[pulumi.Input[str]] = None,
                 rest_api_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ResourceArgs.__new__(ResourceArgs)

            if parent_id is None and not opts.urn:
                raise TypeError("Missing required property 'parent_id'")
            __props__.__dict__["parent_id"] = parent_id
            if path_part is None and not opts.urn:
                raise TypeError("Missing required property 'path_part'")
            __props__.__dict__["path_part"] = path_part
            if rest_api_id is None and not opts.urn:
                raise TypeError("Missing required property 'rest_api_id'")
            __props__.__dict__["rest_api_id"] = rest_api_id
            __props__.__dict__["resource_id"] = None
        replace_on_changes = pulumi.ResourceOptions(replace_on_changes=["parent_id", "path_part", "rest_api_id"])
        opts = pulumi.ResourceOptions.merge(opts, replace_on_changes)
        super(Resource, __self__).__init__(
            'aws-native:apigateway:Resource',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Resource':
        """
        Get an existing Resource resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = ResourceArgs.__new__(ResourceArgs)

        __props__.__dict__["parent_id"] = None
        __props__.__dict__["path_part"] = None
        __props__.__dict__["resource_id"] = None
        __props__.__dict__["rest_api_id"] = None
        return Resource(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="parentId")
    def parent_id(self) -> pulumi.Output[str]:
        """
        The parent resource's identifier.
        """
        return pulumi.get(self, "parent_id")

    @property
    @pulumi.getter(name="pathPart")
    def path_part(self) -> pulumi.Output[str]:
        """
        The last path segment for this resource.
        """
        return pulumi.get(self, "path_part")

    @property
    @pulumi.getter(name="resourceId")
    def resource_id(self) -> pulumi.Output[str]:
        """
        A unique primary identifier for a Resource
        """
        return pulumi.get(self, "resource_id")

    @property
    @pulumi.getter(name="restApiId")
    def rest_api_id(self) -> pulumi.Output[str]:
        """
        The string identifier of the associated RestApi.
        """
        return pulumi.get(self, "rest_api_id")

