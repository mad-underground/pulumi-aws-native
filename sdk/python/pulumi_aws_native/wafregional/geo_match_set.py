# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['GeoMatchSetArgs', 'GeoMatchSet']

@pulumi.input_type
class GeoMatchSetArgs:
    def __init__(__self__, *,
                 name: pulumi.Input[str],
                 geo_match_constraints: Optional[pulumi.Input[Sequence[pulumi.Input['GeoMatchSetGeoMatchConstraintArgs']]]] = None):
        """
        The set of arguments for constructing a GeoMatchSet resource.
        """
        pulumi.set(__self__, "name", name)
        if geo_match_constraints is not None:
            pulumi.set(__self__, "geo_match_constraints", geo_match_constraints)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="geoMatchConstraints")
    def geo_match_constraints(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['GeoMatchSetGeoMatchConstraintArgs']]]]:
        return pulumi.get(self, "geo_match_constraints")

    @geo_match_constraints.setter
    def geo_match_constraints(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['GeoMatchSetGeoMatchConstraintArgs']]]]):
        pulumi.set(self, "geo_match_constraints", value)


warnings.warn("""GeoMatchSet is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.""", DeprecationWarning)


class GeoMatchSet(pulumi.CustomResource):
    warnings.warn("""GeoMatchSet is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.""", DeprecationWarning)

    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 geo_match_constraints: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GeoMatchSetGeoMatchConstraintArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Resource Type definition for AWS::WAFRegional::GeoMatchSet

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: GeoMatchSetArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource Type definition for AWS::WAFRegional::GeoMatchSet

        :param str resource_name: The name of the resource.
        :param GeoMatchSetArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(GeoMatchSetArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 geo_match_constraints: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GeoMatchSetGeoMatchConstraintArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        pulumi.log.warn("""GeoMatchSet is deprecated: GeoMatchSet is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.""")
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = GeoMatchSetArgs.__new__(GeoMatchSetArgs)

            __props__.__dict__["geo_match_constraints"] = geo_match_constraints
            if name is None and not opts.urn:
                raise TypeError("Missing required property 'name'")
            __props__.__dict__["name"] = name
        super(GeoMatchSet, __self__).__init__(
            'aws-native:wafregional:GeoMatchSet',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'GeoMatchSet':
        """
        Get an existing GeoMatchSet resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = GeoMatchSetArgs.__new__(GeoMatchSetArgs)

        __props__.__dict__["geo_match_constraints"] = None
        __props__.__dict__["name"] = None
        return GeoMatchSet(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="geoMatchConstraints")
    def geo_match_constraints(self) -> pulumi.Output[Optional[Sequence['outputs.GeoMatchSetGeoMatchConstraint']]]:
        return pulumi.get(self, "geo_match_constraints")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "name")
