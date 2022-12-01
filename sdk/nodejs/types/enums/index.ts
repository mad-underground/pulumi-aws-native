// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// Export sub-modules:
import * as amplify from "./amplify";
import * as amplifyuibuilder from "./amplifyuibuilder";
import * as apigateway from "./apigateway";
import * as appflow from "./appflow";
import * as applicationinsights from "./applicationinsights";
import * as apprunner from "./apprunner";
import * as athena from "./athena";
import * as auditmanager from "./auditmanager";
import * as batch from "./batch";
import * as billingconductor from "./billingconductor";
import * as budgets from "./budgets";
import * as cassandra from "./cassandra";
import * as ce from "./ce";
import * as cloudformation from "./cloudformation";
import * as cloudfront from "./cloudfront";
import * as cloudtrail from "./cloudtrail";
import * as codeguruprofiler from "./codeguruprofiler";
import * as codegurureviewer from "./codegurureviewer";
import * as codestarnotifications from "./codestarnotifications";
import * as connect from "./connect";
import * as cur from "./cur";
import * as customerprofiles from "./customerprofiles";
import * as databrew from "./databrew";
import * as datasync from "./datasync";
import * as devicefarm from "./devicefarm";
import * as devopsguru from "./devopsguru";
import * as ec2 from "./ec2";
import * as ecr from "./ecr";
import * as ecs from "./ecs";
import * as eks from "./eks";
import * as elasticache from "./elasticache";
import * as emr from "./emr";
import * as events from "./events";
import * as evidently from "./evidently";
import * as finspace from "./finspace";
import * as fms from "./fms";
import * as forecast from "./forecast";
import * as frauddetector from "./frauddetector";
import * as fsx from "./fsx";
import * as gamelift from "./gamelift";
import * as globalaccelerator from "./globalaccelerator";
import * as glue from "./glue";
import * as greengrassv2 from "./greengrassv2";
import * as groundstation from "./groundstation";
import * as healthlake from "./healthlake";
import * as imagebuilder from "./imagebuilder";
import * as inspectorv2 from "./inspectorv2";
import * as iot from "./iot";
import * as iotanalytics from "./iotanalytics";
import * as iotevents from "./iotevents";
import * as iotfleetwise from "./iotfleetwise";
import * as iotsitewise from "./iotsitewise";
import * as iottwinmaker from "./iottwinmaker";
import * as iotwireless from "./iotwireless";
import * as ivs from "./ivs";
import * as kafkaconnect from "./kafkaconnect";
import * as kendra from "./kendra";
import * as kinesis from "./kinesis";
import * as kinesisanalyticsv2 from "./kinesisanalyticsv2";
import * as kinesisfirehose from "./kinesisfirehose";
import * as kinesisvideo from "./kinesisvideo";
import * as kms from "./kms";
import * as lakeformation from "./lakeformation";
import * as lambda from "./lambda";
import * as lex from "./lex";
import * as lightsail from "./lightsail";
import * as location from "./location";
import * as logs from "./logs";
import * as lookoutequipment from "./lookoutequipment";
import * as lookoutmetrics from "./lookoutmetrics";
import * as m2 from "./m2";
import * as macie from "./macie";
import * as mediaconnect from "./mediaconnect";
import * as mediapackage from "./mediapackage";
import * as mediatailor from "./mediatailor";
import * as memorydb from "./memorydb";
import * as msk from "./msk";
import * as mwaa from "./mwaa";
import * as networkfirewall from "./networkfirewall";
import * as nimblestudio from "./nimblestudio";
import * as oam from "./oam";
import * as organizations from "./organizations";
import * as panorama from "./panorama";
import * as personalize from "./personalize";
import * as pinpoint from "./pinpoint";
import * as quicksight from "./quicksight";
import * as rds from "./rds";
import * as redshift from "./redshift";
import * as redshiftserverless from "./redshiftserverless";
import * as refactorspaces from "./refactorspaces";
import * as resiliencehub from "./resiliencehub";
import * as resourceexplorer2 from "./resourceexplorer2";
import * as resourcegroups from "./resourcegroups";
import * as robomaker from "./robomaker";
import * as rolesanywhere from "./rolesanywhere";
import * as route53 from "./route53";
import * as route53recoverycontrol from "./route53recoverycontrol";
import * as route53resolver from "./route53resolver";
import * as rum from "./rum";
import * as s3 from "./s3";
import * as s3outposts from "./s3outposts";
import * as sagemaker from "./sagemaker";
import * as scheduler from "./scheduler";
import * as servicecatalog from "./servicecatalog";
import * as servicecatalogappregistry from "./servicecatalogappregistry";
import * as signer from "./signer";
import * as ssm from "./ssm";
import * as ssmcontacts from "./ssmcontacts";
import * as ssmincidents from "./ssmincidents";
import * as sso from "./sso";
import * as stepfunctions from "./stepfunctions";
import * as supportapp from "./supportapp";
import * as timestream from "./timestream";
import * as transfer from "./transfer";
import * as wafv2 from "./wafv2";
import * as wisdom from "./wisdom";
import * as workspaces from "./workspaces";

