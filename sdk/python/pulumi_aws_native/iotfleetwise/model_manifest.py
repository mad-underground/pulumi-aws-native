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
from ._enums import *
from ._inputs import *

__all__ = ['ModelManifestArgs', 'ModelManifest']

@pulumi.input_type
class ModelManifestArgs:
    def __init__(__self__, *,
                 signal_catalog_arn: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 nodes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 status: Optional[pulumi.Input['ModelManifestManifestStatus']] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['ModelManifestTagArgs']]]] = None):
        """
        The set of arguments for constructing a ModelManifest resource.
        """
        pulumi.set(__self__, "signal_catalog_arn", signal_catalog_arn)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if nodes is not None:
            pulumi.set(__self__, "nodes", nodes)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="signalCatalogArn")
    def signal_catalog_arn(self) -> pulumi.Input[str]:
        return pulumi.get(self, "signal_catalog_arn")

    @signal_catalog_arn.setter
    def signal_catalog_arn(self, value: pulumi.Input[str]):
        pulumi.set(self, "signal_catalog_arn", value)

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
    @pulumi.getter
    def nodes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "nodes")

    @nodes.setter
    def nodes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "nodes", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input['ModelManifestManifestStatus']]:
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input['ModelManifestManifestStatus']]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ModelManifestTagArgs']]]]:
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ModelManifestTagArgs']]]]):
        pulumi.set(self, "tags", value)


warnings.warn("""ModelManifest is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.""", DeprecationWarning)


class ModelManifest(pulumi.CustomResource):
    warnings.warn("""ModelManifest is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.""", DeprecationWarning)

    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 nodes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 signal_catalog_arn: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input['ModelManifestManifestStatus']] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ModelManifestTagArgs']]]]] = None,
                 __props__=None):
        """
        Definition of AWS::IoTFleetWise::ModelManifest Resource Type

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ModelManifestArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Definition of AWS::IoTFleetWise::ModelManifest Resource Type

        :param str resource_name: The name of the resource.
        :param ModelManifestArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ModelManifestArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 nodes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 signal_catalog_arn: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input['ModelManifestManifestStatus']] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ModelManifestTagArgs']]]]] = None,
                 __props__=None):
        pulumi.log.warn("""ModelManifest is deprecated: ModelManifest is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.""")
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ModelManifestArgs.__new__(ModelManifestArgs)

            __props__.__dict__["description"] = description
            __props__.__dict__["name"] = name
            __props__.__dict__["nodes"] = nodes
            if signal_catalog_arn is None and not opts.urn:
                raise TypeError("Missing required property 'signal_catalog_arn'")
            __props__.__dict__["signal_catalog_arn"] = signal_catalog_arn
            __props__.__dict__["status"] = status
            __props__.__dict__["tags"] = tags
            __props__.__dict__["arn"] = None
            __props__.__dict__["creation_time"] = None
            __props__.__dict__["last_modification_time"] = None
        super(ModelManifest, __self__).__init__(
            'aws-native:iotfleetwise:ModelManifest',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'ModelManifest':
        """
        Get an existing ModelManifest resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = ModelManifestArgs.__new__(ModelManifestArgs)

        __props__.__dict__["arn"] = None
        __props__.__dict__["creation_time"] = None
        __props__.__dict__["description"] = None
        __props__.__dict__["last_modification_time"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["nodes"] = None
        __props__.__dict__["signal_catalog_arn"] = None
        __props__.__dict__["status"] = None
        __props__.__dict__["tags"] = None
        return ModelManifest(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="creationTime")
    def creation_time(self) -> pulumi.Output[str]:
        return pulumi.get(self, "creation_time")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="lastModificationTime")
    def last_modification_time(self) -> pulumi.Output[str]:
        return pulumi.get(self, "last_modification_time")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def nodes(self) -> pulumi.Output[Optional[Sequence[str]]]:
        return pulumi.get(self, "nodes")

    @property
    @pulumi.getter(name="signalCatalogArn")
    def signal_catalog_arn(self) -> pulumi.Output[str]:
        return pulumi.get(self, "signal_catalog_arn")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[Optional['ModelManifestManifestStatus']]:
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence['outputs.ModelManifestTag']]]:
        return pulumi.get(self, "tags")
