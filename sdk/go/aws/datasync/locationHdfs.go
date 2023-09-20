// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package datasync

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Resource schema for AWS::DataSync::LocationHDFS.
type LocationHdfs struct {
	pulumi.CustomResourceState

	// ARN(s) of the agent(s) to use for an HDFS location.
	AgentArns pulumi.StringArrayOutput `pulumi:"agentArns"`
	// The authentication mode used to determine identity of user.
	AuthenticationType LocationHdfsAuthenticationTypeOutput `pulumi:"authenticationType"`
	// Size of chunks (blocks) in bytes that the data is divided into when stored in the HDFS cluster.
	BlockSize pulumi.IntPtrOutput `pulumi:"blockSize"`
	// The Base64 string representation of the Keytab file.
	KerberosKeytab pulumi.StringPtrOutput `pulumi:"kerberosKeytab"`
	// The string representation of the Krb5Conf file, or the presigned URL to access the Krb5.conf file within an S3 bucket.
	KerberosKrb5Conf pulumi.StringPtrOutput `pulumi:"kerberosKrb5Conf"`
	// The unique identity, or principal, to which Kerberos can assign tickets.
	KerberosPrincipal pulumi.StringPtrOutput `pulumi:"kerberosPrincipal"`
	// The identifier for the Key Management Server where the encryption keys that encrypt data inside HDFS clusters are stored.
	KmsKeyProviderUri pulumi.StringPtrOutput `pulumi:"kmsKeyProviderUri"`
	// The Amazon Resource Name (ARN) of the HDFS location.
	LocationArn pulumi.StringOutput `pulumi:"locationArn"`
	// The URL of the HDFS location that was described.
	LocationUri pulumi.StringOutput `pulumi:"locationUri"`
	// An array of Name Node(s) of the HDFS location.
	NameNodes        LocationHdfsNameNodeArrayOutput       `pulumi:"nameNodes"`
	QopConfiguration LocationHdfsQopConfigurationPtrOutput `pulumi:"qopConfiguration"`
	// Number of copies of each block that exists inside the HDFS cluster.
	ReplicationFactor pulumi.IntPtrOutput `pulumi:"replicationFactor"`
	// The user name that has read and write permissions on the specified HDFS cluster.
	SimpleUser pulumi.StringPtrOutput `pulumi:"simpleUser"`
	// The subdirectory in HDFS that is used to read data from the HDFS source location or write data to the HDFS destination.
	Subdirectory pulumi.StringPtrOutput `pulumi:"subdirectory"`
	// An array of key-value pairs to apply to this resource.
	Tags LocationHdfsTagArrayOutput `pulumi:"tags"`
}

