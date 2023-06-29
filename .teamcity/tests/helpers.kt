package tests

import ClientConfiguration

fun TestConfiguration() : ClientConfiguration {
    return ClientConfiguration("custId", "org", "org2", "billingAccount", "billingAccount2", "masterBillingAccount", "project", "orgDomain", "projectNumber", "region", "serviceAccount", "zone", "credentials")
}