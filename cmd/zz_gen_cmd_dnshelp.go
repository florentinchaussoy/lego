package cmd

// CODE GENERATED AUTOMATICALLY
// THIS FILE MUST NOT BE EDITED BY HAND

import (
	"fmt"
	"io"
	"sort"
	"strings"
	"text/tabwriter"
)

func allDNSCodes() string {
	providers := []string{
		"manual",
		"acme-dns",
		"alidns",
		"allinkl",
		"arvancloud",
		"auroradns",
		"autodns",
		"azure",
		"bindman",
		"bluecat",
		"checkdomain",
		"civo",
		"clouddns",
		"cloudflare",
		"cloudns",
		"cloudxns",
		"conoha",
		"constellix",
		"desec",
		"designate",
		"digitalocean",
		"dnsimple",
		"dnsmadeeasy",
		"dnspod",
		"dode",
		"domeneshop",
		"dreamhost",
		"duckdns",
		"dyn",
		"dynu",
		"easydns",
		"edgedns",
		"epik",
		"exec",
		"exoscale",
		"freemyip",
		"gandi",
		"gandiv5",
		"gcloud",
		"gcore",
		"glesys",
		"godaddy",
		"hetzner",
		"hostingde",
		"hosttech",
		"httpreq",
		"hurricane",
		"hyperone",
		"ibmcloud",
		"iij",
		"iijdpf",
		"infoblox",
		"infomaniak",
		"internetbs",
		"inwx",
		"ionos",
		"iwantmyname",
		"joker",
		"liara",
		"lightsail",
		"linode",
		"liquidweb",
		"loopia",
		"luadns",
		"mydnsjp",
		"mythicbeasts",
		"namecheap",
		"namedotcom",
		"namesilo",
		"nearlyfreespeech",
		"netcup",
		"netlify",
		"nicmanager",
		"nifcloud",
		"njalla",
		"ns1",
		"oraclecloud",
		"otc",
		"ovh",
		"pdns",
		"porkbun",
		"rackspace",
		"regru",
		"rfc2136",
		"rimuhosting",
		"route53",
		"safedns",
		"sakuracloud",
		"scaleway",
		"selectel",
		"servercow",
		"simply",
		"sonic",
		"stackpath",
		"tencentcloud",
		"transip",
		"ultradns",
		"variomedia",
		"vegadns",
		"vercel",
		"versio",
		"vinyldns",
		"vkcloud",
		"vscale",
		"vultr",
		"wedos",
		"yandex",
		"yandexcloud",
		"zoneee",
		"zonomi",
	}
	sort.Strings(providers)
	return strings.Join(providers, ", ")
}

