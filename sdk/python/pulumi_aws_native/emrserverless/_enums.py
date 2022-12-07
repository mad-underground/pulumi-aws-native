# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from enum import Enum

__all__ = [
    'ApplicationArchitecture',
]


class ApplicationArchitecture(str, Enum):
    """
    The cpu architecture of an application.
    """
    ARM64 = "ARM64"
    X8664 = "X86_64"
