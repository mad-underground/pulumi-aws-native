// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoTWireless.Outputs
{

    [OutputType]
    public sealed class DeviceProfileLoRaWanDeviceProfile
    {
        public readonly int? ClassBTimeout;
        public readonly int? ClassCTimeout;
        public readonly ImmutableArray<int> FactoryPresetFreqsList;
        public readonly string? MacVersion;
        public readonly int? MaxDutyCycle;
        public readonly int? MaxEirp;
        public readonly int? PingSlotDr;
        public readonly int? PingSlotFreq;
        public readonly int? PingSlotPeriod;
        public readonly string? RegParamsRevision;
        public readonly string? RfRegion;
        public readonly int? RxDataRate2;
        public readonly int? RxDelay1;
        public readonly int? RxDrOffset1;
        public readonly int? RxFreq2;
        public readonly bool? Supports32BitFCnt;
        public readonly bool? SupportsClassB;
        public readonly bool? SupportsClassC;
        public readonly bool? SupportsJoin;

        [OutputConstructor]
        private DeviceProfileLoRaWanDeviceProfile(
            int? classBTimeout,

            int? classCTimeout,

            ImmutableArray<int> factoryPresetFreqsList,

            string? macVersion,

            int? maxDutyCycle,

            int? maxEirp,

            int? pingSlotDr,

            int? pingSlotFreq,

            int? pingSlotPeriod,

            string? regParamsRevision,

            string? rfRegion,

            int? rxDataRate2,

            int? rxDelay1,

            int? rxDrOffset1,

            int? rxFreq2,

            bool? supports32BitFCnt,

            bool? supportsClassB,

            bool? supportsClassC,

            bool? supportsJoin)
        {
            ClassBTimeout = classBTimeout;
            ClassCTimeout = classCTimeout;
            FactoryPresetFreqsList = factoryPresetFreqsList;
            MacVersion = macVersion;
            MaxDutyCycle = maxDutyCycle;
            MaxEirp = maxEirp;
            PingSlotDr = pingSlotDr;
            PingSlotFreq = pingSlotFreq;
            PingSlotPeriod = pingSlotPeriod;
            RegParamsRevision = regParamsRevision;
            RfRegion = rfRegion;
            RxDataRate2 = rxDataRate2;
            RxDelay1 = rxDelay1;
            RxDrOffset1 = rxDrOffset1;
            RxFreq2 = rxFreq2;
            Supports32BitFCnt = supports32BitFCnt;
            SupportsClassB = supportsClassB;
            SupportsClassC = supportsClassC;
            SupportsJoin = supportsJoin;
        }
    }
}
