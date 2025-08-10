package payment

import (
	"database/sql"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/global"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/consts"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/database"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/vo"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/pkg/response"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/pkg/utils"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/pkg/utils/crypto"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/pkg/utils/ip"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/pkg/utils/payment"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/pkg/utils/sendto"
	utiltime "github.com/thanhoanganhtuan/DoAnChuyenNganh/pkg/utils/util_time"
	"go.uber.org/zap"
)

type serviceImpl struct {
	sqlc *database.Queries
	db   *sql.DB
}

func New(sqlc *database.Queries, db *sql.DB) Service {
	return &serviceImpl{
		sqlc: sqlc,
		db:   db,
	}
}

func (p *serviceImpl) PostRefund(ctx *gin.Context, in *vo.PostRefundInput) {
	panic("PostRefund")
	// loc, _ := time.LoadLocation("Asia/Ho_Chi_Minh")
	// now := time.Now().In(loc)

	// vnpRequestID := now.Format("150405")
	// vnpCreateDate := now.Format("20060102150405")
	// ipAddr := ip.GetClientIP(ctx)
	// vnpOrderInfo := "Hoan tien GD ma:" + in.OrderID
	// vnpTransactionNo := "0"

	// // Create data string for signature
	// data := strings.Join([]string{
	// 	vnpRequestID,
	// 	"2.1.0",
	// 	"refund",
	// 	global.Config.Payment.VnpTmnCode,
	// 	in.TransType,
	// 	in.OrderID,
	// 	strconv.Itoa(in.Amount * 100),
	// 	vnpTransactionNo,
	// 	in.TransDate,
	// 	in.User,
	// 	vnpCreateDate,
	// 	ipAddr,
	// 	vnpOrderInfo,
	// }, "|")

	// vnpSecureHash := crypto.CreateHMACSignature(data, global.Config.Payment.VnpHashSecret)

	// dataObj := vo.RefundDataObj{
	// 	VnpRequestID:       vnpRequestID,
	// 	VnpVersion:         "2.1.0",
	// 	VnpCommand:         "refund",
	// 	VnpTmnCode:         global.Config.Payment.VnpTmnCode,
	// 	VnpTransactionType: in.TransType,
	// 	VnpTxnRef:          in.OrderID,
	// 	VnpAmount:          in.Amount * 100,
	// 	VnpTransactionNo:   vnpTransactionNo,
	// 	VnpCreateBy:        in.User,
	// 	VnpOrderInfo:       vnpOrderInfo,
	// 	VnpTransactionDate: in.TransDate,
	// 	VnpCreateDate:      vnpCreateDate,
	// 	VnpIpAddr:          ipAddr,
	// 	VnpSecureHash:      vnpSecureHash,
	// }

	// // Make HTTP request to VNPay API
	// response, err := payment.MakeAPIRequest(global.Config.Payment.VnpApi, dataObj)
	// if err != nil {
	// 	log.Printf("Error processing refund: %v", err)
	// 	ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process refund"})
	// 	return
	// }

	// log.Printf("Refund response: %s", response)
	// ctx.JSON(http.StatusOK, gin.H{"message": "Refund request sent successfully"})
}

func (p *serviceImpl) PostQueryDR(ctx *gin.Context, in *vo.PostQueryDRInput) {
	panic("PostQueryDR")
	// loc, _ := time.LoadLocation("Asia/Ho_Chi_Minh")
	// now := time.Now().In(loc)

	// vnpRequestID := now.Format("150405")
	// vnpCreateDate := now.Format("20060102150405")
	// ipAddr := ip.GetClientIP(ctx)

	// vnpOrderInfo := "Truy van GD ma:" + in.OrderID

	// // Create data string for signature
	// data := strings.Join([]string{
	// 	vnpRequestID,
	// 	"2.1.0",
	// 	"querydr",
	// 	global.Config.Payment.VnpTmnCode,
	// 	in.OrderID,
	// 	in.TransDate,
	// 	vnpCreateDate,
	// 	ipAddr,
	// 	vnpOrderInfo,
	// }, "|")

	// vnpSecureHash := crypto.CreateHMACSignature(data, global.Config.Payment.VnpHashSecret)

	// dataObj := vo.QueryDataObj{
	// 	VnpRequestID:       vnpRequestID,
	// 	VnpVersion:         "2.1.0",
	// 	VnpCommand:         "querydr",
	// 	VnpTmnCode:         global.Config.Payment.VnpTmnCode,
	// 	VnpTxnRef:          in.OrderID,
	// 	VnpOrderInfo:       vnpOrderInfo,
	// 	VnpTransactionDate: in.TransDate,
	// 	VnpCreateDate:      vnpCreateDate,
	// 	VnpIpAddr:          ipAddr,
	// 	VnpSecureHash:      vnpSecureHash,
	// }

	// // Make HTTP request to VNPay API
	// response, err := payment.MakeAPIRequest(global.Config.Payment.VnpApi, dataObj)
	// if err != nil {
	// 	log.Printf("Error querying transaction: %v", err)
	// 	ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query transaction"})
	// 	return
	// }

	// log.Printf("Query response: %s", response)
	// ctx.JSON(http.StatusOK, gin.H{"message": "Query sent successfully"})
}

