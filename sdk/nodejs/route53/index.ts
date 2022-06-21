// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./cidrCollection";
export * from "./dnssec";
export * from "./getCidrCollection";
export * from "./getHealthCheck";
export * from "./getHostedZone";
export * from "./getKeySigningKey";
export * from "./getRecordSet";
export * from "./getRecordSetGroup";
export * from "./healthCheck";
export * from "./hostedZone";
export * from "./keySigningKey";
export * from "./recordSet";
export * from "./recordSetGroup";

// Export enums:
export * from "../types/enums/route53";

// Import resources to register:
import { CidrCollection } from "./cidrCollection";
import { DNSSEC } from "./dnssec";
import { HealthCheck } from "./healthCheck";
import { HostedZone } from "./hostedZone";
import { KeySigningKey } from "./keySigningKey";
import { RecordSet } from "./recordSet";
import { RecordSetGroup } from "./recordSetGroup";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "aws-native:route53:CidrCollection":
                return new CidrCollection(name, <any>undefined, { urn })
            case "aws-native:route53:DNSSEC":
                return new DNSSEC(name, <any>undefined, { urn })
            case "aws-native:route53:HealthCheck":
                return new HealthCheck(name, <any>undefined, { urn })
            case "aws-native:route53:HostedZone":
                return new HostedZone(name, <any>undefined, { urn })
            case "aws-native:route53:KeySigningKey":
                return new KeySigningKey(name, <any>undefined, { urn })
            case "aws-native:route53:RecordSet":
                return new RecordSet(name, <any>undefined, { urn })
            case "aws-native:route53:RecordSetGroup":
                return new RecordSetGroup(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("aws-native", "route53", _module)