func displayDNSHelp(w io.Writer, name string) error {
	w = tabwriter.NewWriter(w, 0, 0, 2, ' ', 0)
	ew := &errWriter{w: w}

	switch name {
	case "acme-dns":
		// generated from: providers/dns/acmedns/acmedns.toml
		ew.writeln(`Configuration for Joohoi's ACME-DNS.`)
		ew.writeln(`Code:	'acme-dns'`)
		ew.writeln(`Since:	'v1.1.0'`)
		ew.writeln()

		ew.writeln(`Credentials:`)
		ew.writeln(`	- "ACME_DNS_API_BASE":	The ACME-DNS API address`)
		ew.writeln(`	- "ACME_DNS_STORAGE_PATH":	The ACME-DNS JSON account data file. A per-domain account will be registered/persisted to this file and used for TXT updates.`)
		ew.writeln()

		ew.writeln()
		ew.writeln(`More information: https://go-acme.github.io/lego/dns/acme-dns`)

	case "alidns":
		// generated from: providers/dns/alidns/alidns.toml
		ew.writeln(`Configuration for Alibaba Cloud DNS.`)
		ew.writeln(`Code:	'alidns'`)
		ew.writeln(`Since:	'v1.1.0'`)
		ew.writeln()

		ew.writeln(`Credentials:`)
		ew.writeln(`	- "ALICLOUD_ACCESS_KEY":	Access key ID`)
		ew.writeln(`	- "ALICLOUD_RAM_ROLE":	Your instance RAM role (https://www.alibabacloud.com/help/doc-detail/54579.htm)`)
		ew.writeln(`	- "ALICLOUD_SECRET_KEY":	Access Key secret`)
		ew.writeln(`	- "ALICLOUD_SECURITY_TOKEN":	STS Security Token (optional)`)
		ew.writeln()

		ew.writeln(`Additional Configuration:`)
		ew.writeln(`	- "ALICLOUD_HTTP_TIMEOUT":	API request timeout`)
		ew.writeln(`	- "ALICLOUD_POLLING_INTERVAL":	Time between DNS propagation check`)
		ew.writeln(`	- "ALICLOUD_PROPAGATION_TIMEOUT":	Maximum waiting time for DNS propagation`)
		ew.writeln(`	- "ALICLOUD_TTL":	The TTL of the TXT record used for the DNS challenge`)

		ew.writeln()
		ew.writeln(`More information: https://go-acme.github.io/lego/dns/alidns`)

	case "allinkl":
		// generated from: providers/dns/allinkl/allinkl.toml
		ew.writeln(`Configuration for all-inkl.`)
		ew.writeln(`Code:	'allinkl'`)
		ew.writeln(`Since:	'v4.5.0'`)
		ew.writeln()

		ew.writeln(`Credentials:`)
		ew.writeln(`	- "ALL_INKL_LOGIN":	KAS login`)
		ew.writeln(`	- "ALL_INKL_PASSWORD":	KAS password`)
		ew.writeln()

		ew.writeln(`Additional Configuration:`)
		ew.writeln(`	- "ALL_INKL_HTTP_TIMEOUT":	API request timeout`)
		ew.writeln(`	- "ALL_INKL_POLLING_INTERVAL":	Time between DNS propagation check`)
		ew.writeln(`	- "ALL_INKL_PROPAGATION_TIMEOUT":	Maximum waiting time for DNS propagation`)

		ew.writeln()
		ew.writeln(`More information: https://go-acme.github.io/lego/dns/allinkl`)

	case "arvancloud":
		// generated from: providers/dns/arvancloud/arvancloud.toml
		ew.writeln(`Configuration for ArvanCloud.`)
		ew.writeln(`Code:	'arvancloud'`)
		ew.writeln(`Since:	'v3.8.0'`)
		ew.writeln()

		ew.writeln(`Credentials:`)
		ew.writeln(`	- "ARVANCLOUD_API_KEY":	API key`)
		ew.writeln()

		ew.writeln(`Additional Configuration:`)
		ew.writeln(`	- "ARVANCLOUD_HTTP_TIMEOUT":	API request timeout`)
		ew.writeln(`	- "ARVANCLOUD_POLLING_INTERVAL":	Time between DNS propagation check`)
		ew.writeln(`	- "ARVANCLOUD_PROPAGATION_TIMEOUT":	Maximum waiting time for DNS propagation`)
		ew.writeln(`	- "ARVANCLOUD_TTL":	The TTL of the TXT record used for the DNS challenge`)

		ew.writeln()
		ew.writeln(`More information: https://go-acme.github.io/lego/dns/arvancloud`)

	case "auroradns":
		// generated from: providers/dns/auroradns/auroradns.toml
		ew.writeln(`Configuration for Aurora DNS.`)
		ew.writeln(`Code:	'auroradns'`)
		ew.writeln(`Since:	'v0.4.0'`)
		ew.writeln()

		ew.writeln(`Credentials:`)
		ew.writeln(`	- "AURORA_API_KEY":	API key or username to used`)
		ew.writeln(`	- "AURORA_SECRET":	Secret password to be used`)
		ew.writeln()

		ew.writeln(`Additional Configuration:`)
		ew.writeln(`	- "AURORA_ENDPOINT":	API endpoint URL`)
		ew.writeln(`	- "AURORA_POLLING_INTERVAL":	Time between DNS propagation check`)
		ew.writeln(`	- "AURORA_PROPAGATION_TIMEOUT":	Maximum waiting time for DNS propagation`)
		ew.writeln(`	- "AURORA_TTL":	The TTL of the TXT record used for the DNS challenge`)

		ew.writeln()
		ew.writeln(`More information: https://go-acme.github.io/lego/dns/auroradns`)

	case "autodns":
		// generated from: providers/dns/autodns/autodns.toml
		ew.writeln(`Configuration for Autodns.`)
		ew.writeln(`Code:	'autodns'`)
		ew.writeln(`Since:	'v3.2.0'`)
		ew.writeln()

		ew.writeln(`Credentials:`)
		ew.writeln(`	- "AUTODNS_API_PASSWORD":	User Password`)
		ew.writeln(`	- "AUTODNS_API_USER":	Username`)
		ew.writeln()

		ew.writeln(`Additional Configuration:`)
		ew.writeln(`	- "AUTODNS_CONTEXT":	API context (4 for production, 1 for testing. Defaults to 4)`)
		ew.writeln(`	- "AUTODNS_ENDPOINT":	API endpoint URL, defaults to https://api.autodns.com/v1/`)
		ew.writeln(`	- "AUTODNS_HTTP_TIMEOUT":	API request timeout, defaults to 30 seconds`)
		ew.writeln(`	- "AUTODNS_POLLING_INTERVAL":	Time between DNS propagation check`)
		ew.writeln(`	- "AUTODNS_PROPAGATION_TIMEOUT":	Maximum waiting time for DNS propagation`)
		ew.writeln(`	- "AUTODNS_TTL":	The TTL of the TXT record used for the DNS challenge`)

		ew.writeln()
		ew.writeln(`More information: https://go-acme.github.io/lego/dns/autodns`)

	case "azure":
		// generated from: providers/dns/azure/azure.toml
		ew.writeln(`Configuration for Azure.`)
		ew.writeln(`Code:	'azure'`)
		ew.writeln(`Since:	'v0.4.0'`)
		ew.writeln()

		ew.writeln(`Credentials:`)
		ew.writeln(`	- "AZURE_CLIENT_ID":	Client ID`)
		ew.writeln(`	- "AZURE_CLIENT_SECRET":	Client secret`)
		ew.writeln(`	- "AZURE_ENVIRONMENT":	Azure environment, one of: public, usgovernment, german, and china`)
		ew.writeln(`	- "AZURE_RESOURCE_GROUP":	Resource group`)
		ew.writeln(`	- "AZURE_SUBSCRIPTION_ID":	Subscription ID`)
		ew.writeln(`	- "AZURE_TENANT_ID":	Tenant ID`)
		ew.writeln(`	- "instance metadata service":	If the credentials are **not** set via the environment, then it will attempt to get a bearer token via the [instance metadata service](https://docs.microsoft.com/en-us/azure/virtual-machines/windows/instance-metadata-service).`)
		ew.writeln()

		ew.writeln(`Additional Configuration:`)
		ew.writeln(`	- "AZURE_METADATA_ENDPOINT":	Metadata Service endpoint URL`)
		ew.writeln(`	- "AZURE_POLLING_INTERVAL":	Time between DNS propagation check`)
		ew.writeln(`	- "AZURE_PRIVATE_ZONE":	Set to true to use Azure Private DNS Zones and not public`)
		ew.writeln(`	- "AZURE_PROPAGATION_TIMEOUT":	Maximum waiting time for DNS propagation`)
		ew.writeln(`	- "AZURE_TTL":	The TTL of the TXT record used for the DNS challenge`)
		ew.writeln(`	- "AZURE_ZONE_NAME":	Zone name to use inside Azure DNS service to add the TXT record in`)

		ew.writeln()
		ew.writeln(`More information: https://go-acme.github.io/lego/dns/azure`)

	case "bindman":
		// generated from: providers/dns/bindman/bindman.toml
		ew.writeln(`Configuration for Bindman.`)
		ew.writeln(`Code:	'bindman'`)
		ew.writeln(`Since:	'v2.6.0'`)
		ew.writeln()

		ew.writeln(`Credentials:`)
		ew.writeln(`	- "BINDMAN_MANAGER_ADDRESS":	The server URL, should have scheme, hostname, and port (if required) of the Bindman-DNS Manager server`)
		ew.writeln()

		ew.writeln(`Additional Configuration:`)
		ew.writeln(`	- "BINDMAN_HTTP_TIMEOUT":	API request timeout`)
		ew.writeln(`	- "BINDMAN_POLLING_INTERVAL":	Time between DNS propagation check`)
		ew.writeln(`	- "BINDMAN_PROPAGATION_TIMEOUT":	Maximum waiting time for DNS propagation`)

		ew.writeln()
		ew.writeln(`More information: https://go-acme.github.io/lego/dns/bindman`)

	case "bluecat":
		// generated from: providers/dns/bluecat/bluecat.toml
		ew.writeln(`Configuration for Bluecat.`)
		ew.writeln(`Code:	'bluecat'`)
		ew.writeln(`Since:	'v0.5.0'`)
		ew.writeln()

		ew.writeln(`Credentials:`)
		ew.writeln(`	- "BLUECAT_CONFIG_NAME":	Configuration name`)
		ew.writeln(`	- "BLUECAT_DNS_VIEW":	External DNS View Name`)
		ew.writeln(`	- "BLUECAT_PASSWORD":	API password`)
		ew.writeln(`	- "BLUECAT_SERVER_URL":	The server URL, should have scheme, hostname, and port (if required) of the authoritative Bluecat BAM serve`)
		ew.writeln(`	- "BLUECAT_USER_NAME":	API username`)
		ew.writeln()

		ew.writeln(`Additional Configuration:`)
		ew.writeln(`	- "BLUECAT_HTTP_TIMEOUT":	API request timeout`)
		ew.writeln(`	- "BLUECAT_POLLING_INTERVAL":	Time between DNS propagation check`)
		ew.writeln(`	- "BLUECAT_PROPAGATION_TIMEOUT":	Maximum waiting time for DNS propagation`)
		ew.writeln(`	- "BLUECAT_TTL":	The TTL of the TXT record used for the DNS challenge`)

		ew.writeln()
		ew.writeln(`More information: https://go-acme.github.io/lego/dns/bluecat`)

	case "checkdomain":
		// generated from: providers/dns/checkdomain/checkdomain.toml
		ew.writeln(`Configuration for Checkdomain.`)
		ew.writeln(`Code:	'checkdomain'`)
		ew.writeln(`Since:	'v3.3.0'`)
		ew.writeln()

		ew.writeln(`Credentials:`)
		ew.writeln(`	- "CHECKDOMAIN_TOKEN":	API token`)
		ew.writeln()

		ew.writeln(`Additional Configuration:`)
		ew.writeln(`	- "CHECKDOMAIN_ENDPOINT":	API endpoint URL, defaults to https://api.checkdomain.de`)
		ew.writeln(`	- "CHECKDOMAIN_HTTP_TIMEOUT":	API request timeout, defaults to 30 seconds`)
		ew.writeln(`	- "CHECKDOMAIN_POLLING_INTERVAL":	Time between DNS propagation check`)
		ew.writeln(`	- "CHECKDOMAIN_PROPAGATION_TIMEOUT":	Maximum waiting time for DNS propagation`)
		ew.writeln(`	- "CHECKDOMAIN_TTL":	The TTL of the TXT record used for the DNS challenge`)

		ew.writeln()
		ew.writeln(`More information: https://go-acme.github.io/lego/dns/checkdomain`)

	case "civo":
		// generated from: providers/dns/civo/civo.toml
		ew.writeln(`Configuration for Civo.`)
		ew.writeln(`Code:	'civo'`)
		ew.writeln(`Since:	'v4.9.0'`)
		ew.writeln()

		ew.writeln(`Credentials:`)
		ew.writeln(`	- "CIVO_TOKEN":	Authentication token`)
		ew.writeln()

		ew.writeln(`Additional Configuration:`)
		ew.writeln(`	- "CIVO_POLLING_INTERVAL":	Time between DNS propagation check`)
		ew.writeln(`	- "CIVO_PROPAGATION_TIMEOUT":	Maximum waiting time for DNS propagation`)
		ew.writeln(`	- "CIVO_TTL":	The TTL of the TXT record used for the DNS challenge`)

		ew.writeln()
		ew.writeln(`More information: https://go-acme.github.io/lego/dns/civo`)

	case "clouddns":
		// generated from: providers/dns/clouddns/clouddns.toml
		ew.writeln(`Configuration for CloudDNS.`)
		ew.writeln(`Code:	'clouddns'`)
		ew.writeln(`Since:	'v3.6.0'`)
		ew.writeln()

		ew.writeln(`Credentials:`)
		ew.writeln(`	- "CLOUDDNS_CLIENT_ID":	Client ID`)
		ew.writeln(`	- "CLOUDDNS_EMAIL":	Account email`)
		ew.writeln(`	- "CLOUDDNS_PASSWORD":	Account password`)
		ew.writeln()

		ew.writeln(`Additional Configuration:`)
		ew.writeln(`	- "CLOUDDNS_HTTP_TIMEOUT":	API request timeout`)
		ew.writeln(`	- "CLOUDDNS_POLLING_INTERVAL":	Time between DNS propagation check`)
		ew.writeln(`	- "CLOUDDNS_PROPAGATION_TIMEOUT":	Maximum waiting time for DNS propagation`)
		ew.writeln(`	- "CLOUDDNS_TTL":	The TTL of the TXT record used for the DNS challenge`)

		ew.writeln()
		ew.writeln(`More information: https://go-acme.github.io/lego/dns/clouddns`)

	case "cloudflare":
		// generated from: providers/dns/cloudflare/cloudflare.toml
		ew.writeln(`Configuration for Cloudflare.`)
		ew.writeln(`Code:	'cloudflare'`)
		ew.writeln(`Since:	'v0.3.0'`)
		ew.writeln()

		ew.writeln(`Credentials:`)
		ew.writeln(`	- "CF_API_EMAIL":	Account email`)
		ew.writeln(`	- "CF_API_KEY":	API key`)
		ew.writeln(`	- "CF_DNS_API_TOKEN":	API token with DNS:Edit permission (since v3.1.0)`)
		ew.writeln(`	- "CF_ZONE_API_TOKEN":	API token with Zone:Read permission (since v3.1.0)`)
		ew.writeln(`	- "CLOUDFLARE_API_KEY":	Alias to CF_API_KEY`)
		ew.writeln(`	- "CLOUDFLARE_DNS_API_TOKEN":	Alias to CF_DNS_API_TOKEN`)
		ew.writeln(`	- "CLOUDFLARE_EMAIL":	Alias to CF_API_EMAIL`)
		ew.writeln(`	- "CLOUDFLARE_ZONE_API_TOKEN":	Alias to CF_ZONE_API_TOKEN`)
		ew.writeln()

		ew.writeln(`Additional Configuration:`)
		ew.writeln(`	- "CLOUDFLARE_HTTP_TIMEOUT":	API request timeout`)
		ew.writeln(`	- "CLOUDFLARE_POLLING_INTERVAL":	Time between DNS propagation check`)
		ew.writeln(`	- "CLOUDFLARE_PROPAGATION_TIMEOUT":	Maximum waiting time for DNS propagation`)
		ew.writeln(`	- "CLOUDFLARE_TTL":	The TTL of the TXT record used for the DNS challenge`)

		ew.writeln()
		ew.writeln(`More information: https://go-acme.github.io/lego/dns/cloudflare`)

	case "cloudns":
		// generated from: providers/dns/cloudns/cloudns.toml
		ew.writeln(`Configuration for ClouDNS.`)
		ew.writeln(`Code:	'cloudns'`)
		ew.writeln(`Since:	'v2.3.0'`)
		ew.writeln()

		ew.writeln(`Credentials:`)
		ew.writeln(`	- "CLOUDNS_AUTH_ID":	The API user ID`)
		ew.writeln(`	- "CLOUDNS_AUTH_PASSWORD":	The password for API user ID`)
		ew.writeln()

		ew.writeln(`Additional Configuration:`)
		ew.writeln(`	- "CLOUDNS_HTTP_TIMEOUT":	API request timeout`)
		ew.writeln(`	- "CLOUDNS_POLLING_INTERVAL":	Time between DNS propagation check`)
		ew.writeln(`	- "CLOUDNS_PROPAGATION_TIMEOUT":	Maximum waiting time for DNS propagation`)
		ew.writeln(`	- "CLOUDNS_SUB_AUTH_ID":	The API sub user ID`)
		ew.writeln(`	- "CLOUDNS_TTL":	The TTL of the TXT record used for the DNS challenge`)

		ew.writeln()
		ew.writeln(`More information: https://go-acme.github.io/lego/dns/cloudns`)

	case "cloudxns":
		// generated from: providers/dns/cloudxns/cloudxns.toml
		ew.writeln(`Configuration for CloudXNS.`)
		ew.writeln(`Code:	'cloudxns'`)
		ew.writeln(`Since:	'v0.5.0'`)
		ew.writeln()

		ew.writeln(`Credentials:`)
		ew.writeln(`	- "CLOUDXNS_API_KEY":	The API key`)
		ew.writeln(`	- "CLOUDXNS_SECRET_KEY":	The API secret key`)
		ew.writeln()

		ew.writeln(`Additional Configuration:`)
		ew.writeln(`	- "CLOUDXNS_HTTP_TIMEOUT":	API request timeout`)
		ew.writeln(`	- "CLOUDXNS_POLLING_INTERVAL":	Time between DNS propagation check`)
		ew.writeln(`	- "CLOUDXNS_PROPAGATION_TIMEOUT":	Maximum waiting time for DNS propagation`)
		ew.writeln(`	- "CLOUDXNS_TTL":	The TTL of the TXT record used for the DNS challenge`)

		ew.writeln()
		ew.writeln(`More information: https://go-acme.github.io/lego/dns/cloudxns`)

	case "conoha":
		// generated from: providers/dns/conoha/conoha.toml
		ew.writeln(`Configuration for ConoHa.`)
		ew.writeln(`Code:	'conoha'`)
		ew.writeln(`Since:	'v1.2.0'`)
		ew.writeln()

		ew.writeln(`Credentials:`)
		ew.writeln(`	- "CONOHA_API_PASSWORD":	The API password`)
		ew.writeln(`	- "CONOHA_API_USERNAME":	The API username`)
		ew.writeln(`	- "CONOHA_TENANT_ID":	Tenant ID`)
		ew.writeln()

		ew.writeln(`Additional Configuration:`)
		ew.writeln(`	- "CONOHA_HTTP_TIMEOUT":	API request timeout`)
		ew.writeln(`	- "CONOHA_POLLING_INTERVAL":	Time between DNS propagation check`)
		ew.writeln(`	- "CONOHA_PROPAGATION_TIMEOUT":	Maximum waiting time for DNS propagation`)
		ew.writeln(`	- "CONOHA_REGION":	The region`)
		ew.writeln(`	- "CONOHA_TTL":	The TTL of the TXT record used for the DNS challenge`)

		ew.writeln()
		ew.writeln(`More information: https://go-acme.github.io/lego/dns/conoha`)

	case "constellix":
		// generated from: providers/dns/constellix/constellix.toml
		ew.writeln(`Configuration for Constellix.`)
		ew.writeln(`Code:	'constellix'`)
		ew.writeln(`Since:	'v3.4.0'`)
		ew.writeln()

		ew.writeln(`Credentials:`)
		ew.writeln(`	- "CONSTELLIX_API_KEY":	User API key`)
		ew.writeln(`	- "CONSTELLIX_SECRET_KEY":	User secret key`)
		ew.writeln()

		ew.writeln(`Additional Configuration:`)
		ew.writeln(`	- "CONSTELLIX_HTTP_TIMEOUT":	API request timeout`)
		ew.writeln(`	- "CONSTELLIX_POLLING_INTERVAL":	Time between DNS propagation check`)
		ew.writeln(`	- "CONSTELLIX_PROPAGATION_TIMEOUT":	Maximum waiting time for DNS propagation`)
		ew.writeln(`	- "CONSTELLIX_TTL":	The TTL of the TXT record used for the DNS challenge`)

		ew.writeln()
		ew.writeln(`More information: https://go-acme.github.io/lego/dns/constellix`)

	case "desec":
		// generated from: providers/dns/desec/desec.toml
		ew.writeln(`Configuration for deSEC.io.`)
		ew.writeln(`Code:	'desec'`)
		ew.writeln(`Since:	'v3.7.0'`)
		ew.writeln()

		ew.writeln(`Credentials:`)
		ew.writeln(`	- "DESEC_TOKEN":	Domain token`)
		ew.writeln()

		ew.writeln(`Additional Configuration:`)
		ew.writeln(`	- "DESEC_HTTP_TIMEOUT":	API request timeout`)
		ew.writeln(`	- "DESEC_POLLING_INTERVAL":	Time between DNS propagation check`)
		ew.writeln(`	- "DESEC_PROPAGATION_TIMEOUT":	Maximum waiting time for DNS propagation`)
		ew.writeln(`	- "DESEC_TTL":	The TTL of the TXT record used for the DNS challenge`)

		ew.writeln()
		ew.writeln(`More information: https://go-acme.github.io/lego/dns/desec`)

	case "designate":
		// generated from: providers/dns/designate/designate.toml
		ew.writeln(`Configuration for Designate DNSaaS for Openstack.`)
		ew.writeln(`Code:	'designate'`)
		ew.writeln(`Since:	'v2.2.0'`)
		ew.writeln()

		ew.writeln(`Credentials:`)
		ew.writeln(`	- "OS_APPLICATION_CREDENTIAL_ID":	Application credential ID`)
		ew.writeln(`	- "OS_APPLICATION_CREDENTIAL_NAME":	Application credential name`)
		ew.writeln(`	- "OS_APPLICATION_CREDENTIAL_SECRET":	Application credential secret`)
		ew.writeln(`	- "OS_AUTH_URL":	Identity endpoint URL`)
		ew.writeln(`	- "OS_PASSWORD":	Password`)
		ew.writeln(`	- "OS_PROJECT_NAME":	Project name`)
		ew.writeln(`	- "OS_REGION_NAME":	Region name`)
		ew.writeln(`	- "OS_USERNAME":	Username`)
		ew.writeln(`	- "OS_USER_ID":	User ID`)
		ew.writeln()

		ew.writeln(`Additional Configuration:`)
		ew.writeln(`	- "DESIGNATE_POLLING_INTERVAL":	Time between DNS propagation check`)
		ew.writeln(`	- "DESIGNATE_PROPAGATION_TIMEOUT":	Maximum waiting time for DNS propagation`)
		ew.writeln(`	- "DESIGNATE_TTL":	The TTL of the TXT record used for the DNS challenge`)
		ew.writeln(`	- "OS_PROJECT_ID":	Project ID`)
		ew.writeln(`	- "OS_TENANT_NAME":	Tenant name (deprecated see OS_PROJECT_NAME and OS_PROJECT_ID)`)

		ew.writeln()
		ew.writeln(`More information: https://go-acme.github.io/lego/dns/designate`)

	case "digitalocean":
		// generated from: providers/dns/digitalocean/digitalocean.toml
		ew.writeln(`Configuration for Digital Ocean.`)
		ew.writeln(`Code:	'digitalocean'`)
		ew.writeln(`Since:	'v0.3.0'`)
		ew.writeln()

		ew.writeln(`Credentials:`)
		ew.writeln(`	- "DO_AUTH_TOKEN":	Authentication token`)
		ew.writeln()

		ew.writeln(`Additional Configuration:`)
		ew.writeln(`	- "DO_API_URL":	The URL of the API`)
		ew.writeln(`	- "DO_HTTP_TIMEOUT":	API request timeout`)
		ew.writeln(`	- "DO_POLLING_INTERVAL":	Time between DNS propagation check`)
		ew.writeln(`	- "DO_PROPAGATION_TIMEOUT":	Maximum waiting time for DNS propagation`)
		ew.writeln(`	- "DO_TTL":	The TTL of the TXT record used for the DNS challenge`)

		ew.writeln()
		ew.writeln(`More information: https://go-acme.github.io/lego/dns/digitalocean`)

	case "dnsimple":
		// generated from: providers/dns/dnsimple/dnsimple.toml
		ew.writeln(`Configuration for DNSimple.`)
		ew.writeln(`Code:	'dnsimple'`)
		ew.writeln(`Since:	'v0.3.0'`)
		ew.writeln()

		ew.writeln(`Credentials:`)
		ew.writeln(`	- "DNSIMPLE_OAUTH_TOKEN":	OAuth token`)
		ew.writeln()

		ew.writeln(`Additional Configuration:`)
		ew.writeln(`	- "DNSIMPLE_BASE_URL":	API endpoint URL`)
		ew.writeln(`	- "DNSIMPLE_POLLING_INTERVAL":	Time between DNS propagation check`)
		ew.writeln(`	- "DNSIMPLE_PROPAGATION_TIMEOUT":	Maximum waiting time for DNS propagation`)
		ew.writeln(`	- "DNSIMPLE_TTL":	The TTL of the TXT record used for the DNS challenge`)

		ew.writeln()
		ew.writeln(`More information: https://go-acme.github.io/lego/dns/dnsimple`)

	case "dnsmadeeasy":
		// generated from: providers/dns/dnsmadeeasy/dnsmadeeasy.toml
		ew.writeln(`Configuration for DNS Made Easy.`)
		ew.writeln(`Code:	'dnsmadeeasy'`)
		ew.writeln(`Since:	'v0.4.0'`)
		ew.writeln()

		ew.writeln(`Credentials:`)
		ew.writeln(`	- "DNSMADEEASY_API_KEY":	The API key`)
		ew.writeln(`	- "DNSMADEEASY_API_SECRET":	The API Secret key`)
		ew.writeln()

		ew.writeln(`Additional Configuration:`)
		ew.writeln(`	- "DNSMADEEASY_HTTP_TIMEOUT":	API request timeout`)
		ew.writeln(`	- "DNSMADEEASY_POLLING_INTERVAL":	Time between DNS propagation check`)
		ew.writeln(`	- "DNSMADEEASY_PROPAGATION_TIMEOUT":	Maximum waiting time for DNS propagation`)
		ew.writeln(`	- "DNSMADEEASY_SANDBOX":	Activate the sandbox (boolean)`)
		ew.writeln(`	- "DNSMADEEASY_TTL":	The TTL of the TXT record used for the DNS challenge`)

		ew.writeln()
		ew.writeln(`More information: https://go-acme.github.io/lego/dns/dnsmadeeasy`)

	case "dnspod":
		// generated from: providers/dns/dnspod/dnspod.toml
		ew.writeln(`Configuration for DNSPod (deprecated).`)
		ew.writeln(`Code:	'dnspod'`)
		ew.writeln(`Since:	'v0.4.0'`)
		ew.writeln()

		ew.writeln(`Credentials:`)
		ew.writeln(`	- "DNSPOD_API_KEY":	The user token`)
		ew.writeln()

		ew.writeln(`Additional Configuration:`)
		ew.writeln(`	- "DNSPOD_HTTP_TIMEOUT":	API request timeout`)
		ew.writeln(`	- "DNSPOD_POLLING_INTERVAL":	Time between DNS propagation check`)
		ew.writeln(`	- "DNSPOD_PROPAGATION_TIMEOUT":	Maximum waiting time for DNS propagation`)
		ew.writeln(`	- "DNSPOD_TTL":	The TTL of the TXT record used for the DNS challenge`)

		ew.writeln()
		ew.writeln(`More information: https://go-acme.github.io/lego/dns/dnspod`)

	case "dode":
		// generated from: providers/dns/dode/dode.toml
		ew.writeln(`Configuration for Domain Offensive (do.de).`)
		ew.writeln(`Code:	'dode'`)
		ew.writeln(`Since:	'v2.4.0'`)
		ew.writeln()

		ew.writeln(`Credentials:`)
		ew.writeln(`	- "DODE_TOKEN":	API token`)
		ew.writeln()

		ew.writeln(`Additional Configuration:`)
		ew.writeln(`	- "DODE_HTTP_TIMEOUT":	API request timeout`)
		ew.writeln(`	- "DODE_POLLING_INTERVAL":	Time between DNS propagation check`)
		ew.writeln(`	- "DODE_PROPAGATION_TIMEOUT":	Maximum waiting time for DNS propagation`)
		ew.writeln(`	- "DODE_SEQUENCE_INTERVAL":	Time between sequential requests`)
		ew.writeln(`	- "DODE_TTL":	The TTL of the TXT record used for the DNS challenge`)

		ew.writeln()
		ew.writeln(`More information: https://go-acme.github.io/lego/dns/dode`)

	case "domeneshop":
		// generated from: providers/dns/domeneshop/domeneshop.toml
		ew.writeln(`Configuration for Domeneshop.`)
		ew.writeln(`Code:	'domeneshop'`)
		ew.writeln(`Since:	'v4.3.0'`)
		ew.writeln()

		ew.writeln(`Credentials:`)
		ew.writeln(`	- "DOMENESHOP_API_SECRET":	API secret`)
		ew.writeln(`	- "DOMENESHOP_API_TOKEN":	API token`)
		ew.writeln()

		ew.writeln(`Additional Configuration:`)
		ew.writeln(`	- "DOMENESHOP_HTTP_TIMEOUT":	API request timeout`)
		ew.writeln(`	- "DOMENESHOP_POLLING_INTERVAL":	Time between DNS propagation check`)
		ew.writeln(`	- "DOMENESHOP_PROPAGATION_TIMEOUT":	Maximum waiting time for DNS propagation`)

		ew.writeln()
		ew.writeln(`More information: https://go-acme.github.io/lego/dns/domeneshop`)

	case "dreamhost":
		// generated from: providers/dns/dreamhost/dreamhost.toml
		ew.writeln(`Configuration for DreamHost.`)
		ew.writeln(`Code:	'dreamhost'`)
		ew.writeln(`Since:	'v1.1.0'`)
		ew.writeln()

		ew.writeln(`Credentials:`)
		ew.writeln(`	- "DREAMHOST_API_KEY":	The API key`)
		ew.writeln()

		ew.writeln(`Additional Configuration:`)
		ew.writeln(`	- "DREAMHOST_HTTP_TIMEOUT":	API request timeout`)
		ew.writeln(`	- "DREAMHOST_POLLING_INTERVAL":	Time between DNS propagation check`)
		ew.writeln(`	- "DREAMHOST_PROPAGATION_TIMEOUT":	Maximum waiting time for DNS propagation`)
		ew.writeln(`	- "DREAMHOST_TTL":	The TTL of the TXT record used for the DNS challenge`)

		ew.writeln()
		ew.writeln(`More information: https://go-acme.github.io/lego/dns/dreamhost`)

	case "duckdns":
		// generated from: providers/dns/duckdns/duckdns.toml
		ew.writeln(`Configuration for Duck DNS.`)
		ew.writeln(`Code:	'duckdns'`)
		ew.writeln(`Since:	'v0.5.0'`)
		ew.writeln()

		ew.writeln(`Credentials:`)
		ew.writeln(`	- "DUCKDNS_TOKEN":	Account token`)
		ew.writeln()

		ew.writeln(`Additional Configuration:`)
		ew.writeln(`	- "DUCKDNS_HTTP_TIMEOUT":	API request timeout`)
		ew.writeln(`	- "DUCKDNS_POLLING_INTERVAL":	Time between DNS propagation check`)
		ew.writeln(`	- "DUCKDNS_PROPAGATION_TIMEOUT":	Maximum waiting time for DNS propagation`)
		ew.writeln(`	- "DUCKDNS_SEQUENCE_INTERVAL":	Time between sequential requests`)
		ew.writeln(`	- "DUCKDNS_TTL":	The TTL of the TXT record used for the DNS challenge`)

		ew.writeln()
		ew.writeln(`More information: https://go-acme.github.io/lego/dns/duckdns`)

	case "dyn":
		// generated from: providers/dns/dyn/dyn.toml
		ew.writeln(`Configuration for Dyn.`)
		ew.writeln(`Code:	'dyn'`)
		ew.writeln(`Since:	'v0.3.0'`)
		ew.writeln()

		ew.writeln(`Credentials:`)
		ew.writeln(`	- "DYN_CUSTOMER_NAME":	Customer name`)
		ew.writeln(`	- "DYN_PASSWORD":	Password`)
		ew.writeln(`	- "DYN_USER_NAME":	User name`)
		ew.writeln()

		ew.writeln(`Additional Configuration:`)
		ew.writeln(`	- "DYN_HTTP_TIMEOUT":	API request timeout`)
		ew.writeln(`	- "DYN_POLLING_INTERVAL":	Time between DNS propagation check`)
		ew.writeln(`	- "DYN_PROPAGATION_TIMEOUT":	Maximum waiting time for DNS propagation`)
		ew.writeln(`	- "DYN_TTL":	The TTL of the TXT record used for the DNS challenge`)

		ew.writeln()
		ew.writeln(`More information: https://go-acme.github.io/lego/dns/dyn`)

	case "dynu":
		// generated from: providers/dns/dynu/dynu.toml
		ew.writeln(`Configuration for Dynu.`)
		ew.writeln(`Code:	'dynu'`)
		ew.writeln(`Since:	'v3.5.0'`)
		ew.writeln()

		ew.writeln(`Credentials:`)
		ew.writeln(`	- "DYNU_API_KEY":	API key`)
		ew.writeln()

		ew.writeln(`Additional Configuration:`)
		ew.writeln(`	- "DYNU_HTTP_TIMEOUT":	API request timeout`)
		ew.writeln(`	- "DYNU_POLLING_INTERVAL":	Time between DNS propagation check`)
		ew.writeln(`	- "DYNU_PROPAGATION_TIMEOUT":	Maximum waiting time for DNS propagation`)
		ew.writeln(`	- "DYNU_TTL":	The TTL of the TXT record used for the DNS challenge`)

		ew.writeln()
		ew.writeln(`More information: https://go-acme.github.io/lego/dns/dynu`)

	case "easydns":
		// generated from: providers/dns/easydns/easydns.toml
		ew.writeln(`Configuration for EasyDNS.`)
		ew.writeln(`Code:	'easydns'`)
		ew.writeln(`Since:	'v2.6.0'`)
		ew.writeln()

		ew.writeln(`Credentials:`)
		ew.writeln(`	- "EASYDNS_KEY":	API Key`)
		ew.writeln(`	- "EASYDNS_TOKEN":	API Token`)
		ew.writeln()

		ew.writeln(`Additional Configuration:`)
		ew.writeln(`	- "EASYDNS_ENDPOINT":	The endpoint URL of the API Server`)
		ew.writeln(`	- "EASYDNS_HTTP_TIMEOUT":	API request timeout`)
		ew.writeln(`	- "EASYDNS_POLLING_INTERVAL":	Time between DNS propagation check`)
		ew.writeln(`	- "EASYDNS_PROPAGATION_TIMEOUT":	Maximum waiting time for DNS propagation`)
		ew.writeln(`	- "EASYDNS_SEQUENCE_INTERVAL":	Time between sequential requests`)
		ew.writeln(`	- "EASYDNS_TTL":	The TTL of the TXT record used for the DNS challenge`)

		ew.writeln()
		ew.writeln(`More information: https://go-acme.github.io/lego/dns/easydns`)

	case "edgedns":
		// generated from: providers/dns/edgedns/edgedns.toml
		ew.writeln(`Configuration for Akamai EdgeDNS.`)
		ew.writeln(`Code:	'edgedns'`)
		ew.writeln(`Since:	'v3.9.0'`)
		ew.writeln()

		ew.writeln(`Credentials:`)
		ew.writeln(`	- "AKAMAI_ACCESS_TOKEN":	Access token, managed by the Akamai EdgeGrid client`)
		ew.writeln(`	- "AKAMAI_CLIENT_SECRET":	Client secret, managed by the Akamai EdgeGrid client`)
		ew.writeln(`	- "AKAMAI_CLIENT_TOKEN":	Client token, managed by the Akamai EdgeGrid client`)
		ew.writeln(`	- "AKAMAI_EDGERC":	Path to the .edgerc file, managed by the Akamai EdgeGrid client`)
		ew.writeln(`	- "AKAMAI_EDGERC_SECTION":	Configuration section, managed by the Akamai EdgeGrid client`)
		ew.writeln(`	- "AKAMAI_HOST":	API host, managed by the Akamai EdgeGrid client`)
		ew.writeln()

		ew.writeln(`Additional Configuration:`)
		ew.writeln(`	- "AKAMAI_POLLING_INTERVAL":	Time between DNS propagation check. Default: 15 seconds`)
		ew.writeln(`	- "AKAMAI_PROPAGATION_TIMEOUT":	Maximum waiting time for DNS propagation. Default: 3 minutes`)
		ew.writeln(`	- "AKAMAI_TTL":	The TTL of the TXT record used for the DNS challenge`)

		ew.writeln()
		ew.writeln(`More information: https://go-acme.github.io/lego/dns/edgedns`)

	case "epik":
		// generated from: providers/dns/epik/epik.toml
		ew.writeln(`Configuration for Epik.`)
		ew.writeln(`Code:	'epik'`)
		ew.writeln(`Since:	'v4.5.0'`)
		ew.writeln()

		ew.writeln(`Credentials:`)
		ew.writeln(`	- "EPIK_SIGNATURE":	Epik API signature (https://registrar.epik.com/account/api-settings/)`)
		ew.writeln()

		ew.writeln(`Additional Configuration:`)
		ew.writeln(`	- "EPIK_HTTP_TIMEOUT":	API request timeout`)
		ew.writeln(`	- "EPIK_POLLING_INTERVAL":	Time between DNS propagation check`)
		ew.writeln(`	- "EPIK_PROPAGATION_TIMEOUT":	Maximum waiting time for DNS propagation`)
		ew.writeln(`	- "EPIK_TTL":	The TTL of the TXT record used for the DNS challenge`)

		ew.writeln()
		ew.writeln(`More information: https://go-acme.github.io/lego/dns/epik`)

	case "exec":
		// generated from: providers/dns/exec/exec.toml
		ew.writeln(`Configuration for External program.`)
		ew.writeln(`Code:	'exec'`)
		ew.writeln(`Since:	'v0.5.0'`)
		ew.writeln()

		ew.writeln()
		ew.writeln(`More information: https://go-acme.github.io/lego/dns/exec`)

	case "exoscale":
		// generated from: providers/dns/exoscale/exoscale.toml
		ew.writeln(`Configuration for Exoscale.`)
		ew.writeln(`Code:	'exoscale'`)
		ew.writeln(`Since:	'v0.4.0'`)
		ew.writeln()

		ew.writeln(`Credentials:`)
		ew.writeln(`	- "EXOSCALE_API_KEY":	API key`)
		ew.writeln(`	- "EXOSCALE_API_SECRET":	API secret`)
		ew.writeln()

		ew.writeln(`Additional Configuration:`)
		ew.writeln(`	- "EXOSCALE_API_ZONE":	API zone`)
		ew.writeln(`	- "EXOSCALE_ENDPOINT":	API endpoint URL`)
		ew.writeln(`	- "EXOSCALE_HTTP_TIMEOUT":	API request timeout`)
		ew.writeln(`	- "EXOSCALE_POLLING_INTERVAL":	Time between DNS propagation check`)
		ew.writeln(`	- "EXOSCALE_PROPAGATION_TIMEOUT":	Maximum waiting time for DNS propagation`)
		ew.writeln(`	- "EXOSCALE_TTL":	The TTL of the TXT record used for the DNS challenge`)

		ew.writeln()
		ew.writeln(`More information: https://go-acme.github.io/lego/dns/exoscale`)

	case "freemyip":
		// generated from: providers/dns/freemyip/freemyip.toml
		ew.writeln(`Configuration for freemyip.com.`)
		ew.writeln(`Code:	'freemyip'`)
		ew.writeln(`Since:	'v4.5.0'`)
		ew.writeln()

		ew.writeln(`Credentials:`)
		ew.writeln(`	- "FREEMYIP_TOKEN":	Account token`)
		ew.writeln()

		ew.writeln(`Additional Configuration:`)
		ew.writeln(`	- "FREEMYIP_HTTP_TIMEOUT":	API request timeout`)
		ew.writeln(`	- "FREEMYIP_POLLING_INTERVAL":	Time between DNS propagation check`)
		ew.writeln(`	- "FREEMYIP_PROPAGATION_TIMEOUT":	Maximum waiting time for DNS propagation`)
		ew.writeln(`	- "FREEMYIP_SEQUENCE_INTERVAL":	Time between sequential requests`)
		ew.writeln(`	- "FREEMYIP_TTL":	The TTL of the TXT record used for the DNS challenge`)

		ew.writeln()
		ew.writeln(`More information: https://go-acme.github.io/lego/dns/freemyip`)

	case "gandi":
		// generated from: providers/dns/gandi/gandi.toml
		ew.writeln(`Configuration for Gandi.`)
		ew.writeln(`Code:	'gandi'`)
		ew.writeln(`Since:	'v0.3.0'`)
		ew.writeln()

		ew.writeln(`Credentials:`)
		ew.writeln(`	- "GANDI_API_KEY":	API key`)
		ew.writeln()

		ew.writeln(`Additional Configuration:`)
		ew.writeln(`	- "GANDI_HTTP_TIMEOUT":	API request timeout`)
		ew.writeln(`	- "GANDI_POLLING_INTERVAL":	Time between DNS propagation check`)
		ew.writeln(`	- "GANDI_PROPAGATION_TIMEOUT":	Maximum waiting time for DNS propagation`)
		ew.writeln(`	- "GANDI_TTL":	The TTL of the TXT record used for the DNS challenge`)

		ew.writeln()
		ew.writeln(`More information: https://go-acme.github.io/lego/dns/gandi`)

	case "gandiv5":
		// generated from: providers/dns/gandiv5/gandiv5.toml
		ew.writeln(`Configuration for Gandi Live DNS (v5).`)
		ew.writeln(`Code:	'gandiv5'`)
		ew.writeln(`Since:	'v0.5.0'`)
		ew.writeln()

		ew.writeln(`Credentials:`)
		ew.writeln(`	- "GANDIV5_API_KEY":	API key`)
		ew.writeln()

		ew.writeln(`Additional Configuration:`)
		ew.writeln(`	- "GANDIV5_HTTP_TIMEOUT":	API request timeout`)
		ew.writeln(`	- "GANDIV5_POLLING_INTERVAL":	Time between DNS propagation check`)
		ew.writeln(`	- "GANDIV5_PROPAGATION_TIMEOUT":	Maximum waiting time for DNS propagation`)
		ew.writeln(`	- "GANDIV5_TTL":	The TTL of the TXT record used for the DNS challenge`)

		ew.writeln()
		ew.writeln(`More information: https://go-acme.github.io/lego/dns/gandiv5`)

	case "gcloud":
		// generated from: providers/dns/gcloud/gcloud.toml
		ew.writeln(`Configuration for Google Cloud.`)
		ew.writeln(`Code:	'gcloud'`)
		ew.writeln(`Since:	'v0.3.0'`)
		ew.writeln()

		ew.writeln(`Credentials:`)
		ew.writeln(`	- "Application Default Credentials":	[Documentation](https://cloud.google.com/docs/authentication/production#providing_credentials_to_your_application)`)
		ew.writeln(`	- "GCE_PROJECT":	Project name (by default, the project name is auto-detected by using the metadata service)`)
		ew.writeln(`	- "GCE_SERVICE_ACCOUNT":	Account`)
		ew.writeln(`	- "GCE_SERVICE_ACCOUNT_FILE":	Account file path`)
		ew.writeln()

		ew.writeln(`Additional Configuration:`)
		ew.writeln(`	- "GCE_ALLOW_PRIVATE_ZONE":	Allows requested domain to be in private DNS zone, works only with a private ACME server (by default: false)`)
		ew.writeln(`	- "GCE_POLLING_INTERVAL":	Time between DNS propagation check`)
		ew.writeln(`	- "GCE_PROPAGATION_TIMEOUT":	Maximum waiting time for DNS propagation`)
		ew.writeln(`	- "GCE_TTL":	The TTL of the TXT record used for the DNS challenge`)

		ew.writeln()
		ew.writeln(`More information: https://go-acme.github.io/lego/dns/gcloud`)

	case "gcore":
		// generated from: providers/dns/gcore/gcore.toml
		ew.writeln(`Configuration for G-Core Labs.`)
		ew.writeln(`Code:	'gcore'`)
		ew.writeln(`Since:	'v4.5.0'`)
		ew.writeln()

		ew.writeln(`Credentials:`)
		ew.writeln(`	- "GCORE_PERMANENT_API_TOKEN":	Permanent API tokene (https://gcorelabs.com/blog/permanent-api-token-explained/)`)
		ew.writeln()

		ew.writeln(`Additional Configuration:`)
		ew.writeln(`	- "GCORE_HTTP_TIMEOUT":	API request timeout`)
		ew.writeln(`	- "GCORE_POLLING_INTERVAL":	Time between DNS propagation check`)
		ew.writeln(`	- "GCORE_PROPAGATION_TIMEOUT":	Maximum waiting time for DNS propagation`)
		ew.writeln(`	- "GCORE_TTL":	The TTL of the TXT record used for the DNS challenge`)

		ew.writeln()
		ew.writeln(`More information: https://go-acme.github.io/lego/dns/gcore`)

	case "glesys":
		// generated from: providers/dns/glesys/glesys.toml
		ew.writeln(`Configuration for Glesys.`)
		ew.writeln(`Code:	'glesys'`)
		ew.writeln(`Since:	'v0.5.0'`)
		ew.writeln()

		ew.writeln(`Credentials:`)
		ew.writeln(`	- "GLESYS_API_KEY":	API key`)
		ew.writeln(`	- "GLESYS_API_USER":	API user`)
		ew.writeln()

		ew.writeln(`Additional Configuration:`)
		ew.writeln(`	- "GLESYS_HTTP_TIMEOUT":	API request timeout`)
		ew.writeln(`	- "GLESYS_POLLING_INTERVAL":	Time between DNS propagation check`)
		ew.writeln(`	- "GLESYS_PROPAGATION_TIMEOUT":	Maximum waiting time for DNS propagation`)
		ew.writeln(`	- "GLESYS_TTL":	The TTL of the TXT record used for the DNS challenge`)

		ew.writeln()
		ew.writeln(`More information: https://go-acme.github.io/lego/dns/glesys`)

	case "godaddy":
		// generated from: providers/dns/godaddy/godaddy.toml
		ew.writeln(`Configuration for Go Daddy.`)
		ew.writeln(`Code:	'godaddy'`)
		ew.writeln(`Since:	'v0.5.0'`)
		ew.writeln()

		ew.writeln(`Credentials:`)
		ew.writeln(`	- "GODADDY_API_KEY":	API key`)
		ew.writeln(`	- "GODADDY_API_SECRET":	API secret`)
		ew.writeln()

		ew.writeln(`Additional Configuration:`)
		ew.writeln(`	- "GODADDY_HTTP_TIMEOUT":	API request timeout`)
		ew.writeln(`	- "GODADDY_POLLING_INTERVAL":	Time between DNS propagation check`)
		ew.writeln(`	- "GODADDY_PROPAGATION_TIMEOUT":	Maximum waiting time for DNS propagation`)
		ew.writeln(`	- "GODADDY_TTL":	The TTL of the TXT record used for the DNS challenge`)

		ew.writeln()
		ew.writeln(`More information: https://go-acme.github.io/lego/dns/godaddy`)

	case "hetzner":
		// generated from: providers/dns/hetzner/hetzner.toml
		ew.writeln(`Configuration for Hetzner.`)
		ew.writeln(`Code:	'hetzner'`)
		ew.writeln(`Since:	'v3.7.0'`)
		ew.writeln()

		ew.writeln(`Credentials:`)
		ew.writeln(`	- "HETZNER_API_KEY":	API key`)
		ew.writeln()

		ew.writeln(`Additional Configuration:`)
		ew.writeln(`	- "HETZNER_HTTP_TIMEOUT":	API request timeout`)
		ew.writeln(`	- "HETZNER_POLLING_INTERVAL":	Time between DNS propagation check`)
		ew.writeln(`	- "HETZNER_PROPAGATION_TIMEOUT":	Maximum waiting time for DNS propagation`)
		ew.writeln(`	- "HETZNER_TTL":	The TTL of the TXT record used for the DNS challenge`)

		ew.writeln()
		ew.writeln(`More information: https://go-acme.github.io/lego/dns/hetzner`)

	case "hostingde":
		// generated from: providers/dns/hostingde/hostingde.toml
		ew.writeln(`Configuration for Hosting.de.`)
		ew.writeln(`Code:	'hostingde'`)
		ew.writeln(`Since:	'v1.1.0'`)
		ew.writeln()

		ew.writeln(`Credentials:`)
		ew.writeln(`	- "HOSTINGDE_API_KEY":	API key`)
		ew.writeln()

		ew.writeln(`Additional Configuration:`)
		ew.writeln(`	- "HOSTINGDE_HTTP_TIMEOUT":	API request timeout`)
		ew.writeln(`	- "HOSTINGDE_POLLING_INTERVAL":	Time between DNS propagation check`)
		ew.writeln(`	- "HOSTINGDE_PROPAGATION_TIMEOUT":	Maximum waiting time for DNS propagation`)
		ew.writeln(`	- "HOSTINGDE_TTL":	The TTL of the TXT record used for the DNS challenge`)
		ew.writeln(`	- "HOSTINGDE_ZONE_NAME":	Zone name in ACE format`)

		ew.writeln()
		ew.writeln(`More information: https://go-acme.github.io/lego/dns/hostingde`)

	case "hosttech":
		// generated from: providers/dns/hosttech/hosttech.toml
		ew.writeln(`Configuration for Hosttech.`)
		ew.writeln(`Code:	'hosttech'`)
		ew.writeln(`Since:	'v4.5.0'`)
		ew.writeln()

		ew.writeln(`Credentials:`)
		ew.writeln(`	- "HOSTTECH_API_KEY":	API login`)
		ew.writeln(`	- "HOSTTECH_PASSWORD":	API password`)
		ew.writeln()

		ew.writeln(`Additional Configuration:`)
		ew.writeln(`	- "HOSTTECH_HTTP_TIMEOUT":	API request timeout`)
		ew.writeln(`	- "HOSTTECH_POLLING_INTERVAL":	Time between DNS propagation check`)
		ew.writeln(`	- "HOSTTECH_PROPAGATION_TIMEOUT":	Maximum waiting time for DNS propagation`)
		ew.writeln(`	- "HOSTTECH_TTL":	The TTL of the TXT record used for the DNS challenge`)

		ew.writeln()
		ew.writeln(`More information: https://go-acme.github.io/lego/dns/hosttech`)

	case "httpreq":
		// generated from: providers/dns/httpreq/httpreq.toml
		ew.writeln(`Configuration for HTTP request.`)
		ew.writeln(`Code:	'httpreq'`)
		ew.writeln(`Since:	'v2.0.0'`)
		ew.writeln()

		ew.writeln(`Credentials:`)
		ew.writeln(`	- "HTTPREQ_ENDPOINT":	The URL of the server`)
		ew.writeln(`	- "HTTPREQ_MODE":	'RAW', none`)
		ew.writeln()

		ew.writeln(`Additional Configuration:`)
		ew.writeln(`	- "HTTPREQ_HTTP_TIMEOUT":	API request timeout`)
		ew.writeln(`	- "HTTPREQ_PASSWORD":	Basic authentication password`)
		ew.writeln(`	- "HTTPREQ_POLLING_INTERVAL":	Time between DNS propagation check`)
		ew.writeln(`	- "HTTPREQ_PROPAGATION_TIMEOUT":	Maximum waiting time for DNS propagation`)
		ew.writeln(`	- "HTTPREQ_USERNAME":	Basic authentication username`)

		ew.writeln()
		ew.writeln(`More information: https://go-acme.github.io/lego/dns/httpreq`)

	case "hurricane":
		// generated from: providers/dns/hurricane/hurricane.toml
		ew.writeln(`Configuration for Hurricane Electric DNS.`)
		ew.writeln(`Code:	'hurricane'`)
		ew.writeln(`Since:	'v4.3.0'`)
		ew.writeln()

		ew.writeln(`Credentials:`)
		ew.writeln(`	- "HURRICANE_TOKENS":	TXT record names and tokens`)
		ew.writeln()

		ew.writeln()
		ew.writeln(`More information: https://go-acme.github.io/lego/dns/hurricane`)

	case "hyperone":
		// generated from: providers/dns/hyperone/hyperone.toml
		ew.writeln(`Configuration for HyperOne.`)
		ew.writeln(`Code:	'hyperone'`)
		ew.writeln(`Since:	'v3.9.0'`)
		ew.writeln()

		ew.writeln(`Additional Configuration:`)
		ew.writeln(`	- "HYPERONE_API_URL":	Allows to pass custom API Endpoint to be used in the challenge (default https://api.hyperone.com/v2)`)
		ew.writeln(`	- "HYPERONE_LOCATION_ID":	Specifies location (region) to be used in API calls. (default pl-waw-1)`)
		ew.writeln(`	- "HYPERONE_PASSPORT_LOCATION":	Allows to pass custom passport file location (default ~/.h1/passport.json)`)
		ew.writeln(`	- "HYPERONE_POLLING_INTERVAL":	Time between DNS propagation check`)
		ew.writeln(`	- "HYPERONE_PROPAGATION_TIMEOUT":	Maximum waiting time for DNS propagation`)
		ew.writeln(`	- "HYPERONE_TTL":	The TTL of the TXT record used for the DNS challenge`)

		ew.writeln()
		ew.writeln(`More information: https://go-acme.github.io/lego/dns/hyperone`)

	case "ibmcloud":
		// generated from: providers/dns/ibmcloud/ibmcloud.toml
		ew.writeln(`Configuration for IBM Cloud (SoftLayer).`)
		ew.writeln(`Code:	'ibmcloud'`)
		ew.writeln(`Since:	'v4.5.0'`)
		ew.writeln()

		ew.writeln(`Credentials:`)
		ew.writeln(`	- "SOFTLAYER_API_KEY":	Classic Infrastructure API key`)
		ew.writeln(`	- "SOFTLAYER_USERNAME":	User name (IBM Cloud is <accountID>_<emailAddress>)`)
		ew.writeln()

		ew.writeln(`Additional Configuration:`)
		ew.writeln(`	- "SOFTLAYER_POLLING_INTERVAL":	Time between DNS propagation check`)
		ew.writeln(`	- "SOFTLAYER_PROPAGATION_TIMEOUT":	Maximum waiting time for DNS propagation`)
		ew.writeln(`	- "SOFTLAYER_TIMEOUT":	API request timeout`)
		ew.writeln(`	- "SOFTLAYER_TTL":	The TTL of the TXT record used for the DNS challenge`)

		ew.writeln()
		ew.writeln(`More information: https://go-acme.github.io/lego/dns/ibmcloud`)

	case "iij":
		// generated from: providers/dns/iij/iij.toml
		ew.writeln(`Configuration for Internet Initiative Japan.`)
		ew.writeln(`Code:	'iij'`)
		ew.writeln(`Since:	'v1.1.0'`)
		ew.writeln()

		ew.writeln(`Credentials:`)
		ew.writeln(`	- "IIJ_API_ACCESS_KEY":	API access key`)
		ew.writeln(`	- "IIJ_API_SECRET_KEY":	API secret key`)
		ew.writeln(`	- "IIJ_DO_SERVICE_CODE":	DO service code`)
		ew.writeln()

		ew.writeln(`Additional Configuration:`)
		ew.writeln(`	- "IIJ_POLLING_INTERVAL":	Time between DNS propagation check`)
		ew.writeln(`	- "IIJ_PROPAGATION_TIMEOUT":	Maximum waiting time for DNS propagation`)
		ew.writeln(`	- "IIJ_TTL":	The TTL of the TXT record used for the DNS challenge`)

		ew.writeln()
		ew.writeln(`More information: https://go-acme.github.io/lego/dns/iij`)

	case "iijdpf":
		// generated from: providers/dns/iijdpf/iijdpf.toml
		ew.writeln(`Configuration for IIJ DNS Platform Service.`)
		ew.writeln(`Code:	'iijdpf'`)
		ew.writeln(`Since:	'v4.7.0'`)
		ew.writeln()

		ew.writeln(`Credentials:`)
		ew.writeln(`	- "IIJ_DPF_API_TOKEN":	API token`)
		ew.writeln(`	- "IIJ_DPF_DPM_SERVICE_CODE":	IIJ Managed DNS Service's service code`)
		ew.writeln()

		ew.writeln(`Additional Configuration:`)
		ew.writeln(`	- "IIJ_DPF_API_ENDPOINT":	API endpoint URL, defaults to https://api.dns-platform.jp/dpf/v1`)
		ew.writeln(`	- "IIJ_DPF_POLLING_INTERVAL":	Time between DNS propagation check, defaults to 5 second`)
		ew.writeln(`	- "IIJ_DPF_PROPAGATION_TIMEOUT":	Maximum waiting time for DNS propagation, defaults to 660 second`)
		ew.writeln(`	- "IIJ_DPF_TTL":	The TTL of the TXT record used for the DNS challenge, default to 300`)

		ew.writeln()
		ew.writeln(`More information: https://go-acme.github.io/lego/dns/iijdpf`)

	case "infoblox":
		// generated from: providers/dns/infoblox/infoblox.toml
		ew.writeln(`Configuration for Infoblox.`)
		ew.writeln(`Code:	'infoblox'`)
		ew.writeln(`Since:	'v4.4.0'`)
		ew.writeln()

		ew.writeln(`Credentials:`)
		ew.writeln(`	- "INFOBLOX_HOST":	Host URI`)
		ew.writeln(`	- "INFOBLOX_PASSWORD":	Account Password`)
		ew.writeln(`	- "INFOBLOX_USERNAME":	Account Username`)
		ew.writeln()

		ew.writeln(`Additional Configuration:`)
		ew.writeln(`	- "INFOBLOX_DNS_VIEW":	The view for the TXT records, default: External`)
		ew.writeln(`	- "INFOBLOX_HTTP_TIMEOUT":	HTTP request timeout`)
		ew.writeln(`	- "INFOBLOX_POLLING_INTERVAL":	Time between DNS propagation check`)
		ew.writeln(`	- "INFOBLOX_PORT":	The port for the infoblox grid manager, default: 443`)
		ew.writeln(`	- "INFOBLOX_PROPAGATION_TIMEOUT":	Maximum waiting time for DNS propagation`)
		ew.writeln(`	- "INFOBLOX_SSL_VERIFY":	Whether or not to verify the TLS certificate, default: true`)
		ew.writeln(`	- "INFOBLOX_TTL":	The TTL of the TXT record used for the DNS challenge`)
		ew.writeln(`	- "INFOBLOX_WAPI_VERSION":	The version of WAPI being used, default: 2.11`)

		ew.writeln()
		ew.writeln(`More information: https://go-acme.github.io/lego/dns/infoblox`)

	case "infomaniak":
		// generated from: providers/dns/infomaniak/infomaniak.toml
		ew.writeln(`Configuration for Infomaniak.`)
		ew.writeln(`Code:	'infomaniak'`)
		ew.writeln(`Since:	'v4.1.0'`)
		ew.writeln()

		ew.writeln(`Credentials:`)
		ew.writeln(`	- "INFOMANIAK_ACCESS_TOKEN":	Access token`)
		ew.writeln()

		ew.writeln(`Additional Configuration:`)
		ew.writeln(`	- "INFOMANIAK_ENDPOINT":	https://api.infomaniak.com`)
		ew.writeln(`	- "INFOMANIAK_HTTP_TIMEOUT":	API request timeout`)
		ew.writeln(`	- "INFOMANIAK_POLLING_INTERVAL":	Time between DNS propagation check`)
		ew.writeln(`	- "INFOMANIAK_PROPAGATION_TIMEOUT":	Maximum waiting time for DNS propagation`)
		ew.writeln(`	- "INFOMANIAK_TTL":	The TTL of the TXT record used for the DNS challenge in seconds`)

		ew.writeln()
		ew.writeln(`More information: https://go-acme.github.io/lego/dns/infomaniak`)

	case "internetbs":
		// generated from: providers/dns/internetbs/internetbs.toml
		ew.writeln(`Configuration for Internet.bs.`)
		ew.writeln(`Code:	'internetbs'`)
		ew.writeln(`Since:	'v4.5.0'`)
		ew.writeln()

		ew.writeln(`Credentials:`)
		ew.writeln(`	- "INTERNET_BS_API_KEY":	API key`)
		ew.writeln(`	- "INTERNET_BS_PASSWORD":	API password`)
		ew.writeln()

		ew.writeln(`Additional Configuration:`)
		ew.writeln(`	- "INTERNET_BS_HTTP_TIMEOUT":	API request timeout`)
		ew.writeln(`	- "INTERNET_BS_POLLING_INTERVAL":	Time between DNS propagation check`)
		ew.writeln(`	- "INTERNET_BS_PROPAGATION_TIMEOUT":	Maximum waiting time for DNS propagation`)
		ew.writeln(`	- "INTERNET_BS_TTL":	The TTL of the TXT record used for the DNS challenge`)

		ew.writeln()
		ew.writeln(`More information: https://go-acme.github.io/lego/dns/internetbs`)

	case "inwx":
		// generated from: providers/dns/inwx/inwx.toml
		ew.writeln(`Configuration for INWX.`)
		ew.writeln(`Code:	'inwx'`)
		ew.writeln(`Since:	'v2.0.0'`)
		ew.writeln()

		ew.writeln(`Credentials:`)
		ew.writeln(`	- "INWX_PASSWORD":	Password`)
		ew.writeln(`	- "INWX_USERNAME":	Username`)
		ew.writeln()

		ew.writeln(`Additional Configuration:`)
		ew.writeln(`	- "INWX_POLLING_INTERVAL":	Time between DNS propagation check`)
		ew.writeln(`	- "INWX_PROPAGATION_TIMEOUT":	Maximum waiting time for DNS propagation (default 360s)`)
		ew.writeln(`	- "INWX_SANDBOX":	Activate the sandbox (boolean)`)
		ew.writeln(`	- "INWX_SHARED_SECRET":	shared secret related to 2FA`)
		ew.writeln(`	- "INWX_TTL":	The TTL of the TXT record used for the DNS challenge`)

		ew.writeln()
		ew.writeln(`More information: https://go-acme.github.io/lego/dns/inwx`)

	case "ionos":
		// generated from: providers/dns/ionos/ionos.toml
		ew.writeln(`Configuration for Ionos.`)
		ew.writeln(`Code:	'ionos'`)
		ew.writeln(`Since:	'v4.2.0'`)
		ew.writeln()

		ew.writeln(`Credentials:`)
		ew.writeln(`	- "IONOS_API_KEY":	API key '<prefix>.<secret>' https://developer.hosting.ionos.com/docs/getstarted`)
		ew.writeln()

		ew.writeln(`Additional Configuration:`)
		ew.writeln(`	- "IONOS_HTTP_TIMEOUT":	API request timeout`)
		ew.writeln(`	- "IONOS_POLLING_INTERVAL":	Time between DNS propagation check`)
		ew.writeln(`	- "IONOS_PROPAGATION_TIMEOUT":	Maximum waiting time for DNS propagation`)
		ew.writeln(`	- "IONOS_TTL":	The TTL of the TXT record used for the DNS challenge`)

		ew.writeln()
		ew.writeln(`More information: https://go-acme.github.io/lego/dns/ionos`)

	case "iwantmyname":
		// generated from: providers/dns/iwantmyname/iwantmyname.toml
		ew.writeln(`Configuration for iwantmyname.`)
		ew.writeln(`Code:	'iwantmyname'`)
		ew.writeln(`Since:	'v4.7.0'`)
		ew.writeln()

		ew.writeln(`Credentials:`)
		ew.writeln(`	- "IWANTMYNAME_PASSWORD":	API password`)
		ew.writeln(`	- "IWANTMYNAME_USERNAME":	API username`)
		ew.writeln()

		ew.writeln(`Additional Configuration:`)
		ew.writeln(`	- "IWANTMYNAME_HTTP_TIMEOUT":	API request timeout`)
		ew.writeln(`	- "IWANTMYNAME_POLLING_INTERVAL":	Time between DNS propagation check`)
		ew.writeln(`	- "IWANTMYNAME_PROPAGATION_TIMEOUT":	Maximum waiting time for DNS propagation`)
		ew.writeln(`	- "IWANTMYNAME_TTL":	The TTL of the TXT record used for the DNS challenge`)

		ew.writeln()
		ew.writeln(`More information: https://go-acme.github.io/lego/dns/iwantmyname`)

	case "joker":
		// generated from: providers/dns/joker/joker.toml
		ew.writeln(`Configuration for Joker.`)
		ew.writeln(`Code:	'joker'`)
		ew.writeln(`Since:	'v2.6.0'`)
		ew.writeln()

		ew.writeln(`Credentials:`)
		ew.writeln(`	- "JOKER_API_KEY":	API key (only with DMAPI mode)`)
		ew.writeln(`	- "JOKER_API_MODE":	'DMAPI' or 'SVC'. DMAPI is for resellers accounts. (Default: DMAPI)`)
		ew.writeln(`	- "JOKER_PASSWORD":	Joker.com password`)
		ew.writeln(`	- "JOKER_USERNAME":	Joker.com username`)
		ew.writeln()

		ew.writeln(`Additional Configuration:`)
		ew.writeln(`	- "JOKER_HTTP_TIMEOUT":	API request timeout`)
		ew.writeln(`	- "JOKER_POLLING_INTERVAL":	Time between DNS propagation check`)
		ew.writeln(`	- "JOKER_PROPAGATION_TIMEOUT":	Maximum waiting time for DNS propagation`)
		ew.writeln(`	- "JOKER_SEQUENCE_INTERVAL":	Time between sequential requests (only with 'SVC' mode)`)
		ew.writeln(`	- "JOKER_TTL":	The TTL of the TXT record used for the DNS challenge`)

		ew.writeln()
		ew.writeln(`More information: https://go-acme.github.io/lego/dns/joker`)

	case "liara":
		// generated from: providers/dns/liara/liara.toml
		ew.writeln(`Configuration for Liara.`)
		ew.writeln(`Code:	'liara'`)
		ew.writeln(`Since:	'v4.10.0'`)
		ew.writeln()

		ew.writeln(`Credentials:`)
		ew.writeln(`	- "LIARA_API_KEY":	The API key`)
		ew.writeln()

		ew.writeln(`Additional Configuration:`)
		ew.writeln(`	- "LIARA_HTTP_TIMEOUT":	API request timeout`)
		ew.writeln(`	- "LIARA_POLLING_INTERVAL":	Time between DNS propagation check`)
		ew.writeln(`	- "LIARA_PROPAGATION_TIMEOUT":	Maximum waiting time for DNS propagation`)
		ew.writeln(`	- "LIARA_TTL":	The TTL of the TXT record used for the DNS challenge`)

		ew.writeln()
		ew.writeln(`More information: https://go-acme.github.io/lego/dns/liara`)

	case "lightsail":
		// generated from: providers/dns/lightsail/lightsail.toml
		ew.writeln(`Configuration for Amazon Lightsail.`)
		ew.writeln(`Code:	'lightsail'`)
		ew.writeln(`Since:	'v0.5.0'`)
		ew.writeln()

		ew.writeln(`Credentials:`)
		ew.writeln(`	- "AWS_ACCESS_KEY_ID":	Managed by the AWS client. Access key ID ('AWS_ACCESS_KEY_ID_FILE' is not supported, use 'AWS_SHARED_CREDENTIALS_FILE' instead)`)
		ew.writeln(`	- "AWS_SECRET_ACCESS_KEY":	Managed by the AWS client. Secret access key ('AWS_SECRET_ACCESS_KEY_FILE' is not supported, use 'AWS_SHARED_CREDENTIALS_FILE' instead)`)
		ew.writeln(`	- "DNS_ZONE":	Domain name of the DNS zone`)
		ew.writeln()

		ew.writeln(`Additional Configuration:`)
		ew.writeln(`	- "AWS_SHARED_CREDENTIALS_FILE":	Managed by the AWS client. Shared credentials file.`)
		ew.writeln(`	- "LIGHTSAIL_POLLING_INTERVAL":	Time between DNS propagation check`)
		ew.writeln(`	- "LIGHTSAIL_PROPAGATION_TIMEOUT":	Maximum waiting time for DNS propagation`)

		ew.writeln()
		ew.writeln(`More information: https://go-acme.github.io/lego/dns/lightsail`)

	case "linode":
		// generated from: providers/dns/linode/linode.toml
		ew.writeln(`Configuration for Linode (v4).`)
		ew.writeln(`Code:	'linode'`)
		ew.writeln(`Since:	'v1.1.0'`)
		ew.writeln()

		ew.writeln(`Credentials:`)
		ew.writeln(`	- "LINODE_TOKEN":	API token`)
		ew.writeln()

		ew.writeln(`Additional Configuration:`)
		ew.writeln(`	- "LINODE_HTTP_TIMEOUT":	API request timeout`)
		ew.writeln(`	- "LINODE_POLLING_INTERVAL":	Time between DNS propagation check`)
		ew.writeln(`	- "LINODE_PROPAGATION_TIMEOUT":	Maximum waiting time for DNS propagation`)
		ew.writeln(`	- "LINODE_TTL":	The TTL of the TXT record used for the DNS challenge`)

		ew.writeln()
		ew.writeln(`More information: https://go-acme.github.io/lego/dns/linode`)

	case "liquidweb":
		// generated from: providers/dns/liquidweb/liquidweb.toml
		ew.writeln(`Configuration for Liquid Web.`)
		ew.writeln(`Code:	'liquidweb'`)
		ew.writeln(`Since:	'v3.1.0'`)
		ew.writeln()

		ew.writeln(`Credentials:`)
		ew.writeln(`	- "LIQUID_WEB_PASSWORD":	Storm API Password`)
		ew.writeln(`	- "LIQUID_WEB_USERNAME":	Storm API Username`)
		ew.writeln(`	- "LIQUID_WEB_ZONE":	DNS Zone`)
		ew.writeln()

		ew.writeln(`Additional Configuration:`)
		ew.writeln(`	- "LIQUID_WEB_HTTP_TIMEOUT":	Maximum waiting time for the DNS records to be created (not verified)`)
		ew.writeln(`	- "LIQUID_WEB_POLLING_INTERVAL":	Time between DNS propagation check`)
		ew.writeln(`	- "LIQUID_WEB_PROPAGATION_TIMEOUT":	Maximum waiting time for DNS propagation`)
		ew.writeln(`	- "LIQUID_WEB_TTL":	The TTL of the TXT record used for the DNS challenge`)
		ew.writeln(`	- "LIQUID_WEB_URL":	Storm API endpoint`)

		ew.writeln()
		ew.writeln(`More information: https://go-acme.github.io/lego/dns/liquidweb`)

	case "loopia":
		// generated from: providers/dns/loopia/loopia.toml
		ew.writeln(`Configuration for Loopia.`)
		ew.writeln(`Code:	'loopia'`)
		ew.writeln(`Since:	'v4.2.0'`)
		ew.writeln()

		ew.writeln(`Credentials:`)
		ew.writeln(`	- "LOOPIA_API_PASSWORD":	API password`)
		ew.writeln(`	- "LOOPIA_API_USER":	API username`)
		ew.writeln()

		ew.writeln(`Additional Configuration:`)
		ew.writeln(`	- "LOOPIA_API_URL":	API endpoint. Ex: https://api.loopia.se/RPCSERV or https://api.loopia.rs/RPCSERV`)
		ew.writeln(`	- "LOOPIA_HTTP_TIMEOUT":	API request timeout`)
		ew.writeln(`	- "LOOPIA_POLLING_INTERVAL":	Time between DNS propagation check`)
		ew.writeln(`	- "LOOPIA_PROPAGATION_TIMEOUT":	Maximum waiting time for DNS propagation`)
		ew.writeln(`	- "LOOPIA_TTL":	The TTL of the TXT record used for the DNS challenge`)

		ew.writeln()
		ew.writeln(`More information: https://go-acme.github.io/lego/dns/loopia`)

	case "luadns":
		// generated from: providers/dns/luadns/luadns.toml
		ew.writeln(`Configuration for LuaDNS.`)
		ew.writeln(`Code:	'luadns'`)
		ew.writeln(`Since:	'v3.7.0'`)
		ew.writeln()

		ew.writeln(`Credentials:`)
		ew.writeln(`	- "LUADNS_API_TOKEN":	API token`)
		ew.writeln(`	- "LUADNS_API_USERNAME":	Username (your email)`)
		ew.writeln()

		ew.writeln(`Additional Configuration:`)
		ew.writeln(`	- "LUADNS_HTTP_TIMEOUT":	API request timeout`)
		ew.writeln(`	- "LUADNS_POLLING_INTERVAL":	Time between DNS propagation check`)
		ew.writeln(`	- "LUADNS_PROPAGATION_TIMEOUT":	Maximum waiting time for DNS propagation`)
		ew.writeln(`	- "LUADNS_TTL":	The TTL of the TXT record used for the DNS challenge`)

		ew.writeln()
		ew.writeln(`More information: https://go-acme.github.io/lego/dns/luadns`)

	case "mydnsjp":
		// generated from: providers/dns/mydnsjp/mydnsjp.toml
		ew.writeln(`Configuration for MyDNS.jp.`)
		ew.writeln(`Code:	'mydnsjp'`)
		ew.writeln(`Since:	'v1.2.0'`)
		ew.writeln()

		ew.writeln(`Credentials:`)
		ew.writeln(`	- "MYDNSJP_MASTER_ID":	Master ID`)
		ew.writeln(`	- "MYDNSJP_PASSWORD":	Password`)
		ew.writeln()

		ew.writeln(`Additional Configuration:`)
		ew.writeln(`	- "MYDNSJP_HTTP_TIMEOUT":	API request timeout`)
		ew.writeln(`	- "MYDNSJP_POLLING_INTERVAL":	Time between DNS propagation check`)
		ew.writeln(`	- "MYDNSJP_PROPAGATION_TIMEOUT":	Maximum waiting time for DNS propagation`)
		ew.writeln(`	- "MYDNSJP_TTL":	The TTL of the TXT record used for the DNS challenge`)

		ew.writeln()
		ew.writeln(`More information: https://go-acme.github.io/lego/dns/mydnsjp`)

	case "mythicbeasts":
		// generated from: providers/dns/mythicbeasts/mythicbeasts.toml
		ew.writeln(`Configuration for MythicBeasts.`)
		ew.writeln(`Code:	'mythicbeasts'`)
		ew.writeln(`Since:	'v0.3.7'`)
		ew.writeln()

		ew.writeln(`Credentials:`)
		ew.writeln(`	- "MYTHICBEASTS_PASSWORD":	Password`)
		ew.writeln(`	- "MYTHICBEASTS_USERNAME":	User name`)
		ew.writeln()

		ew.writeln(`Additional Configuration:`)
		ew.writeln(`	- "MYTHICBEASTS_API_ENDPOINT":	The endpoint for the API (must implement v2)`)
		ew.writeln(`	- "MYTHICBEASTS_AUTH_API_ENDPOINT":	The endpoint for Mythic Beasts' Authentication`)
		ew.writeln(`	- "MYTHICBEASTS_HTTP_TIMEOUT":	API request timeout`)
		ew.writeln(`	- "MYTHICBEASTS_POLLING_INTERVAL":	Time between DNS propagation check`)
		ew.writeln(`	- "MYTHICBEASTS_PROPAGATION_TIMEOUT":	Maximum waiting time for DNS propagation`)
		ew.writeln(`	- "MYTHICBEASTS_TTL":	The TTL of the TXT record used for the DNS challenge`)

		ew.writeln()
		ew.writeln(`More information: https://go-acme.github.io/lego/dns/mythicbeasts`)

	case "namecheap":
		// generated from: providers/dns/namecheap/namecheap.toml
		ew.writeln(`Configuration for Namecheap.`)
		ew.writeln(`Code:	'namecheap'`)
		ew.writeln(`Since:	'v0.3.0'`)
		ew.writeln()

		ew.writeln(`Credentials:`)
		ew.writeln(`	- "NAMECHEAP_API_KEY":	API key`)
		ew.writeln(`	- "NAMECHEAP_API_USER":	API user`)
		ew.writeln()

		ew.writeln(`Additional Configuration:`)
		ew.writeln(`	- "NAMECHEAP_HTTP_TIMEOUT":	API request timeout`)
		ew.writeln(`	- "NAMECHEAP_POLLING_INTERVAL":	Time between DNS propagation check`)
		ew.writeln(`	- "NAMECHEAP_PROPAGATION_TIMEOUT":	Maximum waiting time for DNS propagation`)
		ew.writeln(`	- "NAMECHEAP_SANDBOX":	Activate the sandbox (boolean)`)
		ew.writeln(`	- "NAMECHEAP_TTL":	The TTL of the TXT record used for the DNS challenge`)

		ew.writeln()
		ew.writeln(`More information: https://go-acme.github.io/lego/dns/namecheap`)

	case "namedotcom":
		// generated from: providers/dns/namedotcom/namedotcom.toml
		ew.writeln(`Configuration for Name.com.`)
		ew.writeln(`Code:	'namedotcom'`)
		ew.writeln(`Since:	'v0.5.0'`)
		ew.writeln()

		ew.writeln(`Credentials:`)
		ew.writeln(`	- "NAMECOM_API_TOKEN":	API token`)
		ew.writeln(`	- "NAMECOM_USERNAME":	Username`)
		ew.writeln()

		ew.writeln(`Additional Configuration:`)
		ew.writeln(`	- "NAMECOM_HTTP_TIMEOUT":	API request timeout`)
		ew.writeln(`	- "NAMECOM_POLLING_INTERVAL":	Time between DNS propagation check`)
		ew.writeln(`	- "NAMECOM_PROPAGATION_TIMEOUT":	Maximum waiting time for DNS propagation`)
		ew.writeln(`	- "NAMECOM_TTL":	The TTL of the TXT record used for the DNS challenge`)

		ew.writeln()
		ew.writeln(`More information: https://go-acme.github.io/lego/dns/namedotcom`)

	case "namesilo":
		// generated from: providers/dns/namesilo/namesilo.toml
		ew.writeln(`Configuration for Namesilo.`)
		ew.writeln(`Code:	'namesilo'`)
		ew.writeln(`Since:	'v2.7.0'`)
		ew.writeln()

		ew.writeln(`Credentials:`)
		ew.writeln(`	- "NAMESILO_API_KEY":	Client ID`)
		ew.writeln()

		ew.writeln(`Additional Configuration:`)
		ew.writeln(`	- "NAMESILO_POLLING_INTERVAL":	Time between DNS propagation check`)
		ew.writeln(`	- "NAMESILO_PROPAGATION_TIMEOUT":	Maximum waiting time for DNS propagation, it is better to set larger than 15m`)
		ew.writeln(`	- "NAMESILO_TTL":	The TTL of the TXT record used for the DNS challenge, should be in [3600, 2592000]`)

		ew.writeln()
		ew.writeln(`More information: https://go-acme.github.io/lego/dns/namesilo`)

	case "nearlyfreespeech":
		// generated from: providers/dns/nearlyfreespeech/nearlyfreespeech.toml
		ew.writeln(`Configuration for NearlyFreeSpeech.NET.`)
		ew.writeln(`Code:	'nearlyfreespeech'`)
		ew.writeln(`Since:	'v4.8.0'`)
		ew.writeln()

		ew.writeln(`Credentials:`)
		ew.writeln(`	- "NEARLYFREESPEECH_API_KEY":	API Key for API requests`)
		ew.writeln(`	- "NEARLYFREESPEECH_LOGIN":	Username for API requests`)
		ew.writeln()

		ew.writeln(`Additional Configuration:`)
		ew.writeln(`	- "NEARLYFREESPEECH_HTTP_TIMEOUT":	API request timeout`)
		ew.writeln(`	- "NEARLYFREESPEECH_POLLING_INTERVAL":	Time between DNS propagation check`)
		ew.writeln(`	- "NEARLYFREESPEECH_PROPAGATION_TIMEOUT":	Maximum waiting time for DNS propagation`)
		ew.writeln(`	- "NEARLYFREESPEECH_SEQUENCE_INTERVAL":	Time between sequential requests`)
		ew.writeln(`	- "NEARLYFREESPEECH_TTL":	The TTL of the TXT record used for the DNS challenge`)

		ew.writeln()
		ew.writeln(`More information: https://go-acme.github.io/lego/dns/nearlyfreespeech`)

	case "netcup":
		// generated from: providers/dns/netcup/netcup.toml
		ew.writeln(`Configuration for Netcup.`)
		ew.writeln(`Code:	'netcup'`)
		ew.writeln(`Since:	'v1.1.0'`)
		ew.writeln()

		ew.writeln(`Credentials:`)
		ew.writeln(`	- "NETCUP_API_KEY":	API key`)
		ew.writeln(`	- "NETCUP_API_PASSWORD":	API password`)
		ew.writeln(`	- "NETCUP_CUSTOMER_NUMBER":	Customer number`)
		ew.writeln()

		ew.writeln(`Additional Configuration:`)
		ew.writeln(`	- "NETCUP_HTTP_TIMEOUT":	API request timeout`)
		ew.writeln(`	- "NETCUP_POLLING_INTERVAL":	Time between DNS propagation check`)
		ew.writeln(`	- "NETCUP_PROPAGATION_TIMEOUT":	Maximum waiting time for DNS propagation`)
		ew.writeln(`	- "NETCUP_TTL":	The TTL of the TXT record used for the DNS challenge`)

		ew.writeln()
		ew.writeln(`More information: https://go-acme.github.io/lego/dns/netcup`)

	case "netlify":
		// generated from: providers/dns/netlify/netlify.toml
		ew.writeln(`Configuration for Netlify.`)
		ew.writeln(`Code:	'netlify'`)
		ew.writeln(`Since:	'v3.7.0'`)
		ew.writeln()

		ew.writeln(`Credentials:`)
		ew.writeln(`	- "NETLIFY_TOKEN":	Token`)
		ew.writeln()

		ew.writeln(`Additional Configuration:`)
		ew.writeln(`	- "NETLIFY_HTTP_TIMEOUT":	API request timeout`)
		ew.writeln(`	- "NETLIFY_POLLING_INTERVAL":	Time between DNS propagation check`)
		ew.writeln(`	- "NETLIFY_PROPAGATION_TIMEOUT":	Maximum waiting time for DNS propagation`)
		ew.writeln(`	- "NETLIFY_TTL":	The TTL of the TXT record used for the DNS challenge`)

		ew.writeln()
		ew.writeln(`More information: https://go-acme.github.io/lego/dns/netlify`)

	case "nicmanager":
		// generated from: providers/dns/nicmanager/nicmanager.toml
		ew.writeln(`Configuration for Nicmanager.`)
		ew.writeln(`Code:	'nicmanager'`)
		ew.writeln(`Since:	'v4.5.0'`)
		ew.writeln()

		ew.writeln(`Credentials:`)
		ew.writeln(`	- "NICMANAGER_API_EMAIL":	Email-based login`)
		ew.writeln(`	- "NICMANAGER_API_LOGIN":	Login, used for Username-based login`)
		ew.writeln(`	- "NICMANAGER_API_PASSWORD":	Password, always required`)
		ew.writeln(`	- "NICMANAGER_API_USERNAME":	Username, used for Username-based login`)
		ew.writeln()

		ew.writeln(`Additional Configuration:`)
		ew.writeln(`	- "NICMANAGER_API_MODE":	mode: 'anycast' or 'zone' (default: 'anycast')`)
		ew.writeln(`	- "NICMANAGER_API_OTP":	TOTP Secret (optional)`)
		ew.writeln(`	- "NICMANAGER_HTTP_TIMEOUT":	API request timeout`)
		ew.writeln(`	- "NICMANAGER_POLLING_INTERVAL":	Time between DNS propagation check`)
		ew.writeln(`	- "NICMANAGER_PROPAGATION_TIMEOUT":	Maximum waiting time for DNS propagation`)
		ew.writeln(`	- "NICMANAGER_TTL":	The TTL of the TXT record used for the DNS challenge`)

		ew.writeln()
		ew.writeln(`More information: https://go-acme.github.io/lego/dns/nicmanager`)

	case "nifcloud":
		// generated from: providers/dns/nifcloud/nifcloud.toml
		ew.writeln(`Configuration for NIFCloud.`)
		ew.writeln(`Code:	'nifcloud'`)
		ew.writeln(`Since:	'v1.1.0'`)
		ew.writeln()

		ew.writeln(`Credentials:`)
		ew.writeln(`	- "NIFCLOUD_ACCESS_KEY_ID":	Access key`)
		ew.writeln(`	- "NIFCLOUD_SECRET_ACCESS_KEY":	Secret access key`)
		ew.writeln()

		ew.writeln(`Additional Configuration:`)
		ew.writeln(`	- "NIFCLOUD_HTTP_TIMEOUT":	API request timeout`)
		ew.writeln(`	- "NIFCLOUD_POLLING_INTERVAL":	Time between DNS propagation check`)
		ew.writeln(`	- "NIFCLOUD_PROPAGATION_TIMEOUT":	Maximum waiting time for DNS propagation`)
		ew.writeln(`	- "NIFCLOUD_TTL":	The TTL of the TXT record used for the DNS challenge`)

		ew.writeln()
		ew.writeln(`More information: https://go-acme.github.io/lego/dns/nifcloud`)

	case "njalla":
		// generated from: providers/dns/njalla/njalla.toml
		ew.writeln(`Configuration for Njalla.`)
		ew.writeln(`Code:	'njalla'`)
		ew.writeln(`Since:	'v4.3.0'`)
		ew.writeln()

		ew.writeln(`Credentials:`)
		ew.writeln(`	- "NJALLA_TOKEN":	API token`)
		ew.writeln()

		ew.writeln(`Additional Configuration:`)
		ew.writeln(`	- "NJALLA_HTTP_TIMEOUT":	API request timeout`)
		ew.writeln(`	- "NJALLA_POLLING_INTERVAL":	Time between DNS propagation check`)
		ew.writeln(`	- "NJALLA_PROPAGATION_TIMEOUT":	Maximum waiting time for DNS propagation`)
		ew.writeln(`	- "NJALLA_TTL":	The TTL of the TXT record used for the DNS challenge`)

		ew.writeln()
		ew.writeln(`More information: https://go-acme.github.io/lego/dns/njalla`)

	case "ns1":
		// generated from: providers/dns/ns1/ns1.toml
		ew.writeln(`Configuration for NS1.`)
		ew.writeln(`Code:	'ns1'`)
		ew.writeln(`Since:	'v0.4.0'`)
		ew.writeln()

		ew.writeln(`Credentials:`)
		ew.writeln(`	- "NS1_API_KEY":	API key`)
		ew.writeln()

		ew.writeln(`Additional Configuration:`)
		ew.writeln(`	- "NS1_HTTP_TIMEOUT":	API request timeout`)
		ew.writeln(`	- "NS1_POLLING_INTERVAL":	Time between DNS propagation check`)
		ew.writeln(`	- "NS1_PROPAGATION_TIMEOUT":	Maximum waiting time for DNS propagation`)
		ew.writeln(`	- "NS1_TTL":	The TTL of the TXT record used for the DNS challenge`)

		ew.writeln()
		ew.writeln(`More information: https://go-acme.github.io/lego/dns/ns1`)

	case "oraclecloud":
		// generated from: providers/dns/oraclecloud/oraclecloud.toml
		ew.writeln(`Configuration for Oracle Cloud.`)
		ew.writeln(`Code:	'oraclecloud'`)
		ew.writeln(`Since:	'v2.3.0'`)
		ew.writeln()

		ew.writeln(`Credentials:`)
		ew.writeln(`	- "OCI_COMPARTMENT_OCID":	Compartment OCID`)
		ew.writeln(`	- "OCI_PRIVKEY_FILE":	Private key file`)
		ew.writeln(`	- "OCI_PRIVKEY_PASS":	Private key password`)
		ew.writeln(`	- "OCI_PUBKEY_FINGERPRINT":	Public key fingerprint`)
		ew.writeln(`	- "OCI_REGION":	Region`)
		ew.writeln(`	- "OCI_TENANCY_OCID":	Tenancy OCID`)
		ew.writeln(`	- "OCI_USER_OCID":	User OCID`)
		ew.writeln()

		ew.writeln(`Additional Configuration:`)
		ew.writeln(`	- "OCI_POLLING_INTERVAL":	Time between DNS propagation check`)
		ew.writeln(`	- "OCI_PROPAGATION_TIMEOUT":	Maximum waiting time for DNS propagation`)
		ew.writeln(`	- "OCI_TTL":	The TTL of the TXT record used for the DNS challenge`)

		ew.writeln()
		ew.writeln(`More information: https://go-acme.github.io/lego/dns/oraclecloud`)

	case "otc":
		// generated from: providers/dns/otc/otc.toml
		ew.writeln(`Configuration for Open Telekom Cloud.`)
		ew.writeln(`Code:	'otc'`)
		ew.writeln(`Since:	'v0.4.1'`)
		ew.writeln()

		ew.writeln(`Credentials:`)
		ew.writeln(`	- "OTC_DOMAIN_NAME":	Domain name`)
		ew.writeln(`	- "OTC_IDENTITY_ENDPOINT":	Identity endpoint URL`)
		ew.writeln(`	- "OTC_PASSWORD":	Password`)
		ew.writeln(`	- "OTC_PROJECT_NAME":	Project name`)
		ew.writeln(`	- "OTC_USER_NAME":	User name`)
		ew.writeln()

		ew.writeln(`Additional Configuration:`)
		ew.writeln(`	- "OTC_HTTP_TIMEOUT":	API request timeout`)
		ew.writeln(`	- "OTC_POLLING_INTERVAL":	Time between DNS propagation check`)
		ew.writeln(`	- "OTC_PROPAGATION_TIMEOUT":	Maximum waiting time for DNS propagation`)
		ew.writeln(`	- "OTC_TTL":	The TTL of the TXT record used for the DNS challenge`)

		ew.writeln()
		ew.writeln(`More information: https://go-acme.github.io/lego/dns/otc`)

	case "ovh":
		// generated from: providers/dns/ovh/ovh.toml
		ew.writeln(`Configuration for OVH.`)
		ew.writeln(`Code:	'ovh'`)
		ew.writeln(`Since:	'v0.4.0'`)
		ew.writeln()

		ew.writeln(`Credentials:`)
		ew.writeln(`	- "OVH_APPLICATION_KEY":	Application key`)
		ew.writeln(`	- "OVH_APPLICATION_SECRET":	Application secret`)
		ew.writeln(`	- "OVH_CONSUMER_KEY":	Consumer key`)
		ew.writeln(`	- "OVH_ENDPOINT":	Endpoint URL (ovh-eu or ovh-ca)`)
		ew.writeln()

		ew.writeln(`Additional Configuration:`)
		ew.writeln(`	- "OVH_HTTP_TIMEOUT":	API request timeout`)
		ew.writeln(`	- "OVH_POLLING_INTERVAL":	Time between DNS propagation check`)
		ew.writeln(`	- "OVH_PROPAGATION_TIMEOUT":	Maximum waiting time for DNS propagation`)
		ew.writeln(`	- "OVH_TTL":	The TTL of the TXT record used for the DNS challenge`)

		ew.writeln()
		ew.writeln(`More information: https://go-acme.github.io/lego/dns/ovh`)

	case "pdns":
		// generated from: providers/dns/pdns/pdns.toml
		ew.writeln(`Configuration for PowerDNS.`)
		ew.writeln(`Code:	'pdns'`)
		ew.writeln(`Since:	'v0.4.0'`)
		ew.writeln()

		ew.writeln(`Credentials:`)
		ew.writeln(`	- "PDNS_API_KEY":	API key`)
		ew.writeln(`	- "PDNS_API_URL":	API URL`)
		ew.writeln()

		ew.writeln(`Additional Configuration:`)
		ew.writeln(`	- "PDNS_HTTP_TIMEOUT":	API request timeout`)
		ew.writeln(`	- "PDNS_POLLING_INTERVAL":	Time between DNS propagation check`)
		ew.writeln(`	- "PDNS_PROPAGATION_TIMEOUT":	Maximum waiting time for DNS propagation`)
		ew.writeln(`	- "PDNS_SERVER_NAME":	Name of the server in the URL, 'localhost' by default`)
		ew.writeln(`	- "PDNS_TTL":	The TTL of the TXT record used for the DNS challenge`)

		ew.writeln()
		ew.writeln(`More information: https://go-acme.github.io/lego/dns/pdns`)

	case "porkbun":
		// generated from: providers/dns/porkbun/porkbun.toml
		ew.writeln(`Configuration for Porkbun.`)
		ew.writeln(`Code:	'porkbun'`)
		ew.writeln(`Since:	'v4.4.0'`)
		ew.writeln()

		ew.writeln(`Credentials:`)
		ew.writeln(`	- "PORKBUN_API_KEY":	API key`)
		ew.writeln(`	- "PORKBUN_SECRET_API_KEY":	secret API key`)
		ew.writeln()

		ew.writeln(`Additional Configuration:`)
		ew.writeln(`	- "PORKBUN_HTTP_TIMEOUT":	API request timeout`)
		ew.writeln(`	- "PORKBUN_POLLING_INTERVAL":	Time between DNS propagation check`)
		ew.writeln(`	- "PORKBUN_PROPAGATION_TIMEOUT":	Maximum waiting time for DNS propagation`)
		ew.writeln(`	- "PORKBUN_TTL":	The TTL of the TXT record used for the DNS challenge`)

		ew.writeln()
		ew.writeln(`More information: https://go-acme.github.io/lego/dns/porkbun`)

	case "rackspace":
		// generated from: providers/dns/rackspace/rackspace.toml
		ew.writeln(`Configuration for Rackspace.`)
		ew.writeln(`Code:	'rackspace'`)
		ew.writeln(`Since:	'v0.4.0'`)
		ew.writeln()

		ew.writeln(`Credentials:`)
		ew.writeln(`	- "RACKSPACE_API_KEY":	API key`)
		ew.writeln(`	- "RACKSPACE_USER":	API user`)
		ew.writeln()

		ew.writeln(`Additional Configuration:`)
		ew.writeln(`	- "RACKSPACE_HTTP_TIMEOUT":	API request timeout`)
		ew.writeln(`	- "RACKSPACE_POLLING_INTERVAL":	Time between DNS propagation check`)
		ew.writeln(`	- "RACKSPACE_PROPAGATION_TIMEOUT":	Maximum waiting time for DNS propagation`)
		ew.writeln(`	- "RACKSPACE_TTL":	The TTL of the TXT record used for the DNS challenge`)

		ew.writeln()
		ew.writeln(`More information: https://go-acme.github.io/lego/dns/rackspace`)

	case "regru":
		// generated from: providers/dns/regru/regru.toml
		ew.writeln(`Configuration for reg.ru.`)
		ew.writeln(`Code:	'regru'`)
		ew.writeln(`Since:	'v3.5.0'`)
		ew.writeln()

		ew.writeln(`Credentials:`)
		ew.writeln(`	- "REGRU_PASSWORD":	API password`)
		ew.writeln(`	- "REGRU_USERNAME":	API username`)
		ew.writeln()

		ew.writeln(`Additional Configuration:`)
		ew.writeln(`	- "REGRU_HTTP_TIMEOUT":	API request timeout`)
		ew.writeln(`	- "REGRU_POLLING_INTERVAL":	Time between DNS propagation check`)
		ew.writeln(`	- "REGRU_PROPAGATION_TIMEOUT":	Maximum waiting time for DNS propagation`)
		ew.writeln(`	- "REGRU_TTL":	The TTL of the TXT record used for the DNS challenge`)

		ew.writeln()
		ew.writeln(`More information: https://go-acme.github.io/lego/dns/regru`)

	case "rfc2136":
		// generated from: providers/dns/rfc2136/rfc2136.toml
		ew.writeln(`Configuration for RFC2136.`)
		ew.writeln(`Code:	'rfc2136'`)
		ew.writeln(`Since:	'v0.3.0'`)
		ew.writeln()

		ew.writeln(`Credentials:`)
		ew.writeln(`	- "RFC2136_NAMESERVER":	Network address in the form "host" or "host:port"`)
		ew.writeln(`	- "RFC2136_TSIG_ALGORITHM":	TSIG algorithm. See [miekg/dns#tsig.go](https://github.com/miekg/dns/blob/master/tsig.go) for supported values. To disable TSIG authentication, leave the 'RFC2136_TSIG*' variables unset.`)
		ew.writeln(`	- "RFC2136_TSIG_KEY":	Name of the secret key as defined in DNS server configuration. To disable TSIG authentication, leave the 'RFC2136_TSIG*' variables unset.`)
		ew.writeln(`	- "RFC2136_TSIG_SECRET":	Secret key payload. To disable TSIG authentication, leave the' RFC2136_TSIG*' variables unset.`)
		ew.writeln()

		ew.writeln(`Additional Configuration:`)
		ew.writeln(`	- "RFC2136_DNS_TIMEOUT":	API request timeout`)
		ew.writeln(`	- "RFC2136_POLLING_INTERVAL":	Time between DNS propagation check`)
		ew.writeln(`	- "RFC2136_PROPAGATION_TIMEOUT":	Maximum waiting time for DNS propagation`)
		ew.writeln(`	- "RFC2136_SEQUENCE_INTERVAL":	Time between sequential requests`)
		ew.writeln(`	- "RFC2136_TTL":	The TTL of the TXT record used for the DNS challenge`)

		ew.writeln()
		ew.writeln(`More information: https://go-acme.github.io/lego/dns/rfc2136`)

	case "rimuhosting":
		// generated from: providers/dns/rimuhosting/rimuhosting.toml
		ew.writeln(`Configuration for RimuHosting.`)
		ew.writeln(`Code:	'rimuhosting'`)
		ew.writeln(`Since:	'v0.3.5'`)
		ew.writeln()

		ew.writeln(`Credentials:`)
		ew.writeln(`	- "RIMUHOSTING_API_KEY":	User API key`)
		ew.writeln()

		ew.writeln(`Additional Configuration:`)
		ew.writeln(`	- "RIMUHOSTING_HTTP_TIMEOUT":	API request timeout`)
		ew.writeln(`	- "RIMUHOSTING_POLLING_INTERVAL":	Time between DNS propagation check`)
		ew.writeln(`	- "RIMUHOSTING_PROPAGATION_TIMEOUT":	Maximum waiting time for DNS propagation`)
		ew.writeln(`	- "RIMUHOSTING_TTL":	The TTL of the TXT record used for the DNS challenge`)

		ew.writeln()
		ew.writeln(`More information: https://go-acme.github.io/lego/dns/rimuhosting`)

	case "route53":
		// generated from: providers/dns/route53/route53.toml
		ew.writeln(`Configuration for Amazon Route 53.`)
		ew.writeln(`Code:	'route53'`)
		ew.writeln(`Since:	'v0.3.0'`)
		ew.writeln()

		ew.writeln(`Credentials:`)
		ew.writeln(`	- "AWS_ACCESS_KEY_ID":	Managed by the AWS client. Access key ID ('AWS_ACCESS_KEY_ID_FILE' is not supported, use 'AWS_SHARED_CREDENTIALS_FILE' instead)`)
		ew.writeln(`	- "AWS_ASSUME_ROLE_ARN":	Managed by the AWS Role ARN ('AWS_ASSUME_ROLE_ARN' is not supported)`)
		ew.writeln(`	- "AWS_HOSTED_ZONE_ID":	Override the hosted zone ID.`)
		ew.writeln(`	- "AWS_PROFILE":	Managed by the AWS client ('AWS_PROFILE_FILE' is not supported)`)
		ew.writeln(`	- "AWS_REGION":	Managed by the AWS client ('AWS_REGION_FILE' is not supported)`)
		ew.writeln(`	- "AWS_SDK_LOAD_CONFIG":	Managed by the AWS client. Retrieve the region from the CLI config file ('AWS_SDK_LOAD_CONFIG_FILE' is not supported)`)
		ew.writeln(`	- "AWS_SECRET_ACCESS_KEY":	Managed by the AWS client. Secret access key ('AWS_SECRET_ACCESS_KEY_FILE' is not supported, use 'AWS_SHARED_CREDENTIALS_FILE' instead)`)
		ew.writeln()

		ew.writeln(`Additional Configuration:`)
		ew.writeln(`	- "AWS_MAX_RETRIES":	The number of maximum returns the service will use to make an individual API request`)
		ew.writeln(`	- "AWS_POLLING_INTERVAL":	Time between DNS propagation check`)
		ew.writeln(`	- "AWS_PROPAGATION_TIMEOUT":	Maximum waiting time for DNS propagation`)
		ew.writeln(`	- "AWS_SHARED_CREDENTIALS_FILE":	Managed by the AWS client. Shared credentials file.`)
		ew.writeln(`	- "AWS_TTL":	The TTL of the TXT record used for the DNS challenge`)

		ew.writeln()
		ew.writeln(`More information: https://go-acme.github.io/lego/dns/route53`)

	case "safedns":
		// generated from: providers/dns/safedns/safedns.toml
		ew.writeln(`Configuration for UKFast SafeDNS.`)
		ew.writeln(`Code:	'safedns'`)
		ew.writeln(`Since:	'v4.6.0'`)
		ew.writeln()

		ew.writeln(`Credentials:`)
		ew.writeln(`	- "SAFEDNS_AUTH_TOKEN":	Authentication token`)
		ew.writeln()

		ew.writeln(`Additional Configuration:`)
		ew.writeln(`	- "SAFEDNS_HTTP_TIMEOUT":	API request timeout`)
		ew.writeln(`	- "SAFEDNS_POLLING_INTERVAL":	Time between DNS propagation check`)
		ew.writeln(`	- "SAFEDNS_PROPAGATION_TIMEOUT":	Maximum waiting time for DNS propagation`)
		ew.writeln(`	- "SAFEDNS_TTL":	The TTL of the TXT record used for the DNS challenge`)

		ew.writeln()
		ew.writeln(`More information: https://go-acme.github.io/lego/dns/safedns`)

	case "sakuracloud":
		// generated from: providers/dns/sakuracloud/sakuracloud.toml
		ew.writeln(`Configuration for Sakura Cloud.`)
		ew.writeln(`Code:	'sakuracloud'`)
		ew.writeln(`Since:	'v1.1.0'`)
		ew.writeln()

		ew.writeln(`Credentials:`)
		ew.writeln(`	- "SAKURACLOUD_ACCESS_TOKEN":	Access token`)
		ew.writeln(`	- "SAKURACLOUD_ACCESS_TOKEN_SECRET":	Access token secret`)
		ew.writeln()

		ew.writeln(`Additional Configuration:`)
		ew.writeln(`	- "SAKURACLOUD_HTTP_TIMEOUT":	API request timeout`)
		ew.writeln(`	- "SAKURACLOUD_POLLING_INTERVAL":	Time between DNS propagation check`)
		ew.writeln(`	- "SAKURACLOUD_PROPAGATION_TIMEOUT":	Maximum waiting time for DNS propagation`)
		ew.writeln(`	- "SAKURACLOUD_TTL":	The TTL of the TXT record used for the DNS challenge`)

		ew.writeln()
		ew.writeln(`More information: https://go-acme.github.io/lego/dns/sakuracloud`)

	case "scaleway":
		// generated from: providers/dns/scaleway/scaleway.toml
		ew.writeln(`Configuration for Scaleway.`)
		ew.writeln(`Code:	'scaleway'`)
		ew.writeln(`Since:	'v3.4.0'`)
		ew.writeln()

		ew.writeln(`Credentials:`)
		ew.writeln(`	- "SCALEWAY_API_TOKEN":	API token`)
		ew.writeln(`	- "SCALEWAY_PROJECT_ID":	Project to use (optional)`)
		ew.writeln()

		ew.writeln(`Additional Configuration:`)
		ew.writeln(`	- "SCALEWAY_POLLING_INTERVAL":	Time between DNS propagation check`)
		ew.writeln(`	- "SCALEWAY_PROPAGATION_TIMEOUT":	Maximum waiting time for DNS propagation`)
		ew.writeln(`	- "SCALEWAY_TTL":	The TTL of the TXT record used for the DNS challenge`)

		ew.writeln()
		ew.writeln(`More information: https://go-acme.github.io/lego/dns/scaleway`)

	case "selectel":
		// generated from: providers/dns/selectel/selectel.toml
		ew.writeln(`Configuration for Selectel.`)
		ew.writeln(`Code:	'selectel'`)
		ew.writeln(`Since:	'v1.2.0'`)
		ew.writeln()

		ew.writeln(`Credentials:`)
		ew.writeln(`	- "SELECTEL_API_TOKEN":	API token`)
		ew.writeln()

		ew.writeln(`Additional Configuration:`)
		ew.writeln(`	- "SELECTEL_BASE_URL":	API endpoint URL`)
		ew.writeln(`	- "SELECTEL_HTTP_TIMEOUT":	API request timeout`)
		ew.writeln(`	- "SELECTEL_POLLING_INTERVAL":	Time between DNS propagation check`)
		ew.writeln(`	- "SELECTEL_PROPAGATION_TIMEOUT":	Maximum waiting time for DNS propagation`)
		ew.writeln(`	- "SELECTEL_TTL":	The TTL of the TXT record used for the DNS challenge`)

		ew.writeln()
		ew.writeln(`More information: https://go-acme.github.io/lego/dns/selectel`)

	case "servercow":
		// generated from: providers/dns/servercow/servercow.toml
		ew.writeln(`Configuration for Servercow.`)
		ew.writeln(`Code:	'servercow'`)
		ew.writeln(`Since:	'v3.4.0'`)
		ew.writeln()

		ew.writeln(`Credentials:`)
		ew.writeln(`	- "SERVERCOW_PASSWORD":	API password`)
		ew.writeln(`	- "SERVERCOW_USERNAME":	API username`)
		ew.writeln()

		ew.writeln(`Additional Configuration:`)
		ew.writeln(`	- "SERVERCOW_HTTP_TIMEOUT":	API request timeout`)
		ew.writeln(`	- "SERVERCOW_POLLING_INTERVAL":	Time between DNS propagation check`)
		ew.writeln(`	- "SERVERCOW_PROPAGATION_TIMEOUT":	Maximum waiting time for DNS propagation`)
		ew.writeln(`	- "SERVERCOW_TTL":	The TTL of the TXT record used for the DNS challenge`)

		ew.writeln()
		ew.writeln(`More information: https://go-acme.github.io/lego/dns/servercow`)

	case "simply":
		// generated from: providers/dns/simply/simply.toml
		ew.writeln(`Configuration for Simply.com.`)
		ew.writeln(`Code:	'simply'`)
		ew.writeln(`Since:	'v4.4.0'`)
		ew.writeln()

		ew.writeln(`Credentials:`)
		ew.writeln(`	- "SIMPLY_ACCOUNT_NAME":	Account name`)
		ew.writeln(`	- "SIMPLY_API_KEY":	API key`)
		ew.writeln()

		ew.writeln(`Additional Configuration:`)
		ew.writeln(`	- "SIMPLY_HTTP_TIMEOUT":	API request timeout`)
		ew.writeln(`	- "SIMPLY_POLLING_INTERVAL":	Time between DNS propagation check`)
		ew.writeln(`	- "SIMPLY_PROPAGATION_TIMEOUT":	Maximum waiting time for DNS propagation`)
		ew.writeln(`	- "SIMPLY_TTL":	The TTL of the TXT record used for the DNS challenge`)

		ew.writeln()
		ew.writeln(`More information: https://go-acme.github.io/lego/dns/simply`)

	case "sonic":
		// generated from: providers/dns/sonic/sonic.toml
		ew.writeln(`Configuration for Sonic.`)
		ew.writeln(`Code:	'sonic'`)
		ew.writeln(`Since:	'v4.4.0'`)
		ew.writeln()

		ew.writeln(`Credentials:`)
		ew.writeln(`	- "SONIC_API_KEY":	API Key`)
		ew.writeln(`	- "SONIC_USER_ID":	User ID`)
		ew.writeln()

		ew.writeln(`Additional Configuration:`)
		ew.writeln(`	- "SONIC_HTTP_TIMEOUT":	API request timeout`)
		ew.writeln(`	- "SONIC_POLLING_INTERVAL":	Time between DNS propagation check`)
		ew.writeln(`	- "SONIC_PROPAGATION_TIMEOUT":	Maximum waiting time for DNS propagation`)
		ew.writeln(`	- "SONIC_SEQUENCE_INTERVAL":	Time between sequential requests`)
		ew.writeln(`	- "SONIC_TTL":	The TTL of the TXT record used for the DNS challenge`)

		ew.writeln()
		ew.writeln(`More information: https://go-acme.github.io/lego/dns/sonic`)

	case "stackpath":
		// generated from: providers/dns/stackpath/stackpath.toml
		ew.writeln(`Configuration for Stackpath.`)
		ew.writeln(`Code:	'stackpath'`)
		ew.writeln(`Since:	'v1.1.0'`)
		ew.writeln()

		ew.writeln(`Credentials:`)
		ew.writeln(`	- "STACKPATH_CLIENT_ID":	Client ID`)
		ew.writeln(`	- "STACKPATH_CLIENT_SECRET":	Client secret`)
		ew.writeln(`	- "STACKPATH_STACK_ID":	Stack ID`)
		ew.writeln()

		ew.writeln(`Additional Configuration:`)
		ew.writeln(`	- "STACKPATH_POLLING_INTERVAL":	Time between DNS propagation check`)
		ew.writeln(`	- "STACKPATH_PROPAGATION_TIMEOUT":	Maximum waiting time for DNS propagation`)
		ew.writeln(`	- "STACKPATH_TTL":	The TTL of the TXT record used for the DNS challenge`)

		ew.writeln()
		ew.writeln(`More information: https://go-acme.github.io/lego/dns/stackpath`)

	case "tencentcloud":
		// generated from: providers/dns/tencentcloud/tencentcloud.toml
		ew.writeln(`Configuration for Tencent Cloud DNS.`)
		ew.writeln(`Code:	'tencentcloud'`)
		ew.writeln(`Since:	'v4.6.0'`)
		ew.writeln()

		ew.writeln(`Credentials:`)
		ew.writeln(`	- "TENCENTCLOUD_SECRET_ID":	Access key ID`)
		ew.writeln(`	- "TENCENTCLOUD_SECRET_KEY":	Access Key secret`)
		ew.writeln()

		ew.writeln(`Additional Configuration:`)
		ew.writeln(`	- "TENCENTCLOUD_HTTP_TIMEOUT":	API request timeout`)
		ew.writeln(`	- "TENCENTCLOUD_POLLING_INTERVAL":	Time between DNS propagation check`)
		ew.writeln(`	- "TENCENTCLOUD_PROPAGATION_TIMEOUT":	Maximum waiting time for DNS propagation`)
		ew.writeln(`	- "TENCENTCLOUD_REGION":	Region`)
		ew.writeln(`	- "TENCENTCLOUD_SESSION_TOKEN":	Access Key token`)
		ew.writeln(`	- "TENCENTCLOUD_TTL":	The TTL of the TXT record used for the DNS challenge`)

		ew.writeln()
		ew.writeln(`More information: https://go-acme.github.io/lego/dns/tencentcloud`)

	case "transip":
		// generated from: providers/dns/transip/transip.toml
		ew.writeln(`Configuration for TransIP.`)
		ew.writeln(`Code:	'transip'`)
		ew.writeln(`Since:	'v2.0.0'`)
		ew.writeln()

		ew.writeln(`Credentials:`)
		ew.writeln(`	- "TRANSIP_ACCOUNT_NAME":	Account name`)
		ew.writeln(`	- "TRANSIP_PRIVATE_KEY_PATH":	Private key path`)
		ew.writeln()

		ew.writeln(`Additional Configuration:`)
		ew.writeln(`	- "TRANSIP_POLLING_INTERVAL":	Time between DNS propagation check`)
		ew.writeln(`	- "TRANSIP_PROPAGATION_TIMEOUT":	Maximum waiting time for DNS propagation`)
		ew.writeln(`	- "TRANSIP_TTL":	The TTL of the TXT record used for the DNS challenge`)

		ew.writeln()
		ew.writeln(`More information: https://go-acme.github.io/lego/dns/transip`)

	case "ultradns":
		// generated from: providers/dns/ultradns/ultradns.toml
		ew.writeln(`Configuration for Ultradns.`)
		ew.writeln(`Code:	'ultradns'`)
		ew.writeln(`Since:	'v4.10.0'`)
		ew.writeln()

		ew.writeln(`Credentials:`)
		ew.writeln(`	- "ULTRADNS_PASSWORD":	API Password`)
		ew.writeln(`	- "ULTRADNS_USERNAME":	API Username`)
		ew.writeln()

		ew.writeln(`Additional Configuration:`)
		ew.writeln(`	- "ULTRADNS_ENDPOINT":	API endpoint URL, defaults to https://api.ultradns.com/`)
		ew.writeln(`	- "ULTRADNS_POLLING_INTERVAL":	Time between DNS propagation check`)
		ew.writeln(`	- "ULTRADNS_PROPAGATION_TIMEOUT":	Maximum waiting time for DNS propagation`)
		ew.writeln(`	- "ULTRADNS_TTL":	The TTL of the TXT record used for the DNS challenge`)

		ew.writeln()
		ew.writeln(`More information: https://go-acme.github.io/lego/dns/ultradns`)

	case "variomedia":
		// generated from: providers/dns/variomedia/variomedia.toml
		ew.writeln(`Configuration for Variomedia.`)
		ew.writeln(`Code:	'variomedia'`)
		ew.writeln(`Since:	'v4.8.0'`)
		ew.writeln()

		ew.writeln(`Credentials:`)
		ew.writeln(`	- "VARIOMEDIA_API_TOKEN":	API token`)
		ew.writeln()

		ew.writeln(`Additional Configuration:`)
		ew.writeln(`	- "DODE_SEQUENCE_INTERVAL":	Time between sequential requests`)
		ew.writeln(`	- "VARIOMEDIA_HTTP_TIMEOUT":	API request timeout`)
		ew.writeln(`	- "VARIOMEDIA_POLLING_INTERVAL":	Time between DNS propagation check`)
		ew.writeln(`	- "VARIOMEDIA_PROPAGATION_TIMEOUT":	Maximum waiting time for DNS propagation`)
		ew.writeln(`	- "VARIOMEDIA_TTL":	The TTL of the TXT record used for the DNS challenge`)

		ew.writeln()
		ew.writeln(`More information: https://go-acme.github.io/lego/dns/variomedia`)

	case "vegadns":
		// generated from: providers/dns/vegadns/vegadns.toml
		ew.writeln(`Configuration for VegaDNS.`)
		ew.writeln(`Code:	'vegadns'`)
		ew.writeln(`Since:	'v1.1.0'`)
		ew.writeln()

		ew.writeln(`Credentials:`)
		ew.writeln(`	- "SECRET_VEGADNS_KEY":	API key`)
		ew.writeln(`	- "SECRET_VEGADNS_SECRET":	API secret`)
		ew.writeln(`	- "VEGADNS_URL":	API endpoint URL`)
		ew.writeln()

		ew.writeln(`Additional Configuration:`)
		ew.writeln(`	- "VEGADNS_POLLING_INTERVAL":	Time between DNS propagation check`)
		ew.writeln(`	- "VEGADNS_PROPAGATION_TIMEOUT":	Maximum waiting time for DNS propagation`)
		ew.writeln(`	- "VEGADNS_TTL":	The TTL of the TXT record used for the DNS challenge`)

		ew.writeln()
		ew.writeln(`More information: https://go-acme.github.io/lego/dns/vegadns`)

	case "vercel":
		// generated from: providers/dns/vercel/vercel.toml
		ew.writeln(`Configuration for Vercel.`)
		ew.writeln(`Code:	'vercel'`)
		ew.writeln(`Since:	'v4.7.0'`)
		ew.writeln()

		ew.writeln(`Credentials:`)
		ew.writeln(`	- "VERCEL_API_TOKEN":	Authentication token`)
		ew.writeln()

		ew.writeln(`Additional Configuration:`)
		ew.writeln(`	- "VERCEL_HTTP_TIMEOUT":	API request timeout`)
		ew.writeln(`	- "VERCEL_POLLING_INTERVAL":	Time between DNS propagation check`)
		ew.writeln(`	- "VERCEL_PROPAGATION_TIMEOUT":	Maximum waiting time for DNS propagation`)
		ew.writeln(`	- "VERCEL_TEAM_ID":	Team ID (ex: team_xxxxxxxxxxxxxxxxxxxxxxxx)`)
		ew.writeln(`	- "VERCEL_TTL":	The TTL of the TXT record used for the DNS challenge`)

		ew.writeln()
		ew.writeln(`More information: https://go-acme.github.io/lego/dns/vercel`)

	case "versio":
		// generated from: providers/dns/versio/versio.toml
		ew.writeln(`Configuration for Versio.[nl|eu|uk].`)
		ew.writeln(`Code:	'versio'`)
		ew.writeln(`Since:	'v2.7.0'`)
		ew.writeln()

		ew.writeln(`Credentials:`)
		ew.writeln(`	- "VERSIO_PASSWORD":	Basic authentication password`)
		ew.writeln(`	- "VERSIO_USERNAME":	Basic authentication username`)
		ew.writeln()

		ew.writeln(`Additional Configuration:`)
		ew.writeln(`	- "VERSIO_ENDPOINT":	The endpoint URL of the API Server`)
		ew.writeln(`	- "VERSIO_HTTP_TIMEOUT":	API request timeout`)
		ew.writeln(`	- "VERSIO_POLLING_INTERVAL":	Time between DNS propagation check`)
		ew.writeln(`	- "VERSIO_PROPAGATION_TIMEOUT":	Maximum waiting time for DNS propagation`)
		ew.writeln(`	- "VERSIO_SEQUENCE_INTERVAL":	Time between sequential requests, default 60s`)
		ew.writeln(`	- "VERSIO_TTL":	The TTL of the TXT record used for the DNS challenge`)

		ew.writeln()
		ew.writeln(`More information: https://go-acme.github.io/lego/dns/versio`)

	case "vinyldns":
		// generated from: providers/dns/vinyldns/vinyldns.toml
		ew.writeln(`Configuration for VinylDNS.`)
		ew.writeln(`Code:	'vinyldns'`)
		ew.writeln(`Since:	'v4.4.0'`)
		ew.writeln()

		ew.writeln(`Credentials:`)
		ew.writeln(`	- "VINYLDNS_ACCESS_KEY":	The VinylDNS API key`)
		ew.writeln(`	- "VINYLDNS_HOST":	The VinylDNS API URL`)
		ew.writeln(`	- "VINYLDNS_SECRET_KEY":	The VinylDNS API Secret key`)
		ew.writeln()

		ew.writeln(`Additional Configuration:`)
		ew.writeln(`	- "VINYLDNS_POLLING_INTERVAL":	Time between DNS propagation check`)
		ew.writeln(`	- "VINYLDNS_PROPAGATION_TIMEOUT":	Maximum waiting time for DNS propagation`)
		ew.writeln(`	- "VINYLDNS_TTL":	The TTL of the TXT record used for the DNS challenge`)

		ew.writeln()
		ew.writeln(`More information: https://go-acme.github.io/lego/dns/vinyldns`)

	case "vkcloud":
		// generated from: providers/dns/vkcloud/vkcloud.toml
		ew.writeln(`Configuration for VK Cloud.`)
		ew.writeln(`Code:	'vkcloud'`)
		ew.writeln(`Since:	'v4.9.0'`)
		ew.writeln()

		ew.writeln(`Credentials:`)
		ew.writeln(`	- "VK_CLOUD_PASSWORD":	Password for VK Cloud account`)
		ew.writeln(`	- "VK_CLOUD_PROJECT_ID":	String ID of project in VK Cloud`)
		ew.writeln(`	- "VK_CLOUD_USERNAME":	Email of VK Cloud account`)
		ew.writeln()

		ew.writeln(`Additional Configuration:`)
		ew.writeln(`	- "VK_CLOUD_DNS_ENDPOINT":	URL of DNS API. Defaults to https://mcs.mail.ru/public-dns but can be changed for usage with private clouds`)
		ew.writeln(`	- "VK_CLOUD_DOMAIN_NAME":	Openstack users domain name. Defaults to 'users' but can be changed for usage with private clouds`)
		ew.writeln(`	- "VK_CLOUD_IDENTITY_ENDPOINT":	URL of OpenStack Auth API, Defaults to https://infra.mail.ru:35357/v3/ but can be changed for usage with private clouds`)
		ew.writeln(`	- "VK_CLOUD_POLLING_INTERVAL":	Time between DNS propagation check`)
		ew.writeln(`	- "VK_CLOUD_PROPAGATION_TIMEOUT":	Maximum waiting time for DNS propagation`)
		ew.writeln(`	- "VK_CLOUD_TTL":	The TTL of the TXT record used for the DNS challenge`)

		ew.writeln()
		ew.writeln(`More information: https://go-acme.github.io/lego/dns/vkcloud`)

	case "vscale":
		// generated from: providers/dns/vscale/vscale.toml
		ew.writeln(`Configuration for Vscale.`)
		ew.writeln(`Code:	'vscale'`)
		ew.writeln(`Since:	'v2.0.0'`)
		ew.writeln()

		ew.writeln(`Credentials:`)
		ew.writeln(`	- "VSCALE_API_TOKEN":	API token`)
		ew.writeln()

		ew.writeln(`Additional Configuration:`)
		ew.writeln(`	- "VSCALE_BASE_URL":	API endpoint URL`)
		ew.writeln(`	- "VSCALE_HTTP_TIMEOUT":	API request timeout`)
		ew.writeln(`	- "VSCALE_POLLING_INTERVAL":	Time between DNS propagation check`)
		ew.writeln(`	- "VSCALE_PROPAGATION_TIMEOUT":	Maximum waiting time for DNS propagation`)
		ew.writeln(`	- "VSCALE_TTL":	The TTL of the TXT record used for the DNS challenge`)

		ew.writeln()
		ew.writeln(`More information: https://go-acme.github.io/lego/dns/vscale`)

	case "vultr":
		// generated from: providers/dns/vultr/vultr.toml
		ew.writeln(`Configuration for Vultr.`)
		ew.writeln(`Code:	'vultr'`)
		ew.writeln(`Since:	'v0.3.1'`)
		ew.writeln()

		ew.writeln(`Credentials:`)
		ew.writeln(`	- "VULTR_API_KEY":	API key`)
		ew.writeln()

		ew.writeln(`Additional Configuration:`)
		ew.writeln(`	- "VULTR_HTTP_TIMEOUT":	API request timeout`)
		ew.writeln(`	- "VULTR_POLLING_INTERVAL":	Time between DNS propagation check`)
		ew.writeln(`	- "VULTR_PROPAGATION_TIMEOUT":	Maximum waiting time for DNS propagation`)
		ew.writeln(`	- "VULTR_TTL":	The TTL of the TXT record used for the DNS challenge`)

		ew.writeln()
		ew.writeln(`More information: https://go-acme.github.io/lego/dns/vultr`)

	case "wedos":
		// generated from: providers/dns/wedos/wedos.toml
		ew.writeln(`Configuration for WEDOS.`)
		ew.writeln(`Code:	'wedos'`)
		ew.writeln(`Since:	'v4.4.0'`)
		ew.writeln()

		ew.writeln(`Credentials:`)
		ew.writeln(`	- "WEDOS_USERNAME":	Username is the same as for the admin account`)
		ew.writeln(`	- "WEDOS_WAPI_PASSWORD":	Password needs to be generated and IP allowed in the admin interface`)
		ew.writeln()

		ew.writeln(`Additional Configuration:`)
		ew.writeln(`	- "WEDOS_HTTP_TIMEOUT":	API request timeout`)
		ew.writeln(`	- "WEDOS_POLLING_INTERVAL":	Time between DNS propagation check`)
		ew.writeln(`	- "WEDOS_PROPAGATION_TIMEOUT":	Maximum waiting time for DNS propagation`)
		ew.writeln(`	- "WEDOS_TTL":	The TTL of the TXT record used for the DNS challenge`)

		ew.writeln()
		ew.writeln(`More information: https://go-acme.github.io/lego/dns/wedos`)

	case "yandex":
		// generated from: providers/dns/yandex/yandex.toml
		ew.writeln(`Configuration for Yandex PDD.`)
		ew.writeln(`Code:	'yandex'`)
		ew.writeln(`Since:	'v3.7.0'`)
		ew.writeln()

		ew.writeln(`Credentials:`)
		ew.writeln(`	- "YANDEX_PDD_TOKEN":	Basic authentication username`)
		ew.writeln()

		ew.writeln(`Additional Configuration:`)
		ew.writeln(`	- "YANDEX_HTTP_TIMEOUT":	API request timeout`)
		ew.writeln(`	- "YANDEX_POLLING_INTERVAL":	Time between DNS propagation check`)
		ew.writeln(`	- "YANDEX_PROPAGATION_TIMEOUT":	Maximum waiting time for DNS propagation`)
		ew.writeln(`	- "YANDEX_TTL":	The TTL of the TXT record used for the DNS challenge`)

		ew.writeln()
		ew.writeln(`More information: https://go-acme.github.io/lego/dns/yandex`)

	case "yandexcloud":
		// generated from: providers/dns/yandexcloud/yandexcloud.toml
		ew.writeln(`Configuration for Yandex Cloud.`)
		ew.writeln(`Code:	'yandexcloud'`)
		ew.writeln(`Since:	'v4.9.0'`)
		ew.writeln()

		ew.writeln(`Credentials:`)
		ew.writeln(`	- "YANDEX_CLOUD_FOLDER_ID":	The string id of folder (aka project) in Yandex Cloud`)
		ew.writeln(`	- "YANDEX_CLOUD_IAM_TOKEN":	The base64 encoded json which contains inforamtion about iam token of serivce account with 'dns.admin' permissions`)
		ew.writeln()

		ew.writeln(`Additional Configuration:`)
		ew.writeln(`	- "YANDEX_CLOUD_POLLING_INTERVAL":	Time between DNS propagation check`)
		ew.writeln(`	- "YANDEX_CLOUD_PROPAGATION_TIMEOUT":	Maximum waiting time for DNS propagation`)
		ew.writeln(`	- "YANDEX_CLOUD_TTL":	The TTL of the TXT record used for the DNS challenge`)

		ew.writeln()
		ew.writeln(`More information: https://go-acme.github.io/lego/dns/yandexcloud`)

	case "zoneee":
		// generated from: providers/dns/zoneee/zoneee.toml
		ew.writeln(`Configuration for Zone.ee.`)
		ew.writeln(`Code:	'zoneee'`)
		ew.writeln(`Since:	'v2.1.0'`)
		ew.writeln()

		ew.writeln(`Credentials:`)
		ew.writeln(`	- "ZONEEE_API_KEY":	API key`)
		ew.writeln(`	- "ZONEEE_API_USER":	API user`)
		ew.writeln()

		ew.writeln(`Additional Configuration:`)
		ew.writeln(`	- "ZONEEE_ENDPOINT":	API endpoint URL`)
		ew.writeln(`	- "ZONEEE_HTTP_TIMEOUT":	API request timeout`)
		ew.writeln(`	- "ZONEEE_POLLING_INTERVAL":	Time between DNS propagation check`)
		ew.writeln(`	- "ZONEEE_PROPAGATION_TIMEOUT":	Maximum waiting time for DNS propagation`)
		ew.writeln(`	- "ZONEEE_TTL":	The TTL of the TXT record used for the DNS challenge`)

		ew.writeln()
		ew.writeln(`More information: https://go-acme.github.io/lego/dns/zoneee`)

	case "zonomi":
		// generated from: providers/dns/zonomi/zonomi.toml
		ew.writeln(`Configuration for Zonomi.`)
		ew.writeln(`Code:	'zonomi'`)
		ew.writeln(`Since:	'v3.5.0'`)
		ew.writeln()

		ew.writeln(`Credentials:`)
		ew.writeln(`	- "ZONOMI_API_KEY":	User API key`)
		ew.writeln()

		ew.writeln(`Additional Configuration:`)
		ew.writeln(`	- "ZONOMI_HTTP_TIMEOUT":	API request timeout`)
		ew.writeln(`	- "ZONOMI_POLLING_INTERVAL":	Time between DNS propagation check`)
		ew.writeln(`	- "ZONOMI_PROPAGATION_TIMEOUT":	Maximum waiting time for DNS propagation`)
		ew.writeln(`	- "ZONOMI_TTL":	The TTL of the TXT record used for the DNS challenge`)

		ew.writeln()
		ew.writeln(`More information: https://go-acme.github.io/lego/dns/zonomi`)

	case "manual":
		ew.writeln(`Solving the DNS-01 challenge using CLI prompt.`)
	default:
		return fmt.Errorf("%q is not yet supported", name)
	}

	if flusher, ok := w.(interface{ Flush() error }); ok {
		return flusher.Flush()
	}
	return nil
}
