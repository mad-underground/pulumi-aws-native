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

__all__ = [
    'GetVdmAttributesResult',
    'AwaitableGetVdmAttributesResult',
    'get_vdm_attributes',
    'get_vdm_attributes_output',
]

@pulumi.output_type
class GetVdmAttributesResult:
    def __init__(__self__, dashboard_attributes=None, guardian_attributes=None, vdm_attributes_resource_id=None):
        if dashboard_attributes and not isinstance(dashboard_attributes, dict):
            raise TypeError("Expected argument 'dashboard_attributes' to be a dict")
        pulumi.set(__self__, "dashboard_attributes", dashboard_attributes)
        if guardian_attributes and not isinstance(guardian_attributes, dict):
            raise TypeError("Expected argument 'guardian_attributes' to be a dict")
        pulumi.set(__self__, "guardian_attributes", guardian_attributes)
        if vdm_attributes_resource_id and not isinstance(vdm_attributes_resource_id, str):
            raise TypeError("Expected argument 'vdm_attributes_resource_id' to be a str")
        pulumi.set(__self__, "vdm_attributes_resource_id", vdm_attributes_resource_id)

    @property
    @pulumi.getter(name="dashboardAttributes")
    def dashboard_attributes(self) -> Optional['outputs.VdmAttributesDashboardAttributes']:
        return pulumi.get(self, "dashboard_attributes")

    @property
    @pulumi.getter(name="guardianAttributes")
    def guardian_attributes(self) -> Optional['outputs.VdmAttributesGuardianAttributes']:
        return pulumi.get(self, "guardian_attributes")

    @property
    @pulumi.getter(name="vdmAttributesResourceId")
    def vdm_attributes_resource_id(self) -> Optional[str]:
        """
        Unique identifier for this resource
        """
        return pulumi.get(self, "vdm_attributes_resource_id")


class AwaitableGetVdmAttributesResult(GetVdmAttributesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetVdmAttributesResult(
            dashboard_attributes=self.dashboard_attributes,
            guardian_attributes=self.guardian_attributes,
            vdm_attributes_resource_id=self.vdm_attributes_resource_id)


def get_vdm_attributes(vdm_attributes_resource_id: Optional[str] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetVdmAttributesResult:
    """
    Resource Type definition for AWS::SES::VdmAttributes


    :param str vdm_attributes_resource_id: Unique identifier for this resource
    """
    __args__ = dict()
    __args__['vdmAttributesResourceId'] = vdm_attributes_resource_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:ses:getVdmAttributes', __args__, opts=opts, typ=GetVdmAttributesResult).value

    return AwaitableGetVdmAttributesResult(
        dashboard_attributes=pulumi.get(__ret__, 'dashboard_attributes'),
        guardian_attributes=pulumi.get(__ret__, 'guardian_attributes'),
        vdm_attributes_resource_id=pulumi.get(__ret__, 'vdm_attributes_resource_id'))


@_utilities.lift_output_func(get_vdm_attributes)
def get_vdm_attributes_output(vdm_attributes_resource_id: Optional[pulumi.Input[str]] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetVdmAttributesResult]:
    """
    Resource Type definition for AWS::SES::VdmAttributes


    :param str vdm_attributes_resource_id: Unique identifier for this resource
    """
    ...
