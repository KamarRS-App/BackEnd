package helper

import (
	"os"
	"time"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
)

func CreateInvoice(kodeDaftar string, biayaRegistrasi int64, metodePembayaran string) *coreapi.ChargeResponse {
	midtrans.ServerKey = os.Getenv("MIDTRANS_SERVER_KEY")
	midtrans.Environment = midtrans.Sandbox

	switch {
	case metodePembayaran == "transer_va_bca":
		chargeReq := &coreapi.ChargeReq{
			PaymentType: "bank_transer",
			TransactionDetails: midtrans.TransactionDetails{
				OrderID:  kodeDaftar,
				GrossAmt: biayaRegistrasi,
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
	case metodePembayaran == "transer_va_permata":
		chargeReq := &coreapi.ChargeReq{
			PaymentType: "bank_transer",
			TransactionDetails: midtrans.TransactionDetails{
				OrderID:  kodeDaftar,
				GrossAmt: biayaRegistrasi,
			},
			BankTransfer: &coreapi.BankTransferDetails{
				Bank: midtrans.BankPermata,
			},
			CustomExpiry: &coreapi.CustomExpiry{
				OrderTime:      time.Now().Format(time.RFC822),
				ExpiryDuration: 12,
				Unit:           "hour",
			},
		}
		coreApiRes, _ := coreapi.ChargeTransaction(chargeReq)
		return coreApiRes
	case metodePembayaran == "transer_va_bni":
		chargeReq := &coreapi.ChargeReq{
			PaymentType: "bank_transer",
			TransactionDetails: midtrans.TransactionDetails{
				OrderID:  kodeDaftar,
				GrossAmt: biayaRegistrasi,
			},
			BankTransfer: &coreapi.BankTransferDetails{
				Bank: midtrans.BankBni,
			},
			CustomExpiry: &coreapi.CustomExpiry{
				OrderTime:      time.Now().Format(time.RFC822),
				ExpiryDuration: 12,
				Unit:           "hour",
			},
		}
		coreApiRes, _ := coreapi.ChargeTransaction(chargeReq)
		return coreApiRes
	case metodePembayaran == "transer_va_bri":
		chargeReq := &coreapi.ChargeReq{
			PaymentType: "bank_transer",
			TransactionDetails: midtrans.TransactionDetails{
				OrderID:  kodeDaftar,
				GrossAmt: biayaRegistrasi,
			},
			BankTransfer: &coreapi.BankTransferDetails{
				Bank: midtrans.BankBri,
			},
			CustomExpiry: &coreapi.CustomExpiry{
				OrderTime:      time.Now().Format(time.RFC822),
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

func UpdateMidtransPayment(kodeDaftar string) *coreapi.TransactionStatusResponse {
	midtrans.ServerKey = os.Getenv("MIDTRANS_SERVER_KEY")
	midtrans.Environment = midtrans.Sandbox

	res, _ := coreapi.CheckTransaction(kodeDaftar)
	return res
}
