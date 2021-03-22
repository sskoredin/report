package iiko_biz_client

const (
	apiUrl  = "/api/0"
	authUrl = apiUrl + "/auth/access_token"

	customersAPI                                 = apiUrl + "/customers"
	customerAddCategory                          = customersAPI + "%s/add_category"
	customerRemoveCategory                       = customersAPI + "%s/remove_category"
	customersGetCustomersByOrganizationAndPeriod = customersAPI + "/get_customers_by_organization_and_by_period"
	customersGetCustomerByID                     = customersAPI + "/get_customer_by_id"
	customersWithDrawBalance                     = customersAPI + "/withdraw_balance"
	customerCreateOrReplace                      = customersAPI + "/create_or_update"

	organizationsList             = apiUrl + "/organization/list"
	organisationApi               = apiUrl + "/organization/%s"
	organizationGuestCategory     = organisationApi + "/guest_categories"
	organizationTransactionReport = organisationApi + "/transactions_report"

	deliveryDiscounts = apiUrl + "deliverySettings/deliveryDiscounts"
	getNomenclature   = apiUrl + "nomenclature/%s"
	addOrder          = apiUrl + "/orders/add?"
	getPaymentsTypes  = apiUrl + "/rmsSettings/getPaymentTypes"

	nomenclature = apiUrl + "/nomenclature/%s"
)
