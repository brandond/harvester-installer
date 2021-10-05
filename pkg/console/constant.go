package console

const (
	titlePanel            = "title"
	debugPanel            = "debug"
	diskPanel             = "disk"
	askCreatePanel        = "askCreate"
	serverURLPanel        = "serverUrl"
	passwordPanel         = "osPassword"
	passwordConfirmPanel  = "osPasswordConfirm"
	sshKeyPanel           = "sshKey"
	tokenPanel            = "token"
	proxyPanel            = "proxy"
	askInterfacePanel     = "askInterface"
	askBondModePanel      = "askBondMode"
	bondNotePanel         = "bondNote"
	askNetworkMethodPanel = "askNetworkMethod"
	hostNamePanel         = "hostname"
	addressPanel          = "address"
	gatewayPanel          = "gateway"
	dnsServersPanel       = "dnsServers"
	networkValidatorPanel = "networkValidator"
	cloudInitPanel        = "cloudInit"
	validatorPanel        = "validator"
	notePanel             = "note"
	installPanel          = "install"
	footerPanel           = "footer"
	spinnerPanel          = "spinner"
	confirmInstallPanel   = "confirmInstall"
	confirmUpgradePanel   = "confirmUpgrade"
	upgradePanel          = "upgrade"
	askVipMethodPanel     = "askVipMethodPanel"
	vipPanel              = "vipPanel"
	vipTextPanel          = "vipTextPanel"
	ntpServersPanel       = "ntpServersPanel"

	networkTitle          = "Configure network connection"
	askBondModeLabel      = "Bond Mode"
	askInterfaceLabel     = "Management Bond"
	askNetworkMethodLabel = "IPv4 Method"
	hostNameLabel         = "HostName"
	addressLabel          = "IPv4 Address"
	gatewayLabel          = "Gateway"
	dnsServersLabel       = "DNS Servers"
	ntpServersLabel       = "NTP Servers"

	networkMethodDHCPText   = "Automatic (DHCP)"
	networkMethodStaticText = "Static"

	vipTitle          = "Configure VIP"
	vipLabel          = "VIP"
	askVipMethodLabel = "VIP Mode"

	clusterTokenCreateNote = "Note: The token is used for adding nodes to the cluster"
	clusterTokenJoinNote   = "Note: Input the token of the existing cluster"
	serverURLNote          = "Note: Input IP/domain name of the management node"
	proxyNote              = "Note: In the form of \"http://[[user][:pass]@]host[:port]/\"."
	sshKeyNote             = "For example: https://github.com/<username>.keys"
	ntpServersNote         = "Note: You can use comma to add more NTP servers."
	dnsServersNote         = "Note: You can use comma to add more DNS servers."
	bondNote               = "Note: Bond Mode can be ignored for a single NIC."

	authorizedFile = "/home/rancher/.ssh/authorized_keys"
)
