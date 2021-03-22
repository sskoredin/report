package enum

type TransactionType int

const (
	Undefined                   TransactionType = 0
	CloseOrder                  TransactionType = 1
	RefillWallet                TransactionType = 2
	ResetRefillWallet           TransactionType = 3
	PayFromWallet               TransactionType = 4
	CancelPayFromWallet         TransactionType = 5
	RefillWalletFromOrder       TransactionType = 6
	CancelRefillWalletFromOrder TransactionType = 7
	AutomaticRefillWallet       TransactionType = 8
	ResetAutomaticRefillWallet  TransactionType = 9
	AbortOrder                  TransactionType = 10
	RemoveGuestCategory         TransactionType = 11
	SetGuestCategory            TransactionType = 12
	DiscountSum                 TransactionType = 13
	CouponActivation            TransactionType = 14
	RefillWalletFromApi         TransactionType = 15
)

func NewTransactionType(s string) TransactionType {
	switch s {
	case "CloseOrder":
		return CloseOrder
	case "RefillWallet":
		return RefillWallet
	case "ResetRefillWallet":
		return ResetRefillWallet
	case "PayFromWallet":
		return PayFromWallet
	case "CancelPayFromWallet":
		return CancelPayFromWallet
	case "RefillWalletFromOrder":
		return RefillWalletFromOrder
	case "CancelRefillWalletFromOrder":
		return CancelRefillWalletFromOrder
	case "AutomaticRefillWallet":
		return AutomaticRefillWallet
	case "ResetAutomaticRefillWallet":
		return ResetAutomaticRefillWallet
	case "AbortOrder":
		return AbortOrder
	case "RemoveGuestCategory":
		return RemoveGuestCategory
	case "SetGuestCategory":
		return SetGuestCategory
	case "DiscountSum":
		return DiscountSum
	case "CouponActivation":
		return CouponActivation
	case "RefillWalletFromApi":
		return RefillWalletFromApi
	default:
		return Undefined
	}
}
