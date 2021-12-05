# sap-api-integrations-business-partner-reads-customer
sap-api-integrations-business-partner-reads-customer は、外部システム(特にエッジコンピューティング環境)をSAPと統合することを目的に、SAP API で BP(ビジネスパートナ) - 得意先 データを取得するマイクロサービスです。    
sap-api-integrations-business-partner-reads-customer には、サンプルのAPI Json フォーマットが含まれています。   
sap-api-integrations-business-partner-reads-customer は、オンプレミス版である（＝クラウド版ではない）SAPS4HANA API の利用を前提としています。クラウド版APIを利用する場合は、ご注意ください。   
https://api.sap.com/api/OP_API_BUSINESS_PARTNER_SRV/overview   

## 動作環境  
sap-api-integrations-business-partner-reads-customer は、主にエッジコンピューティング環境における動作にフォーカスしています。  
使用する際は、事前に下記の通り エッジコンピューティングの動作環境（推奨/必須）を用意してください。  
・ エッジ Kubernetes （推奨）    
・ AION のリソース （推奨)    
・ OS: LinuxOS （必須）    
・ CPU: ARM/AMD/Intel（いずれか必須）    

## クラウド環境での利用
sap-api-integrations-business-partner-reads-customer は、外部システムがクラウド環境である場合にSAPと統合するときにおいても、利用可能なように設計されています。  

## 本レポジトリ が 対応する API サービス
sap-api-integrations-business-partner-reads-customer が対応する APIサービス は、次のものです。

* APIサービス概要説明 URL: https://api.sap.com/api/OP_API_BUSINESS_PARTNER_SRV/overview    
* APIサービス名(=baseURL): API_BUSINESS_PARTNER

## 本レポジトリ に 含まれる API名
sap-api-integrations-business-partner-reads-customer には、次の API をコールするためのリソースが含まれています。  

* A_BusinessPartnerRole（ビジネスパートナ - ロール）
* A_BusinessPartnerAddress（ビジネスパートナ - アドレス）
* A_CustomerSalesArea（ビジネスパートナ - 得意先販売エリア）
* A_CustomerCompany（ビジネスパートナ - 得意先会社コード）

## API への 値入力条件 の 初期値
sap-api-integrations-business-partner-reads-customer において、API への値入力条件の初期値は、入力ファイルレイアウトの種別毎に、次の通りとなっています。  

### SDC レイアウト

* inoutSDC.BusinessPartner.BusinessPartner（ビジネスパートナ）
* inoutSDC.BusinessPartner.BusinessPartnerRole（ビジネスパートナロール）
* inoutSDC.BusinessPartner.Address.AddressID（アドレスID）
* inoutSDC.BusinessPartner.SalesArea.SalesOrganization（販売組織）
* inoutSDC.BusinessPartner.SalesArea.DistributionChannel（流通チャネル）
* inoutSDC.BusinessPartner.SalesArea.Division（部門）
* inoutSDC.Customer（得意先コード ※ビジネスパートナの販売エリア・会社コード関連のAPIをコールするときにビジネスパートナではなく得意先コードとしての項目値が必要です。通常は、ビジネスパートナの値＝得意先コードの値、となります）
* inoutSDC.BusinessPartner.Company.CompanyCode（会社コード）