func (p *serviceImpl) VNPayIPN(ctx *gin.Context) {
	panic("VNPayIPN")
	// fmt.Print("VNPayIPN")
	// global.Logger.Info("VNPayIPN")

	// vnpParams := make(vo.VNPayParams)

	// // Get all query parameters
	// for key, values := range ctx.Request.URL.Query() {
	// 	if len(values) > 0 {
	// 		vnpParams[key] = values[0]
	// 	}
	// }

	// secureHash := vnpParams["vnp_SecureHash"]
	// orderID := vnpParams["vnp_TxnRef"]
	// rspCode := vnpParams["vnp_ResponseCode"]

	// // Remove hash fields for verification
	// delete(vnpParams, "vnp_SecureHash")
	// delete(vnpParams, "vnp_SecureHashType")

	// // Sort parameters and verify signature
	// sortedParams := payment.SortObject(vnpParams)
	// signData := payment.CreateQueryString(sortedParams)
	// signed := crypto.CreateHMACSignature(signData, global.Config.Payment.VnpHashSecret)

	// // Payment status simulation
	// paymentStatus := "0" // 0: Initial, 1: Success, 2: Failed
	// checkOrderID := true // Check if order exists in database
	// checkAmount := true  // Check if amount matches

	// if secureHash == signed {
	// 	if checkOrderID {
	// 		if checkAmount {
	// 			if paymentStatus == "0" {
	// 				if rspCode == "00" {
	// 					// Payment successful
	// 					// Update payment status to success in database
	// 					fmt.Printf("Payment successful for order: %s\n", orderID)
	// 					ctx.JSON(http.StatusOK, vo.VNPayResponse{
	// 						RspCode: "00",
	// 						Message: "Success",
	// 					})
	// 				} else {
	// 					// Payment failed
	// 					// Update payment status to failed in database
	// 					fmt.Printf("Payment failed for order: %s\n", orderID)
	// 					ctx.JSON(http.StatusOK, vo.VNPayResponse{
	// 						RspCode: "00",
	// 						Message: "Success",
	// 					})
	// 				}
	// 			} else {
	// 				ctx.JSON(http.StatusOK, vo.VNPayResponse{
	// 					RspCode: "02",
	// 					Message: "This order has been updated to the payment status",
	// 				})
	// 			}
	// 		} else {
	// 			ctx.JSON(http.StatusOK, vo.VNPayResponse{
	// 				RspCode: "04",
	// 				Message: "Amount invalid",
	// 			})
	// 		}
	// 	} else {
	// 		ctx.JSON(http.StatusOK, vo.VNPayResponse{
	// 			RspCode: "01",
	// 			Message: "Order not found",
	// 		})
	// 	}
	// } else {
	// 	ctx.JSON(http.StatusOK, vo.VNPayResponse{
	// 		RspCode: "97",
	// 		Message: "Checksum failed",
	// 	})
	// }
}

