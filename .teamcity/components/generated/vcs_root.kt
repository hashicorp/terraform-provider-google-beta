// this file is auto-generated with mmv1, any changes made here will be overwritten

import jetbrains.buildServer.configs.kotlin.v2019_2.vcs.GitVcsRoot

object providerRepository : GitVcsRoot({
    name = "terraform-provider-google-beta"
    url = "https://github.com/hashicorp/terraform-provider-google-beta.git"
    agentCleanPolicy = AgentCleanPolicy.ON_BRANCH_CHANGE
    agentCleanFilesPolicy = AgentCleanFilesPolicy.ALL_UNTRACKED
    branchSpec = "+:*"
    branch = "refs/heads/megan_tc_config"
    authMethod = anonymous()
})
