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
    'GetThingResult',
    'AwaitableGetThingResult',
    'get_thing',
    'get_thing_output',
]

@pulumi.output_type
class GetThingResult:
    def __init__(__self__, arn=None, attribute_payload=None, id=None):
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        pulumi.set(__self__, "arn", arn)
        if attribute_payload and not isinstance(attribute_payload, dict):
            raise TypeError("Expected argument 'attribute_payload' to be a dict")
        pulumi.set(__self__, "attribute_payload", attribute_payload)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)

    @property
    @pulumi.getter
    def arn(self) -> Optional[str]:
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="attributePayload")
    def attribute_payload(self) -> Optional['outputs.ThingAttributePayload']:
        return pulumi.get(self, "attribute_payload")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        return pulumi.get(self, "id")


class AwaitableGetThingResult(GetThingResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetThingResult(
            arn=self.arn,
            attribute_payload=self.attribute_payload,
            id=self.id)


def get_thing(thing_name: Optional[str] = None,
              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetThingResult:
    """
    Resource Type definition for AWS::IoT::Thing
    """
    __args__ = dict()
    __args__['thingName'] = thing_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:iot:getThing', __args__, opts=opts, typ=GetThingResult).value

    return AwaitableGetThingResult(
        arn=pulumi.get(__ret__, 'arn'),
        attribute_payload=pulumi.get(__ret__, 'attribute_payload'),
        id=pulumi.get(__ret__, 'id'))


@_utilities.lift_output_func(get_thing)
def get_thing_output(thing_name: Optional[pulumi.Input[str]] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetThingResult]:
    """
    Resource Type definition for AWS::IoT::Thing
    """
    ...
