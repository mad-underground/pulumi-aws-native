// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package dms

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::DMS::Endpoint
//
// Deprecated: Endpoint is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type Endpoint struct {
	pulumi.CustomResourceState

	CertificateArn             pulumi.StringPtrOutput                      `pulumi:"certificateArn"`
	DatabaseName               pulumi.StringPtrOutput                      `pulumi:"databaseName"`
	DocDbSettings              EndpointDocDbSettingsPtrOutput              `pulumi:"docDbSettings"`
	DynamoDbSettings           EndpointDynamoDbSettingsPtrOutput           `pulumi:"dynamoDbSettings"`
	ElasticsearchSettings      EndpointElasticsearchSettingsPtrOutput      `pulumi:"elasticsearchSettings"`
	EndpointIdentifier         pulumi.StringPtrOutput                      `pulumi:"endpointIdentifier"`
	EndpointType               pulumi.StringOutput                         `pulumi:"endpointType"`
	EngineName                 pulumi.StringOutput                         `pulumi:"engineName"`
	ExternalId                 pulumi.StringOutput                         `pulumi:"externalId"`
	ExtraConnectionAttributes  pulumi.StringPtrOutput                      `pulumi:"extraConnectionAttributes"`
	IbmDb2Settings             EndpointIbmDb2SettingsPtrOutput             `pulumi:"ibmDb2Settings"`
	KafkaSettings              EndpointKafkaSettingsPtrOutput              `pulumi:"kafkaSettings"`
	KinesisSettings            EndpointKinesisSettingsPtrOutput            `pulumi:"kinesisSettings"`
	KmsKeyId                   pulumi.StringPtrOutput                      `pulumi:"kmsKeyId"`
	MicrosoftSqlServerSettings EndpointMicrosoftSqlServerSettingsPtrOutput `pulumi:"microsoftSqlServerSettings"`
	MongoDbSettings            EndpointMongoDbSettingsPtrOutput            `pulumi:"mongoDbSettings"`
	MySqlSettings              EndpointMySqlSettingsPtrOutput              `pulumi:"mySqlSettings"`
	NeptuneSettings            EndpointNeptuneSettingsPtrOutput            `pulumi:"neptuneSettings"`
	OracleSettings             EndpointOracleSettingsPtrOutput             `pulumi:"oracleSettings"`
	Password                   pulumi.StringPtrOutput                      `pulumi:"password"`
	Port                       pulumi.IntPtrOutput                         `pulumi:"port"`
	PostgreSqlSettings         EndpointPostgreSqlSettingsPtrOutput         `pulumi:"postgreSqlSettings"`
	RedisSettings              EndpointRedisSettingsPtrOutput              `pulumi:"redisSettings"`
	RedshiftSettings           EndpointRedshiftSettingsPtrOutput           `pulumi:"redshiftSettings"`
	ResourceIdentifier         pulumi.StringPtrOutput                      `pulumi:"resourceIdentifier"`
	S3Settings                 EndpointS3SettingsPtrOutput                 `pulumi:"s3Settings"`
	ServerName                 pulumi.StringPtrOutput                      `pulumi:"serverName"`
	SslMode                    pulumi.StringPtrOutput                      `pulumi:"sslMode"`
	SybaseSettings             EndpointSybaseSettingsPtrOutput             `pulumi:"sybaseSettings"`
	Tags                       EndpointTagArrayOutput                      `pulumi:"tags"`
	Username                   pulumi.StringPtrOutput                      `pulumi:"username"`
}

