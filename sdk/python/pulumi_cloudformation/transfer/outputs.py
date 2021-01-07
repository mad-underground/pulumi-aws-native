# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables
from . import outputs
from .. import outputs as _root_outputs

__all__ = [
    'ServerAttributes',
    'ServerEndpointDetails',
    'ServerIdentityProviderDetails',
    'ServerProperties',
    'ServerProtocol',
    'UserAttributes',
    'UserHomeDirectoryMapEntry',
    'UserProperties',
    'UserSshPublicKey',
]

@pulumi.output_type
class ServerAttributes(dict):
    def __init__(__self__, *,
                 arn: str,
                 server_id: str):
        pulumi.set(__self__, "arn", arn)
        pulumi.set(__self__, "server_id", server_id)

    @property
    @pulumi.getter(name="Arn")
    def arn(self) -> str:
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="ServerId")
    def server_id(self) -> str:
        return pulumi.get(self, "server_id")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ServerEndpointDetails(dict):
    """
    http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-server-endpointdetails.html
    """
    def __init__(__self__, *,
                 address_allocation_ids: Optional[Sequence[str]] = None,
                 security_group_ids: Optional[Sequence[str]] = None,
                 subnet_ids: Optional[Sequence[str]] = None,
                 vpc_endpoint_id: Optional[str] = None,
                 vpc_id: Optional[str] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-server-endpointdetails.html
        :param Sequence[str] address_allocation_ids: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-server-endpointdetails.html#cfn-transfer-server-endpointdetails-addressallocationids
        :param Sequence[str] security_group_ids: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-server-endpointdetails.html#cfn-transfer-server-endpointdetails-securitygroupids
        :param Sequence[str] subnet_ids: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-server-endpointdetails.html#cfn-transfer-server-endpointdetails-subnetids
        :param str vpc_endpoint_id: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-server-endpointdetails.html#cfn-transfer-server-endpointdetails-vpcendpointid
        :param str vpc_id: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-server-endpointdetails.html#cfn-transfer-server-endpointdetails-vpcid
        """
        if address_allocation_ids is not None:
            pulumi.set(__self__, "address_allocation_ids", address_allocation_ids)
        if security_group_ids is not None:
            pulumi.set(__self__, "security_group_ids", security_group_ids)
        if subnet_ids is not None:
            pulumi.set(__self__, "subnet_ids", subnet_ids)
        if vpc_endpoint_id is not None:
            pulumi.set(__self__, "vpc_endpoint_id", vpc_endpoint_id)
        if vpc_id is not None:
            pulumi.set(__self__, "vpc_id", vpc_id)

    @property
    @pulumi.getter(name="AddressAllocationIds")
    def address_allocation_ids(self) -> Optional[Sequence[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-server-endpointdetails.html#cfn-transfer-server-endpointdetails-addressallocationids
        """
        return pulumi.get(self, "address_allocation_ids")

    @property
    @pulumi.getter(name="SecurityGroupIds")
    def security_group_ids(self) -> Optional[Sequence[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-server-endpointdetails.html#cfn-transfer-server-endpointdetails-securitygroupids
        """
        return pulumi.get(self, "security_group_ids")

    @property
    @pulumi.getter(name="SubnetIds")
    def subnet_ids(self) -> Optional[Sequence[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-server-endpointdetails.html#cfn-transfer-server-endpointdetails-subnetids
        """
        return pulumi.get(self, "subnet_ids")

    @property
    @pulumi.getter(name="VpcEndpointId")
    def vpc_endpoint_id(self) -> Optional[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-server-endpointdetails.html#cfn-transfer-server-endpointdetails-vpcendpointid
        """
        return pulumi.get(self, "vpc_endpoint_id")

    @property
    @pulumi.getter(name="VpcId")
    def vpc_id(self) -> Optional[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-server-endpointdetails.html#cfn-transfer-server-endpointdetails-vpcid
        """
        return pulumi.get(self, "vpc_id")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ServerIdentityProviderDetails(dict):
    """
    http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-server-identityproviderdetails.html
    """
    def __init__(__self__, *,
                 invocation_role: str,
                 url: str):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-server-identityproviderdetails.html
        :param str invocation_role: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-server-identityproviderdetails.html#cfn-transfer-server-identityproviderdetails-invocationrole
        :param str url: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-server-identityproviderdetails.html#cfn-transfer-server-identityproviderdetails-url
        """
        pulumi.set(__self__, "invocation_role", invocation_role)
        pulumi.set(__self__, "url", url)

    @property
    @pulumi.getter(name="InvocationRole")
    def invocation_role(self) -> str:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-server-identityproviderdetails.html#cfn-transfer-server-identityproviderdetails-invocationrole
        """
        return pulumi.get(self, "invocation_role")

    @property
    @pulumi.getter(name="Url")
    def url(self) -> str:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-server-identityproviderdetails.html#cfn-transfer-server-identityproviderdetails-url
        """
        return pulumi.get(self, "url")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ServerProperties(dict):
    """
    http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transfer-server.html
    """
    def __init__(__self__, *,
                 certificate: Optional[str] = None,
                 endpoint_details: Optional['outputs.ServerEndpointDetails'] = None,
                 endpoint_type: Optional[str] = None,
                 identity_provider_details: Optional['outputs.ServerIdentityProviderDetails'] = None,
                 identity_provider_type: Optional[str] = None,
                 logging_role: Optional[str] = None,
                 protocols: Optional[Sequence['outputs.ServerProtocol']] = None,
                 security_policy_name: Optional[str] = None,
                 tags: Optional[Sequence['_root_outputs.Tag']] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transfer-server.html
        :param str certificate: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transfer-server.html#cfn-transfer-server-certificate
        :param 'ServerEndpointDetailsArgs' endpoint_details: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transfer-server.html#cfn-transfer-server-endpointdetails
        :param str endpoint_type: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transfer-server.html#cfn-transfer-server-endpointtype
        :param 'ServerIdentityProviderDetailsArgs' identity_provider_details: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transfer-server.html#cfn-transfer-server-identityproviderdetails
        :param str identity_provider_type: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transfer-server.html#cfn-transfer-server-identityprovidertype
        :param str logging_role: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transfer-server.html#cfn-transfer-server-loggingrole
        :param Sequence['ServerProtocolArgs'] protocols: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transfer-server.html#cfn-transfer-server-protocols
        :param str security_policy_name: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transfer-server.html#cfn-transfer-server-securitypolicyname
        :param Sequence['_root_inputs.TagArgs'] tags: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transfer-server.html#cfn-transfer-server-tags
        """
        if certificate is not None:
            pulumi.set(__self__, "certificate", certificate)
        if endpoint_details is not None:
            pulumi.set(__self__, "endpoint_details", endpoint_details)
        if endpoint_type is not None:
            pulumi.set(__self__, "endpoint_type", endpoint_type)
        if identity_provider_details is not None:
            pulumi.set(__self__, "identity_provider_details", identity_provider_details)
        if identity_provider_type is not None:
            pulumi.set(__self__, "identity_provider_type", identity_provider_type)
        if logging_role is not None:
            pulumi.set(__self__, "logging_role", logging_role)
        if protocols is not None:
            pulumi.set(__self__, "protocols", protocols)
        if security_policy_name is not None:
            pulumi.set(__self__, "security_policy_name", security_policy_name)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="Certificate")
    def certificate(self) -> Optional[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transfer-server.html#cfn-transfer-server-certificate
        """
        return pulumi.get(self, "certificate")

    @property
    @pulumi.getter(name="EndpointDetails")
    def endpoint_details(self) -> Optional['outputs.ServerEndpointDetails']:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transfer-server.html#cfn-transfer-server-endpointdetails
        """
        return pulumi.get(self, "endpoint_details")

    @property
    @pulumi.getter(name="EndpointType")
    def endpoint_type(self) -> Optional[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transfer-server.html#cfn-transfer-server-endpointtype
        """
        return pulumi.get(self, "endpoint_type")

    @property
    @pulumi.getter(name="IdentityProviderDetails")
    def identity_provider_details(self) -> Optional['outputs.ServerIdentityProviderDetails']:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transfer-server.html#cfn-transfer-server-identityproviderdetails
        """
        return pulumi.get(self, "identity_provider_details")

    @property
    @pulumi.getter(name="IdentityProviderType")
    def identity_provider_type(self) -> Optional[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transfer-server.html#cfn-transfer-server-identityprovidertype
        """
        return pulumi.get(self, "identity_provider_type")

    @property
    @pulumi.getter(name="LoggingRole")
    def logging_role(self) -> Optional[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transfer-server.html#cfn-transfer-server-loggingrole
        """
        return pulumi.get(self, "logging_role")

    @property
    @pulumi.getter(name="Protocols")
    def protocols(self) -> Optional[Sequence['outputs.ServerProtocol']]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transfer-server.html#cfn-transfer-server-protocols
        """
        return pulumi.get(self, "protocols")

    @property
    @pulumi.getter(name="SecurityPolicyName")
    def security_policy_name(self) -> Optional[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transfer-server.html#cfn-transfer-server-securitypolicyname
        """
        return pulumi.get(self, "security_policy_name")

    @property
    @pulumi.getter(name="Tags")
    def tags(self) -> Optional[Sequence['_root_outputs.Tag']]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transfer-server.html#cfn-transfer-server-tags
        """
        return pulumi.get(self, "tags")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ServerProtocol(dict):
    def __init__(__self__):
        pass

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class UserAttributes(dict):
    def __init__(__self__, *,
                 arn: str,
                 server_id: str,
                 user_name: str):
        pulumi.set(__self__, "arn", arn)
        pulumi.set(__self__, "server_id", server_id)
        pulumi.set(__self__, "user_name", user_name)

    @property
    @pulumi.getter(name="Arn")
    def arn(self) -> str:
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="ServerId")
    def server_id(self) -> str:
        return pulumi.get(self, "server_id")

    @property
    @pulumi.getter(name="UserName")
    def user_name(self) -> str:
        return pulumi.get(self, "user_name")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class UserHomeDirectoryMapEntry(dict):
    """
    http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-user-homedirectorymapentry.html
    """
    def __init__(__self__, *,
                 entry: str,
                 target: str):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-user-homedirectorymapentry.html
        :param str entry: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-user-homedirectorymapentry.html#cfn-transfer-user-homedirectorymapentry-entry
        :param str target: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-user-homedirectorymapentry.html#cfn-transfer-user-homedirectorymapentry-target
        """
        pulumi.set(__self__, "entry", entry)
        pulumi.set(__self__, "target", target)

    @property
    @pulumi.getter(name="Entry")
    def entry(self) -> str:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-user-homedirectorymapentry.html#cfn-transfer-user-homedirectorymapentry-entry
        """
        return pulumi.get(self, "entry")

    @property
    @pulumi.getter(name="Target")
    def target(self) -> str:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-user-homedirectorymapentry.html#cfn-transfer-user-homedirectorymapentry-target
        """
        return pulumi.get(self, "target")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class UserProperties(dict):
    """
    http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transfer-user.html
    """
    def __init__(__self__, *,
                 role: str,
                 server_id: str,
                 user_name: str,
                 home_directory: Optional[str] = None,
                 home_directory_mappings: Optional[Sequence['outputs.UserHomeDirectoryMapEntry']] = None,
                 home_directory_type: Optional[str] = None,
                 policy: Optional[str] = None,
                 ssh_public_keys: Optional[Sequence['outputs.UserSshPublicKey']] = None,
                 tags: Optional[Sequence['_root_outputs.Tag']] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transfer-user.html
        :param str role: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transfer-user.html#cfn-transfer-user-role
        :param str server_id: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transfer-user.html#cfn-transfer-user-serverid
        :param str user_name: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transfer-user.html#cfn-transfer-user-username
        :param str home_directory: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transfer-user.html#cfn-transfer-user-homedirectory
        :param Sequence['UserHomeDirectoryMapEntryArgs'] home_directory_mappings: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transfer-user.html#cfn-transfer-user-homedirectorymappings
        :param str home_directory_type: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transfer-user.html#cfn-transfer-user-homedirectorytype
        :param str policy: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transfer-user.html#cfn-transfer-user-policy
        :param Sequence['UserSshPublicKeyArgs'] ssh_public_keys: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transfer-user.html#cfn-transfer-user-sshpublickeys
        :param Sequence['_root_inputs.TagArgs'] tags: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transfer-user.html#cfn-transfer-user-tags
        """
        pulumi.set(__self__, "role", role)
        pulumi.set(__self__, "server_id", server_id)
        pulumi.set(__self__, "user_name", user_name)
        if home_directory is not None:
            pulumi.set(__self__, "home_directory", home_directory)
        if home_directory_mappings is not None:
            pulumi.set(__self__, "home_directory_mappings", home_directory_mappings)
        if home_directory_type is not None:
            pulumi.set(__self__, "home_directory_type", home_directory_type)
        if policy is not None:
            pulumi.set(__self__, "policy", policy)
        if ssh_public_keys is not None:
            pulumi.set(__self__, "ssh_public_keys", ssh_public_keys)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="Role")
    def role(self) -> str:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transfer-user.html#cfn-transfer-user-role
        """
        return pulumi.get(self, "role")

    @property
    @pulumi.getter(name="ServerId")
    def server_id(self) -> str:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transfer-user.html#cfn-transfer-user-serverid
        """
        return pulumi.get(self, "server_id")

    @property
    @pulumi.getter(name="UserName")
    def user_name(self) -> str:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transfer-user.html#cfn-transfer-user-username
        """
        return pulumi.get(self, "user_name")

    @property
    @pulumi.getter(name="HomeDirectory")
    def home_directory(self) -> Optional[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transfer-user.html#cfn-transfer-user-homedirectory
        """
        return pulumi.get(self, "home_directory")

    @property
    @pulumi.getter(name="HomeDirectoryMappings")
    def home_directory_mappings(self) -> Optional[Sequence['outputs.UserHomeDirectoryMapEntry']]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transfer-user.html#cfn-transfer-user-homedirectorymappings
        """
        return pulumi.get(self, "home_directory_mappings")

    @property
    @pulumi.getter(name="HomeDirectoryType")
    def home_directory_type(self) -> Optional[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transfer-user.html#cfn-transfer-user-homedirectorytype
        """
        return pulumi.get(self, "home_directory_type")

    @property
    @pulumi.getter(name="Policy")
    def policy(self) -> Optional[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transfer-user.html#cfn-transfer-user-policy
        """
        return pulumi.get(self, "policy")

    @property
    @pulumi.getter(name="SshPublicKeys")
    def ssh_public_keys(self) -> Optional[Sequence['outputs.UserSshPublicKey']]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transfer-user.html#cfn-transfer-user-sshpublickeys
        """
        return pulumi.get(self, "ssh_public_keys")

    @property
    @pulumi.getter(name="Tags")
    def tags(self) -> Optional[Sequence['_root_outputs.Tag']]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transfer-user.html#cfn-transfer-user-tags
        """
        return pulumi.get(self, "tags")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class UserSshPublicKey(dict):
    def __init__(__self__):
        pass

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


