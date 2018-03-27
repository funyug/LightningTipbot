package lnd

import (
	"google.golang.org/grpc/credentials"
	"gopkg.in/macaroon.v2"
	"github.com/lightningnetwork/lnd/macaroons"
	"google.golang.org/grpc"
	"path/filepath"
	"github.com/btcsuite/btcutil"
	"log"
	"os"
	"io/ioutil"
	"github.com/lightningnetwork/lnd/lnrpc"
	"context"
)

const (
	defaultTLSCertFilename  = "tls.cert"
	defaultMacaroonFilename = "admin.macaroon"
)

var (
	defaultLndDir       = btcutil.AppDataDir("lnd", false)
	defaultTLSCertPath  = filepath.Join(defaultLndDir, defaultTLSCertFilename)
	defaultMacaroonPath = filepath.Join(defaultLndDir, defaultMacaroonFilename)
)

var Client lnrpc.LightningClient

func fatal(err error) {
	log.Print(os.Stderr, "[lncli] %v\n", err)
	os.Exit(1)
}

func Connect(conn *grpc.ClientConn) *grpc.ClientConn {
	creds, err := credentials.NewClientTLSFromFile(defaultTLSCertPath, "")
	if err != nil {
		fatal(err)
	}

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(creds),
	}

	macBytes, err := ioutil.ReadFile(defaultMacaroonPath)
	if err != nil {
		fatal(err)
	}
	mac := &macaroon.Macaroon{}
	if err = mac.UnmarshalBinary(macBytes); err != nil {
		fatal(err)
	}

	// Now we append the macaroon credentials to the dial options.
	cred := macaroons.NewMacaroonCredential(mac)
	opts = append(opts, grpc.WithPerRPCCredentials(cred))

	conn, err = grpc.Dial("localhost:10009", opts...)
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	return conn
}

func AddInvoice(amount int64) (*lnrpc.AddInvoiceResponse,error) {
	invoice := &lnrpc.Invoice{
		Value: amount,
	}
	response, err := Client.AddInvoice(context.Background(),invoice);
	if err != nil {
		log.Println(err)
		return nil,err
	}
	log.Println(response)
	return response, err
}

func DecodePayReq(payreq string) (*lnrpc.PayReq,error) {
	req := &lnrpc.PayReqString{
		PayReq: payreq,
	}
	response, err := Client.DecodePayReq(context.Background(),req);
	if err != nil {
		log.Println(err)
		return nil,err
	}
	log.Println(response)
	return response, err
}


func SendPaymentSync(payreq string) (*lnrpc.SendResponse,error) {
	req := &lnrpc.SendRequest{
		PaymentRequest: payreq,
	}
	response, err := Client.SendPaymentSync(context.Background(),req);
	if err != nil {
		log.Println(err)
		return nil,err
	}
	log.Println(response)
	return response, err
}