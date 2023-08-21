## CMDB云主机字段
后期会做多云的资产管理，各云商提供的API中返回字段略有不同，且部分字段也可能并没有实际使用含义。需要进行统一，仅保留实用字段。
注：各厂商含义相同但叫法不同的字段也需进行统一。

[腾讯云API文档](https://cloud.tencent.com/document/api/213/15753
[腾讯云VPC文档](https://console.cloud.tencent.com/api/explorer?Product=vpc&Version=2017-03-12&Action=DescribeVpcInstances)
- 当前主机信息根据所在Region和VPC进行同步，需传入accessKey、secretKey、region、vpcId进行同步。参考以上文档

#### api使用方法请参照swagger api文档
```azure
/cmdb/instance/{instanceId}                              GET
/cmdb/instance                                           POST
/cmdb/instances                                          GET
/cmdb/instance/{instanceId}                              DELETE
/cmdb/instance/{instanceId}                              PUT
/cmdb/sync/tencent                                       POST
```


### CloudHost(待丰富)
| Field            | Type    | Meaning              |
|------------------|---------|----------------------|
| VpcId            | String  | VPC实例ID。             |
| SubnetId         | String  | 子网实例ID。              |
| InstanceId       | String  | 云主机实例ID              |
| InstanceName     | String  | 云主机名称。               |
| InstanceState    | String  | 云主机状态。               |
| CPU              | Integer | 实例的CPU核数，单位：核。       |
| Memory           | Integer | 实例内存容量，单位：GB。        |
| CreatedTime      | String  | 创建时间。                |
| InstanceType     | String  | 实例机型。                |
| EniLimit         | Integer | 实例弹性网卡配额（包含主网卡）。     |
| EniIpLimit       | Integer | 实例弹性网卡内网IP配额（包含主网卡）。 |
| InstanceEniCount | Integer | 实例已绑定弹性网卡的个数（包含主网卡）。 |


## 云主机同步功能接口
```go
type CloudSyncer interface {
	Get(id string) (*CloudHost,error)
	Create(host *CloudHost) (*CloudHost,error)
	List() ([]*CloudHost,error)
	Delete(id string) (*CloudHost,error)
	Update(id string, host *Host) (*CloudHost,error)
}
```

#### TODO
- 丰富测试用例
- 同步VPC列表以及地域列表
- 云商主机同步方法抽象为接口