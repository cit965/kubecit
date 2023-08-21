package data

import (
	"context"
	"fmt"
	"reflect"

	"github.com/go-kratos/kratos/v2/log"

	"kubecit/ent/cloudhost"
	"kubecit/internal/biz"
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
	h, err := c.data.db.CloudHost.Query().Where(cloudhost.InstanceIdEQ(id)).First(ctx)
	if err != nil {
		return nil, err
	}
	res := &biz.CloudHost{}
	err = ConvertType(h, res)
	return res, err
}

// Create
func (c *cloudHostRepo) Create(ctx context.Context, host *biz.CloudHost) (*biz.CloudHost, error) {
	h, err := c.data.db.CloudHost.Create().SetVpcId(host.VpcId).SetSubnetId(host.SubnetId).SetInstanceState(host.InstanceState).
		SetInstanceId(host.InstanceId).SetInstanceName(host.InstanceName).SetCPU(host.CPU).
		SetMemory(host.Memory).SetCreatedTime(host.CreatedTime).SetInstanceType(host.InstanceType).
		SetEniLimit(host.EniLimit).SetEnilpLimit(host.EnilpLimit).SetInstanceEniCount(host.InstanceEniCount).Save(ctx)
	if err != nil {
		fmt.Println(err)
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
		resList = append(resList, res)
	}
	return resList, err
}

// Delete
func (c *cloudHostRepo) Delete(ctx context.Context, id string) (*biz.CloudHost, error) {
	host, err := c.data.db.CloudHost.Query().Where(cloudhost.InstanceIdEQ(id)).First(ctx)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	rows, err := c.data.db.CloudHost.Delete().Where(cloudhost.InstanceIdEQ(id)).Exec(ctx)
	if err != nil || rows != 1 {
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
	query := c.data.db.CloudHost.Update().Where(cloudhost.InstanceIdEQ(id))
	if host.VpcId != "" {
		query.SetVpcId(host.VpcId)
	}
	if host.SubnetId != "" {
		query.SetSubnetId(host.SubnetId)
	}
	if host.InstanceId != "" {
		query.SetInstanceId(host.InstanceId)
	}
	if host.InstanceName != "" {
		query.SetInstanceName(host.InstanceName)
	}
	if host.InstanceState != "" {
		query.SetInstanceState(host.InstanceState)
	}
	if host.CPU != 0 {
		query.SetCPU(host.CPU)
	}
	if host.Memory != 0 {
		query.SetMemory(host.Memory)
	}
	if host.CreatedTime != "" {
		query.SetCreatedTime(host.CreatedTime)
	}
	if host.InstanceType != "" {
		query.SetInstanceType(host.InstanceType)
	}
	if host.EniLimit != 0 {
		query.SetEniLimit(host.EniLimit)
	}
	if host.EnilpLimit != 0 {
		query.SetEnilpLimit(host.EnilpLimit)
	}
	if host.InstanceEniCount != 0 {
		query.SetInstanceEniCount(host.InstanceEniCount)
	}

	_, err := query.Save(ctx)
	if err != nil {
		fmt.Println("update error: ", err)
		return nil, err
	}
	res, err := c.Get(ctx, id)
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
