package hubspot

type ObjectType string

type ObjectTypeID string

const (
	// Defines the obeject type 'deals'
	ObjectTypeContacts ObjectType = "contacts"

	// Defines the obeject type 'deals'
	ObjectTypeDeals ObjectType = "deals"

	// Defines the obeject type 'deals'
	ObjectTypeTickets ObjectType = "tickets"
)

const (
	ObjectTypeIDContacts            ObjectTypeID = "0-1"
	ObjectTypeIDCompanies           ObjectTypeID = "0-2"
	ObjectTypeIDDeals               ObjectTypeID = "0-3"
	ObjectTypeIDTickets             ObjectTypeID = "0-5"
	ObjectTypeIDCalls               ObjectTypeID = "0-48"
	ObjectTypeIDEmails              ObjectTypeID = "0-49"
	ObjectTypeIDMeetings            ObjectTypeID = "0-47"
	ObjectTypeIDNotes               ObjectTypeID = "0-4"
	ObjectTypeIDTasks               ObjectTypeID = "0-27"
	ObjectTypeIDProducts            ObjectTypeID = "0-7"
	ObjectTypeIDInvoices            ObjectTypeID = "0-52"
	ObjectTypeIDLineItems           ObjectTypeID = "0-8"
	ObjectTypeIDPayments            ObjectTypeID = "0-101"
	ObjectTypeIDQuotes              ObjectTypeID = "0-14"
	ObjectTypeIDSubscriptions       ObjectTypeID = "0-69"
	ObjectTypeIDCommunications      ObjectTypeID = "0-18"
	ObjectTypeIDPostalMail          ObjectTypeID = "0-116"
	ObjectTypeIDMarketingEvents     ObjectTypeID = "0-54"
	ObjectTypeIDFeedbackSubmissions ObjectTypeID = "0-19"
)