// NewEndpoint registers a new resource with the given unique name, arguments, and options.
func NewEndpoint(ctx *pulumi.Context,
	name string, args *EndpointArgs, opts ...pulumi.ResourceOption) (*Endpoint, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EndpointType == nil {
		return nil, errors.New("invalid value for required argument 'EndpointType'")
	}
	if args.EngineName == nil {
		return nil, errors.New("invalid value for required argument 'EngineName'")
	}
	var resource Endpoint
	err := ctx.RegisterResource("aws-native:dms:Endpoint", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEndpoint gets an existing Endpoint resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEndpoint(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EndpointState, opts ...pulumi.ResourceOption) (*Endpoint, error) {
	var resource Endpoint
	err := ctx.ReadResource("aws-native:dms:Endpoint", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Endpoint resources.
type endpointState struct {
}

type EndpointState struct {
}

func (EndpointState) ElementType() reflect.Type {
	return reflect.TypeOf((*endpointState)(nil)).Elem()
}

type endpointArgs struct {
	CertificateArn             *string                             `pulumi:"certificateArn"`
	DatabaseName               *string                             `pulumi:"databaseName"`
	DocDbSettings              *EndpointDocDbSettings              `pulumi:"docDbSettings"`
	DynamoDbSettings           *EndpointDynamoDbSettings           `pulumi:"dynamoDbSettings"`
	ElasticsearchSettings      *EndpointElasticsearchSettings      `pulumi:"elasticsearchSettings"`
	EndpointIdentifier         *string                             `pulumi:"endpointIdentifier"`
	EndpointType               string                              `pulumi:"endpointType"`
	EngineName                 string                              `pulumi:"engineName"`
	ExtraConnectionAttributes  *string                             `pulumi:"extraConnectionAttributes"`
	IbmDb2Settings             *EndpointIbmDb2Settings             `pulumi:"ibmDb2Settings"`
	KafkaSettings              *EndpointKafkaSettings              `pulumi:"kafkaSettings"`
	KinesisSettings            *EndpointKinesisSettings            `pulumi:"kinesisSettings"`
	KmsKeyId                   *string                             `pulumi:"kmsKeyId"`
	MicrosoftSqlServerSettings *EndpointMicrosoftSqlServerSettings `pulumi:"microsoftSqlServerSettings"`
	MongoDbSettings            *EndpointMongoDbSettings            `pulumi:"mongoDbSettings"`
	MySqlSettings              *EndpointMySqlSettings              `pulumi:"mySqlSettings"`
	NeptuneSettings            *EndpointNeptuneSettings            `pulumi:"neptuneSettings"`
	OracleSettings             *EndpointOracleSettings             `pulumi:"oracleSettings"`
	Password                   *string                             `pulumi:"password"`
	Port                       *int                                `pulumi:"port"`
	PostgreSqlSettings         *EndpointPostgreSqlSettings         `pulumi:"postgreSqlSettings"`
	RedisSettings              *EndpointRedisSettings              `pulumi:"redisSettings"`
	RedshiftSettings           *EndpointRedshiftSettings           `pulumi:"redshiftSettings"`
	ResourceIdentifier         *string                             `pulumi:"resourceIdentifier"`
	S3Settings                 *EndpointS3Settings                 `pulumi:"s3Settings"`
	ServerName                 *string                             `pulumi:"serverName"`
	SslMode                    *string                             `pulumi:"sslMode"`
	SybaseSettings             *EndpointSybaseSettings             `pulumi:"sybaseSettings"`
	Tags                       []EndpointTag                       `pulumi:"tags"`
	Username                   *string                             `pulumi:"username"`
}

// The set of arguments for constructing a Endpoint resource.
type EndpointArgs struct {
	CertificateArn             pulumi.StringPtrInput
	DatabaseName               pulumi.StringPtrInput
	DocDbSettings              EndpointDocDbSettingsPtrInput
	DynamoDbSettings           EndpointDynamoDbSettingsPtrInput
	ElasticsearchSettings      EndpointElasticsearchSettingsPtrInput
	EndpointIdentifier         pulumi.StringPtrInput
	EndpointType               pulumi.StringInput
	EngineName                 pulumi.StringInput
	ExtraConnectionAttributes  pulumi.StringPtrInput
	IbmDb2Settings             EndpointIbmDb2SettingsPtrInput
	KafkaSettings              EndpointKafkaSettingsPtrInput
	KinesisSettings            EndpointKinesisSettingsPtrInput
	KmsKeyId                   pulumi.StringPtrInput
	MicrosoftSqlServerSettings EndpointMicrosoftSqlServerSettingsPtrInput
	MongoDbSettings            EndpointMongoDbSettingsPtrInput
	MySqlSettings              EndpointMySqlSettingsPtrInput
	NeptuneSettings            EndpointNeptuneSettingsPtrInput
	OracleSettings             EndpointOracleSettingsPtrInput
	Password                   pulumi.StringPtrInput
	Port                       pulumi.IntPtrInput
	PostgreSqlSettings         EndpointPostgreSqlSettingsPtrInput
	RedisSettings              EndpointRedisSettingsPtrInput
	RedshiftSettings           EndpointRedshiftSettingsPtrInput
	ResourceIdentifier         pulumi.StringPtrInput
	S3Settings                 EndpointS3SettingsPtrInput
	ServerName                 pulumi.StringPtrInput
	SslMode                    pulumi.StringPtrInput
	SybaseSettings             EndpointSybaseSettingsPtrInput
	Tags                       EndpointTagArrayInput
	Username                   pulumi.StringPtrInput
}

func (EndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*endpointArgs)(nil)).Elem()
}

type EndpointInput interface {
	pulumi.Input

	ToEndpointOutput() EndpointOutput
	ToEndpointOutputWithContext(ctx context.Context) EndpointOutput
}

func (*Endpoint) ElementType() reflect.Type {
	return reflect.TypeOf((*Endpoint)(nil))
}

func (i *Endpoint) ToEndpointOutput() EndpointOutput {
	return i.ToEndpointOutputWithContext(context.Background())
}

func (i *Endpoint) ToEndpointOutputWithContext(ctx context.Context) EndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointOutput)
}

type EndpointOutput struct{ *pulumi.OutputState }

func (EndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Endpoint)(nil))
}

func (o EndpointOutput) ToEndpointOutput() EndpointOutput {
	return o
}

func (o EndpointOutput) ToEndpointOutputWithContext(ctx context.Context) EndpointOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(EndpointOutput{})
}
