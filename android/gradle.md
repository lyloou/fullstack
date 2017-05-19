```java
apply plugin: 'com.android.application'

android {
    compileSdkVersion 25
    buildToolsVersion '25.0.2'
    defaultConfig {
        applicationId "com.example.demo"
        minSdkVersion 15
        targetSdkVersion 25
        versionCode 1
        versionName "1.0"
        testInstrumentationRunner "android.support.test.runner.AndroidJUnitRunner"
    }

    signingConfigs {
        debug {
        }

        release {
            try {
                storeFile file(STORE_FILE)
                storePassword STORE_PASSWORD
                keyAlias KEY_ALIAS
                keyPassword KEY_PASSWORD
            } catch (ex) {
                throw new InvalidPreferencesFormatException("You should define KEYSTORE_PASSWORD and KEY_PASSWORD in gradle.properties." + ex)
            }

        }
    }

    buildTypes {
        debug {
            minifyEnabled false
            applicationIdSuffix '.debug'
            versionNameSuffix '-DEBUG'
            signingConfig signingConfigs.debug
            resValue 'string', 'app_name', 'Demo-Debug'
            buildConfigField "boolean", "TEST_ENV", "true"
        }

        release {
            minifyEnabled true
            proguardFiles getDefaultProguardFile('proguard-android.txt'), 'proguard-rules.pro'
            shrinkResources true
            zipAlignEnabled true
            signingConfig signingConfigs.release
            resValue 'string', 'app_name', 'Demo'
            buildConfigField "boolean", "TEST_ENV", "false"
        }
    }

    applicationVariants.all { variant ->
        variant.outputs.each { output ->
            def outputFile = output.outputFile
            if (outputFile != null) {
                if(outputFile.name.endsWith('-debug.apk')){
                    def fileName = "Demo_${releaseTime()}_V${defaultConfig.versionName}_DEBUG.apk"
                    output.outputFile = new File(outputFile.parent, fileName)
                } else if(outputFile.name.endsWith('-release.apk')){
                    def fileName = "Demo_${releaseTime()}_V${defaultConfig.versionName}.apk"
                    output.outputFile = new File(outputFile.parent, fileName)
                }

            }

        }

    }
}

def releaseTime() {
    return new Date().format("yyyMMdd", TimeZone.getTimeZone("UTC"))
}
```
