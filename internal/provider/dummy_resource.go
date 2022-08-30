package provider

import (
	"context"
	"encoding/json"
	"os"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var _ schema.UpdateContextFunc = dummyUpdate

func dummyResource() *schema.Resource {
	return &schema.Resource{
		Description: "Dummy resource",

		CreateContext: dummyCreate,
		ReadContext:   dummyRead,
		UpdateContext: dummyUpdate,
		DeleteContext: dummyDelete,

		CustomizeDiff: func(ctx context.Context, diff *schema.ResourceDiff, _ any) error {
			if os.Getenv("CUSTOMIZE_DIFF_PANIC") != "" {
				panic("customize_diff called with CUSTOMIZE_DIFF_PANIC=1")
			}
			return nil
		},

		Schema: map[string]*schema.Schema{
			"job": {
				Type:     schema.TypeList,
				Elem:     newJobSchema(),
				MaxItems: 1,
				Required: true,
			},
		},
	}
}

func newJobSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"q": {
				Type:     schema.TypeString,
				Required: true,
				StateFunc: func(v any) string {
					return strings.TrimSpace(v.(string))
				},
			},
		},
	}
}

func dummyCreate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	jobs := d.Get("job").([]any)
	if len(jobs) > 1 {
		jobsMarshalled, err := json.Marshal(jobs)
		if err != nil {
			return diag.FromErr(err)
		}

		panic("got unexpected jobs: " + string(jobsMarshalled))
	}

	d.SetId("created")
	return nil
}

func dummyUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	return nil
}

func dummyRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	return nil
}

func dummyDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	return nil
}
