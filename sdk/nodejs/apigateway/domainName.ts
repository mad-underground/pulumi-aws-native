// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-domainname.html
 */
export class DomainName extends pulumi.CustomResource {
    /**
     * Get an existing DomainName resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): DomainName {
        return new DomainName(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:ApiGateway:DomainName';

    /**
     * Returns true if the given object is an instance of DomainName.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DomainName {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DomainName.__pulumiType;
    }

    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-domainname.html#cfn-apigateway-domainname-certificatearn
     */
    public readonly CertificateArn!: pulumi.Output<string | undefined>;
    public /*out*/ readonly DistributionDomainName!: pulumi.Output<string>;
    public /*out*/ readonly DistributionHostedZoneId!: pulumi.Output<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-domainname.html#cfn-apigateway-domainname-domainname
     */
    public readonly DomainName!: pulumi.Output<string | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-domainname.html#cfn-apigateway-domainname-endpointconfiguration
     */
    public readonly EndpointConfiguration!: pulumi.Output<outputs.ApiGateway.DomainNameEndpointConfiguration | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-domainname.html#cfn-apigateway-domainname-mutualtlsauthentication
     */
    public readonly MutualTlsAuthentication!: pulumi.Output<outputs.ApiGateway.DomainNameMutualTlsAuthentication | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-domainname.html#cfn-apigateway-domainname-regionalcertificatearn
     */
    public readonly RegionalCertificateArn!: pulumi.Output<string | undefined>;
    public /*out*/ readonly RegionalDomainName!: pulumi.Output<string>;
    public /*out*/ readonly RegionalHostedZoneId!: pulumi.Output<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-domainname.html#cfn-apigateway-domainname-securitypolicy
     */
    public readonly SecurityPolicy!: pulumi.Output<string | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-domainname.html#cfn-apigateway-domainname-tags
     */
    public readonly Tags!: pulumi.Output<outputs.Tag[] | undefined>;

    /**
     * Create a DomainName resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: DomainNameArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            inputs["CertificateArn"] = args ? args.CertificateArn : undefined;
            inputs["DomainName"] = args ? args.DomainName : undefined;
            inputs["EndpointConfiguration"] = args ? args.EndpointConfiguration : undefined;
            inputs["MutualTlsAuthentication"] = args ? args.MutualTlsAuthentication : undefined;
            inputs["RegionalCertificateArn"] = args ? args.RegionalCertificateArn : undefined;
            inputs["SecurityPolicy"] = args ? args.SecurityPolicy : undefined;
            inputs["Tags"] = args ? args.Tags : undefined;
            inputs["DistributionDomainName"] = undefined /*out*/;
            inputs["DistributionHostedZoneId"] = undefined /*out*/;
            inputs["RegionalDomainName"] = undefined /*out*/;
            inputs["RegionalHostedZoneId"] = undefined /*out*/;
        } else {
            inputs["CertificateArn"] = undefined /*out*/;
            inputs["DistributionDomainName"] = undefined /*out*/;
            inputs["DistributionHostedZoneId"] = undefined /*out*/;
            inputs["DomainName"] = undefined /*out*/;
            inputs["EndpointConfiguration"] = undefined /*out*/;
            inputs["MutualTlsAuthentication"] = undefined /*out*/;
            inputs["RegionalCertificateArn"] = undefined /*out*/;
            inputs["RegionalDomainName"] = undefined /*out*/;
            inputs["RegionalHostedZoneId"] = undefined /*out*/;
            inputs["SecurityPolicy"] = undefined /*out*/;
            inputs["Tags"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(DomainName.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a DomainName resource.
 */
export interface DomainNameArgs {
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-domainname.html#cfn-apigateway-domainname-certificatearn
     */
    readonly CertificateArn?: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-domainname.html#cfn-apigateway-domainname-domainname
     */
    readonly DomainName?: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-domainname.html#cfn-apigateway-domainname-endpointconfiguration
     */
    readonly EndpointConfiguration?: pulumi.Input<inputs.ApiGateway.DomainNameEndpointConfiguration>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-domainname.html#cfn-apigateway-domainname-mutualtlsauthentication
     */
    readonly MutualTlsAuthentication?: pulumi.Input<inputs.ApiGateway.DomainNameMutualTlsAuthentication>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-domainname.html#cfn-apigateway-domainname-regionalcertificatearn
     */
    readonly RegionalCertificateArn?: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-domainname.html#cfn-apigateway-domainname-securitypolicy
     */
    readonly SecurityPolicy?: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-domainname.html#cfn-apigateway-domainname-tags
     */
    readonly Tags?: pulumi.Input<pulumi.Input<inputs.Tag>[]>;
}
