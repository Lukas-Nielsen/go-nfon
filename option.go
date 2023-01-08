package nfon

type linkRel string
type dataName string

// predefind values
const (
	CLIR_ENABLED_SHOW_BASE_NUMBER     = "SHOW_BASE_NUMBER"
	CLIR_ENABLED_SHOW_COMPLETE_NUMBER = "SHOW_COMPLETE_NUMBER"
	CLIR_ENABLED_DISABLED             = "DISABLED"
	CLIR_ENABLED_SIGNAL_PILOT         = "SIGNAL_PILOT"

	CLICK_TO_DIAL_STATE_DEFAULT = "DEFAULT"
	CLICK_TO_DIAL_STATE_OFF     = "OFF"
	CLICK_TO_DIAL_STATE_ON      = "ON"

	RECORDING_MODE_NEVER    = "NEVER"
	RECORDING_MODE_ONDEMAND = "ONDEMAND"

	MEDIA_GATEWAY_EMERGENCY_DIALPLAN_PRIORITY_OFF   = "OFF"
	MEDIA_GATEWAY_EMERGENCY_DIALPLAN_PRIORITY_ONE   = "ONE"
	MEDIA_GATEWAY_EMERGENCY_DIALPLAN_PRIORITY_TWO   = "TWO"
	MEDIA_GATEWAY_EMERGENCY_DIALPLAN_PRIORITY_THREE = "THREE"

	AUTODIAL_TIMEOUT_0  = 0
	AUTODIAL_TIMEOUT_2  = 2
	AUTODIAL_TIMEOUT_5  = 5
	AUTODIAL_TIMEOUT_10 = 10
	AUTODIAL_TIMEOUT_15 = 15

	LANGUAGE_DE = "de"
	LANGUAGE_EN = "en"
	LANGUAGE_FR = "fr"
	LANGUAGE_IT = "it"
	LANGUAGE_NL = "nl"
	LANGUAGE_PL = "pl"
	LANGUAGE_HR = "hr"
	LANGUAGE_ES = "es"

	DIAL_PREFIX_0 = "0"
	DIAL_PREFIX_1 = "9"

	NUMBERGUESSING_LENGTH_0 = 0
	NUMBERGUESSING_LENGTH_4 = 4
	NUMBERGUESSING_LENGTH_6 = 6
)

// possible link rels
const (
	BLACKLIST_PROFILE        linkRel = "blacklistProfile"
	CUSTOMER_CONTRACT        linkRel = "customerContract"
	INTERNAL_OUTGOING        linkRel = "internalOutgoing"
	PREFERRED_OUTBOUND_TRUNK linkRel = "preferredOutboundTrunk"
	EMERGENCY_SITE           linkRel = "emergencySite" // INHERIT
	CTI_INFO                 linkRel = "ctiInfo"
)

// possible data names
const (
	EXTENSION_NUMBER                          dataName = "extensionNumber"
	DISPLAY_NAME                              dataName = "displayName"
	ACCESS_CENTRAL_PHONE_BOOK                 dataName = "accessCentralPhoneBook"
	AUTODIAL_TIMEOUT                          dataName = "autodialTimeout"
	INTERCOM_ENABLED                          dataName = "intercomEnabled"
	NUMBERGUESSING_LENGTH                     dataName = "numberguessingLength"
	CALL_WAITING_INDICATION                   dataName = "callWaitingIndication "
	REPLICATE_AGENT                           dataName = "replicateAgent "
	CLIR_ENABLED                              dataName = "clirEnabled"
	NCONTROL_ENABLED                          dataName = "ncontrolEnabled"
	CCBS                                      dataName = "ccbs "
	PHONE_BOOK_HIDE                           dataName = "phoneBookHide"
	TIMEOUT_AFTER_SIP_TRANSFER                dataName = "timeoutAfterSipTransfer "
	COST_CENTER                               dataName = "costCenter "
	ABANDON_OTHER_SOFTPHONES                  dataName = "abandonOtherSoftphones"
	CLICK_TO_DIAL_STATE                       dataName = "clickToDialState"
	RECORDING_MODE                            dataName = "recordingMode"
	MEDIA_GATEWAY_EMERGENCY_DIALPLAN_PRIORITY dataName = "mediaGatewayEmergencyDialplanPriority "
	LANGUAGE                                  dataName = "language "
	DIAL_PREFIX                               dataName = "dialPrefix "
	N_MEETING                                 dataName = "nMeeting "
)

type Option struct {
	link map[linkRel]string
	data map[dataName]any
}

func (o *Option) SetData(name dataName, value any) *Option {
	o.data[name] = value
	return o
}

func (o *Option) SetLink(rel linkRel, href string) *Option {
	o.link[rel] = href
	return o
}
