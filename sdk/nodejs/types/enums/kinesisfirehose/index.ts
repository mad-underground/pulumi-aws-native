// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const DeliveryStreamAmazonOpenSearchServerlessDestinationConfigurationS3BackupMode = {
    FailedDocumentsOnly: "FailedDocumentsOnly",
    AllDocuments: "AllDocuments",
} as const;

export type DeliveryStreamAmazonOpenSearchServerlessDestinationConfigurationS3BackupMode = (typeof DeliveryStreamAmazonOpenSearchServerlessDestinationConfigurationS3BackupMode)[keyof typeof DeliveryStreamAmazonOpenSearchServerlessDestinationConfigurationS3BackupMode];

export const DeliveryStreamAmazonopensearchserviceDestinationConfigurationIndexRotationPeriod = {
    NoRotation: "NoRotation",
    OneHour: "OneHour",
    OneDay: "OneDay",
    OneWeek: "OneWeek",
    OneMonth: "OneMonth",
} as const;

export type DeliveryStreamAmazonopensearchserviceDestinationConfigurationIndexRotationPeriod = (typeof DeliveryStreamAmazonopensearchserviceDestinationConfigurationIndexRotationPeriod)[keyof typeof DeliveryStreamAmazonopensearchserviceDestinationConfigurationIndexRotationPeriod];

export const DeliveryStreamAmazonopensearchserviceDestinationConfigurationS3BackupMode = {
    FailedDocumentsOnly: "FailedDocumentsOnly",
    AllDocuments: "AllDocuments",
} as const;

export type DeliveryStreamAmazonopensearchserviceDestinationConfigurationS3BackupMode = (typeof DeliveryStreamAmazonopensearchserviceDestinationConfigurationS3BackupMode)[keyof typeof DeliveryStreamAmazonopensearchserviceDestinationConfigurationS3BackupMode];

export const DeliveryStreamAuthenticationConfigurationConnectivity = {
    Public: "PUBLIC",
    Private: "PRIVATE",
} as const;

export type DeliveryStreamAuthenticationConfigurationConnectivity = (typeof DeliveryStreamAuthenticationConfigurationConnectivity)[keyof typeof DeliveryStreamAuthenticationConfigurationConnectivity];

export const DeliveryStreamDocumentIdOptionsDefaultDocumentIdFormat = {
    FirehoseDefault: "FIREHOSE_DEFAULT",
    NoDocumentId: "NO_DOCUMENT_ID",
} as const;

export type DeliveryStreamDocumentIdOptionsDefaultDocumentIdFormat = (typeof DeliveryStreamDocumentIdOptionsDefaultDocumentIdFormat)[keyof typeof DeliveryStreamDocumentIdOptionsDefaultDocumentIdFormat];

export const DeliveryStreamElasticsearchDestinationConfigurationIndexRotationPeriod = {
    NoRotation: "NoRotation",
    OneHour: "OneHour",
    OneDay: "OneDay",
    OneWeek: "OneWeek",
    OneMonth: "OneMonth",
} as const;

export type DeliveryStreamElasticsearchDestinationConfigurationIndexRotationPeriod = (typeof DeliveryStreamElasticsearchDestinationConfigurationIndexRotationPeriod)[keyof typeof DeliveryStreamElasticsearchDestinationConfigurationIndexRotationPeriod];

export const DeliveryStreamElasticsearchDestinationConfigurationS3BackupMode = {
    FailedDocumentsOnly: "FailedDocumentsOnly",
    AllDocuments: "AllDocuments",
} as const;

export type DeliveryStreamElasticsearchDestinationConfigurationS3BackupMode = (typeof DeliveryStreamElasticsearchDestinationConfigurationS3BackupMode)[keyof typeof DeliveryStreamElasticsearchDestinationConfigurationS3BackupMode];

export const DeliveryStreamEncryptionConfigurationInputKeyType = {
    AwsOwnedCmk: "AWS_OWNED_CMK",
    CustomerManagedCmk: "CUSTOMER_MANAGED_CMK",
} as const;

