package calculator

type DiscountCalculator struct {
	minimumPurchaseAmount int
	disountAmount         int
}

func NewDiscountCalculator(minimumPurchaseAmount int, discountAmount int) *DiscountCalculator {
	return &DiscountCalculator{
		minimumPurchaseAmount: minimumPurchaseAmount,
		disountAmount:         discountAmount,
	}
}
func (c *DiscountCalculator) Calculate(purchaseAmount int) int {
	if purchaseAmount > c.minimumPurchaseAmount {
		return purchaseAmount - c.disountAmount
	}
	return purchaseAmount

}
