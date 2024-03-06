# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from .. import _inputs as _root_inputs
from .. import outputs as _root_outputs

__all__ = ['TrafficMirrorSessionArgs', 'TrafficMirrorSession']

@pulumi.input_type
class TrafficMirrorSessionArgs:
    def __init__(__self__, *,
                 network_interface_id: pulumi.Input[str],
                 session_number: pulumi.Input[int],
                 traffic_mirror_filter_id: pulumi.Input[str],
                 traffic_mirror_target_id: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 packet_length: Optional[pulumi.Input[int]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]] = None,
                 virtual_network_id: Optional[pulumi.Input[int]] = None):
        """
        The set of arguments for constructing a TrafficMirrorSession resource.
        """
        pulumi.set(__self__, "network_interface_id", network_interface_id)
        pulumi.set(__self__, "session_number", session_number)
        pulumi.set(__self__, "traffic_mirror_filter_id", traffic_mirror_filter_id)
        pulumi.set(__self__, "traffic_mirror_target_id", traffic_mirror_target_id)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if packet_length is not None:
            pulumi.set(__self__, "packet_length", packet_length)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if virtual_network_id is not None:
            pulumi.set(__self__, "virtual_network_id", virtual_network_id)

    @property
    @pulumi.getter(name="networkInterfaceId")
    def network_interface_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "network_interface_id")

    @network_interface_id.setter
    def network_interface_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "network_interface_id", value)

    @property
    @pulumi.getter(name="sessionNumber")
    def session_number(self) -> pulumi.Input[int]:
        return pulumi.get(self, "session_number")

    @session_number.setter
    def session_number(self, value: pulumi.Input[int]):
        pulumi.set(self, "session_number", value)

    @property
    @pulumi.getter(name="trafficMirrorFilterId")
    def traffic_mirror_filter_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "traffic_mirror_filter_id")

    @traffic_mirror_filter_id.setter
    def traffic_mirror_filter_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "traffic_mirror_filter_id", value)

    @property
    @pulumi.getter(name="trafficMirrorTargetId")
    def traffic_mirror_target_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "traffic_mirror_target_id")

    @traffic_mirror_target_id.setter
    def traffic_mirror_target_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "traffic_mirror_target_id", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="packetLength")
    def packet_length(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "packet_length")

    @packet_length.setter
    def packet_length(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "packet_length", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]]:
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="virtualNetworkId")
    def virtual_network_id(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "virtual_network_id")

    @virtual_network_id.setter
    def virtual_network_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "virtual_network_id", value)


warnings.warn("""TrafficMirrorSession is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.""", DeprecationWarning)


class TrafficMirrorSession(pulumi.CustomResource):
    warnings.warn("""TrafficMirrorSession is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.""", DeprecationWarning)

    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 network_interface_id: Optional[pulumi.Input[str]] = None,
                 packet_length: Optional[pulumi.Input[int]] = None,
                 session_number: Optional[pulumi.Input[int]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['_root_inputs.TagArgs']]]]] = None,
                 traffic_mirror_filter_id: Optional[pulumi.Input[str]] = None,
                 traffic_mirror_target_id: Optional[pulumi.Input[str]] = None,
                 virtual_network_id: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        """
        Resource Type definition for AWS::EC2::TrafficMirrorSession

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: TrafficMirrorSessionArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource Type definition for AWS::EC2::TrafficMirrorSession

        :param str resource_name: The name of the resource.
        :param TrafficMirrorSessionArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(TrafficMirrorSessionArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 network_interface_id: Optional[pulumi.Input[str]] = None,
                 packet_length: Optional[pulumi.Input[int]] = None,
                 session_number: Optional[pulumi.Input[int]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['_root_inputs.TagArgs']]]]] = None,
                 traffic_mirror_filter_id: Optional[pulumi.Input[str]] = None,
                 traffic_mirror_target_id: Optional[pulumi.Input[str]] = None,
                 virtual_network_id: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        pulumi.log.warn("""TrafficMirrorSession is deprecated: TrafficMirrorSession is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.""")
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = TrafficMirrorSessionArgs.__new__(TrafficMirrorSessionArgs)

            __props__.__dict__["description"] = description
            if network_interface_id is None and not opts.urn:
                raise TypeError("Missing required property 'network_interface_id'")
            __props__.__dict__["network_interface_id"] = network_interface_id
            __props__.__dict__["packet_length"] = packet_length
            if session_number is None and not opts.urn:
                raise TypeError("Missing required property 'session_number'")
            __props__.__dict__["session_number"] = session_number
            __props__.__dict__["tags"] = tags
            if traffic_mirror_filter_id is None and not opts.urn:
                raise TypeError("Missing required property 'traffic_mirror_filter_id'")
            __props__.__dict__["traffic_mirror_filter_id"] = traffic_mirror_filter_id
            if traffic_mirror_target_id is None and not opts.urn:
                raise TypeError("Missing required property 'traffic_mirror_target_id'")
            __props__.__dict__["traffic_mirror_target_id"] = traffic_mirror_target_id
            __props__.__dict__["virtual_network_id"] = virtual_network_id
            __props__.__dict__["aws_id"] = None
        replace_on_changes = pulumi.ResourceOptions(replace_on_changes=["network_interface_id"])
        opts = pulumi.ResourceOptions.merge(opts, replace_on_changes)
        super(TrafficMirrorSession, __self__).__init__(
            'aws-native:ec2:TrafficMirrorSession',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'TrafficMirrorSession':
        """
        Get an existing TrafficMirrorSession resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = TrafficMirrorSessionArgs.__new__(TrafficMirrorSessionArgs)

        __props__.__dict__["aws_id"] = None
        __props__.__dict__["description"] = None
        __props__.__dict__["network_interface_id"] = None
        __props__.__dict__["packet_length"] = None
        __props__.__dict__["session_number"] = None
        __props__.__dict__["tags"] = None
        __props__.__dict__["traffic_mirror_filter_id"] = None
        __props__.__dict__["traffic_mirror_target_id"] = None
        __props__.__dict__["virtual_network_id"] = None
        return TrafficMirrorSession(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="awsId")
    def aws_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "aws_id")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="networkInterfaceId")
    def network_interface_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "network_interface_id")

    @property
    @pulumi.getter(name="packetLength")
    def packet_length(self) -> pulumi.Output[Optional[int]]:
        return pulumi.get(self, "packet_length")

    @property
    @pulumi.getter(name="sessionNumber")
    def session_number(self) -> pulumi.Output[int]:
        return pulumi.get(self, "session_number")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence['_root_outputs.Tag']]]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="trafficMirrorFilterId")
    def traffic_mirror_filter_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "traffic_mirror_filter_id")

    @property
    @pulumi.getter(name="trafficMirrorTargetId")
    def traffic_mirror_target_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "traffic_mirror_target_id")

    @property
    @pulumi.getter(name="virtualNetworkId")
    def virtual_network_id(self) -> pulumi.Output[Optional[int]]:
        return pulumi.get(self, "virtual_network_id")