export type DeliveryStreamEncryptionConfigurationInputKeyType = (typeof DeliveryStreamEncryptionConfigurationInputKeyType)[keyof typeof DeliveryStreamEncryptionConfigurationInputKeyType];

export const DeliveryStreamEncryptionConfigurationNoEncryptionConfig = {
    NoEncryption: "NoEncryption",
} as const;

export type DeliveryStreamEncryptionConfigurationNoEncryptionConfig = (typeof DeliveryStreamEncryptionConfigurationNoEncryptionConfig)[keyof typeof DeliveryStreamEncryptionConfigurationNoEncryptionConfig];

export const DeliveryStreamExtendedS3DestinationConfigurationCompressionFormat = {
    Uncompressed: "UNCOMPRESSED",
    Gzip: "GZIP",
    Zip: "ZIP",
    Snappy: "Snappy",
    HadoopSnappy: "HADOOP_SNAPPY",
} as const;

export type DeliveryStreamExtendedS3DestinationConfigurationCompressionFormat = (typeof DeliveryStreamExtendedS3DestinationConfigurationCompressionFormat)[keyof typeof DeliveryStreamExtendedS3DestinationConfigurationCompressionFormat];

export const DeliveryStreamExtendedS3DestinationConfigurationS3BackupMode = {
    Disabled: "Disabled",
    Enabled: "Enabled",
} as const;

export type DeliveryStreamExtendedS3DestinationConfigurationS3BackupMode = (typeof DeliveryStreamExtendedS3DestinationConfigurationS3BackupMode)[keyof typeof DeliveryStreamExtendedS3DestinationConfigurationS3BackupMode];

export const DeliveryStreamHttpEndpointRequestConfigurationContentEncoding = {
    None: "NONE",
    Gzip: "GZIP",
} as const;

export type DeliveryStreamHttpEndpointRequestConfigurationContentEncoding = (typeof DeliveryStreamHttpEndpointRequestConfigurationContentEncoding)[keyof typeof DeliveryStreamHttpEndpointRequestConfigurationContentEncoding];

export const DeliveryStreamProcessorType = {
    RecordDeAggregation: "RecordDeAggregation",
    Lambda: "Lambda",
    MetadataExtraction: "MetadataExtraction",
    AppendDelimiterToRecord: "AppendDelimiterToRecord",
} as const;

export type DeliveryStreamProcessorType = (typeof DeliveryStreamProcessorType)[keyof typeof DeliveryStreamProcessorType];

export const DeliveryStreamRedshiftDestinationConfigurationS3BackupMode = {
    Disabled: "Disabled",
    Enabled: "Enabled",
} as const;

export type DeliveryStreamRedshiftDestinationConfigurationS3BackupMode = (typeof DeliveryStreamRedshiftDestinationConfigurationS3BackupMode)[keyof typeof DeliveryStreamRedshiftDestinationConfigurationS3BackupMode];

export const DeliveryStreamS3DestinationConfigurationCompressionFormat = {
    Uncompressed: "UNCOMPRESSED",
    Gzip: "GZIP",
    Zip: "ZIP",
    Snappy: "Snappy",
    HadoopSnappy: "HADOOP_SNAPPY",
} as const;

export type DeliveryStreamS3DestinationConfigurationCompressionFormat = (typeof DeliveryStreamS3DestinationConfigurationCompressionFormat)[keyof typeof DeliveryStreamS3DestinationConfigurationCompressionFormat];

export const DeliveryStreamSplunkDestinationConfigurationHecEndpointType = {
    Raw: "Raw",
    Event: "Event",
} as const;

export type DeliveryStreamSplunkDestinationConfigurationHecEndpointType = (typeof DeliveryStreamSplunkDestinationConfigurationHecEndpointType)[keyof typeof DeliveryStreamSplunkDestinationConfigurationHecEndpointType];

export const DeliveryStreamType = {
    DirectPut: "DirectPut",
    KinesisStreamAsSource: "KinesisStreamAsSource",
    MskasSource: "MSKAsSource",
} as const;

export type DeliveryStreamType = (typeof DeliveryStreamType)[keyof typeof DeliveryStreamType];
