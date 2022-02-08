// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::SageMaker::Domain
 */
export function getDomain(args: GetDomainArgs, opts?: pulumi.InvokeOptions): Promise<GetDomainResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:sagemaker:getDomain", {
        "domainId": args.domainId,
    }, opts);
}

export interface GetDomainArgs {
    /**
     * The domain name.
     */
    domainId: string;
}

export interface GetDomainResult {
    /**
     * The default user settings.
     */
    readonly defaultUserSettings?: outputs.sagemaker.DomainUserSettings;
    /**
     * The Amazon Resource Name (ARN) of the created domain.
     */
    readonly domainArn?: string;
    /**
     * The domain name.
     */
    readonly domainId?: string;
    /**
     * The ID of the Amazon Elastic File System (EFS) managed by this Domain.
     */
    readonly homeEfsFileSystemId?: string;
    /**
     * The SSO managed application instance ID.
     */
    readonly singleSignOnManagedApplicationInstanceId?: string;
    /**
     * The URL to the created domain.
     */
    readonly url?: string;
}

export function getDomainOutput(args: GetDomainOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDomainResult> {
    return pulumi.output(args).apply(a => getDomain(a, opts))
}

export interface GetDomainOutputArgs {
    /**
     * The domain name.
     */
    domainId: pulumi.Input<string>;
}
