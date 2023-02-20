# Go: Vandar Payment Gateway

This app is a middle-ware app for Vandar.io Internet Payment Gateway. you cabt set your gateway and port in config.yaml and set api key in payment/vandar/const.go


# Step 1. Get Token
- get token from vandar to redirect paymnent base_url/get-token
```json
{
	"amount" :10000,
	"mobile" : "09353917307"
}
```

- Response 
```json
{
  "status": 1,
  "token": "MU4MBIYGUDZNVUS",
  "errors": null
}
```
## Step 2. Redirect

- in first step get token from vandar to redirect paymnent base_url/redirect?token=
```
http://vandagateway.local:8080/redirect?token=MU4MBIYGUDZNVUS
```

## Step 2. Callback

- in first step get token from vandar to redirect paymnent base_url/callback
Failed Callback:
```
http://vandagateway.local:8080/callback?token=MU4MBIYGUDZNVUS&payment_status=FAILED
```
OK Callback
```
http://vandagateway.local:8080/callback?token=2YKWWXU9PLUHICC&payment_status=OK
```
```
{
  "status": 1,
  "amount": "10000.00",
  "transId": 167688571489,
  "refnumber": "GmshtyjwKSsQnapH6427NHGkBHPu1ch2ai7B2n8o2D",
  "trackingCode": "543008",
  "factorNumber": null,
  "mobile": "09353917307",
  "description": null,
  "cardNumber": "504172******5053",
  "CID": "3AF51353AD57716389FF1781E352015E14FF43A88FD737C4C890B933316CD52D",
  "createdAt": "2023-02-20 21:42:38",
  "paymentDate": "2023-02-20 21:42:52",
  "code": 1
}
```
## Step 3. Transaction Detail

- get transaction detail  base_url/transaction-detail

```
{
	"status": 1,
	"amount": "10000.00",
	"transId": 167688571489,
	"refnumber": "GmshtyjwKSsQnapH6427NHGkBHPu1ch2ai7B2n8o2D",
	"trackingCode": "543008",
	"factorNumber": null,
	"mobile": "09353917307",
	"description": null,
	"cardNumber": "504172******5053",
	"CID": "3AF51353AD57716389FF1781E352015E14FF43A88FD737C4C890B933316CD52D",
	"createdAt": "2023-02-20 21:42:38",
	"paymentDate": "2023-02-20 21:46:09",
	"code": 1
}
```

## Step 4. Verify Transaction

- verify transaction paymnent base_url/verify
```json
{
	"token" :"2YKWWXU9PLUHICC"
}
```
```json
{
	"status": 1,
	"errors": null,
	"amount": "10000.00",
	"realAmount": 10000,
	"wage": "0",
	"transId": 167688571489,
	"factorNumber": "",
	"mobile": "09353917307",
	"description": "",
	"cardNumber": "504172******5053",
	"paymentDate": "2023-02-20 21:46:44",
	"cid": "3AF51353AD57716389FF1781E352015E14FF43A88FD737C4C890B933316CD52D",
	"message": "ok"
}
```



