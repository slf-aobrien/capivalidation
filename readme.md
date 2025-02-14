# Compass API Validation

## What is it
This is a quality assurance tool built in GoLang.  I downloaded the swagger definition file from located [here](https://dev-corp-wk-ca1-k8s.sunlifecorp.com/sit-compass-integration-ns/compassintegrationapp/swagger-ui/index.html#) it is located at the root level called `capi_swagger.json`

I then used swagger-go to generate the client code.  This was a large trial and error but I ended up createing a client directory initializing it as a module and then executing the command something like:  
`swagger generate client -f ../capi_swagger.json -t ./` <- Probably not perfect but it's close.

## How it works
After that it was a series of trial and error with using copilot :|  

The compass api's are... strange at times and require a lot of finessing with the request.  I ended up downloading the community edition of [burp suite](https://portswigger.net/burp/communitydownload) to interrogate and modify the requests being sent.  It led to a lot of difficult but interesting code mostly in the cmd.go `retrieveMembersForMemberGroup` function.  Forcing a response type from the server was done very unintuitively:
```
	transport.Consumers = map[string]runtime.Consumer{
		"application/json": runtime.JSONConsumer(),
	}

	// Set the Accept header explicitly
	transport.DefaultAuthentication = runtime.ClientAuthInfoWriterFunc(
        func(req runtime.ClientRequest, reg strfmt.Registry) error {
		    req.SetHeaderParam("Accept", "application/json")
		    return nil
	    }
    )
```
and more peculiarly the data that get's sent back from this call doesn't parse well with the generated client!  This secion failed to parse:
```
"person":{
    "firstName":"Mbr1",
    "lastName":"871675-1",
    "birthDate":"12/12/1981",
    "birthDateAsDate":376981200000
}
```
What was going on was that the `birthDateAsDate` field in the JSON response is being sent as a number (timestamp), but your Go struct expects it to be a strfmt.DateTime type, which is a string representation of a date-time.

To resolve this, I had to change the type of the birthDateAsDate field in your PersonDTORes struct to int64 to match the format of the data being received, and this method  
`func (m *PersonDTORes) validateBirthDateAsDate(formats strfmt.Registry) error {`
had to be modifed to remove the ability to parse the data as though it was a string.

This is **important** for you to understand because if the client is regenerated you will have this problem again!

Runnint the client at this particular moment returns the following output:
```
SM58@USMA018TCO client % go run cmd.go
2025/02/14 14:19:01 DCX - Scenario 1
2025/02/14 14:19:01 Policy Number 871675
2025/02/14 14:19:01 Current Date: 02/14/2025
2025/02/14 14:19:01 Member Group Description: COBRA Employees-Dental or no Dental:2750676
2025/02/14 14:19:02 Member Count: 1
2025/02/14 14:19:02     Member: "Person 5", "Tester": "001931155"
2025/02/14 14:19:02 Member Group Description: All Employees-Prepaid:2750677
2025/02/14 14:19:03 Member Count: 5
2025/02/14 14:19:03     Member: "Person 10", "Tester": "001931160"
2025/02/14 14:19:03     Member: "Person 4", "Tester": "001931154"
2025/02/14 14:19:03     Member: "Person 7", "Tester": "001931157"
2025/02/14 14:19:03     Member: "Person 8", "Tester": "001931158"
2025/02/14 14:19:03     Member: "Person3", "Tester": "001931153"
2025/02/14 14:19:03 Member Group Description: COBRA Employees-Prepaid:2750679
2025/02/14 14:19:03 Member Count: 0
2025/02/14 14:19:03 Member Group Description: All Employees-Dental or no Dental:2750678
2025/02/14 14:19:03 Member Count: 10
2025/02/14 14:19:03     Member: "Mbr1", "871675-1": "001923911"
2025/02/14 14:19:03     Member: "Termed", "Benefit": "001964589"
2025/02/14 14:19:03     Member: "Term EE", "Current Ben": "001964887"
2025/02/14 14:19:03     Member: "Benefit", "Effective": "001964266"
2025/02/14 14:19:03     Member: "Termed", "Numb2": "001964593"
2025/02/14 14:19:03     Member: "Termed", "Reinstate": "001964594"
2025/02/14 14:19:03     Member: "Dental", "Spouse": "001964888"
2025/02/14 14:19:03     Member: "Person 6", "Tester": "001931156"
2025/02/14 14:19:03     Member: "Person 9", "Tester": "001931159"
2025/02/14 14:19:03     Member: "Person2", "Tester": "001931152"
2025/02/14 14:19:03 Elapsed Time: 2.349669ms
```

## Environments
Since I generated the client from the sit/qar version of the file that is the default environment the client runs against.  If you want to change the environment that the client will communicate with there are two ways the right way and the wrong way.

The right way is to create a new client like this:
```
	// Create a new transport with the updated host and basepath
	transport = httptransport.New("dev-corp-wk-ca1-k8s.sunlifecorp.com", 
                                "/dev-compass-integration-ns", 
                                apiclient.DefaultSchemes)
	client := apiclient.New(transport, strfmt.Default)

	resp, err := client.System.GetDomainUsingGET(params)
```

The wrong way is to edit the api_documentation code:
```
const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file

	DefaultHost string = "dev-corp-wk-ca1-k8s.sunlifecorp.com"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file

	DefaultBasePath string = "/sit-compass-integration-ns"
	//DefaultBasePath string = "/dev-compass-integration-ns"
)
```

## Docker 
I did install swagger-go on my laptop but my laptop, for some reason, wouldn't run the swagger library.  I don't know why.  So I used docker and vscode's [devcontainer](https://code.visualstudio.com/docs/devcontainers/containers) to make it work.  If you don't know how to get devcontainer working check it out, it's pretty easy and very powerful.

Also once I got the client code generated, my docker container refuesed to talk to compass, I don't know  why - I can't blame it - but still it was odd.  I ended up going back to developing on my local lapotop environment at that point.  I would recommend that you do that as well.

Cheers!

-Aaron