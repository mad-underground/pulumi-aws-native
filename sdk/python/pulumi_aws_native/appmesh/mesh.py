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
from ._inputs import *

__all__ = ['MeshArgs', 'Mesh']

@pulumi.input_type
class MeshArgs:
    def __init__(__self__, *,
                 mesh_name: Optional[pulumi.Input[str]] = None,
                 spec: Optional[pulumi.Input['MeshSpecArgs']] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]] = None):
        """
        The set of arguments for constructing a Mesh resource.
        """
        if mesh_name is not None:
            pulumi.set(__self__, "mesh_name", mesh_name)
        if spec is not None:
            pulumi.set(__self__, "spec", spec)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="meshName")
    def mesh_name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "mesh_name")

    @mesh_name.setter
    def mesh_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "mesh_name", value)

    @property
    @pulumi.getter
    def spec(self) -> Optional[pulumi.Input['MeshSpecArgs']]:
        return pulumi.get(self, "spec")

    @spec.setter
    def spec(self, value: Optional[pulumi.Input['MeshSpecArgs']]):
        pulumi.set(self, "spec", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]]:
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]]):
        pulumi.set(self, "tags", value)


warnings.warn("""Mesh is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.""", DeprecationWarning)


class Mesh(pulumi.CustomResource):
    warnings.warn("""Mesh is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.""", DeprecationWarning)

    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 mesh_name: Optional[pulumi.Input[str]] = None,
                 spec: Optional[pulumi.Input[pulumi.InputType['MeshSpecArgs']]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['_root_inputs.TagArgs']]]]] = None,
                 __props__=None):
        """
        Resource Type definition for AWS::AppMesh::Mesh

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[MeshArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource Type definition for AWS::AppMesh::Mesh

        :param str resource_name: The name of the resource.
        :param MeshArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(MeshArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 mesh_name: Optional[pulumi.Input[str]] = None,
                 spec: Optional[pulumi.Input[pulumi.InputType['MeshSpecArgs']]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['_root_inputs.TagArgs']]]]] = None,
                 __props__=None):
        pulumi.log.warn("""Mesh is deprecated: Mesh is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.""")
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = MeshArgs.__new__(MeshArgs)

            __props__.__dict__["mesh_name"] = mesh_name
            __props__.__dict__["spec"] = spec
            __props__.__dict__["tags"] = tags
            __props__.__dict__["arn"] = None
            __props__.__dict__["aws_id"] = None
            __props__.__dict__["mesh_owner"] = None
            __props__.__dict__["resource_owner"] = None
            __props__.__dict__["uid"] = None
        replace_on_changes = pulumi.ResourceOptions(replace_on_changes=["mesh_name"])
        opts = pulumi.ResourceOptions.merge(opts, replace_on_changes)
        super(Mesh, __self__).__init__(
            'aws-native:appmesh:Mesh',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Mesh':
        """
        Get an existing Mesh resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = MeshArgs.__new__(MeshArgs)

        __props__.__dict__["arn"] = None
        __props__.__dict__["aws_id"] = None
        __props__.__dict__["mesh_name"] = None
        __props__.__dict__["mesh_owner"] = None
        __props__.__dict__["resource_owner"] = None
        __props__.__dict__["spec"] = None
        __props__.__dict__["tags"] = None
        __props__.__dict__["uid"] = None
        return Mesh(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="awsId")
    def aws_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "aws_id")

    @property
    @pulumi.getter(name="meshName")
    def mesh_name(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "mesh_name")

    @property
    @pulumi.getter(name="meshOwner")
    def mesh_owner(self) -> pulumi.Output[str]:
        return pulumi.get(self, "mesh_owner")

    @property
    @pulumi.getter(name="resourceOwner")
    def resource_owner(self) -> pulumi.Output[str]:
        return pulumi.get(self, "resource_owner")

    @property
    @pulumi.getter
    def spec(self) -> pulumi.Output[Optional['outputs.MeshSpec']]:
        return pulumi.get(self, "spec")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence['_root_outputs.Tag']]]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def uid(self) -> pulumi.Output[str]:
        return pulumi.get(self, "uid")

