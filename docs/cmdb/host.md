## CMDB云主机字段
后期会做多云的资产管理，各云商提供的API中返回字段略有不同，且部分字段也可能并没有实际使用含义。需要进行统一，仅保留实用字段。
注：各厂商含义相同但叫法不同的字段也需进行统一。

[腾讯云API文档](https://cloud.tencent.com/document/api/213/15753)

### CloudHost
| Field              | Type      | Examples                                |
|--------------------|-----------|-----------------------------------------|
| UUID               | string    | e85f1388-0422-410d-8e50-bef540e78c18    |
| State              | string    | RUNNING                                 |
| IPV6AddressPrivate | []string  | 2001:0db8:86a3:08d3:1319:8a2e:0370:7344 |
| IPV4AddressPrivate | []string  | 172.10.11.82                            |
| IPV6AddressPublic  | []string  | 2001:0db8:86a3:08d3:1319:8a2e:0370:7344 |
| IPV4AddressPublic  | []string  | 123.207.11.190                          |
| Memory             | int       | 32768(MB)                               |
| CPU                | int       | 8                                       |
| CreatedTime        | time.Time | 2020-09-22T00:00:00+00:00               |
| ExpiredTime        | time.Time | 2020-09-22T00:00:00+00:00               |
| InstanceName       | string    | ins-xlsyru2j                            |
| ImageName          | string    | Centos7.9_enhanced                      |
| OSType             | string    | Centos7.9                               |
| Manufacturer       | string    | TencentCloud                            |
| Zone               | string    | ap-guangzhou-1                          |
| SystemDisk         | *Disk     |                                         |
| DataDisks          | []*Disk   |                                         |
| SecurityGroups     | []string  | VPC-PRD,VPC-Serivce                     |
| BillType           | string    | NOTIFY_AND_MANUAL_RENEW                 |
| ChargeType         | string    | POSTPAID_BY_HOUR                        |
| IsActive           | bool      | True                                    |
| InstanceType       | string    | S2.SMALL2                               |

### Disk
| Field    | Type   | Examples  |
|----------|--------|-----------|
| DiskType | String | CLOUD_SSD |
| DiskSize | int    | 100       |

## 云主机同步功能接口
```go
type CloudSyncer interface {
	Get(id string) (*CloudHost,error)
	Create(host *CloudHost) (*CloudHost,error)
	List() ([]*CloudHost,error)
	// not real delete, just add flag
	Delete(id string) (*CloudHost,error)
	Update(id string, host *Host) (*CloudHost,error)
}
```