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
    'GetMlTransformResult',
    'AwaitableGetMlTransformResult',
    'get_ml_transform',
    'get_ml_transform_output',
]

@pulumi.output_type
class GetMlTransformResult:
    def __init__(__self__, description=None, glue_version=None, id=None, max_capacity=None, max_retries=None, name=None, number_of_workers=None, role=None, tags=None, timeout=None, transform_encryption=None, transform_parameters=None, worker_type=None):
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if glue_version and not isinstance(glue_version, str):
            raise TypeError("Expected argument 'glue_version' to be a str")
        pulumi.set(__self__, "glue_version", glue_version)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if max_capacity and not isinstance(max_capacity, float):
            raise TypeError("Expected argument 'max_capacity' to be a float")
        pulumi.set(__self__, "max_capacity", max_capacity)
        if max_retries and not isinstance(max_retries, int):
            raise TypeError("Expected argument 'max_retries' to be a int")
        pulumi.set(__self__, "max_retries", max_retries)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if number_of_workers and not isinstance(number_of_workers, int):
            raise TypeError("Expected argument 'number_of_workers' to be a int")
        pulumi.set(__self__, "number_of_workers", number_of_workers)
        if role and not isinstance(role, str):
            raise TypeError("Expected argument 'role' to be a str")
        pulumi.set(__self__, "role", role)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if timeout and not isinstance(timeout, int):
            raise TypeError("Expected argument 'timeout' to be a int")
        pulumi.set(__self__, "timeout", timeout)
        if transform_encryption and not isinstance(transform_encryption, dict):
            raise TypeError("Expected argument 'transform_encryption' to be a dict")
        pulumi.set(__self__, "transform_encryption", transform_encryption)
        if transform_parameters and not isinstance(transform_parameters, dict):
            raise TypeError("Expected argument 'transform_parameters' to be a dict")
        pulumi.set(__self__, "transform_parameters", transform_parameters)
        if worker_type and not isinstance(worker_type, str):
            raise TypeError("Expected argument 'worker_type' to be a str")
        pulumi.set(__self__, "worker_type", worker_type)

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="glueVersion")
    def glue_version(self) -> Optional[str]:
        return pulumi.get(self, "glue_version")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="maxCapacity")
    def max_capacity(self) -> Optional[float]:
        return pulumi.get(self, "max_capacity")

    @property
    @pulumi.getter(name="maxRetries")
    def max_retries(self) -> Optional[int]:
        return pulumi.get(self, "max_retries")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="numberOfWorkers")
    def number_of_workers(self) -> Optional[int]:
        return pulumi.get(self, "number_of_workers")

    @property
    @pulumi.getter
    def role(self) -> Optional[str]:
        return pulumi.get(self, "role")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Any]:
        """
        Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::Glue::MLTransform` for more information about the expected schema for this property.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def timeout(self) -> Optional[int]:
        return pulumi.get(self, "timeout")

    @property
    @pulumi.getter(name="transformEncryption")
    def transform_encryption(self) -> Optional['outputs.MlTransformTransformEncryption']:
        return pulumi.get(self, "transform_encryption")

    @property
    @pulumi.getter(name="transformParameters")
    def transform_parameters(self) -> Optional['outputs.MlTransformTransformParameters']:
        return pulumi.get(self, "transform_parameters")

    @property
    @pulumi.getter(name="workerType")
    def worker_type(self) -> Optional[str]:
        return pulumi.get(self, "worker_type")


class AwaitableGetMlTransformResult(GetMlTransformResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetMlTransformResult(
            description=self.description,
            glue_version=self.glue_version,
            id=self.id,
            max_capacity=self.max_capacity,
            max_retries=self.max_retries,
            name=self.name,
            number_of_workers=self.number_of_workers,
            role=self.role,
            tags=self.tags,
            timeout=self.timeout,
            transform_encryption=self.transform_encryption,
            transform_parameters=self.transform_parameters,
            worker_type=self.worker_type)


def get_ml_transform(id: Optional[str] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetMlTransformResult:
    """
    Resource Type definition for AWS::Glue::MLTransform
    """
    __args__ = dict()
    __args__['id'] = id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:glue:getMlTransform', __args__, opts=opts, typ=GetMlTransformResult).value

    return AwaitableGetMlTransformResult(
        description=pulumi.get(__ret__, 'description'),
        glue_version=pulumi.get(__ret__, 'glue_version'),
        id=pulumi.get(__ret__, 'id'),
        max_capacity=pulumi.get(__ret__, 'max_capacity'),
        max_retries=pulumi.get(__ret__, 'max_retries'),
        name=pulumi.get(__ret__, 'name'),
        number_of_workers=pulumi.get(__ret__, 'number_of_workers'),
        role=pulumi.get(__ret__, 'role'),
        tags=pulumi.get(__ret__, 'tags'),
        timeout=pulumi.get(__ret__, 'timeout'),
        transform_encryption=pulumi.get(__ret__, 'transform_encryption'),
        transform_parameters=pulumi.get(__ret__, 'transform_parameters'),
        worker_type=pulumi.get(__ret__, 'worker_type'))


@_utilities.lift_output_func(get_ml_transform)
def get_ml_transform_output(id: Optional[pulumi.Input[str]] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetMlTransformResult]:
    """
    Resource Type definition for AWS::Glue::MLTransform
    """
    ...
