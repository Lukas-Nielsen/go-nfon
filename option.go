package nfon

type LinkRel string
type DataName string

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
	BLACKLIST_PROFILE        LinkRel = "blacklistProfile"
	CUSTOMER_CONTRACT        LinkRel = "customerContract"
	INTERNAL_OUTGOING        LinkRel = "internalOutgoing"
	PREFERRED_OUTBOUND_TRUNK LinkRel = "preferredOutboundTrunk"
	EMERGENCY_SITE           LinkRel = "emergencySite" // INHERIT
	CTI_INFO                 LinkRel = "ctiInfo"
	TARGET                   LinkRel = "target"
)

// possible data names
const (
	EXTENSION_NUMBER                          DataName = "extensionNumber"
	DISPLAY_NAME                              DataName = "displayName"
	ACCESS_CENTRAL_PHONE_BOOK                 DataName = "accessCentralPhoneBook"
	AUTODIAL_TIMEOUT                          DataName = "autodialTimeout"
	INTERCOM_ENABLED                          DataName = "intercomEnabled"
	NUMBERGUESSING_LENGTH                     DataName = "numberguessingLength"
	CALL_WAITING_INDICATION                   DataName = "callWaitingIndication"
	REPLICATE_AGENT                           DataName = "replicateAgent"
	CLIR_ENABLED                              DataName = "clirEnabled"
	NCONTROL_ENABLED                          DataName = "ncontrolEnabled"
	CCBS                                      DataName = "ccbs"
	PHONE_BOOK_HIDE                           DataName = "phoneBookHide"
	TIMEOUT_AFTER_SIP_TRANSFER                DataName = "timeoutAfterSipTransfer"
	COST_CENTER                               DataName = "costCenter"
	ABANDON_OTHER_SOFTPHONES                  DataName = "abandonOtherSoftphones"
	CLICK_TO_DIAL_STATE                       DataName = "clickToDialState"
	RECORDING_MODE                            DataName = "recordingMode"
	MEDIA_GATEWAY_EMERGENCY_DIALPLAN_PRIORITY DataName = "mediaGatewayEmergencyDialplanPriority"
	LANGUAGE                                  DataName = "language"
	DIAL_PREFIX                               DataName = "dialPrefix"
	N_MEETING                                 DataName = "nMeeting"
	KEY_NUMBER                                DataName = "keyNumber"
	TYPE                                      DataName = "type"
	FUNCTION_CODE                             DataName = "functionCode"
)

type Option struct {
	link map[LinkRel]string
	data map[DataName]any
}

func (o *Option) SetData(name DataName, value any) *Option {
	o.data[name] = value
	return o
}

func (o *Option) SetLink(rel LinkRel, href string) *Option {
	o.link[rel] = href
	return o
}
