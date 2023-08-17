package data

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"kubecit/ent/cloudhost"
	"kubecit/internal/biz"
	"reflect"
	"strings"
)

type cloudHostRepo struct {
	data *Data
	log  *log.Helper
}

// NewCloudHostRepo .
func NewCloudHostRepo(data *Data, logger log.Logger) biz.CloudHostRepo {
	return &cloudHostRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// Get
func (c *cloudHostRepo) Get(ctx context.Context, id string) (*biz.CloudHost, error) {
	h, err := c.data.db.CloudHost.Query().Where(cloudhost.UUIDEQ(id)).Only(ctx)
	if err != nil {
		return nil, err
	}
	res := &biz.CloudHost{}
	err = ConvertType(h, res)
	res.IPV4AddressPrivate = strings.Split(h.Ipv4AddressPrivate, ",")
	res.IPV6AddressPublic = strings.Split(h.Ipv6AddressPublic, ",")
	res.IPV4AddressPublic = strings.Split(h.Ipv4AddressPublic, ",")
	res.IPV6AddressPrivate = strings.Split(h.Ipv6AddressPrivate, ",")
	res.SecurityGroups = strings.Split(h.SecurityGroups, ",")
	return res, err
}

// Create
func (c *cloudHostRepo) Create(ctx context.Context, host *biz.CloudHost) (*biz.CloudHost, error) {

	h, err := c.data.db.CloudHost.Create().SetCreatedTime(host.CreatedTime).SetExpiredTime(host.ExpiredTime).
		SetBillType(host.BillType).SetInstanceType(host.InstanceType).SetUUID(host.UUID).
		SetChargeType(host.ChargeType).SetCPU(host.CPU).SetMemory(host.Memory).SetImageName(host.ImageName).
		SetIpv4AddressPrivate(strings.Join(host.IPV4AddressPrivate, ",")).SetIpv4AddressPublic(strings.Join(host.IPV4AddressPublic, ",")).
		SetIpv6AddressPrivate(strings.Join(host.IPV6AddressPrivate, ",")).SetIpv6AddressPublic(strings.Join(host.IPV6AddressPublic, ",")).
		SetState(host.State).SetInstanceName(host.InstanceName).SetOsType(host.OSType).
		SetManufacturer(host.Manufacturer).SetZone(host.Zone).SetSecurityGroups(strings.Join(host.SecurityGroups, ",")).Save(ctx)
	if err != nil {
		return nil, err
	}

	res := &biz.CloudHost{}
	err = ConvertType(h, res)

	return res, err
}

// List
func (c *cloudHostRepo) List(ctx context.Context) ([]*biz.CloudHost, error) {
	hosts, err := c.data.db.CloudHost.Query().All(ctx)
	if err != nil {
		return nil, err
	}

	resList := make([]*biz.CloudHost, 0)
	for _, host := range hosts {
		res := &biz.CloudHost{}
		err = ConvertType(host, res)
		if err != nil {
			return nil, err
		}
		res.IPV4AddressPrivate = strings.Split(host.Ipv4AddressPrivate, ",")
		res.IPV6AddressPublic = strings.Split(host.Ipv6AddressPublic, ",")
		res.IPV4AddressPublic = strings.Split(host.Ipv4AddressPublic, ",")
		res.IPV6AddressPrivate = strings.Split(host.Ipv6AddressPrivate, ",")
		res.SecurityGroups = strings.Split(host.SecurityGroups, ",")
		resList = append(resList, res)
	}
	return resList, err
}

// Delete
func (c *cloudHostRepo) Delete(ctx context.Context, id string) (*biz.CloudHost, error) {
	rows, err := c.data.db.CloudHost.Update().SetIsActive(false).Where(cloudhost.UUID(id)).Save(ctx)
	if err != nil || rows != 1 {
		fmt.Println(err)
		return nil, err
	}
	host, err := c.data.db.CloudHost.Query().Where(cloudhost.UUID(id)).First(ctx)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	res := &biz.CloudHost{}
	err = ConvertType(host, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *cloudHostRepo) Update(ctx context.Context, id string, host *biz.CloudHost) (*biz.CloudHost, error) {
	query := c.data.db.CloudHost.Update().Where(cloudhost.UUIDEQ(id))
	if host.IsActive != false {
		query.SetIsActive(host.IsActive)
	}
	if host.UUID != "" {
		query.SetUUID(host.UUID)
	}
	if host.CPU != 0 {
		query.SetCPU(host.CPU)
	}
	if host.State != "" {
		query.SetState(host.State)
	}
	if len(host.IPV6AddressPrivate) != 0 {
		query.SetIpv6AddressPrivate(strings.Join(host.IPV6AddressPrivate, ","))
	}
	if len(host.IPV6AddressPublic) != 0 {
		query.SetIpv6AddressPublic(strings.Join(host.IPV6AddressPublic, ","))
	}
	if len(host.IPV4AddressPrivate) != 0 {
		query.SetIpv4AddressPrivate(strings.Join(host.IPV4AddressPrivate, ","))
	}
	if len(host.IPV4AddressPublic) != 0 {
		query.SetIpv4AddressPublic(strings.Join(host.IPV4AddressPublic, ","))
	}
	if len(host.SecurityGroups) != 0 {
		query.SetSecurityGroups(strings.Join(host.SecurityGroups, ","))
	}
	if host.Memory != 0 {
		query.SetMemory(host.Memory)
	}
	if host.InstanceType != "" {
		query.SetInstanceType(host.InstanceType)
	}
	if host.InstanceName != "" {
		query.SetInstanceName(host.InstanceName)
	}
	if host.ChargeType != "" {
		query.SetChargeType(host.ChargeType)
	}
	if host.Zone != "" {
		query.SetZone(host.Zone)
	}
	if host.BillType != "" {
		query.SetBillType(host.BillType)
	}
	if host.OSType != "" {
		query.SetOsType(host.OSType)
	}
	if host.Manufacturer != "" {
		query.SetManufacturer(host.Manufacturer)
	}
	if host.ImageName != "" {
		query.SetImageName(host.ImageName)
	}

	_, err := query.Save(ctx)
	if err != nil {
		fmt.Println("update error: ", err)
		return nil, err
	}
	res, err := c.Get(ctx, host.UUID)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Sync TODO
func (c *cloudHostRepo) Sync(ctx context.Context) (bool, error) {
	return true, nil
}

func ConvertType(src, dest interface{}) error {
	dv, sv := reflect.ValueOf(dest), reflect.ValueOf(src)
	if dv.Kind() != reflect.Ptr {
		return fmt.Errorf("need pointer")
	}
	if sv.Kind() != reflect.Ptr {
		return fmt.Errorf("need pointer")
	}
	dv = dv.Elem()
	sv = sv.Elem()

	if !dv.CanAddr() {
		return fmt.Errorf("can't write to dest")
	}

	st := sv.Type()
	dt := dv.Type()
	for i := 0; i < st.NumField(); i++ {
		fieldName := st.Field(i).Name
		sField, _ := st.FieldByName(fieldName)
		sFieldValue := sv.FieldByName(fieldName)
		dField, _ := dt.FieldByName(fieldName)
		dFieldValue := dv.FieldByName(fieldName)
		if sField.Name == dField.Name && dFieldValue.CanSet() && sField.Type == dField.Type {
			dFieldValue.Set(sFieldValue)
		} else if sField.Name == dField.Name && sFieldValue.CanConvert(dFieldValue.Type()) {
			dFieldValue.Set(sFieldValue.Convert(dFieldValue.Type()))
		}
	}
	return nil
}
