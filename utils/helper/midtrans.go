package helper

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
)

func CreateInvoice(orderId string, price int64, payment_method string) *coreapi.ChargeResponse {
	if _, exist := os.LookupEnv("SECRET"); !exist {
		if err := godotenv.Load(".env"); err != nil {
			log.Fatal(err)
		}
	}

	midtrans.ServerKey = os.Getenv("MIDTRANS_SERVER_KEY")
	midtrans.Environment = midtrans.Sandbox

	switch {
	case payment_method == "transer_va_bca":
		chargeReq := &coreapi.ChargeReq{
			PaymentType: "bank_transer",
			TransactionDetails: midtrans.TransactionDetails{
				OrderID:  orderId,
				GrossAmt: price,
			},
			BankTransfer: &coreapi.BankTransferDetails{
				Bank: midtrans.BankBca,
			},
			CustomExpiry: &coreapi.CustomExpiry{
				OrderTime:      time.Now().Format(time.RFC822),
				ExpiryDuration: 12,
				Unit:           "hour",
			},
		}
		coreApiRes, _ := coreapi.ChargeTransaction(chargeReq)
		return coreApiRes
	case payment_method == "qris":
		chargeReq := &coreapi.ChargeReq{
			PaymentType: "qris",
			TransactionDetails: midtrans.TransactionDetails{
				OrderID:  orderId,
				GrossAmt: price,
			},
			CustomExpiry: &coreapi.CustomExpiry{
				OrderTime:      time.Now().Format(time.RFC822),
				ExpiryDuration: 12,
				Unit:           "hour",
			},
		}
		coreApiRes, _ := coreapi.ChargeTransaction(chargeReq)
		return coreApiRes
	}
	return &coreapi.ChargeResponse{}
}
