load("@envoy_mobile//bazel:kotlin_test.bzl", "envoy_mobile_android_test", "envoy_mobile_kt_test")

licenses(["notice"])  # Apache 2

envoy_mobile_kt_test(
    name = "envoy_configuration_test",
    srcs = [
        "EnvoyConfigurationTest.kt",
    ],
    deps = [
        "//library/kotlin/io/envoyproxy/envoymobile:envoy_interfaces_lib",
    ],
)

envoy_mobile_kt_test(
    name = "envoy_native_resource_registry_test",
    srcs = [
        "EnvoyNativeResourceRegistryTest.kt",
    ],
    deps = [
        "//library/java/io/envoyproxy/envoymobile/engine:envoy_base_engine_lib",
    ],
)

envoy_mobile_kt_test(
    name = "jvm_bridge_utility_test",
    srcs = [
        "JvmBridgeUtilityTest.kt",
    ],
    deps = [
        "//library/java/io/envoyproxy/envoymobile/engine:envoy_base_engine_lib",
    ],
)

envoy_mobile_android_test(
    name = "android_network_monitor_tests",
    srcs = [
        "AndroidNetworkMonitorTest.java",
    ],
    native_deps = [
        "//library/common/jni:libndk_envoy_jni.so",
        "//library/common/jni:libndk_envoy_jni_jnilib",
    ],
    deps = [
        "//library/java/io/envoyproxy/envoymobile/engine:envoy_base_engine_lib",
        "//library/java/io/envoyproxy/envoymobile/engine:envoy_engine_lib",
        "//library/kotlin/io/envoyproxy/envoymobile:envoy_interfaces_lib",
    ],
)