export {
    amplify,
    amplifyuibuilder,
    apigateway,
    appflow,
    applicationinsights,
    apprunner,
    athena,
    auditmanager,
    batch,
    billingconductor,
    budgets,
    cassandra,
    ce,
    cloudformation,
    cloudfront,
    cloudtrail,
    codeguruprofiler,
    codegurureviewer,
    codestarnotifications,
    connect,
    cur,
    customerprofiles,
    databrew,
    datasync,
    devicefarm,
    devopsguru,
    ec2,
    ecr,
    ecs,
    eks,
    elasticache,
    emr,
    events,
    evidently,
    finspace,
    fms,
    forecast,
    frauddetector,
    fsx,
    gamelift,
    globalaccelerator,
    glue,
    greengrassv2,
    groundstation,
    healthlake,
    imagebuilder,
    inspectorv2,
    iot,
    iotanalytics,
    iotevents,
    iotfleetwise,
    iotsitewise,
    iottwinmaker,
    iotwireless,
    ivs,
    kafkaconnect,
    kendra,
    kinesis,
    kinesisanalyticsv2,
    kinesisfirehose,
    kinesisvideo,
    kms,
    lakeformation,
    lambda,
    lex,
    lightsail,
    location,
    logs,
    lookoutequipment,
    lookoutmetrics,
    m2,
    macie,
    mediaconnect,
    mediapackage,
    mediatailor,
    memorydb,
    msk,
    mwaa,
    networkfirewall,
    nimblestudio,
    oam,
    organizations,
    panorama,
    personalize,
    pinpoint,
    quicksight,
    rds,
    redshift,
    redshiftserverless,
    refactorspaces,
    resiliencehub,
    resourceexplorer2,
    resourcegroups,
    robomaker,
    rolesanywhere,
    route53,
    route53recoverycontrol,
    route53resolver,
    rum,
    s3,
    s3outposts,
    sagemaker,
    scheduler,
    servicecatalog,
    servicecatalogappregistry,
    signer,
    ssm,
    ssmcontacts,
    ssmincidents,
    sso,
    stepfunctions,
    supportapp,
    timestream,
    transfer,
    wafv2,
    wisdom,
    workspaces,
};

export const Region = {
    /**
     * Africa (Cape Town)
     */
    AFSouth1: "af-south-1",
    /**
     * Asia Pacific (Hong Kong)
     */
    APEast1: "ap-east-1",
    /**
     * Asia Pacific (Tokyo)
     */
    APNortheast1: "ap-northeast-1",
    /**
     * Asia Pacific (Seoul)
     */
    APNortheast2: "ap-northeast-2",
    /**
     * Asia Pacific (Osaka)
     */
    APNortheast3: "ap-northeast-3",
    /**
     * Asia Pacific (Mumbai)
     */
    APSouth1: "ap-south-1",
    /**
     * Asia Pacific (Singapore)
     */
    APSoutheast1: "ap-southeast-1",
    /**
     * Asia Pacific (Sydney)
     */
    APSoutheast2: "ap-southeast-2",
    /**
     * Canada (Central)
     */
    CACentral: "ca-central-1",
    /**
     * China (Beijing)
     */
    CNNorth1: "cn-north-1",
    /**
     * China (Ningxia)
     */
    CNNorthwest1: "cn-northwest-1",
    /**
     * Europe (Frankfurt)
     */
    EUCentral1: "eu-central-1",
    /**
     * Europe (Stockholm)
     */
    EUNorth1: "eu-north-1",
    /**
     * Europe (Ireland)
     */
    EUWest1: "eu-west-1",
    /**
     * Europe (London)
     */
    EUWest2: "eu-west-2",
    /**
     * Europe (Paris)
     */
    EUWest3: "eu-west-3",
    /**
     * Europe (Milan)
     */
    EUSouth1: "eu-south-1",
    /**
     * Middle East (Bahrain)
     */
    MESouth1: "me-south-1",
    /**
     * South America (São Paulo)
     */
    SAEast1: "sa-east-1",
    /**
     * AWS GovCloud (US-East)
     */
    USGovEast1: "us-gov-east-1",
    /**
     * AWS GovCloud (US-West)
     */
    USGovWest1: "us-gov-west-1",
    /**
     * US East (N. Virginia)
     */
    USEast1: "us-east-1",
    /**
     * US East (Ohio)
     */
    USEast2: "us-east-2",
    /**
     * US West (N. California)
     */
    USWest1: "us-west-1",
    /**
     * US West (Oregon)
     */
    USWest2: "us-west-2",
} as const;

/**
 * A Region represents any valid Amazon region that may be targeted with deployments.
 */
export type Region = (typeof Region)[keyof typeof Region];
