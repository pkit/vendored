// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package appstream_test

import (
	"bytes"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/appstream"
)

var _ time.Duration
var _ bytes.Buffer

func ExampleAppStream_AssociateFleet() {
	sess := session.Must(session.NewSession())

	svc := appstream.New(sess)

	params := &appstream.AssociateFleetInput{
		FleetName: aws.String("String"), // Required
		StackName: aws.String("String"), // Required
	}
	resp, err := svc.AssociateFleet(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleAppStream_CreateFleet() {
	sess := session.Must(session.NewSession())

	svc := appstream.New(sess)

	params := &appstream.CreateFleetInput{
		ComputeCapacity: &appstream.ComputeCapacity{ // Required
			DesiredInstances: aws.Int64(1), // Required
		},
		ImageName:                   aws.String("String"), // Required
		InstanceType:                aws.String("String"), // Required
		Name:                        aws.String("Name"),   // Required
		Description:                 aws.String("Description"),
		DisconnectTimeoutInSeconds:  aws.Int64(1),
		DisplayName:                 aws.String("DisplayName"),
		EnableDefaultInternetAccess: aws.Bool(true),
		MaxUserDurationInSeconds:    aws.Int64(1),
		VpcConfig: &appstream.VpcConfig{
			SubnetIds: []*string{
				aws.String("String"), // Required
				// More values...
			},
		},
	}
	resp, err := svc.CreateFleet(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleAppStream_CreateStack() {
	sess := session.Must(session.NewSession())

	svc := appstream.New(sess)

	params := &appstream.CreateStackInput{
		Name:        aws.String("String"), // Required
		Description: aws.String("Description"),
		DisplayName: aws.String("DisplayName"),
		StorageConnectors: []*appstream.StorageConnector{
			{ // Required
				ConnectorType:      aws.String("StorageConnectorType"), // Required
				ResourceIdentifier: aws.String("ResourceIdentifier"),
			},
			// More values...
		},
	}
	resp, err := svc.CreateStack(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleAppStream_CreateStreamingURL() {
	sess := session.Must(session.NewSession())

	svc := appstream.New(sess)

	params := &appstream.CreateStreamingURLInput{
		FleetName:      aws.String("String"), // Required
		StackName:      aws.String("String"), // Required
		UserId:         aws.String("UserId"), // Required
		ApplicationId:  aws.String("String"),
		SessionContext: aws.String("String"),
		Validity:       aws.Int64(1),
	}
	resp, err := svc.CreateStreamingURL(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleAppStream_DeleteFleet() {
	sess := session.Must(session.NewSession())

	svc := appstream.New(sess)

	params := &appstream.DeleteFleetInput{
		Name: aws.String("String"), // Required
	}
	resp, err := svc.DeleteFleet(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleAppStream_DeleteStack() {
	sess := session.Must(session.NewSession())

	svc := appstream.New(sess)

	params := &appstream.DeleteStackInput{
		Name: aws.String("String"), // Required
	}
	resp, err := svc.DeleteStack(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleAppStream_DescribeFleets() {
	sess := session.Must(session.NewSession())

	svc := appstream.New(sess)

	params := &appstream.DescribeFleetsInput{
		Names: []*string{
			aws.String("String"), // Required
			// More values...
		},
		NextToken: aws.String("String"),
	}
	resp, err := svc.DescribeFleets(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleAppStream_DescribeImages() {
	sess := session.Must(session.NewSession())

	svc := appstream.New(sess)

	params := &appstream.DescribeImagesInput{
		Names: []*string{
			aws.String("String"), // Required
			// More values...
		},
	}
	resp, err := svc.DescribeImages(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleAppStream_DescribeSessions() {
	sess := session.Must(session.NewSession())

	svc := appstream.New(sess)

	params := &appstream.DescribeSessionsInput{
		FleetName:          aws.String("String"), // Required
		StackName:          aws.String("String"), // Required
		AuthenticationType: aws.String("AuthenticationType"),
		Limit:              aws.Int64(1),
		NextToken:          aws.String("String"),
		UserId:             aws.String("UserId"),
	}
	resp, err := svc.DescribeSessions(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleAppStream_DescribeStacks() {
	sess := session.Must(session.NewSession())

	svc := appstream.New(sess)

	params := &appstream.DescribeStacksInput{
		Names: []*string{
			aws.String("String"), // Required
			// More values...
		},
		NextToken: aws.String("String"),
	}
	resp, err := svc.DescribeStacks(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleAppStream_DisassociateFleet() {
	sess := session.Must(session.NewSession())

	svc := appstream.New(sess)

	params := &appstream.DisassociateFleetInput{
		FleetName: aws.String("String"), // Required
		StackName: aws.String("String"), // Required
	}
	resp, err := svc.DisassociateFleet(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleAppStream_ExpireSession() {
	sess := session.Must(session.NewSession())

	svc := appstream.New(sess)

	params := &appstream.ExpireSessionInput{
		SessionId: aws.String("String"), // Required
	}
	resp, err := svc.ExpireSession(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleAppStream_ListAssociatedFleets() {
	sess := session.Must(session.NewSession())

	svc := appstream.New(sess)

	params := &appstream.ListAssociatedFleetsInput{
		StackName: aws.String("String"), // Required
		NextToken: aws.String("String"),
	}
	resp, err := svc.ListAssociatedFleets(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleAppStream_ListAssociatedStacks() {
	sess := session.Must(session.NewSession())

	svc := appstream.New(sess)

	params := &appstream.ListAssociatedStacksInput{
		FleetName: aws.String("String"), // Required
		NextToken: aws.String("String"),
	}
	resp, err := svc.ListAssociatedStacks(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleAppStream_StartFleet() {
	sess := session.Must(session.NewSession())

	svc := appstream.New(sess)

	params := &appstream.StartFleetInput{
		Name: aws.String("String"), // Required
	}
	resp, err := svc.StartFleet(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleAppStream_StopFleet() {
	sess := session.Must(session.NewSession())

	svc := appstream.New(sess)

	params := &appstream.StopFleetInput{
		Name: aws.String("String"), // Required
	}
	resp, err := svc.StopFleet(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleAppStream_UpdateFleet() {
	sess := session.Must(session.NewSession())

	svc := appstream.New(sess)

	params := &appstream.UpdateFleetInput{
		Name: aws.String("String"), // Required
		ComputeCapacity: &appstream.ComputeCapacity{
			DesiredInstances: aws.Int64(1), // Required
		},
		DeleteVpcConfig:             aws.Bool(true),
		Description:                 aws.String("Description"),
		DisconnectTimeoutInSeconds:  aws.Int64(1),
		DisplayName:                 aws.String("DisplayName"),
		EnableDefaultInternetAccess: aws.Bool(true),
		ImageName:                   aws.String("String"),
		InstanceType:                aws.String("String"),
		MaxUserDurationInSeconds:    aws.Int64(1),
		VpcConfig: &appstream.VpcConfig{
			SubnetIds: []*string{
				aws.String("String"), // Required
				// More values...
			},
		},
	}
	resp, err := svc.UpdateFleet(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleAppStream_UpdateStack() {
	sess := session.Must(session.NewSession())

	svc := appstream.New(sess)

	params := &appstream.UpdateStackInput{
		Name: aws.String("String"), // Required
		DeleteStorageConnectors: aws.Bool(true),
		Description:             aws.String("Description"),
		DisplayName:             aws.String("DisplayName"),
		StorageConnectors: []*appstream.StorageConnector{
			{ // Required
				ConnectorType:      aws.String("StorageConnectorType"), // Required
				ResourceIdentifier: aws.String("ResourceIdentifier"),
			},
			// More values...
		},
	}
	resp, err := svc.UpdateStack(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}
