// this file is auto-generated with mmv1, any changes made here will be overwritten

import jetbrains.buildServer.configs.kotlin.v2019_2.BuildType
import jetbrains.buildServer.configs.kotlin.v2019_2.Project

const val providerName = "google-beta"

fun GoogleBeta(environment: String, configuration : ClientConfiguration) : Project {
    return Project{
        vcsRoot(providerRepository)

        var buildConfigs = buildConfigurationsForPackages(packages, providerName, "google-beta", environment, configuration)
        buildConfigs.forEach { buildConfiguration ->
            buildType(buildConfiguration)
        }
    }
}

fun buildConfigurationsForPackages(packages: Map<String, String>, providerName : String, path : String, environment: String, config : ClientConfiguration): List<BuildType> {
    var list = ArrayList<BuildType>()

    packages.forEach { (packageName, displayName) ->
        if (packageName == "services") {
            buildConfigurationsForPackages(services, providerName, path+"/"+packageName, environment, config)
            continue
        } else {
            var defaultTestConfig = testConfiguration()

            var package = packageDetails(packageName, displayName, environment)
            var buildConfig = package.buildConfiguration(providerName, path, runNightly, testConfig.startHour, testConfig.parallelism, testConfig.daysOfWeek, testConfig.daysOfMonth)

            buildConfig.params.ConfigureGoogleSpecificTestParameters(environment, config)

            list.add(buildConfig)
        }
    }

    return list
}

class testConfiguration(parallelism: Int = defaultParallelism, startHour: Int = defaultStartHour, daysOfWeek: String = defaultDaysOfWeek, daysOfMonth: String = defaultDaysOfMonth, useAltSubscription: Boolean = false, useDevTestSubscription: Boolean = false) {
    var parallelism = parallelism
    var startHour = startHour
    var daysOfWeek = daysOfWeek
    var daysOfMonth = daysOfMonth
    var useAltSubscription = useAltSubscription
    var useDevTestSubscription = useDevTestSubscription
}
