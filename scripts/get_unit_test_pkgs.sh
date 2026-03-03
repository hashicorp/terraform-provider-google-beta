#!/usr/bin/env bash
# scripts/get_unit_test_pkgs.sh

# Ask Go for BOTH the Import Path and the exact Directory Path on disk, separated by a pipe '|'
go list -e -f '{{.ImportPath}}|{{.Dir}}' ./... | grep -v "/scripts" | while IFS='|' read -r pkg dir; do

    # For NON-SERVICE packages:
    if [[ "$pkg" != *"/google-beta/services/"* ]]; then
        # Keep ONLY if it has at least one test of ANY kind.
        # `grep -q` will return true as soon as it finds one "func Test"
        if cat "$dir"/*_test.go 2>/dev/null | grep -q "^func Test"; then
            echo "$pkg"
        fi
        continue
    fi

    # 2. For SERVICE packages: 
    # Keep ONLY if it has at least one Unit test (a test that does NOT start with TestAcc).
    if cat "$dir"/*_test.go 2>/dev/null | grep "^func Test" | grep -v -q "^func TestAcc"; then
        echo "$pkg"
    fi

done