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

__all__ = ['DecoderManifestArgs', 'DecoderManifest']

@pulumi.input_type
class DecoderManifestArgs:
    def __init__(__self__, *,
                 model_manifest_arn: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 network_interfaces: Optional[pulumi.Input[Sequence[pulumi.Input[Union['DecoderManifestCanNetworkInterfaceArgs', 'DecoderManifestObdNetworkInterfaceArgs']]]]] = None,
                 signal_decoders: Optional[pulumi.Input[Sequence[pulumi.Input[Union['DecoderManifestCanSignalDecoderArgs', 'DecoderManifestObdSignalDecoderArgs']]]]] = None,
                 status: Optional[pulumi.Input['DecoderManifestManifestStatus']] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]] = None):
        """
        The set of arguments for constructing a DecoderManifest resource.
        """
        pulumi.set(__self__, "model_manifest_arn", model_manifest_arn)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if network_interfaces is not None:
            pulumi.set(__self__, "network_interfaces", network_interfaces)
        if signal_decoders is not None:
            pulumi.set(__self__, "signal_decoders", signal_decoders)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="modelManifestArn")
    def model_manifest_arn(self) -> pulumi.Input[str]:
        return pulumi.get(self, "model_manifest_arn")

    @model_manifest_arn.setter
    def model_manifest_arn(self, value: pulumi.Input[str]):
        pulumi.set(self, "model_manifest_arn", value)

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
    @pulumi.getter(name="networkInterfaces")
    def network_interfaces(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[Union['DecoderManifestCanNetworkInterfaceArgs', 'DecoderManifestObdNetworkInterfaceArgs']]]]]:
        return pulumi.get(self, "network_interfaces")

    @network_interfaces.setter
    def network_interfaces(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[Union['DecoderManifestCanNetworkInterfaceArgs', 'DecoderManifestObdNetworkInterfaceArgs']]]]]):
        pulumi.set(self, "network_interfaces", value)

    @property
    @pulumi.getter(name="signalDecoders")
    def signal_decoders(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[Union['DecoderManifestCanSignalDecoderArgs', 'DecoderManifestObdSignalDecoderArgs']]]]]:
        return pulumi.get(self, "signal_decoders")

    @signal_decoders.setter
    def signal_decoders(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[Union['DecoderManifestCanSignalDecoderArgs', 'DecoderManifestObdSignalDecoderArgs']]]]]):
        pulumi.set(self, "signal_decoders", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input['DecoderManifestManifestStatus']]:
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input['DecoderManifestManifestStatus']]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]]:
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]]):
        pulumi.set(self, "tags", value)


warnings.warn("""DecoderManifest is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.""", DeprecationWarning)


class DecoderManifest(pulumi.CustomResource):
    warnings.warn("""DecoderManifest is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.""", DeprecationWarning)

    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 model_manifest_arn: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 network_interfaces: Optional[pulumi.Input[Sequence[pulumi.Input[Union[pulumi.InputType['DecoderManifestCanNetworkInterfaceArgs'], pulumi.InputType['DecoderManifestObdNetworkInterfaceArgs']]]]]] = None,
                 signal_decoders: Optional[pulumi.Input[Sequence[pulumi.Input[Union[pulumi.InputType['DecoderManifestCanSignalDecoderArgs'], pulumi.InputType['DecoderManifestObdSignalDecoderArgs']]]]]] = None,
                 status: Optional[pulumi.Input['DecoderManifestManifestStatus']] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['_root_inputs.TagArgs']]]]] = None,
                 __props__=None):
        """
        Definition of AWS::IoTFleetWise::DecoderManifest Resource Type

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DecoderManifestArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Definition of AWS::IoTFleetWise::DecoderManifest Resource Type

        :param str resource_name: The name of the resource.
        :param DecoderManifestArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DecoderManifestArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 model_manifest_arn: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 network_interfaces: Optional[pulumi.Input[Sequence[pulumi.Input[Union[pulumi.InputType['DecoderManifestCanNetworkInterfaceArgs'], pulumi.InputType['DecoderManifestObdNetworkInterfaceArgs']]]]]] = None,
                 signal_decoders: Optional[pulumi.Input[Sequence[pulumi.Input[Union[pulumi.InputType['DecoderManifestCanSignalDecoderArgs'], pulumi.InputType['DecoderManifestObdSignalDecoderArgs']]]]]] = None,
                 status: Optional[pulumi.Input['DecoderManifestManifestStatus']] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['_root_inputs.TagArgs']]]]] = None,
                 __props__=None):
        pulumi.log.warn("""DecoderManifest is deprecated: DecoderManifest is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.""")
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DecoderManifestArgs.__new__(DecoderManifestArgs)

            __props__.__dict__["description"] = description
            if model_manifest_arn is None and not opts.urn:
                raise TypeError("Missing required property 'model_manifest_arn'")
            __props__.__dict__["model_manifest_arn"] = model_manifest_arn
            __props__.__dict__["name"] = name
            __props__.__dict__["network_interfaces"] = network_interfaces
            __props__.__dict__["signal_decoders"] = signal_decoders
            __props__.__dict__["status"] = status
            __props__.__dict__["tags"] = tags
            __props__.__dict__["arn"] = None
            __props__.__dict__["creation_time"] = None
            __props__.__dict__["last_modification_time"] = None
        replace_on_changes = pulumi.ResourceOptions(replace_on_changes=["model_manifest_arn", "name"])
        opts = pulumi.ResourceOptions.merge(opts, replace_on_changes)
        super(DecoderManifest, __self__).__init__(
            'aws-native:iotfleetwise:DecoderManifest',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'DecoderManifest':
        """
        Get an existing DecoderManifest resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = DecoderManifestArgs.__new__(DecoderManifestArgs)

        __props__.__dict__["arn"] = None
        __props__.__dict__["creation_time"] = None
        __props__.__dict__["description"] = None
        __props__.__dict__["last_modification_time"] = None
        __props__.__dict__["model_manifest_arn"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["network_interfaces"] = None
        __props__.__dict__["signal_decoders"] = None
        __props__.__dict__["status"] = None
        __props__.__dict__["tags"] = None
        return DecoderManifest(resource_name, opts=opts, __props__=__props__)

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
    @pulumi.getter(name="modelManifestArn")
    def model_manifest_arn(self) -> pulumi.Output[str]:
        return pulumi.get(self, "model_manifest_arn")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="networkInterfaces")
    def network_interfaces(self) -> pulumi.Output[Optional[Sequence[Any]]]:
        return pulumi.get(self, "network_interfaces")

    @property
    @pulumi.getter(name="signalDecoders")
    def signal_decoders(self) -> pulumi.Output[Optional[Sequence[Any]]]:
        return pulumi.get(self, "signal_decoders")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[Optional['DecoderManifestManifestStatus']]:
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence['_root_outputs.Tag']]]:
        return pulumi.get(self, "tags")

