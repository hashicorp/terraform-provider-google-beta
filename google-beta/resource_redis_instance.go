func resourceRedisInstance() *schema.Resource {
	return &schema.Resource{
		Create: resourceRedisInstanceCreate,
		Read:   resourceRedisInstanceRead,
		Update: resourceRedisInstanceUpdate,
		Delete: resourceRedisInstanceDelete,

		Importer: &schema.ResourceImporter{
			State: resourceRedisInstanceImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"memory_size_gb": {
				Type:        schema.TypeInt,
				Required:    true,
				ForceNew:    true,
				Description: ``,
			},

			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: ``,
			},

			"alternative_location_id": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
			},

			"authorized_network": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
			},

			"display_name": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
			},

			"labels": {
				Type:        schema.TypeMap,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},

			"location_id": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
			},

			"project": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Description: ``,
			},

			"redis_configs": {
				Type:        schema.TypeMap,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},

			"redis_version": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
			},

			"region": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Description: ``,
			},

			"reserved_ip_range": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
			},

			"tier": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
			},

			"current_location_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: ``,
			},

			"host": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: ``,
			},

			"port": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: ``,
			},
		},
	}
}

func resourceRedisInstanceCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	region, err := getRegion(d, config)
	if err != nil {
		return err
	}

	name := d.Get("name").(string)

	obj := &redis.Instance{
		MemorySizeGb:          dcl.Int64(int64(d.Get("memory_size_gb").(int))),
		Name:                  dcl.String(name),
		AlternativeLocationId: dcl.String(d.Get("alternative_location_id").(string)),
		AuthorizedNetwork:     dcl.String(d.Get("authorized_network").(string)),
		DisplayName:           dcl.String(d.Get("display_name").(string)),
		Labels:                expandMap(d.Get("labels"), d, config),
		LocationId:            dcl.String(d.Get("location_id").(string)),
		Project:               dcl.String(project),
		RedisConfigs:          expandMap(d.Get("redis_configs"), d, config),
		RedisVersion:          dcl.String(d.Get("redis_version").(string)),
		Region:                dcl.String(region),
		ReservedIPRange:       dcl.String(d.Get("reserved_ip_range").(string)),
		Tier:                  redis.InstanceTierEnumRef(d.Get("tier").(string)),
	}

	id, err := replaceVars(d, config, "projects/{{project}}/locations/{{region}}/instances/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	res, err := config.clientRedis.ApplyInstance(obj, CreateDirective...)
	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error creating Instance: %s", err)
	}

	log.Printf("[DEBUG] Finished creating Instance %q: %#v", d.Id(), res)

	return resourceRedisInstanceRead(d, meta)
}

func resourceRedisInstanceRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	region, err := getRegion(d, config)
	if err != nil {
		return err
	}

	name := d.Get("name").(string)

	res, err := config.clientRedis.GetInstance(project, region, name)
	if err != nil {
		// Resource not found
		d.SetId("")
		return err
	}

	if err = d.Set("memory_size_gb", res.MemorySizeGb); err != nil {
		return fmt.Errorf("error setting memory_size_gb in state: %s", err)
	}
	if err = d.Set("name", res.Name); err != nil {
		return fmt.Errorf("error setting name in state: %s", err)
	}
	if err = d.Set("alternative_location_id", res.AlternativeLocationId); err != nil {
		return fmt.Errorf("error setting alternative_location_id in state: %s", err)
	}
	if err = d.Set("authorized_network", res.AuthorizedNetwork); err != nil {
		return fmt.Errorf("error setting authorized_network in state: %s", err)
	}
	if err = d.Set("display_name", res.DisplayName); err != nil {
		return fmt.Errorf("error setting display_name in state: %s", err)
	}
	if err = d.Set("labels", res.Labels); err != nil {
		return fmt.Errorf("error setting labels in state: %s", err)
	}
	if err = d.Set("location_id", res.LocationId); err != nil {
		return fmt.Errorf("error setting location_id in state: %s", err)
	}
	if err = d.Set("project", res.Project); err != nil {
		return fmt.Errorf("error setting project in state: %s", err)
	}
	if err = d.Set("redis_configs", res.RedisConfigs); err != nil {
		return fmt.Errorf("error setting redis_configs in state: %s", err)
	}
	if err = d.Set("redis_version", res.RedisVersion); err != nil {
		return fmt.Errorf("error setting redis_version in state: %s", err)
	}
	if err = d.Set("region", res.Region); err != nil {
		return fmt.Errorf("error setting region in state: %s", err)
	}
	if err = d.Set("reserved_ip_range", res.ReservedIPRange); err != nil {
		return fmt.Errorf("error setting reserved_ip_range in state: %s", err)
	}
	if err = d.Set("tier", res.Tier); err != nil {
		return fmt.Errorf("error setting tier in state: %s", err)
	}
	if err = d.Set("current_location_id", res.CurrentLocationId); err != nil {
		return fmt.Errorf("error setting current_location_id in state: %s", err)
	}
	if err = d.Set("host", res.Host); err != nil {
		return fmt.Errorf("error setting host in state: %s", err)
	}
	if err = d.Set("port", res.Port); err != nil {
		return fmt.Errorf("error setting port in state: %s", err)
	}
	return nil
}

func resourceRedisInstanceUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	region, err := getRegion(d, config)
	if err != nil {
		return err
	}

	name := d.Get("name").(string)

	obj := &redis.Instance{
		MemorySizeGb:          dcl.Int64(int64(d.Get("memory_size_gb").(int))),
		Name:                  dcl.String(name),
		AlternativeLocationId: dcl.String(d.Get("alternative_location_id").(string)),
		AuthorizedNetwork:     dcl.String(d.Get("authorized_network").(string)),
		DisplayName:           dcl.String(d.Get("display_name").(string)),
		Labels:                expandMap(d.Get("labels"), d, config),
		LocationId:            dcl.String(d.Get("location_id").(string)),
		Project:               dcl.String(project),
		RedisConfigs:          expandMap(d.Get("redis_configs"), d, config),
		RedisVersion:          dcl.String(d.Get("redis_version").(string)),
		Region:                dcl.String(region),
		ReservedIPRange:       dcl.String(d.Get("reserved_ip_range").(string)),
		Tier:                  redis.InstanceTierEnumRef(d.Get("tier").(string)),
	}

	res, err := config.clientRedis.ApplyInstance(obj, UpdateDirective...)
	if err != nil {
		return fmt.Errorf("Error updating Instance: %s", err)
	}

	log.Printf("[DEBUG] Finished creating Instance %q: %#v", d.Id(), res)

	return resourceRedisInstanceRead(d, meta)
}

func resourceRedisInstanceDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	region, err := getRegion(d, config)
	if err != nil {
		return err
	}

	name := d.Get("name").(string)

	log.Printf("[DEBUG] Deleting Instance %q", d.Id())
	err = config.clientRedis.DeleteInstance(project, region, name)
	if err != nil {
		return fmt.Errorf("Error deleting Instance: %s", err)
	}

	log.Printf("[DEBUG] Finished deleting Instance %q: %#v", d.Id())
	return nil
}

func resourceRedisInstanceImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{
		"projects/(?P<project>[^/]+)/locations/(?P<region>[^/]+)/instances/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<region>[^/]+)/(?P<name>[^/]+)",
		"(?P<region>[^/]+)/(?P<name>[^/]+)",
		"(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVars(d, config, "projects/{{project}}/locations/{{region}}/instances/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}
