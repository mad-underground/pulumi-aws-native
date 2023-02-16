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
from ._inputs import *

__all__ = ['ApplicationArgs', 'Application']

@pulumi.input_type
class ApplicationArgs:
    def __init__(__self__, *,
                 inputs: pulumi.Input[Sequence[pulumi.Input['ApplicationInputArgs']]],
                 application_code: Optional[pulumi.Input[str]] = None,
                 application_description: Optional[pulumi.Input[str]] = None,
                 application_name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Application resource.
        """
        pulumi.set(__self__, "inputs", inputs)
        if application_code is not None:
            pulumi.set(__self__, "application_code", application_code)
        if application_description is not None:
            pulumi.set(__self__, "application_description", application_description)
        if application_name is not None:
            pulumi.set(__self__, "application_name", application_name)

    @property
    @pulumi.getter
    def inputs(self) -> pulumi.Input[Sequence[pulumi.Input['ApplicationInputArgs']]]:
        return pulumi.get(self, "inputs")

    @inputs.setter
    def inputs(self, value: pulumi.Input[Sequence[pulumi.Input['ApplicationInputArgs']]]):
        pulumi.set(self, "inputs", value)

    @property
    @pulumi.getter(name="applicationCode")
    def application_code(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "application_code")

    @application_code.setter
    def application_code(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "application_code", value)

    @property
    @pulumi.getter(name="applicationDescription")
    def application_description(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "application_description")

    @application_description.setter
    def application_description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "application_description", value)

    @property
    @pulumi.getter(name="applicationName")
    def application_name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "application_name")

    @application_name.setter
    def application_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "application_name", value)


warnings.warn("""Application is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.""", DeprecationWarning)


class Application(pulumi.CustomResource):
    warnings.warn("""Application is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.""", DeprecationWarning)

    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 application_code: Optional[pulumi.Input[str]] = None,
                 application_description: Optional[pulumi.Input[str]] = None,
                 application_name: Optional[pulumi.Input[str]] = None,
                 inputs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ApplicationInputArgs']]]]] = None,
                 __props__=None):
        """
        Resource Type definition for AWS::KinesisAnalytics::Application

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ApplicationArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource Type definition for AWS::KinesisAnalytics::Application

        :param str resource_name: The name of the resource.
        :param ApplicationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ApplicationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 application_code: Optional[pulumi.Input[str]] = None,
                 application_description: Optional[pulumi.Input[str]] = None,
                 application_name: Optional[pulumi.Input[str]] = None,
                 inputs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ApplicationInputArgs']]]]] = None,
                 __props__=None):
        pulumi.log.warn("""Application is deprecated: Application is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.""")
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ApplicationArgs.__new__(ApplicationArgs)

            __props__.__dict__["application_code"] = application_code
            __props__.__dict__["application_description"] = application_description
            __props__.__dict__["application_name"] = application_name
            if inputs is None and not opts.urn:
                raise TypeError("Missing required property 'inputs'")
            __props__.__dict__["inputs"] = inputs
        super(Application, __self__).__init__(
            'aws-native:kinesisanalytics:Application',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Application':
        """
        Get an existing Application resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = ApplicationArgs.__new__(ApplicationArgs)

        __props__.__dict__["application_code"] = None
        __props__.__dict__["application_description"] = None
        __props__.__dict__["application_name"] = None
        __props__.__dict__["inputs"] = None
        return Application(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="applicationCode")
    def application_code(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "application_code")

    @property
    @pulumi.getter(name="applicationDescription")
    def application_description(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "application_description")

    @property
    @pulumi.getter(name="applicationName")
    def application_name(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "application_name")

    @property
    @pulumi.getter
    def inputs(self) -> pulumi.Output[Sequence['outputs.ApplicationInput']]:
        return pulumi.get(self, "inputs")