// NewLocationHdfs registers a new resource with the given unique name, arguments, and options.
func NewLocationHdfs(ctx *pulumi.Context,
	name string, args *LocationHdfsArgs, opts ...pulumi.ResourceOption) (*LocationHdfs, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AgentArns == nil {
		return nil, errors.New("invalid value for required argument 'AgentArns'")
	}
	if args.AuthenticationType == nil {
		return nil, errors.New("invalid value for required argument 'AuthenticationType'")
	}
	if args.NameNodes == nil {
		return nil, errors.New("invalid value for required argument 'NameNodes'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource LocationHdfs
	err := ctx.RegisterResource("aws-native:datasync:LocationHdfs", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLocationHdfs gets an existing LocationHdfs resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLocationHdfs(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LocationHdfsState, opts ...pulumi.ResourceOption) (*LocationHdfs, error) {
	var resource LocationHdfs
	err := ctx.ReadResource("aws-native:datasync:LocationHdfs", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LocationHdfs resources.
type locationHdfsState struct {
}

type LocationHdfsState struct {
}

func (LocationHdfsState) ElementType() reflect.Type {
	return reflect.TypeOf((*locationHdfsState)(nil)).Elem()
}

type locationHdfsArgs struct {
	// ARN(s) of the agent(s) to use for an HDFS location.
	AgentArns []string `pulumi:"agentArns"`
	// The authentication mode used to determine identity of user.
	AuthenticationType LocationHdfsAuthenticationType `pulumi:"authenticationType"`
	// Size of chunks (blocks) in bytes that the data is divided into when stored in the HDFS cluster.
	BlockSize *int `pulumi:"blockSize"`
	// The Base64 string representation of the Keytab file.
	KerberosKeytab *string `pulumi:"kerberosKeytab"`
	// The string representation of the Krb5Conf file, or the presigned URL to access the Krb5.conf file within an S3 bucket.
	KerberosKrb5Conf *string `pulumi:"kerberosKrb5Conf"`
	// The unique identity, or principal, to which Kerberos can assign tickets.
	KerberosPrincipal *string `pulumi:"kerberosPrincipal"`
	// The identifier for the Key Management Server where the encryption keys that encrypt data inside HDFS clusters are stored.
	KmsKeyProviderUri *string `pulumi:"kmsKeyProviderUri"`
	// An array of Name Node(s) of the HDFS location.
	NameNodes        []LocationHdfsNameNode        `pulumi:"nameNodes"`
	QopConfiguration *LocationHdfsQopConfiguration `pulumi:"qopConfiguration"`
	// Number of copies of each block that exists inside the HDFS cluster.
	ReplicationFactor *int `pulumi:"replicationFactor"`
	// The user name that has read and write permissions on the specified HDFS cluster.
	SimpleUser *string `pulumi:"simpleUser"`
	// The subdirectory in HDFS that is used to read data from the HDFS source location or write data to the HDFS destination.
	Subdirectory *string `pulumi:"subdirectory"`
	// An array of key-value pairs to apply to this resource.
	Tags []LocationHdfsTag `pulumi:"tags"`
}

// The set of arguments for constructing a LocationHdfs resource.
type LocationHdfsArgs struct {
	// ARN(s) of the agent(s) to use for an HDFS location.
	AgentArns pulumi.StringArrayInput
	// The authentication mode used to determine identity of user.
	AuthenticationType LocationHdfsAuthenticationTypeInput
	// Size of chunks (blocks) in bytes that the data is divided into when stored in the HDFS cluster.
	BlockSize pulumi.IntPtrInput
	// The Base64 string representation of the Keytab file.
	KerberosKeytab pulumi.StringPtrInput
	// The string representation of the Krb5Conf file, or the presigned URL to access the Krb5.conf file within an S3 bucket.
	KerberosKrb5Conf pulumi.StringPtrInput
	// The unique identity, or principal, to which Kerberos can assign tickets.
	KerberosPrincipal pulumi.StringPtrInput
	// The identifier for the Key Management Server where the encryption keys that encrypt data inside HDFS clusters are stored.
	KmsKeyProviderUri pulumi.StringPtrInput
	// An array of Name Node(s) of the HDFS location.
	NameNodes        LocationHdfsNameNodeArrayInput
	QopConfiguration LocationHdfsQopConfigurationPtrInput
	// Number of copies of each block that exists inside the HDFS cluster.
	ReplicationFactor pulumi.IntPtrInput
	// The user name that has read and write permissions on the specified HDFS cluster.
	SimpleUser pulumi.StringPtrInput
	// The subdirectory in HDFS that is used to read data from the HDFS source location or write data to the HDFS destination.
	Subdirectory pulumi.StringPtrInput
	// An array of key-value pairs to apply to this resource.
	Tags LocationHdfsTagArrayInput
}

func (LocationHdfsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*locationHdfsArgs)(nil)).Elem()
}

type LocationHdfsInput interface {
	pulumi.Input

	ToLocationHdfsOutput() LocationHdfsOutput
	ToLocationHdfsOutputWithContext(ctx context.Context) LocationHdfsOutput
}

func (*LocationHdfs) ElementType() reflect.Type {
	return reflect.TypeOf((**LocationHdfs)(nil)).Elem()
}

func (i *LocationHdfs) ToLocationHdfsOutput() LocationHdfsOutput {
	return i.ToLocationHdfsOutputWithContext(context.Background())
}

func (i *LocationHdfs) ToLocationHdfsOutputWithContext(ctx context.Context) LocationHdfsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocationHdfsOutput)
}

func (i *LocationHdfs) ToOutput(ctx context.Context) pulumix.Output[*LocationHdfs] {
	return pulumix.Output[*LocationHdfs]{
		OutputState: i.ToLocationHdfsOutputWithContext(ctx).OutputState,
	}
}

type LocationHdfsOutput struct{ *pulumi.OutputState }

func (LocationHdfsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LocationHdfs)(nil)).Elem()
}

func (o LocationHdfsOutput) ToLocationHdfsOutput() LocationHdfsOutput {
	return o
}

func (o LocationHdfsOutput) ToLocationHdfsOutputWithContext(ctx context.Context) LocationHdfsOutput {
	return o
}

func (o LocationHdfsOutput) ToOutput(ctx context.Context) pulumix.Output[*LocationHdfs] {
	return pulumix.Output[*LocationHdfs]{
		OutputState: o.OutputState,
	}
}

