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
    'GetDocumentResult',
    'AwaitableGetDocumentResult',
    'get_document',
    'get_document_output',
]

@pulumi.output_type
class GetDocumentResult:
    def __init__(__self__, content=None, document_format=None, requires=None, tags=None, target_type=None, version_name=None):
        if content and not isinstance(content, dict):
            raise TypeError("Expected argument 'content' to be a dict")
        pulumi.set(__self__, "content", content)
        if document_format and not isinstance(document_format, str):
            raise TypeError("Expected argument 'document_format' to be a str")
        pulumi.set(__self__, "document_format", document_format)
        if requires and not isinstance(requires, list):
            raise TypeError("Expected argument 'requires' to be a list")
        pulumi.set(__self__, "requires", requires)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)
        if target_type and not isinstance(target_type, str):
            raise TypeError("Expected argument 'target_type' to be a str")
        pulumi.set(__self__, "target_type", target_type)
        if version_name and not isinstance(version_name, str):
            raise TypeError("Expected argument 'version_name' to be a str")
        pulumi.set(__self__, "version_name", version_name)

    @property
    @pulumi.getter
    def content(self) -> Optional[Any]:
        """
        The content for the Systems Manager document in JSON, YAML or String format.
        """
        return pulumi.get(self, "content")

    @property
    @pulumi.getter(name="documentFormat")
    def document_format(self) -> Optional['DocumentFormat']:
        """
        Specify the document format for the request. The document format can be either JSON or YAML. JSON is the default format.
        """
        return pulumi.get(self, "document_format")

    @property
    @pulumi.getter
    def requires(self) -> Optional[Sequence['outputs.DocumentRequires']]:
        """
        A list of SSM documents required by a document. For example, an ApplicationConfiguration document requires an ApplicationConfigurationSchema document.
        """
        return pulumi.get(self, "requires")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['outputs.DocumentTag']]:
        """
        Optional metadata that you assign to a resource. Tags enable you to categorize a resource in different ways, such as by purpose, owner, or environment.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="targetType")
    def target_type(self) -> Optional[str]:
        """
        Specify a target type to define the kinds of resources the document can run on.
        """
        return pulumi.get(self, "target_type")

    @property
    @pulumi.getter(name="versionName")
    def version_name(self) -> Optional[str]:
        """
        An optional field specifying the version of the artifact you are creating with the document. This value is unique across all versions of a document, and cannot be changed.
        """
        return pulumi.get(self, "version_name")


class AwaitableGetDocumentResult(GetDocumentResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDocumentResult(
            content=self.content,
            document_format=self.document_format,
            requires=self.requires,
            tags=self.tags,
            target_type=self.target_type,
            version_name=self.version_name)


def get_document(name: Optional[str] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDocumentResult:
    """
    The AWS::SSM::Document resource is an SSM document in AWS Systems Manager that defines the actions that Systems Manager performs, which can be used to set up and run commands on your instances.


    :param str name: A name for the Systems Manager document.
    """
    __args__ = dict()
    __args__['name'] = name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('aws-native:ssm:getDocument', __args__, opts=opts, typ=GetDocumentResult).value

    return AwaitableGetDocumentResult(
        content=pulumi.get(__ret__, 'content'),
        document_format=pulumi.get(__ret__, 'document_format'),
        requires=pulumi.get(__ret__, 'requires'),
        tags=pulumi.get(__ret__, 'tags'),
        target_type=pulumi.get(__ret__, 'target_type'),
        version_name=pulumi.get(__ret__, 'version_name'))


@_utilities.lift_output_func(get_document)
def get_document_output(name: Optional[pulumi.Input[str]] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetDocumentResult]:
    """
    The AWS::SSM::Document resource is an SSM document in AWS Systems Manager that defines the actions that Systems Manager performs, which can be used to set up and run commands on your instances.


    :param str name: A name for the Systems Manager document.
    """
    ...