func (p *serviceImpl) VNPayReturn(ctx *gin.Context) (codeStatus int, err error) {
	vnpParams := make(vo.VNPayParams)

	// Get all query parameters
	for key, values := range ctx.Request.URL.Query() {
		if len(values) > 0 {
			vnpParams[key] = values[0]
		}
	}

	secureHash := vnpParams["vnp_SecureHash"]
	orderIDExternal := vnpParams["vnp_TxnRef"]
	responseCode := vnpParams["vnp_ResponseCode"]
	amount := vnpParams["vnp_Amount"]
	bankCode := vnpParams["vnp_BankCode"]
	transactionNo := vnpParams["vnp_TransactionNo"]
	payDate := vnpParams["vnp_PayDate"]

	// Remove hash fields for verification
	delete(vnpParams, "vnp_SecureHash")
	delete(vnpParams, "vnp_SecureHashType")

	// Sort parameters and verify signature
	sortedParams := payment.SortObject(vnpParams)
	signData := payment.CreateQueryString(sortedParams)
	signed := crypto.CreateHMACSignature(signData, global.Config.Payment.VnpHashSecret)

	if secureHash != signed {
		// Signature is valid - check with database and return result
		// code = vnpParams["vnp_ResponseCode"]
		// TODO: update order status
		now := utiltime.GetTimeNow()
		err = p.sqlc.UpdateOrderStatus(ctx, database.UpdateOrderStatusParams{
			OrderStatus:     database.EcommerceGoOrderOrderStatusPaymentFailed,
			UpdatedAt:       now,
			OrderIDExternal: orderIDExternal,
		})

		if err != nil {
			return response.ErrCodeInternalServerError, err
		}

		// TODO: redirect to frontend
		p.redirectToReactWithError(ctx, "INVALID SIGNATURE", "Không đúng", orderIDExternal)
		return
	}

	// Convert amount to VND
	amountInt, _ := strconv.Atoi(amount)
	// amountVND := amountInt / 100
	amountVND := decimal.NewFromInt(int64(amountInt)).Div(decimal.NewFromInt(100))

	// TODO: update order status
	now := utiltime.GetTimeNow()

	var (
		orderStatus   string
		paymentStatus string
	)

	if responseCode == "00" {
		orderStatus = string(database.EcommerceGoOrderOrderStatusPaymentSuccess)
		paymentStatus = string(database.EcommerceGoPaymentPaymentStatusSuccess)
	} else {
		orderStatus = string(database.EcommerceGoOrderOrderStatusPaymentFailed)
		paymentStatus = string(database.EcommerceGoPaymentPaymentStatusFailed)
	}

	err = p.sqlc.UpdateOrderStatus(ctx, database.UpdateOrderStatusParams{
		OrderStatus:     database.EcommerceGoOrderOrderStatus(orderStatus),
		UpdatedAt:       now,
		OrderIDExternal: orderIDExternal,
	})

	if err != nil {
		return response.ErrCodeInternalServerError, err
	}

	// TODO: get order id
	orderAndUserID, err := p.sqlc.GetOrderIdAndUserIdByOrderIdExternal(ctx, orderIDExternal)
	if err != nil {
		return response.ErrCodeInternalServerError, err
	}

	// TODO: Create payment
	paymentID := uuid.NewString()
	now = utiltime.GetTimeNow()
	err = p.sqlc.CreatePayment(ctx, database.CreatePaymentParams{
		ID:            paymentID,
		OrderID:       orderAndUserID.ID,
		PaymentStatus: database.EcommerceGoPaymentPaymentStatus(paymentStatus),
		PaymentMethod: database.EcommerceGoPaymentPaymentMethodCard,
		TotalPrice:    amountVND,
		TransactionID: sql.NullString{
			String: transactionNo,
			Valid:  true,
		},
		CreatedAt: now,
		UpdatedAt: now,
	})

	if err != nil {
		return response.ErrCodeInternalServerError, err
	}

	p.redirectToReactWithResult(ctx, vo.PaymentResultData{
		OrderIDExternal: orderIDExternal,
		ResponseCode:    responseCode,
		Amount:          amountVND.String(),
		BankCode:        bankCode,
		TransactionNo:   transactionNo,
		PayDate:         payDate,
	})

	formatted, err := utiltime.FormatVNPayTime(payDate)
	if err != nil {
		return response.ErrCodeInternalServerError, fmt.Errorf("format time from vnpay failed: %s", err)
	}

	// TODO: get email of user payment
	emailAndUsername, err := p.sqlc.GetEmailAndUsernameByID(ctx, orderAndUserID.UserID)
	if err != nil {
		return response.ErrCodeInternalServerError, fmt.Errorf("get email of user info failed: %s", err)
	}

	// TODO: send email
	err = sendto.SendEmail([]string{emailAndUsername.Account}, "payment_confirmation.html", map[string]interface{}{
		"username":     emailAndUsername.UserName,
		"orderId":      orderIDExternal,
		"amount":       amountVND,
		"orderDate":    formatted,
		"responseCode": responseCode,
	}, consts.PAYMENT_CONFIRMATION)
	if err != nil {
		return response.ErrCodeInternalServerError, fmt.Errorf("send email failed: %s", err)
	}

	return response.ErrCodeSuccessfully, nil
}