// ARN(s) of the agent(s) to use for an HDFS location.
func (o LocationHdfsOutput) AgentArns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *LocationHdfs) pulumi.StringArrayOutput { return v.AgentArns }).(pulumi.StringArrayOutput)
}

// The authentication mode used to determine identity of user.
func (o LocationHdfsOutput) AuthenticationType() LocationHdfsAuthenticationTypeOutput {
	return o.ApplyT(func(v *LocationHdfs) LocationHdfsAuthenticationTypeOutput { return v.AuthenticationType }).(LocationHdfsAuthenticationTypeOutput)
}

// Size of chunks (blocks) in bytes that the data is divided into when stored in the HDFS cluster.
func (o LocationHdfsOutput) BlockSize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *LocationHdfs) pulumi.IntPtrOutput { return v.BlockSize }).(pulumi.IntPtrOutput)
}

// The Base64 string representation of the Keytab file.
func (o LocationHdfsOutput) KerberosKeytab() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocationHdfs) pulumi.StringPtrOutput { return v.KerberosKeytab }).(pulumi.StringPtrOutput)
}

// The string representation of the Krb5Conf file, or the presigned URL to access the Krb5.conf file within an S3 bucket.
func (o LocationHdfsOutput) KerberosKrb5Conf() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocationHdfs) pulumi.StringPtrOutput { return v.KerberosKrb5Conf }).(pulumi.StringPtrOutput)
}

// The unique identity, or principal, to which Kerberos can assign tickets.
func (o LocationHdfsOutput) KerberosPrincipal() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocationHdfs) pulumi.StringPtrOutput { return v.KerberosPrincipal }).(pulumi.StringPtrOutput)
}

// The identifier for the Key Management Server where the encryption keys that encrypt data inside HDFS clusters are stored.
func (o LocationHdfsOutput) KmsKeyProviderUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocationHdfs) pulumi.StringPtrOutput { return v.KmsKeyProviderUri }).(pulumi.StringPtrOutput)
}

// The Amazon Resource Name (ARN) of the HDFS location.
func (o LocationHdfsOutput) LocationArn() pulumi.StringOutput {
	return o.ApplyT(func(v *LocationHdfs) pulumi.StringOutput { return v.LocationArn }).(pulumi.StringOutput)
}

// The URL of the HDFS location that was described.
func (o LocationHdfsOutput) LocationUri() pulumi.StringOutput {
	return o.ApplyT(func(v *LocationHdfs) pulumi.StringOutput { return v.LocationUri }).(pulumi.StringOutput)
}

// An array of Name Node(s) of the HDFS location.
func (o LocationHdfsOutput) NameNodes() LocationHdfsNameNodeArrayOutput {
	return o.ApplyT(func(v *LocationHdfs) LocationHdfsNameNodeArrayOutput { return v.NameNodes }).(LocationHdfsNameNodeArrayOutput)
}

func (o LocationHdfsOutput) QopConfiguration() LocationHdfsQopConfigurationPtrOutput {
	return o.ApplyT(func(v *LocationHdfs) LocationHdfsQopConfigurationPtrOutput { return v.QopConfiguration }).(LocationHdfsQopConfigurationPtrOutput)
}

// Number of copies of each block that exists inside the HDFS cluster.
func (o LocationHdfsOutput) ReplicationFactor() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *LocationHdfs) pulumi.IntPtrOutput { return v.ReplicationFactor }).(pulumi.IntPtrOutput)
}

// The user name that has read and write permissions on the specified HDFS cluster.
func (o LocationHdfsOutput) SimpleUser() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocationHdfs) pulumi.StringPtrOutput { return v.SimpleUser }).(pulumi.StringPtrOutput)
}

// The subdirectory in HDFS that is used to read data from the HDFS source location or write data to the HDFS destination.
func (o LocationHdfsOutput) Subdirectory() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocationHdfs) pulumi.StringPtrOutput { return v.Subdirectory }).(pulumi.StringPtrOutput)
}

// An array of key-value pairs to apply to this resource.
func (o LocationHdfsOutput) Tags() LocationHdfsTagArrayOutput {
	return o.ApplyT(func(v *LocationHdfs) LocationHdfsTagArrayOutput { return v.Tags }).(LocationHdfsTagArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LocationHdfsInput)(nil)).Elem(), &LocationHdfs{})
	pulumi.RegisterOutputType(LocationHdfsOutput{})
}