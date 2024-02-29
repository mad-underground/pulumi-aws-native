# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from ._enums import *

__all__ = [
    'EnabledBaselineParameter',
    'EnabledControlParameter',
]

@pulumi.output_type
class EnabledBaselineParameter(dict):
    def __init__(__self__, *,
                 key: Optional[str] = None,
                 value: Optional[Any] = None):
        if key is not None:
            pulumi.set(__self__, "key", key)
        if value is not None:
            pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> Optional[str]:
        return pulumi.get(self, "key")

    @property
    @pulumi.getter
    def value(self) -> Optional[Any]:
        return pulumi.get(self, "value")


@pulumi.output_type
class EnabledControlParameter(dict):
    def __init__(__self__, *,
                 key: str,
                 value: Any):
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> str:
        return pulumi.get(self, "key")

    @property
    @pulumi.getter
    def value(self) -> Any:
        return pulumi.get(self, "value")


