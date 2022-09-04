# sap-api-integrations-business-partner-creates-customer
sap-api-integrations-business-partner-creates-customer は、外部システム(特にエッジコンピューティング環境)をSAPと統合することを目的に、SAP API で BP(ビジネスパートナ) - 得意先 データ を登録するマイクロサービスです。  
sap-api-integrations-business-partner-creates-customer には、サンプルのAPI Json フォーマットが含まれています。  
sap-api-integrations-business-partner-creates-customer は、オンプレミス版である（＝クラウド版ではない）SAPS4HANA API の利用を前提としています。クラウド版APIを利用する場合は、ご注意ください。  
https://api.sap.com/api/OP_API_BUSINESS_PARTNER_SRV/overview  

## 動作環境  
sap-api-integrations-business-partner-creates-customer は、主にエッジコンピューティング環境における動作にフォーカスしています。  
使用する際は、事前に下記の通り エッジコンピューティングの動作環境（推奨/必須）を用意してください。  
・ エッジ Kubernetes （推奨）   
・ AION のリソース （推奨)   
・ OS: LinuxOS （必須）   
・ CPU: ARM/AMD/Intel（いずれか必須）  

## クラウド環境での利用
sap-api-integrations-business-partner-creates-customer は、外部システムがクラウド環境である場合にSAPと統合するときにおいても、利用可能なように設計されています。


## 本レポジトリ が 対応する API サービス
sap-api-integrations-business-partner-creates-customer が対応する APIサービス は、次のものです。

* APIサービス概要説明 URL: https://api.sap.com/api/OP_API_BUSINESS_PARTNER_SRV/overview  
* APIサービス名(=baseURL): API_BUSINESS_PARTNER

## 本レポジトリ に 含まれる API名
sap-api-integrations-business-partner-creates-customer には、次の API をコールするためのリソースが含まれています。  

* A_BusinessPartner（ビジネスパートナ - 一般）
* A_BusinessPartnerRole（ビジネスパートナ - ロール）
* A_BusinessPartnerAddress（ビジネスパートナ - アドレス）
* A_BusinessPartnerBank（ビジネスパートナ - 銀行）
* A_Customer（ビジネスパートナ - 得意先）
* A_CustomerSalesArea（ビジネスパートナ - 得意先販売エリア）
* A_CustSalesPartnerFunc（ビジネスパートナ - 得意先機能）
* A_CustomerCompany（ビジネスパートナ - 得意先会社コード）

## SAP API Bussiness Hub の API の選択的コール

Latona および AION の SAP 関連リソースでは、Inputs フォルダ下の sample.json の accepter に登録したいデータの種別（＝APIの種別）を入力し、指定することができます。  
なお、同 accepter にAll(もしくは空白)の値を入力することで、全データ（＝全APIの種別）をまとめて登録することができます。  

* sample.jsonの記載例(1)  

accepter において 下記の例のように、データの種別（＝APIの種別）を指定します。  
ここでは、"General" が指定されています。    
  
```
"api_schema": "SAPBusinessPartnerCreatesCustomer",
"accepter": ["General"],
"business_partner_code": "1000140",
"deleted": false
```
  
* 全データを取得する際のsample.jsonの記載例(2)  

全データを取得する場合、sample.json は以下のように記載します。  

```
"api_schema": "SAPBusinessPartnerCreatesCustomer",
"accepter": ["All"],
"business_partner_code": "1000140",
"deleted": false
```

## 指定されたデータ種別のコール

accepter における データ種別 の指定に基づいて SAP_API_Caller 内の caller.go で API がコールされます。  
caller.go の func() 毎 の 以下の箇所が、指定された API をコールするソースコードです。  

```
func (c *SAPAPICaller) AsyncPostBPC(
	general *requests.General,
	role *requests.Role,
	address *requests.Address,
	bank *requests.Bank,
	customer *requests.Customer,
	salesArea *requests.SalesArea,
	partnerFunction *requests.PartnerFunction,
	company *requests.Company,
	accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "General":
			func() {
				c.General(general)
				wg.Done()
			}()
		case "Role":
			func() {
				c.Role(role)
				wg.Done()
			}()
		case "Bank":
			func() {
				c.Bank(bank)
				wg.Done()
			}()
		case "Address":
			func() {
				c.Address(address)
				wg.Done()
			}()
		case "Customer":
			func() {
				c.Customer(customer)
				wg.Done()
			}()
		case "SalesArea":
			func() {
				c.SalesArea(salesArea)
				wg.Done()
			}()
		case "PartnerFunction":
			func() {
				c.PartnerFunction(partnerFunction)
				wg.Done()
			}()
		case "Company":
			func() {
				c.Company(company)
				wg.Done()
			}()
		default:
			wg.Done()
		}
	}

	wg.Wait()
}
```

## Output  
本マイクロサービスでは、[golang-logging-library-for-sap](https://github.com/latonaio/golang-logging-library-for-sap) により、以下のようなデータがJSON形式で出力されます。  
以下の sample.json の例は、SAP BP(ビジネスパートナ) - 得意先 データ が登録された結果の JSON の例です。  
以下の項目のうち、"XXXXX" ～ "XXXXX" は、/SAP_API_Output_Formatter/type.go 内 の Type General {} による出力結果です。"cursor" ～ "time"は、golang-logging-library による 定型フォーマットの出力結果です。  

```
	"cursor": "/Users/latona2/bitbucket/sap-api-integrations-business-partner-creates-customer/SAP_API_Caller/caller.go#L50",
	"function": "sap-api-integrations-business-partner-creates-customer/SAP_API_Caller.(*SAPAPICaller).Header",
	"level": "INFO",
	"message": "[{XXXXXXXXXXXXXXXXXXXXXXXXXXXXX}]",
	"time": "2021-12-11T15:33:00.054455+09:00"
```
