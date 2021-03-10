# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables
from .. import _inputs as _root_inputs
from .. import outputs as _root_outputs

__all__ = ['Project']


class Project(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dataset_name: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 recipe_name: Optional[pulumi.Input[str]] = None,
                 role_arn: Optional[pulumi.Input[str]] = None,
                 sample: Optional[pulumi.Input[Union[Any, str]]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['_root_inputs.TagArgs']]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-databrew-project.html

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] dataset_name: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-databrew-project.html#cfn-databrew-project-datasetname
        :param pulumi.Input[str] name: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-databrew-project.html#cfn-databrew-project-name
        :param pulumi.Input[str] recipe_name: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-databrew-project.html#cfn-databrew-project-recipename
        :param pulumi.Input[str] role_arn: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-databrew-project.html#cfn-databrew-project-rolearn
        :param pulumi.Input[Union[Any, str]] sample: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-databrew-project.html#cfn-databrew-project-sample
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['_root_inputs.TagArgs']]]] tags: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-databrew-project.html#cfn-databrew-project-tags
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            if dataset_name is None and not opts.urn:
                raise TypeError("Missing required property 'dataset_name'")
            __props__['dataset_name'] = dataset_name
            if name is None and not opts.urn:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            if recipe_name is None and not opts.urn:
                raise TypeError("Missing required property 'recipe_name'")
            __props__['recipe_name'] = recipe_name
            if role_arn is None and not opts.urn:
                raise TypeError("Missing required property 'role_arn'")
            __props__['role_arn'] = role_arn
            __props__['sample'] = sample
            __props__['tags'] = tags
        super(Project, __self__).__init__(
            'aws-native:DataBrew:Project',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Project':
        """
        Get an existing Project resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return Project(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="DatasetName")
    def dataset_name(self) -> pulumi.Output[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-databrew-project.html#cfn-databrew-project-datasetname
        """
        return pulumi.get(self, "dataset_name")

    @property
    @pulumi.getter(name="Name")
    def name(self) -> pulumi.Output[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-databrew-project.html#cfn-databrew-project-name
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="RecipeName")
    def recipe_name(self) -> pulumi.Output[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-databrew-project.html#cfn-databrew-project-recipename
        """
        return pulumi.get(self, "recipe_name")

    @property
    @pulumi.getter(name="RoleArn")
    def role_arn(self) -> pulumi.Output[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-databrew-project.html#cfn-databrew-project-rolearn
        """
        return pulumi.get(self, "role_arn")

    @property
    @pulumi.getter(name="Sample")
    def sample(self) -> pulumi.Output[Optional[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-databrew-project.html#cfn-databrew-project-sample
        """
        return pulumi.get(self, "sample")

    @property
    @pulumi.getter(name="Tags")
    def tags(self) -> pulumi.Output[Optional[Sequence['_root_outputs.Tag']]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-databrew-project.html#cfn-databrew-project-tags
        """
        return pulumi.get(self, "tags")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

