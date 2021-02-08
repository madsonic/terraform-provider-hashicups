package hashicups

import (
	"context"
	"strconv"

	hc "github.com/hashicorp-demoapp/hashicups-client-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceOrder() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceOrderRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"coffee_id": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"coffee_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"coffee_description": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"coffee_image": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"coffee_price": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"coffee_teaser": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"quantity": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceOrderRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	c := m.(*hc.Client)

	orderID := strconv.Itoa(d.Get("id").(int))

	order, err := c.GetOrder(orderID)
	if err != nil {
		return diag.FromErr(err)
	}

	orderItems := flattenOrderItemsData(&order.Items)

	if err := d.Set("items", orderItems); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(orderID)

	return diags
}

func flattenOrderItemsData(orderItems *[]hc.OrderItem) []interface{} {
	if orderItems == nil {
		return make([]interface{}, 0)
	}

	items := make([]interface{}, len(*orderItems), len(*orderItems))

	for i, orderItem := range *orderItems {
		m := make(map[string]interface{})

		m["coffee_id"] = orderItem.Coffee.ID
		m["coffee_name"] = orderItem.Coffee.Name
		m["coffee_description"] = orderItem.Coffee.Description
		m["coffee_teaser"] = orderItem.Coffee.Teaser
		m["coffee_price"] = orderItem.Coffee.Price
		m["coffee_image"] = orderItem.Coffee.Image
		m["quantity"] = orderItem.Quantity

		items[i] = m
	}

	return items
}
