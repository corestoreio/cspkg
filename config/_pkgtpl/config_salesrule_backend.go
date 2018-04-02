// +build ignore

package salesrule

import (
	"github.com/corestoreio/pkg/config/cfgmodel"
	"github.com/corestoreio/pkg/config/element"
)

// Backend will be initialized in the init() function together with ConfigStructure.
var Backend *PkgBackend

// PkgBackend just exported for the sake of documentation. See fields
// for more information. The PkgBackend handles the reading and writing
// of configuration values within this package.
type PkgBackend struct {
	cfgmodel.PkgBackend
	// PromoAutoGeneratedCouponCodesLength => Code Length.
	// Excluding prefix, suffix and separators.
	// Path: promo/auto_generated_coupon_codes/length
	PromoAutoGeneratedCouponCodesLength cfgmodel.Str

	// PromoAutoGeneratedCouponCodesFormat => Code Format.
	// Path: promo/auto_generated_coupon_codes/format
	// SourceModel: Magento\SalesRule\Model\System\Config\Source\Coupon\Format
	PromoAutoGeneratedCouponCodesFormat cfgmodel.Str

	// PromoAutoGeneratedCouponCodesPrefix => Code Prefix.
	// Path: promo/auto_generated_coupon_codes/prefix
	PromoAutoGeneratedCouponCodesPrefix cfgmodel.Str

	// PromoAutoGeneratedCouponCodesSuffix => Code Suffix.
	// Path: promo/auto_generated_coupon_codes/suffix
	PromoAutoGeneratedCouponCodesSuffix cfgmodel.Str

	// PromoAutoGeneratedCouponCodesDash => Dash Every X Characters.
	// If empty no separation.
	// Path: promo/auto_generated_coupon_codes/dash
	PromoAutoGeneratedCouponCodesDash cfgmodel.Str

	// RssCatalogDiscounts => Coupons/Discounts.
	// Path: rss/catalog/discounts
	// SourceModel: Magento\Config\Model\Config\Source\Enabledisable
	RssCatalogDiscounts cfgmodel.Bool
}

// NewBackend initializes the global Backend variable. See init()
func NewBackend(cfgStruct element.Sections) *PkgBackend {
	return (&PkgBackend{}).init(cfgStruct)
}

func (pp *PkgBackend) init(cfgStruct element.Sections) *PkgBackend {
	pp.Lock()
	defer pp.Unlock()
	pp.PromoAutoGeneratedCouponCodesLength = cfgmodel.NewStr(`promo/auto_generated_coupon_codes/length`, cfgmodel.WithFieldFromSectionSlice(cfgStruct))
	pp.PromoAutoGeneratedCouponCodesFormat = cfgmodel.NewStr(`promo/auto_generated_coupon_codes/format`, cfgmodel.WithFieldFromSectionSlice(cfgStruct))
	pp.PromoAutoGeneratedCouponCodesPrefix = cfgmodel.NewStr(`promo/auto_generated_coupon_codes/prefix`, cfgmodel.WithFieldFromSectionSlice(cfgStruct))
	pp.PromoAutoGeneratedCouponCodesSuffix = cfgmodel.NewStr(`promo/auto_generated_coupon_codes/suffix`, cfgmodel.WithFieldFromSectionSlice(cfgStruct))
	pp.PromoAutoGeneratedCouponCodesDash = cfgmodel.NewStr(`promo/auto_generated_coupon_codes/dash`, cfgmodel.WithFieldFromSectionSlice(cfgStruct))
	pp.RssCatalogDiscounts = cfgmodel.NewBool(`rss/catalog/discounts`, cfgmodel.WithFieldFromSectionSlice(cfgStruct))

	return pp
}
