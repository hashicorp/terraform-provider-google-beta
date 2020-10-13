behavior "assign_random_reviewer" "random" {
  reviewers            = [
    "danawillow",
    "megan07",
    "rileykarson",
    "ndmckinley",
    "slevenick",
    "c2thorn",
    "scottsuarez",
  ]
  only_non_maintainers = true
}

behavior "regexp_issue_labeler" "panic_label" {
    regexp = "panic:"
    labels = ["crash", "bug"]
}

behavior "remove_labels_on_reply" "remove_stale" {
    labels = ["waiting-response", "stale"]
    only_non_maintainers = true
}

behavior "pull_request_size_labeler" "size" {
    label_prefix = "size/"
    label_map = {
        "size/XS" = {
            from = 0
            to = 30
        }
        "size/S" = {
            from = 31
            to = 60
        }
        "size/M" = {
            from = 61
            to = 150
        }
        "size/L" = {
            from = 151
            to = 300
        }
        "size/XL" = {
            from = 301
            to = 1000
        }
        "size/XXL" = {
            from = 1001
            to = 0
        }
    }
}
