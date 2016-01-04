package idcf

import (
	"log"

	"github.com/atsaki/go-idcf/dns"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceIDCFDNSRecord() *schema.Resource {
	return &schema.Resource{
		Create: resourceIDCFDNSRecordCreate,
		Read:   resourceIDCFDNSRecordRead,
		Update: resourceIDCFDNSRecordUpdate,
		Delete: resourceIDCFDNSRecordDelete,

		Schema: map[string]*schema.Schema{
			"zone_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"value": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"ttl": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"priority": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceIDCFDNSRecordCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*IDCFClient).dns
	zoneid := d.Get("zone_id").(string)
	name := d.Get("name").(string)

	p := dns.NewCreateRecordParamter(
		zoneid, name, d.Get("type").(string), d.Get("value").(string))

	if v, ok := d.GetOk("ttl"); ok {
		p.SetTTL(v.(int))
	}

	if v, ok := d.GetOk("priority"); ok {
		p.SetPriority(v.(int))
	}

	log.Printf("[DEBUG] Creating IDCF DNS record: %s", name)
	r, err := client.CreateRecord(p)
	if err != nil {
		return err
	}

	d.SetId(r.UUID)

	return resourceIDCFDNSRecordUpdate(d, meta)
}

func resourceIDCFDNSRecordRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*IDCFClient).dns
	zoneid := d.Get("zone_id").(string)
	recordid := d.Id()

	r, err := client.Record(zoneid, recordid)
	if err != nil {
		return err
	}

	d.Set("name", r.Name)
	d.Set("type", r.Type)
	d.Set("value", r.Content)
	d.Set("ttl", r.TTL)
	d.Set("priority", r.Priority)

	return nil
}

func resourceIDCFDNSRecordUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*IDCFClient).dns
	zoneid := d.Get("zone_id").(string)
	recordid := d.Id()

	p := dns.NewUpdateRecordParamter(zoneid, recordid)
	changed := false

	if d.HasChange("name") {
		p.SetName(d.Get("name").(string))
		changed = true
	}

	if d.HasChange("type") {
		p.SetType(d.Get("type").(string))
		changed = true
	}

	if d.HasChange("value") {
		p.SetContent(d.Get("value").(string))
		changed = true
	}

	if d.HasChange("ttl") {
		p.SetTTL(d.Get("ttl").(int))
		changed = true
	}

	if d.HasChange("priority") {
		p.SetPriority(d.Get("priority").(int))
		changed = true
	}

	if changed {
		log.Printf("[DEBUG] Updating IDCF DNS record: %s (%s)", d.Get("name"), recordid)
		_, err := client.UpdateRecord(p)
		if err != nil {
			return err
		}
	}

	return resourceIDCFDNSRecordRead(d, meta)
}

func resourceIDCFDNSRecordDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*IDCFClient).dns
	zoneid := d.Get("zone_id").(string)
	recordid := d.Id()
	log.Printf("[DEBUG] Deleting IDCF DNS record: %s (%s)", d.Get("name"), recordid)
	err := client.DeleteRecord(zoneid, recordid)
	if err != nil {
		return err
	}
	d.SetId("")
	return nil
}
