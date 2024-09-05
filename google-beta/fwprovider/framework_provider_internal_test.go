// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package fwprovider_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/fwprovider"
)

func TestFrameworkProvider_impl(t *testing.T) {
	var _ provider.ProviderWithMetaSchema = fwprovider.New()
}
