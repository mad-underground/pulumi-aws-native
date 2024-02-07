// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package lex

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Amazon Lex conversational bot performing automated tasks such as ordering a pizza, booking a hotel, and so on.
type Bot struct {
	pulumi.CustomResourceState

	Arn pulumi.StringOutput `pulumi:"arn"`
	// Specifies whether to build the bot locales after bot creation completes.
	AutoBuildBotLocales pulumi.BoolPtrOutput   `pulumi:"autoBuildBotLocales"`
	BotFileS3Location   BotS3LocationPtrOutput `pulumi:"botFileS3Location"`
	// List of bot locales
	BotLocales BotLocaleArrayOutput `pulumi:"botLocales"`
	// A list of tags to add to the bot, which can only be added at bot creation.
	BotTags BotTagArrayOutput `pulumi:"botTags"`
	// Data privacy setting of the Bot.
	DataPrivacy DataPrivacyPropertiesOutput `pulumi:"dataPrivacy"`
	Description pulumi.StringPtrOutput      `pulumi:"description"`
	// IdleSessionTTLInSeconds of the resource
	IdleSessionTtlInSeconds pulumi.IntOutput                 `pulumi:"idleSessionTtlInSeconds"`
	Name                    pulumi.StringOutput              `pulumi:"name"`
	RoleArn                 pulumi.StringOutput              `pulumi:"roleArn"`
	TestBotAliasSettings    BotTestBotAliasSettingsPtrOutput `pulumi:"testBotAliasSettings"`
	// A list of tags to add to the test alias for a bot, , which can only be added at bot/bot alias creation.
	TestBotAliasTags BotTagArrayOutput `pulumi:"testBotAliasTags"`
}

// NewBot registers a new resource with the given unique name, arguments, and options.
func NewBot(ctx *pulumi.Context,
	name string, args *BotArgs, opts ...pulumi.ResourceOption) (*Bot, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DataPrivacy == nil {
		return nil, errors.New("invalid value for required argument 'DataPrivacy'")
	}
	if args.IdleSessionTtlInSeconds == nil {
		return nil, errors.New("invalid value for required argument 'IdleSessionTtlInSeconds'")
	}
	if args.RoleArn == nil {
		return nil, errors.New("invalid value for required argument 'RoleArn'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Bot
	err := ctx.RegisterResource("aws-native:lex:Bot", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBot gets an existing Bot resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBot(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BotState, opts ...pulumi.ResourceOption) (*Bot, error) {
	var resource Bot
	err := ctx.ReadResource("aws-native:lex:Bot", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Bot resources.
type botState struct {
}

type BotState struct {
}

func (BotState) ElementType() reflect.Type {
	return reflect.TypeOf((*botState)(nil)).Elem()
}

type botArgs struct {
	// Specifies whether to build the bot locales after bot creation completes.
	AutoBuildBotLocales *bool          `pulumi:"autoBuildBotLocales"`
	BotFileS3Location   *BotS3Location `pulumi:"botFileS3Location"`
	// List of bot locales
	BotLocales []BotLocale `pulumi:"botLocales"`
	// A list of tags to add to the bot, which can only be added at bot creation.
	BotTags []BotTag `pulumi:"botTags"`
	// Data privacy setting of the Bot.
	DataPrivacy DataPrivacyProperties `pulumi:"dataPrivacy"`
	Description *string               `pulumi:"description"`
	// IdleSessionTTLInSeconds of the resource
	IdleSessionTtlInSeconds int                      `pulumi:"idleSessionTtlInSeconds"`
	Name                    *string                  `pulumi:"name"`
	RoleArn                 string                   `pulumi:"roleArn"`
	TestBotAliasSettings    *BotTestBotAliasSettings `pulumi:"testBotAliasSettings"`
	// A list of tags to add to the test alias for a bot, , which can only be added at bot/bot alias creation.
	TestBotAliasTags []BotTag `pulumi:"testBotAliasTags"`
}

// The set of arguments for constructing a Bot resource.
type BotArgs struct {
	// Specifies whether to build the bot locales after bot creation completes.
	AutoBuildBotLocales pulumi.BoolPtrInput
	BotFileS3Location   BotS3LocationPtrInput
	// List of bot locales
	BotLocales BotLocaleArrayInput
	// A list of tags to add to the bot, which can only be added at bot creation.
	BotTags BotTagArrayInput
	// Data privacy setting of the Bot.
	DataPrivacy DataPrivacyPropertiesInput
	Description pulumi.StringPtrInput
	// IdleSessionTTLInSeconds of the resource
	IdleSessionTtlInSeconds pulumi.IntInput
	Name                    pulumi.StringPtrInput
	RoleArn                 pulumi.StringInput
	TestBotAliasSettings    BotTestBotAliasSettingsPtrInput
	// A list of tags to add to the test alias for a bot, , which can only be added at bot/bot alias creation.
	TestBotAliasTags BotTagArrayInput
}

func (BotArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*botArgs)(nil)).Elem()
}

type BotInput interface {
	pulumi.Input

	ToBotOutput() BotOutput
	ToBotOutputWithContext(ctx context.Context) BotOutput
}

func (*Bot) ElementType() reflect.Type {
	return reflect.TypeOf((**Bot)(nil)).Elem()
}

func (i *Bot) ToBotOutput() BotOutput {
	return i.ToBotOutputWithContext(context.Background())
}

func (i *Bot) ToBotOutputWithContext(ctx context.Context) BotOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BotOutput)
}

type BotOutput struct{ *pulumi.OutputState }

func (BotOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Bot)(nil)).Elem()
}

