#!/usr/bin/env bash
# scripts/get_unit_test_pkgs.sh

ALL_PKGS=$(go list -e ./... | grep -v "github.com/hashicorp/terraform-provider-google-beta/scripts")

for pkg in $ALL_PKGS; do
    # If it is NOT in the services directory, keep it automatically
    if [[ "$pkg" != *"/google-beta/services/"* ]]; then
        echo "$pkg"
        continue
    fi

    # First, convert the Go import path into a local directory path.
    # Example: github.com/hashicorp/.../services/iam -> ./google-beta/services/iam
    dir="./${pkg#$MOD_PATH/}"

    # - cat all _test.go files in that directory (hiding errors if none exist)
    # - grep for any function starting with "func Test"
    # - exclude functions starting with "func TestAcc"
    # - If grep finds at least one match, we keep the package!
    if cat "$dir"/*_test.go 2>/dev/null | grep "^func Test" | grep -v -q "^func TestAcc"; then
        echo "$pkg"
    fi
done
