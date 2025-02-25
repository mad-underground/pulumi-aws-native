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
from .. import _inputs as _root_inputs
from .. import outputs as _root_outputs
from ._enums import *
from ._inputs import *

__all__ = ['DataSourceArgs', 'DataSource']

@pulumi.input_type
class DataSourceArgs:
    def __init__(__self__, *,
                 alternate_data_source_parameters: Optional[pulumi.Input[Sequence[pulumi.Input['DataSourceParametersArgs']]]] = None,
                 aws_account_id: Optional[pulumi.Input[str]] = None,
                 credentials: Optional[pulumi.Input['DataSourceCredentialsArgs']] = None,
                 data_source_id: Optional[pulumi.Input[str]] = None,
                 data_source_parameters: Optional[pulumi.Input['DataSourceParametersArgs']] = None,
                 error_info: Optional[pulumi.Input['DataSourceErrorInfoArgs']] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 permissions: Optional[pulumi.Input[Sequence[pulumi.Input['DataSourceResourcePermissionArgs']]]] = None,
                 ssl_properties: Optional[pulumi.Input['DataSourceSslPropertiesArgs']] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]] = None,
                 type: Optional[pulumi.Input['DataSourceType']] = None,
                 vpc_connection_properties: Optional[pulumi.Input['DataSourceVpcConnectionPropertiesArgs']] = None):
        """
        The set of arguments for constructing a DataSource resource.
        :param pulumi.Input[Sequence[pulumi.Input['DataSourceParametersArgs']]] alternate_data_source_parameters: <p>A set of alternate data source parameters that you want to share for the credentials
                           stored with this data source. The credentials are applied in tandem with the data source
                           parameters when you copy a data source by using a create or update request. The API
                           operation compares the <code>DataSourceParameters</code> structure that's in the request
                           with the structures in the <code>AlternateDataSourceParameters</code> allow list. If the
                           structures are an exact match, the request is allowed to use the credentials from this
                           existing data source. If the <code>AlternateDataSourceParameters</code> list is null,
                           the <code>Credentials</code> originally used with this <code>DataSourceParameters</code>
                           are automatically allowed.</p>
        :param pulumi.Input[str] name: <p>A display name for the data source.</p>
        :param pulumi.Input[Sequence[pulumi.Input['DataSourceResourcePermissionArgs']]] permissions: <p>A list of resource permissions on the data source.</p>
        :param pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]] tags: <p>Contains a map of the key-value pairs for the resource tag or tags assigned to the data source.</p>
        """
        if alternate_data_source_parameters is not None:
            pulumi.set(__self__, "alternate_data_source_parameters", alternate_data_source_parameters)
        if aws_account_id is not None:
            pulumi.set(__self__, "aws_account_id", aws_account_id)
        if credentials is not None:
            pulumi.set(__self__, "credentials", credentials)
        if data_source_id is not None:
            pulumi.set(__self__, "data_source_id", data_source_id)
        if data_source_parameters is not None:
            pulumi.set(__self__, "data_source_parameters", data_source_parameters)
        if error_info is not None:
            pulumi.set(__self__, "error_info", error_info)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if permissions is not None:
            pulumi.set(__self__, "permissions", permissions)
        if ssl_properties is not None:
            pulumi.set(__self__, "ssl_properties", ssl_properties)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if vpc_connection_properties is not None:
            pulumi.set(__self__, "vpc_connection_properties", vpc_connection_properties)

    @property
    @pulumi.getter(name="alternateDataSourceParameters")
    def alternate_data_source_parameters(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['DataSourceParametersArgs']]]]:
        """
        <p>A set of alternate data source parameters that you want to share for the credentials
                    stored with this data source. The credentials are applied in tandem with the data source
                    parameters when you copy a data source by using a create or update request. The API
                    operation compares the <code>DataSourceParameters</code> structure that's in the request
                    with the structures in the <code>AlternateDataSourceParameters</code> allow list. If the
                    structures are an exact match, the request is allowed to use the credentials from this
                    existing data source. If the <code>AlternateDataSourceParameters</code> list is null,
                    the <code>Credentials</code> originally used with this <code>DataSourceParameters</code>
                    are automatically allowed.</p>
        """
        return pulumi.get(self, "alternate_data_source_parameters")

    @alternate_data_source_parameters.setter
    def alternate_data_source_parameters(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['DataSourceParametersArgs']]]]):
        pulumi.set(self, "alternate_data_source_parameters", value)

    @property
    @pulumi.getter(name="awsAccountId")
    def aws_account_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "aws_account_id")

    @aws_account_id.setter
    def aws_account_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "aws_account_id", value)

    @property
    @pulumi.getter
    def credentials(self) -> Optional[pulumi.Input['DataSourceCredentialsArgs']]:
        return pulumi.get(self, "credentials")

    @credentials.setter
    def credentials(self, value: Optional[pulumi.Input['DataSourceCredentialsArgs']]):
        pulumi.set(self, "credentials", value)

    @property
    @pulumi.getter(name="dataSourceId")
    def data_source_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "data_source_id")

    @data_source_id.setter
    def data_source_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "data_source_id", value)

    @property
    @pulumi.getter(name="dataSourceParameters")
    def data_source_parameters(self) -> Optional[pulumi.Input['DataSourceParametersArgs']]:
        return pulumi.get(self, "data_source_parameters")

    @data_source_parameters.setter
    def data_source_parameters(self, value: Optional[pulumi.Input['DataSourceParametersArgs']]):
        pulumi.set(self, "data_source_parameters", value)

    @property
    @pulumi.getter(name="errorInfo")
    def error_info(self) -> Optional[pulumi.Input['DataSourceErrorInfoArgs']]:
        return pulumi.get(self, "error_info")

    @error_info.setter
    def error_info(self, value: Optional[pulumi.Input['DataSourceErrorInfoArgs']]):
        pulumi.set(self, "error_info", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        <p>A display name for the data source.</p>
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def permissions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['DataSourceResourcePermissionArgs']]]]:
        """
        <p>A list of resource permissions on the data source.</p>
        """
        return pulumi.get(self, "permissions")

    @permissions.setter
    def permissions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['DataSourceResourcePermissionArgs']]]]):
        pulumi.set(self, "permissions", value)

    @property
    @pulumi.getter(name="sslProperties")
    def ssl_properties(self) -> Optional[pulumi.Input['DataSourceSslPropertiesArgs']]:
        return pulumi.get(self, "ssl_properties")

    @ssl_properties.setter
    def ssl_properties(self, value: Optional[pulumi.Input['DataSourceSslPropertiesArgs']]):
        pulumi.set(self, "ssl_properties", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]]:
        """
        <p>Contains a map of the key-value pairs for the resource tag or tags assigned to the data source.</p>
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['_root_inputs.TagArgs']]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input['DataSourceType']]:
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input['DataSourceType']]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="vpcConnectionProperties")
    def vpc_connection_properties(self) -> Optional[pulumi.Input['DataSourceVpcConnectionPropertiesArgs']]:
        return pulumi.get(self, "vpc_connection_properties")

    @vpc_connection_properties.setter
    def vpc_connection_properties(self, value: Optional[pulumi.Input['DataSourceVpcConnectionPropertiesArgs']]):
        pulumi.set(self, "vpc_connection_properties", value)


class DataSource(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 alternate_data_source_parameters: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DataSourceParametersArgs']]]]] = None,
                 aws_account_id: Optional[pulumi.Input[str]] = None,
                 credentials: Optional[pulumi.Input[pulumi.InputType['DataSourceCredentialsArgs']]] = None,
                 data_source_id: Optional[pulumi.Input[str]] = None,
                 data_source_parameters: Optional[pulumi.Input[pulumi.InputType['DataSourceParametersArgs']]] = None,
                 error_info: Optional[pulumi.Input[pulumi.InputType['DataSourceErrorInfoArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 permissions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DataSourceResourcePermissionArgs']]]]] = None,
                 ssl_properties: Optional[pulumi.Input[pulumi.InputType['DataSourceSslPropertiesArgs']]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['_root_inputs.TagArgs']]]]] = None,
                 type: Optional[pulumi.Input['DataSourceType']] = None,
                 vpc_connection_properties: Optional[pulumi.Input[pulumi.InputType['DataSourceVpcConnectionPropertiesArgs']]] = None,
                 __props__=None):
        """
        Definition of the AWS::QuickSight::DataSource Resource Type.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DataSourceParametersArgs']]]] alternate_data_source_parameters: <p>A set of alternate data source parameters that you want to share for the credentials
                           stored with this data source. The credentials are applied in tandem with the data source
                           parameters when you copy a data source by using a create or update request. The API
                           operation compares the <code>DataSourceParameters</code> structure that's in the request
                           with the structures in the <code>AlternateDataSourceParameters</code> allow list. If the
                           structures are an exact match, the request is allowed to use the credentials from this
                           existing data source. If the <code>AlternateDataSourceParameters</code> list is null,
                           the <code>Credentials</code> originally used with this <code>DataSourceParameters</code>
                           are automatically allowed.</p>
        :param pulumi.Input[str] name: <p>A display name for the data source.</p>
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DataSourceResourcePermissionArgs']]]] permissions: <p>A list of resource permissions on the data source.</p>
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['_root_inputs.TagArgs']]]] tags: <p>Contains a map of the key-value pairs for the resource tag or tags assigned to the data source.</p>
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[DataSourceArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Definition of the AWS::QuickSight::DataSource Resource Type.

        :param str resource_name: The name of the resource.
        :param DataSourceArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DataSourceArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 alternate_data_source_parameters: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DataSourceParametersArgs']]]]] = None,
                 aws_account_id: Optional[pulumi.Input[str]] = None,
                 credentials: Optional[pulumi.Input[pulumi.InputType['DataSourceCredentialsArgs']]] = None,
                 data_source_id: Optional[pulumi.Input[str]] = None,
                 data_source_parameters: Optional[pulumi.Input[pulumi.InputType['DataSourceParametersArgs']]] = None,
                 error_info: Optional[pulumi.Input[pulumi.InputType['DataSourceErrorInfoArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 permissions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DataSourceResourcePermissionArgs']]]]] = None,
                 ssl_properties: Optional[pulumi.Input[pulumi.InputType['DataSourceSslPropertiesArgs']]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['_root_inputs.TagArgs']]]]] = None,
                 type: Optional[pulumi.Input['DataSourceType']] = None,
                 vpc_connection_properties: Optional[pulumi.Input[pulumi.InputType['DataSourceVpcConnectionPropertiesArgs']]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DataSourceArgs.__new__(DataSourceArgs)

            __props__.__dict__["alternate_data_source_parameters"] = alternate_data_source_parameters
            __props__.__dict__["aws_account_id"] = aws_account_id
            __props__.__dict__["credentials"] = credentials
            __props__.__dict__["data_source_id"] = data_source_id
            __props__.__dict__["data_source_parameters"] = data_source_parameters
            __props__.__dict__["error_info"] = error_info
            __props__.__dict__["name"] = name
            __props__.__dict__["permissions"] = permissions
            __props__.__dict__["ssl_properties"] = ssl_properties
            __props__.__dict__["tags"] = tags
            __props__.__dict__["type"] = type
            __props__.__dict__["vpc_connection_properties"] = vpc_connection_properties
            __props__.__dict__["arn"] = None
            __props__.__dict__["created_time"] = None
            __props__.__dict__["last_updated_time"] = None
            __props__.__dict__["status"] = None
        replace_on_changes = pulumi.ResourceOptions(replace_on_changes=["aws_account_id", "data_source_id", "type"])
        opts = pulumi.ResourceOptions.merge(opts, replace_on_changes)
        super(DataSource, __self__).__init__(
            'aws-native:quicksight:DataSource',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'DataSource':
        """
        Get an existing DataSource resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = DataSourceArgs.__new__(DataSourceArgs)

        __props__.__dict__["alternate_data_source_parameters"] = None
        __props__.__dict__["arn"] = None
        __props__.__dict__["aws_account_id"] = None
        __props__.__dict__["created_time"] = None
        __props__.__dict__["credentials"] = None
        __props__.__dict__["data_source_id"] = None
        __props__.__dict__["data_source_parameters"] = None
        __props__.__dict__["error_info"] = None
        __props__.__dict__["last_updated_time"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["permissions"] = None
        __props__.__dict__["ssl_properties"] = None
        __props__.__dict__["status"] = None
        __props__.__dict__["tags"] = None
        __props__.__dict__["type"] = None
        __props__.__dict__["vpc_connection_properties"] = None
        return DataSource(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="alternateDataSourceParameters")
    def alternate_data_source_parameters(self) -> pulumi.Output[Optional[Sequence['outputs.DataSourceParameters']]]:
        """
        <p>A set of alternate data source parameters that you want to share for the credentials
                    stored with this data source. The credentials are applied in tandem with the data source
                    parameters when you copy a data source by using a create or update request. The API
                    operation compares the <code>DataSourceParameters</code> structure that's in the request
                    with the structures in the <code>AlternateDataSourceParameters</code> allow list. If the
                    structures are an exact match, the request is allowed to use the credentials from this
                    existing data source. If the <code>AlternateDataSourceParameters</code> list is null,
                    the <code>Credentials</code> originally used with this <code>DataSourceParameters</code>
                    are automatically allowed.</p>
        """
        return pulumi.get(self, "alternate_data_source_parameters")

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        <p>The Amazon Resource Name (ARN) of the data source.</p>
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="awsAccountId")
    def aws_account_id(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "aws_account_id")

    @property
    @pulumi.getter(name="createdTime")
    def created_time(self) -> pulumi.Output[str]:
        """
        <p>The time that this data source was created.</p>
        """
        return pulumi.get(self, "created_time")

    @property
    @pulumi.getter
    def credentials(self) -> pulumi.Output[Optional['outputs.DataSourceCredentials']]:
        return pulumi.get(self, "credentials")

    @property
    @pulumi.getter(name="dataSourceId")
    def data_source_id(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "data_source_id")

    @property
    @pulumi.getter(name="dataSourceParameters")
    def data_source_parameters(self) -> pulumi.Output[Optional['outputs.DataSourceParameters']]:
        return pulumi.get(self, "data_source_parameters")

    @property
    @pulumi.getter(name="errorInfo")
    def error_info(self) -> pulumi.Output[Optional['outputs.DataSourceErrorInfo']]:
        return pulumi.get(self, "error_info")

    @property
    @pulumi.getter(name="lastUpdatedTime")
    def last_updated_time(self) -> pulumi.Output[str]:
        """
        <p>The last time that this data source was updated.</p>
        """
        return pulumi.get(self, "last_updated_time")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[Optional[str]]:
        """
        <p>A display name for the data source.</p>
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def permissions(self) -> pulumi.Output[Optional[Sequence['outputs.DataSourceResourcePermission']]]:
        """
        <p>A list of resource permissions on the data source.</p>
        """
        return pulumi.get(self, "permissions")

    @property
    @pulumi.getter(name="sslProperties")
    def ssl_properties(self) -> pulumi.Output[Optional['outputs.DataSourceSslProperties']]:
        return pulumi.get(self, "ssl_properties")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output['DataSourceResourceStatus']:
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence['_root_outputs.Tag']]]:
        """
        <p>Contains a map of the key-value pairs for the resource tag or tags assigned to the data source.</p>
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[Optional['DataSourceType']]:
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="vpcConnectionProperties")
    def vpc_connection_properties(self) -> pulumi.Output[Optional['outputs.DataSourceVpcConnectionProperties']]:
        return pulumi.get(self, "vpc_connection_properties")

