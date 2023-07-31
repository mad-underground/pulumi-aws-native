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
    'GetImageVersionResult',
    'AwaitableGetImageVersionResult',
    'get_image_version',
    'get_image_version_output',
]

@pulumi.output_type
class GetImageVersionResult:
    def __init__(__self__, container_image=None, horovod=None, image_arn=None, image_version_arn=None, job_type=None, m_l_framework=None, processor=None, programming_lang=None, release_notes=None, vendor_guidance=None, version=None):
        if container_image and not isinstance(container_image, str):
            raise TypeError("Expected argument 'container_image' to be a str")
        pulumi.set(__self__, "container_image", container_image)
        if horovod and not isinstance(horovod, bool):
            raise TypeError("Expected argument 'horovod' to be a bool")
        pulumi.set(__self__, "horovod", horovod)
        if image_arn and not isinstance(image_arn, str):
            raise TypeError("Expected argument 'image_arn' to be a str")
        pulumi.set(__self__, "image_arn", image_arn)
        if image_version_arn and not isinstance(image_version_arn, str):
            raise TypeError("Expected argument 'image_version_arn' to be a str")
        pulumi.set(__self__, "image_version_arn", image_version_arn)
        if job_type and not isinstance(job_type, str):
            raise TypeError("Expected argument 'job_type' to be a str")
        pulumi.set(__self__, "job_type", job_type)
        if m_l_framework and not isinstance(m_l_framework, str):
            raise TypeError("Expected argument 'm_l_framework' to be a str")
        pulumi.set(__self__, "m_l_framework", m_l_framework)
        if processor and not isinstance(processor, str):
            raise TypeError("Expected argument 'processor' to be a str")
        pulumi.set(__self__, "processor", processor)
        if programming_lang and not isinstance(programming_lang, str):
            raise TypeError("Expected argument 'programming_lang' to be a str")
        pulumi.set(__self__, "programming_lang", programming_lang)
        if release_notes and not isinstance(release_notes, str):
            raise TypeError("Expected argument 'release_notes' to be a str")
        pulumi.set(__self__, "release_notes", release_notes)
        if vendor_guidance and not isinstance(vendor_guidance, str):
            raise TypeError("Expected argument 'vendor_guidance' to be a str")
        pulumi.set(__self__, "vendor_guidance", vendor_guidance)
        if version and not isinstance(version, int):
            raise TypeError("Expected argument 'version' to be a int")
        pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter(name="containerImage")
    def container_image(self) -> Optional[str]:
        return pulumi.get(self, "container_image")

    @property
    @pulumi.getter
    def horovod(self) -> Optional[bool]:
        return pulumi.get(self, "horovod")

    @property
    @pulumi.getter(name="imageArn")
    def image_arn(self) -> Optional[str]:
        return pulumi.get(self, "image_arn")

    @property
    @pulumi.getter(name="imageVersionArn")
    def image_version_arn(self) -> Optional[str]:
        return pulumi.get(self, "image_version_arn")

    @property
    @pulumi.getter(name="jobType")
    def job_type(self) -> Optional['ImageVersionJobType']:
        return pulumi.get(self, "job_type")

    @property
    @pulumi.getter(name="mLFramework")
    def m_l_framework(self) -> Optional[str]:
        return pulumi.get(self, "m_l_framework")

    @property
    @pulumi.getter
    def processor(self) -> Optional['ImageVersionProcessor']:
        return pulumi.get(self, "processor")

    @property
    @pulumi.getter(name="programmingLang")
    def programming_lang(self) -> Optional[str]:
        return pulumi.get(self, "programming_lang")

    @property
    @pulumi.getter(name="releaseNotes")
    def release_notes(self) -> Optional[str]:
        return pulumi.get(self, "release_notes")

    @property
    @pulumi.getter(name="vendorGuidance")
    def vendor_guidance(self) -> Optional['ImageVersionVendorGuidance']:
        return pulumi.get(self, "vendor_guidance")

    @property
    @pulumi.getter
    def version(self) -> Optional[int]:
        return pulumi.get(self, "version")


class AwaitableGetImageVersionResult(GetImageVersionResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetImageVersionResult(
            container_image=self.container_image,
            horovod=self.horovod,
            image_arn=self.image_arn,
            image_version_arn=self.image_version_arn,
            job_type=self.job_type,
            m_l_framework=self.m_l_framework,
            processor=self.processor,
            programming_lang=self.programming_lang,
            release_notes=self.release_notes,
            vendor_guidance=self.vendor_guidance,
            version=self.version)


def get_image_version(image_version_arn: Optional[str] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetImageVersionResult:
    """
    Resource Type definition for AWS::SageMaker::ImageVersion
    """
    __args__ = dict()
    __args__['imageVersionArn'] = image_version_arn
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:sagemaker:getImageVersion', __args__, opts=opts, typ=GetImageVersionResult).value

    return AwaitableGetImageVersionResult(
        container_image=pulumi.get(__ret__, 'container_image'),
        horovod=pulumi.get(__ret__, 'horovod'),
        image_arn=pulumi.get(__ret__, 'image_arn'),
        image_version_arn=pulumi.get(__ret__, 'image_version_arn'),
        job_type=pulumi.get(__ret__, 'job_type'),
        m_l_framework=pulumi.get(__ret__, 'm_l_framework'),
        processor=pulumi.get(__ret__, 'processor'),
        programming_lang=pulumi.get(__ret__, 'programming_lang'),
        release_notes=pulumi.get(__ret__, 'release_notes'),
        vendor_guidance=pulumi.get(__ret__, 'vendor_guidance'),
        version=pulumi.get(__ret__, 'version'))


@_utilities.lift_output_func(get_image_version)
def get_image_version_output(image_version_arn: Optional[pulumi.Input[str]] = None,
                             opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetImageVersionResult]:
    """
    Resource Type definition for AWS::SageMaker::ImageVersion
    """
    ...
