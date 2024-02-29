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

__all__ = ['NetworkInterfaceAttachmentArgs', 'NetworkInterfaceAttachment']

@pulumi.input_type
class NetworkInterfaceAttachmentArgs:
    def __init__(__self__, *,
                 device_index: pulumi.Input[str],
                 instance_id: pulumi.Input[str],
                 network_interface_id: pulumi.Input[str],
                 delete_on_termination: Optional[pulumi.Input[bool]] = None,
                 ena_srd_specification: Optional[pulumi.Input['NetworkInterfaceAttachmentEnaSrdSpecificationArgs']] = None):
        """
        The set of arguments for constructing a NetworkInterfaceAttachment resource.
        :param pulumi.Input[str] device_index: The network interface's position in the attachment order. For example, the first attached network interface has a DeviceIndex of 0.
        :param pulumi.Input[str] instance_id: The ID of the instance to which you will attach the ENI.
        :param pulumi.Input[str] network_interface_id: The ID of the ENI that you want to attach.
        :param pulumi.Input[bool] delete_on_termination: Whether to delete the network interface when the instance terminates. By default, this value is set to true.
        """
        pulumi.set(__self__, "device_index", device_index)
        pulumi.set(__self__, "instance_id", instance_id)
        pulumi.set(__self__, "network_interface_id", network_interface_id)
        if delete_on_termination is not None:
            pulumi.set(__self__, "delete_on_termination", delete_on_termination)
        if ena_srd_specification is not None:
            pulumi.set(__self__, "ena_srd_specification", ena_srd_specification)

    @property
    @pulumi.getter(name="deviceIndex")
    def device_index(self) -> pulumi.Input[str]:
        """
        The network interface's position in the attachment order. For example, the first attached network interface has a DeviceIndex of 0.
        """
        return pulumi.get(self, "device_index")

    @device_index.setter
    def device_index(self, value: pulumi.Input[str]):
        pulumi.set(self, "device_index", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Input[str]:
        """
        The ID of the instance to which you will attach the ENI.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="networkInterfaceId")
    def network_interface_id(self) -> pulumi.Input[str]:
        """
        The ID of the ENI that you want to attach.
        """
        return pulumi.get(self, "network_interface_id")

    @network_interface_id.setter
    def network_interface_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "network_interface_id", value)

    @property
    @pulumi.getter(name="deleteOnTermination")
    def delete_on_termination(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to delete the network interface when the instance terminates. By default, this value is set to true.
        """
        return pulumi.get(self, "delete_on_termination")

    @delete_on_termination.setter
    def delete_on_termination(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "delete_on_termination", value)

    @property
    @pulumi.getter(name="enaSrdSpecification")
    def ena_srd_specification(self) -> Optional[pulumi.Input['NetworkInterfaceAttachmentEnaSrdSpecificationArgs']]:
        return pulumi.get(self, "ena_srd_specification")

    @ena_srd_specification.setter
    def ena_srd_specification(self, value: Optional[pulumi.Input['NetworkInterfaceAttachmentEnaSrdSpecificationArgs']]):
        pulumi.set(self, "ena_srd_specification", value)


class NetworkInterfaceAttachment(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 delete_on_termination: Optional[pulumi.Input[bool]] = None,
                 device_index: Optional[pulumi.Input[str]] = None,
                 ena_srd_specification: Optional[pulumi.Input[pulumi.InputType['NetworkInterfaceAttachmentEnaSrdSpecificationArgs']]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 network_interface_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Resource Type definition for AWS::EC2::NetworkInterfaceAttachment

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] delete_on_termination: Whether to delete the network interface when the instance terminates. By default, this value is set to true.
        :param pulumi.Input[str] device_index: The network interface's position in the attachment order. For example, the first attached network interface has a DeviceIndex of 0.
        :param pulumi.Input[str] instance_id: The ID of the instance to which you will attach the ENI.
        :param pulumi.Input[str] network_interface_id: The ID of the ENI that you want to attach.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: NetworkInterfaceAttachmentArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource Type definition for AWS::EC2::NetworkInterfaceAttachment

        :param str resource_name: The name of the resource.
        :param NetworkInterfaceAttachmentArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(NetworkInterfaceAttachmentArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 delete_on_termination: Optional[pulumi.Input[bool]] = None,
                 device_index: Optional[pulumi.Input[str]] = None,
                 ena_srd_specification: Optional[pulumi.Input[pulumi.InputType['NetworkInterfaceAttachmentEnaSrdSpecificationArgs']]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 network_interface_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = NetworkInterfaceAttachmentArgs.__new__(NetworkInterfaceAttachmentArgs)

            __props__.__dict__["delete_on_termination"] = delete_on_termination
            if device_index is None and not opts.urn:
                raise TypeError("Missing required property 'device_index'")
            __props__.__dict__["device_index"] = device_index
            __props__.__dict__["ena_srd_specification"] = ena_srd_specification
            if instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'instance_id'")
            __props__.__dict__["instance_id"] = instance_id
            if network_interface_id is None and not opts.urn:
                raise TypeError("Missing required property 'network_interface_id'")
            __props__.__dict__["network_interface_id"] = network_interface_id
            __props__.__dict__["attachment_id"] = None
        replace_on_changes = pulumi.ResourceOptions(replace_on_changes=["device_index", "instance_id", "network_interface_id"])
        opts = pulumi.ResourceOptions.merge(opts, replace_on_changes)
        super(NetworkInterfaceAttachment, __self__).__init__(
            'aws-native:ec2:NetworkInterfaceAttachment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'NetworkInterfaceAttachment':
        """
        Get an existing NetworkInterfaceAttachment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = NetworkInterfaceAttachmentArgs.__new__(NetworkInterfaceAttachmentArgs)

        __props__.__dict__["attachment_id"] = None
        __props__.__dict__["delete_on_termination"] = None
        __props__.__dict__["device_index"] = None
        __props__.__dict__["ena_srd_specification"] = None
        __props__.__dict__["instance_id"] = None
        __props__.__dict__["network_interface_id"] = None
        return NetworkInterfaceAttachment(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="attachmentId")
    def attachment_id(self) -> pulumi.Output[str]:
        """
        The ID of the network interface attachment.
        """
        return pulumi.get(self, "attachment_id")

    @property
    @pulumi.getter(name="deleteOnTermination")
    def delete_on_termination(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether to delete the network interface when the instance terminates. By default, this value is set to true.
        """
        return pulumi.get(self, "delete_on_termination")

    @property
    @pulumi.getter(name="deviceIndex")
    def device_index(self) -> pulumi.Output[str]:
        """
        The network interface's position in the attachment order. For example, the first attached network interface has a DeviceIndex of 0.
        """
        return pulumi.get(self, "device_index")

    @property
    @pulumi.getter(name="enaSrdSpecification")
    def ena_srd_specification(self) -> pulumi.Output[Optional['outputs.NetworkInterfaceAttachmentEnaSrdSpecification']]:
        return pulumi.get(self, "ena_srd_specification")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[str]:
        """
        The ID of the instance to which you will attach the ENI.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="networkInterfaceId")
    def network_interface_id(self) -> pulumi.Output[str]:
        """
        The ID of the ENI that you want to attach.
        """
        return pulumi.get(self, "network_interface_id")

