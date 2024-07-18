package hubspot

import (
	"fmt"
)

// Associations service
type AssociationsService struct {
	service
}

type ListAssociationResponse struct {
	Associations *[]Association `json:"results,omitempty"`
}

type Association struct {
	ID    string              `json:"id,omitempty"`
	Type  string              `json:"type,omitempty"`
	To    *AssociationTo      `json:"to,omitempty"`
	Types *[]AssociationTypes `json:"types,omitempty"`
}

type AssociationTypes struct {
	AssociationCategory AssociationCategory `json:"associationCategory,omitempty"`
	AssociationTypeID   AssociationTypeID   `json:"associationTypeId,omitempty"`
}

type AssociationTo struct {
	ID string `json:"id,omitempty"`
}

type AssociationCategory string

type AssociationTypeID int32

const (
	AssociationCategoryHubspot    AssociationCategory = "HUBSPOT_DEFINED"
	AssociationCategoryUser       AssociationCategory = "USER_DEFINED"
	AssociationCategoryIntegrator AssociationCategory = "INTEGRATOR_DEFINED"
)

const (
	// Contact to Object Association IDs
	AssociationTypeContactToContact        AssociationTypeID = 449
	AssociationTypeContactToCompany        AssociationTypeID = 279
	AssociationTypeContactToCompanyPrimary AssociationTypeID = 1
	AssociationTypeContactToDeal           AssociationTypeID = 4
	AssociationTypeContactToTicket         AssociationTypeID = 15
	AssociationTypeContactToCall           AssociationTypeID = 193
	AssociationTypeContactToEmail          AssociationTypeID = 197
	AssociationTypeContactToMeeting        AssociationTypeID = 199
	AssociationTypeContactToNote           AssociationTypeID = 201
	AssociationTypeContactToTask           AssociationTypeID = 203
	AssociationTypeContactToCommunication  AssociationTypeID = 82
	AssociationTypeContactToPostalMail     AssociationTypeID = 454
	AssociationTypeContactToCart           AssociationTypeID = 587
	AssociationTypeContactToOrder          AssociationTypeID = 508
	AssociationTypeContactToInvoice        AssociationTypeID = 178
	AssociationTypeContactToPayment        AssociationTypeID = 388
	AssociationTypeContactToSubscription   AssociationTypeID = 296

	// Company to Object Association IDs
	AssociationTypeCompanyToCompany        AssociationTypeID = 450
	AssociationTypeChildToParentCompany    AssociationTypeID = 14
	AssociationTypeParentToChildCompany    AssociationTypeID = 13
	AssociationTypeCompanyToContact        AssociationTypeID = 280
	AssociationTypeCompanyToContactPrimary AssociationTypeID = 2
	AssociationTypeCompanyToDeal           AssociationTypeID = 342
	AssociationTypeCompanyToDealPrimary    AssociationTypeID = 6
	AssociationTypeCompanyToTicket         AssociationTypeID = 340
	AssociationTypeCompanyToTicketPrimary  AssociationTypeID = 25
	AssociationTypeCompanyToCall           AssociationTypeID = 181
	AssociationTypeCompanyToEmail          AssociationTypeID = 185
	AssociationTypeCompanyToMeeting        AssociationTypeID = 187
	AssociationTypeCompanyToNote           AssociationTypeID = 189
	AssociationTypeCompanyToTask           AssociationTypeID = 191
	AssociationTypeCompanyToCommunication  AssociationTypeID = 88
	AssociationTypeCompanyToPostalMail     AssociationTypeID = 460
	AssociationTypeCompanyToInvoice        AssociationTypeID = 180
	AssociationTypeCompanyToOrder          AssociationTypeID = 510
	AssociationTypeCompanyToPayment        AssociationTypeID = 390
	AssociationTypeCompanyToSubscription   AssociationTypeID = 298

	// Deal to Object Association IDs
	AssociationTypeDealToDeal           AssociationTypeID = 451
	AssociationTypeDealToContact        AssociationTypeID = 3
	AssociationTypeDealToCompany        AssociationTypeID = 341
	AssociationTypeDealToCompanyPrimary AssociationTypeID = 5
	AssociationTypeDealToTicket         AssociationTypeID = 27
	AssociationTypeDealToCall           AssociationTypeID = 205
	AssociationTypeDealToEmail          AssociationTypeID = 209
	AssociationTypeDealToMeeting        AssociationTypeID = 211
	AssociationTypeDealToNote           AssociationTypeID = 213
	AssociationTypeDealToTask           AssociationTypeID = 215
	AssociationTypeDealToCommunication  AssociationTypeID = 86
	AssociationTypeDealToPostalMail     AssociationTypeID = 458
	AssociationTypeDealToDealSplit      AssociationTypeID = 313
	AssociationTypeDealToLineItem       AssociationTypeID = 19
	AssociationTypeDealToInvoice        AssociationTypeID = 176
	AssociationTypeDealToOrder          AssociationTypeID = 511
	AssociationTypeDealToPayment        AssociationTypeID = 392
	AssociationTypeDealToProduct        AssociationTypeID = 630
	AssociationTypeDealToQuote          AssociationTypeID = 63
	AssociationTypeDealToSubscription   AssociationTypeID = 300

	// Ticket to Object Association IDs
	AssociationTypeTicketToTicket         AssociationTypeID = 452
	AssociationTypeTicketToContact        AssociationTypeID = 16
	AssociationTypeTicketToCompany        AssociationTypeID = 339
	AssociationTypeTicketToCompanyPrimary AssociationTypeID = 26
	AssociationTypeTicketToDeal           AssociationTypeID = 28
	AssociationTypeTicketToCall           AssociationTypeID = 219
	AssociationTypeTicketToEmail          AssociationTypeID = 223
	AssociationTypeTicketToMeeting        AssociationTypeID = 225
	AssociationTypeTicketToNote           AssociationTypeID = 227
	AssociationTypeTicketToTask           AssociationTypeID = 229
	AssociationTypeTicketToCommunication  AssociationTypeID = 84
	AssociationTypeTicketToPostalMail     AssociationTypeID = 456
	AssociationTypeTicketToThread         AssociationTypeID = 32
	AssociationTypeTicketToConversation   AssociationTypeID = 278
	AssociationTypeTicketToOrder          AssociationTypeID = 526

	// Lead to Object Association IDs
	AssociationTypeLeadToPrimaryContact AssociationTypeID = 578
	AssociationTypeLeadToCall           AssociationTypeID = 596
	AssociationTypeLeadToEmail          AssociationTypeID = 598
	AssociationTypeLeadToMeeting        AssociationTypeID = 600
	AssociationTypeLeadToCommunication  AssociationTypeID = 602
	AssociationTypeLeadToContact        AssociationTypeID = 608
	AssociationTypeLeadToCompany        AssociationTypeID = 610
	AssociationTypeLeadToTask           AssociationTypeID = 646

	// Call to Object Association IDs
	AssociationTypeCallToContact AssociationTypeID = 194
	AssociationTypeCallToCompany AssociationTypeID = 182
	AssociationTypeCallToDeal    AssociationTypeID = 206
	AssociationTypeCallToTicket  AssociationTypeID = 220

	// Email to Object Association IDs
	AssociationTypeEmailToContact AssociationTypeID = 198
	AssociationTypeEmailToCompany AssociationTypeID = 186
	AssociationTypeEmailToDeal    AssociationTypeID = 210
	AssociationTypeEmailToTicket  AssociationTypeID = 224

	// Meeting to Object Association IDs
	AssociationTypeMeetingToContact AssociationTypeID = 200
	AssociationTypeMeetingToCompany AssociationTypeID = 188
	AssociationTypeMeetingToDeal    AssociationTypeID = 212
	AssociationTypeMeetingToTicket  AssociationTypeID = 226

	// Note to Object Association IDs
	AssociationTypeNoteToContact AssociationTypeID = 202
	AssociationTypeNoteToCompany AssociationTypeID = 190
	AssociationTypeNoteToDeal    AssociationTypeID = 214
	AssociationTypeNoteToTicket  AssociationTypeID = 228

	// Postal Mail to Object Association IDs
	AssociationTypePostalMailToContact AssociationTypeID = 453
	AssociationTypePostalMailToCompany AssociationTypeID = 459
	AssociationTypePostalMailToDeal    AssociationTypeID = 457
	AssociationTypePostalMailToTicket  AssociationTypeID = 455

	// Quote to Object Association IDs
	AssociationTypeQuoteToContact              AssociationTypeID = 69
	AssociationTypeQuoteToCompany              AssociationTypeID = 71
	AssociationTypeQuoteToDeal                 AssociationTypeID = 64
	AssociationTypeQuoteToLineItem             AssociationTypeID = 67
	AssociationTypeQuoteToQuoteTemplate        AssociationTypeID = 286
	AssociationTypeQuoteToDiscount             AssociationTypeID = 362
	AssociationTypeQuoteToFee                  AssociationTypeID = 364
	AssociationTypeQuoteToTax                  AssociationTypeID = 366
	AssociationTypeContactSignerForESignatures AssociationTypeID = 702
	AssociationTypeQuoteToCart                 AssociationTypeID = 733
	AssociationTypeQuoteToInvoice              AssociationTypeID = 408
	AssociationTypeQuoteToOrder                AssociationTypeID = 731
	AssociationTypeQuoteToPayment              AssociationTypeID = 398
	AssociationTypeQuoteToSubscription         AssociationTypeID = 304

	// Task to Object Association IDs
	AssociationTypeTaskToContact AssociationTypeID = 204
	AssociationTypeTaskToCompany AssociationTypeID = 192
	AssociationTypeTaskToDeal    AssociationTypeID = 216
	AssociationTypeTaskToTicket  AssociationTypeID = 230

	// Communication to Object Association IDs
	AssociationTypeCommunicationToContact AssociationTypeID = 81
	AssociationTypeCommunicationToCompany AssociationTypeID = 87
	AssociationTypeCommunicationToDeal    AssociationTypeID = 85
	AssociationTypeCommunicationToTicket  AssociationTypeID = 83

	// Order to Object Association IDs
	AssociationTypeOrderToCart         AssociationTypeID = 593
	AssociationTypeOrderToContact      AssociationTypeID = 507
	AssociationTypeOrderToCompany      AssociationTypeID = 509
	AssociationTypeOrderToDeal         AssociationTypeID = 512
	AssociationTypeOrderToDiscount     AssociationTypeID = 519
	AssociationTypeOrderToDiscountCode AssociationTypeID = 521
	AssociationTypeOrderToInvoice      AssociationTypeID = 518
	AssociationTypeOrderToLineItem     AssociationTypeID = 513
	AssociationTypeOrderToPayment      AssociationTypeID = 523
	AssociationTypeOrderToQuote        AssociationTypeID = 730
	AssociationTypeOrderToSubscription AssociationTypeID = 516
	AssociationTypeOrderToTask         AssociationTypeID = 726
	AssociationTypeOrderToTicket       AssociationTypeID = 525

	// Cart to Object Association IDs
	AssociationTypeCartToContact  AssociationTypeID = 586
	AssociationTypeCartToDiscount AssociationTypeID = 588
	AssociationTypeCartToLineItem AssociationTypeID = 590
	AssociationTypeCartToOrder    AssociationTypeID = 592
	AssociationTypeCartToQuote    AssociationTypeID = 732
	AssociationTypeCartToTask     AssociationTypeID = 728
	AssociationTypeCartToTicket   AssociationTypeID = 594
)

// List associations by ID.
func (service *AssociationsService) List(objectType ObjectTypeID, id string, objectToType ObjectTypeID, opts QueryValues) (*ListAssociationResponse, *Response, error) {
	_url := fmt.Sprintf("/crm/%s/objects/%s/%s/associations/%s", *service.revision, objectType, id, objectToType)

	req, _ := service.client.NewRequest("GET", _url, opts, nil)

	data := new(ListAssociationResponse)
	response, err := service.client.Do(req, data)

	if err != nil {
		return nil, response, err
	}

	return data, response, nil
}