func (p *serviceImpl) CreatePaymentURL(ctx *gin.Context, in *vo.CreatePaymentURLInput) (codeStatus int, out *vo.CreatePaymentURLOutput, err error) {
	out = &vo.CreatePaymentURLOutput{}

	// TODO: get userId from context
	userID, ok := utils.GetUserIDFromGin(ctx)
	if !ok {
		return response.ErrCodeUnauthorized, nil, fmt.Errorf("userID not found in context")
	}

	// TODO: check user exists
	exists, err := p.sqlc.CheckUserBaseExistsById(ctx, userID)
	if err != nil {
		return response.ErrCodeInternalServerError, nil, fmt.Errorf("get user base failed: %s", err)
	}

	if !exists {
		return response.ErrCodeUnauthorized, nil, fmt.Errorf("user base not found")
	}

	// TODO: create payment url
	loc, _ := time.LoadLocation("Asia/Ho_Chi_Minh")
	now := time.Now().In(loc)

	createDate := now.Format("20060102150405")
	expireDate := now.Add(15 * time.Minute).Format("20060102150405")
	orderIDExternal := now.Format("02150405")

	ipAddr := ip.GetClientIP(ctx)

	layout := "02-01-2006"
	checkInDate, err1 := time.Parse(layout, in.CheckIn)
	checkOutDate, err2 := time.Parse(layout, in.CheckOut)
	if err1 != nil || err2 != nil {
		return response.ErrCodeInternalServerError, nil, err1
	}

	duration := checkOutDate.Sub(checkInDate)

	numDays := int64(duration.Hours() / 24)

	if numDays <= 0 {
		return response.ErrCodeInternalServerError, nil, fmt.Errorf("check_out must come after check_in")
	}

	// TODO: get total price
	var totalPrice decimal.Decimal
	totalPrice = decimal.Zero
	for _, roomSelected := range in.RoomSelected {
		accommodationDetail, err := p.sqlc.GetAccommodationDetail(ctx, database.GetAccommodationDetailParams{
			ID:              roomSelected.ID,
			AccommodationID: in.AccommodationID,
		})

		if err != nil {
			return response.ErrCodeInternalServerError, nil, err
		}
		quantity := decimal.NewFromInt(int64(roomSelected.Quantity))
		roomSubtotal := accommodationDetail.Price.Mul(quantity)
		totalPrice = totalPrice.Add(roomSubtotal)
	}

	numDaysDecimal := decimal.NewFromInt(numDays)

	totalPrice = totalPrice.Mul(numDaysDecimal)

	vnpParams := map[string]string{
		"vnp_Version":    "2.1.0",
		"vnp_Command":    "pay",
		"vnp_TmnCode":    global.Config.Payment.VnpTmnCode,
		"vnp_Locale":     "vn",
		"vnp_CurrCode":   "VND",
		"vnp_TxnRef":     orderIDExternal,
		"vnp_OrderInfo":  "Thanh toan cho ma GD:" + orderIDExternal,
		"vnp_OrderType":  "other",
		"vnp_Amount":     totalPrice.Mul(decimal.NewFromInt(100)).String(),
		"vnp_ReturnUrl":  global.Config.Payment.VnpReturnUrl,
		"vnp_IpAddr":     ipAddr,
		"vnp_CreateDate": createDate,
		"vnp_ExpireDate": expireDate,
	}

	sortedParams := payment.SortObject(vnpParams)

	signData := payment.CreateSignData(sortedParams)

	signature := crypto.CreateHMACSignature(signData, global.Config.Payment.VnpHashSecret)

	sortedParams["vnp_SecureHash"] = signature

	finalURL := global.Config.Payment.VnpUrl + "?" + payment.CreateQueryString(sortedParams)

	// TODO: save order to database
	orderID := uuid.NewString()

	checkIn, err := utiltime.ConvertISOToUnixTimestamp(in.CheckIn)
	if err != nil {
		return response.ErrCodeInternalServerError, nil, fmt.Errorf("convert ISO to Unix failed: %s", err)
	}

	checkOut, err := utiltime.ConvertISOToUnixTimestamp(in.CheckOut)
	if err != nil {
		return response.ErrCodeInternalServerError, nil, fmt.Errorf("convert ISO to Unix failed: %s", err)
	}

	createdAt := utiltime.GetTimeNow()
	// createdAt := uint64(time.Date(2025, time.June, 3, 0, 0, 0, 0, time.UTC).UnixMilli())

	// TODO: create order
	err = p.sqlc.CreateOrder(ctx, database.CreateOrderParams{
		ID:              orderID,
		UserID:          userID,
		FinalTotal:      totalPrice,
		OrderStatus:     database.EcommerceGoOrderOrderStatusPendingPayment,
		AccommodationID: in.AccommodationID,
		OrderIDExternal: orderIDExternal,
		VoucherID: sql.NullString{
			String: "",
			Valid:  false,
		},
		CheckinDate:  checkIn,
		CheckoutDate: checkOut,
		CreatedAt:    createdAt,
		UpdatedAt:    createdAt,
	})

	if err != nil {
		return response.ErrCodeInternalServerError, nil, err
	}

	// TODO: lấy thông tin của accommodation detail
	for _, roomSelected := range in.RoomSelected {
		accommodationDetail, err := p.sqlc.GetAccommodationDetail(ctx, database.GetAccommodationDetailParams{
			ID:              roomSelected.ID,
			AccommodationID: in.AccommodationID,
		})
		if err != nil {
			return response.ErrCodeInternalServerError, nil, err
		}

		// TODO: get available rooms by quantity
		availableRooms, err := p.sqlc.GetAccommodationRoomsAvailableByQuantity(ctx, database.GetAccommodationRoomsAvailableByQuantityParams{
			CheckOut:            checkOut,
			CheckIn:             checkIn,
			AccommodationTypeID: roomSelected.ID,
			Limit:               int32(roomSelected.Quantity),
		})

		if err != nil {
			return response.ErrCodeInternalServerError, nil, fmt.Errorf("get accommodation rooms failed: %s", err)
		}

		if len(availableRooms) < int(roomSelected.Quantity) {
			return response.ErrCodeAccommodationRoomNotFound, nil, fmt.Errorf("not enough available rooms: need %d, found %d", roomSelected.Quantity, len(availableRooms))
		}

		orderDetailID := uuid.NewString()
		err = p.sqlc.CreateOrderDetail(ctx, database.CreateOrderDetailParams{
			ID:                    orderDetailID,
			OrderID:               orderID,
			Quantity:              roomSelected.Quantity,
			Price:                 accommodationDetail.Price.Mul(decimal.NewFromInt(int64(roomSelected.Quantity))),
			AccommodationDetailID: accommodationDetail.ID,
			CreatedAt:             createdAt,
			UpdatedAt:             createdAt,
		})

		if err != nil {
			return response.ErrCodeInternalServerError, nil, err
		}

		// TODO: create order room booking for each available room
		for _, availableRoom := range availableRooms {
			orderRoomBookingID := uuid.NewString()
			err = p.sqlc.CreateOrderRoomBooking(ctx, database.CreateOrderRoomBookingParams{
				ID:                  orderRoomBookingID,
				OrderDetailID:       orderDetailID,
				AccommodationRoomID: availableRoom.ID,
				BookingStatus:       database.EcommerceGoOrderRoomBookingBookingStatusReserved,
				CreatedAt:           createdAt,
				UpdatedAt:           createdAt,
			})

			if err != nil {
				return response.ErrCodeInternalServerError, nil, fmt.Errorf("create order room booking failed: %s", err)
			}
		}
	}

	out.Url = finalURL
	return response.ErrCodeSuccessfully, out, nil
}

