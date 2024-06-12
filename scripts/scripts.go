package scripts

import _ "embed"

//go:embed nmap.sh
var Nmap string

//go:embed LinEnum/LinEnum.sh
var LinEnum string

//go:embed linpeas/linpeas.sh
var Linpeas string
