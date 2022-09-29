package destiny2

import (
	"github.com/crit/raid-helper/worker/internal/bungie"
	"github.com/crit/raid-helper/worker/internal/errors"
)

const (
	XurVendorHash = "2190858386"
)

type (
	VendorResult struct{} // TODO: what are the fields to return?

	vendorPayload struct{}
)

func Vendors(client *bungie.Client) (any, error) {
	var res bungie.Result

	// 400 = DestinyComponentTypeVendors (summary about the vendors)
	// 401 = DestinyComponentTypeVendorCategories (category of items on the vendor)
	// 402 = DestinyComponentTypeVendorSales (items being sold by the vendor)
	err := client.Get(&res, "/Destiny2/Vendors/?components=400,401,402")
	if err != nil {
		return nil, errors.Wrap(err)
	}

	return res, nil
}