func (p *serviceImpl) redirectToReactWithResult(ctx *gin.Context, data vo.PaymentResultData) {
	params := url.Values{}
	params.Set("order_id", data.OrderIDExternal)
	params.Set("response_code", data.ResponseCode)
	params.Set("amount", data.Amount)

	if data.BankCode != "" {
		params.Set("bank_code", data.BankCode)
	}

	if data.PayDate != "" {
		params.Set("pay_date", data.PayDate)
	}

	if data.TransactionNo != "" {
		params.Set("transaction_no", data.TransactionNo)
	}

	// TODO: build react frontend url
	frontendURL := fmt.Sprintf("%s/payment/result?%s",
		global.Config.Frontend.Url, params.Encode())

	fmt.Printf("redirectToReactWithResult success: %s\n", frontendURL)
	global.Logger.Info("redirectToReactWithResult success: ", zap.String("info", frontendURL))

	ctx.Redirect(http.StatusFound, frontendURL)
}

func (p *serviceImpl) redirectToReactWithError(ctx *gin.Context, errorCode, message, orderID string) {
	params := url.Values{}
	params.Set("status", "error")
	params.Set("error_code", errorCode)
	params.Set("message", message)

	if orderID != "" {
		params.Set("order_id", orderID)
	}

	// TODO: build react frontend url
	frontendURL := fmt.Sprintf("%s/payment/result?%s",
		global.Config.Frontend.Url, params.Encode())

	fmt.Printf("redirectToReactWithError success: %s\n", frontendURL)
	global.Logger.Info("redirectToReactWithError success: ", zap.String("info", frontendURL))
	ctx.Redirect(http.StatusFound, frontendURL)
}
