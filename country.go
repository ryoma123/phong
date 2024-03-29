package phong

// country struct
type country struct {
	flag string
	name string
}

// countryList variable
var countryList = map[string]*country{
	"AF": &country{flag: "🇦🇫", name: "Afghanistan"},
	"AX": &country{flag: "🇦🇽", name: "Åland Islands"},
	"AL": &country{flag: "🇦🇱", name: "Albania"},
	"DZ": &country{flag: "🇩🇿", name: "Algeria"},
	"AS": &country{flag: "🇦🇸", name: "American Samoa"},
	"AD": &country{flag: "🇦🇩", name: "Andorra"},
	"AO": &country{flag: "🇦🇴", name: "Angola"},
	"AI": &country{flag: "🇦🇮", name: "Anguilla"},
	"AQ": &country{flag: "🇦🇶", name: "Antarctica"},
	"AG": &country{flag: "🇦🇬", name: "Antigua and Barbuda"},
	"AR": &country{flag: "🇦🇷", name: "Argentina"},
	"AM": &country{flag: "🇦🇲", name: "Armenia"},
	"AW": &country{flag: "🇦🇼", name: "Aruba"},
	"AU": &country{flag: "🇦🇺", name: "Australia"},
	"AT": &country{flag: "🇦🇹", name: "Austria"},
	"AZ": &country{flag: "🇦🇿", name: "Azerbaijan"},
	"BS": &country{flag: "🇧🇸", name: "Bahamas"},
	"BH": &country{flag: "🇧🇭", name: "Bahrain"},
	"BD": &country{flag: "🇧🇩", name: "Bangladesh"},
	"BB": &country{flag: "🇧🇧", name: "Barbados"},
	"BY": &country{flag: "🇧🇾", name: "Belarus"},
	"BE": &country{flag: "🇧🇪", name: "Belgium"},
	"BZ": &country{flag: "🇧🇿", name: "Belize"},
	"BJ": &country{flag: "🇧🇯", name: "Benin"},
	"BM": &country{flag: "🇧🇲", name: "Bermuda"},
	"BT": &country{flag: "🇧🇹", name: "Bhutan"},
	"BO": &country{flag: "🇧🇴", name: "Bolivia (Plurinational State of)"},
	"BQ": &country{flag: "🇧🇶", name: "Bonaire, Sint Eustatius and Saba"},
	"BA": &country{flag: "🇧🇦", name: "Bosnia and Herzegovina"},
	"BW": &country{flag: "🇧🇼", name: "Botswana"},
	"BV": &country{flag: "🇳🇴", name: "Bouvet Island"},
	"BR": &country{flag: "🇧🇷", name: "Brazil"},
	"IO": &country{flag: "🇮🇴", name: "British Indian Ocean Territory"},
	"BN": &country{flag: "🇧🇳", name: "Brunei Darussalam"},
	"BG": &country{flag: "🇧🇬", name: "Bulgaria"},
	"BF": &country{flag: "🇧🇫", name: "Burkina Faso"},
	"BI": &country{flag: "🇧🇮", name: "Burundi"},
	"CV": &country{flag: "🇨🇻", name: "Cabo Verde"},
	"KH": &country{flag: "🇰🇭", name: "Cambodia"},
	"CM": &country{flag: "🇨🇲", name: "Cameroon"},
	"CA": &country{flag: "🇨🇦", name: "Canada"},
	"KY": &country{flag: "🇰🇾", name: "Cayman Islands"},
	"CF": &country{flag: "🇨🇫", name: "Central African Republic"},
	"TD": &country{flag: "🇹🇩", name: "Chad"},
	"CL": &country{flag: "🇨🇱", name: "Chile"},
	"CN": &country{flag: "🇨🇳", name: "China"},
	"CX": &country{flag: "🇨🇽", name: "Christmas Island"},
	"CC": &country{flag: "🇨🇨", name: "Cocos (Keeling) Islands"},
	"CO": &country{flag: "🇨🇴", name: "Colombia"},
	"KM": &country{flag: "🇰🇲", name: "Comoros"},
	"CG": &country{flag: "🇨🇬", name: "Congo"},
	"CD": &country{flag: "🇨🇩", name: "Congo, Democratic Republic of the"},
	"CK": &country{flag: "🇨🇰", name: "Cook Islands"},
	"CR": &country{flag: "🇨🇷", name: "Costa Rica"},
	"CI": &country{flag: "🇨🇮", name: "Côte d'Ivoire"},
	"HR": &country{flag: "🇭🇷", name: "Croatia"},
	"CU": &country{flag: "🇨🇺", name: "Cuba"},
	"CW": &country{flag: "🇨🇼", name: "Curaçao"},
	"CY": &country{flag: "🇨🇾", name: "Cyprus"},
	"CZ": &country{flag: "🇨🇿", name: "Czechia"},
	"DK": &country{flag: "🇩🇰", name: "Denmark"},
	"DJ": &country{flag: "🇩🇯", name: "Djibouti"},
	"DM": &country{flag: "🇩🇲", name: "Dominica"},
	"DO": &country{flag: "🇩🇴", name: "Dominican Republic"},
	"EC": &country{flag: "🇪🇨", name: "Ecuador"},
	"EG": &country{flag: "🇪🇬", name: "Egypt"},
	"SV": &country{flag: "🇸🇻", name: "El Salvador"},
	"GQ": &country{flag: "🇬🇶", name: "Equatorial Guinea"},
	"ER": &country{flag: "🇪🇷", name: "Eritrea"},
	"EE": &country{flag: "🇪🇪", name: "Estonia"},
	"SZ": &country{flag: "🇸🇿", name: "Eswatini"},
	"ET": &country{flag: "🇪🇹", name: "Ethiopia"},
	"FK": &country{flag: "🇫🇰", name: "Falkland Islands (Malvinas)"},
	"FO": &country{flag: "🇫🇴", name: "Faroe Islands"},
	"FJ": &country{flag: "🇫🇯", name: "Fiji"},
	"FI": &country{flag: "🇫🇮", name: "Finland"},
	"FR": &country{flag: "🇫🇷", name: "France"},
	"GF": &country{flag: "🇬🇫", name: "French Guiana"},
	"PF": &country{flag: "🇵🇫", name: "French Polynesia"},
	"TF": &country{flag: "🇹🇫", name: "French Southern Territories"},
	"GA": &country{flag: "🇬🇦", name: "Gabon"},
	"GM": &country{flag: "🇬🇲", name: "Gambia"},
	"GE": &country{flag: "🇬🇪", name: "Georgia"},
	"DE": &country{flag: "🇩🇪", name: "Germany"},
	"GH": &country{flag: "🇬🇭", name: "Ghana"},
	"GI": &country{flag: "🇬🇮", name: "Gibraltar"},
	"GR": &country{flag: "🇬🇷", name: "Greece"},
	"GL": &country{flag: "🇬🇱", name: "Greenland"},
	"GD": &country{flag: "🇬🇩", name: "Grenada"},
	"GP": &country{flag: "🇬🇵", name: "Guadeloupe"},
	"GU": &country{flag: "🇬🇺", name: "Guam"},
	"GT": &country{flag: "🇬🇹", name: "Guatemala"},
	"GG": &country{flag: "🇬🇬", name: "Guernsey"},
	"GN": &country{flag: "🇬🇳", name: "Guinea"},
	"GW": &country{flag: "🇬🇼", name: "Guinea-Bissau"},
	"GY": &country{flag: "🇬🇾", name: "Guyana"},
	"HT": &country{flag: "🇭🇹", name: "Haiti"},
	"HM": &country{flag: "🇭🇲", name: "Heard Island and McDonald Islands"},
	"VA": &country{flag: "🇻🇦", name: "Holy See"},
	"HN": &country{flag: "🇭🇳", name: "Honduras"},
	"HK": &country{flag: "🇭🇰", name: "Hong Kong"},
	"HU": &country{flag: "🇭🇺", name: "Hungary"},
	"IS": &country{flag: "🇮🇸", name: "Iceland"},
	"IN": &country{flag: "🇮🇳", name: "India"},
	"ID": &country{flag: "🇮🇩", name: "Indonesia"},
	"IR": &country{flag: "🇮🇷", name: "Iran (Islamic Republic of)"},
	"IQ": &country{flag: "🇮🇶", name: "Iraq"},
	"IE": &country{flag: "🇮🇪", name: "Ireland"},
	"IM": &country{flag: "🇮🇲", name: "Isle of Man"},
	"IL": &country{flag: "🇮🇱", name: "Israel"},
	"IT": &country{flag: "🇮🇹", name: "Italy"},
	"JM": &country{flag: "🇯🇲", name: "Jamaica"},
	"JP": &country{flag: "🇯🇵", name: "Japan"},
	"JE": &country{flag: "🇯🇪", name: "Jersey"},
	"JO": &country{flag: "🇯🇴", name: "Jordan"},
	"KZ": &country{flag: "🇰🇿", name: "Kazakhstan"},
	"KE": &country{flag: "🇰🇪", name: "Kenya"},
	"KI": &country{flag: "🇰🇮", name: "Kiribati"},
	"KP": &country{flag: "🇰🇵", name: "Korea (Democratic People's Republic of)"},
	"KR": &country{flag: "🇰🇷", name: "Korea, Republic of"},
	"KW": &country{flag: "🇰🇼", name: "Kuwait"},
	"KG": &country{flag: "🇰🇬", name: "Kyrgyzstan"},
	"LA": &country{flag: "🇱🇦", name: "Lao People's Democratic Republic"},
	"LV": &country{flag: "🇱🇻", name: "Latvia"},
	"LB": &country{flag: "🇱🇧", name: "Lebanon"},
	"LS": &country{flag: "🇱🇸", name: "Lesotho"},
	"LR": &country{flag: "🇱🇷", name: "Liberia"},
	"LY": &country{flag: "🇱🇾", name: "Libya"},
	"LI": &country{flag: "🇱🇮", name: "Liechtenstein"},
	"LT": &country{flag: "🇱🇹", name: "Lithuania"},
	"LU": &country{flag: "🇱🇺", name: "Luxembourg"},
	"MO": &country{flag: "🇲🇴", name: "Macao"},
	"MG": &country{flag: "🇲🇬", name: "Madagascar"},
	"MW": &country{flag: "🇲🇼", name: "Malawi"},
	"MY": &country{flag: "🇲🇾", name: "Malaysia"},
	"MV": &country{flag: "🇲🇻", name: "Maldives"},
	"ML": &country{flag: "🇲🇱", name: "Mali"},
	"MT": &country{flag: "🇲🇹", name: "Malta"},
	"MH": &country{flag: "🇲🇭", name: "Marshall Islands"},
	"MQ": &country{flag: "🇲🇶", name: "Martinique"},
	"MR": &country{flag: "🇲🇷", name: "Mauritania"},
	"MU": &country{flag: "🇲🇺", name: "Mauritius"},
	"YT": &country{flag: "🇾🇹", name: "Mayotte"},
	"MX": &country{flag: "🇲🇽", name: "Mexico"},
	"FM": &country{flag: "🇫🇲", name: "Micronesia (Federated States of)"},
	"MD": &country{flag: "🇲🇩", name: "Moldova, Republic of"},
	"MC": &country{flag: "🇲🇨", name: "Monaco"},
	"MN": &country{flag: "🇲🇳", name: "Mongolia"},
	"ME": &country{flag: "🇲🇪", name: "Montenegro"},
	"MS": &country{flag: "🇲🇸", name: "Montserrat"},
	"MA": &country{flag: "🇲🇦", name: "Morocco"},
	"MZ": &country{flag: "🇲🇿", name: "Mozambique"},
	"MM": &country{flag: "🇲🇲", name: "Myanmar"},
	"NA": &country{flag: "🇳🇦", name: "Namibia"},
	"NR": &country{flag: "🇳🇷", name: "Nauru"},
	"NP": &country{flag: "🇳🇵", name: "Nepal"},
	"NL": &country{flag: "🇳🇱", name: "Netherlands"},
	"NC": &country{flag: "🇳🇨", name: "New Caledonia"},
	"NZ": &country{flag: "🇳🇿", name: "New Zealand"},
	"NI": &country{flag: "🇳🇮", name: "Nicaragua"},
	"NE": &country{flag: "🇳🇪", name: "Niger"},
	"NG": &country{flag: "🇳🇬", name: "Nigeria"},
	"NU": &country{flag: "🇳🇺", name: "Niue"},
	"NF": &country{flag: "🇳🇫", name: "Norfolk Island"},
	"MK": &country{flag: "🇲🇰", name: "North Macedonia"},
	"MP": &country{flag: "🇲🇵", name: "Northern Mariana Islands"},
	"NO": &country{flag: "🇳🇴", name: "Norway"},
	"OM": &country{flag: "🇴🇲", name: "Oman"},
	"PK": &country{flag: "🇵🇰", name: "Pakistan"},
	"PW": &country{flag: "🇵🇼", name: "Palau"},
	"PS": &country{flag: "🇵🇸", name: "Palestine, State of"},
	"PA": &country{flag: "🇵🇦", name: "Panama"},
	"PG": &country{flag: "🇵🇬", name: "Papua New Guinea"},
	"PY": &country{flag: "🇵🇾", name: "Paraguay"},
	"PE": &country{flag: "🇵🇪", name: "Peru"},
	"PH": &country{flag: "🇵🇭", name: "Philippines"},
	"PN": &country{flag: "🇵🇳", name: "Pitcairn"},
	"PL": &country{flag: "🇵🇱", name: "Poland"},
	"PT": &country{flag: "🇵🇹", name: "Portugal"},
	"PR": &country{flag: "🇵🇷", name: "Puerto Rico"},
	"QA": &country{flag: "🇶🇦", name: "Qatar"},
	"RE": &country{flag: "🇷🇪", name: "Réunion"},
	"RO": &country{flag: "🇷🇴", name: "Romania"},
	"RU": &country{flag: "🇷🇺", name: "Russian Federation"},
	"RW": &country{flag: "🇷🇼", name: "Rwanda"},
	"BL": &country{flag: "🇧🇱", name: "Saint Barthélemy"},
	"SH": &country{flag: "🇸🇭", name: "Saint Helena, Ascension and Tristan da Cunha"},
	"KN": &country{flag: "🇰🇳", name: "Saint Kitts and Nevis"},
	"LC": &country{flag: "🇱🇨", name: "Saint Lucia"},
	"MF": &country{flag: "🇫🇷", name: "Saint Martin (French part)"},
	"PM": &country{flag: "🇵🇲", name: "Saint Pierre and Miquelon"},
	"VC": &country{flag: "🇻🇨", name: "Saint Vincent and the Grenadines"},
	"WS": &country{flag: "🇼🇸", name: "Samoa"},
	"SM": &country{flag: "🇸🇲", name: "San Marino"},
	"ST": &country{flag: "🇸🇹", name: "Sao Tome and Principe"},
	"SA": &country{flag: "🇸🇦", name: "Saudi Arabia"},
	"SN": &country{flag: "🇸🇳", name: "Senegal"},
	"RS": &country{flag: "🇷🇸", name: "Serbia"},
	"SC": &country{flag: "🇸🇨", name: "Seychelles"},
	"SL": &country{flag: "🇸🇱", name: "Sierra Leone"},
	"SG": &country{flag: "🇸🇬", name: "Singapore"},
	"SX": &country{flag: "🇸🇽", name: "Sint Maarten (Dutch part)"},
	"SK": &country{flag: "🇸🇰", name: "Slovakia"},
	"SI": &country{flag: "🇸🇮", name: "Slovenia"},
	"SB": &country{flag: "🇸🇧", name: "Solomon Islands"},
	"SO": &country{flag: "🇸🇴", name: "Somalia"},
	"ZA": &country{flag: "🇿🇦", name: "South Africa"},
	"GS": &country{flag: "🇬🇸", name: "South Georgia and the South Sandwich Islands"},
	"SS": &country{flag: "🇸🇸", name: "South Sudan"},
	"ES": &country{flag: "🇪🇸", name: "Spain"},
	"LK": &country{flag: "🇱🇰", name: "Sri Lanka"},
	"SD": &country{flag: "🇸🇩", name: "Sudan"},
	"SR": &country{flag: "🇸🇷", name: "Suriname"},
	"SJ": &country{flag: "🇳🇴", name: "Svalbard and Jan Mayen"},
	"SE": &country{flag: "🇸🇪", name: "Sweden"},
	"CH": &country{flag: "🇨🇭", name: "Switzerland"},
	"SY": &country{flag: "🇸🇾", name: "Syrian Arab Republic"},
	"TW": &country{flag: "🇹🇼", name: "Taiwan, Province of China"},
	"TJ": &country{flag: "🇹🇯", name: "Tajikistan"},
	"TZ": &country{flag: "🇹🇿", name: "Tanzania, United Republic of"},
	"TH": &country{flag: "🇹🇭", name: "Thailand"},
	"TL": &country{flag: "🇹🇱", name: "Timor-Leste"},
	"TG": &country{flag: "🇹🇬", name: "Togo"},
	"TK": &country{flag: "🇹🇰", name: "Tokelau"},
	"TO": &country{flag: "🇹🇴", name: "Tonga"},
	"TT": &country{flag: "🇹🇹", name: "Trinidad and Tobago"},
	"TN": &country{flag: "🇹🇳", name: "Tunisia"},
	"TR": &country{flag: "🇹🇷", name: "Turkey"},
	"TM": &country{flag: "🇹🇲", name: "Turkmenistan"},
	"TC": &country{flag: "🇹🇨", name: "Turks and Caicos Islands"},
	"TV": &country{flag: "🇹🇻", name: "Tuvalu"},
	"UG": &country{flag: "🇺🇬", name: "Uganda"},
	"UA": &country{flag: "🇺🇦", name: "Ukraine"},
	"AE": &country{flag: "🇦🇪", name: "United Arab Emirates"},
	"GB": &country{flag: "🇬🇧", name: "United Kingdom of Great Britain and Northern Ireland"},
	"US": &country{flag: "🇺🇸", name: "United States of America"},
	"UM": &country{flag: "🇺🇸", name: "United States Minor Outlying Islands"},
	"UY": &country{flag: "🇺🇾", name: "Uruguay"},
	"UZ": &country{flag: "🇺🇿", name: "Uzbekistan"},
	"VU": &country{flag: "🇻🇺", name: "Vanuatu"},
	"VE": &country{flag: "🇻🇪", name: "Venezuela (Bolivarian Republic of)"},
	"VN": &country{flag: "🇻🇳", name: "Viet Nam"},
	"VG": &country{flag: "🇻🇬", name: "Virgin Islands (British)"},
	"VI": &country{flag: "🇻🇮", name: "Virgin Islands (U.S.)"},
	"WF": &country{flag: "🇼🇫", name: "Wallis and Futuna"},
	"EH": &country{flag: "🇪🇭", name: "Western Sahara"},
	"YE": &country{flag: "🇾🇪", name: "Yemen"},
	"ZM": &country{flag: "🇿🇲", name: "Zambia"},
	"ZW": &country{flag: "🇿🇼", name: "Zimbabwe"},
}

// newCountry returns country
func newCountry(region string) *country {
	return countryList[region]
}
