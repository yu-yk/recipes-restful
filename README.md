# recipes-restful

```
├── api                  the api server backed by gin
│   ├── handler          handler and middleware
│   ├── route            url mapping
│   └── server           api server struct
├── recipe               recipe package
│   ├── model            recipe struct mapping with db
│   ├── mysql_repo       implementation of mysql storage
│   ├── service_test     service layer logic test
│   ├── service          service layer logic
│   └── test_repo        implementation of test storage 
└── main                 entry point
```