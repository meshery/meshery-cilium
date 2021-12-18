module github.com/layer5io/meshery-cilium

go 1.13

replace (
	github.com/kudobuilder/kuttl => github.com/layer5io/kuttl v0.4.1-0.20200723152044-916f10574334
	github.com/spf13/afero => github.com/spf13/afero v1.5.1 // Until viper bug is resolved #1161
	gopkg.in/ini.v1 => github.com/go-ini/ini v1.62.0
)

require (
	github.com/layer5io/meshery-adapter-library v0.1.25
	github.com/layer5io/meshery-cilium v0.1.0
	github.com/layer5io/meshkit v0.2.34
	github.com/layer5io/service-mesh-performance v0.3.3
)
