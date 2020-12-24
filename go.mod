module github.com/bi-zone/ruwireguard-windows

go 1.15

require (
	github.com/bi-zone/ruwireguard-go v0.0.0-20201222151552-0de9ac51051e
	github.com/lxn/walk v0.0.0-20201209144500-98655d01b2f1
	github.com/lxn/win v0.0.0-20201111105847-2a20daff6a55
	golang.org/x/crypto v0.0.0-20201208171446-5f87f3452ae9
	golang.org/x/net v0.0.0-20201216054612-986b41b23924 // indirect
	golang.org/x/sys v0.0.0-20201223074533-0d417f636930
	golang.org/x/text v0.3.5-0.20201208001344-75a595aef632
)

replace (
	github.com/lxn/walk => golang.zx2c4.com/wireguard/windows v0.0.0-20201130211600-76ef01985b1c
	github.com/lxn/win => golang.zx2c4.com/wireguard/windows v0.0.0-20201107183008-659a4e955570
)
