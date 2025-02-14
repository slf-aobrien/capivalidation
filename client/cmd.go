package main

import (
	apiclient "client/client"
	"client/client/member"
	"client/client/policy"
	"client/client/system"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

func main() {
	retrieveDomains("ZZSTCD")

	// startTime := time.Now()
	// retrieveMemberGroupsForPolicy("871675", "8", "sw57")
	// endTime := time.Now()
	// log.Print("Elapsed Time: ", (endTime.Sub(startTime) / 1000))
	//etrieveMembersForMemberGroup("871675", "2750678", "02/15/2025")
}
func retrieveMemberGroupsForPolicy(policyNumber string, flag string, user string) {
	//https://dev-corp-wk-ca1-k8s.sunlifecorp.com/sit-compass-integration-ns/compassintegrationapp/
	// 		policy/policyInfo?flag=8&policyNumber=871675&requestApplication=sw57&requestUser=sw57
	params := policy.NewGetPolicyUsingGETParams()
	params.PolicyNumber = &policyNumber
	params.Flag = &flag
	params.RequestApplication = user
	params.RequestUser = &user

	resp, err := apiclient.Default.Policy.GetPolicyUsingGET(params)
	if err != nil {
		log.Print("Error was not null!?")
		log.Print(err)
		return
	}

	log.Print(resp.Payload.Payload.Name)
	log.Print("Policy Number ", resp.Payload.Payload.PolicyNumber)

	currentDate := time.Now().Format("01/02/2006")
	log.Print("Current Date: ", currentDate)
	for _, mgb := range resp.Payload.Payload.MemberGroupList {
		log.Printf("Member Group Description: %v:%v", mgb.Description, mgb.Key)
		retrieveMembersForMemberGroup(policyNumber, mgb.Key, currentDate)
	}

}

func retrieveMembersForMemberGroup(policyNumber string, memberGroupKey string, effectiveDate string) {

	// Load Burp Suite CA certificate
	caCertDER, err := os.ReadFile("../burpCert.der")
	if err != nil {
		log.Fatal(err)
	}

	// Convert DER to PEM format
	caCertPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "CERTIFICATE",
		Bytes: caCertDER,
	})

	caCertPool := x509.NewCertPool()
	if !caCertPool.AppendCertsFromPEM(caCertPEM) {
		log.Fatal("Failed to append Burp Suite CA certificate to pool")
	}

	proxyURL, err := url.Parse("http://127.0.0.1:8080")
	if err != nil {
		log.Fatal("Failed to parse proxy URL:", err)
	}

	transport := httptransport.New(apiclient.DefaultHost, apiclient.DefaultBasePath, apiclient.DefaultSchemes)
	transport.Transport = &http.Transport{
		Proxy: http.ProxyURL(proxyURL),
		TLSClientConfig: &tls.Config{
			RootCAs: caCertPool,
		},
	}
	//transport.DefaultAuthentication = httptransport.BearerToken("your-token-here")
	//transport.Consumers["application/aeb.cas.member.list.v1+json"] = runtime.XMLConsumer() // Add custom consumer
	transport.Consumers = map[string]runtime.Consumer{
		"application/json": runtime.JSONConsumer(),
	}

	// Set the Accept header explicitly
	transport.DefaultAuthentication = runtime.ClientAuthInfoWriterFunc(func(req runtime.ClientRequest, reg strfmt.Registry) error {
		req.SetHeaderParam("Accept", "application/json")
		return nil
	})

	client := apiclient.New(transport, strfmt.Default)

	//https://dev-corp-wk-ca1-k8s.sunlifecorp.com/sit-compass-integration-ns/compassintegrationapp/
	// member/list?effectiveDate=02%2F15%2F2025&memberGroupKeys=2750678&policyNumber=871675
	params := member.NewGetMemberListUsingGETParams()
	params.PolicyNumber = &policyNumber
	params.MemberGroupKeys = &memberGroupKey
	params.EffectiveDate = &effectiveDate

	resp, err := client.Member.GetMemberListUsingGET(params)
	if err != nil {
		log.Print("Error was not null!?")
		log.Print(err)
		return
	}
	log.Print("Member Count: ", len(resp.Payload.Payload.MemberList))
	for _, member := range resp.Payload.Payload.MemberList {
		log.Printf("\tMember: %#v, %#v: %#v\n", member.Person.FirstName, member.Person.LastName, member.MemberNumber)
	}
}

func retrieveDomains(domain string) {
	// Create a custom transport that sets the Accept header to application/xml
	transport := httptransport.New(apiclient.DefaultHost, apiclient.DefaultBasePath, apiclient.DefaultSchemes)

	params := system.NewGetDomainUsingGETParams()
	domainName := "ZZSTCD"
	user := "sw57"
	format := "DETAIL"
	level := "ALL"
	params.DomainName = &domainName
	params.RequestUser = &user
	params.RequestApplication = &user
	params.VarianceFormat = &format
	params.VarianceLevel = &level

	// Create a new transport with the updated host and basepath
	transport = httptransport.New("dev-corp-wk-ca1-k8s.sunlifecorp.com", "/dev-compass-integration-ns", apiclient.DefaultSchemes)
	client := apiclient.New(transport, strfmt.Default)

	//resp, err := client.System.GetHoldaysDatesUsingGET(nil)
	resp, err := client.System.GetDomainUsingGET(params)
	if err != nil {
		log.Print("Error was not null!?")
		log.Print(err)
		log.Printf("transport.Consumers: %v\n", transport.Consumers)
		return
	}
	log.Printf("%#v\n", resp.Payload)
	log.Print(resp.Payload.Payload.Name)
	for _, value := range resp.Payload.Payload.ValueList {
		log.Printf("%#v\n", value)
	}
}
