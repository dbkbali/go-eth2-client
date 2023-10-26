module github.com/dbkbali/go-eth2-client

go 1.20

require (
	github.com/ferranbt/fastssz v0.1.3
	github.com/goccy/go-yaml v1.9.2
	github.com/golang/snappy v0.0.4
	github.com/holiman/uint256 v1.2.2
	github.com/huandu/go-clone/generic v1.6.0
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.16.0
	github.com/prysmaticlabs/go-bitfield v0.0.0-20210809151128-385d8c5e3fb7
	github.com/r3labs/sse/v2 v2.10.0
	github.com/rs/zerolog v1.29.1
	github.com/stretchr/testify v1.8.4
	go.opentelemetry.io/otel v1.16.0
	golang.org/x/crypto v0.10.0
)

require (
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/cilium/ebpf v0.12.2 // indirect
	github.com/cosiner/argv v0.1.0 // indirect
	github.com/cpuguy83/go-md2man/v2 v2.0.3 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/derekparker/trie v0.0.0-20230829180723-39f4de51ef7d // indirect
	github.com/fatih/color v1.10.0 // indirect
	github.com/go-delve/delve v1.21.1 // indirect
	github.com/go-delve/liner v1.2.3-0.20220127212407-d32d89dd2a5d // indirect
	github.com/go-logr/logr v1.2.4 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/google/go-dap v0.11.0 // indirect
	github.com/hashicorp/golang-lru v1.0.2 // indirect
	github.com/huandu/go-clone v1.6.0 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/klauspost/cpuid/v2 v2.0.9 // indirect
	github.com/konsorten/go-windows-terminal-sequences v1.0.3 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mattn/go-runewidth v0.0.15 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.4 // indirect
	github.com/minio/sha256-simd v1.0.0 // indirect
	github.com/mitchellh/mapstructure v1.3.2 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/prometheus/client_model v0.3.0 // indirect
	github.com/prometheus/common v0.42.0 // indirect
	github.com/prometheus/procfs v0.10.1 // indirect
	github.com/rivo/uniseg v0.4.4 // indirect
	github.com/rogpeppe/go-internal v1.11.0 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/shurcooL/sanitized_anchor_name v1.0.0 // indirect
	github.com/sirupsen/logrus v1.9.3 // indirect
	github.com/spf13/cobra v1.7.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	go.opentelemetry.io/otel/metric v1.16.0 // indirect
	go.opentelemetry.io/otel/trace v1.16.0 // indirect
	go.starlark.net v0.0.0-20231016134836-22325403fcb3 // indirect
	golang.org/x/arch v0.5.0 // indirect
	golang.org/x/exp v0.0.0-20231006140011-7918f672742d // indirect
	golang.org/x/net v0.10.0 // indirect
	golang.org/x/sys v0.13.0 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
	google.golang.org/protobuf v1.30.0 // indirect
	gopkg.in/cenkalti/backoff.v1 v1.1.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

retract (
	v1.15.2 // Retraction for 1.15.1.
	v1.15.1 // Incorrect release number.
)