func (o BotOutput) ToBotOutput() BotOutput {
	return o
}

func (o BotOutput) ToBotOutputWithContext(ctx context.Context) BotOutput {
	return o
}

func (o BotOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *Bot) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// Specifies whether to build the bot locales after bot creation completes.
func (o BotOutput) AutoBuildBotLocales() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Bot) pulumi.BoolPtrOutput { return v.AutoBuildBotLocales }).(pulumi.BoolPtrOutput)
}

func (o BotOutput) BotFileS3Location() BotS3LocationPtrOutput {
	return o.ApplyT(func(v *Bot) BotS3LocationPtrOutput { return v.BotFileS3Location }).(BotS3LocationPtrOutput)
}

// List of bot locales
func (o BotOutput) BotLocales() BotLocaleArrayOutput {
	return o.ApplyT(func(v *Bot) BotLocaleArrayOutput { return v.BotLocales }).(BotLocaleArrayOutput)
}

// A list of tags to add to the bot, which can only be added at bot creation.
func (o BotOutput) BotTags() BotTagArrayOutput {
	return o.ApplyT(func(v *Bot) BotTagArrayOutput { return v.BotTags }).(BotTagArrayOutput)
}

// Data privacy setting of the Bot.
func (o BotOutput) DataPrivacy() DataPrivacyPropertiesOutput {
	return o.ApplyT(func(v *Bot) DataPrivacyPropertiesOutput { return v.DataPrivacy }).(DataPrivacyPropertiesOutput)
}

func (o BotOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Bot) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// IdleSessionTTLInSeconds of the resource
func (o BotOutput) IdleSessionTtlInSeconds() pulumi.IntOutput {
	return o.ApplyT(func(v *Bot) pulumi.IntOutput { return v.IdleSessionTtlInSeconds }).(pulumi.IntOutput)
}

func (o BotOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Bot) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o BotOutput) RoleArn() pulumi.StringOutput {
	return o.ApplyT(func(v *Bot) pulumi.StringOutput { return v.RoleArn }).(pulumi.StringOutput)
}

func (o BotOutput) TestBotAliasSettings() BotTestBotAliasSettingsPtrOutput {
	return o.ApplyT(func(v *Bot) BotTestBotAliasSettingsPtrOutput { return v.TestBotAliasSettings }).(BotTestBotAliasSettingsPtrOutput)
}

// A list of tags to add to the test alias for a bot, , which can only be added at bot/bot alias creation.
func (o BotOutput) TestBotAliasTags() BotTagArrayOutput {
	return o.ApplyT(func(v *Bot) BotTagArrayOutput { return v.TestBotAliasTags }).(BotTagArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BotInput)(nil)).Elem(), &Bot{})
	pulumi.RegisterOutputType(BotOutput{})
}
