# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from .. import outputs as _root_outputs
from ._enums import *

__all__ = [
    'GetAgreementResult',
    'AwaitableGetAgreementResult',
    'get_agreement',
    'get_agreement_output',
]

@pulumi.output_type
class GetAgreementResult:
    def __init__(__self__, access_role=None, agreement_id=None, arn=None, base_directory=None, description=None, local_profile_id=None, partner_profile_id=None, status=None, tags=None):
        if access_role and not isinstance(access_role, str):
            raise TypeError("Expected argument 'access_role' to be a str")
        pulumi.set(__self__, "access_role", access_role)
        if agreement_id and not isinstance(agreement_id, str):
            raise TypeError("Expected argument 'agreement_id' to be a str")
        pulumi.set(__self__, "agreement_id", agreement_id)
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        pulumi.set(__self__, "arn", arn)
        if base_directory and not isinstance(base_directory, str):
            raise TypeError("Expected argument 'base_directory' to be a str")
        pulumi.set(__self__, "base_directory", base_directory)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if local_profile_id and not isinstance(local_profile_id, str):
            raise TypeError("Expected argument 'local_profile_id' to be a str")
        pulumi.set(__self__, "local_profile_id", local_profile_id)
        if partner_profile_id and not isinstance(partner_profile_id, str):
            raise TypeError("Expected argument 'partner_profile_id' to be a str")
        pulumi.set(__self__, "partner_profile_id", partner_profile_id)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="accessRole")
    def access_role(self) -> Optional[str]:
        """
        Specifies the access role for the agreement.
        """
        return pulumi.get(self, "access_role")

    @property
    @pulumi.getter(name="agreementId")
    def agreement_id(self) -> Optional[str]:
        """
        A unique identifier for the agreement.
        """
        return pulumi.get(self, "agreement_id")

    @property
    @pulumi.getter
    def arn(self) -> Optional[str]:
        """
        Specifies the unique Amazon Resource Name (ARN) for the agreement.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="baseDirectory")
    def base_directory(self) -> Optional[str]:
        """
        Specifies the base directory for the agreement.
        """
        return pulumi.get(self, "base_directory")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        A textual description for the agreement.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="localProfileId")
    def local_profile_id(self) -> Optional[str]:
        """
        A unique identifier for the local profile.
        """
        return pulumi.get(self, "local_profile_id")

    @property
    @pulumi.getter(name="partnerProfileId")
    def partner_profile_id(self) -> Optional[str]:
        """
        A unique identifier for the partner profile.
        """
        return pulumi.get(self, "partner_profile_id")

    @property
    @pulumi.getter
    def status(self) -> Optional['AgreementStatus']:
        """
        Specifies the status of the agreement.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['_root_outputs.Tag']]:
        """
        Key-value pairs that can be used to group and search for agreements. Tags are metadata attached to agreements for any purpose.
        """
        return pulumi.get(self, "tags")


class AwaitableGetAgreementResult(GetAgreementResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAgreementResult(
            access_role=self.access_role,
            agreement_id=self.agreement_id,
            arn=self.arn,
            base_directory=self.base_directory,
            description=self.description,
            local_profile_id=self.local_profile_id,
            partner_profile_id=self.partner_profile_id,
            status=self.status,
            tags=self.tags)


def get_agreement(agreement_id: Optional[str] = None,
                  server_id: Optional[str] = None,
                  opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAgreementResult:
    """
    Resource Type definition for AWS::Transfer::Agreement


    :param str agreement_id: A unique identifier for the agreement.
    :param str server_id: A unique identifier for the server.
    """
    __args__ = dict()
    __args__['agreementId'] = agreement_id
    __args__['serverId'] = server_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:transfer:getAgreement', __args__, opts=opts, typ=GetAgreementResult).value

    return AwaitableGetAgreementResult(
        access_role=pulumi.get(__ret__, 'access_role'),
        agreement_id=pulumi.get(__ret__, 'agreement_id'),
        arn=pulumi.get(__ret__, 'arn'),
        base_directory=pulumi.get(__ret__, 'base_directory'),
        description=pulumi.get(__ret__, 'description'),
        local_profile_id=pulumi.get(__ret__, 'local_profile_id'),
        partner_profile_id=pulumi.get(__ret__, 'partner_profile_id'),
        status=pulumi.get(__ret__, 'status'),
        tags=pulumi.get(__ret__, 'tags'))


@_utilities.lift_output_func(get_agreement)
def get_agreement_output(agreement_id: Optional[pulumi.Input[str]] = None,
                         server_id: Optional[pulumi.Input[str]] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetAgreementResult]:
    """
    Resource Type definition for AWS::Transfer::Agreement


    :param str agreement_id: A unique identifier for the agreement.
    :param str server_id: A unique identifier for the server.
    """
    ...
