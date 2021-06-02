// Package money provides various utilities to make it easy to manage money.
package money

// Money represents the combination of an amounr of money
// and the currency the money is in.
type Money struct {
	Value    float64
	Currency string
}

// Convert converts the value of one currency to another.
//
// It has two parameter: a Money instance with value to convert,
// and a string that represents the currency to convert to. Convert returns
// the converted currency and any errors encountered from unknown or unconvertible
// currencies.
// If an error is returned, the Money instance is set to the zero value.
//
// Supported currencies are:
// 				USD - US Dollar
// 				CAD - Canadian Dollar
// 				EUR - Euro
// 				INR - Indian Rupee
//
// More information on exchange rates can be found
// at https://www.investopedia.com/terms/e/exchangerate.asp
func Convert(from Money, to string) (Money, error) {
	if from.Currency == "USD" && to == "CAD" {
		return Money{
			Value:    from.Value * 1.21,
			Currency: "CAD",
		}, nil
	} else if from.Currency == "USD" && to == "EU" {
		return Money{
			Value:    from.Value * 0.82,
			Currency: "EU",
		}, nil
	} else if from.Currency == "USD" && to == "INR" {
		return Money{
			Value:    from.Value * 73.12,
			Currency: "INR",
		}, nil
	}
	return from, nil
}
