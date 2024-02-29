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
from .. import outputs as _root_outputs

__all__ = [
    'GetEnabledBaselineResult',
    'AwaitableGetEnabledBaselineResult',
    'get_enabled_baseline',
    'get_enabled_baseline_output',
]

@pulumi.output_type
class GetEnabledBaselineResult:
    def __init__(__self__, baseline_version=None, enabled_baseline_identifier=None, parameters=None, tags=None):
        if baseline_version and not isinstance(baseline_version, str):
            raise TypeError("Expected argument 'baseline_version' to be a str")
        pulumi.set(__self__, "baseline_version", baseline_version)
        if enabled_baseline_identifier and not isinstance(enabled_baseline_identifier, str):
            raise TypeError("Expected argument 'enabled_baseline_identifier' to be a str")
        pulumi.set(__self__, "enabled_baseline_identifier", enabled_baseline_identifier)
        if parameters and not isinstance(parameters, list):
            raise TypeError("Expected argument 'parameters' to be a list")
        pulumi.set(__self__, "parameters", parameters)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="baselineVersion")
    def baseline_version(self) -> Optional[str]:
        return pulumi.get(self, "baseline_version")

    @property
    @pulumi.getter(name="enabledBaselineIdentifier")
    def enabled_baseline_identifier(self) -> Optional[str]:
        return pulumi.get(self, "enabled_baseline_identifier")

    @property
    @pulumi.getter
    def parameters(self) -> Optional[Sequence['outputs.EnabledBaselineParameter']]:
        return pulumi.get(self, "parameters")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['_root_outputs.Tag']]:
        return pulumi.get(self, "tags")


class AwaitableGetEnabledBaselineResult(GetEnabledBaselineResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetEnabledBaselineResult(
            baseline_version=self.baseline_version,
            enabled_baseline_identifier=self.enabled_baseline_identifier,
            parameters=self.parameters,
            tags=self.tags)


def get_enabled_baseline(enabled_baseline_identifier: Optional[str] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetEnabledBaselineResult:
    """
    Definition of AWS::ControlTower::EnabledBaseline Resource Type
    """
    __args__ = dict()
    __args__['enabledBaselineIdentifier'] = enabled_baseline_identifier
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:controltower:getEnabledBaseline', __args__, opts=opts, typ=GetEnabledBaselineResult).value

    return AwaitableGetEnabledBaselineResult(
        baseline_version=pulumi.get(__ret__, 'baseline_version'),
        enabled_baseline_identifier=pulumi.get(__ret__, 'enabled_baseline_identifier'),
        parameters=pulumi.get(__ret__, 'parameters'),
        tags=pulumi.get(__ret__, 'tags'))


@_utilities.lift_output_func(get_enabled_baseline)
def get_enabled_baseline_output(enabled_baseline_identifier: Optional[pulumi.Input[str]] = None,
                                opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetEnabledBaselineResult]:
    """
    Definition of AWS::ControlTower::EnabledBaseline Resource Type
    """
    ...
