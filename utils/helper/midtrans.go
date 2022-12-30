package helper

import (
	"os"
	// "time"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
)

func CreateInvoice(kodeDaftar string, biayaRegistrasi int64, metodePembayaran string) *coreapi.ChargeResponse {
	midtrans.ServerKey = os.Getenv("MIDTRANS_SERVER_KEY")
	midtrans.ClientKey = os.Getenv("MIDTRANS_CLIENT_KEY")
	midtrans.Environment = midtrans.Sandbox
	c := coreapi.Client{}
	c.New(os.Getenv("MIDTRANS_SERVER_KEY"), midtrans.Sandbox)

	switch {
	case metodePembayaran == "transfer_va_bca":
		chargeReq := &coreapi.ChargeReq{
			PaymentType: "bank_transfer",
			TransactionDetails: midtrans.TransactionDetails{
				OrderID:  kodeDaftar,
				GrossAmt: biayaRegistrasi,
			},
			BankTransfer: &coreapi.BankTransferDetails{
				Bank: midtrans.BankBca,
			},
			CustomExpiry: &coreapi.CustomExpiry{
				// OrderTime:      time.Now().Format("2016-12-07 11:54:12 +0700"),
				ExpiryDuration: 12,
				Unit:           "hour",
			},
		}
		coreApiRes, _ := coreapi.ChargeTransaction(chargeReq)
		return coreApiRes
	case metodePembayaran == "transfer_va_permata":
		chargeReq := &coreapi.ChargeReq{
			PaymentType: "bank_transfer",
			TransactionDetails: midtrans.TransactionDetails{
				OrderID:  kodeDaftar,
				GrossAmt: biayaRegistrasi,
			},
			BankTransfer: &coreapi.BankTransferDetails{
				Bank: midtrans.BankPermata,
			},
			CustomExpiry: &coreapi.CustomExpiry{
				// OrderTime:      time.Now().Format("2016-12-07 11:54:12 +0700"),
				ExpiryDuration: 12,
				Unit:           "hour",
			},
		}
		coreApiRes, _ := coreapi.ChargeTransaction(chargeReq)
		return coreApiRes
	case metodePembayaran == "transfer_va_bni":
		chargeReq := &coreapi.ChargeReq{
			PaymentType: "bank_transfer",
			TransactionDetails: midtrans.TransactionDetails{
				OrderID:  kodeDaftar,
				GrossAmt: biayaRegistrasi,
			},
			BankTransfer: &coreapi.BankTransferDetails{
				Bank: midtrans.BankBni,
			},
			CustomExpiry: &coreapi.CustomExpiry{
				// OrderTime:      time.Now().Format("2016-12-07 11:54:12 +0700"),
				ExpiryDuration: 12,
				Unit:           "hour",
			},
		}
		coreApiRes, _ := coreapi.ChargeTransaction(chargeReq)
		return coreApiRes
	case metodePembayaran == "transfer_va_bri":
		chargeReq := &coreapi.ChargeReq{
			PaymentType: "bank_transfer",
			TransactionDetails: midtrans.TransactionDetails{
				OrderID:  kodeDaftar,
				GrossAmt: biayaRegistrasi,
			},
			BankTransfer: &coreapi.BankTransferDetails{
				Bank: midtrans.BankBri,
			},
			CustomExpiry: &coreapi.CustomExpiry{
				// OrderTime:      time.Now().Format("2016-12-07 11:54:12 +0700"),
				ExpiryDuration: 12,
				Unit:           "hour",
			},
		}
		coreApiRes, _ := coreapi.ChargeTransaction(chargeReq)
		return coreApiRes
	case metodePembayaran == "qris":
		chargeReq := &coreapi.ChargeReq{
			PaymentType: "qris",
			TransactionDetails: midtrans.TransactionDetails{
				OrderID:  kodeDaftar,
				GrossAmt: biayaRegistrasi,
			},
			CustomExpiry: &coreapi.CustomExpiry{
				// OrderTime:      time.Now().Format("2016-12-07 11:54:12 +0700"),
				ExpiryDuration: 12,
				Unit:           "hour",
			},
		}
		coreApiRes, _ := coreapi.ChargeTransaction(chargeReq)
		return coreApiRes
	}
	return &coreapi.ChargeResponse{}
}

func UpdateMidtransPayment(kodeDaftar string) *coreapi.TransactionStatusResponse {
	midtrans.ServerKey = os.Getenv("MIDTRANS_SERVER_KEY")
	midtrans.ClientKey = os.Getenv("MIDTRANS_CLIENT_KEY")
	midtrans.Environment = midtrans.Sandbox
	c := coreapi.Client{}
	c.New(os.Getenv("MIDTRANS_SERVER_KEY"), midtrans.Sandbox)

	res, _ := c.CheckTransaction(kodeDaftar)
	return res
}
