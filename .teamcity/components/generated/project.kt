// this file is auto-generated with mmv1, any changes made here will be overwritten

import jetbrains.buildServer.configs.kotlin.v2019_2.BuildType
import jetbrains.buildServer.configs.kotlin.v2019_2.Project
import jetbrains.buildServer.configs.kotlin.v2019_2.projectFeatures.VersionedSettings

const val providerName = "google-beta"

fun GoogleBeta(environment: String, configuration : ClientConfiguration) : Project {
    return Project{
        vcsRoot(providerRepository)

        var buildConfigs = buildConfigurationsForPackages(packages, providerName, "google-beta", environment, configuration)
        buildConfigs.forEach { buildConfiguration ->
            buildType(buildConfiguration)
        }

        versionedSettings(vs)
    }
}

fun buildConfigurationsForPackages(packages: Map<String, String>, providerName : String, path : String, environment: String, config : ClientConfiguration): List<BuildType> {
    var list = ArrayList<BuildType>()

    packages.forEach { (packageName, displayName) ->
        if (packageName == "services") {
            buildConfigurationsForPackages(services, providerName, path+"/"+packageName, environment, config)
        } else {
            var defaultTestConfig = testConfiguration()

            var pkg = packageDetails(packageName, displayName, environment)
            var buildConfig = pkg.buildConfiguration(providerName, path, true, defaultTestConfig.startHour, defaultTestConfig.parallelism, defaultTestConfig.daysOfWeek, defaultTestConfig.daysOfMonth)

            buildConfig.params.ConfigureGoogleSpecificTestParameters(environment, config)

            list.add(buildConfig)
        }
    }

    return list
}

class testConfiguration(parallelism: Int = defaultParallelism, startHour: Int = defaultStartHour, daysOfWeek: String = defaultDaysOfWeek, daysOfMonth: String = defaultDaysOfMonth) {
    var parallelism = parallelism
    var startHour = startHour
    var daysOfWeek = daysOfWeek
    var daysOfMonth = daysOfMonth
}

object vs : VersionedSettings({
    storeSecureParamsOutsideOfVcs = true
    allowEditingOfProjectSettings = true
})
