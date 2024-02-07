// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package route53

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Represents a key signing key (KSK) associated with a hosted zone. You can only have two KSKs per hosted zone.
type KeySigningKey struct {
	pulumi.CustomResourceState

	// The unique string (ID) used to identify a hosted zone.
	HostedZoneId pulumi.StringOutput `pulumi:"hostedZoneId"`
	// The Amazon resource name (ARN) for a customer managed key (CMK) in AWS Key Management Service (KMS). The KeyManagementServiceArn must be unique for each key signing key (KSK) in a single hosted zone.
	KeyManagementServiceArn pulumi.StringOutput `pulumi:"keyManagementServiceArn"`
	// An alphanumeric string used to identify a key signing key (KSK). Name must be unique for each key signing key in the same hosted zone.
	Name pulumi.StringOutput `pulumi:"name"`
	// A string specifying the initial status of the key signing key (KSK). You can set the value to ACTIVE or INACTIVE.
	Status KeySigningKeyStatusOutput `pulumi:"status"`
}

// NewKeySigningKey registers a new resource with the given unique name, arguments, and options.
func NewKeySigningKey(ctx *pulumi.Context,
	name string, args *KeySigningKeyArgs, opts ...pulumi.ResourceOption) (*KeySigningKey, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.HostedZoneId == nil {
		return nil, errors.New("invalid value for required argument 'HostedZoneId'")
	}
	if args.KeyManagementServiceArn == nil {
		return nil, errors.New("invalid value for required argument 'KeyManagementServiceArn'")
	}
	if args.Status == nil {
		return nil, errors.New("invalid value for required argument 'Status'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"hostedZoneId",
		"keyManagementServiceArn",
		"name",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource KeySigningKey
	err := ctx.RegisterResource("aws-native:route53:KeySigningKey", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetKeySigningKey gets an existing KeySigningKey resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetKeySigningKey(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *KeySigningKeyState, opts ...pulumi.ResourceOption) (*KeySigningKey, error) {
	var resource KeySigningKey
	err := ctx.ReadResource("aws-native:route53:KeySigningKey", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering KeySigningKey resources.
type keySigningKeyState struct {
}

type KeySigningKeyState struct {
}

func (KeySigningKeyState) ElementType() reflect.Type {
	return reflect.TypeOf((*keySigningKeyState)(nil)).Elem()
}

type keySigningKeyArgs struct {
	// The unique string (ID) used to identify a hosted zone.
	HostedZoneId string `pulumi:"hostedZoneId"`
	// The Amazon resource name (ARN) for a customer managed key (CMK) in AWS Key Management Service (KMS). The KeyManagementServiceArn must be unique for each key signing key (KSK) in a single hosted zone.
	KeyManagementServiceArn string `pulumi:"keyManagementServiceArn"`
	// An alphanumeric string used to identify a key signing key (KSK). Name must be unique for each key signing key in the same hosted zone.
	Name *string `pulumi:"name"`
	// A string specifying the initial status of the key signing key (KSK). You can set the value to ACTIVE or INACTIVE.
	Status KeySigningKeyStatus `pulumi:"status"`
}

// The set of arguments for constructing a KeySigningKey resource.
type KeySigningKeyArgs struct {
	// The unique string (ID) used to identify a hosted zone.
	HostedZoneId pulumi.StringInput
	// The Amazon resource name (ARN) for a customer managed key (CMK) in AWS Key Management Service (KMS). The KeyManagementServiceArn must be unique for each key signing key (KSK) in a single hosted zone.
	KeyManagementServiceArn pulumi.StringInput
	// An alphanumeric string used to identify a key signing key (KSK). Name must be unique for each key signing key in the same hosted zone.
	Name pulumi.StringPtrInput
	// A string specifying the initial status of the key signing key (KSK). You can set the value to ACTIVE or INACTIVE.
	Status KeySigningKeyStatusInput
}

func (KeySigningKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*keySigningKeyArgs)(nil)).Elem()
}

type KeySigningKeyInput interface {
	pulumi.Input

	ToKeySigningKeyOutput() KeySigningKeyOutput
	ToKeySigningKeyOutputWithContext(ctx context.Context) KeySigningKeyOutput
}

func (*KeySigningKey) ElementType() reflect.Type {
	return reflect.TypeOf((**KeySigningKey)(nil)).Elem()
}

func (i *KeySigningKey) ToKeySigningKeyOutput() KeySigningKeyOutput {
	return i.ToKeySigningKeyOutputWithContext(context.Background())
}

func (i *KeySigningKey) ToKeySigningKeyOutputWithContext(ctx context.Context) KeySigningKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeySigningKeyOutput)
}

type KeySigningKeyOutput struct{ *pulumi.OutputState }

func (KeySigningKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KeySigningKey)(nil)).Elem()
}

func (o KeySigningKeyOutput) ToKeySigningKeyOutput() KeySigningKeyOutput {
	return o
}

func (o KeySigningKeyOutput) ToKeySigningKeyOutputWithContext(ctx context.Context) KeySigningKeyOutput {
	return o
}

// The unique string (ID) used to identify a hosted zone.
func (o KeySigningKeyOutput) HostedZoneId() pulumi.StringOutput {
	return o.ApplyT(func(v *KeySigningKey) pulumi.StringOutput { return v.HostedZoneId }).(pulumi.StringOutput)
}

// The Amazon resource name (ARN) for a customer managed key (CMK) in AWS Key Management Service (KMS). The KeyManagementServiceArn must be unique for each key signing key (KSK) in a single hosted zone.
func (o KeySigningKeyOutput) KeyManagementServiceArn() pulumi.StringOutput {
	return o.ApplyT(func(v *KeySigningKey) pulumi.StringOutput { return v.KeyManagementServiceArn }).(pulumi.StringOutput)
}

// An alphanumeric string used to identify a key signing key (KSK). Name must be unique for each key signing key in the same hosted zone.
func (o KeySigningKeyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *KeySigningKey) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// A string specifying the initial status of the key signing key (KSK). You can set the value to ACTIVE or INACTIVE.
func (o KeySigningKeyOutput) Status() KeySigningKeyStatusOutput {
	return o.ApplyT(func(v *KeySigningKey) KeySigningKeyStatusOutput { return v.Status }).(KeySigningKeyStatusOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*KeySigningKeyInput)(nil)).Elem(), &KeySigningKey{})
	pulumi.RegisterOutputType(KeySigningKeyOutput{})
}
