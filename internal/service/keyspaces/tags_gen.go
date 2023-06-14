// Code generated by internal/generate/tags/main.go; DO NOT EDIT.
package keyspaces

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/keyspaces"
	"github.com/aws/aws-sdk-go/service/keyspaces/keyspacesiface"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	tftags "github.com/hashicorp/terraform-provider-aws/internal/tags"
	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

// ListTags lists keyspaces service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func ListTags(ctx context.Context, conn keyspacesiface.KeyspacesAPI, identifier string) (tftags.KeyValueTags, error) {
	input := &keyspaces.ListTagsForResourceInput{
		ResourceArn: aws.String(identifier),
	}

	output, err := conn.ListTagsForResourceWithContext(ctx, input)

	if err != nil {
		return tftags.New(ctx, nil), err
	}

	return KeyValueTags(ctx, output.Tags), nil
}

// ListTags lists keyspaces service tags and set them in Context.
// It is called from outside this package.
func (p *servicePackage) ListTags(ctx context.Context, meta any, identifier string) error {
	tags, err := ListTags(ctx, meta.(*conns.AWSClient).KeyspacesConn(ctx), identifier)

	if err != nil {
		return err
	}

	if inContext, ok := tftags.FromContext(ctx); ok {
		inContext.TagsOut = types.Some(tags)
	}

	return nil
}

// []*SERVICE.Tag handling

// Tags returns keyspaces service tags.
func Tags(tags tftags.KeyValueTags) []*keyspaces.Tag {
	result := make([]*keyspaces.Tag, 0, len(tags))

	for k, v := range tags.Map() {
		tag := &keyspaces.Tag{
			Key:   aws.String(k),
			Value: aws.String(v),
		}

		result = append(result, tag)
	}

	return result
}

// KeyValueTags creates tftags.KeyValueTags from keyspaces service tags.
func KeyValueTags(ctx context.Context, tags []*keyspaces.Tag) tftags.KeyValueTags {
	m := make(map[string]*string, len(tags))

	for _, tag := range tags {
		m[aws.StringValue(tag.Key)] = tag.Value
	}

	return tftags.New(ctx, m)
}

// GetTagsIn returns keyspaces service tags from Context.
// nil is returned if there are no input tags.
func GetTagsIn(ctx context.Context) []*keyspaces.Tag {
	if inContext, ok := tftags.FromContext(ctx); ok {
		if tags := Tags(inContext.TagsIn.UnwrapOrDefault()); len(tags) > 0 {
			return tags
		}
	}

	return nil
}

// SetTagsOut sets keyspaces service tags in Context.
func SetTagsOut(ctx context.Context, tags []*keyspaces.Tag) {
	if inContext, ok := tftags.FromContext(ctx); ok {
		inContext.TagsOut = types.Some(KeyValueTags(ctx, tags))
	}
}

// UpdateTags updates keyspaces service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func UpdateTags(ctx context.Context, conn keyspacesiface.KeyspacesAPI, identifier string, oldTagsMap, newTagsMap any) error {
	oldTags := tftags.New(ctx, oldTagsMap)
	newTags := tftags.New(ctx, newTagsMap)

	removedTags := oldTags.Removed(newTags)
	removedTags = removedTags.IgnoreSystem(names.Keyspaces)
	if len(removedTags) > 0 {
		input := &keyspaces.UntagResourceInput{
			ResourceArn: aws.String(identifier),
			Tags:        Tags(removedTags),
		}

		_, err := conn.UntagResourceWithContext(ctx, input)

		if err != nil {
			return fmt.Errorf("untagging resource (%s): %w", identifier, err)
		}
	}

	updatedTags := oldTags.Updated(newTags)
	updatedTags = updatedTags.IgnoreSystem(names.Keyspaces)
	if len(updatedTags) > 0 {
		input := &keyspaces.TagResourceInput{
			ResourceArn: aws.String(identifier),
			Tags:        Tags(updatedTags),
		}

		_, err := conn.TagResourceWithContext(ctx, input)

		if err != nil {
			return fmt.Errorf("tagging resource (%s): %w", identifier, err)
		}
	}

	return nil
}

// UpdateTags updates keyspaces service tags.
// It is called from outside this package.
func (p *servicePackage) UpdateTags(ctx context.Context, meta any, identifier string, oldTags, newTags any) error {
	return UpdateTags(ctx, meta.(*conns.AWSClient).KeyspacesConn(ctx), identifier, oldTags, newTags)
}
