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

__all__ = [
    'GetLocationObjectStorageResult',
    'AwaitableGetLocationObjectStorageResult',
    'get_location_object_storage',
    'get_location_object_storage_output',
]

@pulumi.output_type
class GetLocationObjectStorageResult:
    def __init__(__self__, access_key=None, agent_arns=None, location_arn=None, location_uri=None, server_certificate=None, server_port=None, server_protocol=None, tags=None):
        if access_key and not isinstance(access_key, str):
            raise TypeError("Expected argument 'access_key' to be a str")
        pulumi.set(__self__, "access_key", access_key)
        if agent_arns and not isinstance(agent_arns, list):
            raise TypeError("Expected argument 'agent_arns' to be a list")
        pulumi.set(__self__, "agent_arns", agent_arns)
        if location_arn and not isinstance(location_arn, str):
            raise TypeError("Expected argument 'location_arn' to be a str")
        pulumi.set(__self__, "location_arn", location_arn)
        if location_uri and not isinstance(location_uri, str):
            raise TypeError("Expected argument 'location_uri' to be a str")
        pulumi.set(__self__, "location_uri", location_uri)
        if server_certificate and not isinstance(server_certificate, str):
            raise TypeError("Expected argument 'server_certificate' to be a str")
        pulumi.set(__self__, "server_certificate", server_certificate)
        if server_port and not isinstance(server_port, int):
            raise TypeError("Expected argument 'server_port' to be a int")
        pulumi.set(__self__, "server_port", server_port)
        if server_protocol and not isinstance(server_protocol, str):
            raise TypeError("Expected argument 'server_protocol' to be a str")
        pulumi.set(__self__, "server_protocol", server_protocol)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="accessKey")
    def access_key(self) -> Optional[str]:
        """
        Optional. The access key is used if credentials are required to access the self-managed object storage server.
        """
        return pulumi.get(self, "access_key")

    @property
    @pulumi.getter(name="agentArns")
    def agent_arns(self) -> Optional[Sequence[str]]:
        """
        The Amazon Resource Name (ARN) of the agents associated with the self-managed object storage server location.
        """
        return pulumi.get(self, "agent_arns")

    @property
    @pulumi.getter(name="locationArn")
    def location_arn(self) -> Optional[str]:
        """
        The Amazon Resource Name (ARN) of the location that is created.
        """
        return pulumi.get(self, "location_arn")

    @property
    @pulumi.getter(name="locationUri")
    def location_uri(self) -> Optional[str]:
        """
        The URL of the object storage location that was described.
        """
        return pulumi.get(self, "location_uri")

    @property
    @pulumi.getter(name="serverCertificate")
    def server_certificate(self) -> Optional[str]:
        """
        X.509 PEM content containing a certificate authority or chain to trust.
        """
        return pulumi.get(self, "server_certificate")

    @property
    @pulumi.getter(name="serverPort")
    def server_port(self) -> Optional[int]:
        """
        The port that your self-managed server accepts inbound network traffic on.
        """
        return pulumi.get(self, "server_port")

    @property
    @pulumi.getter(name="serverProtocol")
    def server_protocol(self) -> Optional['LocationObjectStorageServerProtocol']:
        """
        The protocol that the object storage server uses to communicate.
        """
        return pulumi.get(self, "server_protocol")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['outputs.LocationObjectStorageTag']]:
        """
        An array of key-value pairs to apply to this resource.
        """
        return pulumi.get(self, "tags")


class AwaitableGetLocationObjectStorageResult(GetLocationObjectStorageResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetLocationObjectStorageResult(
            access_key=self.access_key,
            agent_arns=self.agent_arns,
            location_arn=self.location_arn,
            location_uri=self.location_uri,
            server_certificate=self.server_certificate,
            server_port=self.server_port,
            server_protocol=self.server_protocol,
            tags=self.tags)


def get_location_object_storage(location_arn: Optional[str] = None,
                                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetLocationObjectStorageResult:
    """
    Resource schema for AWS::DataSync::LocationObjectStorage.


    :param str location_arn: The Amazon Resource Name (ARN) of the location that is created.
    """
    __args__ = dict()
    __args__['locationArn'] = location_arn
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:datasync:getLocationObjectStorage', __args__, opts=opts, typ=GetLocationObjectStorageResult).value

    return AwaitableGetLocationObjectStorageResult(
        access_key=pulumi.get(__ret__, 'access_key'),
        agent_arns=pulumi.get(__ret__, 'agent_arns'),
        location_arn=pulumi.get(__ret__, 'location_arn'),
        location_uri=pulumi.get(__ret__, 'location_uri'),
        server_certificate=pulumi.get(__ret__, 'server_certificate'),
        server_port=pulumi.get(__ret__, 'server_port'),
        server_protocol=pulumi.get(__ret__, 'server_protocol'),
        tags=pulumi.get(__ret__, 'tags'))


@_utilities.lift_output_func(get_location_object_storage)
def get_location_object_storage_output(location_arn: Optional[pulumi.Input[str]] = None,
                                       opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetLocationObjectStorageResult]:
    """
    Resource schema for AWS::DataSync::LocationObjectStorage.


    :param str location_arn: The Amazon Resource Name (ARN) of the location that is created.
    """
    ...
