Pelago聚合卡产品文档
1.业务目标
整合卡渠道接口内化为pelago自己的卡业务接口，再对外提供卡业务相关接口供商户对接

业务关系图





2.业务架构图


3.业务资金流

3.1 商户大账户模式


3.2 支付网关代收模式

4.功能列表


5.解决方案
5.1 商户入驻
将商户添加在pelago card系统, 建立商户ID , 完成APIkey 的配置
第一阶段，先手动添加在数据库
第二阶段再做管理后台进行APIkey的配置
商户API请求要金额APIKey的验签

5.1.1 后端数据表
商户添加后，需要保存以下字段，为后续路由



5.2 商户路由配置




5.3 卡Bin路由管理
解释：
卡bin，表示卡号的前6位数
一个渠道可以提供多个卡bin


5.4 商户创建持卡人账户
5.4.1 接口目的
该接口提供给商户创建持卡人，商户抛送的持卡人数据先存在pelago后端表里，再根据商户调用的卡名称查找卡bin，基于卡bin查找渠道，再将卡信息抛送给渠道进行持卡人申请



5.4.2 时序图


5.4.3 商户创建持卡人的接口表

这里的数据来自于商户创建申请，由商户调Pelago card 创建申请人接口传入，pelago card 需要做存储


5.4.4 支持商户查询持卡人信息


商户传入参数：


Pelago 出参给商户：


5.5 商户确认账户开卡
开卡时点：当account_id对应的bin_code状态是active 且 持卡人applicator_status 为active 支持开卡操作
开卡pelago card需要生成 application fee订单，并走网关让C端用户进行支付
开卡目的：通过商户调用确认开卡，pelago向提供卡bin的渠道申请开卡，并把卡数据进行存储，支持商户查询开卡结果
开卡成功后，向商户推送开卡结果

5.5.1 开卡需要收取 “application fee"




支付成功后，网关回调pelago card 支付成功，pelago card调渠道开卡接口
按照渠道的开卡接口传参数

5.5.2 商户调开卡接口——商户入参数

商户调pelago开卡接口，需要传入字段


Pelago需要调商户对应的bin 所在渠道的接口进行开卡确认

5.5.3 Pelago card 需要存储开卡信息表


5.5.4 商户查询/Pelago card 推送开卡结果

商户查询时输入参数


Pelago card向渠道查询卡详细信息
当商户调pelago card查询开卡详情信息，pelago card需要调 该商户Bin所在的渠道查询卡详细接口进行查询， Pelago card 不可以存储卡详情，渠道回调数据同步返回给商户


5.6 商户拉取卡消费记录

5.6.1 商户拉取卡消费数据

商户输入参数


Pelago向渠道入参数
基于渠道要求

渠道向pelago出参数
pelago card对渠道回调的数据进行保存并返回给商户

5.6.2 Pelago card 后端消费交易表：
Pelago card需要存储卡交易相关数据字段


5.6.3 Pelago card 向商户出参卡消费数据


5.7 消费receipt获取积分
5.7.1 商户请求pelago生成receipt
商户可以调这个接口，pelago返回一个字符串，商户可以在自己的页面上渲染出二维码，供消费者下载保存，消费者可以把receipt 上传到polyflow 获取积分

商户入参：


5.7.2 Pelago receipt 字符串包含字段
Pelago card收到商户请求生成receipt请求，生成一个字符串并返回商户，由商户在自己的平台渲染成二维码
需要包括列表字段：

pelago出参数：


5.7.3 Polyflow向 pelago card系统校验receipt信息是否有效
Polyflow 验证签名通过
Polyflow验证信息包括：



5.8 商户调充值卡接口-网关模式


5.8.1 网关商户入驻逻辑
卡平台商户向网关单独入驻，获得商户ID，单独配置APIKey等
Pelago card存商户网关的ID和APIKey，用于后续向网关调起支付
当商户调充值接口，Pelago card向网关请求收单接口，并按照要求生成订单，等待支付结果回调

5.8.2 card 调网关充值时序图

5.8.3 商户调 Pelago Card 充值接口

商户给Pelago card top-up 传入参数：


5.8.4 Pelago card  业务订单表


5.8.5 卡-网关充值接口联调

Pelago card给 Pelago gateway 入参
当商户传入订单数据成功，pelago card 需要基于费用生成pelago card的业务订单


Pelago Payment Gateway 返回 checkout url


Pelago Payment Gateway 回调支付成功结果给卡业务


Pelago Payment Gateway 生成交易流水
基于pelago网关的收单逻辑生成

Pelago Payment Gateway  记账逻辑
记账记到商户账户


5.8.6 收单结算金额提现至渠道大账户

网关收单成功，金额结算至商户账户余额
管理员对商户余额走网关逻辑进行提现至渠道
商户可以设置3个提现白名单，白名单地址可以是渠道的地址（网关的逻辑）


5.9 商户调充值卡接口-大账户模式

5.9.1 大账户模式定义
商户走自己对持卡人的充值支付方式，但是当商户请求抛送到pelago card, 需要从商户的大账户余额扣除一笔费用

实际运营方式：
商户线下/线上给pelago 转账 ，线上转账通过crypto完成，需要pelago payment gateway 参与提供充值地址
Pelago给渠道转账

5.9.2 充值时序图


5.9.3 商户调充值接口

商户给Pelago card人参：


5.9.4 查询渠道大账户余额
Pelago card 按照渠道查询接口进行查询


5.9.5 商户充值大账户余额
商户入驻网关成功后，网关给商户开通payout业务
网关生成商户deposit钱包地址
商户往钱包充值
网关收到渠道回调充值成功通知
充值成功后商户账户增加余额
Admin 管理员提交提现到渠道大账户

5.9.6 网关增加 admin 管理员可提现商户余额
网关新功能


5.10 冻结卡

商户给Pelago Card 入参数

Pelago Card 给渠道入参数
基于渠道接口要求调冻结卡的接口，协助商户冻结卡片

冻结卡后续操作
后续冻结卡的操作取决于发卡渠道的控制

5.11 删除卡

商户给Pelago Card 入参数

Pelago Card 给渠道入参数
基于渠道接口要求调删除卡的接口，协助商户删除卡片


5.12 查询卡信息

商户给Pelago Card 入参数

Pelago Card 入参数给渠道


Pelago Card 返回给商户参数


5.13 费用管理-网关模式——C端用户付费

5.13.1  收费流程和时点

5.13.1.1 Application Fee/ 开卡费
开卡时候向持卡人收取，向网关抛送订单，参考 5.5.1 商户调开卡接口——商户入参数

5.13.1.2 Top-up fee / 充值费
商户调充值接口时候向持卡人收取，向网关抛送订单， 参考 5.8 商户调充值卡接口-网关模式

5.13.1.3 Monthly fee / 月费
概述：

时序图

月费订单表
参考： 5.8.4 Pelago card  业务订单表

5.13.2 基于商户ID配置收费类目表



5.14 费用管理-大账户模式——向商户大账户扣费

